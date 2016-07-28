package main

import (
	"testing"
	"time"

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
		resource "baremetal_identity_user" "t" {
			name = "name!"
			description = "desc!"
		}
	`
	t, _ := time.Parse("2006-Jan-02", "2006-Jan-02")
	user := &BareMetalIdentity{
		ID:            "id!",
		Name:          "name!",
		Description:   "desc!",
		CompartmentID: "cid!",
		State:         "state!",
		TimeModified:  t,
		TimeCreated:   t,
	}
	s.Client.On("CreateUser", "name!", "desc!").Return(user, nil)

	rname := "baremetal_identity_user.t"
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rname, "name", "name!"),
					resource.TestCheckResourceAttr(rname, "description", "desc!"),
					resource.TestCheckResourceAttr(rname, "compartment_id", "cid!"),
					resource.TestCheckResourceAttr(rname, "state", "state!"),
					resource.TestCheckResourceAttr(rname, "time_modified", t.String()),
					resource.TestCheckResourceAttr(rname, "time_created", t.String()),
				),
			},
		},
	})
}

func TestResourceIdentityUserTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserTestSuite))
}
