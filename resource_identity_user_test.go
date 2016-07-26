package main

import (
	"crypto/rsa"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// MockClient is used to access Oracle BareMetal Services during tests
type MockClient struct {
}

// New creates a new BareMetalClient instance
func (mock MockClient) New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey) BareMetalClient {
	return nil
}

// UserCreate method to create an user
func (mock MockClient) UserCreate(name, description string) (id string, err error) {
	// TODO: return a random string
	return "SOME_USER_ID", nil
}

func TestAccIdentityUserCreate(t *testing.T) {
	var testAccProviders map[string]terraform.ResourceProvider
	var testAccProvider *schema.Provider

	client := MockClient{}

	testAccProvider = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"baremetal_identity_user": ResourceIdentityUser(client),
		},
	}

	testAccProviders = map[string]terraform.ResourceProvider{
		"baremetal": testAccProvider,
	}

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
