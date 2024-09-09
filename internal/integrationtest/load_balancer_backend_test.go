// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	BackendRequiredOnlyResource = BackendResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", acctest.Required, acctest.Create, backendRepresentation)

	backendDataSourceRepresentation = map[string]interface{}{
		"backendset_name":  acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: backendDataSourceFilterRepresentation}}
	backendDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_load_balancer_backend.test_backend.name}`}},
	}

	backendRepresentation = map[string]interface{}{
		"backendset_name":  acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
		"ip_address":       acctest.Representation{RepType: acctest.Required, Create: `10.0.0.3`},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"port":             acctest.Representation{RepType: acctest.Required, Create: `10`},
		"backup":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"drain":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"max_connections":  acctest.Representation{RepType: acctest.Optional, Create: `375`, Update: `450`},
		"offline":          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"weight":           acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	BackendResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_certificate", "test_certificate", acctest.Required, acctest.Create, certificateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerBackendResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerBackendResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_backend.test_backend"
	datasourceName := "data.oci_load_balancer_backends.test_backends"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BackendResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", acctest.Optional, acctest.Create, backendRepresentation), "loadbalancer", "backend", t)

	acctest.ResourceTest(t, testAccCheckLoadBalancerBackendDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BackendResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", acctest.Required, acctest.Create, backendRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backendset_name"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BackendResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BackendResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", acctest.Optional, acctest.Create, backendRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backendset_name"),
				resource.TestCheckResourceAttr(resourceName, "backup", "false"),
				resource.TestCheckResourceAttr(resourceName, "drain", "false"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "max_connections", "375"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "offline", "false"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "weight", "10"),

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
			Config: config + compartmentIdVariableStr + BackendResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", acctest.Optional, acctest.Update, backendRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backendset_name"),
				resource.TestCheckResourceAttr(resourceName, "backup", "true"),
				resource.TestCheckResourceAttr(resourceName, "drain", "true"),
				resource.TestCheckResourceAttr(resourceName, "ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "max_connections", "450"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "offline", "true"),
				resource.TestCheckResourceAttr(resourceName, "port", "10"),
				resource.TestCheckResourceAttr(resourceName, "weight", "11"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_backends", "test_backends", acctest.Optional, acctest.Update, backendDataSourceRepresentation) +
				compartmentIdVariableStr + BackendResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend", "test_backend", acctest.Optional, acctest.Update, backendRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "backendset_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "backends.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.backup", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.drain", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.ip_address", "10.0.0.3"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.max_connections", "450"),
				resource.TestCheckResourceAttrSet(datasourceName, "backends.0.name"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.offline", "true"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.port", "10"),
				resource.TestCheckResourceAttr(datasourceName, "backends.0.weight", "11"),
			),
		},
		// verify resource import
		{
			Config:            config + BackendRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"backendset_name",
				"state",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckLoadBalancerBackendDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LoadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_backend" {
			noResourceFound = false
			request := oci_load_balancer.GetBackendRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.BackendName = &value
			}

			if value, ok := rs.Primary.Attributes["backendset_name"]; ok {
				request.BackendSetName = &value
			}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")

			_, err := client.GetBackend(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("LoadBalancerBackend") {
		resource.AddTestSweepers("LoadBalancerBackend", &resource.Sweeper{
			Name:         "LoadBalancerBackend",
			Dependencies: acctest.DependencyGraph["backend"],
			F:            sweepLoadBalancerBackendResource,
		})
	}
}

func sweepLoadBalancerBackendResource(compartment string) error {
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()
	backendIds, err := getBackendIds(compartment)
	if err != nil {
		return err
	}
	for _, backendId := range backendIds {
		if ok := acctest.SweeperDefaultResourceId[backendId]; !ok {
			deleteBackendRequest := oci_load_balancer.DeleteBackendRequest{}

			deleteBackendRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeleteBackend(context.Background(), deleteBackendRequest)
			if error != nil {
				fmt.Printf("Error deleting Backend %s %s, It is possible that the resource is already deleted. Please verify manually \n", backendId, error)
				continue
			}
		}
	}
	return nil
}

func getBackendIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BackendId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()

	listBackendsRequest := oci_load_balancer.ListBackendsRequest{}

	backendsetNames, error := getBackendSetIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting backendsetName required for Backend resource requests \n")
	}
	for _, backendsetName := range backendsetNames {
		listBackendsRequest.BackendSetName = &backendsetName

		loadBalancerIds, error := getLoadBalancerIds(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting loadBalancerId required for Backend resource requests \n")
		}
		for _, loadBalancerId := range loadBalancerIds {
			listBackendsRequest.LoadBalancerId = &loadBalancerId

			listBackendsResponse, err := loadBalancerClient.ListBackends(context.Background(), listBackendsRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting Backend list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, backend := range listBackendsResponse.Items {
				id := *backend.Name
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BackendId", id)
			}

		}
	}
	return resourceIds, nil
}
