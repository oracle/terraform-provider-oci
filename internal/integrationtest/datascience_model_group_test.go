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
	DatascienceModelGroupRequiredOnlyResource = DatascienceModelGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group", "test_model_group", acctest.Required, acctest.Create, DatascienceModelGroupRepresentation)

	DatascienceModelGroupResourceConfig = DatascienceModelGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group", "test_model_group_optional", acctest.Optional, acctest.Update, DatascienceModelGroupRepresentation)

	DatascienceModelGroupSingularDataSourceRepresentation = map[string]interface{}{
		"model_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model_group.test_model_group.id}`},
	}

	DatascienceModelGroupDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":                     acctest.Representation{RepType: acctest.Optional, Create: `createdBy`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_group.test_model_group.id}`},
		"model_group_version_history_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_group_version_history.test_model_group_version_history.id}`},
		"project_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.test_project.id}`},
		"member_model_entries":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelGroupMemberModelEntriesRepresentation},
		"model_group_details":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelGroupModelGroupDetailsRepresentation},
		"state":                          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelGroupDataSourceFilterRepresentation}}
	DatascienceModelGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_model_group.test_model_group.id}`}},
	}

	DatascienceModelGroupRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"create_type":                    acctest.Representation{RepType: acctest.Required, Create: `CREATE`},
		"project_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"member_model_entries":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelGroupMemberModelEntriesRepresentation},
		"model_group_details":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelGroupModelGroupDetailsRepresentation},
		"model_group_version_history_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_group_version_history.test_model_group_version_history.id}`},
		"version_label":                  acctest.Representation{RepType: acctest.Optional, Create: `versionLabel`, Update: `versionLabel2`},
	}
	DatascienceModelGroupMemberModelEntriesRepresentation = map[string]interface{}{
		"member_model_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelGroupMemberModelEntriesMemberModelDetailsRepresentation},
	}
	DatascienceModelGroupModelGroupCloneSourceDetailsRepresentation = map[string]interface{}{
		"model_group_clone_source_type":          acctest.Representation{RepType: acctest.Required, Create: `MODEL_GROUP`},
		"source_id":                              acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_source.test_source.id}`},
		"modify_model_group_details":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelGroupModelGroupCloneSourceDetailsModifyModelGroupDetailsRepresentation},
		"patch_model_group_member_model_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelGroupModelGroupCloneSourceDetailsPatchModelGroupMemberModelDetailsRepresentation},
	}
	DatascienceModelGroupModelGroupDetailsRepresentation = map[string]interface{}{
		"type":                 acctest.Representation{RepType: acctest.Required, Create: `HOMOGENEOUS`},
		"base_model_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model.test_model.id}`},
		"custom_metadata_list": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelGroupModelGroupDetailsCustomMetadataListRepresentation},
	}
	DatascienceModelGroupMemberModelEntriesMemberModelDetailsRepresentation = map[string]interface{}{
		"inference_key": acctest.Representation{RepType: acctest.Optional, Create: `inferenceKey`},
		"model_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
	}
	DatascienceModelGroupModelGroupCloneSourceDetailsModifyModelGroupDetailsRepresentation = map[string]interface{}{
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
		"model_group_details":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelGroupModelGroupCloneSourceDetailsModifyModelGroupDetailsModelGroupDetailsRepresentation},
		"member_model_entries":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelGroupMemberModelEntriesRepresentation},
		"model_group_version_history_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model_group_version_history.test_model_group_version_history.id}`},
		"version_label":                  acctest.Representation{RepType: acctest.Optional, Create: `versionLabel`},
	}
	DatascienceModelGroupModelGroupCloneSourceDetailsPatchModelGroupMemberModelDetailsRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelGroupModelGroupCloneSourceDetailsPatchModelGroupMemberModelDetailsItemsRepresentation},
	}
	DatascienceModelGroupModelGroupDetailsCustomMetadataListRepresentation = map[string]interface{}{
		"category":    acctest.Representation{RepType: acctest.Optional, Create: `category`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`},
		"value":       acctest.Representation{RepType: acctest.Optional, Create: `value`},
	}
	DatascienceModelGroupModelGroupCloneSourceDetailsModifyModelGroupDetailsModelGroupDetailsRepresentation = map[string]interface{}{
		"type":                 acctest.Representation{RepType: acctest.Required, Create: `HOMOGENEOUS`},
		"base_model_id":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model.test_model.id}`},
		"custom_metadata_list": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelGroupModelGroupCloneSourceDetailsModifyModelGroupDetailsModelGroupDetailsCustomMetadataListRepresentation},
	}
	DatascienceModelGroupModelGroupCloneSourceDetailsPatchModelGroupMemberModelDetailsItemsRepresentation = map[string]interface{}{
		"operation": acctest.Representation{RepType: acctest.Required, Create: `INSERT`},
		"values":    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelGroupModelGroupCloneSourceDetailsPatchModelGroupMemberModelDetailsItemsValuesRepresentation},
	}
	DatascienceModelGroupModelGroupCloneSourceDetailsModifyModelGroupDetailsModelGroupDetailsCustomMetadataListRepresentation = map[string]interface{}{
		"category":    acctest.Representation{RepType: acctest.Optional, Create: `category`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `key`},
		"value":       acctest.Representation{RepType: acctest.Optional, Create: `value`},
	}
	DatascienceModelGroupModelGroupCloneSourceDetailsPatchModelGroupMemberModelDetailsItemsValuesRepresentation = map[string]interface{}{
		"model_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
		"inference_key": acctest.Representation{RepType: acctest.Optional, Create: `inferenceKey`},
	}

	DatascienceModelGroupResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group_version_history", "test_model_group_version_history", acctest.Required, acctest.Create, DatascienceModelGroupVersionHistoryRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, DatascienceModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceModelGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_model_group.test_model_group"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceModelGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group", "test_model_group", acctest.Optional, acctest.Create, DatascienceModelGroupRepresentation), "datascience", "model-group", t)

	acctest.ResourceTest(t, testAccCheckDatascienceModelGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_group", "test_model_group", acctest.Required, acctest.Create, DatascienceModelGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "member_model_entries.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_group_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_group_details.0.type", "HOMOGENEOUS"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
	})
}

func testAccCheckDatascienceModelGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_model_group" {
			noResourceFound = false
			request := oci_datascience.GetModelGroupRequest{}

			tmp := rs.Primary.ID
			request.ModelGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetModelGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ModelGroupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceModelGroup") {
		resource.AddTestSweepers("DatascienceModelGroup", &resource.Sweeper{
			Name:         "DatascienceModelGroup",
			Dependencies: acctest.DependencyGraph["model-group"],
			F:            sweepDatascienceModelGroupResource,
		})
	}
}

func sweepDatascienceModelGroupResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	modelGroupIds, err := getDatascienceModelGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, modelGroupId := range modelGroupIds {
		if ok := acctest.SweeperDefaultResourceId[modelGroupId]; !ok {
			deleteModelGroupRequest := oci_datascience.DeleteModelGroupRequest{}

			deleteModelGroupRequest.ModelGroupId = &modelGroupId

			deleteModelGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteModelGroup(context.Background(), deleteModelGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting ModelGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelGroupId, DatascienceModelGroupSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceModelGroupSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceModelGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listModelGroupsRequest := oci_datascience.ListModelGroupsRequest{}
	listModelGroupsRequest.CompartmentId = &compartmentId
	listModelGroupsRequest.LifecycleState = oci_datascience.ListModelGroupsLifecycleStateActive
	listModelGroupsResponse, err := dataScienceClient.ListModelGroups(context.Background(), listModelGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ModelGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, modelGroup := range listModelGroupsResponse.Items {
		id := *modelGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ModelGroupId", id)
	}
	return resourceIds, nil
}

func DatascienceModelGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelGroupResponse, ok := response.Response.(oci_datascience.GetModelGroupResponse); ok {
		return modelGroupResponse.LifecycleState != oci_datascience.ModelGroupLifecycleStateDeleted
	}
	return false
}

func DatascienceModelGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetModelGroup(context.Background(), oci_datascience.GetModelGroupRequest{
		ModelGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
