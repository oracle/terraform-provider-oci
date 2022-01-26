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
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v56/databasetools"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DatabaseToolsPrivateEndpointRequiredOnlyResource = DatabaseToolsPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_database_tools_private_endpoint", acctest.Required, acctest.Create, databaseToolsPrivateEndpointRepresentation)

	DatabaseToolsPrivateEndpointResourceConfig = DatabaseToolsPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_database_tools_private_endpoint", acctest.Optional, acctest.Update, databaseToolsPrivateEndpointRepresentation)

	databaseToolsPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_private_endpoint.test_database_tools_private_endpoint.id}`},
	}

	databaseToolsPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `MyPE`, Update: `displayName2`},
		"endpoint_service_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_tools_database_tools_endpoint_services.test_database_tools_endpoint_services.database_tools_endpoint_service_collection.0.items.0.id}`},
		"state":               acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"subnet_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.id}`},
		"filter":              acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseToolsPrivateEndpointDataSourceFilterRepresentation}}
	databaseToolsPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_private_endpoint.test_database_tools_private_endpoint.id}`}},
	}

	databaseToolsPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `MyPE`, Update: `displayName2`},
		"endpoint_service_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_tools_database_tools_endpoint_services.test_database_tools_endpoint_services.database_tools_endpoint_service_collection.0.items.0.id}`},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `Private Endpoint for mySubnet`, Update: `description2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":             acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"private_endpoint_ip": acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.4`},
		//"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesDatabaseToolsPrivateEndpointRepresentation},
	}

	//ignoreChangesDatabaseToolsPrivateEndpointRepresentation = map[string]interface{}{   // On R1 only
	//	"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	//}

	DatabaseToolsPrivateEndpointResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestDatabaseToolsDatabaseToolsPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_private_endpoint.test_database_tools_private_endpoint"
	datasourceName := "data.oci_database_tools_database_tools_private_endpoints.test_database_tools_private_endpoints"
	singularDatasourceName := "data.oci_database_tools_database_tools_private_endpoint.test_database_tools_private_endpoint"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseToolsPrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_database_tools_private_endpoint", acctest.Optional, acctest.Create, databaseToolsPrivateEndpointRepresentation), "databasetools", "databaseToolsPrivateEndpoint", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseToolsDatabaseToolsPrivateEndpointDestroy,
		Steps: []resource.TestStep{
			// 0. verify create
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsPrivateEndpointResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_service", "test_database_tools_endpoint_service", acctest.Required, acctest.Create, databaseToolsEndpointServiceSingularDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_database_tools_private_endpoint", acctest.Required, acctest.Create, databaseToolsPrivateEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyPE"),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_service_id"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// 1. delete before next create
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsPrivateEndpointResourceDependencies,
			},
			// 2. verify create with optionals
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsPrivateEndpointResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_service", "test_database_tools_endpoint_service", acctest.Required, acctest.Create, databaseToolsEndpointServiceSingularDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_database_tools_private_endpoint", acctest.Optional, acctest.Create, databaseToolsPrivateEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "Private Endpoint for mySubnet"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyPE"),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_service_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "private_endpoint_ip", "10.0.0.4"),
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

			// 3. verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseToolsPrivateEndpointResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_service", "test_database_tools_endpoint_service", acctest.Required, acctest.Create, databaseToolsEndpointServiceSingularDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_database_tools_private_endpoint", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(databaseToolsPrivateEndpointRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "description", "Private Endpoint for mySubnet"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "MyPE"),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_service_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "private_endpoint_ip", "10.0.0.4"),
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

			// 4. verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsPrivateEndpointResourceDependencies +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_service", "test_database_tools_endpoint_service", acctest.Required, acctest.Create, databaseToolsEndpointServiceSingularDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_database_tools_private_endpoint", acctest.Optional, acctest.Update, databaseToolsPrivateEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(resourceName, "endpoint_service_id"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "private_endpoint_ip", "10.0.0.4"),
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
			// 5. verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_service", "test_database_tools_endpoint_service", acctest.Required, acctest.Create, databaseToolsEndpointServiceSingularDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoints", "test_database_tools_private_endpoints", acctest.Optional, acctest.Update, databaseToolsPrivateEndpointDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseToolsPrivateEndpointResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_database_tools_private_endpoint", acctest.Optional, acctest.Update, databaseToolsPrivateEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "endpoint_service_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "subnet_id"),

					resource.TestCheckResourceAttr(datasourceName, "database_tools_private_endpoint_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "database_tools_private_endpoint_collection.0.items.#", "1"),
				),
			},
			// 6. verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_private_endpoint", "test_database_tools_private_endpoint", acctest.Required, acctest.Create, databaseToolsPrivateEndpointSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DatabaseToolsPrivateEndpointResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_private_endpoint_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "additional_fqdns.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "private_endpoint_ip", "10.0.0.4"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_vnic_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "reverse_connection_configuration.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vcn_id"),
				),
			},
			// 7. remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + DatabaseToolsPrivateEndpointResourceConfig +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_endpoint_services", "test_database_tools_endpoint_services", acctest.Required, acctest.Create, databaseToolsEndpointServiceDataSourceRepresentation),
			},
			// 8. verify resource import
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

func testAccCheckDatabaseToolsDatabaseToolsPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_database_tools_private_endpoint" {
			noResourceFound = false
			request := oci_database_tools.GetDatabaseToolsPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.DatabaseToolsPrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")

			response, err := client.GetDatabaseToolsPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_tools.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseToolsDatabaseToolsPrivateEndpoint") {
		resource.AddTestSweepers("DatabaseToolsDatabaseToolsPrivateEndpoint", &resource.Sweeper{
			Name:         "DatabaseToolsDatabaseToolsPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["databaseToolsPrivateEndpoint"],
			F:            sweepDatabaseToolsDatabaseToolsPrivateEndpointResource,
		})
	}
}

func sweepDatabaseToolsDatabaseToolsPrivateEndpointResource(compartment string) error {
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()
	databaseToolsPrivateEndpointIds, err := getDatabaseToolsPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsPrivateEndpointId := range databaseToolsPrivateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsPrivateEndpointId]; !ok {
			deleteDatabaseToolsPrivateEndpointRequest := oci_database_tools.DeleteDatabaseToolsPrivateEndpointRequest{}

			deleteDatabaseToolsPrivateEndpointRequest.DatabaseToolsPrivateEndpointId = &databaseToolsPrivateEndpointId

			deleteDatabaseToolsPrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")
			_, error := databaseToolsClient.DeleteDatabaseToolsPrivateEndpoint(context.Background(), deleteDatabaseToolsPrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsPrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsPrivateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseToolsPrivateEndpointId, databaseToolsPrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				databaseToolsPrivateEndpointSweepResponseFetchOperation, "database_tools", true)
		}
	}
	return nil
}

func getDatabaseToolsPrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsPrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()

	listDatabaseToolsPrivateEndpointsRequest := oci_database_tools.ListDatabaseToolsPrivateEndpointsRequest{}
	listDatabaseToolsPrivateEndpointsRequest.CompartmentId = &compartmentId
	listDatabaseToolsPrivateEndpointsRequest.LifecycleState = oci_database_tools.ListDatabaseToolsPrivateEndpointsLifecycleStateActive
	listDatabaseToolsPrivateEndpointsResponse, err := databaseToolsClient.ListDatabaseToolsPrivateEndpoints(context.Background(), listDatabaseToolsPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseToolsPrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseToolsPrivateEndpoint := range listDatabaseToolsPrivateEndpointsResponse.Items {
		id := *databaseToolsPrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsPrivateEndpointId", id)
	}
	return resourceIds, nil
}

func databaseToolsPrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseToolsPrivateEndpointResponse, ok := response.Response.(oci_database_tools.GetDatabaseToolsPrivateEndpointResponse); ok {
		return databaseToolsPrivateEndpointResponse.LifecycleState != oci_database_tools.LifecycleStateDeleted
	}
	return false
}

func databaseToolsPrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseToolsClient().GetDatabaseToolsPrivateEndpoint(context.Background(), oci_database_tools.GetDatabaseToolsPrivateEndpointRequest{
		DatabaseToolsPrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
