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
	User         *baremtlsdk.Resource
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
	s.User = &baremtlsdk.Resource{
		ID:            "id!",
		Name:          "name!",
		Description:   "desc!",
		CompartmentID: "cid!",
		State:         baremtlsdk.ResourceCreated,
		TimeCreated:   s.TimeCreated,
		TimeModified:  s.TimeCreated,
	}
}

func (s *ResourceIdentityUserTestSuite) TestCreateResourceIdentityUser() {
	s.Client.On("CreateUser", "name!", "desc!").Return(s.User, nil)
	s.Client.On("GetUser", "id!").Return(s.User, nil)
	s.Client.On("DeleteUser", "id!").Return(nil)

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

func (s *ResourceIdentityUserTestSuite) TestCreateResourceIdentityUserPolling() {
	s.User.State = baremtlsdk.ResourceCreating
	s.Client.On("CreateUser", "name!", "desc!").Return(s.User, nil)
	s.Client.On("GetUser", "id!").Return(s.User, nil).Once()

	u := *s.User
	u.State = baremtlsdk.ResourceCreated
	s.Client.On("GetUser", "id!").Return(&u, nil)

	s.Client.On("DeleteUser", "id!").Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremtlsdk.ResourceCreated),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestUpdateResourceIdentityUserDescription() {
	s.Client.On("CreateUser", "name!", "desc!").Return(s.User, nil)
	s.Client.On("GetUser", "id!").Return(s.User, nil)
	s.Client.On("DeleteUser", "id!").Return(nil)

	c := `
		resource "baremetal_identity_user" "t" {
			name = "name!"
			description = "newdesc!"
		}
	`
	t := s.TimeCreated.Add(5 * time.Minute)
	u := &baremtlsdk.Resource{
		Description:  "newdesc!",
		TimeModified: t,
	}
	s.Client.On("UpdateUser", "id!", "newdesc!").Return(u, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: c,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "description", "newdesc!"),
					resource.TestCheckResourceAttr(s.ResourceName, "time_modified", t.String()),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestUpdateResourceIdentityUserNameShouldCreateNew() {
	s.Client.On("CreateUser", "name!", "desc!").Return(s.User, nil)
	s.Client.On("GetUser", "id!").Return(s.User, nil)
	s.Client.On("DeleteUser", "id!").Return(nil)

	c := `
		resource "baremetal_identity_user" "t" {
			name = "newname!"
			description = "desc!"
		}
	`
	u := &baremtlsdk.Resource{
		ID:   "newid!",
		Name: "newname!",
	}
	s.Client.On("CreateUser", "newname!", "desc!").Return(u, nil)
	s.Client.On("GetUser", "newid!").Return(u, nil)
	s.Client.On("DeleteUser", "newid!").Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: c,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "newname!"),
				),
			},
		},
	})
}

func (s *ResourceIdentityUserTestSuite) TestDeleteResourceIdentityUser() {
	s.Client.On("CreateUser", "name!", "desc!").Return(s.User, nil)
	s.Client.On("GetUser", "id!").Return(s.User, nil)
	s.Client.On("DeleteUser", "id!").Return(nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
			},
			resource.TestStep{
				Config: "",
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeleteUser", "id!")
}

func TestResourceIdentityUserTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityUserTestSuite))
}
