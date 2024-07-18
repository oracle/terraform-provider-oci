// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	//"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafeDiscoveryJobRequiredOnlyResource = DataSafeDiscoveryJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_job", "test_discovery_job", acctest.Required, acctest.Create, discoveryJobRepresentation)

	DataSafeDiscoveryJobResourceConfig = DataSafeDiscoveryJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_job", "test_discovery_job", acctest.Optional, acctest.Update, discoveryJobRepresentation)

	DataSafediscoveryJobSingularDataSourceRepresentation = map[string]interface{}{
		"discovery_job_id": acctest.Representation{RepType: acctest.Required, Create: `${var.discovery_job_id}`},
	}

	DataSafediscoveryJobDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"discovery_job_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.discovery_job_id}`},
		"sensitive_data_model_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
		"state":                   acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: discoveryJobDataSourceFilterRepresentation}}
	discoveryJobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_discovery_job.test_discovery_job.id}`}},
	}

	discoveryJobRepresentation = map[string]interface{}{
		"compartment_id":                            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"sensitive_data_model_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_sensitive_data_model.test_sensitive_data_model.id}`},
		"discovery_type":                            acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
		"freeform_tags":                             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_app_defined_relation_discovery_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_include_all_schemas":                    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_include_all_sensitive_types":            acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_sample_data_collection_enabled":         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"schemas_for_discovery":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`ZX`}},
		"sensitive_type_ids_for_discovery":          acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.sensitive_type_id}`}},
		"tables_for_discovery":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: DataSafeDiscoveryJobTablesForDiscoveryRepresentation},
	}
	DataSafeDiscoveryJobTablesForDiscoveryRepresentation = map[string]interface{}{
		"schema_name": acctest.Representation{RepType: acctest.Required, Create: `ZX`},
		"table_names": acctest.Representation{RepType: acctest.Optional, Create: []string{`ZX_API_REGISTRATIONS`}},
	}

	DataSafeDiscoveryJobResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Required, acctest.Create, sensitiveDataModelRepresentation)

	ignoreDiscoveryJobSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`, `freeform_tags`}},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeDiscoveryJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeDiscoveryJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	discoveryJobId := utils.GetEnvSettingWithBlankDefault("data_safe_discovery_job_ocid")
	discoveryJobIdVariableStr := fmt.Sprintf("variable \"discovery_job_id\" { default = \"%s\" }\n", discoveryJobId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	sensitiveTypeId := utils.GetEnvSettingWithBlankDefault("sensitive_type_id")
	sensitiveTypeIdVariableStr := fmt.Sprintf("variable \"sensitive_type_id\" { default = \"%s\" }\n", sensitiveTypeId)

	resourceName := "oci_data_safe_discovery_job.test_discovery_job"
	datasourceName := "data.oci_data_safe_discovery_jobs.test_discovery_jobs"
	singularDatasourceName := "data.oci_data_safe_discovery_job.test_discovery_job"
	//
	var resId string
	var resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeDiscoveryJobResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_job", "test_discovery_job", acctest.Optional, acctest.Create, discoveryJobRepresentation), "datasafe", "discoveryJob", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeDiscoveryJobResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_job", "test_discovery_job", acctest.Required, acctest.Create, discoveryJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeDiscoveryJobResourceDependencies + targetIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DataSafeDiscoveryJobResourceDependencies + targetIdVariableStr + sensitiveTypeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_job", "test_discovery_job", acctest.Optional, acctest.Create, discoveryJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "discovery_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_app_defined_relation_discovery_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_include_all_schemas", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_include_all_sensitive_types", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sample_data_collection_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "tables_for_discovery.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tables_for_discovery.0.table_names.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(resourceName, "time_started"),
				resource.TestCheckResourceAttrSet(resourceName, "total_columns_scanned"),
				resource.TestCheckResourceAttrSet(resourceName, "total_deleted_sensitive_columns"),
				resource.TestCheckResourceAttrSet(resourceName, "total_modified_sensitive_columns"),
				resource.TestCheckResourceAttrSet(resourceName, "total_new_sensitive_columns"),
				resource.TestCheckResourceAttrSet(resourceName, "total_objects_scanned"),
				resource.TestCheckResourceAttrSet(resourceName, "total_schemas_scanned"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DataSafeDiscoveryJobResourceDependencies + targetIdVariableStr + sensitiveTypeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_job", "test_discovery_job", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(discoveryJobRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "discovery_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_app_defined_relation_discovery_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_include_all_schemas", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_include_all_sensitive_types", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_sample_data_collection_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "sensitive_data_model_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(resourceName, "time_started"),
				resource.TestCheckResourceAttrSet(resourceName, "total_columns_scanned"),
				resource.TestCheckResourceAttrSet(resourceName, "total_deleted_sensitive_columns"),
				resource.TestCheckResourceAttrSet(resourceName, "total_modified_sensitive_columns"),
				resource.TestCheckResourceAttrSet(resourceName, "total_new_sensitive_columns"),
				resource.TestCheckResourceAttrSet(resourceName, "total_objects_scanned"),
				resource.TestCheckResourceAttrSet(resourceName, "total_schemas_scanned"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config + targetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_discovery_jobs", "test_discovery_jobs", acctest.Required, acctest.Create, DataSafediscoveryJobDataSourceRepresentation) +
				compartmentIdVariableStr + discoveryJobIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "discovery_job_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + targetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_discovery_job", "test_discovery_job", acctest.Required, acctest.Create, DataSafediscoveryJobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + discoveryJobIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "discovery_job_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "discovery_type", "ALL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_app_defined_relation_discovery_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_include_all_schemas", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_include_all_sensitive_types", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_sample_data_collection_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tables_for_discovery.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tables_for_discovery.0.schema_name", "ZX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tables_for_discovery.0.table_names.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_columns_scanned"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_deleted_sensitive_columns"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_modified_sensitive_columns"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_new_sensitive_columns"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_objects_scanned"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_schemas_scanned"),
			),
		},

		// verify resource import
		{
			Config:                  config + targetIdVariableStr + DataSafeDiscoveryJobRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{`state`},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeDiscoveryJobDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_discovery_job" {
			noResourceFound = false
			request := oci_data_safe.GetDiscoveryJobRequest{}

			tmp := rs.Primary.ID
			request.DiscoveryJobId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetDiscoveryJob(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.DiscoveryLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeDiscoveryJob") {
		resource.AddTestSweepers("DataSafeDiscoveryJob", &resource.Sweeper{
			Name:         "DataSafeDiscoveryJob",
			Dependencies: acctest.DependencyGraph["discoveryJob"],
			F:            sweepDataSafeDiscoveryJobResource,
		})
	}
}

func sweepDataSafeDiscoveryJobResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	discoveryJobIds, err := getDataSafeDiscoveryJobIds(compartment)
	if err != nil {
		return err
	}
	for _, discoveryJobId := range discoveryJobIds {
		if ok := acctest.SweeperDefaultResourceId[discoveryJobId]; !ok {
			deleteDiscoveryJobRequest := oci_data_safe.DeleteDiscoveryJobRequest{}

			deleteDiscoveryJobRequest.DiscoveryJobId = &discoveryJobId

			deleteDiscoveryJobRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteDiscoveryJob(context.Background(), deleteDiscoveryJobRequest)
			if error != nil {
				fmt.Printf("Error deleting DiscoveryJob %s %s, It is possible that the resource is already deleted. Please verify manually \n", discoveryJobId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &discoveryJobId, DataSafediscoveryJobsSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafediscoveryJobsSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeDiscoveryJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DiscoveryJobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listDiscoveryJobsRequest := oci_data_safe.ListDiscoveryJobsRequest{}
	listDiscoveryJobsRequest.CompartmentId = &compartmentId
	listDiscoveryJobsRequest.LifecycleState = oci_data_safe.ListDiscoveryJobsLifecycleStateActive
	listDiscoveryJobsResponse, err := dataSafeClient.ListDiscoveryJobs(context.Background(), listDiscoveryJobsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DiscoveryJob list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, discoveryJob := range listDiscoveryJobsResponse.Items {
		id := *discoveryJob.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DiscoveryJobId", id)
	}
	return resourceIds, nil
}

func DataSafediscoveryJobsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if discoveryJobResponse, ok := response.Response.(oci_data_safe.GetDiscoveryJobResponse); ok {
		return discoveryJobResponse.LifecycleState != oci_data_safe.DiscoveryLifecycleStateDeleted
	}
	return false
}

func DataSafediscoveryJobsSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetDiscoveryJob(context.Background(), oci_data_safe.GetDiscoveryJobRequest{
		DiscoveryJobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
