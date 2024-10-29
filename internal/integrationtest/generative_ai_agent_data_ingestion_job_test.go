// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiAgentDataIngestionJobRequiredOnlyResource = GenerativeAiAgentDataIngestionJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_ingestion_job", "test_data_ingestion_job", acctest.Required, acctest.Create, GenerativeAiAgentDataIngestionJobRepresentation)

	GenerativeAiAgentDataIngestionJobResourceConfig = GenerativeAiAgentDataIngestionJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_ingestion_job", "test_data_ingestion_job", acctest.Optional, acctest.Update, GenerativeAiAgentDataIngestionJobRepresentation)

	GenerativeAiAgentDataIngestionJobSingularDataSourceRepresentation = map[string]interface{}{
		"data_ingestion_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_agent_data_ingestion_job.test_data_ingestion_job.id}`},
	}

	GenerativeAiAgentDataIngestionJobDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"data_source_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.dataSource_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `SUCCEEDED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiAgentDataIngestionJobDataSourceFilterRepresentation}}
	GenerativeAiAgentDataIngestionJobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_agent_data_ingestion_job.test_data_ingestion_job.id}`}},
	}

	GenerativeAiAgentDataIngestionJobRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_source_id": acctest.Representation{RepType: acctest.Required, Create: `${var.dataSource_id}`},
		// "defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
	}

	GenerativeAiAgentDataIngestionJobResourceDependencies = ``
)

// issue-routing-tag: generative_ai_agent/default
func TestGenerativeAiAgentDataIngestionJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiAgentDataIngestionJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dataSourceId := utils.GetEnvSettingWithBlankDefault("dataSource_ocid_for_create")
	dataSourceIdVariableStr := fmt.Sprintf("variable \"dataSource_id\" { default = \"%s\" }\n", dataSourceId)

	resourceName := "oci_generative_ai_agent_data_ingestion_job.test_data_ingestion_job"
	datasourceName := "data.oci_generative_ai_agent_data_ingestion_jobs.test_data_ingestion_jobs"
	singularDatasourceName := "data.oci_generative_ai_agent_data_ingestion_job.test_data_ingestion_job"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dataSourceIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_ingestion_job", "test_data_ingestion_job", acctest.Optional, acctest.Create, GenerativeAiAgentDataIngestionJobRepresentation), "generativeaiagent", "dataIngestionJob", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiAgentDataIngestionJobDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + dataSourceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_ingestion_job", "test_data_ingestion_job", acctest.Required, acctest.Create, GenerativeAiAgentDataIngestionJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "data_source_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + dataSourceIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + dataSourceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_ingestion_job", "test_data_ingestion_job", acctest.Optional, acctest.Create, GenerativeAiAgentDataIngestionJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_ingestion_job_statistics.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_data_ingestion_jobs", "test_data_ingestion_jobs", acctest.Optional, acctest.Update, GenerativeAiAgentDataIngestionJobDataSourceRepresentation) +
				compartmentIdVariableStr + dataSourceIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_agent_data_ingestion_job", "test_data_ingestion_job", acctest.Optional, acctest.Update, GenerativeAiAgentDataIngestionJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "data_source_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),

				resource.TestCheckResourceAttr(datasourceName, "data_ingestion_job_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "data_ingestion_job_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_agent_data_ingestion_job", "test_data_ingestion_job", acctest.Required, acctest.Create, GenerativeAiAgentDataIngestionJobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dataSourceIdVariableStr + GenerativeAiAgentDataIngestionJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "data_ingestion_job_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_ingestion_job_statistics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + GenerativeAiAgentDataIngestionJobRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckGenerativeAiAgentDataIngestionJobDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_agent_data_ingestion_job" {
			noResourceFound = false
			request := oci_generative_ai_agent.GetDataIngestionJobRequest{}

			tmp := rs.Primary.ID
			request.DataIngestionJobId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")

			response, err := client.GetDataIngestionJob(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai_agent.DataIngestionJobLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiAgentDataIngestionJob") {
		resource.AddTestSweepers("GenerativeAiAgentDataIngestionJob", &resource.Sweeper{
			Name:         "GenerativeAiAgentDataIngestionJob",
			Dependencies: acctest.DependencyGraph["dataIngestionJob"],
			F:            sweepGenerativeAiAgentDataIngestionJobResource,
		})
	}
}

func sweepGenerativeAiAgentDataIngestionJobResource(compartment string) error {
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()
	dataIngestionJobIds, err := getGenerativeAiAgentDataIngestionJobIds(compartment)
	if err != nil {
		return err
	}
	for _, dataIngestionJobId := range dataIngestionJobIds {
		if ok := acctest.SweeperDefaultResourceId[dataIngestionJobId]; !ok {
			deleteDataIngestionJobRequest := oci_generative_ai_agent.DeleteDataIngestionJobRequest{}

			deleteDataIngestionJobRequest.DataIngestionJobId = &dataIngestionJobId

			deleteDataIngestionJobRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai_agent")
			_, error := generativeAiAgentClient.DeleteDataIngestionJob(context.Background(), deleteDataIngestionJobRequest)
			if error != nil {
				fmt.Printf("Error deleting DataIngestionJob %s %s, It is possible that the resource is already deleted. Please verify manually \n", dataIngestionJobId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dataIngestionJobId, GenerativeAiAgentDataIngestionJobSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiAgentDataIngestionJobSweepResponseFetchOperation, "generative_ai_agent", true)
		}
	}
	return nil
}

func getGenerativeAiAgentDataIngestionJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DataIngestionJobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiAgentClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiAgentClient()

	listDataIngestionJobsRequest := oci_generative_ai_agent.ListDataIngestionJobsRequest{}
	listDataIngestionJobsRequest.CompartmentId = &compartmentId
	listDataIngestionJobsRequest.LifecycleState = oci_generative_ai_agent.DataIngestionJobLifecycleStateSucceeded
	listDataIngestionJobsResponse, err := generativeAiAgentClient.ListDataIngestionJobs(context.Background(), listDataIngestionJobsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DataIngestionJob list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dataIngestionJob := range listDataIngestionJobsResponse.Items {
		id := *dataIngestionJob.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DataIngestionJobId", id)
	}
	return resourceIds, nil
}

func GenerativeAiAgentDataIngestionJobSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dataIngestionJobResponse, ok := response.Response.(oci_generative_ai_agent.GetDataIngestionJobResponse); ok {
		return dataIngestionJobResponse.LifecycleState != oci_generative_ai_agent.DataIngestionJobLifecycleStateDeleted
	}
	return false
}

func GenerativeAiAgentDataIngestionJobSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiAgentClient().GetDataIngestionJob(context.Background(), oci_generative_ai_agent.GetDataIngestionJobRequest{
		DataIngestionJobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
