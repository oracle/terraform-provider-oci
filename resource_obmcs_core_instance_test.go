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

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

type ResourceCoreInstanceTestSuite struct {
	suite.Suite
	Client       client.BareMetalClient
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	TimeCreated  baremetal.Time
	Config       string
	ResourceName string
	Res          *baremetal.Instance
	DeletedRes   *baremetal.Instance
}

func (s *ResourceCoreInstanceTestSuite) SetupTest() {
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

	s.Config = instanceConfig + `
	data "baremetal_core_instances" "s" {
      		compartment_id = "${var.compartment_id}"
      		availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
    	}`

	s.Config += testProviderConfig()

	s.ResourceName = "baremetal_core_instance.t"
}

func (s *ResourceCoreInstanceTestSuite) TestCreateResourceCoreInstance() {

	resource.UnitTest(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "instance_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", baremetal.ResourceRunning),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip"),
					resource.TestCheckResourceAttrSet("data.baremetal_core_instances.s", "instances.#"),
				),
			},
		},
	})
}

func TestIsStatefulResource(t *testing.T) {
	var sr crud.StatefulResource
	sr = &InstanceResourceCrud{}
	if sr == nil {
		t.Fail()
	}
}

func TestResourceCoreInstanceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreInstanceTestSuite))
}
