// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// configuration to create ipv4 lbaas and update to ipv6
	loadBalancerRepresentationIPV6 = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `example_load_balancer`},
		"shape":                 acctest.Representation{RepType: acctest.Required, Create: `100Mbps`},
		"subnet_ids":            acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.ipv6_subnet_1.id}`}},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"ip_mode":               acctest.Representation{RepType: acctest.Optional, Create: `IPV4`, Update: `IPV6`},
		"is_request_id_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}

	loadBalancerReservedIpsRepresentationIP6 = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_ipv6.reserved_ipv6.id}`, Update: `${oci_core_ipv6.reserved_ipv6.id}`},
	}

	loadBalancerRepresentationIPV6ReservedIp = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `example_load_balancer`, Update: `example_load_balancer_updated`},
		"shape":                 acctest.Representation{RepType: acctest.Required, Create: `100Mbps`},
		"subnet_ids":            acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.ipv6_subnet_1.id}`}, Update: []string{`${oci_core_subnet.ipv6_subnet_1.id}`}},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Biology"}},
		"ip_mode":               acctest.Representation{RepType: acctest.Optional, Create: `IPV4`, Update: `IPV6`},
		"is_request_id_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
		"reserved_ips":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: loadBalancerReservedIpsRepresentationIP6},
		"ipv6subnet_cidr":       acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `${local.ipv6SubnetCidr1}`},
	}

	loadBalancerRepresentationIPV6_2 = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `example_load_balancer`},
		"shape":                 acctest.Representation{RepType: acctest.Required, Create: `100Mbps`},
		"subnet_ids":            acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.ipv6_subnet_1.id}`}},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"ip_mode":               acctest.Representation{RepType: acctest.Optional, Create: `IPV4`, Update: `IPV6`},
		"ipv6subnet_cidr":       acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `${local.ipv6SubnetCidr1}`},
		"is_request_id_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}

	// configuration to create with default ipv4 and update ip_mode and ipv6subnetcidr
	loadBalancerRepresentationIPV6_3 = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `example_load_balancer`},
		"shape":                 acctest.Representation{RepType: acctest.Required, Create: `100Mbps`},
		"subnet_ids":            acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.ipv6_subnet_1.id}`}},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"ip_mode":               acctest.Representation{RepType: acctest.Optional, Create: `IPV4`, Update: `IPV6`},
		"ipv6subnet_cidr":       acctest.Representation{RepType: acctest.Optional, Create: `fc00:1000::/64`, Update: `${local.ipv6SubnetCidr1}`},
		"is_request_id_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`}, "lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerLoadBalancerResourceIPv6_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerLoadBalancerResourceIPv6_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	//compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_load_balancer_load_balancer.test_load_balancer"
	//	datasourceName := "data.oci_load_balancer_load_balancers.test_load_balancers"

	var resId1 string
	var resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	//acctest.SaveConfigContent(config+compartmentIdVariableStr+LoadBalancerResourceDependencies+
	//	acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Create, loadBalancerRepresentationIPV6), "loadbalancer", "loadBalancer", t)

	acctest.ResourceTest(t, testAccCheckLoadBalancerLoadBalancerDestroy, []resource.TestStep{

		// verify create with default ip_mode
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentationIPV6),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
				resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ip_mode", "IPV4"),

				func(s *terraform.State) (err error) {
					resId1, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Create, loadBalancerRepresentationIPV6),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_request_id_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "request_id_header", "X-Request-Id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId1, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId1, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify update to ip_mode
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Update, loadBalancerRepresentationIPV6),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_mode", "IPV6"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId1 != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{

			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies,
		},
		// verify create with default ip_mode and empty ipv6cidr
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentationIPV6_2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
				resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ip_mode", "IPV4"),

				func(s *terraform.State) (err error) {
					resId1, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify update to ip_mode and ipv6cidr from empty to value
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Update, loadBalancerRepresentationIPV6_2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_mode", "IPV6"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId1 != resId2 {
						return fmt.Errorf("Resource was recreated instead of being updated when ipv6subnetcidr was initially empty.")
					}
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies,
		},

		// create with default ip_mode and ipv6subnetcidr
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentationIPV6_3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
				resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ip_mode", "IPV4"),

				func(s *terraform.State) (err error) {
					resId1, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to ip_mode and ipv6prefixcidr
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Update, loadBalancerRepresentationIPV6_3),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ip_mode", "IPV6"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId1 != resId2 {
						return fmt.Errorf("Resource was supposed to be updated when updating ipmode and ipv6prefixcidr.")
					}
					return err
				},
			),
		},
	})
}
