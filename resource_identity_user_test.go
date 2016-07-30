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
	u := &baremtlclient.Resource{
		ID:            "id!",
		Name:          "name!",
		Description:   "desc!",
		CompartmentID: "cid!",
		State:         "state!",
		TimeModified:  t,
		TimeCreated:   t,
	}
	s.Client.On("CreateUser", "name!", "desc!").Return(u, nil)
	s.Client.On("GetUser", "id!").Return(u, nil)

	rname := "baremetal_identity_user.t"
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rname, "name", u.Name),
					resource.TestCheckResourceAttr(rname, "description", u.Description),
					resource.TestCheckResourceAttr(rname, "compartment_id", u.CompartmentID),
					resource.TestCheckResourceAttr(rname, "state", u.State),
					resource.TestCheckResourceAttr(rname, "time_modified", u.TimeModified.String()),
					resource.TestCheckResourceAttr(rname, "time_created", u.TimeCreated.String()),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestUpdateResourceIdentityUser() {
	t, _ := time.Parse("2006-Jan-02", "2006-Jan-02")
	c1 := `
		resource "baremetal_identity_user" "t" {
			name = "name!"
			description = "desc!"
		}
	`
	u1 := &baremtlclient.Resource{
		ID:            "id!",
		Name:          "name!",
		Description:   "desc!",
		CompartmentID: "cid!",
		State:         "state!",
		TimeModified:  t,
		TimeCreated:   t,
	}
	s.Client.On("CreateUser", "name!", "desc!").Return(u1, nil)
	s.Client.On("GetUser", "id!").Return(u1, nil)

	c2 := `
		resource "baremetal_identity_user" "t" {
			name = "name!"
			description = "newdesc!"
		}
	`
	s.Client.On("UpdateUser", "newdesc!")

	rname := "baremetal_identity_user.t"
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: c1,
			},
			resource.TestStep{
				Config: c2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rname, "description", "newdesc!"),
				),
			},
		},
	})
}

func TestResourceIdentityUserTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserTestSuite))
}
