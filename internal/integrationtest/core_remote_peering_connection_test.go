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
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreRemotePeeringConnectionRequiredOnlyResource = CoreRemotePeeringConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", acctest.Required, acctest.Create, CoreRemotePeeringConnectionRepresentation)

	CoreCoreRemotePeeringConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"drg_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_drg.test_drg.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreRemotePeeringConnectionDataSourceFilterRepresentation}}
	CoreRemotePeeringConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_remote_peering_connection.test_remote_peering_connection.id}`}},
	}

	CoreRemotePeeringConnectionRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"drg_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_core_drg.test_drg.id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CoreRemotePeeringConnectionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_drg", "test_drg", acctest.Required, acctest.Create, CoreDrgRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: core/default
func TestCoreRemotePeeringConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreRemotePeeringConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_remote_peering_connection.test_remote_peering_connection"
	datasourceName := "data.oci_core_remote_peering_connections.test_remote_peering_connections"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreRemotePeeringConnectionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", acctest.Optional, acctest.Create, CoreRemotePeeringConnectionRepresentation), "core", "remotePeeringConnection", t)

	acctest.ResourceTest(t, testAccCheckCoreRemotePeeringConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreRemotePeeringConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", acctest.Required, acctest.Create, CoreRemotePeeringConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreRemotePeeringConnectionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreRemotePeeringConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", acctest.Optional, acctest.Create, CoreRemotePeeringConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
				resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreRemotePeeringConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreRemotePeeringConnectionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
				resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + CoreRemotePeeringConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", acctest.Optional, acctest.Update, CoreRemotePeeringConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
				resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_remote_peering_connections", "test_remote_peering_connections", acctest.Optional, acctest.Update, CoreCoreRemotePeeringConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + CoreRemotePeeringConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", acctest.Optional, acctest.Update, CoreRemotePeeringConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),

				resource.TestCheckResourceAttr(datasourceName, "remote_peering_connections.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "remote_peering_connections.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "remote_peering_connections.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.drg_id"),
				resource.TestCheckResourceAttr(datasourceName, "remote_peering_connections.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.is_cross_tenancy_peering"),
				resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.peering_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "remote_peering_connections.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreRemotePeeringConnectionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreRemotePeeringConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_remote_peering_connection" {
			noResourceFound = false
			request := oci_core.GetRemotePeeringConnectionRequest{}

			tmp := rs.Primary.ID
			request.RemotePeeringConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetRemotePeeringConnection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.RemotePeeringConnectionLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("CoreRemotePeeringConnection") {
		resource.AddTestSweepers("CoreRemotePeeringConnection", &resource.Sweeper{
			Name:         "CoreRemotePeeringConnection",
			Dependencies: acctest.DependencyGraph["remotePeeringConnection"],
			F:            sweepCoreRemotePeeringConnectionResource,
		})
	}
}

func sweepCoreRemotePeeringConnectionResource(compartment string) error {
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()
	remotePeeringConnectionIds, err := getCoreRemotePeeringConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, remotePeeringConnectionId := range remotePeeringConnectionIds {
		if ok := acctest.SweeperDefaultResourceId[remotePeeringConnectionId]; !ok {
			deleteRemotePeeringConnectionRequest := oci_core.DeleteRemotePeeringConnectionRequest{}

			deleteRemotePeeringConnectionRequest.RemotePeeringConnectionId = &remotePeeringConnectionId

			deleteRemotePeeringConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteRemotePeeringConnection(context.Background(), deleteRemotePeeringConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting RemotePeeringConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", remotePeeringConnectionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &remotePeeringConnectionId, CoreRemotePeeringConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				CoreRemotePeeringConnectionSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getCoreRemotePeeringConnectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RemotePeeringConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := acctest.GetTestClients(&schema.ResourceData{}).VirtualNetworkClient()

	listRemotePeeringConnectionsRequest := oci_core.ListRemotePeeringConnectionsRequest{}
	listRemotePeeringConnectionsRequest.CompartmentId = &compartmentId
	listRemotePeeringConnectionsResponse, err := virtualNetworkClient.ListRemotePeeringConnections(context.Background(), listRemotePeeringConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting RemotePeeringConnection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, remotePeeringConnection := range listRemotePeeringConnectionsResponse.Items {
		id := *remotePeeringConnection.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RemotePeeringConnectionId", id)
	}
	return resourceIds, nil
}

func CoreRemotePeeringConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if remotePeeringConnectionResponse, ok := response.Response.(oci_core.GetRemotePeeringConnectionResponse); ok {
		return remotePeeringConnectionResponse.LifecycleState != oci_core.RemotePeeringConnectionLifecycleStateTerminated
	}
	return false
}

func CoreRemotePeeringConnectionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.VirtualNetworkClient().GetRemotePeeringConnection(context.Background(), oci_core.GetRemotePeeringConnectionRequest{
		RemotePeeringConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
