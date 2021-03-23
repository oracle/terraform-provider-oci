// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v37/common"
	oci_core "github.com/oracle/oci-go-sdk/v37/core"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	RemotePeeringConnectionRequiredOnlyResource = RemotePeeringConnectionResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", Required, Create, remotePeeringConnectionRepresentation)

	remotePeeringConnectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"drg_id":         Representation{repType: Optional, create: `${oci_core_drg.test_drg.id}`},
		"filter":         RepresentationGroup{Required, remotePeeringConnectionDataSourceFilterRepresentation}}
	remotePeeringConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_core_remote_peering_connection.test_remote_peering_connection.id}`}},
	}

	remotePeeringConnectionRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"drg_id":         Representation{repType: Required, create: `${oci_core_drg.test_drg.id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	RemotePeeringConnectionResourceDependencies = generateResourceFromRepresentationMap("oci_core_drg", "test_drg", Required, Create, drgRepresentation) +
		DefinedTagsDependencies
)

func TestCoreRemotePeeringConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreRemotePeeringConnectionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_remote_peering_connection.test_remote_peering_connection"
	datasourceName := "data.oci_core_remote_peering_connections.test_remote_peering_connections"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+RemotePeeringConnectionResourceDependencies+
		generateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", Optional, Create, remotePeeringConnectionRepresentation), "core", "remotePeeringConnection", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckCoreRemotePeeringConnectionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + RemotePeeringConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", Required, Create, remotePeeringConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + RemotePeeringConnectionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + RemotePeeringConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", Optional, Create, remotePeeringConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RemotePeeringConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", Optional, Create,
						representationCopyWithNewProperties(remotePeeringConnectionRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				Config: config + compartmentIdVariableStr + RemotePeeringConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", Optional, Update, remotePeeringConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "drg_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "is_cross_tenancy_peering"),
					resource.TestCheckResourceAttrSet(resourceName, "peering_status"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
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
					generateDataSourceFromRepresentationMap("oci_core_remote_peering_connections", "test_remote_peering_connections", Optional, Update, remotePeeringConnectionDataSourceRepresentation) +
					compartmentIdVariableStr + RemotePeeringConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_remote_peering_connection", "test_remote_peering_connection", Optional, Update, remotePeeringConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "drg_id"),

					resource.TestCheckResourceAttr(datasourceName, "remote_peering_connections.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "remote_peering_connections.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "remote_peering_connections.0.defined_tags.%", "1"),
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
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckCoreRemotePeeringConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).virtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_remote_peering_connection" {
			noResourceFound = false
			request := oci_core.GetRemotePeeringConnectionRequest{}

			tmp := rs.Primary.ID
			request.RemotePeeringConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("CoreRemotePeeringConnection") {
		resource.AddTestSweepers("CoreRemotePeeringConnection", &resource.Sweeper{
			Name:         "CoreRemotePeeringConnection",
			Dependencies: DependencyGraph["remotePeeringConnection"],
			F:            sweepCoreRemotePeeringConnectionResource,
		})
	}
}

func sweepCoreRemotePeeringConnectionResource(compartment string) error {
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()
	remotePeeringConnectionIds, err := getRemotePeeringConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, remotePeeringConnectionId := range remotePeeringConnectionIds {
		if ok := SweeperDefaultResourceId[remotePeeringConnectionId]; !ok {
			deleteRemotePeeringConnectionRequest := oci_core.DeleteRemotePeeringConnectionRequest{}

			deleteRemotePeeringConnectionRequest.RemotePeeringConnectionId = &remotePeeringConnectionId

			deleteRemotePeeringConnectionRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "core")
			_, error := virtualNetworkClient.DeleteRemotePeeringConnection(context.Background(), deleteRemotePeeringConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting RemotePeeringConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", remotePeeringConnectionId, error)
				continue
			}
			waitTillCondition(testAccProvider, &remotePeeringConnectionId, remotePeeringConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				remotePeeringConnectionSweepResponseFetchOperation, "core", true)
		}
	}
	return nil
}

func getRemotePeeringConnectionIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "RemotePeeringConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	virtualNetworkClient := GetTestClients(&schema.ResourceData{}).virtualNetworkClient()

	listRemotePeeringConnectionsRequest := oci_core.ListRemotePeeringConnectionsRequest{}
	listRemotePeeringConnectionsRequest.CompartmentId = &compartmentId
	listRemotePeeringConnectionsResponse, err := virtualNetworkClient.ListRemotePeeringConnections(context.Background(), listRemotePeeringConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting RemotePeeringConnection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, remotePeeringConnection := range listRemotePeeringConnectionsResponse.Items {
		id := *remotePeeringConnection.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "RemotePeeringConnectionId", id)
	}
	return resourceIds, nil
}

func remotePeeringConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if remotePeeringConnectionResponse, ok := response.Response.(oci_core.GetRemotePeeringConnectionResponse); ok {
		return remotePeeringConnectionResponse.LifecycleState != oci_core.RemotePeeringConnectionLifecycleStateTerminated
	}
	return false
}

func remotePeeringConnectionSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.virtualNetworkClient().GetRemotePeeringConnection(context.Background(), oci_core.GetRemotePeeringConnectionRequest{
		RemotePeeringConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
