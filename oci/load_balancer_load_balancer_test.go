// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v28/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v28/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	LoadBalancerRequiredOnlyResource = LoadBalancerResourceDependencies +
		generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation)

	loadBalancerDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"detail":         Representation{repType: Optional, create: `detail`},
		"display_name":   Representation{repType: Optional, create: `example_load_balancer`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, loadBalancerDataSourceFilterRepresentation}}
	loadBalancerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_load_balancer_load_balancer.test_load_balancer.id}`}},
	}

	LoadBalancerResourceConfig = LoadBalancerResourceDependencies +
		generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Optional, Create, loadBalancerRepresentation)

	loadBalancerRepresentation = map[string]interface{}{
		"compartment_id":             Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":               Representation{repType: Required, create: `example_load_balancer`, update: `displayName2`},
		"shape":                      Representation{repType: Required, create: `100Mbps`, update: `400Mbps`},
		"subnet_ids":                 Representation{repType: Required, create: []string{`${oci_core_subnet.lb_test_subnet_1.id}`, `${oci_core_subnet.lb_test_subnet_2.id}`}},
		"defined_tags":               Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":              Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_private":                 Representation{repType: Optional, create: `false`},
		"reserved_ips":               RepresentationGroup{Optional, loadBalancerReservedIpsRepresentation},
		"network_security_group_ids": Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group1.id}`}, update: []string{}},
		"lifecycle":                  RepresentationGroup{Required, ignoreChangesLBRepresentation},
	}

	ignoreChangesLBRepresentation = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}

	loadBalancerReservedIpsRepresentation = map[string]interface{}{
		"id": Representation{repType: Optional, create: `${oci_core_public_ip.test_public_ip.id}`},
	}

	LoadBalancerSubnetDependencies = generateResourceFromRepresentationMap("oci_core_vcn", "test_lb_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
		"dns_label": Representation{repType: Required, create: `dnslabel`},
	})) +
		`
	data "oci_load_balancer_shapes" "t" {
		compartment_id = "${var.compartment_id}"
	}

	data "oci_identity_availability_domains" "ADs" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_subnet" "lb_test_subnet_1" {
		#Required
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		cidr_block = "10.0.0.0/24"
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_vcn.test_lb_vcn.id}"
		display_name        = "lbTestSubnet"
		security_list_ids = ["${oci_core_vcn.test_lb_vcn.default_security_list_id}"]
	}
	
	resource "oci_core_subnet" "lb_test_subnet_2" {
		#Required
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.1.name}"
		cidr_block = "10.0.1.0/24"
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_vcn.test_lb_vcn.id}"
		display_name        = "lbTestSubnet2"
		security_list_ids = ["${oci_core_vcn.test_lb_vcn.default_security_list_id}"]
	}
`
	LoadBalancerReservedIpDependencies = generateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", Required, Create, publicIpRepresentation)

	LoadBalancerResourceDependencies = LoadBalancerSubnetDependencies + LoadBalancerReservedIpDependencies +
		generateResourceFromRepresentationMap("oci_core_network_security_group",
			"test_network_security_group1", Required, Create, representationCopyWithNewProperties(networkSecurityGroupRepresentation, map[string]interface{}{
				"vcn_id": Representation{repType: Required, create: `${oci_core_vcn.test_lb_vcn.id}`},
			})) +
		DefinedTagsDependencies
)

func TestLoadBalancerLoadBalancerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerLoadBalancerResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_load_balancer_load_balancer.test_load_balancer"
	datasourceName := "data.oci_load_balancer_load_balancers.test_load_balancers"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerLoadBalancerDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Optional, Create, loadBalancerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//Commenting this out as we are ignoring the changes to the tags in the resource representation.
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "reserved_ips.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "reserved_ips.0.id"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LoadBalancerResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Optional, Create,
						representationCopyWithNewProperties(loadBalancerRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					//Commenting this out as we are ignoring the changes to the tags in the resource representation.
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "reserved_ips.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "reserved_ips.0.id"),
					resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Optional, Update, loadBalancerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//Commenting this out as we are ignoring the changes to the tags in the resource representation.
					//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "reserved_ips.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "reserved_ips.0.id"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "shape", "400Mbps"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_load_balancer_load_balancers", "test_load_balancers", Optional, Update, loadBalancerDataSourceRepresentation) +
					compartmentIdVariableStr + LoadBalancerResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Optional, Update, loadBalancerRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "detail", "detail"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "load_balancers.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.compartment_id", compartmentId),
					//Commenting this out as we are ignoring the changes to the tags in the resource representation.
					//resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.ip_address_details.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.is_private", "false"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.network_security_group_ids.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.shape", "400Mbps"),
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.state"),
					resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.subnet_ids.#", "2"),
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

func testAccCheckLoadBalancerLoadBalancerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_load_balancer" {
			noResourceFound = false
			request := oci_load_balancer.GetLoadBalancerRequest{}

			tmp := rs.Primary.ID
			request.LoadBalancerId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "load_balancer")

			response, err := client.GetLoadBalancer(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_load_balancer.LoadBalancerLifecycleStateDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("LoadBalancerLoadBalancer") {
		resource.AddTestSweepers("LoadBalancerLoadBalancer", &resource.Sweeper{
			Name:         "LoadBalancerLoadBalancer",
			Dependencies: DependencyGraph["loadBalancer"],
			F:            sweepLoadBalancerLoadBalancerResource,
		})
	}
}

func sweepLoadBalancerLoadBalancerResource(compartment string) error {
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()
	loadBalancerIds, err := getLoadBalancerIds(compartment)
	if err != nil {
		return err
	}
	for _, loadBalancerId := range loadBalancerIds {
		if ok := SweeperDefaultResourceId[loadBalancerId]; !ok {
			deleteLoadBalancerRequest := oci_load_balancer.DeleteLoadBalancerRequest{}

			deleteLoadBalancerRequest.LoadBalancerId = &loadBalancerId

			deleteLoadBalancerRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeleteLoadBalancer(context.Background(), deleteLoadBalancerRequest)
			if error != nil {
				fmt.Printf("Error deleting LoadBalancer %s %s, It is possible that the resource is already deleted. Please verify manually \n", loadBalancerId, error)
				continue
			}
			waitTillCondition(testAccProvider, &loadBalancerId, loadBalancerSweepWaitCondition, time.Duration(3*time.Minute),
				loadBalancerSweepResponseFetchOperation, "load_balancer", true)
		}
	}
	return nil
}

func getLoadBalancerIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "LoadBalancerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()

	listLoadBalancersRequest := oci_load_balancer.ListLoadBalancersRequest{}
	listLoadBalancersRequest.CompartmentId = &compartmentId
	listLoadBalancersRequest.LifecycleState = oci_load_balancer.LoadBalancerLifecycleStateActive
	listLoadBalancersResponse, err := loadBalancerClient.ListLoadBalancers(context.Background(), listLoadBalancersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting LoadBalancer list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, loadBalancer := range listLoadBalancersResponse.Items {
		id := *loadBalancer.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "LoadBalancerId", id)
	}
	return resourceIds, nil
}

func loadBalancerSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if loadBalancerResponse, ok := response.Response.(oci_load_balancer.GetLoadBalancerResponse); ok {
		return loadBalancerResponse.LifecycleState != oci_load_balancer.LoadBalancerLifecycleStateDeleted
	}
	return false
}

func loadBalancerSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.loadBalancerClient().GetLoadBalancer(context.Background(), oci_load_balancer.GetLoadBalancerRequest{
		LoadBalancerId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
