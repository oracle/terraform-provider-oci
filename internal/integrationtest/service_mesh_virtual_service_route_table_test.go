// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ServiceMeshVirtualServiceRouteTableRequiredOnlyResource = ServiceMeshVirtualServiceRouteTableResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service_route_table", "test_virtual_service_route_table", acctest.Required, acctest.Create, ServiceMeshVirtualServiceRouteTableRepresentation)

	ServiceMeshVirtualServiceRouteTableResourceConfig = ServiceMeshVirtualServiceRouteTableResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service_route_table", "test_virtual_service_route_table", acctest.Optional, acctest.Update, ServiceMeshVirtualServiceRouteTableRepresentation)

	ServiceMeshServiceMeshVirtualServiceRouteTableSingularDataSourceRepresentation = map[string]interface{}{
		"virtual_service_route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_virtual_service_route_table.test_virtual_service_route_table.id}`},
	}

	ServiceMeshServiceMeshVirtualServiceRouteTableDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_virtual_service_route_table.test_virtual_service_route_table.id}`},
		"name":               acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"virtual_service_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_virtual_service.virtual_service_1.id}`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshVirtualServiceRouteTableDataSourceFilterRepresentation}}
	ServiceMeshVirtualServiceRouteTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_service_mesh_virtual_service_route_table.test_virtual_service_route_table.id}`}},
	}

	ServiceMeshVirtualServiceRouteTableRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`},
		"route_rules":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshVirtualServiceRouteTableRouteRulesRepresentation},
		"virtual_service_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_virtual_service.virtual_service_1.id}`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"priority":           acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	ServiceMeshVirtualServiceRouteTableRouteRulesRepresentation = map[string]interface{}{
		"destinations":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshVirtualServiceRouteTableRouteRulesDestinationsRepresentation},
		"type":                  acctest.Representation{RepType: acctest.Required, Create: `TCP`, Update: `HTTP`},
		"is_grpc":               acctest.Representation{RepType: acctest.Optional, Update: `true`},
		"path":                  acctest.Representation{RepType: acctest.Optional, Update: `/path2`},
		"path_type":             acctest.Representation{RepType: acctest.Optional, Update: `PREFIX`},
		"request_timeout_in_ms": acctest.Representation{RepType: acctest.Optional, Update: `11`},
	}
	ServiceMeshVirtualServiceRouteTableRouteRulesDestinationsRepresentation = map[string]interface{}{
		"virtual_deployment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_virtual_deployment.virtual_deployment_1.id}`},
		"weight":                acctest.Representation{RepType: acctest.Required, Create: `100`},
		"port":                  acctest.Representation{RepType: acctest.Optional, Create: `910`, Update: `911`},
	}

	ServiceMeshVirtualServiceRouteTableResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "mesh1", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ServiceMeshMeshRepresentation, map[string]interface{}{
			"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_deployment", "virtual_deployment_1", acctest.Required, acctest.Create, ServiceMeshVirtualDeploymentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service", "virtual_service_1", acctest.Required, acctest.Create, ServiceMeshVirtualServiceRepresentation)
)

// issue-routing-tag: service_mesh/default
func TestServiceMeshVirtualServiceRouteTableResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceMeshVirtualServiceRouteTableResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	certificateAuthorityId := utils.GetEnvSettingWithBlankDefault("certificate_authority_id")
	certificateAuthorityIdVariableStr := fmt.Sprintf("variable \"certificate_authority_id\" { default = \"%s\" }\n", certificateAuthorityId)

	resourceName := "oci_service_mesh_virtual_service_route_table.test_virtual_service_route_table"
	datasourceName := "data.oci_service_mesh_virtual_service_route_tables.test_virtual_service_route_tables"
	singularDatasourceName := "data.oci_service_mesh_virtual_service_route_table.test_virtual_service_route_table"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ServiceMeshVirtualServiceRouteTableResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service_route_table", "test_virtual_service_route_table", acctest.Optional, acctest.Create, ServiceMeshVirtualServiceRouteTableRepresentation), "servicemesh", "virtualServiceRouteTable", t)

	acctest.ResourceTest(t, testAccCheckServiceMeshVirtualServiceRouteTableDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualServiceRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service_route_table", "test_virtual_service_route_table", acctest.Required, acctest.Create, ServiceMeshVirtualServiceRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
					"destinations.#": "1",
					"type":           "TCP",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_service_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualServiceRouteTableResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualServiceRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service_route_table", "test_virtual_service_route_table", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ServiceMeshVirtualServiceRouteTableRepresentation, map[string]interface{}{
					"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
					"destinations.#": "1",
					"type":           "TCP",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_service_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + ServiceMeshVirtualServiceRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service_route_table", "test_virtual_service_route_table", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ServiceMeshVirtualServiceRouteTableRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
					"destinations.#": "1",
					"type":           "TCP",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_service_id"),

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
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualServiceRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service_route_table", "test_virtual_service_route_table", acctest.Optional, acctest.Update, ServiceMeshVirtualServiceRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "priority", "11"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
					"destinations.#":        "1",
					"is_grpc":               "true",
					"path":                  "/path2",
					"path_type":             "PREFIX",
					"request_timeout_in_ms": "11",
					"type":                  "HTTP",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "virtual_service_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_virtual_service_route_tables", "test_virtual_service_route_tables", acctest.Optional, acctest.Update, ServiceMeshServiceMeshVirtualServiceRouteTableDataSourceRepresentation) +
				certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualServiceRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service_route_table", "test_virtual_service_route_table", acctest.Optional, acctest.Update, ServiceMeshVirtualServiceRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "virtual_service_id"),

				resource.TestCheckResourceAttr(datasourceName, "virtual_service_route_table_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "virtual_service_route_table_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_virtual_service_route_table", "test_virtual_service_route_table", acctest.Required, acctest.Create, ServiceMeshServiceMeshVirtualServiceRouteTableSingularDataSourceRepresentation) +
				certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshVirtualServiceRouteTableResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "virtual_service_route_table_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "priority", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "route_rules", map[string]string{
					"destinations.#":        "1",
					"is_grpc":               "true",
					"path":                  "/path2",
					"path_type":             "PREFIX",
					"request_timeout_in_ms": "11",
					"type":                  "HTTP",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ServiceMeshVirtualServiceRouteTableRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckServiceMeshVirtualServiceRouteTableDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ServiceMeshClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_mesh_virtual_service_route_table" {
			noResourceFound = false
			request := oci_service_mesh.GetVirtualServiceRouteTableRequest{}

			tmp := rs.Primary.ID
			request.VirtualServiceRouteTableId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")

			response, err := client.GetVirtualServiceRouteTable(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_service_mesh.VirtualServiceRouteTableLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ServiceMeshVirtualServiceRouteTable") {
		resource.AddTestSweepers("ServiceMeshVirtualServiceRouteTable", &resource.Sweeper{
			Name:         "ServiceMeshVirtualServiceRouteTable",
			Dependencies: acctest.DependencyGraph["virtualServiceRouteTable"],
			F:            sweepServiceMeshVirtualServiceRouteTableResource,
		})
	}
}

func sweepServiceMeshVirtualServiceRouteTableResource(compartment string) error {
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()
	virtualServiceRouteTableIds, err := getServiceMeshVirtualServiceRouteTableIds(compartment)
	if err != nil {
		return err
	}
	for _, virtualServiceRouteTableId := range virtualServiceRouteTableIds {
		if ok := acctest.SweeperDefaultResourceId[virtualServiceRouteTableId]; !ok {
			deleteVirtualServiceRouteTableRequest := oci_service_mesh.DeleteVirtualServiceRouteTableRequest{}

			deleteVirtualServiceRouteTableRequest.VirtualServiceRouteTableId = &virtualServiceRouteTableId

			deleteVirtualServiceRouteTableRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")
			_, error := serviceMeshClient.DeleteVirtualServiceRouteTable(context.Background(), deleteVirtualServiceRouteTableRequest)
			if error != nil {
				fmt.Printf("Error deleting VirtualServiceRouteTable %s %s, It is possible that the resource is already deleted. Please verify manually \n", virtualServiceRouteTableId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &virtualServiceRouteTableId, ServiceMeshVirtualServiceRouteTableSweepWaitCondition, time.Duration(3*time.Minute),
				ServiceMeshVirtualServiceRouteTableSweepResponseFetchOperation, "service_mesh", true)
		}
	}
	return nil
}

func getServiceMeshVirtualServiceRouteTableIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VirtualServiceRouteTableId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()

	listVirtualServiceRouteTablesRequest := oci_service_mesh.ListVirtualServiceRouteTablesRequest{}
	listVirtualServiceRouteTablesRequest.CompartmentId = &compartmentId
	listVirtualServiceRouteTablesRequest.LifecycleState = oci_service_mesh.VirtualServiceRouteTableLifecycleStateActive
	listVirtualServiceRouteTablesResponse, err := serviceMeshClient.ListVirtualServiceRouteTables(context.Background(), listVirtualServiceRouteTablesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting VirtualServiceRouteTable list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, virtualServiceRouteTable := range listVirtualServiceRouteTablesResponse.Items {
		id := *virtualServiceRouteTable.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VirtualServiceRouteTableId", id)
	}
	return resourceIds, nil
}

func ServiceMeshVirtualServiceRouteTableSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if virtualServiceRouteTableResponse, ok := response.Response.(oci_service_mesh.GetVirtualServiceRouteTableResponse); ok {
		return virtualServiceRouteTableResponse.LifecycleState != oci_service_mesh.VirtualServiceRouteTableLifecycleStateDeleted
	}
	return false
}

func ServiceMeshVirtualServiceRouteTableSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ServiceMeshClient().GetVirtualServiceRouteTable(context.Background(), oci_service_mesh.GetVirtualServiceRouteTableRequest{
		VirtualServiceRouteTableId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
