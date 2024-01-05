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
	ServiceMeshIngressGatewayRouteTableRequiredOnlyResource = ServiceMeshIngressGatewayRouteTableResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway_route_table", "test_ingress_gateway_route_table", acctest.Required, acctest.Create, ServiceMeshIngressGatewayRouteTableRepresentation)

	ServiceMeshIngressGatewayRouteTableResourceConfig = ServiceMeshIngressGatewayRouteTableResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway_route_table", "test_ingress_gateway_route_table", acctest.Optional, acctest.Update, ServiceMeshIngressGatewayRouteTableRepresentation)

	ServiceMeshServiceMeshIngressGatewayRouteTableSingularDataSourceRepresentation = map[string]interface{}{
		"ingress_gateway_route_table_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_ingress_gateway_route_table.test_ingress_gateway_route_table.id}`},
	}

	ServiceMeshServiceMeshIngressGatewayRouteTableDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_ingress_gateway_route_table.test_ingress_gateway_route_table.id}`},
		"ingress_gateway_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_service_mesh_ingress_gateway.ingress_gateway_1.id}`},
		"name":               acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshIngressGatewayRouteTableDataSourceFilterRepresentation}}
	ServiceMeshIngressGatewayRouteTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_service_mesh_ingress_gateway_route_table.test_ingress_gateway_route_table.id}`}},
	}

	ServiceMeshIngressGatewayRouteTableRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ingress_gateway_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_ingress_gateway.ingress_gateway_1.id}`},
		"name":               acctest.Representation{RepType: acctest.Required, Create: `name`},
		"route_rules":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshIngressGatewayRouteTableRouteRulesRepresentation},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":        acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"priority":           acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	ServiceMeshIngressGatewayRouteTableRouteRulesRepresentation = map[string]interface{}{
		"destinations":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ServiceMeshIngressGatewayRouteTableRouteRulesDestinationsRepresentation},
		"type":                    acctest.Representation{RepType: acctest.Required, Create: `TLS_PASSTHROUGH`, Update: `HTTP`},
		"ingress_gateway_host":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: ServiceMeshIngressGatewayRouteTableRouteRulesIngressGatewayHostRepresentation},
		"is_grpc":                 acctest.Representation{RepType: acctest.Optional, Update: `true`},
		"is_host_rewrite_enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`},
		"is_path_rewrite_enabled": acctest.Representation{RepType: acctest.Optional, Update: `true`},
		"path":                    acctest.Representation{RepType: acctest.Optional, Update: `/path2`},
		"path_type":               acctest.Representation{RepType: acctest.Optional, Update: `PREFIX`},
		"request_timeout_in_ms":   acctest.Representation{RepType: acctest.Optional, Update: `11`},
	}
	ServiceMeshIngressGatewayRouteTableRouteRulesDestinationsRepresentation = map[string]interface{}{
		"virtual_service_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_service_mesh_virtual_service.virtual_service_1.id}`},
		"port":               acctest.Representation{RepType: acctest.Optional, Create: `7010`, Update: `7011`},
		"weight":             acctest.Representation{RepType: acctest.Optional, Create: `100`},
	}
	ServiceMeshIngressGatewayRouteTableRouteRulesIngressGatewayHostRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `name`},
		"port": acctest.Representation{RepType: acctest.Optional, Create: `8010`, Update: `8011`},
	}

	ServiceMeshIngressGatewayRouteTableResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_mesh", "mesh1", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ServiceMeshMeshRepresentation, map[string]interface{}{
			"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway", "ingress_gateway_1", acctest.Required, acctest.Create, ServiceMeshIngressGatewayRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_virtual_service", "virtual_service_1", acctest.Required, acctest.Create, ServiceMeshVirtualServiceRepresentation)
)

// issue-routing-tag: service_mesh/default
func TestServiceMeshIngressGatewayRouteTableResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestServiceMeshIngressGatewayRouteTableResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	certificateAuthorityId := utils.GetEnvSettingWithBlankDefault("certificate_authority_id")
	certificateAuthorityIdVariableStr := fmt.Sprintf("variable \"certificate_authority_id\" { default = \"%s\" }\n", certificateAuthorityId)

	resourceName := "oci_service_mesh_ingress_gateway_route_table.test_ingress_gateway_route_table"
	datasourceName := "data.oci_service_mesh_ingress_gateway_route_tables.test_ingress_gateway_route_tables"
	singularDatasourceName := "data.oci_service_mesh_ingress_gateway_route_table.test_ingress_gateway_route_table"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+certificateAuthorityIdVariableStr+compartmentIdVariableStr+ServiceMeshIngressGatewayRouteTableResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway_route_table", "test_ingress_gateway_route_table", acctest.Optional, acctest.Create, ServiceMeshIngressGatewayRouteTableRepresentation), "servicemesh", "ingressGatewayRouteTable", t)

	acctest.ResourceTest(t, testAccCheckServiceMeshIngressGatewayRouteTableDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshIngressGatewayRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway_route_table", "test_ingress_gateway_route_table", acctest.Required, acctest.Create, ServiceMeshIngressGatewayRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "ingress_gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
					"destinations.#": "1",
					"type":           "TLS_PASSTHROUGH",
				},
					[]string{}),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshIngressGatewayRouteTableResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshIngressGatewayRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway_route_table", "test_ingress_gateway_route_table", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithNewProperties(ServiceMeshIngressGatewayRouteTableRepresentation, map[string]interface{}{
					"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMeshDefinedTagsChangesRepresentation}})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ingress_gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
					"destinations.#":         "1",
					"ingress_gateway_host.#": "1",
					"type":                   "TLS_PASSTHROUGH",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr + ServiceMeshIngressGatewayRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway_route_table", "test_ingress_gateway_route_table", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ServiceMeshIngressGatewayRouteTableRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ingress_gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "priority", "10"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
					"destinations.#":         "1",
					"ingress_gateway_host.#": "1",
					"type":                   "TLS_PASSTHROUGH",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshIngressGatewayRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway_route_table", "test_ingress_gateway_route_table", acctest.Optional, acctest.Update, ServiceMeshIngressGatewayRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ingress_gateway_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "priority", "11"),
				resource.TestCheckResourceAttr(resourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "route_rules", map[string]string{
					"destinations.#":          "1",
					"ingress_gateway_host.#":  "1",
					"is_grpc":                 "true",
					"is_host_rewrite_enabled": "true",
					"is_path_rewrite_enabled": "true",
					"path":                    "/path2",
					"path_type":               "PREFIX",
					"request_timeout_in_ms":   "11",
					"type":                    "HTTP",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_ingress_gateway_route_tables", "test_ingress_gateway_route_tables", acctest.Optional, acctest.Update, ServiceMeshServiceMeshIngressGatewayRouteTableDataSourceRepresentation) +
				certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshIngressGatewayRouteTableResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_service_mesh_ingress_gateway_route_table", "test_ingress_gateway_route_table", acctest.Optional, acctest.Update, ServiceMeshIngressGatewayRouteTableRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ingress_gateway_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "ingress_gateway_route_table_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ingress_gateway_route_table_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_service_mesh_ingress_gateway_route_table", "test_ingress_gateway_route_table", acctest.Required, acctest.Create, ServiceMeshServiceMeshIngressGatewayRouteTableSingularDataSourceRepresentation) +
				certificateAuthorityIdVariableStr + compartmentIdVariableStr + ServiceMeshIngressGatewayRouteTableResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ingress_gateway_route_table_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "priority", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "route_rules.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "route_rules", map[string]string{
					"destinations.#":          "1",
					"ingress_gateway_host.#":  "1",
					"is_grpc":                 "true",
					"is_host_rewrite_enabled": "true",
					"is_path_rewrite_enabled": "true",
					"path":                    "/path2",
					"path_type":               "PREFIX",
					"request_timeout_in_ms":   "11",
					"type":                    "HTTP",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ServiceMeshIngressGatewayRouteTableRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckServiceMeshIngressGatewayRouteTableDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ServiceMeshClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_service_mesh_ingress_gateway_route_table" {
			noResourceFound = false
			request := oci_service_mesh.GetIngressGatewayRouteTableRequest{}

			tmp := rs.Primary.ID
			request.IngressGatewayRouteTableId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")

			response, err := client.GetIngressGatewayRouteTable(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_service_mesh.IngressGatewayRouteTableLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ServiceMeshIngressGatewayRouteTable") {
		resource.AddTestSweepers("ServiceMeshIngressGatewayRouteTable", &resource.Sweeper{
			Name:         "ServiceMeshIngressGatewayRouteTable",
			Dependencies: acctest.DependencyGraph["ingressGatewayRouteTable"],
			F:            sweepServiceMeshIngressGatewayRouteTableResource,
		})
	}
}

func sweepServiceMeshIngressGatewayRouteTableResource(compartment string) error {
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()
	ingressGatewayRouteTableIds, err := getServiceMeshIngressGatewayRouteTableIds(compartment)
	if err != nil {
		return err
	}
	for _, ingressGatewayRouteTableId := range ingressGatewayRouteTableIds {
		if ok := acctest.SweeperDefaultResourceId[ingressGatewayRouteTableId]; !ok {
			deleteIngressGatewayRouteTableRequest := oci_service_mesh.DeleteIngressGatewayRouteTableRequest{}

			deleteIngressGatewayRouteTableRequest.IngressGatewayRouteTableId = &ingressGatewayRouteTableId

			deleteIngressGatewayRouteTableRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "service_mesh")
			_, error := serviceMeshClient.DeleteIngressGatewayRouteTable(context.Background(), deleteIngressGatewayRouteTableRequest)
			if error != nil {
				fmt.Printf("Error deleting IngressGatewayRouteTable %s %s, It is possible that the resource is already deleted. Please verify manually \n", ingressGatewayRouteTableId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &ingressGatewayRouteTableId, ServiceMeshIngressGatewayRouteTableSweepWaitCondition, time.Duration(3*time.Minute),
				ServiceMeshIngressGatewayRouteTableSweepResponseFetchOperation, "service_mesh", true)
		}
	}
	return nil
}

func getServiceMeshIngressGatewayRouteTableIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "IngressGatewayRouteTableId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceMeshClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceMeshClient()

	listIngressGatewayRouteTablesRequest := oci_service_mesh.ListIngressGatewayRouteTablesRequest{}
	listIngressGatewayRouteTablesRequest.CompartmentId = &compartmentId
	listIngressGatewayRouteTablesRequest.LifecycleState = oci_service_mesh.IngressGatewayRouteTableLifecycleStateActive
	listIngressGatewayRouteTablesResponse, err := serviceMeshClient.ListIngressGatewayRouteTables(context.Background(), listIngressGatewayRouteTablesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting IngressGatewayRouteTable list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ingressGatewayRouteTable := range listIngressGatewayRouteTablesResponse.Items {
		id := *ingressGatewayRouteTable.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "IngressGatewayRouteTableId", id)
	}
	return resourceIds, nil
}

func ServiceMeshIngressGatewayRouteTableSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ingressGatewayRouteTableResponse, ok := response.Response.(oci_service_mesh.GetIngressGatewayRouteTableResponse); ok {
		return ingressGatewayRouteTableResponse.LifecycleState != oci_service_mesh.IngressGatewayRouteTableLifecycleStateDeleted
	}
	return false
}

func ServiceMeshIngressGatewayRouteTableSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ServiceMeshClient().GetIngressGatewayRouteTable(context.Background(), oci_service_mesh.GetIngressGatewayRouteTableRequest{
		IngressGatewayRouteTableId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
