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
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v58/aianomalydetection"
	"github.com/oracle/oci-go-sdk/v58/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AiAnomalyDetectionDataAssetRequiredOnlyResource = AiAnomalyDetectionDataAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Required, acctest.Create, aiAnomalyDetectionDataAssetRepresentation)

	AiAnomalyDetectionDataAssetResourceConfig = AiAnomalyDetectionDataAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Optional, acctest.Update, aiAnomalyDetectionDataAssetRepresentation)

	aiAnomalyDetectionDataAssetSingularDataSourceRepresentation = map[string]interface{}{
		"data_asset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_anomaly_detection_data_asset.test_data_asset.id}`},
	}

	aiAnomalyDetectionDataAssetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: aiAnomalyDetectionDataAssetDataSourceFilterRepresentation}}
	aiAnomalyDetectionDataAssetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_anomaly_detection_data_asset.test_data_asset.id}`}},
	}

	aiAnomalyDetectionDataAssetRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: dataAssetDataSourceDetailsObjRepresentation},
		"project_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_anomaly_detection_project.test_project.id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	aiAnomalyDetectionDataAssetAtpRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: dataAssetDataSourceDetailsAtpRepresentation},
		"project_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_anomaly_detection_project.test_project.id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_anomaly_detection_ai_private_endpoint.test_private_endpoint.id}`},
	}

	aiAnomalyDetectionDataAssetInfluxRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: dataAssetDataSourceDetailsInfluxRepresentation},
		"project_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_anomaly_detection_project.test_project.id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_anomaly_detection_ai_private_endpoint.test_private_endpoint.id}`},
	}

	dataAssetDataSourceDetailsObjRepresentation = map[string]interface{}{
		"data_source_type": acctest.Representation{RepType: acctest.Required, Create: `ORACLE_OBJECT_STORAGE`},
		"bucket":           acctest.Representation{RepType: acctest.Required, Create: `bucket-test`},
		"namespace":        acctest.Representation{RepType: acctest.Required, Create: `dxterraformtest`},
		"object":           acctest.Representation{RepType: acctest.Required, Create: `latest_training_data.json`},
	}

	dataAssetDataSourceDetailsAtpRepresentation = map[string]interface{}{
		"data_source_type":          acctest.Representation{RepType: acctest.Required, Create: `ORACLE_ATP`},
		"atp_password_secret_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_deslt_secret.test_secret.id}`},
		"atp_user_name":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.name}`},
		"cwallet_file_secret_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"database_name":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database.test_database.name}`},
		"ewallet_file_secret_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"key_store_file_secret_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"ojdbc_file_secret_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"table_name":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_nosql_table.test_table.name}`},
		"tnsnames_file_secret_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"truststore_file_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"wallet_password_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
	}

	dataAssetDataSourceDetailsInfluxRepresentation = map[string]interface{}{
		"data_source_type":   acctest.Representation{RepType: acctest.Required, Create: `INFLUX`},
		"measurement_name":   acctest.Representation{RepType: acctest.Optional, Create: `measurementName`},
		"password_secret_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_vault_secret.test_secret.id}`},
		"url":                acctest.Representation{RepType: acctest.Optional, Create: `url`},
		"user_name":          acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_user.test_user.name}`},
	}

	dataAssetDataSourceDetailsVersionSpecificDetailsRepresentation = map[string]interface{}{
		"influx_version":        acctest.Representation{RepType: acctest.Required, Create: `V_1_8`},
		"bucket":                acctest.Representation{RepType: acctest.Optional, Create: `bucket`},
		"database_name":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database.test_database.name}`},
		"organization_name":     acctest.Representation{RepType: acctest.Optional, Create: `organizationName`},
		"retention_policy_name": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_policy.test_policy.name}`},
	}

	//Change this to only what is required
	AiAnomalyDetectionDataAssetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Required, acctest.Create, aiAnomalyDetectionProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Optional, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Optional, acctest.Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

func TestAiAnomalyDetectionDataAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiAnomalyDetectionDataAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_anomaly_detection_data_asset.test_data_asset"
	datasourceName := "data.oci_ai_anomaly_detection_data_assets.test_data_assets"
	singularDatasourceName := "data.oci_ai_anomaly_detection_data_asset.test_data_asset"

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiAnomalyDetectionDataAssetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Optional, acctest.Create, aiAnomalyDetectionDataAssetRepresentation), "aianomalydetection", "dataAsset", t)

	acctest.ResourceTest(t, testAccCheckAiAnomalyDetectionDataAssetDestroy, []resource.TestStep{
		// verify Create
		{
			//print this
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionDataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Required, acctest.Create, aiAnomalyDetectionDataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.data_source_type", "ORACLE_OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.bucket", "bucket-test"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.namespace", "dxterraformtest"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.object", "latest_training_data.json"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionDataAssetResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionDataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Optional, acctest.Create, aiAnomalyDetectionDataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.bucket", "bucket-test"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.data_source_type", "ORACLE_OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.namespace", "dxterraformtest"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.object", "latest_training_data.json"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiAnomalyDetectionDataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(aiAnomalyDetectionDataAssetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.bucket", "bucket-test"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.data_source_type", "ORACLE_OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.namespace", "dxterraformtest"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.object", "latest_training_data.json"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionDataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Optional, acctest.Update, aiAnomalyDetectionDataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.bucket", "bucket-test"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.data_source_type", "ORACLE_OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.namespace", "dxterraformtest"),
				resource.TestCheckResourceAttr(resourceName, "data_source_details.0.object", "latest_training_data.json"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_data_assets", "test_data_assets", acctest.Optional, acctest.Update, aiAnomalyDetectionDataAssetDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionDataAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Optional, acctest.Update, aiAnomalyDetectionDataAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "data_asset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "data_asset_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Required, acctest.Create, aiAnomalyDetectionDataAssetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionDataAssetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_asset_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionDataAssetResourceConfig,
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

func testAccCheckAiAnomalyDetectionDataAssetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AnomalyDetectionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_anomaly_detection_data_asset" {
			noResourceFound = false
			request := oci_ai_anomaly_detection.GetDataAssetRequest{}

			tmp := rs.Primary.ID
			request.DataAssetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_anomaly_detection")

			response, err := client.GetDataAsset(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_anomaly_detection.DataAssetLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiAnomalyDetectionDataAsset") {
		resource.AddTestSweepers("AiAnomalyDetectionDataAsset", &resource.Sweeper{
			Name:         "AiAnomalyDetectionDataAsset",
			Dependencies: acctest.DependencyGraph["dataAsset"],
			F:            sweepAiAnomalyDetectionDataAssetResource,
		})
	}
}

func sweepAiAnomalyDetectionDataAssetResource(compartment string) error {
	anomalyDetectionClient := acctest.GetTestClients(&schema.ResourceData{}).AnomalyDetectionClient()
	dataAssetIds, err := getDataAssetIds(compartment)
	if err != nil {
		return err
	}
	for _, dataAssetId := range dataAssetIds {
		if ok := acctest.SweeperDefaultResourceId[dataAssetId]; !ok {
			deleteDataAssetRequest := oci_ai_anomaly_detection.DeleteDataAssetRequest{}

			deleteDataAssetRequest.DataAssetId = &dataAssetId

			deleteDataAssetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_anomaly_detection")
			_, error := anomalyDetectionClient.DeleteDataAsset(context.Background(), deleteDataAssetRequest)
			if error != nil {
				fmt.Printf("Error deleting DataAsset %s %s, It is possible that the resource is already deleted. Please verify manually \n", dataAssetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dataAssetId, dataAssetSweepWaitCondition, time.Duration(3*time.Minute),
				dataAssetSweepResponseFetchOperation, "ai_anomaly_detection", true)
		}
	}
	return nil
}

func getDataAssetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DataAssetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	anomalyDetectionClient := acctest.GetTestClients(&schema.ResourceData{}).AnomalyDetectionClient()

	listDataAssetsRequest := oci_ai_anomaly_detection.ListDataAssetsRequest{}
	listDataAssetsRequest.CompartmentId = &compartmentId
	listDataAssetsRequest.LifecycleState = oci_ai_anomaly_detection.DataAssetLifecycleStateActive
	listDataAssetsResponse, err := anomalyDetectionClient.ListDataAssets(context.Background(), listDataAssetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DataAsset list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dataAsset := range listDataAssetsResponse.Items {
		id := *dataAsset.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DataAssetId", id)
	}
	return resourceIds, nil
}

func dataAssetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dataAssetResponse, ok := response.Response.(oci_ai_anomaly_detection.GetDataAssetResponse); ok {
		return dataAssetResponse.LifecycleState != oci_ai_anomaly_detection.DataAssetLifecycleStateDeleted
	}
	return false
}

func dataAssetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AnomalyDetectionClient().GetDataAsset(context.Background(), oci_ai_anomaly_detection.GetDataAssetRequest{
		DataAssetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
