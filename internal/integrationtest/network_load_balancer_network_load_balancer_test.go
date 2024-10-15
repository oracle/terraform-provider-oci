// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	NetworkLoadBalancerNetworkLoadBalancerRequiredOnlyResource = NetworkLoadBalancerNetworkLoadBalancerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerRepresentation)

	NetworkLoadBalancerNetworkLoadBalancerResourceConfig = NetworkLoadBalancerNetworkLoadBalancerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Optional, acctest.Update, NetworkLoadBalancerNetworkLoadBalancerRepresentation)

	NetworkLoadBalancerNetworkLoadBalancerNetworkLoadBalancerSingularDataSourceRepresentation = map[string]interface{}{
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
	}

	NetworkLoadBalancerNetworkLoadBalancerNetworkLoadBalancerDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: NetworkLoadBalancerNetworkLoadBalancerDataSourceFilterRepresentation}}
	NetworkLoadBalancerNetworkLoadBalancerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`}},
	}

	ignoreNlbDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	NetworkLoadBalancerNetworkLoadBalancerRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"subnet_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"assigned_private_ipv4":          acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.50`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_preserve_source_destination": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_private":                     acctest.Representation{RepType: acctest.Required, Create: `false`},
		"is_symmetric_hash_enabled":      acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"nlb_ip_version":                 acctest.Representation{RepType: acctest.Required, Create: `IPV4`},
		"reserved_ips":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: networkLoadBalancerReservedIpsRepresentation},
		"network_security_group_ids":     acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: NetworkLoadBalancerIgnoreChangesRepresentation},
		"security_attributes":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"secAttriZprNlbIAD.secAttriIAD.mode": "enforce", "secAttriZprNlbIAD.secAttriIAD.value": "someVal"}},
	}

	NetworkLoadBalancerSubnetIpv6CidrRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `displayNameSubnetIpv6Cidr`},
		"subnet_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"assigned_private_ipv4":          acctest.Representation{RepType: acctest.Required, Create: `10.0.0.51`},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"is_preserve_source_destination": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_private":                     acctest.Representation{RepType: acctest.Required, Create: `false`},
		"nlb_ip_version":                 acctest.Representation{RepType: acctest.Required, Create: `IPV4_AND_IPV6`},
		"reserved_ips":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: networkLoadBalancerReservedIpsRepresentation},
		"subnet_ipv6cidr":                acctest.Representation{RepType: acctest.Required, Create: `2000:1000:1200:0001::/64`},
		"network_security_group_ids":     acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: NetworkLoadBalancerIgnoreChangesRepresentation},
	}

	NetworkLoadBalancerAssignIpv6Representation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `displayNameAssignIpv6`},
		"subnet_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"assigned_private_ipv4":          acctest.Representation{RepType: acctest.Required, Create: `10.0.0.53`},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"is_preserve_source_destination": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_private":                     acctest.Representation{RepType: acctest.Required, Create: `false`},
		"nlb_ip_version":                 acctest.Representation{RepType: acctest.Required, Create: `IPV4_AND_IPV6`},
		"reserved_ips":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: networkLoadBalancerReservedIpsRepresentation},
		"assigned_ipv6":                  acctest.Representation{RepType: acctest.Required, Create: `2000:1000:1200:0001:0001:0000:0001:000`},
		"network_security_group_ids":     acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: NetworkLoadBalancerIgnoreChangesRepresentation},
	}

	networkLoadBalancerRepresentationIpv6 = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"subnet_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":                   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_preserve_source_destination": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_private":                     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"assigned_ipv6":                  acctest.Representation{RepType: acctest.Optional, Create: `2000:1000:1200:0001:0001:0000:0001:0000`},
		"nlb_ip_version":                 acctest.Representation{RepType: acctest.Optional, Create: `IPV4_AND_IPV6`},
		"network_security_group_ids":     acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"lifecycle":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: NetworkLoadBalancerIgnoreChangesRepresentation},
	}
	networkLoadBalancerReservedIpsRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_public_ip.test_public_ip.id}`},
	}
	NetworkLoadBalancerIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}
	VcnIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`is_oracle_gua_allocation_enabled`, `defined_tags`, `ipv6private_cidr_blocks`, `is_ipv6enabled`}},
	}
	NetworkLoadBalancerReservedIpDependencies                  = acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Required, acctest.Create, CorePublicIpRepresentation)
	NetworkLoadBalancerNetworkLoadBalancerResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"cidr_block":      acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/24`},
			"ipv6cidr_blocks": acctest.Representation{RepType: acctest.Optional, Create: []string{`2000:1000:1200:0001::/64`, `fc00:1000:1200:0001::/64`}},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"cidr_block":                       acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
			"is_ipv6enabled":                   acctest.Representation{RepType: acctest.Required, Create: `true`},
			"is_oracle_gua_allocation_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
			"ipv6private_cidr_blocks":          acctest.Representation{RepType: acctest.Required, Create: []string{`2000:1000:1200::/56`, `fc00:1000:1200::/56`}},
			"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: VcnIgnoreChangesRepresentation},
		})) +
		AvailabilityDomainConfig + DefinedTagsDependencies
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerNetworkLoadBalancerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerNetworkLoadBalancerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_network_load_balancer_network_load_balancer.test_network_load_balancer"
	datasourceName := "data.oci_network_load_balancer_network_load_balancers.test_network_load_balancers"
	singularDatasourceName := "data.oci_network_load_balancer_network_load_balancer.test_network_load_balancer"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckNetworkLoadBalancerNetworkLoadBalancerDestroy, []resource.TestStep{
		// Initialize Tag dependencies: After a tag is created, if it is defined in the resource immediately, a 400-InvalidParameter error due to invalid tags may be returned.
		// However, this error is not observed if we wait for some time. To prevent the issue, a preconfigured 30-second wait is added.
		{
			Config: config + compartmentIdVariableStr + DefinedTagsDependencies,
		},

		// verify Create with optionals
		{
			//wait for 30 sec
			PreConfig: func() {
				time.Sleep(30 * time.Second)
			},
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Optional, acctest.Create, networkLoadBalancerRepresentationIpv6),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
				resource.TestCheckResourceAttr(resourceName, "nlb_ip_version", "IPV4_AND_IPV6"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.is_public", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_addresses.0.ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.2.ip_address", "2000:1000:1200:1:1:0:1:0"),
				resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies,
		},

		// verify Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Optional, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.1.ip_address", "10.0.0.50"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_symmetric_hash_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "nlb_ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.is_public", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_addresses.0.ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_addresses.0.reserved_ip.0.id"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(NetworkLoadBalancerNetworkLoadBalancerRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_symmetric_hash_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "nlb_ip_version", "IPV4"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_addresses.0.ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "ip_addresses.0.reserved_ip.0.id"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),

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
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Optional, acctest.Update, NetworkLoadBalancerNetworkLoadBalancerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.is_public", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
				resource.TestCheckResourceAttr(resourceName, "nlb_ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify removal of NSGs
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(NetworkLoadBalancerNetworkLoadBalancerRepresentation, map[string]interface{}{
						"network_security_group_ids": acctest.Representation{RepType: acctest.Required, Create: []string{}},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.is_public", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "0"),

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
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers", "test_network_load_balancers", acctest.Optional, acctest.Update, NetworkLoadBalancerNetworkLoadBalancerNetworkLoadBalancerDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Optional, acctest.Update, NetworkLoadBalancerNetworkLoadBalancerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "network_load_balancer_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "network_load_balancer_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerNetworkLoadBalancerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceConfig + NetworkLoadBalancerReservedIpDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_load_balancer_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_addresses.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_addresses.0.is_public", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_preserve_source_destination", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_private", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_symmetric_hash_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nlb_ip_version", "IPV4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "network_security_group_ids.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + NetworkLoadBalancerNetworkLoadBalancerRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"assigned_ipv6",
				"assigned_private_ipv4",
				"reserved_ips",
				"subnet_ipv6cidr",
			},
			ResourceName: resourceName,
		},
		// Verify AssignIpv6 create
		{Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies +
			acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerAssignIpv6Representation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameAssignIpv6"),
				resource.TestCheckResourceAttr(resourceName, "assigned_ipv6", "2000:1000:1200:0001:0001:0000:0001:000"),
				resource.TestCheckResourceAttr(resourceName, "assigned_private_ipv4", "10.0.0.53"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// Verify AssignIpv6 during update
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Update, NetworkLoadBalancerAssignIpv6Representation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameAssignIpv6"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "true"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf(".Resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// Verify subnet_ipv6cidr create
		{Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies +
			acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerSubnetIpv6CidrRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameSubnetIpv6Cidr"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ipv6cidr", "2000:1000:1200:0001::/64"),
				resource.TestCheckResourceAttr(resourceName, "assigned_private_ipv4", "10.0.0.51"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		//Verify subnet_ipv6cidr during update
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Update, NetworkLoadBalancerSubnetIpv6CidrRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameSubnetIpv6Cidr"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "true"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf(".Resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
	})
}

func testAccCheckNetworkLoadBalancerNetworkLoadBalancerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkLoadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_load_balancer_network_load_balancer" {
			noResourceFound = false
			request := oci_network_load_balancer.GetNetworkLoadBalancerRequest{}

			tmp := rs.Primary.ID
			request.NetworkLoadBalancerId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_load_balancer")

			response, err := client.GetNetworkLoadBalancer(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_network_load_balancer.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("NetworkLoadBalancerNetworkLoadBalancer") {
		resource.AddTestSweepers("NetworkLoadBalancerNetworkLoadBalancer", &resource.Sweeper{
			Name:         "NetworkLoadBalancerNetworkLoadBalancer",
			Dependencies: acctest.DependencyGraph["networkLoadBalancer"],
			F:            sweepNetworkLoadBalancerNetworkLoadBalancerResource,
		})
	}
}

func sweepNetworkLoadBalancerNetworkLoadBalancerResource(compartment string) error {
	networkLoadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkLoadBalancerClient()
	networkLoadBalancerIds, err := getNetworkLoadBalancerNetworkLoadBalancerIds(compartment)
	if err != nil {
		return err
	}
	for _, networkLoadBalancerId := range networkLoadBalancerIds {
		if ok := acctest.SweeperDefaultResourceId[networkLoadBalancerId]; !ok {
			deleteNetworkLoadBalancerRequest := oci_network_load_balancer.DeleteNetworkLoadBalancerRequest{}

			deleteNetworkLoadBalancerRequest.NetworkLoadBalancerId = &networkLoadBalancerId

			deleteNetworkLoadBalancerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_load_balancer")
			_, error := networkLoadBalancerClient.DeleteNetworkLoadBalancer(context.Background(), deleteNetworkLoadBalancerRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkLoadBalancer %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkLoadBalancerId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &networkLoadBalancerId, NetworkLoadBalancerNetworkLoadBalancerSweepWaitCondition, time.Duration(3*time.Minute),
				NetworkLoadBalancerNetworkLoadBalancerSweepResponseFetchOperation, "network_load_balancer", true)
		}
	}
	return nil
}

func getNetworkLoadBalancerNetworkLoadBalancerIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkLoadBalancerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkLoadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkLoadBalancerClient()

	listNetworkLoadBalancersRequest := oci_network_load_balancer.ListNetworkLoadBalancersRequest{}
	listNetworkLoadBalancersRequest.CompartmentId = &compartmentId
	listNetworkLoadBalancersRequest.LifecycleState = oci_network_load_balancer.ListNetworkLoadBalancersLifecycleStateActive
	listNetworkLoadBalancersResponse, err := networkLoadBalancerClient.ListNetworkLoadBalancers(context.Background(), listNetworkLoadBalancersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NetworkLoadBalancer list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, networkLoadBalancer := range listNetworkLoadBalancersResponse.Items {
		id := *networkLoadBalancer.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkLoadBalancerId", id)
	}
	return resourceIds, nil
}

func NetworkLoadBalancerNetworkLoadBalancerSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if networkLoadBalancerResponse, ok := response.Response.(oci_network_load_balancer.GetNetworkLoadBalancerResponse); ok {
		return networkLoadBalancerResponse.LifecycleState != oci_network_load_balancer.LifecycleStateDeleted
	}
	return false
}

func NetworkLoadBalancerNetworkLoadBalancerSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.NetworkLoadBalancerClient().GetNetworkLoadBalancer(context.Background(), oci_network_load_balancer.GetNetworkLoadBalancerRequest{
		NetworkLoadBalancerId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
