// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/oracle/terraform-provider-baremetal/client/mocks"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
	"strconv"
)

var testPasswordsConfig = `
  data "baremetal_identity_swift_passwords" "p" {
    user_id = "%s"
  }
`

type ResourceIdentitySwiftPasswordsTestSuite struct {
	suite.Suite
	Client      *mocks.BareMetalClient
	Provider    terraform.ResourceProvider
	Providers   map[string]terraform.ResourceProvider
	TimeCreated time.Time
	Config      string
	PasswordsName string
	PasswordList baremetal.ListSwiftPasswords
}

func (s *ResourceIdentitySwiftPasswordsTestSuite) SetupTest() {
	s.Client = &mocks.BareMetalClient{}
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	},
	)
	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}
	s.TimeCreated, _ = time.Parse("2006-Jan-02", "2006-Jan-02")
	s.Config = fmt.Sprintf(testProviderConfig+testPasswordsConfig, "userid")
	s.PasswordsName = "data.baremetal_identity_swift_passwords.p"
	s.PasswordList = baremetal.ListSwiftPasswords{
		SwiftPasswords: []baremetal.SwiftPassword{
			{
				Password: "pass",
				ID: "1",
				UserID: "userid",
				Description: "desc",
				State: "available",
				InactiveStatus: 0,
				TimeCreated: time.Now(),
			},
			{
				Password: "pass",
				ID: "2",
				UserID: "userid",
				Description: "desc",
				State: "available",
				InactiveStatus: 0,
				TimeCreated: time.Now(),
			},
		},
	}

	s.Client.On(
		"ListSwiftPasswords",
		"userid",
	).Return(&s.PasswordList, nil)

}

func (s *ResourceIdentitySwiftPasswordsTestSuite) TestListResourceIdentitySwiftPasswords() {
	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.PasswordsName, "passwords.0.id", s.PasswordList.SwiftPasswords[0].ID),
					resource.TestCheckResourceAttr(s.PasswordsName, "passwords.1.id", s.PasswordList.SwiftPasswords[1].ID),
					resource.TestCheckResourceAttr(s.PasswordsName, "passwords.#", strconv.Itoa(len(s.PasswordList.SwiftPasswords))),
				),
			},
		},
	},
	)
}

func TestResourceIdentitySwiftPasswordsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceIdentitySwiftPasswordsTestSuite))
}
