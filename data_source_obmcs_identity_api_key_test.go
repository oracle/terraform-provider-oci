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

type DatasourceIdentityAPIKeysTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
	List         *baremetal.ListAPIKeyResponses
}

func (s *DatasourceIdentityAPIKeysTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = `
	resource "oci_identity_user" "t" {
			name = "-tf-test"
			description = "automated test user"
		}
		resource "oci_identity_api_key" "t" {
			user_id = "${oci_identity_user.t.id}"
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
	s.ResourceName = "data.oci_identity_api_keys.t"

	b1 := baremetal.APIKey{
		Fingerprint: "fingerprint",
		KeyID:       "id1",
		KeyValue:    "key_value",
		State:       baremetal.ResourceAvailable,
		TimeCreated: time.Now(),
		UserID:      "user_id",
	}

	b2 := b1
	b2.KeyID = "id2"

	s.List = &baremetal.ListAPIKeyResponses{
		Keys: []baremetal.APIKey{b1, b2},
	}
}

func (s *DatasourceIdentityAPIKeysTestSuite) TestReadAPIKeys() {

	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
			},
			{
				Config: s.Config + `
				    data "oci_identity_api_keys" "t" {
				      user_id = "${oci_identity_user.t.id}"
				    }`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "api_keys.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "api_keys.#", "1"),
				),
			},
		},
	},
	)
}

func TestDatasourceIdentityAPIKeysTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceIdentityAPIKeysTestSuite))
}
