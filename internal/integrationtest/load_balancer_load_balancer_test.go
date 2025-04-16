// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LoadBalancerRequiredOnlyResource = LoadBalancerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation)

	loadBalancerDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"detail":         acctest.Representation{RepType: acctest.Optional, Create: `detail`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `example_load_balancer`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: loadBalancerDataSourceFilterRepresentation}}
	loadBalancerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_load_balancer_load_balancer.test_load_balancer.id}`}},
	}

	LoadBalancerResourceConfig = LoadBalancerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Create, loadBalancerRepresentation)

	loadBalancerRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `example_load_balancer`, Update: `displayName2`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `100Mbps`, Update: `400Mbps`},
		"subnet_ids":     acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.lb_test_subnet_1.id}`, `${oci_core_subnet.lb_test_subnet_2.id}`}},

		// For laptop testing, comment this defined_tags out along with the DefinedTagsDependencies + line below
		// Failure to do so results in test failures:  Error: Reference to undeclared resource.
		// "defined_tags": acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},

		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_private":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"security_attributes":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle-zpr.sa-test-lbaas.mode": "enforce", "oracle-zpr.sa-test-lbaas.value": "create-zpr-tersi-lbaas"}, Update: map[string]string{"oracle-zpr.sa-test-lbaas.value": "update-zpr-tersi-lbaas", "oracle-zpr.sa-test-lbaas.mode": "enforce"}},
		"is_request_id_enabled":        acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `true`},
		"request_id_header":            acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `X-MyRequestB-Id`},
		"is_delete_protection_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"reserved_ips":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: loadBalancerReservedIpsRepresentation},
		"network_security_group_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group1.id}`}, Update: []string{}},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}

	loadBalancer2Representation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `example_load_balancer2`, Update: `displayName3`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `100Mbps`, Update: `400Mbps`},
		"subnet_ids":     acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.lb_test_subnet_3.id}`, `${oci_core_subnet.lb_test_subnet_4.id}`}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}

	loadBalancer3Representation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `example_load_balancer3`, Update: `displayName4`},
		"shape":          acctest.Representation{RepType: acctest.Required, Create: `100Mbps`, Update: `400Mbps`},
		"subnet_ids":     acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_subnet.lb_test_subnet_3.id}`, `${oci_core_subnet.lb_test_subnet_4.id}`}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
	}

	ignoreChangesLBRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	loadBalancerReservedIpsRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_public_ip.test_public_ip.id}`},
	}

	LoadBalancerSubnetDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_lb_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
		"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
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
	resource "oci_core_subnet" "lb_test_subnet_3" {
		#Required
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
		cidr_block = "10.0.2.0/24"
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_vcn.test_lb_vcn.id}"
		display_name        = "lbTestSubnet3"
		security_list_ids = ["${oci_core_vcn.test_lb_vcn.default_security_list_id}"]
	}

	resource "oci_core_subnet" "lb_test_subnet_4" {
		#Required
		availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.1.name}"
		cidr_block = "10.0.3.0/24"
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_vcn.test_lb_vcn.id}"
		display_name        = "lbTestSubnet4"
		security_list_ids = ["${oci_core_vcn.test_lb_vcn.default_security_list_id}"]
	}
`

	LoadBalancerReservedIpDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", acctest.Required, acctest.Create, CorePublicIpRepresentation)

	LoadBalancerResourceDependencies = LoadBalancerSubnetDependencies + LoadBalancerReservedIpDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group",
			"test_network_security_group1", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreNetworkSecurityGroupRepresentation, map[string]interface{}{
				"vcn_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_lb_vcn.id}`},
			})) // +
		// For laptop testing comment out this line
		// Failure to do so results in
		//     test_helpers.go:535: Step 1/7 error: Error running apply: exit status 1
		//       [DEBUG] Using modified User-Agent: Terraform/0.12.31 HashiCorp-terraform-exec/0.14.0
		//       Error: 404-NotAuthorizedOrNotFound, Authorization failed or requested resource not found.
		//       Suggestion: Either the resource has been deleted or service Identity Tag Namespace need policy to access this resource.
		// DefinedTagsDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerLoadBalancerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerLoadBalancerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_load_balancer_load_balancer.test_load_balancer"
	datasourceName := "data.oci_load_balancer_load_balancers.test_load_balancers"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LoadBalancerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Create, loadBalancerRepresentation), "loadbalancer", "loadBalancer", t)

	acctest.ResourceTest(t, testAccCheckLoadBalancerLoadBalancerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
				resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),

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
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LoadBalancerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Create, loadBalancerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_delete_protection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_request_id_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "request_id_header", "X-Request-Id"),
				resource.TestCheckResourceAttr(resourceName, "reserved_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.oracle-zpr.sa-test-lbaas.value", "create-zpr-tersi-lbaas"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.oracle-zpr.sa-test-lbaas.mode", "enforce"),
				resource.TestCheckResourceAttrSet(resourceName, "reserved_ips.0.id"),
				resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
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
					acctest.RepresentationCopyWithNewProperties(loadBalancerRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_load_balancer"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_delete_protection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_request_id_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "request_id_header", "X-Request-Id"),
				resource.TestCheckResourceAttr(resourceName, "reserved_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),

				resource.TestCheckResourceAttrSet(resourceName, "reserved_ips.0.id"),
				resource.TestCheckResourceAttr(resourceName, "shape", "100Mbps"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
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
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Update, loadBalancerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_delete_protection_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_request_id_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "request_id_header", "X-MyRequestB-Id"),
				resource.TestCheckResourceAttr(resourceName, "reserved_ips.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.oracle-zpr.sa-test-lbaas.value", "update-zpr-tersi-lbaas"),
				resource.TestCheckResourceAttr(resourceName, "security_attributes.oracle-zpr.sa-test-lbaas.mode", "enforce"),
				resource.TestCheckResourceAttr(resourceName, "shape", "400Mbps"),
				resource.TestCheckResourceAttrSet(resourceName, "reserved_ips.0.id"),
				resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
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
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Optional, acctest.Update, loadBalancerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "detail", "detail"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "load_balancers.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.ip_address_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.is_delete_protection_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.is_private", "false"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.network_security_group_ids.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.is_request_id_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.security_attributes.%", "2"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.request_id_header", "X-MyRequestB-Id"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.shape", "400Mbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "load_balancers.0.subnet_ids.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancers.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:            config + LoadBalancerRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"ip_mode",
				"reserved_ips",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckLoadBalancerLoadBalancerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LoadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_load_balancer" {
			noResourceFound = false
			request := oci_load_balancer.GetLoadBalancerRequest{}

			tmp := rs.Primary.ID
			request.LoadBalancerId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("LoadBalancerLoadBalancer") {
		resource.AddTestSweepers("LoadBalancerLoadBalancer", &resource.Sweeper{
			Name:         "LoadBalancerLoadBalancer",
			Dependencies: acctest.DependencyGraph["loadBalancer"],
			F:            sweepLoadBalancerLoadBalancerResource,
		})
	}
}

func sweepLoadBalancerLoadBalancerResource(compartment string) error {
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()
	loadBalancerIds, err := getLoadBalancerIds(compartment)
	if err != nil {
		return err
	}
	for _, loadBalancerId := range loadBalancerIds {
		if ok := acctest.SweeperDefaultResourceId[loadBalancerId]; !ok {
			deleteLoadBalancerRequest := oci_load_balancer.DeleteLoadBalancerRequest{}

			deleteLoadBalancerRequest.LoadBalancerId = &loadBalancerId

			deleteLoadBalancerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeleteLoadBalancer(context.Background(), deleteLoadBalancerRequest)
			if error != nil {
				fmt.Printf("Error deleting LoadBalancer %s %s, It is possible that the resource is already deleted. Please verify manually \n", loadBalancerId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &loadBalancerId, loadBalancerSweepWaitCondition, time.Duration(3*time.Minute),
				loadBalancerSweepResponseFetchOperation, "load_balancer", true)
		}
	}
	return nil
}

func getLoadBalancerIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LoadBalancerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LoadBalancerId", id)
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

func loadBalancerSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.LoadBalancerClient().GetLoadBalancer(context.Background(), oci_load_balancer.GetLoadBalancerRequest{
		LoadBalancerId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
