// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatascienceModelRequiredOnlyResource = DatascienceModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, DatascienceModelRepresentation)

	DatascienceModelResourceConfig = DatascienceModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Update, DatascienceModelRepresentation)

	DatascienceDatascienceModelSingularDataSourceRepresentation = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
	}

	DatascienceDatascienceModelDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model.test_model.id}`},
		"model_version_set_name": acctest.Representation{RepType: acctest.Optional, Create: ``},
		"project_id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"version_label":          acctest.Representation{RepType: acctest.Optional, Create: ``, Update: ``},
		"filter":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceModelDataSourceFilterRepresentation}}
	DatascienceModelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_model.test_model.id}`}},
	}

	DatascienceModelRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"backup_setting":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelBackupSettingRepresentation},
		"custom_metadata_list":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelCustomMetadataListRepresentation},
		"defined_metadata_list":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelDefinedMetadataListRepresentation},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"input_schema":                 acctest.Representation{RepType: acctest.Optional, Create: `inputSchema`},
		"output_schema":                acctest.Representation{RepType: acctest.Optional, Create: `outputSchema`},
		"retention_setting":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceModelRetentionSettingRepresentation},
		"artifact_content_length":      acctest.Representation{RepType: acctest.Required, Create: `6954`},
		"model_artifact":               acctest.Representation{RepType: acctest.Required, Create: `../../examples/datascience/artifact.zip`},
		"artifact_content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `attachment; filename=tfTestArtifact`},
	}
	DatascienceModelBackupSettingRepresentation = map[string]interface{}{
		"backup_region":              acctest.Representation{RepType: acctest.Required, Create: `us-phoenix-1`, Update: `us-phoenix-1`},
		"is_backup_enabled":          acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `true`},
		"customer_notification_type": acctest.Representation{RepType: acctest.Optional, Create: `NONE`, Update: `ALL`},
	}
	DatascienceModelCustomMetadataListRepresentation = map[string]interface{}{
		"category":    acctest.Representation{RepType: acctest.Optional, Create: `Performance`, Update: `Performance`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `BaseModel1`, Update: `BaseModel1`},
		"value":       acctest.Representation{RepType: acctest.Optional, Create: `xgb`, Update: `xgb`},
	}
	DatascienceModelDefinedMetadataListRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Optional, Create: `UseCaseType`, Update: `UseCaseType`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `ner`, Update: `ner`},
	}
	DatascienceModelRetentionSettingRepresentation = map[string]interface{}{
		"archive_after_days":         acctest.Representation{RepType: acctest.Required, Create: `40`, Update: `40`},
		"customer_notification_type": acctest.Representation{RepType: acctest.Optional, Create: `NONE`, Update: `ALL`},
		"delete_after_days":          acctest.Representation{RepType: acctest.Optional, Create: `45`, Update: `45`},
	}

	DatascienceModelResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies

	DatascienceModelResourceModelVersionSetDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_model_version_set", "test_model_version_set", acctest.Required, acctest.Create, DatascienceModelVersionSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_model.test_model"
	datasourceName := "data.oci_datascience_models.test_models"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceModelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create, DatascienceModelRepresentation), "datascience", "model", t)

	acctest.ResourceTest(t, testAccCheckDatascienceModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, DatascienceModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "artifact_content_length", "6954"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_content_md5"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_last_modified"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceModelResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create, DatascienceModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_operation_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.0.backup_region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.0.customer_notification_type", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.0.is_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "artifact_content_length", "6954"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_content_md5"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_last_modified"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.category", "Performance"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.key", "BaseModel1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.value", "xgb"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.key", "UseCaseType"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.value", "ner"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_schema", "{}"),
				resource.TestCheckResourceAttr(resourceName, "output_schema", "{}"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "retention_operation_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.0.archive_after_days", "40"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.0.customer_notification_type", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.0.delete_after_days", "45"),
				//resource.TestCheckResourceAttr(resourceName, "state", ACTIVE),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceModelRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_operation_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.0.backup_region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.0.customer_notification_type", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.0.is_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "artifact_content_length", "6954"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_content_md5"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_last_modified"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.category", "Performance"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.key", "BaseModel1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.value", "xgb"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.key", "UseCaseType"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.value", "ner"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_schema", "{}"),
				resource.TestCheckResourceAttr(resourceName, "output_schema", "{}"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "retention_operation_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.0.archive_after_days", "40"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.0.customer_notification_type", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.0.delete_after_days", "45"),
				//resource.TestCheckResourceAttr(resourceName, "state", ACTIVE),
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
			Config: config + compartmentIdVariableStr + DatascienceModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Update, DatascienceModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backup_operation_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.0.backup_region", "us-phoenix-1"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.0.customer_notification_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "backup_setting.0.is_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.category", "Performance"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.key", "BaseModel1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.value", "xgb"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.key", "UseCaseType"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.value", "ner"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_schema", "{}"),
				resource.TestCheckResourceAttr(resourceName, "output_schema", "{}"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "retention_operation_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.0.archive_after_days", "40"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.0.customer_notification_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "retention_setting.0.delete_after_days", "45"),
				//resource.TestCheckResourceAttr(resourceName, "state", ACTIVE),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_models", "test_models", acctest.Optional, acctest.Update, DatascienceDatascienceModelDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceModelResourceModelVersionSetDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Update, DatascienceModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "models.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "models.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "models.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "models.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "models.0.version_label", ""),
			),
		},

		// verify resource import
		{
			Config:            config + DatascienceModelRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"artifact_content_disposition",
				"artifact_content_md5",
				"artifact_last_modified",
				"artifact_content_length",
				"empty_model",
				"model_artifact",
				"model_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatascienceModelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_model" {
			noResourceFound = false
			request := oci_datascience.GetModelRequest{}

			tmp := rs.Primary.ID
			request.ModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ModelLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceModel") {
		resource.AddTestSweepers("DatascienceModel", &resource.Sweeper{
			Name:         "DatascienceModel",
			Dependencies: acctest.DependencyGraph["model"],
			F:            sweepDatascienceModelResource,
		})
	}
}

func sweepDatascienceModelResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	modelIds, err := getDatascienceModelIds(compartment)
	if err != nil {
		return err
	}
	for _, modelId := range modelIds {
		if ok := acctest.SweeperDefaultResourceId[modelId]; !ok {
			deleteModelRequest := oci_datascience.DeleteModelRequest{}

			deleteModelRequest.ModelId = &modelId

			deleteModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteModel(context.Background(), deleteModelRequest)
			if error != nil {
				fmt.Printf("Error deleting Model %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelId, DatascienceModelSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceModelSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceModelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listModelsRequest := oci_datascience.ListModelsRequest{}
	listModelsRequest.CompartmentId = &compartmentId
	listModelsRequest.LifecycleState = oci_datascience.ListModelsLifecycleStateActive
	listModelsResponse, err := dataScienceClient.ListModels(context.Background(), listModelsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Model list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, model := range listModelsResponse.Items {
		id := *model.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ModelId", id)
	}
	return resourceIds, nil
}

func DatascienceModelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is ACTIVE beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelResponse, ok := response.Response.(oci_datascience.GetModelResponse); ok {
		return modelResponse.LifecycleState != oci_datascience.ModelLifecycleStateDeleted
	}
	return false
}

func DatascienceModelSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetModel(context.Background(), oci_datascience.GetModelRequest{
		ModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
