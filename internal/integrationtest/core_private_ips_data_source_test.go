// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/stretchr/testify/suite"
)

type DatasourcePrivateIPTestSuite struct {
	suite.Suite
	Config       string
	Providers    map[string]*schema.Provider
	ResourceName string
}

func (s *DatasourcePrivateIPTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + tfresource.TestADs() + tfresource.TestVCN1() + tfresource.TestSubnet1() + tfresource.TestImage1() + tfresource.TestInstance1() + DefinedTagsDependencies + `
	data "oci_core_vnic_attachments" "t" {
		compartment_id = "${var.compartment_id}"
		availability_domain = "${data.oci_identity_availability_domains.t.availability_domains.0.name}"
		instance_id = "${oci_core_instance.t.id}"
	}

	resource "oci_core_private_ip" "t" {
		vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0], "vnic_id")}"
		ip_address = "10.0.1.23"
		defined_tags = "${map(
			"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value"
			)}"
		freeform_tags = { "Department" = "Finance"}
	}`

	s.ResourceName = "data.oci_core_private_ips.t"
}

func (s *DatasourcePrivateIPTestSuite) TestAccCorePrivateIPs_basic() {
	// Define a function closure for verifying the hostname labels from a primary and a secondary private IP
	// The datasource could retrieve them in any order.
	checkPrivateIpHostnameLabels := func(state *terraform.State) error {
		hostnameLabel1, err := acctest.FromInstanceState(state, s.ResourceName, "private_ips.0.hostname_label")
		if err != nil {
			return err
		}

		hostnameLabel2, err := acctest.FromInstanceState(state, s.ResourceName, "private_ips.1.hostname_label")
		if err != nil {
			return err
		}

		instanceHostnameLabel, err := acctest.FromInstanceState(state, "oci_core_instance.t", "create_vnic_details.0.hostname_label")
		if err != nil {
			return err
		}

		if hostnameLabel1 != "" && hostnameLabel2 != "" {
			return fmt.Errorf("Expected one of the private IPs to have a hostname label of empty, but instead got: '%s' and '%s'", hostnameLabel1, hostnameLabel2)
		}

		if hostnameLabel1 != instanceHostnameLabel && hostnameLabel2 != instanceHostnameLabel {
			return fmt.Errorf("Expected one of the private IPs to have a hostname_label of '%s', but instead got: '%s' and '%s'", instanceHostnameLabel, hostnameLabel1, hostnameLabel2)
		}

		return nil
	}

	resource.Test(s.T(), resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 s.Providers,
		Steps: []resource.TestStep{
			{
				Config: s.Config,
			},
			// list by ip address
			{
				Config: s.Config + `
				data "oci_core_private_ips" "t" {
					ip_address = "10.0.1.23"
					subnet_id = "${oci_core_subnet.t.id}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.0.ip_address", "10.0.1.23"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.0.hostname_label", ""),
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.0.freeform_tags.%", "1"),
				),
			},
			// list by vnic id
			{
				Config: s.Config + `
				data "oci_core_private_ips" "t" {
					vnic_id = "${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.#", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.time_created"),
					checkPrivateIpHostnameLabels,
				),
			},
			// list by subnet id
			{
				Config: s.Config + `
				data "oci_core_private_ips" "t" {
					subnet_id = "${oci_core_subnet.t.id}"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "private_ips.#", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.0.time_created"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.ip_address"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.display_name"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.vnic_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.availability_domain"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.compartment_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.is_primary"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.subnet_id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "private_ips.1.time_created"),
					checkPrivateIpHostnameLabels,
				),
			},
		},
	},
	)
}

// issue-routing-tag: core/virtualNetwork
func TestDatasourceCorePrivateIPTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestDatasourceCorePrivateIPTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(DatasourcePrivateIPTestSuite))
}
