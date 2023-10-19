// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	DatascienceJobRequiredOnlyResource = DatascienceJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Required, acctest.Create, DatascienceJobRepresentation)

	DatascienceJobResourceConfig = DatascienceJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Update, DatascienceJobRepresentation)

	DatascienceDatascienceJobSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_job.test_job.id}`},
	}

	DatascienceDatascienceJobDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_job.test_job.created_by}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_job.test_job.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobDataSourceFilterRepresentation},
	}

	DatascienceJobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_job.test_job.id}`}},
	}

	DatascienceJobRepresentation = map[string]interface{}{
		"compartment_id":                           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"job_configuration_details":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobJobConfigurationDetailsRepresentation},
		"job_infrastructure_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobJobInfrastructureConfigurationDetailsRepresentation},
		"project_id":                               acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"job_artifact":                             acctest.Representation{RepType: acctest.Optional, Create: `../../examples/datascience/job-artifact.py`},
		"artifact_content_length":                  acctest.Representation{RepType: acctest.Optional, Create: `1380`}, // wc -c job-artifact.py
		"artifact_content_disposition":             acctest.Representation{RepType: acctest.Optional, Create: `attachment; filename=job-artifact.py`},
		"defined_tags":                             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                              acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                             acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"delete_related_job_runs":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMlJobDefinedTagsChangesRepresentation},
	}
	DatascienceJobJobConfigurationDetailsRepresentation = map[string]interface{}{
		"job_type":                   acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"command_line_arguments":     acctest.Representation{RepType: acctest.Optional, Create: `commandLineArguments`},
		"environment_variables":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"environmentVariables": "environmentVariables"}},
		"maximum_runtime_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DatascienceJobJobInfrastructureConfigurationDetailsRepresentation = map[string]interface{}{
		"block_storage_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `50`, Update: `100`},
		"job_infrastructure_type":   acctest.Representation{RepType: acctest.Required, Create: `STANDALONE`},
		"shape_name":                acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	ignoreMlJobDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `job_infrastructure_configuration_details`}},
	}

	// easier to work with from JobRuns
	mlJobWithArtifactNoLogging = map[string]interface{}{
		"compartment_id":                           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"job_configuration_details":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobJobConfigurationDetailsRepresentation},
		"job_infrastructure_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobJobInfrastructureConfigurationDetailsRepresentation},
		"project_id":                               acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"job_artifact":                             acctest.Representation{RepType: acctest.Required, Create: `../../examples/datascience/job-artifact.py`},
		"artifact_content_length":                  acctest.Representation{RepType: acctest.Required, Create: `1380`}, // wc -c job-artifact.py
		"artifact_content_disposition":             acctest.Representation{RepType: acctest.Required, Create: `attachment; filename=job-artifact.py`},
		"lifecycle":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMlJobDefinedTagsChangesRepresentation},
	}

	DatascienceJobResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", acctest.Required, acctest.Create, CoreCoreShapeDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceJobResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_job.test_job"
	datasourceName := "data.oci_datascience_jobs.test_jobs"
	singularDatasourceName := "data.oci_datascience_job.test_job"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceJobResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Create, DatascienceJobRepresentation), "datascience", "job", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatascienceJobDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + DatascienceJobResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Required, acctest.Create, DatascienceJobRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.shape_name"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + DatascienceJobResourceDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + DatascienceJobResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Create, DatascienceJobRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.shape_name"),
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
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceJobResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(DatascienceJobRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
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
				Config: config + compartmentIdVariableStr + DatascienceJobResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Update, DatascienceJobRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.shape_name"),
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_jobs", "test_jobs", acctest.Optional, acctest.Update, DatascienceDatascienceJobDataSourceRepresentation) +
					compartmentIdVariableStr + DatascienceJobResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Update, DatascienceJobRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(datasourceName, "jobs.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "jobs.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "jobs.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "jobs.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "jobs.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "jobs.0.project_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "jobs.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "jobs.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Required, acctest.Create, DatascienceDatascienceJobSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DatascienceJobResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "job_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// verify resource import
			{
				Config:            config + DatascienceJobRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"artifact_content_disposition",
					"artifact_content_length",
					"lifecycle_details",
					"delete_related_job_runs",
					"job_artifact",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckDatascienceJobDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		fmt.Printf("TYPE OF RS: %s\n", rs.Type)
		if rs.Type == "oci_datascience_job" {
			noResourceFound = false
			request := oci_datascience.GetJobRequest{}

			tmp := rs.Primary.ID
			request.JobId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetJob(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.JobLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceJob") {
		resource.AddTestSweepers("DatascienceJob", &resource.Sweeper{
			Name:         "DatascienceJob",
			Dependencies: acctest.DependencyGraph["job"],
			F:            sweepDatascienceJobResource,
		})
	}
}

func sweepDatascienceJobResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	jobIds, err := getDatascienceJobIds(compartment)
	if err != nil {
		return err
	}
	for _, jobId := range jobIds {
		if ok := acctest.SweeperDefaultResourceId[jobId]; !ok {
			deleteJobRequest := oci_datascience.DeleteJobRequest{}

			deleteJobRequest.JobId = &jobId

			deleteJobRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteJob(context.Background(), deleteJobRequest)
			if error != nil {
				fmt.Printf("Error deleting Job %s %s, It is possible that the resource is already deleted. Please verify manually \n", jobId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &jobId, DatascienceJobSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceJobSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "JobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listJobsRequest := oci_datascience.ListJobsRequest{}
	listJobsRequest.CompartmentId = &compartmentId
	listJobsRequest.LifecycleState = oci_datascience.ListJobsLifecycleStateActive
	listJobsResponse, err := dataScienceClient.ListJobs(context.Background(), listJobsRequest)

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

func DatascienceJobSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if jobResponse, ok := response.Response.(oci_datascience.GetJobResponse); ok {
		return jobResponse.LifecycleState != oci_datascience.JobLifecycleStateDeleted
	}
	return false
}

func DatascienceJobSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetJob(context.Background(), oci_datascience.GetJobRequest{
		JobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
