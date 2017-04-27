package main

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/MustWin/baremetal-sdk-go"
)

func TestAccOBMASVolume_basic(t *testing.T) {
	ri := acctest.RandInt()
	config := testAccOBMASVolume_basic(ri)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckOBMASVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckOBMASVolumeExists("baremetal_core_volume.test"),
				),
			},
		},
	})
}

func testCheckOBMASVolumeExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		name := rs.Primary.ID
		client := testAccProvider.Meta().(*baremetal.Client)
		resp, err := client.GetVolume(name)
		if err != nil {
			return fmt.Errorf("Bad: Error retrieving the Volume: %s", err)
		}

		if resp == nil {
			return fmt.Errorf("Bad: Unable to find Volume: %s", err)
		}

		return nil
	}
}

func testCheckOBMASVolumeDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*baremetal.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "baremetal_core_volume" {
			continue
		}

		name := rs.Primary.ID

		resp, err := client.GetVolume(name)
		if err != nil {
			return nil
		}

		if resp == nil {
			return fmt.Errorf("Volume still exists:\n%#v", resp)
		}
	}

	return nil
}

func testAccOBMASVolume_basic(rInt int) string{
	return fmt.Sprintf(`
resource "baremetal_core_volume" "test" {
  availability_domain = "Uocm:PHX-AD-1"
  compartment_id      = "XXXXX"
  size_in_mbs         = "1024"
}
`)
}