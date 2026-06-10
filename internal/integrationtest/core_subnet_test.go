// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreSubnetRequiredOnlyResource = CoreSubnetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation)

	CoreSubnetResourceConfig = CoreSubnetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, CoreSubnetRepresentation)

	CoreCoreSubnetSingularDataSourceRepresentation = map[string]interface{}{
		"subnet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	CoreCoreSubnetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MySubnet`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vcn_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreSubnetDataSourceFilterRepresentation}}
	CoreSubnetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.test_subnet.id}`}},
	}

	CoreSubnetRepresentation = map[string]interface{}{
		"cidr_block":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/24`, Update: "10.0.0.0/16"},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain":        acctest.Representation{RepType: acctest.Optional, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dhcp_options_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`, Update: `${oci_core_dhcp_options.test_dhcp_options.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `MySubnet`, Update: `displayName2`},
		"dns_label":                  acctest.Representation{RepType: acctest.Optional, Create: `dnslabel`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"prohibit_internet_ingress":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"route_table_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`, Update: `${oci_core_route_table.test_route_table.id}`},
		"security_list_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}, Update: []string{`${oci_core_security_list.test_security_list.id}`}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	CoreSubnetRepresentation2 = map[string]interface{}{
		"ipv4cidr_blocks":            acctest.Representation{RepType: acctest.Required, Create: []string{"10.0.0.0/24", "10.0.1.0/24"}},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"availability_domain":        acctest.Representation{RepType: acctest.Optional, Create: `${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}`},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dhcp_options_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_dhcp_options_id}`, Update: `${oci_core_dhcp_options.test_dhcp_options.id}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `MySubnet`, Update: `displayName2`},
		"dns_label":                  acctest.Representation{RepType: acctest.Optional, Create: `dnslabel`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"prohibit_internet_ingress":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"route_table_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_vcn.test_vcn.default_route_table_id}`, Update: `${oci_core_route_table.test_route_table.id}`},
		"security_list_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_vcn.test_vcn.default_security_list_id}`}, Update: []string{`${oci_core_security_list.test_security_list.id}`}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	CoreSubnetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_dhcp_options", "test_dhcp_options", acctest.Required, acctest.Create, CoreDhcpOptionsRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_internet_gateway", "test_internet_gateway", acctest.Required, acctest.Create, CoreInternetGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_route_table", "test_route_table", acctest.Required, acctest.Create, CoreRouteTableRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, CoreSecurityListRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, CoreCoreServiceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreIpv6VcnRepresentation, map[string]interface{}{
			"dns_label":               acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
			"is_ipv6enabled":          acctest.Representation{RepType: acctest.Optional, Create: `true`},
			"ipv6private_cidr_blocks": acctest.Representation{RepType: acctest.Optional, Create: utils.GenerateIPv6Cidrs(52)},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
	AnotherSecurityListRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_security_list", acctest.Required, acctest.Create, CoreSecurityListRepresentation)
	SubnetRequiredOnlyResourceDependencies  = AvailabilityDomainConfig + CoreVcnResourceConfig
)

// issue-routing-tag: core/virtualNetwork
func TestCoreSubnetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreSubnetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_subnet.test_subnet"
	datasourceName := "data.oci_core_subnets.test_subnets"
	singularDatasourceName := "data.oci_core_subnet.test_subnet"

	// Get subnet CIDR block based on its VCN CIDR Block
	// For example: VCN CIDR Block: 2607:9b80:9a0f:0100::/56, Subnet CIDR Block: 2607:9b80:9a0f:0100::/64
	subnetCidrBlock := `${substr(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn.ipv6cidr_blocks[0]) - 2)}${64}`
	var resId, resId2 string

	var testSteps []resource.TestStep
	testSteps = append(testSteps, resource.TestStep{
		// verify Create
		Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies +
			acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation),
		Check: acctest.ComposeAggregateTestCheckFuncWrapper(
			resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
			resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
			func(s *terraform.State) (err error) {
				resId, err = acctest.FromInstanceState(s, resourceName, "id")
				return err
			},
		),
	},

		// delete before next Create
		resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies,
		},

		// verify Create with optional ipv6cidr_blocks
		resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
					"ipv6cidr_blocks": acctest.Representation{RepType: acctest.Optional, Create: []string{subnetCidrBlock}},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MySubnet"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				testCheckCanonicalResourceAttrPair(resourceName, "ipv6cidr_block", "ipv6cidr_blocks.0"),
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify no diff after creating subnet with ipv6cidr_blocks
		resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
					"ipv6cidr_blocks": acctest.Representation{RepType: acctest.Optional, Create: []string{subnetCidrBlock}},
				})),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},

		// delete before next Create
		resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies,
		},

		// verify Create with optionals ipv4cidr_blocks
		resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, CoreSubnetRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MySubnet"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4cidr_blocks.#", "2"),
				resource.TestCheckNoResourceAttr(resourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// Test to validate that an existing IPv4 only subnet can be appended with some ipv6 blocks via the field ipv6cidr_blocks
		resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation2, map[string]interface{}{
					"ipv6cidr_blocks": acctest.Representation{RepType: acctest.Optional, Update: []string{subnetCidrBlock}},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4cidr_blocks.#", "2"),
				testCheckCanonicalResourceAttrPair(resourceName, "ipv6cidr_block", "ipv6cidr_blocks.0"),
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify no diff after appending ipv6cidr_blocks to an existing ipv4-only subnet
		resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation2, map[string]interface{}{
					"ipv6cidr_blocks": acctest.Representation{RepType: acctest.Optional, Update: []string{subnetCidrBlock}},
				})),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},

		// delete before next Create
		resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies,
		},

		// verify Create with optionals ipv6cidr_block
		resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
					"ipv6cidr_block": acctest.Representation{RepType: acctest.Optional, Create: subnetCidrBlock},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MySubnet"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		})

	var subnetIPv6CidrBlocks []string
	var subnetIpV6CidrBlockAdded = ""
	var addedCidrBlocksCount = 1
	subnetIPv6CidrBlocks = append(subnetIPv6CidrBlocks, subnetCidrBlock)

	// Add and validate 15 ipv6 cidr blocks one by one
	for i := 1000; i <= 1014; i++ {
		addedCidrBlocksCount += 1
		subnetIpV6CidrBlockAdded = fmt.Sprintf("fc00:%d::/64", i)
		subnetIPv6CidrBlocks = append(subnetIPv6CidrBlocks, subnetIpV6CidrBlockAdded)

		testSteps = append(testSteps, resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
					"ipv6cidr_blocks": acctest.Representation{RepType: acctest.Optional, Update: subnetIPv6CidrBlocks},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(addedCidrBlocksCount)),
				resource.TestCheckTypeSetElemAttr(resourceName, "ipv6cidr_blocks.*", convertToCanonical(subnetIpV6CidrBlockAdded)),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		})
	}

	// Remove and validate 15 ipv6 cidr blocks one by one
	for i := 15; i > 0; i-- {
		subnetIPv6CidrBlocks = subnetIPv6CidrBlocks[:len(subnetIPv6CidrBlocks)-1]
		testSteps = append(testSteps, resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
					"ipv6cidr_blocks": acctest.Representation{RepType: acctest.Optional, Update: subnetIPv6CidrBlocks},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(i)),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		})
	}

	// Perform the final set of validations
	testSteps = append(testSteps,
		// verify Update to the compartment (the compartment will be switched back in the next step)
		resource.TestStep{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/24"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MySubnet"),
				resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		resource.TestStep{
			Config: config + compartmentIdVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
					"ipv6cidr_block": acctest.Representation{RepType: acctest.Optional, Update: subnetCidrBlock},
				})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "dhcp_options_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "route_table_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_router_mac"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		resource.TestStep{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_subnets", "test_subnets", acctest.Optional, acctest.Update, CoreCoreSubnetDataSourceRepresentation) +
				compartmentIdVariableStr + CoreSubnetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Update, CoreSubnetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "subnets.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.dhcp_options_id"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.dns_label", "dnslabel"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.ipv4cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.ipv6cidr_block"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subnets.0.prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.route_table_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.subnet_domain_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.vcn_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.virtual_router_ip"),
				resource.TestCheckResourceAttrSet(datasourceName, "subnets.0.virtual_router_mac"),
			),
		},
		// verify singular datasource
		resource.TestStep{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreCoreSubnetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreSubnetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cidr_block", "10.0.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ipv4cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ipv6cidr_block"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "prohibit_public_ip_on_vnic", "false"),
				resource.TestCheckResourceAttr(resourceName, "prohibit_internet_ingress", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_domain_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_router_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_router_mac"),
			),
		},
		// verify resource import
		resource.TestStep{
			Config:                  config + CoreSubnetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		})

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreSubnetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, CoreSubnetRepresentation), "core", "subnet", t)
	acctest.ResourceTest(t, testAccCheckCoreSubnetDestroy, testSteps)
}

// issue-routing-tag: core/virtualNetwork
func TestCoreSubnetResource_SubnetPatch(t *testing.T) {
	httpreplay.SetScenario("TestCoreSubnetResource_SubnetPatch")
	defer httpreplay.SaveScenario()
	acctest.PreCheck(t)

	byoipv6RangeId := acctest.GetEnvSettingWithDefaultVar("byoipv6_range_id", "unknown")
	if byoipv6RangeId == "unknown" {
		t.Skip("TF_VAR_byoipv6_range_id must be set for BYO IPv6 subnet patch acceptance tests")
	}

	resourceName := "oci_core_subnet.patch_subnet"
	vcnResourceName := "oci_core_virtual_network.patch_vcn"

	vcnByoIpv6Cidrs := coreSubnetPatchTestByoIpv6Cidrs()
	initialByoBlock := vcnByoIpv6Cidrs[0]
	firstListBlock := vcnByoIpv6Cidrs[1]
	elevenAdditionalBlocks := append([]string{}, vcnByoIpv6Cidrs[2:13]...)
	blocksAfterBulkAdd := append([]string{firstListBlock}, elevenAdditionalBlocks...)
	beginningReplacementBlock := vcnByoIpv6Cidrs[13]
	blocksAfterBeginningReplace := append([]string{beginningReplacementBlock}, blocksAfterBulkAdd[1:]...)
	middleReplacementBlock := vcnByoIpv6Cidrs[14]
	blocksAfterMiddleReplace := append([]string{}, blocksAfterBeginningReplace...)
	blocksAfterMiddleReplace[len(blocksAfterMiddleReplace)/2] = middleReplacementBlock
	blocksAfterEndReplace := append([]string{}, blocksAfterMiddleReplace[:len(blocksAfterMiddleReplace)-1]...)
	blocksAfterEndReplace = append(blocksAfterEndReplace, firstListBlock)
	blocksAfterMultiReplace := append([]string{}, blocksAfterEndReplace...)
	blocksAfterMultiReplace[0] = vcnByoIpv6Cidrs[0]
	blocksAfterMultiReplace[len(blocksAfterMultiReplace)/2-1] = vcnByoIpv6Cidrs[7]
	blocksAfterMultiReplace[len(blocksAfterMultiReplace)/2] = vcnByoIpv6Cidrs[12]
	blocksAfterReorder := append([]string{blocksAfterMultiReplace[1], blocksAfterMultiReplace[2], blocksAfterMultiReplace[0]}, blocksAfterMultiReplace[3:]...)
	scalarBlockIndexAfterReorder := 2
	blocksAfterScalarRemoval := append([]string{}, blocksAfterReorder[:scalarBlockIndexAfterReorder]...)
	blocksAfterScalarRemoval = append(blocksAfterScalarRemoval, blocksAfterReorder[scalarBlockIndexAfterReorder+1:]...)
	firstBlockAfterScalarRemoval := blocksAfterScalarRemoval[0]
	middleRemovalIndex := len(blocksAfterScalarRemoval) / 2
	blocksAfterMiddleRemoval := append([]string{}, blocksAfterScalarRemoval[:middleRemovalIndex]...)
	blocksAfterMiddleRemoval = append(blocksAfterMiddleRemoval, blocksAfterScalarRemoval[middleRemovalIndex+1:]...)
	blocksAfterEndRemoval := append([]string{}, blocksAfterMiddleRemoval[:len(blocksAfterMiddleRemoval)-1]...)

	acctest.ResourceTest(t, testAccCheckCoreSubnetDestroy, []resource.TestStep{
		// Step 1 - Create a VCN with 15 byoipv6 cidrs and an Oracle-assigned GUA.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", nil, false),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(vcnResourceName, "is_oracle_gua_allocation_enabled", "true"),
				resource.TestCheckResourceAttr(vcnResourceName, "ipv6cidr_blocks.#", "1"),
				resource.TestCheckResourceAttrSet(vcnResourceName, "ipv6cidr_blocks.0"),
				resource.TestCheckResourceAttr(vcnResourceName, "byoipv6cidr_details.#", strconv.Itoa(len(vcnByoIpv6Cidrs))),
				resource.TestCheckResourceAttr(vcnResourceName, "byoipv6cidr_blocks.#", strconv.Itoa(len(vcnByoIpv6Cidrs))),
				testCheckCanonicalTypeSetContains(vcnResourceName, "byoipv6cidr_blocks", []string{initialByoBlock}),
			),
		},
		// Step 2 - Create a SUBNET under the VCN with 0 ipv6 cidrs.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", nil, true),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				testCheckMissingOrZeroResourceAttr(resourceName, "ipv6cidr_blocks.#"),
			),
		},
		// Step 3 - Add 12 ipv6 cidrs to the Subnet via the ipv6cidr_blocks field.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterBulkAdd, true),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					subnetIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: subnetIpv6PlanExpectation{
							blocksFieldChanges: 1,
							blocksAdditions:    12,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(len(blocksAfterBulkAdd))),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6cidr_blocks", blocksAfterBulkAdd),
				testCheckCanonicalListEquals(resourceName, "ipv6cidr_blocks", blocksAfterBulkAdd),
			),
		},
		// Step 4
		{
			Config:             coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterBulkAdd, true),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
		},
		// Step 5 - Replace a cidr block at the beginning of the Subnet's ipv6cidr_blocks field.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterBeginningReplace, true),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					subnetIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: subnetIpv6PlanExpectation{
							blocksFieldChanges:      1,
							blocksAdditions:         1,
							blocksRemovals:          1,
							blocksReplacementGroups: 1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(len(blocksAfterBeginningReplace))),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6cidr_blocks", blocksAfterBeginningReplace),
				testCheckCanonicalListEquals(resourceName, "ipv6cidr_blocks", blocksAfterBeginningReplace),
			),
		},
		// Step 6 - Replace a cidr block in the middle of the Subnet's ipv6cidr_blocks field.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterMiddleReplace, true),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					subnetIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: subnetIpv6PlanExpectation{
							blocksFieldChanges:      1,
							blocksAdditions:         1,
							blocksRemovals:          1,
							blocksReplacementGroups: 1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(len(blocksAfterMiddleReplace))),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6cidr_blocks", blocksAfterMiddleReplace),
				testCheckCanonicalListEquals(resourceName, "ipv6cidr_blocks", blocksAfterMiddleReplace),
			),
		},
		// Step 7 - Replace a cidr block at the end of the Subnet's ipv6cidr_blocks field.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterEndReplace, true),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					subnetIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: subnetIpv6PlanExpectation{
							blocksFieldChanges:      1,
							blocksAdditions:         1,
							blocksRemovals:          1,
							blocksReplacementGroups: 1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(len(blocksAfterEndReplace))),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6cidr_blocks", blocksAfterEndReplace),
				testCheckCanonicalListEquals(resourceName, "ipv6cidr_blocks", blocksAfterEndReplace),
			),
		},
		// Step 8 - Replace three cidr blocks in the Subnet's ipv6cidr_blocks field with entirely new cidrs.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterMultiReplace, true),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					subnetIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: subnetIpv6PlanExpectation{
							blocksFieldChanges:      1,
							blocksAdditions:         3,
							blocksRemovals:          3,
							blocksReplacementGroups: 3,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(len(blocksAfterMultiReplace))),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6cidr_blocks", blocksAfterMultiReplace),
				testCheckCanonicalListEquals(resourceName, "ipv6cidr_blocks", blocksAfterMultiReplace),
			),
		},
		// Step 9 - Reorder cidr blocks in the Subnet's ipv6cidr_blocks field without changing membership.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterReorder, true),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					subnetIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation:     subnetIpv6PlanExpectation{},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(len(blocksAfterReorder))),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6cidr_blocks", blocksAfterReorder),
				testCheckCanonicalListEquals(resourceName, "ipv6cidr_blocks", blocksAfterReorder),
			),
		},
		// Step 10 - Remove block A from ipv6cidr_blocks while ipv6cidr_block still
		// points at A. State should replace ipv6cidr_block with the first remaining
		// cidr from ipv6cidr_blocks.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterScalarRemoval, true),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					subnetIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: subnetIpv6PlanExpectation{
							blocksFieldChanges: 1,
							blocksRemovals:     1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(len(blocksAfterScalarRemoval))),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6cidr_blocks", blocksAfterScalarRemoval),
				testCheckCanonicalListEquals(resourceName, "ipv6cidr_blocks", blocksAfterScalarRemoval),
				testCheckCanonicalResourceAttrEqualsLiteral(resourceName, "ipv6cidr_block", firstBlockAfterScalarRemoval),
			),
		},
		// Step 11 - Remove a cidr block from the middle of the Subnet's ipv6cidr_blocks field.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterMiddleRemoval, true),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					subnetIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: subnetIpv6PlanExpectation{
							blocksFieldChanges: 1,
							blocksRemovals:     1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(len(blocksAfterMiddleRemoval))),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6cidr_blocks", blocksAfterMiddleRemoval),
				testCheckCanonicalListEquals(resourceName, "ipv6cidr_blocks", blocksAfterMiddleRemoval),
			),
		},
		// Step 12 - Verify no diff after removing the middle cidr block.
		{
			Config:             coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterMiddleRemoval, true),
			PlanOnly:           true,
			ExpectNonEmptyPlan: false,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPreRefresh: []plancheck.PlanCheck{
					plancheck.ExpectEmptyPlan(),
				},
			},
		},
		// Step 13 - Remove a cidr block from the end of the Subnet's ipv6cidr_blocks field.
		{
			Config: coreSubnetPatchByoIpv6Config(byoipv6RangeId, vcnByoIpv6Cidrs, "", blocksAfterEndRemoval, true),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
					subnetIpv6PlanCheck{
						resourceAddress: resourceName,
						expectation: subnetIpv6PlanExpectation{
							blocksFieldChanges: 1,
							blocksRemovals:     1,
						},
					},
				},
			},
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ipv6cidr_blocks.#", strconv.Itoa(len(blocksAfterEndRemoval))),
				testCheckCanonicalTypeSetContains(resourceName, "ipv6cidr_blocks", blocksAfterEndRemoval),
				testCheckCanonicalListEquals(resourceName, "ipv6cidr_blocks", blocksAfterEndRemoval),
			),
		},
	})
}

func coreSubnetPatchTestByoIpv6Cidrs() []string {
	cidrs := make([]string, 0, 15)
	for i := 0; i < 15; i++ {
		cidrs = append(cidrs, fmt.Sprintf("2607:f590:0000:%04x::/64", 0x2200+i))
	}
	return cidrs
}

func coreSubnetPatchByoIpv6Config(byoipv6RangeId string, vcnByoIpv6Cidrs []string, subnetIpv6Block string, subnetIpv6Blocks []string, includeSubnet bool) string {
	var config strings.Builder

	config.WriteString(acctest.LegacyTestProviderConfig())
	config.WriteString(`
		data "oci_identity_availability_domains" "patch_ads" {
			compartment_id = "${var.compartment_id}"
		}

		resource "oci_core_virtual_network" "patch_vcn" {
			cidr_block                      = "10.0.0.0/16"
			compartment_id                  = "${var.compartment_id}"
			display_name                    = "patch-vcn"
			dns_label                       = "patchvcn"
			is_ipv6enabled                  = true
			is_oracle_gua_allocation_enabled = true
`)

	for _, cidr := range vcnByoIpv6Cidrs {
		fmt.Fprintf(&config, `
			byoipv6cidr_details {
				byoipv6range_id = %q
				ipv6cidr_block  = %q
			}
`, byoipv6RangeId, cidr)
	}

	config.WriteString(`
		}
`)

	if includeSubnet {
		config.WriteString(`
		resource "oci_core_subnet" "patch_subnet" {
			availability_domain = "${data.oci_identity_availability_domains.patch_ads.availability_domains.0.name}"
			compartment_id      = "${var.compartment_id}"
			vcn_id              = "${oci_core_virtual_network.patch_vcn.id}"
			route_table_id      = "${oci_core_virtual_network.patch_vcn.default_route_table_id}"
			dhcp_options_id     = "${oci_core_virtual_network.patch_vcn.default_dhcp_options_id}"
			security_list_ids   = ["${oci_core_virtual_network.patch_vcn.default_security_list_id}"]
			cidr_block          = "10.0.2.0/24"
			display_name        = "patch-subnet"
`)

		if subnetIpv6Block != "" {
			fmt.Fprintf(&config, "			ipv6cidr_block = %q\n", subnetIpv6Block)
		}

		if subnetIpv6Blocks != nil {
			config.WriteString("			ipv6cidr_blocks = [")
			for i, cidr := range subnetIpv6Blocks {
				if i > 0 {
					config.WriteString(", ")
				}
				fmt.Fprintf(&config, "%q", cidr)
			}
			config.WriteString("]\n")
		}

		config.WriteString(`
		}
`)
	}

	return config.String()
}

func coreSubnetPatchPrivateIpv6Config(subnetIpv6Blocks []string, includeSubnet bool) string {
	var config strings.Builder

	config.WriteString(acctest.LegacyTestProviderConfig())
	config.WriteString(`
		data "oci_identity_availability_domains" "patch_ads" {
			compartment_id = "${var.compartment_id}"
		}

		resource "oci_core_virtual_network" "patch_vcn" {
			cidr_block                       = "10.0.0.0/16"
			compartment_id                   = "${var.compartment_id}"
			display_name                     = "patch-vcn-private"
			dns_label                        = "patchvcnp"
			is_ipv6enabled                   = true
			is_oracle_gua_allocation_enabled = false
			ipv6private_cidr_blocks          = ["fc00:1000::/56"]
		}
`)

	if includeSubnet {
		config.WriteString(`
		resource "oci_core_subnet" "patch_subnet" {
			availability_domain = "${data.oci_identity_availability_domains.patch_ads.availability_domains.0.name}"
			compartment_id      = "${var.compartment_id}"
			vcn_id              = "${oci_core_virtual_network.patch_vcn.id}"
			route_table_id      = "${oci_core_virtual_network.patch_vcn.default_route_table_id}"
			dhcp_options_id     = "${oci_core_virtual_network.patch_vcn.default_dhcp_options_id}"
			security_list_ids   = ["${oci_core_virtual_network.patch_vcn.default_security_list_id}"]
			cidr_block          = "10.0.2.0/24"
			display_name        = "patch-subnet-private"
`)

		if subnetIpv6Blocks != nil {
			config.WriteString("			ipv6cidr_blocks = [")
			for i, cidr := range subnetIpv6Blocks {
				if i > 0 {
					config.WriteString(", ")
				}
				fmt.Fprintf(&config, "%q", cidr)
			}
			config.WriteString("]\n")
		}

		config.WriteString(`
		}
`)
	}

	return config.String()
}

func convertToCanonical(block string) string {
	splitString := strings.Split(block, ":")

	final := []string{"0000", "0000", "0000", "0000", "0000", "0000", "0000", "0000"}

	for i := 0; i < len(splitString)-2; i++ {

		// append 4 - len(i) 0's to the left, and add it to string along with :
		final[i] = strings.Repeat("0", 4-len(splitString[i])) + splitString[i]
	}
	result := strings.Join(final, ":")

	return result + "/64"
}

func testCheckCanonicalResourceAttrPair(resourceName, lhsAttr, rhsAttr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found in state: %s", resourceName)
		}

		lhs, ok := rs.Primary.Attributes[lhsAttr]
		if !ok {
			return fmt.Errorf("attribute not found in state: %s.%s", resourceName, lhsAttr)
		}

		rhs, ok := rs.Primary.Attributes[rhsAttr]
		if !ok {
			return fmt.Errorf("attribute not found in state: %s.%s", resourceName, rhsAttr)
		}

		if convertToCanonical(lhs) != convertToCanonical(rhs) {
			return fmt.Errorf(
				"expected canonical %s (%s) to equal canonical %s (%s)",
				lhsAttr, lhs, rhsAttr, rhs,
			)
		}

		return nil
	}
}

func testAccCheckCoreSubnetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_subnet" {
			noResourceFound = false
			request := oci_core.GetSubnetRequest{}

			tmp := rs.Primary.ID
			request.SubnetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetSubnet(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.SubnetLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CoreSubnet") {
		resource.AddTestSweepers("CoreSubnet", &resource.Sweeper{
			Name:         "CoreSubnet",
			Dependencies: acctest.DependencyGraph["subnet"],
			F:            sweepCoreSubnetResource,
		})
	}
}

func sweepCoreSubnetResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	subnetIds, err := getCoreSubnetIds(compartment)
	if err != nil {
		return err
	}
	for _, subnetId := range subnetIds {
		if ok := acctest.SweeperDefaultResourceId[subnetId]; !ok {
			deleteSubnetRequest := oci_core.DeleteSubnetRequest{}

			deleteSubnetRequest.SubnetId = &subnetId

			deleteSubnetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteSubnet(context.Background(), deleteSubnetRequest)
			if error != nil {
				fmt.Printf("Error deleting Subnet %s %s, It is possible that the resource is already deleted. Please verify manually \n", subnetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &subnetId, CoreSubnetSweepWaitCondition, time.Duration(3*time.Minute),
				CoreSubnetSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreSubnetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubnetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listSubnetsRequest := oci_core.ListSubnetsRequest{}
	listSubnetsRequest.CompartmentId = &compartmentId
	listSubnetsRequest.LifecycleState = oci_core.SubnetLifecycleStateAvailable
	listSubnetsResponse, err := virtualNetworkClient.ListSubnets(context.Background(), listSubnetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Subnet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subnet := range listSubnetsResponse.Items {
		id := *subnet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubnetId", id)
	}
	return resourceIds, nil
}

func CoreSubnetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if subnetResponse, ok := response.Response.(oci_core.GetSubnetResponse); ok {
		return subnetResponse.LifecycleState != oci_core.SubnetLifecycleStateTerminated
	}
	return false
}

func CoreSubnetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetSubnet(context.Background(), oci_core.GetSubnetRequest{
		SubnetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
