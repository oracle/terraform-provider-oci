// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AiLanguageJobRequiredOnlyResource = AiLanguageJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_job", "test_job", acctest.Required, acctest.Create, AiLanguageJobRepresentation)

	//AiLanguageJobResourceConfig = AiLanguageJobResourceDependencies +
	//	acctest.GenerateResourceFromRepresentationMap("oci_ai_language_job", "test_job", acctest.Optional, acctest.Update, AiLanguageJobRepresentation)
	//
	//AiLanguageJobSingularDataSourceRepresentation = map[string]interface{}{
	//	"job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_language_job.test_job.id}`},
	//}

	//AiLanguageJobDataSourceRepresentation = map[string]interface{}{
	//	"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	//	"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
	//	"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_language_job.test_job.id}`},
	//	"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	//	"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageJobDataSourceFilterRepresentation}}
	//AiLanguageJobDataSourceFilterRepresentation = map[string]interface{}{
	//	"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
	//	"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_language_job.test_job.id}`}},
	//}

	AiLanguageJobRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"input_location":         acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageJobInputLocationRepresentation},
		"model_metadata_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageJobModelMetadataDetailsRepresentation},
		"output_location":        acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageJobOutputLocationRepresentation},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayNametet12t`, Update: `displayName3`},
		"input_configuration":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: AiLanguageJobInputConfigurationRepresentation},
		"timeouts": acctest.RepresentationGroup{
			RepType: acctest.Optional,
			Group: map[string]interface{}{
				"create": acctest.Representation{
					RepType: acctest.Optional,
					Create:  "120m",
				},
				"update": acctest.Representation{
					RepType: acctest.Optional,
					Create:  "120m",
				},
				"delete": acctest.Representation{
					RepType: acctest.Optional,
					Create:  "120m",
				},
			},
		},
	}
	AiLanguageJobInputLocationRepresentation = map[string]interface{}{
		"bucket":        acctest.Representation{RepType: acctest.Required, Create: `adipanda-test`},
		"location_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_PREFIX`},
		"namespace":     acctest.Representation{RepType: acctest.Required, Create: `idngwwc5ajp5`},
		//"object_names":  acctest.Representation{RepType: acctest.Required, Create: []string{`hotel.csv`}},
		//"prefix":        acctest.Representation{RepType: acctest.Optional, Create: `test/`},
	}

	AiLanguageJobModelMetadataDetailsRepresentation = map[string]interface{}{
		//"configuration": acctest.RepresentationGroup{RepType: acctest.Optional, Group: AiLanguageJobModelMetadataDetailsConfigurationRepresentation},
		"endpoint_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_language_endpoint.test_endpoint.id}`},
		"language_code": acctest.Representation{RepType: acctest.Required, Create: `en`},
		"model_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_language_model.test_model.id}`},
		"model_type":    acctest.Representation{RepType: acctest.Required, Create: `TEXT_CLASSIFICATION`},
	}

	AiLanguageJobOutputLocationRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `adipanda-test`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `idngwwc5ajp5`},
	}
	AiLanguageJobInputConfigurationRepresentation = map[string]interface{}{
		//"configuration": acctest.Representation{
		//	RepType: acctest.Required,
		//	Create: []interface{}{
		//		map[string]interface{}{
		//			"csv": []interface{}{
		//				map[string]interface{}{
		//					"config": map[string]interface{}{
		//						"inputColumn": "review",
		//						"rowId":       "",
		//						"delimiter":   ",",
		//					},
		//				},
		//			},
		//		},
		//	},
		//},
		"configuration":  acctest.Representation{RepType: acctest.Required, Create: []interface{}{}},
		"document_types": acctest.Representation{RepType: acctest.Optional, Create: []string{`TXT`}},
	}
	//AiLanguageJobModelMetadataDetailsConfigurationRepresentation = map[string]interface{}{
	//	"configuration_map": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"configurationMap": "configurationMap"}, Update: map[string]string{"configurationMap2": "configurationMap2"}},
	//}
	//AiLanguageJobInputConfigurationConfigurationRepresentation = map[string]interface{}{
	//	"CSV": acctest.RepresentationGroup{
	//		RepType: acctest.Optional,
	//		Group: map[string]interface{}{
	//			"config": acctest.Representation{
	//				RepType: acctest.Optional,
	//				Create: map[string]interface{}{
	//					"inputColumn": "review",
	//					"rowId":       nil,
	//					"delimiter":   ",",
	//				},
	//			},
	//		},
	//	},
	//}

	AiLanguageJobResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_language_endpoint", "test_endpoint", acctest.Required, acctest.Create, AiLanguageEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Required, acctest.Create, AiLanguageModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_project", "test_project", acctest.Required, acctest.Create, AiLanguageProjectRepresentation) +
		DefinedTagsDependencies
	//acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
	//acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: ai_language/default
func TestAiLanguageJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiLanguageJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_language_job.test_job"
	//datasourceName := "data.oci_ai_language_jobs.test_jobs"
	//singularDatasourceName := "data.oci_ai_language_job.test_job"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiLanguageJobResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_job", "test_job", acctest.Optional, acctest.Create, AiLanguageJobRepresentation), "ailanguage", "job", t)

	acctest.ResourceTest(t, testAccCheckAiLanguageJobDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiLanguageJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_job", "test_job", acctest.Required, acctest.Create, AiLanguageJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "input_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.bucket", "adipanda-test"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.location_type", "OBJECT_STORAGE_PREFIX"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_names.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "model_metadata_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.bucket", "adipanda-test"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.namespace", "idngwwc5ajp5"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiLanguageJobResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiLanguageJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_job", "test_job", acctest.Optional, acctest.Create, AiLanguageJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNametet12t"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_configuration.0.configuration.%", "0"),
				resource.TestCheckResourceAttr(resourceName, "input_configuration.0.document_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.bucket", "adipanda-test"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.location_type", "OBJECT_STORAGE_PREFIX"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_names.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "model_metadata_details.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "model_metadata_details.0.configuration.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_metadata_details.0.endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "model_metadata_details.0.language_code", "en"),
				resource.TestCheckResourceAttrSet(resourceName, "model_metadata_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_metadata_details.0.model_type", "TEXT_CLASSIFICATION"),
				resource.TestCheckResourceAttr(resourceName, "output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.bucket", "adipanda-test"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.namespace", "idngwwc5ajp5"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiLanguageJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_job", "test_job", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AiLanguageJobRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNametet12t"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_configuration.0.configuration.%", "0"),
				resource.TestCheckResourceAttr(resourceName, "input_configuration.0.document_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.bucket", "adipanda-test"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.location_type", "OBJECT_STORAGE_PREFIX"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_names.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "model_metadata_details.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "model_metadata_details.0.configuration.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_metadata_details.0.endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "model_metadata_details.0.language_code", "en"),
				resource.TestCheckResourceAttrSet(resourceName, "model_metadata_details.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "model_metadata_details.0.model_type", "TEXT_CLASSIFICATION"),
				resource.TestCheckResourceAttr(resourceName, "output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.bucket", "adipanda-test"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.namespace", "idngwwc5ajp5"),

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
		//{
		//	Config: config + compartmentIdVariableStr + AiLanguageJobResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_job", "test_job", acctest.Optional, acctest.Update, AiLanguageJobRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		//		resource.TestCheckResourceAttr(resourceName, "description", "description2"),
		//		resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
		//		resource.TestCheckResourceAttrSet(resourceName, "id"),
		//		resource.TestCheckResourceAttr(resourceName, "input_configuration.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "input_configuration.0.configuration.%", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "input_configuration.0.document_types.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "input_location.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "input_location.0.bucket", "bucket"),
		//		resource.TestCheckResourceAttr(resourceName, "input_location.0.location_type", "OBJECT_STORAGE_PREFIX"),
		//		resource.TestCheckResourceAttr(resourceName, "input_location.0.namespace", "namespace"),
		//		resource.TestCheckResourceAttr(resourceName, "input_location.0.object_names.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "input_location.0.prefix", "prefix"),
		//		resource.TestCheckResourceAttr(resourceName, "model_metadata_details.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "model_metadata_details.0.configuration.%", "1"),
		//		resource.TestCheckResourceAttrSet(resourceName, "model_metadata_details.0.endpoint_id"),
		//		resource.TestCheckResourceAttr(resourceName, "model_metadata_details.0.language_code", "en"),
		//		resource.TestCheckResourceAttrSet(resourceName, "model_metadata_details.0.model_id"),
		//		resource.TestCheckResourceAttr(resourceName, "model_metadata_details.0.model_type", "TEXT_CLASSIFICATION"),
		//		resource.TestCheckResourceAttr(resourceName, "output_location.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "output_location.0.bucket", "bucket"),
		//		resource.TestCheckResourceAttr(resourceName, "output_location.0.namespace", "namespace"),
		//		resource.TestCheckResourceAttr(resourceName, "output_location.0.prefix", "prefix"),
		//
		//		func(s *terraform.State) (err error) {
		//			resId2, err = acctest.FromInstanceState(s, resourceName, "id")
		//			if resId != resId2 {
		//				return fmt.Errorf("Resource recreated when it was supposed to be updated.")
		//			}
		//			return err
		//		},
		//	),
		//},
		//// verify datasource
		//{
		//	Config: config +
		//		acctest.GenerateDataSourceFromRepresentationMap("oci_ai_language_jobs", "test_jobs", acctest.Optional, acctest.Update, AiLanguageJobDataSourceRepresentation) +
		//		compartmentIdVariableStr + AiLanguageJobResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_job", "test_job", acctest.Optional, acctest.Update, AiLanguageJobRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
		//		resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
		//		resource.TestCheckResourceAttr(datasourceName, "id", "id"),
		//		resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
		//
		//		resource.TestCheckResourceAttr(datasourceName, "job_collection.#", "1"),
		//		resource.TestCheckResourceAttr(datasourceName, "job_collection.0.items.#", "1"),
		//	),
		//},
		//// verify singular datasource
		//{
		//	Config: config +
		//		acctest.GenerateDataSourceFromRepresentationMap("oci_ai_language_job", "test_job", acctest.Required, acctest.Create, AiLanguageJobSingularDataSourceRepresentation) +
		//		compartmentIdVariableStr + AiLanguageJobResourceConfig,
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "job_id"),
		//
		//		resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "completed_documents"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "failed_documents"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "input_configuration.#", "1"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "input_configuration.0.configuration.%", "1"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "input_configuration.0.document_types.#", "1"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "input_location.#", "1"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.bucket", "bucket"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.location_type", "OBJECT_STORAGE_PREFIX"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.namespace", "namespace"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.object_names.#", "1"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.prefix", "prefix"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "model_metadata_details.#", "1"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "model_metadata_details.0.configuration.%", "1"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "model_metadata_details.0.language_code", "en"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "model_metadata_details.0.model_type", "TEXT_CLASSIFICATION"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "output_location.#", "1"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "output_location.0.bucket", "bucket"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "output_location.0.namespace", "namespace"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "output_location.0.prefix", "prefix"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "pending_documents"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "percent_complete"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "time_completed"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "total_documents"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "ttl_in_days"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "warnings_count"),
		//	),
		//},
		// verify resource import
		{
			Config:                  config + AiLanguageJobRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAiLanguageJobDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AiServiceLanguageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_language_job" {
			noResourceFound = false
			request := oci_ai_language.GetJobRequest{}

			tmp := rs.Primary.ID
			request.JobId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_language")

			response, err := client.GetJob(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_language.JobLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiLanguageJob") {
		resource.AddTestSweepers("AiLanguageJob", &resource.Sweeper{
			Name:         "AiLanguageJob",
			Dependencies: acctest.DependencyGraph["job"],
			F:            sweepAiLanguageJobResource,
		})
	}
}

func sweepAiLanguageJobResource(compartment string) error {
	aiServiceLanguageClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceLanguageClient()
	jobIds, err := getAiLanguageJobIds(compartment)
	if err != nil {
		return err
	}
	for _, jobId := range jobIds {
		if ok := acctest.SweeperDefaultResourceId[jobId]; !ok {
			deleteJobRequest := oci_ai_language.DeleteJobRequest{}

			deleteJobRequest.JobId = &jobId

			deleteJobRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_language")
			_, error := aiServiceLanguageClient.DeleteJob(context.Background(), deleteJobRequest)
			if error != nil {
				fmt.Printf("Error deleting Job %s %s, It is possible that the resource is already deleted. Please verify manually \n", jobId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &jobId, AiLanguageJobSweepWaitCondition, time.Duration(3*time.Minute),
				AiLanguageJobSweepResponseFetchOperation, "ai_language", true)
		}
	}
	return nil
}

func getAiLanguageJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "JobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	aiServiceLanguageClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceLanguageClient()

	listJobsRequest := oci_ai_language.ListJobsRequest{}
	listJobsRequest.CompartmentId = &compartmentId
	listJobsRequest.LifecycleState = oci_ai_language.JobLifecycleStateSucceeded
	listJobsResponse, err := aiServiceLanguageClient.ListJobs(context.Background(), listJobsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Job list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, job := range listJobsResponse.Items {
		id := *job.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "JobId", id)
	}
	return resourceIds, nil
}

func AiLanguageJobSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if jobResponse, ok := response.Response.(oci_ai_language.GetJobResponse); ok {
		return jobResponse.LifecycleState != oci_ai_language.JobLifecycleStateDeleted
	}
	return false
}

func AiLanguageJobSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AiServiceLanguageClient().GetJob(context.Background(), oci_ai_language.GetJobRequest{
		JobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
