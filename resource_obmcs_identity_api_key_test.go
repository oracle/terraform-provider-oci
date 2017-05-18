// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"




	"github.com/stretchr/testify/suite"
)

type ResourceIdentityAPIKeyTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  time.Time
	Config       string
	ResourceName string
	Res          *baremetal.APIKey
	DeletedRes   *baremetal.APIKey
}

func (s *ResourceIdentityAPIKeyTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) { return s.Client, nil },
	)
	s.Providers = map[string]terraform.ResourceProvider{"baremetal": s.Provider}

	s.Config = `
		resource "baremetal_identity_api_key" "t" {
			user_id = "user_id"
			key_value = "1"
		}
	`
	s.Config += testProviderConfig()
	s.ResourceName = "baremetal_identity_api_key.t"

	s.TimeCreated = time.Now()
	s.Res = &baremetal.APIKey{
		Fingerprint: "fingerprint",
		KeyID:       "key_id",
		KeyValue:    "1",
		State:       baremetal.ResourceActive,
		TimeCreated: s.TimeCreated,
		UserID:      "user_id",
	}

	deletedRes := *s.Res
	s.DeletedRes = &deletedRes
	s.DeletedRes.State = baremetal.ResourceDeleted

	s.Client.On(
		"UploadAPIKey", "user_id", "1", (*baremetal.RetryTokenOptions)(nil),
	).Return(s.Res, nil).Once()

	s.Client.On(
		"DeleteAPIKey", s.Res.UserID, s.Res.Fingerprint, (*baremetal.IfMatchOptions)(nil),
	).Return(nil)
}

func (s *ResourceIdentityAPIKeyTestSuite) TestCreateAPIKey() {
	res := &baremetal.ListAPIKeyResponses{Keys: []baremetal.APIKey{*s.Res}}
	s.Client.On("ListAPIKeys", s.Res.UserID).Return(res, nil).Times(2)

	deletedRes := &baremetal.ListAPIKeyResponses{
		Keys: []baremetal.APIKey{*s.DeletedRes},
	}
	s.Client.On("ListAPIKeys", s.Res.UserID).Return(deletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "user_id", "user_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "key_value", "1"),
				),
			},
		},
	})
}

func (s *ResourceIdentityAPIKeyTestSuite) TestDeleteAPIKey() {
	res := &baremetal.ListAPIKeyResponses{Keys: []baremetal.APIKey{*s.Res}}
	s.Client.On("ListAPIKeys", s.Res.UserID).Return(res, nil).Times(2)

	deletedRes := &baremetal.ListAPIKeyResponses{
		Keys: []baremetal.APIKey{*s.DeletedRes},
	}
	s.Client.On("ListAPIKeys", s.Res.UserID).Return(deletedRes, nil)

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config:  s.Config,
				Destroy: true,
			},
		},
	})

	s.Client.AssertCalled(s.T(), "DeleteAPIKey", s.Res.UserID, s.Res.Fingerprint, (*baremetal.IfMatchOptions)(nil))
}

func TestResourceIdentityAPIKeyTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityAPIKeyTestSuite))
}
