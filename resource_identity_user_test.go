package main

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIdentityUserCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIdentityUser,
				Check: resource.ComposeTestCheckFunc(
					testAccIdentityUserCreated("baremetal_identity_user.users"),
				),
			},
		},
	})
}

func testAccIdentityUserCreated(resourceID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceID]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceID)
		}

		if rs.Primary.ID != "SOME_USER_ID" {
			return fmt.Errorf("Unexpected user_id: %v", rs.Primary.ID)
		}
		return nil
	}
}

var testAccIdentityUser = fmt.Sprintf(`
resource "baremetal_identity_user" "users" {
	name = "test_user"
	description = "A test user"
	compartment_id = "TBD.TBD.TBD"
}
`)
