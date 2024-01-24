// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OdaOdaPrivateEndpointScanProxyRequiredOnlyResource = OdaOdaPrivateEndpointScanProxyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint_scan_proxy", "test_oda_private_endpoint_scan_proxy", acctest.Required, acctest.Create, OdaOdaPrivateEndpointScanProxyRepresentation)

	OdaOdaPrivateEndpointScanProxyResourceConfig = OdaOdaPrivateEndpointScanProxyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint_scan_proxy", "test_oda_private_endpoint_scan_proxy", acctest.Optional, acctest.Update, OdaOdaPrivateEndpointScanProxyRepresentation)

	OdaOdaOdaPrivateEndpointScanProxySingularDataSourceRepresentation = map[string]interface{}{
		"oda_private_endpoint_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_oda_oda_private_endpoint.test_oda_private_endpoint.id}`},
		"oda_private_endpoint_scan_proxy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_oda_oda_private_endpoint_scan_proxy.test_oda_private_endpoint_scan_proxy.id}`},
	}

	OdaOdaOdaPrivateEndpointScanProxyDataSourceRepresentation = map[string]interface{}{
		"oda_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_oda_oda_private_endpoint.test_oda_private_endpoint.id}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `CREATING`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: OdaOdaPrivateEndpointScanProxyDataSourceFilterRepresentation}}

	OdaOdaPrivateEndpointScanProxyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_oda_oda_private_endpoint_scan_proxy.test_oda_private_endpoint_scan_proxy.id}`}},
	}

	OdaOdaPrivateEndpointScanProxyRepresentation = map[string]interface{}{
		"oda_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_oda_oda_private_endpoint.test_oda_private_endpoint.id}`},
		"protocol":                acctest.Representation{RepType: acctest.Required, Create: `TCP`},
		"scan_listener_infos":     acctest.RepresentationGroup{RepType: acctest.Required, Group: OdaOdaPrivateEndpointScanProxyScanListenerInfosRepresentation},
		"scan_listener_type":      acctest.Representation{RepType: acctest.Required, Create: `FQDN`},
	}

	OdaOdaPrivateEndpointScanProxyScanListenerInfosRepresentation = map[string]interface{}{
		"scan_listener_fqdn": acctest.Representation{RepType: acctest.Required, Create: `myprefix1-test-scan`},
		"scan_listener_port": acctest.Representation{RepType: acctest.Required, Create: `1521`},
	}

	OdaOdaPrivateEndpointScanProxyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint", "test_oda_private_endpoint", acctest.Required, acctest.Create, OdaOdaPrivateEndpointRepresentation) +
		OdaOdaPrivateEndpointResourceDependencies
)

// issue-routing-tag: oda/default
func TestOdaOdaPrivateEndpointScanProxyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOdaOdaPrivateEndpointScanProxyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_oda_oda_private_endpoint_scan_proxy.test_oda_private_endpoint_scan_proxy"
	datasourceName := "data.oci_oda_oda_private_endpoint_scan_proxies.test_oda_private_endpoint_scan_proxies"
	singularDatasourceName := "data.oci_oda_oda_private_endpoint_scan_proxy.test_oda_private_endpoint_scan_proxy"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OdaOdaPrivateEndpointScanProxyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint_scan_proxy", "test_oda_private_endpoint_scan_proxy", acctest.Required, acctest.Create, OdaOdaPrivateEndpointScanProxyRepresentation), "oda", "odaPrivateEndpointScanProxy", t)

	acctest.ResourceTest(t, testAccCheckOdaOdaPrivateEndpointScanProxyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OdaOdaPrivateEndpointScanProxyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_oda_oda_private_endpoint_scan_proxy", "test_oda_private_endpoint_scan_proxy", acctest.Required, acctest.Create, OdaOdaPrivateEndpointScanProxyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "oda_private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_infos.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scan_listener_type", "FQDN"),

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
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_oda_oda_private_endpoint_scan_proxies", "test_oda_private_endpoint_scan_proxies", acctest.Optional, acctest.Update, OdaOdaOdaPrivateEndpointScanProxyDataSourceRepresentation) +
				compartmentIdVariableStr + OdaOdaPrivateEndpointScanProxyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "oda_private_endpoint_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "CREATING"),

				resource.TestCheckResourceAttr(datasourceName, "oda_private_endpoint_scan_proxy_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oda_private_endpoint_scan_proxy_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_oda_oda_private_endpoint_scan_proxy", "test_oda_private_endpoint_scan_proxy", acctest.Required, acctest.Create, OdaOdaOdaPrivateEndpointScanProxySingularDataSourceRepresentation) +
				compartmentIdVariableStr + OdaOdaPrivateEndpointScanProxyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oda_private_endpoint_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oda_private_endpoint_scan_proxy_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol", "TCP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_listener_infos.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_listener_infos.0.scan_listener_fqdn", "myprefix1-test-scan"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_listener_infos.0.scan_listener_port", "1521"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_listener_type", "FQDN"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + OdaOdaPrivateEndpointScanProxyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getPrivateEndpointScanProxyImportCompositeId(resourceName),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func getPrivateEndpointScanProxyImportCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("odaPrivateEndpoints/" + rs.Primary.Attributes["oda_private_endpoint_id"] + "/odaPrivateEndpointScanProxies/" + rs.Primary.Attributes["id"]), nil
	}
}

func testAccCheckOdaOdaPrivateEndpointScanProxyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_oda_oda_private_endpoint_scan_proxy" {
			noResourceFound = false
			request := oci_oda.GetOdaPrivateEndpointScanProxyRequest{}

			if value, ok := rs.Primary.Attributes["oda_private_endpoint_id"]; ok {
				request.OdaPrivateEndpointId = &value
			}

			tmp := rs.Primary.ID
			request.OdaPrivateEndpointScanProxyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oda")

			response, err := client.GetOdaPrivateEndpointScanProxy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_oda.OdaPrivateEndpointScanProxyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OdaOdaPrivateEndpointScanProxy") {
		resource.AddTestSweepers("OdaOdaPrivateEndpointScanProxy", &resource.Sweeper{
			Name:         "OdaOdaPrivateEndpointScanProxy",
			Dependencies: acctest.DependencyGraph["odaPrivateEndpointScanProxy"],
			F:            sweepOdaOdaPrivateEndpointScanProxyResource,
		})
	}
}

func sweepOdaOdaPrivateEndpointScanProxyResource(compartment string) error {
	managementClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementClient()
	odaPrivateEndpointScanProxyIds, err := getOdaOdaPrivateEndpointScanProxyIds(compartment)
	if err != nil {
		return err
	}
	for _, odaPrivateEndpointScanProxyId := range odaPrivateEndpointScanProxyIds {
		if ok := acctest.SweeperDefaultResourceId[odaPrivateEndpointScanProxyId]; !ok {
			deleteOdaPrivateEndpointScanProxyRequest := oci_oda.DeleteOdaPrivateEndpointScanProxyRequest{}

			deleteOdaPrivateEndpointScanProxyRequest.OdaPrivateEndpointScanProxyId = &odaPrivateEndpointScanProxyId

			deleteOdaPrivateEndpointScanProxyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "oda")
			_, error := managementClient.DeleteOdaPrivateEndpointScanProxy(context.Background(), deleteOdaPrivateEndpointScanProxyRequest)
			if error != nil {
				fmt.Printf("Error deleting OdaPrivateEndpointScanProxy %s %s, It is possible that the resource is already deleted. Please verify manually \n", odaPrivateEndpointScanProxyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &odaPrivateEndpointScanProxyId, OdaOdaPrivateEndpointScanProxySweepWaitCondition, time.Duration(3*time.Minute),
				OdaOdaPrivateEndpointScanProxySweepResponseFetchOperation, "oda", true)
		}
	}
	return nil
}

func getOdaOdaPrivateEndpointScanProxyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OdaPrivateEndpointScanProxyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementClient()

	listOdaPrivateEndpointScanProxiesRequest := oci_oda.ListOdaPrivateEndpointScanProxiesRequest{}
	odaPrivateEndpointIds, error := getOdaOdaPrivateEndpointIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting odaPrivateEndpointId required for OdaPrivateEndpointScanProxy resource requests \n")
	}
	for _, odaPrivateEndpointId := range odaPrivateEndpointIds {
		listOdaPrivateEndpointScanProxiesRequest.OdaPrivateEndpointId = &odaPrivateEndpointId

		listOdaPrivateEndpointScanProxiesRequest.LifecycleState = oci_oda.OdaPrivateEndpointScanProxyLifecycleStateActive
		listOdaPrivateEndpointScanProxiesResponse, err := managementClient.ListOdaPrivateEndpointScanProxies(context.Background(), listOdaPrivateEndpointScanProxiesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting OdaPrivateEndpointScanProxy list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, odaPrivateEndpointScanProxy := range listOdaPrivateEndpointScanProxiesResponse.Items {
			id := *odaPrivateEndpointScanProxy.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OdaPrivateEndpointScanProxyId", id)
		}

	}
	return resourceIds, nil
}

func OdaOdaPrivateEndpointScanProxySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if odaPrivateEndpointScanProxyResponse, ok := response.Response.(oci_oda.GetOdaPrivateEndpointScanProxyResponse); ok {
		return odaPrivateEndpointScanProxyResponse.LifecycleState != oci_oda.OdaPrivateEndpointScanProxyLifecycleStateDeleted
	}
	return false
}

func OdaOdaPrivateEndpointScanProxySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagementClient().GetOdaPrivateEndpointScanProxy(context.Background(), oci_oda.GetOdaPrivateEndpointScanProxyRequest{
		OdaPrivateEndpointScanProxyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
