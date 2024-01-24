// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreCoreVnicSingularDataSourceRepresentation = map[string]interface{}{
		"vnic_id": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}`},
	}

	CoreVnicResourceConfig         = ``
	VnicResourceConfigDependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	VnicResourceConfigIpv6Dependencies = utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Optional, acctest.Create, instanceWithAssignIPV6ResourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label":       acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
			"ipv6cidr_blocks": acctest.Representation{RepType: acctest.Required, Create: []string{`${substr(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn.ipv6cidr_blocks[0]) - 2)}${64}`}},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
			"is_ipv6enabled": acctest.Representation{RepType: acctest.Required, Create: `true`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreVnicResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVnicResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_vnic.test_vnic"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + `

data "oci_core_vnic_attachments" "t" {
	compartment_id = "${var.compartment_id}"
	instance_id = "${oci_core_instance.test_instance.id}"
}` +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vnic", "test_vnic", acctest.Required, acctest.Create, CoreCoreVnicSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreVnicResourceConfig + VnicResourceConfigDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vnic_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hostname_label"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ipv6addresses.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mac_address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "skip_source_dest_check"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}

// issue-routing-tag: core/virtualNetwork
func TestCoreVnicResource_AssignIpv6(t *testing.T) {
	httpreplay.SetScenario("TestCoreVnicResource_AssignIpv6")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_vnic.test_vnic"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + `

data "oci_core_vnic_attachments" "t" {
	compartment_id = "${var.compartment_id}"
	instance_id = "${oci_core_instance.test_instance.id}"
}` +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_vnic", "test_vnic", acctest.Required, acctest.Create, CoreCoreVnicSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreVnicResourceConfig + VnicResourceConfigIpv6Dependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vnic_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hostname_label"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ipv6addresses.#", "1"),
				resource.TestMatchResourceAttr(singularDatasourceName, "ipv6addresses.0", regexp.MustCompile(".*1000$")),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mac_address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "skip_source_dest_check"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
