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
					testAccIdentityUserCreated(
						"test_user"),
				),
			},
		},
	})
}

func testAccIdentityUserCreated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		/*
			_, ok := s.RootModule().Resources[n]
				if !ok {
					return fmt.Errorf("Not implemented %s", n)
				}
		*/

		return fmt.Errorf("Not implemented.")

	}
}

var testAccIdentityUser = fmt.Sprintf(`
resource "baremetal_identity_user" "test_user" {
	name = "test_user"
	description = "A test user"
	compartment_id = "TBD.TBD.TBD"
}
`)
