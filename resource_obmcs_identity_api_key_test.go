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
		resource "baremetal_identity_user" "t" {
			name = "-tf-user"
			description = "automated test user"
		}
		resource "baremetal_identity_api_key" "t" {
			user_id = "${baremetal_identity_user.t.id}"
			key_value = <<EOF
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtBLQAGmKJ7tpfzYJyqLG
ZDwHL51+d6T8Z00BnP9CFfzxZZZ48PcYSUHuTyCM8mR5JqYLyH6C8tZ/DKqwxUnc
ONgBytG3MM42bgxfHIhsZRj5rCz1oqWlSLuXvgww1kuqWnt6r+NtnXog439YsGTH
RotrTLTdEgOxH0EFP5uHUc9w/Uix7rWU7GB2ra060oeTB/hKpts5U70eI2EI6ec9
1sJdUIj7xNfBJeQQrz4CFUrkyzL06211CFvhmxH2hA9gBKOqC3rGL8XraHZBhGWn
mXlrQB7nNKsJrrv5fHwaPDrAY4iNP2W0q3LRpyNigJ6cgRuGJhHa82iHPmxgIx8m
fwIDAQAB
-----END PUBLIC KEY-----
EOF
		}
	`
	s.Config += testProviderConfig()
	s.ResourceName = "baremetal_identity_api_key.t"

}

func (s *ResourceIdentityAPIKeyTestSuite) TestCreateAPIKey() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "user_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceActive),
				),
			},
		},
	})
}

func (s *ResourceIdentityAPIKeyTestSuite) TestDeleteAPIKey() {

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

}

func TestResourceIdentityAPIKeyTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentityAPIKeyTestSuite))
}
