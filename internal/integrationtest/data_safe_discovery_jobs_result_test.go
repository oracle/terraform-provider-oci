// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"terraform-provider-oci/internal/resourcediscovery"

	"terraform-provider-oci/internal/acctest"
	tf_client "terraform-provider-oci/internal/client"
	"terraform-provider-oci/internal/tfresource"
	"terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"terraform-provider-oci/httpreplay"
)

var (
	DataSafeDiscoveryJobsResultResourceConfig = DataSafeDiscoveryJobsResultResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_jobs_result", "test_discovery_jobs_result", acctest.Optional, acctest.Update, discoveryJobsResultRepresentation)

	DataSafediscoveryJobsResultSingularDataSourceRepresentation = map[string]interface{}{
		"discovery_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_discovery_job.test_discovery_job.id}`},
		"result_key":       acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_discovery_jobs_result.test_discovery_jobs_result.key}`},
	}

	DataSafediscoveryJobsResultDataSourceRepresentation = map[string]interface{}{
		"discovery_job_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_discovery_job.test_discovery_job.id}`},
		"discovery_type":    acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
		"is_result_applied": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"planned_action":    acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
		"filter":            acctest.RepresentationGroup{RepType: acctest.Required, Group: discoveryJobsResultDataSourceFilterRepresentation}}
	discoveryJobsResultDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_discovery_jobs_result.test_discovery_jobs_result.key}`}},
	}

	discoveryJobsResultRepresentation = map[string]interface{}{
		"discovery_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_discovery_job.test_discovery_job.id}`},
	}

	discoveryJobsResultRepresentation2 = map[string]interface{}{
		"discovery_job_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_discovery_job.test_discovery_job.id}`},
	}

	DataSafeDiscoveryJobsResultResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_job", "test_discovery_job", acctest.Optional, acctest.Create, discoveryJobRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_sensitive_data_model", "test_sensitive_data_model", acctest.Required, acctest.Create, sensitiveDataModelRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeDiscoveryJobsResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeDiscoveryJobsResultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	resourceName := "oci_data_safe_discovery_jobs_result.test_discovery_jobs_result"
	datasourceName := "data.oci_data_safe_discovery_jobs_results.test_discovery_jobs_results"
	singularDatasourceName := "data.oci_data_safe_discovery_jobs_result.test_discovery_jobs_result"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeDiscoveryJobsResultResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_jobs_result", "test_discovery_jobs_result", acctest.Required, acctest.Create, discoveryJobsResultRepresentation), "datasafe", "discoveryJobsResult", t)

	acctest.ResourceTest(t, testAccCheckDataSafeDiscoveryJobsResultDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeDiscoveryJobsResultResourceDependencies + targetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_jobs_result", "test_discovery_jobs_result", acctest.Optional, acctest.Create, discoveryJobsResultRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				func(s *terraform.State) (err error) {
					var compositeId string
					compositeId, err = acctest.FromInstanceState(s, resourceName, "id")
					prefix := "oci_data_safe_discovery_jobs_result:"
					fullPath := prefix + compositeId
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&fullPath, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config + targetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_discovery_jobs_results", "test_discovery_jobs_results", acctest.Optional, acctest.Update, DataSafediscoveryJobsResultDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeDiscoveryJobsResultResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_discovery_jobs_result", "test_discovery_jobs_result", acctest.Optional, acctest.Update, discoveryJobsResultRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "discovery_job_id"),
				resource.TestCheckResourceAttr(datasourceName, "discovery_type", "ALL"),
				resource.TestCheckResourceAttr(datasourceName, "is_result_applied", "false"),
				resource.TestCheckResourceAttr(datasourceName, "planned_action", "NONE"),
				resource.TestCheckResourceAttr(datasourceName, "discovery_job_result_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + targetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_discovery_jobs_result", "test_discovery_jobs_result", acctest.Required, acctest.Create, DataSafediscoveryJobsResultSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeDiscoveryJobsResultResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "discovery_job_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "discovery_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "estimated_data_value_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_result_applied"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "planned_action"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "relation_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schema_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_columnkey"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sensitive_type_id"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DataSafeDiscoveryJobsResultResourceConfig + targetIdVariableStr,
		},
		// verify resource import
		{
			Config:                  config + targetIdVariableStr,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeDiscoveryJobsResultDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_discovery_jobs_result" {
			noResourceFound = false
			request := oci_data_safe.GetDiscoveryJobResultRequest{}

			if value, ok := rs.Primary.Attributes["discovery_job_id"]; ok {
				request.DiscoveryJobId = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.ResultKey = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			_, err := client.GetDiscoveryJobResult(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("DataSafeDiscoveryJobsResult") {
		resource.AddTestSweepers("DataSafeDiscoveryJobsResult", &resource.Sweeper{
			Name:         "DataSafeDiscoveryJobsResult",
			Dependencies: acctest.DependencyGraph["discoveryJobsResult"],
			F:            sweepDataSafeDiscoveryJobsResultResource,
		})
	}
}

func sweepDataSafeDiscoveryJobsResultResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	discoveryJobsResultIds, err := getDataSafeDiscoveryJobsResultIds(compartment)
	if err != nil {
		return err
	}
	for _, discoveryJobsResultId := range discoveryJobsResultIds {
		if ok := acctest.SweeperDefaultResourceId[discoveryJobsResultId]; !ok {
			deleteDiscoveryJobResultRequest := oci_data_safe.DeleteDiscoveryJobResultRequest{}

			deleteDiscoveryJobResultRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteDiscoveryJobResult(context.Background(), deleteDiscoveryJobResultRequest)
			if error != nil {
				fmt.Printf("Error deleting DiscoveryJobsResult %s %s, It is possible that the resource is already deleted. Please verify manually \n", discoveryJobsResultId, error)
				continue
			}
		}
	}
	return nil
}

func getDataSafeDiscoveryJobsResultIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DiscoveryJobsResultId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listDiscoveryJobResultsRequest := oci_data_safe.ListDiscoveryJobResultsRequest{}
	// listDiscoveryJobResultsRequest.CompartmentId = &compartmentId

	discoveryJobIds, error := getDataSafeDiscoveryJobIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting discoveryJobId required for DiscoveryJobsResult resource requests \n")
	}
	for _, discoveryJobId := range discoveryJobIds {
		listDiscoveryJobResultsRequest.DiscoveryJobId = &discoveryJobId

		listDiscoveryJobResultsResponse, err := dataSafeClient.ListDiscoveryJobResults(context.Background(), listDiscoveryJobResultsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DiscoveryJobsResult list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, discoveryJobsResult := range listDiscoveryJobResultsResponse.Items {
			id := *discoveryJobsResult.Key
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DiscoveryJobsResultId", id)
		}

	}
	return resourceIds, nil
}
