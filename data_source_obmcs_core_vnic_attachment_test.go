// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type ResourceCoreVnicAttachmentsTestSuite struct {
	suite.Suite
	Client       mockableClient
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *ResourceCoreVnicAttachmentsTestSuite) SetupTest() {
	s.Client = GetTestProvider()
	s.Provider = Provider(func(d *schema.ResourceData) (interface{}, error) {
		return s.Client, nil
	})

	s.Providers = map[string]terraform.ResourceProvider{
		"baremetal": s.Provider,
	}

	s.Config = instanceConfig + `
    data "baremetal_core_vnic_attachments" "s" {
      compartment_id = "${var.compartment_id}"
      availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
      vnic_id = "${baremetal_core_virtual_network.t.id}"
      instance_id = "${baremetal_core_instance.t.id}"
    }
  `
	s.Config += testProviderConfig()
	s.ResourceName = "data.baremetal_core_vnic_attachments.s"

}

func (s *ResourceCoreVnicAttachmentsTestSuite) TestResourceReadCoreVnicAttachments() {
	resource.UnitTest(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.#", "0"),
				),
			},
		},
	},
	)


}

func TestResourceCoreVnicAttachmentsTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceCoreVnicAttachmentsTestSuite))
}
