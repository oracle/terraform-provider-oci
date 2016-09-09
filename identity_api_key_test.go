package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceIdentityAPIKeyTestSuite struct {
	suite.Suite
	Client       *mocks.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
	Res          *baremetal.APIKey
	Opts         []baremetal.Options
}

func (s *ResourceIdentityAPIKeyTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.Config = `
		resource "baremetal_identity_api_key" "t" {
			user_id = "user_id"
			key_value = "1"
		}
	`
	s.Config += testProviderConfig

	s.TimeCreated = time.Now()
	s.ResourceName = "baremetal_identity_api_key.t"

	s.Res = &baremetal.APIKey{
		Fingerprint: "fingerprint",
		KeyID:       "key_id",
		KeyValue:    "1",
		State:       baremetal.ResourceActive,
		TimeCreated: s.TimeCreated,
		UserID:      "user_id",
	}

	s.Client.On("UploadAPIKey", "user_id", "1", []baremetal.Options(nil)).
		Return(s.Res, nil).Once()
	s.Client.On("DeleteAPIKey", s.Res.UserID, s.Res.Fingerprint, []baremetal.Options(nil)).
		Return(nil)
}

func (s *ResourceIdentityAPIKeyTestSuite) TestCreateAPIKey() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "user_id", "user_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "key_value", "1"),
				),
			},
		},
	})
}

// func (s ResourceIdentityAPIKeyTestSuite) TestUpdateVersionForcesNewAPIKey() {
// 	config := `
// 		resource "baremetal_identity_api_key" "t" {
// 			user_id = "user_id"
// 			version = "2"
// 		}
//   `
// 	config += testProviderConfig

// 	res := &baremetal.APIKey{
// 		Password:    "new_password",
// 		TimeCreated: s.TimeCreated,
// 		UserID:      "user_id",
// 	}

// 	s.Client.On("CreateOrResetAPIKey", "user_id", []baremetal.Options(nil)).
// 		Return(res, nil)

// 	resource.UnitTest(s.T(), resource.TestCase{
// 		Providers: s.Providers,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: s.Config,
// 			},
// 			resource.TestStep{
// 				Config: config,
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(s.ResourceName, "password", "new_password"),
// 				),
// 			},
// 		},
// 	})
// }

func TestResourceIdentityAPIKeyTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityAPIKeyTestSuite))
}
