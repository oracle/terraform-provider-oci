package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityUserTestSuite struct {
	suite.Suite
	Client       *MockClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
	User         *baremtlclient.Resource
}

func (s *ResourceIdentityUserTestSuite) SetupTest() {
	s.Client = &MockClient{}
	s.Provider = Provider(s.Client)
	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = `
		resource "baremetal_identity_user" "t" {
			name = "name!"
			description = "desc!"
		}
	`
	s.ResourceName = "baremetal_identity_user.t"
	s.User = &baremtlclient.Resource{
		ID:            "id!",
		Name:          "name!",
		Description:   "desc!",
		CompartmentID: "cid!",
		State:         "state!",
		TimeCreated:   s.TimeCreated,
		TimeModified:  s.TimeCreated,
	}
}

func (s *ResourceIdentityUserTestSuite) TestCreateResourceIdentityUser() {
	s.Client.On("CreateUser", "name!", "desc!").Return(s.User, nil)
	s.Client.On("GetUser", "id!").Return(s.User, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", s.User.Name),
					resource.TestCheckResourceAttr(s.ResourceName, "description", s.User.Description),
					resource.TestCheckResourceAttr(s.ResourceName, "compartment_id", s.User.CompartmentID),
					resource.TestCheckResourceAttr(s.ResourceName, "state", s.User.State),
					resource.TestCheckResourceAttr(s.ResourceName, "time_created", s.User.TimeCreated.String()),
					resource.TestCheckResourceAttr(s.ResourceName, "time_modified", s.User.TimeModified.String()),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestUpdateResourceIdentityUser() {
	s.Client.On("CreateUser", "name!", "desc!").Return(s.User, nil)
	s.Client.On("GetUser", "id!").Return(s.User, nil)

	updatedesc := `
		resource "baremetal_identity_user" "t" {
			name = "name!"
			description = "newdesc!"
		}
	`
	s.Client.On("UpdateUser", "newdesc!")

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: updatedesc,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
				),
			},
		},
	})
}

func TestResourceIdentityUserTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserTestSuite))
}
