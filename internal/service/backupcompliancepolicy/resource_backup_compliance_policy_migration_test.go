package backupcompliancepolicy_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/mongodb/terraform-provider-mongodbatlas/internal/testutil/mig"
)

func TestAccMigrationGenericBackupRSBackupCompliancePolicy_basic(t *testing.T) {
	var (
		projectName    = fmt.Sprintf("testacc-project-%s", acctest.RandString(10))
		orgID          = os.Getenv("MONGODB_ATLAS_ORG_ID")
		projectOwnerID = os.Getenv("MONGODB_ATLAS_PROJECT_OWNER_ID")
		config         = configBasic(projectName, orgID, projectOwnerID)
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { mig.PreCheckBasic(t) },
		CheckDestroy: checkDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: mig.ExternalProviders(),
				Config:            config,
				Check: resource.ComposeTestCheckFunc(
					checkExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "copy_protection_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "encryption_at_rest_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "authorized_user_first_name", "First"),
					resource.TestCheckResourceAttr(resourceName, "authorized_user_last_name", "Last"),
					resource.TestCheckResourceAttr(resourceName, "restore_window_days", "7"),
				),
			},
			mig.TestStep(config),
		},
	})
}
