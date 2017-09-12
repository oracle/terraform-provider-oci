// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreConsoleHistoryTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	Res          *baremetal.ConsoleHistoryMetadata
	ResourceName string
}

func (s *ResourceCoreConsoleHistoryTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	p := s.Provider.(*schema.Provider)
	res := p.ResourcesMap["oci_core_console_history"]
	res.Delete = func(d *schema.ResourceData, m interface{}) (e error) {
		return nil
	}

	s.Providers = map[string]terraform.ResourceProvider{
		"oci": s.Provider,
	}
	s.Config = instanceConfig + `
    resource "oci_core_console_history" "t" {
	instance_id = "${oci_core_instance.t.id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "oci_core_console_history.t"
	s.Res = &baremetal.ConsoleHistoryMetadata{
		AvailabilityDomain: "availability_domain",
		CompartmentID:      "compartmentid",
		DisplayName:        "display_name",
		InstanceID:         "instance_id",
		ID:                 "id",
		State:              baremetal.ResourceSucceeded,
	}
	s.Res.ETag = "etag"
	s.Res.RequestID = "opcrequestid"

}

func (s *ResourceCoreConsoleHistoryTestSuite) TestCreateResourceCoreInstanceConsoleHistory() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
				),
			},
		},
	})
}

func TestResourceCoreConsoleHistoryTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreConsoleHistoryTestSuite))
}
