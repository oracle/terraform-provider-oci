// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/v58/dataflow"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PrivateEndpointRequiredOnlyResource = PrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, privateEndpointRepresentation)

	PrivateEndpointResourceConfig = PrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Update, privateEndpointRepresentation)

	privateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_private_endpoint.test_private_endpoint.id}`},
	}

	privateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `pe_1234`, Update: `displayName2`},
		"display_name_starts_with": acctest.Representation{RepType: acctest.Optional, Create: `displayNameStartsWith`},
		"owner_principal_id":       acctest.Representation{RepType: acctest.Optional, Create: `${var.user_ocid}`},
		"state":                    acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"filter":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: privateEndpointDataSourceFilterRepresentation}}
	privateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dataflow_private_endpoint.test_private_endpoint.id}`}},
	}

	privateEndpointRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dns_zones":      acctest.Representation{RepType: acctest.Required, Create: []string{`custpvtsubnet.oraclevcn.com`}, Update: []string{`db.custpvtsubnet.oraclevcn.com`}},
		"subnet_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `pe_1234`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"max_host_count": acctest.Representation{RepType: acctest.Optional, Create: `256`, Update: `512`},
		"nsg_ids":        acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
	}

	PrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: dataflow/default
func TestDataflowPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	userId := utils.GetEnvSettingWithBlankDefault("user_ocid")
	userIdVariableStr := fmt.Sprintf("variable \"user_ocid\" { default = \"%s\" }\n", userId)

	resourceName := "oci_dataflow_private_endpoint.test_private_endpoint"
	datasourceName := "data.oci_dataflow_private_endpoints.test_private_endpoints"
	singularDatasourceName := "data.oci_dataflow_private_endpoint.test_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Create, privateEndpointRepresentation), "dataflow", "privateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckDataflowPrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, privateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PrivateEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Create, privateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe_1234"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_host_count", "256"),
				resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(privateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "pe_1234"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_host_count", "256"),
				resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + PrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Update, privateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "max_host_count", "512"),
				resource.TestCheckResourceAttrSet(resourceName, "owner_principal_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_private_endpoints", "test_private_endpoints", acctest.Required, acctest.Update, privateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + PrivateEndpointResourceDependencies + userIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Optional, acctest.Update, privateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "private_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "private_endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", acctest.Required, acctest.Create, privateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns_zones.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_host_count", "512"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner_user_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + PrivateEndpointResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataflowPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataFlowClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataflow_private_endpoint" {
			noResourceFound = false
			request := oci_dataflow.GetPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.PrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")

			response, err := client.GetPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dataflow.PrivateEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataflowPrivateEndpoint") {
		resource.AddTestSweepers("DataflowPrivateEndpoint", &resource.Sweeper{
			Name:         "DataflowPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["privateEndpoint"],
			F:            sweepDataflowPrivateEndpointResource,
		})
	}
}

func sweepDataflowPrivateEndpointResource(compartment string) error {
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()
	privateEndpointIds, err := getPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, privateEndpointId := range privateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[privateEndpointId]; !ok {
			deletePrivateEndpointRequest := oci_dataflow.DeletePrivateEndpointRequest{}

			deletePrivateEndpointRequest.PrivateEndpointId = &privateEndpointId

			deletePrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dataflow")
			_, error := dataFlowClient.DeletePrivateEndpoint(context.Background(), deletePrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting PrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", privateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &privateEndpointId, privateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				privateEndpointSweepResponseFetchOperation, "dataflow", true)
		}
	}
	return nil
}

func getPrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataFlowClient := acctest.GetTestClients(&schema.ResourceData{}).DataFlowClient()

	listPrivateEndpointsRequest := oci_dataflow.ListPrivateEndpointsRequest{}
	listPrivateEndpointsRequest.CompartmentId = &compartmentId
	listPrivateEndpointsRequest.LifecycleState = oci_dataflow.ListPrivateEndpointsLifecycleStateInactive
	listPrivateEndpointsResponse, err := dataFlowClient.ListPrivateEndpoints(context.Background(), listPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, privateEndpoint := range listPrivateEndpointsResponse.Items {
		id := *privateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PrivateEndpointId", id)
	}
	return resourceIds, nil
}

func privateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if privateEndpointResponse, ok := response.Response.(oci_dataflow.GetPrivateEndpointResponse); ok {
		return privateEndpointResponse.LifecycleState != oci_dataflow.PrivateEndpointLifecycleStateDeleted
	}
	return false
}

func privateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataFlowClient().GetPrivateEndpoint(context.Background(), oci_dataflow.GetPrivateEndpointRequest{
		PrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
