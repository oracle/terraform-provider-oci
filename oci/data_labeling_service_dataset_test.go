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
	"github.com/oracle/oci-go-sdk/v51/common"
	oci_data_labeling_service "github.com/oracle/oci-go-sdk/v51/datalabelingservice"
	"github.com/oracle/oci-go-sdk/v51/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DatasetRequiredOnlyResource = DatasetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", Required, Create, datasetRepresentation)

	DatasetResourceConfig = DatasetResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", Optional, Update, datasetRepresentation)

	datasetSingularDataSourceRepresentation = map[string]interface{}{
		"dataset_id": Representation{RepType: Required, Create: `${oci_data_labeling_service_dataset.test_dataset.id}`},
	}

	datasetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    Representation{RepType: Required, Create: `${var.compartment_id}`},
		"annotation_format": Representation{RepType: Optional, Create: `BOUNDING_BOX`},
		"display_name":      Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":                Representation{RepType: Optional, Create: `${oci_data_labeling_service_dataset.test_dataset.id}`},
		"state":             Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":            RepresentationGroup{Required, datasetDataSourceFilterRepresentation}}
	datasetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_data_labeling_service_dataset.test_dataset.id}`}},
	}

	datasetRepresentation = map[string]interface{}{
		"annotation_format":      Representation{RepType: Required, Create: `BOUNDING_BOX`},
		"compartment_id":         Representation{RepType: Required, Create: `${var.compartment_id}`},
		"dataset_format_details": RepresentationGroup{Required, datasetDatasetFormatDetailsRepresentation},
		"dataset_source_details": RepresentationGroup{Required, datasetDatasetSourceDetailsRepresentation},
		"label_set":              RepresentationGroup{Required, datasetLabelSetRepresentation},
		"defined_tags":           Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "UpdatedValue")}`},
		"description":            Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":           Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":          Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		//"initial_record_generation_configuration": RepresentationGroup{Optional, datasetInitialRecordGenerationConfigurationRepresentation},
	}
	datasetDatasetFormatDetailsRepresentation = map[string]interface{}{
		"format_type": Representation{RepType: Required, Create: `IMAGE`},
	}
	datasetDatasetSourceDetailsRepresentation = map[string]interface{}{
		"bucket":      Representation{RepType: Required, Create: objectstorageBucket},
		"namespace":   Representation{RepType: Required, Create: objectstorageNamespace},
		"source_type": Representation{RepType: Required, Create: `OBJECT_STORAGE`},
		"prefix":      Representation{RepType: Optional, Create: `prefix`},
	}
	datasetLabelSetRepresentation = map[string]interface{}{
		"items": RepresentationGroup{Required, datasetLabelSetItemsRepresentation},
	}
	//datasetInitialRecordGenerationConfigurationRepresentation = map[string]interface{}{}
	datasetLabelSetItemsRepresentation = map[string]interface{}{
		"name": Representation{RepType: Required, Create: `name`, Update: `name2`},
	}

	//Representation map to Create ObjectStorage bucket
	bucketRepresentationDataset = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"name":           Representation{RepType: Required, Create: objectstorageBucket},
		"namespace":      Representation{RepType: Required, Create: objectstorageNamespace},
	}

	objectstorageNamespace = "${data.oci_objectstorage_namespace.test_namespace.namespace}"
	objectstorageBucket    = "tf_dataset_objectstoragebucket"

	DatasetResourceDependencies = GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentationDataset) +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: data_labeling_service/default
func TestDataLabelingServiceDatasetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataLabelingServiceDatasetResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_Update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_Update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_data_labeling_service_dataset.test_dataset"
	datasourceName := "data.oci_data_labeling_service_datasets.test_datasets"
	singularDatasourceName := "data.oci_data_labeling_service_dataset.test_dataset"

	objectstorageNamespace := getobjectstoragenamespace(compartmentId)
	objectstorageBucket := "tf_dataset_objectstoragebucket"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DatasetResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", Optional, Create, datasetRepresentation), "datalabelingservice", "dataset", t)

	ResourceTest(t, testAccCheckDataLabelingServiceDatasetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatasetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", Required, Create, datasetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", Optional, Create, datasetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "annotation_format", "BOUNDING_BOX"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.0.format_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.bucket", objectstorageBucket),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.namespace", objectstorageNamespace),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.source_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "initial_record_generation_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.0.name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", Optional, Create,
					RepresentationCopyWithNewProperties(datasetRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_Update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "annotation_format", "BOUNDING_BOX"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.0.format_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.bucket", objectstorageBucket),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.namespace", objectstorageNamespace),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.source_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "initial_record_generation_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.0.name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", Optional, Update, datasetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "annotation_format", "BOUNDING_BOX"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_format_details.0.format_type", "IMAGE"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.bucket", objectstorageBucket),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.namespace", objectstorageNamespace),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(resourceName, "dataset_source_details.0.source_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				//resource.TestCheckResourceAttr(resourceName, "initial_record_generation_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "label_set.0.items.0.name", "name2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_data_labeling_service_datasets", "test_datasets", Optional, Update, datasetDataSourceRepresentation) +
				compartmentIdVariableStr + DatasetResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", Optional, Update, datasetRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", Required, Create, datasetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatasetResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "initial_record_generation_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "label_set.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "label_set.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "label_set.0.items.0.name", "name2"),
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
	client := testAccProvider.Meta().(*OracleClients).dataLabelingManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_labeling_service_dataset" {
			noResourceFound = false
			request := oci_data_labeling_service.GetDatasetRequest{}

			tmp := rs.Primary.ID
			request.DatasetId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "data_labeling_service")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("DataLabelingServiceDataset") {
		resource.AddTestSweepers("DataLabelingServiceDataset", &resource.Sweeper{
			Name:         "DataLabelingServiceDataset",
			Dependencies: DependencyGraph["dataset"],
			F:            sweepDataLabelingServiceDatasetResource,
		})
	}
}

func sweepDataLabelingServiceDatasetResource(compartment string) error {
	dataLabelingManagementClient := GetTestClients(&schema.ResourceData{}).dataLabelingManagementClient()
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
		if ok := SweeperDefaultResourceId[datasetId]; !ok {
			deleteDatasetRequest := oci_data_labeling_service.DeleteDatasetRequest{}

			deleteDatasetRequest.DatasetId = &datasetId

			deleteDatasetRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "data_labeling_service")
			_, error := dataLabelingManagementClient.DeleteDataset(context.Background(), deleteDatasetRequest)
			if error != nil {
				fmt.Printf("Error deleting Dataset %s %s, It is possible that the resource is already deleted. Please verify manually \n", datasetId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &datasetId, datasetSweepWaitCondition, time.Duration(3*time.Minute),
				datasetSweepResponseFetchOperation, "data_labeling_service", true)
		}
	}
	return nil
}

func getDatasetIds(compartment string, lifecycleState oci_data_labeling_service.DatasetLifecycleStateEnum) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DatasetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataLabelingManagementClient := GetTestClients(&schema.ResourceData{}).dataLabelingManagementClient()

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
		AddResourceIdToSweeperResourceIdMap(compartmentId, "DatasetId", id)
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

func datasetSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataLabelingManagementClient().GetDataset(context.Background(), oci_data_labeling_service.GetDatasetRequest{
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
	objectStorageClient := GetTestClients(&schema.ResourceData{}).objectStorageClient()
	ctx := context.Background()
	request := objectstorage.GetNamespaceRequest{CompartmentId: common.String(compartment)}
	r, err := objectStorageClient.GetNamespace(ctx, request)
	if err != nil {
		err := fmt.Errorf("Error getting namespace : %v", err)
		fmt.Println(err.Error())
	}
	return *r.Value
}
