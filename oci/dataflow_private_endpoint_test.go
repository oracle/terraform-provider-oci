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
	"github.com/oracle/oci-go-sdk/v40/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/v40/dataflow"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PrivateEndpointRequiredOnlyResource = PrivateEndpointResourceDependencies +
		generateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", Required, Create, privateEndpointRepresentation)

	PrivateEndpointResourceConfig = PrivateEndpointResourceDependencies +
		generateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", Optional, Update, privateEndpointRepresentation)

	privateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"private_endpoint_id": Representation{repType: Required, create: `${oci_dataflow_private_endpoint.test_private_endpoint.id}`},
	}

	privateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id":           Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":             Representation{repType: Required, create: `pe_1234`, update: `displayName2`},
		"display_name_starts_with": Representation{repType: Optional, create: `displayNameStartsWith`},
		"owner_principal_id":       Representation{repType: Optional, create: `${var.user_ocid}`},
		"state":                    Representation{repType: Optional, create: `INACTIVE`},
		"filter":                   RepresentationGroup{Required, privateEndpointDataSourceFilterRepresentation}}
	privateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_dataflow_private_endpoint.test_private_endpoint.id}`}},
	}

	privateEndpointRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"dns_zones":      Representation{repType: Required, create: []string{`custpvtsubnet.oraclevcn.com`}, update: []string{`db.custpvtsubnet.oraclevcn.com`}},
		"subnet_id":      Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":   Representation{repType: Optional, create: `pe_1234`, update: `displayName2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"max_host_count": Representation{repType: Optional, create: `256`, update: `512`},
		"nsg_ids":        Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, update: []string{}},
	}

	PrivateEndpointResourceDependencies = generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

func TestDataflowPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)
	userId := getEnvSettingWithBlankDefault("user_ocid")
	userIdVariableStr := fmt.Sprintf("variable \"user_ocid\" { default = \"%s\" }\n", userId)

	resourceName := "oci_dataflow_private_endpoint.test_private_endpoint"
	datasourceName := "data.oci_dataflow_private_endpoints.test_private_endpoints"
	singularDatasourceName := "data.oci_dataflow_private_endpoint.test_private_endpoint"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+PrivateEndpointResourceDependencies+
		generateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", Optional, Create, privateEndpointRepresentation), "dataflow", "privateEndpoint", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDataflowPrivateEndpointDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + PrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", Required, Create, privateEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "dns_zones.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + PrivateEndpointResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + PrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", Optional, Create, privateEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + PrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", Optional, Create,
						representationCopyWithNewProperties(privateEndpointRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
				Config: config + compartmentIdVariableStr + PrivateEndpointResourceDependencies +
					generateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", Optional, Update, privateEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_dataflow_private_endpoints", "test_private_endpoints", Required, Update, privateEndpointDataSourceRepresentation) +
					compartmentIdVariableStr + PrivateEndpointResourceDependencies + userIdVariableStr +
					generateResourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", Optional, Update, privateEndpointRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

					resource.TestCheckResourceAttr(datasourceName, "private_endpoint_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "private_endpoint_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_dataflow_private_endpoint", "test_private_endpoint", Required, Create, privateEndpointSingularDataSourceRepresentation) +
					compartmentIdVariableStr + PrivateEndpointResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
		},
	})
}

func testAccCheckDataflowPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataFlowClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dataflow_private_endpoint" {
			noResourceFound = false
			request := oci_dataflow.GetPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.PrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dataflow")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DataflowPrivateEndpoint") {
		resource.AddTestSweepers("DataflowPrivateEndpoint", &resource.Sweeper{
			Name:         "DataflowPrivateEndpoint",
			Dependencies: DependencyGraph["privateEndpoint"],
			F:            sweepDataflowPrivateEndpointResource,
		})
	}
}

func sweepDataflowPrivateEndpointResource(compartment string) error {
	dataFlowClient := GetTestClients(&schema.ResourceData{}).dataFlowClient()
	privateEndpointIds, err := getPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, privateEndpointId := range privateEndpointIds {
		if ok := SweeperDefaultResourceId[privateEndpointId]; !ok {
			deletePrivateEndpointRequest := oci_dataflow.DeletePrivateEndpointRequest{}

			deletePrivateEndpointRequest.PrivateEndpointId = &privateEndpointId

			deletePrivateEndpointRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "dataflow")
			_, error := dataFlowClient.DeletePrivateEndpoint(context.Background(), deletePrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting PrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", privateEndpointId, error)
				continue
			}
			waitTillCondition(testAccProvider, &privateEndpointId, privateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				privateEndpointSweepResponseFetchOperation, "dataflow", true)
		}
	}
	return nil
}

func getPrivateEndpointIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "PrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataFlowClient := GetTestClients(&schema.ResourceData{}).dataFlowClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "PrivateEndpointId", id)
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

func privateEndpointSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataFlowClient().GetPrivateEndpoint(context.Background(), oci_dataflow.GetPrivateEndpointRequest{
		PrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
