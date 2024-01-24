// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OpsiOperationsInsightsPrivateEndpointRequiredOnlyResource = OpsiOperationsInsightsPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_private_endpoint", "test_operations_insights_private_endpoint", acctest.Required, acctest.Create, OpsiOperationsInsightsPrivateEndpointRepresentation)

	OpsiOperationsInsightsPrivateEndpointResourceConfig = OpsiOperationsInsightsPrivateEndpointResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_private_endpoint", "test_operations_insights_private_endpoint", acctest.Optional, acctest.Update, OpsiOperationsInsightsPrivateEndpointRepresentation)

	OpsiOpsiOperationsInsightsPrivateEndpointSingularDataSourceRepresentation = map[string]interface{}{
		"operations_insights_private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_operations_insights_private_endpoint.test_operations_insights_private_endpoint.id}`},
	}

	OpsiOpsiOperationsInsightsPrivateEndpointDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `TerraformPe`, Update: `TerraformPe2`},
		"is_used_for_rac_dbs":       acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"opsi_private_endpoint_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_operations_insights_private_endpoint.test_operations_insights_private_endpoint.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"vcn_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.vcn_id}`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiOperationsInsightsPrivateEndpointDataSourceFilterRepresentation}}
	OpsiOperationsInsightsPrivateEndpointDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_operations_insights_private_endpoint.test_operations_insights_private_endpoint.id}`}},
	}

	OpsiOperationsInsightsPrivateEndpointRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `TerraformPe`, Update: `TerraformPe2`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"is_used_for_rac_dbs": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"subnet_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.subnet_id}`},
		"vcn_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.vcn_id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesOperationsInsightsPrivateEndpointRepresentation},
	}

	ignoreChangesOperationsInsightsPrivateEndpointRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	OpsiOperationsInsightsPrivateEndpointResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiOperationsInsightsPrivateEndpointResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiOperationsInsightsPrivateEndpointResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	log.Printf("in test function (%s)", "TestOpsiOperationsInsightsPrivateEndpointResource_basic")

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	vcnId := utils.GetEnvSettingWithBlankDefault("vcn_id")
	vcnIdVariableStr := fmt.Sprintf("variable \"vcn_id\" { default = \"%s\" }\n", vcnId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_opsi_operations_insights_private_endpoint.test_operations_insights_private_endpoint"
	datasourceName := "data.oci_opsi_operations_insights_private_endpoints.test_operations_insights_private_endpoints"
	singularDatasourceName := "data.oci_opsi_operations_insights_private_endpoint.test_operations_insights_private_endpoint"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OpsiOperationsInsightsPrivateEndpointResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_private_endpoint", "test_operations_insights_private_endpoint", acctest.Optional, acctest.Create, OpsiOperationsInsightsPrivateEndpointRepresentation), "operationsinsights", "operationsInsightsPrivateEndpoint", t)

	acctest.ResourceTest(t, testAccCheckOpsiOperationsInsightsPrivateEndpointDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + OpsiOperationsInsightsPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_private_endpoint", "test_operations_insights_private_endpoint", acctest.Required, acctest.Create, OpsiOperationsInsightsPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TerraformPe"),
				resource.TestCheckResourceAttr(resourceName, "is_used_for_rac_dbs", "false"),
				resource.TestCheckResourceAttr(resourceName, "subnet_id", subnetId),
				resource.TestCheckResourceAttr(resourceName, "vcn_id", vcnId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + OpsiOperationsInsightsPrivateEndpointResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + OpsiOperationsInsightsPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_private_endpoint", "test_operations_insights_private_endpoint", acctest.Optional, acctest.Create, OpsiOperationsInsightsPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TerraformPe"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_used_for_rac_dbs", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + vcnIdVariableStr + subnetIdVariableStr + OpsiOperationsInsightsPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_private_endpoint", "test_operations_insights_private_endpoint", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OpsiOperationsInsightsPrivateEndpointRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TerraformPe"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_used_for_rac_dbs", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
			Config: config + compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + OpsiOperationsInsightsPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_private_endpoint", "test_operations_insights_private_endpoint", acctest.Optional, acctest.Update, OpsiOperationsInsightsPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TerraformPe2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_used_for_rac_dbs", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_operations_insights_private_endpoints", "test_operations_insights_private_endpoints", acctest.Optional, acctest.Update, OpsiOpsiOperationsInsightsPrivateEndpointDataSourceRepresentation) +
				compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + OpsiOperationsInsightsPrivateEndpointResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_private_endpoint", "test_operations_insights_private_endpoint", acctest.Optional, acctest.Update, OpsiOperationsInsightsPrivateEndpointRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TerraformPe2"),
				resource.TestCheckResourceAttr(datasourceName, "is_used_for_rac_dbs", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "opsi_private_endpoint_id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vcn_id"),

				resource.TestCheckResourceAttr(datasourceName, "operations_insights_private_endpoint_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "operations_insights_private_endpoint_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_operations_insights_private_endpoint", "test_operations_insights_private_endpoint", acctest.Required, acctest.Create, OpsiOpsiOperationsInsightsPrivateEndpointSingularDataSourceRepresentation) +
				compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + OpsiOperationsInsightsPrivateEndpointResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operations_insights_private_endpoint_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TerraformPe2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_used_for_rac_dbs", "false"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "private_endpoint_status_details"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + vcnIdVariableStr + subnetIdVariableStr + OpsiOperationsInsightsPrivateEndpointResourceConfig,
		},
		// verify resource import
		{
			Config:                  config + OpsiOperationsInsightsPrivateEndpointRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOpsiOperationsInsightsPrivateEndpointDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_operations_insights_private_endpoint" {
			noResourceFound = false
			request := oci_opsi.GetOperationsInsightsPrivateEndpointRequest{}

			tmp := rs.Primary.ID
			request.OperationsInsightsPrivateEndpointId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetOperationsInsightsPrivateEndpoint(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.OperationsInsightsPrivateEndpointLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpsiOperationsInsightsPrivateEndpoint") {
		resource.AddTestSweepers("OpsiOperationsInsightsPrivateEndpoint", &resource.Sweeper{
			Name:         "OpsiOperationsInsightsPrivateEndpoint",
			Dependencies: acctest.DependencyGraph["operationsInsightsPrivateEndpoint"],
			F:            sweepOpsiOperationsInsightsPrivateEndpointResource,
		})
	}
}

func sweepOpsiOperationsInsightsPrivateEndpointResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	operationsInsightsPrivateEndpointIds, err := getOpsiOperationsInsightsPrivateEndpointIds(compartment)
	if err != nil {
		return err
	}
	for _, operationsInsightsPrivateEndpointId := range operationsInsightsPrivateEndpointIds {
		if ok := acctest.SweeperDefaultResourceId[operationsInsightsPrivateEndpointId]; !ok {
			deleteOperationsInsightsPrivateEndpointRequest := oci_opsi.DeleteOperationsInsightsPrivateEndpointRequest{}

			deleteOperationsInsightsPrivateEndpointRequest.OperationsInsightsPrivateEndpointId = &operationsInsightsPrivateEndpointId

			deleteOperationsInsightsPrivateEndpointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteOperationsInsightsPrivateEndpoint(context.Background(), deleteOperationsInsightsPrivateEndpointRequest)
			if error != nil {
				fmt.Printf("Error deleting OperationsInsightsPrivateEndpoint %s %s, It is possible that the resource is already deleted. Please verify manually \n", operationsInsightsPrivateEndpointId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &operationsInsightsPrivateEndpointId, OpsiOperationsInsightsPrivateEndpointSweepWaitCondition, time.Duration(3*time.Minute),
				OpsiOperationsInsightsPrivateEndpointSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getOpsiOperationsInsightsPrivateEndpointIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OperationsInsightsPrivateEndpointId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listOperationsInsightsPrivateEndpointsRequest := oci_opsi.ListOperationsInsightsPrivateEndpointsRequest{}
	listOperationsInsightsPrivateEndpointsRequest.CompartmentId = &compartmentId
	listOperationsInsightsPrivateEndpointsRequest.LifecycleState =
		[]oci_opsi.OperationsInsightsPrivateEndpointLifecycleStateEnum{oci_opsi.OperationsInsightsPrivateEndpointLifecycleStateActive}
	listOperationsInsightsPrivateEndpointsResponse, err := operationsInsightsClient.ListOperationsInsightsPrivateEndpoints(context.Background(), listOperationsInsightsPrivateEndpointsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OperationsInsightsPrivateEndpoint list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, operationsInsightsPrivateEndpoint := range listOperationsInsightsPrivateEndpointsResponse.Items {
		id := *operationsInsightsPrivateEndpoint.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OperationsInsightsPrivateEndpointId", id)
	}
	return resourceIds, nil
}

func OpsiOperationsInsightsPrivateEndpointSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if operationsInsightsPrivateEndpointResponse, ok := response.Response.(oci_opsi.GetOperationsInsightsPrivateEndpointResponse); ok {
		return operationsInsightsPrivateEndpointResponse.LifecycleState != oci_opsi.OperationsInsightsPrivateEndpointLifecycleStateDeleted
	}
	return false
}

func OpsiOperationsInsightsPrivateEndpointSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetOperationsInsightsPrivateEndpoint(context.Background(), oci_opsi.GetOperationsInsightsPrivateEndpointRequest{
		OperationsInsightsPrivateEndpointId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
