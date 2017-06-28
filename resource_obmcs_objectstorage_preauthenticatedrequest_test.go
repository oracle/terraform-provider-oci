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

type ResourcePARTestSuite struct {
	suite.Suite
	Client       mockableClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	TimeExpired  baremetal.Time
	Config       string
	ResourceName string
	Res          *PreauthenticatedRequestResourceCrud
}

func (s *ResourcePARTestSuite) SetupTest() {
	s.Client = GetTestProvider()

	s.Provider = Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return s.Client, nil
		},
	)

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.TimeCreated = baremetal.Time{Time: time.Now()}
	t, _ := time.Parse(time.RFC3339, "2019-11-10T23:00:00Z")
	s.TimeExpired = baremetal.Time{Time: t}

	s.Config = `
		resource "baremetal_objectstorage_preauthrequest" "t" {
			namespace ="internalbriangustafson"
			bucket = "testOne"
			name = "parOne"
			access_type = "AnyObjectWrite"
			time_expires = "2019-11-10T23:00:00Z"
		}`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_objectstorage_preauthrequest.t"

}

func (s *ResourcePARTestSuite) TestCreatePAR() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(s.ResourceName, "name", "parOne"),
				),
			},
		},
	})
}

func (s *ResourcePARTestSuite) TestDeletePAR() {

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

func TestResourcePARTestSuite(t *testing.T) {
	suite.Run(t, new(ResourcePARTestSuite))
}
