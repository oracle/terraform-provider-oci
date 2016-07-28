package main

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUserTestSuite struct {
	suite.Suite
	Client    *MockClient
	Provider  terraform.ResourceProvider
	Providers map[string]terraform.ResourceProvider
}

func (s *ResourceIdentityUserTestSuite) SetupTest() {
	s.Client = &MockClient{}
	s.Provider = Provider(s.Client)
	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
}

func (s *ResourceIdentityUserTestSuite) TestCreateResourceIdentityUser() {
	config := `
		resource "baremetal_identity_user" "test" {
			name = "name!"
			description = "desc!"
			compartment_id = "compartment_id!"
		}
	`
	s.Client.On("CreateUser", "name!", "desc!").Return("id!", nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					s.testPresentInStateAfterCreate,
					s.testIdSetAfterCreate,
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) testPresentInStateAfterCreate(state *terraform.State) error {
	_, ok := state.RootModule().Resources["baremetal_identity_user.test"]
	if !ok {
		return fmt.Errorf("Resource not found.")
	}

	return nil
}

func (s *ResourceIdentityUserTestSuite) testIdSetAfterCreate(state *terraform.State) error {
	rs, _ := state.RootModule().Resources["baremetal_identity_user.test"]
	s.Equal("id!", rs.Primary.ID)

	return nil
}

func TestResourceIdentityUserTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserTestSuite))
}
