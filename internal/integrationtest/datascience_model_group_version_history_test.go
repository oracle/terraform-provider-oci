// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatascienceModelGroupVersionHistoryRequiredOnlyResource = DatascienceModelGroupVersionHistoryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group_version_history", "test_model_group_version_history", acctest.Required, acctest.Create, DatascienceModelGroupVersionHistoryRepresentation)

	DatascienceModelGroupVersionHistoryResourceConfig = DatascienceModelGroupVersionHistoryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group_version_history", "test_model_group_version_history", acctest.Optional, acctest.Update, DatascienceModelGroupVersionHistoryRepresentation)

	DatascienceModelGroupVersionHistorySingularDataSourceRepresentation = map[string]interface{}{
		"model_group_version_history_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model_group_version_history.test_model_group_version_history.id}`},
	}

	DatascienceModelGroupVersionHistoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":     acctest.Representation{RepType: acctest.Optional, Create: `createdBy`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_group_version_history.test_model_group_version_history.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelGroupVersionHistoryDataSourceFilterRepresentation}}
	DatascienceModelGroupVersionHistoryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_model_group_version_history.test_model_group_version_history.id}`}},
	}

	DatascienceModelGroupVersionHistoryRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":            acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"latest_model_group_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_group.test_model_group.id}`},
	}

	DatascienceModelGroupVersionHistoryResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group", "test_model_group", acctest.Required, acctest.Create, DatascienceModelGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, DatascienceModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceModelGroupVersionHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelGroupVersionHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_model_group_version_history.test_model_group_version_history"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceModelGroupVersionHistoryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group_version_history", "test_model_group_version_history", acctest.Optional, acctest.Create, DatascienceModelGroupVersionHistoryRepresentation), "datascience", "modelGroupVersionHistory", t)

	acctest.ResourceTest(t, testAccCheckDatascienceModelGroupVersionHistoryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelGroupVersionHistoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group_version_history", "test_model_group_version_history", acctest.Required, acctest.Create, DatascienceModelGroupVersionHistoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
	})
}

func testAccCheckDatascienceModelGroupVersionHistoryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_model_group_version_history" {
			noResourceFound = false
			request := oci_datascience.GetModelGroupVersionHistoryRequest{}

			tmp := rs.Primary.ID
			request.ModelGroupVersionHistoryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetModelGroupVersionHistory(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ModelGroupVersionHistoryLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceModelGroupVersionHistory") {
		resource.AddTestSweepers("DatascienceModelGroupVersionHistory", &resource.Sweeper{
			Name:         "DatascienceModelGroupVersionHistory",
			Dependencies: acctest.DependencyGraph["modelGroupVersionHistory"],
			F:            sweepDatascienceModelGroupVersionHistoryResource,
		})
	}
}

func sweepDatascienceModelGroupVersionHistoryResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	modelGroupVersionHistoryIds, err := getDatascienceModelGroupVersionHistoryIds(compartment)
	if err != nil {
		return err
	}
	for _, modelGroupVersionHistoryId := range modelGroupVersionHistoryIds {
		if ok := acctest.SweeperDefaultResourceId[modelGroupVersionHistoryId]; !ok {
			deleteModelGroupVersionHistoryRequest := oci_datascience.DeleteModelGroupVersionHistoryRequest{}

			deleteModelGroupVersionHistoryRequest.ModelGroupVersionHistoryId = &modelGroupVersionHistoryId

			deleteModelGroupVersionHistoryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteModelGroupVersionHistory(context.Background(), deleteModelGroupVersionHistoryRequest)
			if error != nil {
				fmt.Printf("Error deleting ModelGroupVersionHistory %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelGroupVersionHistoryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelGroupVersionHistoryId, DatascienceModelGroupVersionHistorySweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceModelGroupVersionHistorySweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceModelGroupVersionHistoryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelGroupVersionHistoryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listModelGroupVersionHistoriesRequest := oci_datascience.ListModelGroupVersionHistoriesRequest{}
	listModelGroupVersionHistoriesRequest.CompartmentId = &compartmentId
	listModelGroupVersionHistoriesRequest.LifecycleState = oci_datascience.ListModelGroupVersionHistoriesLifecycleStateActive
	listModelGroupVersionHistoriesResponse, err := dataScienceClient.ListModelGroupVersionHistories(context.Background(), listModelGroupVersionHistoriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ModelGroupVersionHistory list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, modelGroupVersionHistory := range listModelGroupVersionHistoriesResponse.Items {
		id := *modelGroupVersionHistory.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ModelGroupVersionHistoryId", id)
	}
	return resourceIds, nil
}

func DatascienceModelGroupVersionHistorySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelGroupVersionHistoryResponse, ok := response.Response.(oci_datascience.GetModelGroupVersionHistoryResponse); ok {
		return modelGroupVersionHistoryResponse.LifecycleState != oci_datascience.ModelGroupVersionHistoryLifecycleStateDeleted
	}
	return false
}

func DatascienceModelGroupVersionHistorySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetModelGroupVersionHistory(context.Background(), oci_datascience.GetModelGroupVersionHistoryRequest{
		ModelGroupVersionHistoryId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
