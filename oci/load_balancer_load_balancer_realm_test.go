// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	govLoadBalancerSubnetDependencies = AvailabilityDomainConfig + `
	data "oci_load_balancer_shapes" "t" {
		compartment_id = "${var.compartment_id}"
	}

	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_subnet" "lb_test_subnet_1" {
		#Required
		//availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
		cidr_block = "10.0.0.0/24"
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_vcn.test_vcn.id}"
		display_name        = "lbTestSubnet"
		security_list_ids = ["${oci_core_vcn.test_vcn.default_security_list_id}"]
		ipv6cidr_block = "fd00:aaaa:0123::/64"
	}
`

	govLoadBalancerResourceDependencies = VcnRequiredOnlyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Optional, Create, RepresentationCopyWithNewProperties(VcnRepresentation, map[string]interface{}{
			"ipv6cidr_block": Representation{RepType: Optional, Create: `fd00:aaaa:0123::/48`},
			"is_ipv6enabled": Representation{RepType: Optional, Create: `true`},
		})) +
		VcnResourceDependencies + govLoadBalancerSubnetDependencies +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group1", Required, Create, networkSecurityGroupRepresentation)
)

// issue-routing-tag: load_balancer/default
func TestGovSpecificLoadBalancerLoadBalancerResource_basic(t *testing.T) {
	if !strings.Contains(GetEnvSettingWithBlankDefault("enabled_tests"), "IPv6") {
		t.Skip("DoDIPv6 test not supported in this realm")
	}
	httpreplay.SetScenario("TestGovSpecificLoadBalancerLoadBalancerResource_basic")
	defer httpreplay.SaveScenario()

	provider := TestAccProvider
	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_load_balancer.test_load_balancer"
	datasourceName := "data.oci_load_balancer_load_balancers.test_load_balancers"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { PreCheck() },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerLoadBalancerDestroy,
		Steps: []resource.TestStep{
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + govLoadBalancerResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Optional, Create, RepresentationCopyWithNewProperties(loadBalancerRepresentation, map[string]interface{}{
						"ip_mode": Representation{RepType: Optional, Create: `IPV6`},
					})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_mode", "IPV6"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + govLoadBalancerResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Optional, Update, RepresentationCopyWithNewProperties(loadBalancerRepresentation, map[string]interface{}{
						"ip_mode": Representation{RepType: Optional, Create: `IPV6`},
					})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_mode", "IPV6"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
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
					GenerateDataSourceFromRepresentationMap("oci_load_balancer_load_balancers", "test_load_balancers", Optional, Update, loadBalancerDataSourceRepresentation) +
					compartmentIdVariableStr + govLoadBalancerResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Optional, Update, RepresentationCopyWithNewProperties(loadBalancerRepresentation, map[string]interface{}{
						"ip_mode": Representation{RepType: Optional, Create: `IPV6`},
					})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "detail", "detail"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "load_balancers.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.ip_address_details.#", "2"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.is_private", "false"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.network_security_group_ids.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.shape", "100Mbps"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.time_created"),
				),
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ip_mode",
				},
				ResourceName: resourceName,
			},
		},
	})
}
