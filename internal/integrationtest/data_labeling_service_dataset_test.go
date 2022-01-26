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
	oci_data_labeling_service "github.com/oracle/oci-go-sdk/v56/datalabelingservice"
	"github.com/oracle/oci-go-sdk/v56/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DatasetRequiredOnlyResource = DatasetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Required, acctest.Create, datasetRepresentation)

	DatasetResourceConfig = DatasetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Optional, acctest.Update, datasetRepresentation)

	datasetSingularDataSourceRepresentation = map[string]interface{}{
		"dataset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_labeling_service_dataset.test_dataset.id}`},
	}

	datasetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"annotation_format": acctest.Representation{RepType: acctest.Optional, Create: `BOUNDING_BOX`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_labeling_service_dataset.test_dataset.id}`},
		"state":             acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: datasetDataSourceFilterRepresentation}}
	datasetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_labeling_service_dataset.test_dataset.id}`}},
	}

	datasetIgnoreDefinedTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	datasetRepresentation = map[string]interface{}{
		"annotation_format":      acctest.Representation{RepType: acctest.Required, Create: `BOUNDING_BOX`},
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dataset_format_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: datasetDatasetFormatDetailsRepresentation},
		"dataset_source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: datasetDatasetSourceDetailsRepresentation},
		"label_set":              acctest.RepresentationGroup{RepType: acctest.Required, Group: datasetLabelSetRepresentation},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "UpdatedValue")}`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		//"initial_record_generation_configuration": acctest.RepresentationGroup{acctest.Optional, datasetInitialRecordGenerationConfigurationRepresentation},
		//"labeling_instructions": acctest.Representation{RepType: acctest.Optional, Create: `labelingInstructions`, Update: `labelingInstructions2`},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: datasetIgnoreDefinedTagsChangesRep},
	}
	datasetDatasetFormatDetailsRepresentation = map[string]interface{}{
		"format_type": acctest.Representation{RepType: acctest.Required, Create: `IMAGE`},
	}
	datasetDatasetSourceDetailsRepresentation = map[string]interface{}{
		"bucket":      acctest.Representation{RepType: acctest.Required, Create: objectstorageBucket},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: objectstorageNamespace},
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		"prefix":      acctest.Representation{RepType: acctest.Optional, Create: `prefix`},
	}
	datasetLabelSetRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: datasetLabelSetItemsRepresentation},
	}
	//datasetInitialRecordGenerationConfigurationRepresentation = map[string]interface{}{}
	datasetLabelSetItemsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
	}

	//Representation map to Create ObjectStorage bucket
	bucketRepresentationDataset = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: objectstorageBucket},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: objectstorageNamespace},
	}

	objectstorageNamespace = "${data.oci_objectstorage_namespace.test_namespace.namespace}"
	objectstorageBucket    = "tf_dataset_objectstoragebucket"

	DatasetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentationDataset) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: data_labeling_service/default
func TestDataLabelingServiceDatasetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataLabelingServiceDatasetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id"+
		"\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_Update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_Update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_labeling_service_dataset.test_dataset"
	datasourceName := "data.oci_data_labeling_service_datasets.test_datasets"
	singularDatasourceName := "data.oci_data_labeling_service_dataset.test_dataset"

	objectstorageNamespace := getobjectstoragenamespace(compartmentId)
	objectstorageBucket := "tf_dataset_objectstoragebucket"

	var resId, resId2 string
	// Save TF content to Create resource with optional propeTestCheckResourceAttrrties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatasetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Optional, acctest.Create, datasetRepresentation), "datalabelingservice", "dataset", t)

	acctest.ResourceTest(t, testAccCheckDataLabelingServiceDatasetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatasetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Required, acctest.Create, datasetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "annotation_format", "BOUNDING_BOX"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.0.format_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.bucket", objectstorageBucket),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.namespace", objectstorageNamespace),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.source_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "label_set.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.0.name", "name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatasetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatasetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Optional, acctest.Create, datasetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "annotation_format", "BOUNDING_BOX"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.0.format_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.bucket", objectstorageBucket),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.namespace", objectstorageNamespace),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.source_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "initial_record_generation_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.0.name", "name"),
				//resource.TestCheckResourceAttr(resourceName, "labeling_instructions", "labelingInstructions"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatasetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(datasetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_Update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "annotation_format", "BOUNDING_BOX"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.0.format_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.bucket", objectstorageBucket),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.namespace", objectstorageNamespace),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.source_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "initial_record_generation_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.0.name", "name"),
				//resource.TestCheckResourceAttr(resourceName, "labeling_instructions", "labelingInstructions"),
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

		// verify Updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatasetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Optional, acctest.Update, datasetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "annotation_format", "BOUNDING_BOX"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.0.format_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.bucket", objectstorageBucket),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.namespace", objectstorageNamespace),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.source_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "initial_record_generation_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.0.name", "name2"),
				//resource.TestCheckResourceAttr(resourceName, "labeling_instructions", "labelingInstructions2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_labeling_service_datasets", "test_datasets", acctest.Optional, acctest.Update, datasetDataSourceRepresentation) +
				compartmentIdVariableStr + DatasetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Optional, acctest.Update, datasetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "annotation_format", "BOUNDING_BOX"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				//resource.TestCheckResourceAttr(datasourceName, "id", "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "dataset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dataset_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Required, acctest.Create, datasetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatasetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dataset_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "annotation_format", "BOUNDING_BOX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dataset_format_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dataset_format_details.0.format_type", "IMAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dataset_source_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dataset_source_details.0.bucket", objectstorageBucket),
				resource.TestCheckResourceAttr(singularDatasourceName, "dataset_source_details.0.namespace", objectstorageNamespace),
				resource.TestCheckResourceAttr(singularDatasourceName, "dataset_source_details.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dataset_source_details.0.source_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "initial_record_generation_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "label_set.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "label_set.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "label_set.0.items.0.name", "name2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "labeling_instructions", "labelingInstructions2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DatasetResourceConfig,
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

func testAccCheckDataLabelingServiceDatasetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataLabelingManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_labeling_service_dataset" {
			noResourceFound = false
			request := oci_data_labeling_service.GetDatasetRequest{}

			tmp := rs.Primary.ID
			request.DatasetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_labeling_service")

			response, err := client.GetDataset(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_labeling_service.DatasetLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataLabelingServiceDataset") {
		resource.AddTestSweepers("DataLabelingServiceDataset", &resource.Sweeper{
			Name:         "DataLabelingServiceDataset",
			Dependencies: acctest.DependencyGraph["dataset"],
			F:            sweepDataLabelingServiceDatasetResource,
		})
	}
}

func sweepDataLabelingServiceDatasetResource(compartment string) error {
	dataLabelingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DataLabelingManagementClient()
	activeDatasetIds, err := getDatasetIds(compartment, oci_data_labeling_service.DatasetLifecycleStateActive)
	if err != nil {
		return err
	}
	nonActiveDatasetIds, err := getDatasetIds(compartment, oci_data_labeling_service.DatasetLifecycleStateNeedsAttention)
	if err != nil {
		return err
	}
	datasetIds := append(activeDatasetIds, nonActiveDatasetIds...)
	for _, datasetId := range datasetIds {
		if ok := acctest.SweeperDefaultResourceId[datasetId]; !ok {
			deleteDatasetRequest := oci_data_labeling_service.DeleteDatasetRequest{}

			deleteDatasetRequest.DatasetId = &datasetId

			deleteDatasetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_labeling_service")
			_, error := dataLabelingManagementClient.DeleteDataset(context.Background(), deleteDatasetRequest)
			if error != nil {
				fmt.Printf("Error deleting Dataset %s %s, It is possible that the resource is already deleted. Please verify manually \n", datasetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &datasetId, datasetSweepWaitCondition, time.Duration(3*time.Minute),
				datasetSweepResponseFetchOperation, "data_labeling_service", true)
		}
	}
	return nil
}

func getDatasetIds(compartment string, lifecycleState oci_data_labeling_service.DatasetLifecycleStateEnum) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatasetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataLabelingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DataLabelingManagementClient()

	listDatasetsRequest := oci_data_labeling_service.ListDatasetsRequest{}
	listDatasetsRequest.CompartmentId = &compartmentId
	listDatasetsRequest.LifecycleState = lifecycleState
	listDatasetsResponse, err := dataLabelingManagementClient.ListDatasets(context.Background(), listDatasetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Dataset list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dataset := range listDatasetsResponse.Items {
		id := *dataset.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatasetId", id)
	}
	return resourceIds, nil
}

func datasetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if datasetResponse, ok := response.Response.(oci_data_labeling_service.GetDatasetResponse); ok {
		return datasetResponse.LifecycleState != oci_data_labeling_service.DatasetLifecycleStateDeleted
	}
	return false
}

func datasetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataLabelingManagementClient().GetDataset(context.Background(), oci_data_labeling_service.GetDatasetRequest{
		DatasetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

//Function to get ObjectStorage Namespace
func getobjectstoragenamespace(compartmentId string) string {
	compartment := compartmentId
	objectStorageClient := acctest.GetTestClients(&schema.ResourceData{}).ObjectStorageClient()
	ctx := context.Background()
	request := objectstorage.GetNamespaceRequest{CompartmentId: common.String(compartment)}
	r, err := objectStorageClient.GetNamespace(ctx, request)
	if err != nil {
		err := fmt.Errorf("Error getting namespace : %v", err)
		fmt.Println(err.Error())
	}
	return *r.Value
}
