// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	baremetal "github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVnicAttachmentTestSuite struct {
	suite.Suite
	Client       *baremetal.Client
	Config       string
	Provider     terraform.ResourceProvider
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVnicAttachmentTestSuite) SetupTest() {
	s.Client = testAccClient
	s.Provider = testAccProvider
	s.Providers = testAccProviders
	s.Config = testProviderConfig() + instanceConfig + `
    data "oci_core_vnic_attachments" "s" {
		compartment_id = "${var.compartment_id}"
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		instance_id = "${oci_core_instance.t.id}"
    }`
	s.ResourceName = "data.oci_core_vnic_attachments.s"
}

func (s *DatasourceCoreVnicAttachmentTestSuite) TestAccDatasourceCoreVnicAttachment_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            s.Config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_attachments.#"),
				),
			},
		},
	},
	)
}

func TestDatasourceCoreVnicAttachmentTestSuite(t *testing.T) {
	suite.Run(t, new(DatasourceCoreVnicAttachmentTestSuite))
}
