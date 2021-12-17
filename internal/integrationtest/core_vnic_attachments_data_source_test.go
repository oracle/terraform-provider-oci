// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVnicAttachmentTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVnicAttachmentTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + instanceConfig + `
    data "oci_core_vnic_attachments" "s" {
		compartment_id = "${var.compartment_id}"
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		instance_id = "${oci_core_instance.t.id}"
		filter {
			name = "instance_id"
			values = ["${oci_core_instance.t.id}"]
		}
    }`
	s.ResourceName = "data.oci_core_vnic_attachments.s"
}

func (s *DatasourceCoreVnicAttachmentTestSuite) TestAccDatasourceCoreVnicAttachment_basic() {
	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "vnic_attachments.#", "1"),
				),
			},
		},
	},
	)
}

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestDatasourceCoreVnicAttachmentTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreVnicAttachmentTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreVnicAttachmentTestSuite))
}
