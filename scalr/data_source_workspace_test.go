package scalr

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccTFEWorkspaceDataSource_basic(t *testing.T) {
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTFEWorkspaceDataSourceConfig(rInt),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.scalr_workspace.foobar",
						"id",
						fmt.Sprintf("existing-org/workspace-test-%d", rInt),
					),
					resource.TestCheckResourceAttr(
						"data.scalr_workspace.foobar", "name", fmt.Sprintf("workspace-test-%d", rInt)),
					resource.TestCheckResourceAttr(
						"data.scalr_workspace.foobar", "organization", "existing-org"),
					resource.TestCheckResourceAttr(
						"data.scalr_workspace.foobar", "auto_apply", "true"),
					resource.TestCheckResourceAttr(
						"data.scalr_workspace.foobar", "queue_all_runs", "false"),
					resource.TestCheckResourceAttr(
						"data.scalr_workspace.foobar", "terraform_version", "0.12.19"),
					resource.TestCheckResourceAttr(
						"data.scalr_workspace.foobar", "working_directory", "terraform/test"),

					resource.TestCheckResourceAttrSet("data.scalr_workspace.foobar", "external_id"),
				),
			},
		},
	})
}

func testAccTFEWorkspaceDataSourceConfig(rInt int) string {
	return fmt.Sprintf(`
resource "scalr_workspace" "foobar" {
  name                  = "workspace-test-%d"
  organization          = "existing-org"
  auto_apply            = true
  queue_all_runs        = false
  terraform_version     = "0.12.19"
  working_directory     = "terraform/test"
}

data "scalr_workspace" "foobar" {
  name         = "${scalr_workspace.foobar.name}"
  organization = "existing-org"
}`, rInt)
}
