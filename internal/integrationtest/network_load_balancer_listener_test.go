// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	NetworkLoadBalancerListenerRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Required, acctest.Create, NetworkLoadBalancerListenerRepresentation)

	NetworkLoadBalancerListenerResourceConfig = NetworkLoadBalancerListenerResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Optional, acctest.Update, NetworkLoadBalancerListenerRepresentation)

	NetworkLoadBalancerNetworkLoadBalancerListenerDataSourceRepresentation = map[string]interface{}{
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
	}

	NetworkLoadBalancerNetworkLoadBalancerListenerSingularDataSourceRepresentation = map[string]interface{}{
		"listener_name":            acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_listener.test_listener.name}`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
	}

	NetworkLoadBalancerListenerRepresentation = map[string]interface{}{
		"default_backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `example_listener`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"port":                     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"is_ppv2enabled":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tcp_idle_timeout":         acctest.Representation{RepType: acctest.Optional, Create: `180`, Update: `300`},
		"udp_idle_timeout":         acctest.Representation{RepType: acctest.Optional, Create: `180`, Update: `300`},
		"protocol":                 acctest.Representation{RepType: acctest.Required, Create: `TCP_AND_UDP`, Update: `TCP_AND_UDP`},
		"ip_version":               acctest.Representation{RepType: acctest.Optional, Create: `IPV4`},
	}

	NetworkLoadBalancerTCPListenerRepresentation = map[string]interface{}{
		"default_backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `example_listener`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"port":                     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"is_ppv2enabled":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"tcp_idle_timeout":         acctest.Representation{RepType: acctest.Optional, Create: `180`, Update: `180`},
		"udp_idle_timeout":         acctest.Representation{RepType: acctest.Optional, Update: `240`},
		"protocol":                 acctest.Representation{RepType: acctest.Required, Create: `TCP`, Update: `TCP_AND_UDP`},
		"ip_version":               acctest.Representation{RepType: acctest.Optional, Create: `IPV4`},
	}

	NetworkLoadBalancerUDPListenerRepresentation = map[string]interface{}{
		"default_backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `example_listener`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"port":                     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"is_ppv2enabled":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"udp_idle_timeout":         acctest.Representation{RepType: acctest.Optional, Create: `180`, Update: `300`},
		"protocol":                 acctest.Representation{RepType: acctest.Required, Create: `UDP`},
		"ip_version":               acctest.Representation{RepType: acctest.Optional, Create: `IPV4`},
	}

	NetworkLoadBalancerL3IPListenerRepresentation = map[string]interface{}{
		"default_backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_backend_set.test_backend_set.name}`},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `example_listener`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"port":                     acctest.Representation{RepType: acctest.Required, Create: `0`, Update: `0`},
		"tcp_idle_timeout":         acctest.Representation{RepType: acctest.Optional, Create: `180`, Update: `240`},
		"udp_idle_timeout":         acctest.Representation{RepType: acctest.Optional, Create: `180`, Update: `300`},
		"l3ip_idle_timeout":        acctest.Representation{RepType: acctest.Optional, Create: `200`, Update: `400`},
		"protocol":                 acctest.Representation{RepType: acctest.Required, Create: `L3IP`},
		"ip_version":               acctest.Representation{RepType: acctest.Optional, Create: `IPV4`},
		"is_ppv2enabled":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	NetworkLoadBalancerListenerResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(NetworkLoadBalancerBackendSetRepresentation, map[string]interface{}{
				"is_preserve_source": acctest.Representation{RepType: acctest.Optional, Create: `true`},
				"policy":             acctest.Representation{RepType: acctest.Required, Create: `TWO_TUPLE`, Update: `THREE_TUPLE`},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerRepresentation)
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerListenerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerListenerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_load_balancer_listener.test_listener"
	datasourceName := "data.oci_network_load_balancer_listeners.test_listeners"
	singularDatasourceName := "data.oci_network_load_balancer_listener.test_listener"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ListenerResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Optional, acctest.Create, NetworkLoadBalancerListenerRepresentation), "networkloadbalancer", "listener", t)

	acctest.ResourceTest(t, testAccCheckNetworkLoadBalancerListenerDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Required, acctest.Create, NetworkLoadBalancerListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_listener"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP_AND_UDP"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies,
		},

		// verify TCP Listener create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Optional, acctest.Create, NetworkLoadBalancerTCPListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "is_ppv2enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_listener"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "tcp_idle_timeout", "180"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Optional, acctest.Update, NetworkLoadBalancerTCPListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttr(resourceName, "is_ppv2enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_listener"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "11"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP_AND_UDP"),
				resource.TestCheckResourceAttr(resourceName, "tcp_idle_timeout", "180"),
				resource.TestCheckResourceAttr(resourceName, "udp_idle_timeout", "240"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies,
		},

		// verify UDP Listener create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Optional, acctest.Create, NetworkLoadBalancerUDPListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "is_ppv2enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_listener"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "UDP"),
				resource.TestCheckResourceAttr(resourceName, "udp_idle_timeout", "180"),
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
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Optional, acctest.Update, NetworkLoadBalancerUDPListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttr(resourceName, "is_ppv2enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_listener"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "11"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "UDP"),
				resource.TestCheckResourceAttr(resourceName, "udp_idle_timeout", "300"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies,
		},

		// verify L3IP Listener create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Optional, acctest.Create, NetworkLoadBalancerL3IPListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "is_ppv2enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_listener"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "0"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "L3IP"),
				resource.TestCheckResourceAttr(resourceName, "tcp_idle_timeout", "180"),
				resource.TestCheckResourceAttr(resourceName, "udp_idle_timeout", "180"),
				resource.TestCheckResourceAttr(resourceName, "l3ip_idle_timeout", "200"),

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
		// verify L3IP updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Optional, acctest.Update, NetworkLoadBalancerL3IPListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttr(resourceName, "is_ppv2enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_listener"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "0"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "L3IP"),
				resource.TestCheckResourceAttr(resourceName, "tcp_idle_timeout", "240"),
				resource.TestCheckResourceAttr(resourceName, "udp_idle_timeout", "300"),
				resource.TestCheckResourceAttr(resourceName, "l3ip_idle_timeout", "400"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Optional, acctest.Create, NetworkLoadBalancerListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "is_ppv2enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_listener"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP_AND_UDP"),
				resource.TestCheckResourceAttr(resourceName, "tcp_idle_timeout", "180"),
				resource.TestCheckResourceAttr(resourceName, "udp_idle_timeout", "180"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerListenerResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Optional, acctest.Update, NetworkLoadBalancerListenerRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "default_backend_set_name"),
				resource.TestCheckResourceAttr(resourceName, "is_ppv2enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_listener"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "11"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP_AND_UDP"),
				resource.TestCheckResourceAttr(resourceName, "tcp_idle_timeout", "300"),
				resource.TestCheckResourceAttr(resourceName, "udp_idle_timeout", "300"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_load_balancer_listeners", "test_listeners", acctest.Optional, acctest.Update, NetworkLoadBalancerNetworkLoadBalancerListenerDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkLoadBalancerListenerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "listener_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "listener_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_load_balancer_listener", "test_listener", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerListenerSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkLoadBalancerListenerResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "listener_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_ppv2enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_listener"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol", "TCP_AND_UDP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tcp_idle_timeout", "300"),
				resource.TestCheckResourceAttr(singularDatasourceName, "udp_idle_timeout", "300"),
			),
		},
		// verify resource import
		{
			Config:                  config + NetworkLoadBalancerListenerRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkLoadBalancerListenerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkLoadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_load_balancer_listener" {
			noResourceFound = false
			request := oci_network_load_balancer.GetListenerRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.ListenerName = &value
			}

			if value, ok := rs.Primary.Attributes["network_load_balancer_id"]; ok {
				request.NetworkLoadBalancerId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_load_balancer")

			_, err := client.GetListener(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("NetworkLoadBalancerListener") {
		resource.AddTestSweepers("NetworkLoadBalancerListener", &resource.Sweeper{
			Name:         "NetworkLoadBalancerListener",
			Dependencies: acctest.DependencyGraph["listener"],
			F:            sweepNetworkLoadBalancerListenerResource,
		})
	}
}

func sweepNetworkLoadBalancerListenerResource(compartment string) error {
	networkLoadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkLoadBalancerClient()
	listenerIds, err := getNetworkLoadBalancerListenerIds(compartment)
	if err != nil {
		return err
	}
	for _, listenerId := range listenerIds {
		if ok := acctest.SweeperDefaultResourceId[listenerId]; !ok {
			deleteListenerRequest := oci_network_load_balancer.DeleteListenerRequest{}

			deleteListenerRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_load_balancer")
			_, error := networkLoadBalancerClient.DeleteListener(context.Background(), deleteListenerRequest)
			if error != nil {
				fmt.Printf("Error deleting Listener %s %s, It is possible that the resource is already deleted. Please verify manually \n", listenerId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkLoadBalancerListenerIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ListenerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkLoadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkLoadBalancerClient()

	listListenersRequest := oci_network_load_balancer.ListListenersRequest{}

	networkLoadBalancerIds, error := getNetworkLoadBalancerNetworkLoadBalancerIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkLoadBalancerId required for Listener resource requests \n")
	}
	for _, networkLoadBalancerId := range networkLoadBalancerIds {
		listListenersRequest.NetworkLoadBalancerId = &networkLoadBalancerId

		listListenersResponse, err := networkLoadBalancerClient.ListListeners(context.Background(), listListenersRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Listener list for NLB id : %s , %s \n", networkLoadBalancerId, err)
		}
		for _, listener := range listListenersResponse.Items {
			id := *listener.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ListenerId", id)
			acctest.SweeperDefaultResourceId[*listener.Name] = true
		}

	}
	return resourceIds, nil
}
