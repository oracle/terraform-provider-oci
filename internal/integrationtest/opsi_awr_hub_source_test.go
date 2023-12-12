// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OpsiAwrHubSourceRequiredOnlyResource = OpsiAwrHubSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source", "test_awr_hub_source", acctest.Required, acctest.Create, OpsiAwrHubSourceRepresentation)

	OpsiAwrHubSourceResourceConfig = OpsiAwrHubSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source", "test_awr_hub_source", acctest.Optional, acctest.Update, OpsiAwrHubSourceRepresentation)

	OpsiAwrHubSourceSingularDataSourceRepresentation = map[string]interface{}{
		"awr_hub_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_awr_hub_source.test_awr_hub_source.id}`},
	}

	OpsiAwrHubSourceDataSourceRepresentation = map[string]interface{}{
		"awr_hub_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"awr_hub_source_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_awr_hub_source.test_awr_hub_source.id}`},
		"compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":              acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"source_type":       acctest.Representation{RepType: acctest.Optional, Create: []string{`ADW_S`}},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"status":            acctest.Representation{RepType: acctest.Optional, Create: []string{`NOT_REGISTERED`}},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiAwrHubSourceDataSourceFilterRepresentation}}

	OpsiAwrHubSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_awr_hub_source.test_awr_hub_source.id}`}},
	}

	OpsiAwrHubSourceRepresentation = map[string]interface{}{
		"awr_hub_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"type":                   acctest.Representation{RepType: acctest.Required, Create: `ADW_S`, Update: `ATP_S`},
		"associated_opsi_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.associated_opsi_id}`},
		"associated_resource_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.associated_resource_id}`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	OpsiAwrHubSourceResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Required, acctest.Create, OpsiAwrHubRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiAwrHubSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiAwrHubSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	associatedOpsiId := utils.GetEnvSettingWithBlankDefault("associated_opsi_id")
	associatedOpsiIdVariableStr := fmt.Sprintf("variable \"associated_opsi_id\" { default = \"%s\" }\n", associatedOpsiId)

	associatedResourceId := utils.GetEnvSettingWithBlankDefault("associated_resource_id")
	associatedResourceIdVariableStr := fmt.Sprintf("variable \"associated_resource_id\" { default = \"%s\" }\n", associatedResourceId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_opsi_awr_hub_source.test_awr_hub_source"
	datasourceName := "data.oci_opsi_awr_hub_sources.test_awr_hub_sources"
	singularDatasourceName := "data.oci_opsi_awr_hub_source.test_awr_hub_source"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+associatedOpsiIdVariableStr+associatedResourceIdVariableStr+OpsiAwrHubSourceResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source", "test_awr_hub_source", acctest.Optional, acctest.Create, OpsiAwrHubSourceRepresentation), "operationsinsights", "awrHubSource", t)

	acctest.ResourceTest(t, testAccCheckOpsiAwrHubSourceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OpsiAwrHubSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source", "test_awr_hub_source", acctest.Required, acctest.Create, OpsiAwrHubSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "awr_hub_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "type", "ADW_S"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + associatedOpsiIdVariableStr + associatedResourceIdVariableStr + OpsiAwrHubSourceResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + associatedOpsiIdVariableStr + associatedResourceIdVariableStr + OpsiAwrHubSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source", "test_awr_hub_source", acctest.Optional, acctest.Create, OpsiAwrHubSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "associated_opsi_id"),
				resource.TestCheckResourceAttrSet(resourceName, "associated_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "awr_hub_id"),
				resource.TestCheckResourceAttrSet(resourceName, "awr_hub_opsi_source_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "source_mail_box_url"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ADW_S"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + associatedOpsiIdVariableStr + associatedResourceIdVariableStr + OpsiAwrHubSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source", "test_awr_hub_source", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OpsiAwrHubSourceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "associated_opsi_id"),
				resource.TestCheckResourceAttrSet(resourceName, "associated_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "awr_hub_id"),
				resource.TestCheckResourceAttrSet(resourceName, "awr_hub_opsi_source_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "source_mail_box_url"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ADW_S"),

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
			Config: config + compartmentIdVariableStr + associatedOpsiIdVariableStr + associatedResourceIdVariableStr + OpsiAwrHubSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source", "test_awr_hub_source", acctest.Optional, acctest.Update, OpsiAwrHubSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "associated_opsi_id"),
				resource.TestCheckResourceAttrSet(resourceName, "associated_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "awr_hub_id"),
				resource.TestCheckResourceAttrSet(resourceName, "awr_hub_opsi_source_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "source_mail_box_url"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "status"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ATP_S"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub_sources", "test_awr_hub_sources", acctest.Optional, acctest.Update, OpsiAwrHubSourceDataSourceRepresentation) +
				compartmentIdVariableStr + associatedOpsiIdVariableStr + associatedResourceIdVariableStr + OpsiAwrHubSourceResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub_source", "test_awr_hub_source", acctest.Optional, acctest.Update, OpsiAwrHubSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "awr_hub_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "awr_hub_source_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "source_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "status.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "awr_hub_source_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "awr_hub_source_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub_source", "test_awr_hub_source", acctest.Required, acctest.Create, OpsiAwrHubSourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + associatedOpsiIdVariableStr + associatedResourceIdVariableStr + OpsiAwrHubSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "awr_hub_source_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "awr_hub_opsi_source_id"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "awr_source_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "hours_since_last_import"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "is_registered_with_awr_hub"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "max_snapshot_identifier"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "min_snapshot_identifier"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_mail_box_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_first_snapshot_generated"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_snapshot_generated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ATP_S"),
			),
		},
		// verify resource import
		{
			Config:                  config + OpsiAwrHubSourceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOpsiAwrHubSourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_awr_hub_source" {
			noResourceFound = false
			request := oci_opsi.GetAwrHubSourceRequest{}

			tmp := rs.Primary.ID
			request.AwrHubSourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			response, err := client.GetAwrHubSource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.AwrHubSourceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OpsiAwrHubSource") {
		resource.AddTestSweepers("OpsiAwrHubSource", &resource.Sweeper{
			Name:         "OpsiAwrHubSource",
			Dependencies: acctest.DependencyGraph["awrHubSource"],
			F:            sweepOpsiAwrHubSourceResource,
		})
	}
}

func sweepOpsiAwrHubSourceResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	awrHubSourceIds, err := getOpsiAwrHubSourceIds(compartment)
	if err != nil {
		return err
	}
	for _, awrHubSourceId := range awrHubSourceIds {
		if ok := acctest.SweeperDefaultResourceId[awrHubSourceId]; !ok {
			deleteAwrHubSourceRequest := oci_opsi.DeleteAwrHubSourceRequest{}

			deleteAwrHubSourceRequest.AwrHubSourceId = &awrHubSourceId

			deleteAwrHubSourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteAwrHubSource(context.Background(), deleteAwrHubSourceRequest)
			if error != nil {
				fmt.Printf("Error deleting AwrHubSource %s %s, It is possible that the resource is already deleted. Please verify manually \n", awrHubSourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &awrHubSourceId, OpsiAwrHubSourceSweepWaitCondition, time.Duration(3*time.Minute),
				OpsiAwrHubSourceSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getOpsiAwrHubSourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AwrHubSourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

	listAwrHubSourcesRequest := oci_opsi.ListAwrHubSourcesRequest{}
	listAwrHubSourcesRequest.CompartmentId = &compartmentId

	awrHubIds, error := getOpsiAwrHubIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting awrHubId required for AwrHubSource resource requests \n")
	}
	for _, awrHubId := range awrHubIds {
		listAwrHubSourcesRequest.AwrHubId = &awrHubId
		tmp := make([]oci_opsi.AwrHubSourceLifecycleStateEnum, 1)
		tmp[0] = oci_opsi.AwrHubSourceLifecycleStateActive
		listAwrHubSourcesRequest.LifecycleState = tmp
		listAwrHubSourcesResponse, err := operationsInsightsClient.ListAwrHubSources(context.Background(), listAwrHubSourcesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting AwrHubSource list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, awrHubSource := range listAwrHubSourcesResponse.Items {
			id := *awrHubSource.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AwrHubSourceId", id)
		}

	}
	return resourceIds, nil
}

func OpsiAwrHubSourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if awrHubSourceResponse, ok := response.Response.(oci_opsi.GetAwrHubSourceResponse); ok {
		return awrHubSourceResponse.LifecycleState != oci_opsi.AwrHubSourceLifecycleStateDeleted
	}
	return false
}

func OpsiAwrHubSourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetAwrHubSource(context.Background(), oci_opsi.GetAwrHubSourceRequest{
		AwrHubSourceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
