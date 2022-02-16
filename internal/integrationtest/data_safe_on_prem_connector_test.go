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
	oci_data_safe "github.com/oracle/oci-go-sdk/v58/datasafe"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	OnPremConnectorRequiredOnlyResource = OnPremConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", acctest.Required, acctest.Create, onPremConnectorRepresentation)

	OnPremConnectorResourceConfig = OnPremConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", acctest.Optional, acctest.Update, onPremConnectorRepresentation)

	onPremConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"on_prem_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_on_prem_connector.test_on_prem_connector.id}`},
	}

	onPremConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":                      acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree":         acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"display_name":                      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"on_prem_connector_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_on_prem_connector.test_on_prem_connector.id}`},
		"on_prem_connector_lifecycle_state": acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"filter":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: onPremConnectorDataSourceFilterRepresentation}}
	onPremConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_on_prem_connector.test_on_prem_connector.id}`}},
	}

	onPremConnectorRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDS},
	}

	ignoreDefinedTagsDS = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	OnPremConnectorResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeOnPremConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeOnPremConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_safe_on_prem_connector.test_on_prem_connector"
	datasourceName := "data.oci_data_safe_on_prem_connectors.test_on_prem_connectors"
	singularDatasourceName := "data.oci_data_safe_on_prem_connector.test_on_prem_connector"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDataSafeOnPremConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OnPremConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", acctest.Required, acctest.Create, onPremConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OnPremConnectorResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OnPremConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", acctest.Optional, acctest.Create, onPremConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OnPremConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(onPremConnectorRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + OnPremConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", acctest.Optional, acctest.Update, onPremConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_on_prem_connectors", "test_on_prem_connectors", acctest.Optional, acctest.Update, onPremConnectorDataSourceRepresentation) +
				compartmentIdVariableStr + OnPremConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", acctest.Optional, acctest.Update, onPremConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "on_prem_connector_id"),
				resource.TestCheckResourceAttr(datasourceName, "on_prem_connector_lifecycle_state", "INACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "on_prem_connectors.0.created_version"),
				resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "on_prem_connectors.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "on_prem_connectors.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "on_prem_connectors.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "on_prem_connectors.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_on_prem_connector", "test_on_prem_connector", acctest.Required, acctest.Create, onPremConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OnPremConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "on_prem_connector_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_version"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + OnPremConnectorResourceConfig,
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

func testAccCheckDataSafeOnPremConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_on_prem_connector" {
			noResourceFound = false
			request := oci_data_safe.GetOnPremConnectorRequest{}

			tmp := rs.Primary.ID
			request.OnPremConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetOnPremConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.OnPremConnectorLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeOnPremConnector") {
		resource.AddTestSweepers("DataSafeOnPremConnector", &resource.Sweeper{
			Name:         "DataSafeOnPremConnector",
			Dependencies: acctest.DependencyGraph["onPremConnector"],
			F:            sweepDataSafeOnPremConnectorResource,
		})
	}
}

func sweepDataSafeOnPremConnectorResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	onPremConnectorIds, err := getOnPremConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, onPremConnectorId := range onPremConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[onPremConnectorId]; !ok {
			deleteOnPremConnectorRequest := oci_data_safe.DeleteOnPremConnectorRequest{}

			deleteOnPremConnectorRequest.OnPremConnectorId = &onPremConnectorId

			deleteOnPremConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteOnPremConnector(context.Background(), deleteOnPremConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting OnPremConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", onPremConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &onPremConnectorId, onPremConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				onPremConnectorSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getOnPremConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OnPremConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listOnPremConnectorsRequest := oci_data_safe.ListOnPremConnectorsRequest{}
	listOnPremConnectorsRequest.CompartmentId = &compartmentId
	listOnPremConnectorsResponse, err := dataSafeClient.ListOnPremConnectors(context.Background(), listOnPremConnectorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OnPremConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, onPremConnector := range listOnPremConnectorsResponse.Items {
		id := *onPremConnector.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OnPremConnectorId", id)
	}
	return resourceIds, nil
}

func onPremConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if onPremConnectorResponse, ok := response.Response.(oci_data_safe.GetOnPremConnectorResponse); ok {
		return string(onPremConnectorResponse.LifecycleState) != string(oci_data_safe.OnPremConnectorLifecycleStateDeleted)
	}
	return false
}

func onPremConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetOnPremConnector(context.Background(), oci_data_safe.GetOnPremConnectorRequest{
		OnPremConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
