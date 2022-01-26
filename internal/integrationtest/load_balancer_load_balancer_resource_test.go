// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/loadbalancer"
	"github.com/stretchr/testify/suite"
)

type ResourceLoadBalancerLBTestSuite struct {
	suite.Suite
	Providers    map[string]terraform.ResourceProvider
	Config       string
	ResourceName string
}

var (
	loadBalancerFlexRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `example_load_balancer`, Update: `displayName2`},
		"shape":                      acctest.Representation{RepType: acctest.Required, Create: `flexible`},
		"subnet_ids":                 acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.lb_test_subnet_1.id}`, `${oci_core_subnet.lb_test_subnet_2.id}`}},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_private":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"reserved_ips":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: loadBalancerReservedIpsRepresentation},
		"shape_details":              acctest.RepresentationGroup{RepType: acctest.Required, Group: loadBalancerShapeDetailsRepresentation},
		"network_security_group_ids": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group1.id}`}, Update: []string{}},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}

	loadBalancerShapeDetailsRepresentation = map[string]interface{}{
		"maximum_bandwidth_in_mbps": acctest.Representation{RepType: acctest.Required, Create: `100`},
		"minimum_bandwidth_in_mbps": acctest.Representation{RepType: acctest.Required, Create: `10`},
	}
)

func (s *ResourceLoadBalancerLBTestSuite) SetupTest() {
	s.Providers = acctest.TestAccProviders
	acctest.PreCheck(s.T())
	s.Config = acctest.LegacyTestProviderConfig() + `
	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}
	
	resource "oci_core_virtual_network" "t" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.0.0.0/16"
		display_name = "-tf-vcn"
	}
	
	resource "oci_core_subnet" "t" {
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		cidr_block          = "10.0.0.0/24"
		display_name        = "-tf-subnet"
	}
	
	resource "oci_core_subnet" "t2" {
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.t.id}"
		availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
		route_table_id      = "${oci_core_virtual_network.t.default_route_table_id}"
		dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
		security_list_ids = ["${oci_core_virtual_network.t.default_security_list_id}"]
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet2"
	}
	
	data "oci_load_balancer_shapes" "t" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_network_security_group" "test_network_security_group" {
         compartment_id  = "${var.compartment_id}"
		 vcn_id            = "${oci_core_virtual_network.t.id}"
         display_name      =  "displayName"
    }

	resource "oci_core_network_security_group" "test_network_security_group2" {
		compartment_id = "${var.compartment_id}"
		vcn_id            = "${oci_core_virtual_network.t.id}"
	}
	`
	s.ResourceName = "oci_load_balancer.t"
}

func (s *ResourceLoadBalancerLBTestSuite) TestAccResourceLoadBalancerLB_basicPrivate() {
	var resId, resId2 string
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test Create
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids = ["${oci_core_subnet.t.id}"]
					display_name = "-tf-lb"
					is_private = true
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_mode"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "nsg_ids.#", "0"),
					func(ts *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						return err
					},
				),
			},
			// test Update without nsgIds
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape          = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids     = ["${oci_core_subnet.t.id}"]
					display_name   = "-tf-lb-updated"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb-updated"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId2 != resId {
							return fmt.Errorf("resource recreated when it should not have been")
						}
						return err
					},
				),
			},
			// test Update with nsgIds
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape          = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids     = ["${oci_core_subnet.t.id}"]
					display_name   = "-tf-lb-updated"
					network_security_group_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb-updated"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "network_security_group_ids.#", "1"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId2 != resId {
							return fmt.Errorf("resource recreated when it should not have been")
						}
						return err
					},
				),
			},
			// test Update with removing nsgIds
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape          = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids     = ["${oci_core_subnet.t.id}"]
					display_name   = "-tf-lb-updated"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb-updated"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "true"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "network_security_group_ids.#", "0"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId2 != resId {
							return fmt.Errorf("resource recreated when it should not have been")
						}
						return err
					},
				),
			},
			// verify force Update
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape          = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids     = ["${oci_core_subnet.t.id}", "${oci_core_subnet.t2.id}"]
					display_name   = "-tf-lb-updated"
					is_private 	   = false
					network_security_group_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb-updated"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_mode"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
					resource.TestCheckResourceAttr(s.ResourceName, "network_security_group_ids.#", "1"),
					func(ts *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(ts, s.ResourceName, "id")
						if resId2 == resId {
							return fmt.Errorf("resource was not recreated as expected")
						}
						return err
					},
				),
			},
		},
	})
}

func (s *ResourceLoadBalancerLBTestSuite) TestAccResourceLoadBalancerLB_basicPublic() {
	resource.Test(s.T(), resource.TestCase{
		Providers: s.Providers,
		Steps: []resource.TestStep{
			// test Create
			{
				Config: s.Config + `
				resource "oci_load_balancer" "t" {
					shape = "${data.oci_load_balancer_shapes.t.shapes.0.name}"
					compartment_id = "${var.compartment_id}"
					subnet_ids = ["${oci_core_subnet.t.id}", "${oci_core_subnet.t2.id}"]
					display_name = "-tf-lb"
				}`,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(s.ResourceName, "display_name", "-tf-lb"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "id"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "shape"),
					resource.TestCheckResourceAttr(s.ResourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_address_details.#"),
					resource.TestCheckResourceAttrSet(s.ResourceName, "ip_mode"),
					resource.TestCheckResourceAttr(s.ResourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(s.ResourceName, "state", string(loadbalancer.LoadBalancerLifecycleStateActive)),
					resource.TestCheckResourceAttrSet(s.ResourceName, "time_created"),
				),
			},
		},
	})
}

// issue-routing-tag: load_balancer/default
func TestResourceLoadBalancerLBTestSuite(t *testing.T) {
	httpreplay.SetScenario("TestResourceLoadBalancerLBTestSuite")
	defer httpreplay.SaveScenario()
	suite.Run(t, new(ResourceLoadBalancerLBTestSuite))
}

// issue-routing-tag: load_balancer/default
func TestResourceLoadBalancerLoadBalancerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestResourceLoadBalancerLoadBalancerResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_load_balancer_load_balancer.test_load_balancer"
	datasourceName := "data.oci_load_balancer_load_balancers.test_load_balancers"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerLoadBalancerDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerFlexRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttr(resourceName, "shape", "flexible"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.0.maximum_bandwidth_in_mbps", "100"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.0.minimum_bandwidth_in_mbps", "10"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies,
			},
			// verify Create 100Mbps
			{
				Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(loadBalancerFlexRepresentation, []string{"shape_details"}), map[string]interface{}{
							"shape": acctest.Representation{RepType: acctest.Required, Create: `100Mbps`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//Commenting this out as we are ignoring the changes to the tags in the resource representation.
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "reserved_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
					resource.TestCheckResourceAttrSet(resourceName, "reserved_ips.0.id"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify Update to flexshape
			{
				Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Create, loadBalancerFlexRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//Commenting this out as we are ignoring the changes to the tags in the resource representation.
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "reserved_ips.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "reserved_ips.0.id"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "flexible"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.0.maximum_bandwidth_in_mbps", "100"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.0.minimum_bandwidth_in_mbps", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LoadBalancerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(loadBalancerFlexRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					//Commenting this out as we are ignoring the changes to the tags in the resource representation.
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "reserved_ips.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "reserved_ips.0.id"),
					resource.TestCheckResourceAttr(resourceName, "shape", "flexible"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.0.maximum_bandwidth_in_mbps", "100"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.0.minimum_bandwidth_in_mbps", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Update, loadBalancerFlexRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//Commenting this out as we are ignoring the changes to the tags in the resource representation.
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "reserved_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "flexible"),
					resource.TestCheckResourceAttrSet(resourceName, "reserved_ips.0.id"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.0.maximum_bandwidth_in_mbps", "100"),
					resource.TestCheckResourceAttr(resourceName, "shape_details.0.minimum_bandwidth_in_mbps", "10"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
					acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_load_balancers", "test_load_balancers", acctest.Optional, acctest.Update, loadBalancerDataSourceRepresentation) +
					compartmentIdVariableStr + LoadBalancerResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Update, loadBalancerFlexRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "detail", "detail"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "load_balancers.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.compartment_id", compartmentId),
					//Commenting this out as we are ignoring the changes to the tags in the resource representation.
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.ip_address_details.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.is_private", "false"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.network_security_group_ids.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.shape", "flexible"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.shape_details.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.shape_details.0.maximum_bandwidth_in_mbps", "100"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.shape_details.0.minimum_bandwidth_in_mbps", "10"),
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
					"reserved_ips",
				},
				ResourceName: resourceName,
			},
		},
	})
}
