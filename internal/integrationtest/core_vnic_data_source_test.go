// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/core"
	"github.com/stretchr/testify/suite"
)

type DatasourceCoreVnicTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]terraform.ResourceProvider
	ResourceName string
}

func (s *DatasourceCoreVnicTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + instanceDnsConfig
	s.ResourceName = "data.oci_core_vnic.t"
}

func (s *DatasourceCoreVnicTestSuite) TestAccDatasourceCoreAttachVnic_basic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config + `
				data "oci_core_vnic_attachments" "t" {
					compartment_id = "${var.compartment_id}"
					instance_id = "${oci_core_instance.t.id}"
				}
				data "oci_core_vnic" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "mac_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "public_ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-instance-vnic"),
					resource.TestCheckResourceAttr(s.ResourceName, "hostname_label", "testinstance"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_primary", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "skip_source_dest_check", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(core.VnicLifecycleStateAvailable)),
					resource.TestCheckResourceAttr(s.ResourceName, "freeform_tags.%", "1"),
				),
			},
		},
	})
}

// issue-routing-tag: core/virtualNetwork
func TestDatasourceCoreVnicTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCoreVnicTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourceCoreVnicTestSuite))
}
