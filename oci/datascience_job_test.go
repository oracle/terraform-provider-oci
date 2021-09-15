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
	"github.com/oracle/oci-go-sdk/v47/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v47/datascience"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	mlJobRequiredOnlyResource = mlJobResourceDependencies +
		generateResourceFromRepresentationMap("oci_datascience_job", "test_job", Required, Create, mlJobRepresentation)

	mlJobResourceConfig = mlJobResourceDependencies +
		generateResourceFromRepresentationMap("oci_datascience_job", "test_job", Optional, Update, mlJobRepresentation)

	mlJobSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": Representation{repType: Required, create: `${oci_datascience_job.test_job.id}`},
	}

	mlJobDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"created_by":     Representation{repType: Optional, create: `${oci_datascience_job.test_job.created_by}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"id":             Representation{repType: Optional, create: `${oci_datascience_job.test_job.id}`},
		"project_id":     Representation{repType: Optional, create: `${oci_datascience_project.test_project.id}`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, mlJobDataSourceFilterRepresentation},
	}

	mlJobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_datascience_job.test_job.id}`}},
	}

	mlJobRepresentation = map[string]interface{}{
		"compartment_id":                           Representation{repType: Required, create: `${var.compartment_id}`},
		"job_configuration_details":                RepresentationGroup{Required, jobJobConfigurationDetailsRepresentation},
		"job_infrastructure_configuration_details": RepresentationGroup{Required, jobJobInfrastructureConfigurationDetailsRepresentation},
		"project_id":                               Representation{repType: Required, create: `${oci_datascience_project.test_project.id}`},
		"job_artifact":                             Representation{repType: Optional, create: `../examples/datascience/job-artifact.py`},
		"artifact_content_length":                  Representation{repType: Optional, create: `1380`}, // wc -c job-artifact.py
		"artifact_content_disposition":             Representation{repType: Optional, create: `attachment; filename=job-artifact.py`},
		"defined_tags":                             Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                              Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":                             Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":                            Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"delete_related_job_runs":                  Representation{repType: Optional, create: `false`, update: `true`},
		"job_log_configuration_details":            RepresentationGroup{Optional, jobJobLogConfigurationDetailsRepresentation},
		"lifecycle":                                RepresentationGroup{Required, ignoreMlJobDefinedTagsChangesRepresentation},
	}
	jobJobConfigurationDetailsRepresentation = map[string]interface{}{
		"job_type":                   Representation{repType: Required, create: `DEFAULT`},
		"command_line_arguments":     Representation{repType: Optional, create: `commandLineArguments`},
		"environment_variables":      Representation{repType: Optional, create: map[string]string{"environmentVariables": "environmentVariables"}},
		"maximum_runtime_in_minutes": Representation{repType: Optional, create: `10`},
	}
	jobJobInfrastructureConfigurationDetailsRepresentation = map[string]interface{}{
		"block_storage_size_in_gbs": Representation{repType: Required, create: `51`, update: `52`},
		"job_infrastructure_type":   Representation{repType: Required, create: `STANDALONE`},
		"shape_name":                Representation{repType: Required, create: `VM.Standard2.2`, update: `VM.Standard2.4`},
		"subnet_id":                 Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
	}
	jobJobLogConfigurationDetailsRepresentation = map[string]interface{}{
		"enable_auto_log_creation": Representation{repType: Optional, create: `true`},
		"enable_logging":           Representation{repType: Optional, create: `true`},
		"log_group_id":             Representation{repType: Optional, create: `${oci_logging_log_group.test_log_group.id}`},
	}

	ignoreMlJobDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": Representation{repType: Required, create: []string{`defined_tags`}},
	}

	// easier to work with from JobRuns
	mlJobWithArtifactNoLogging = map[string]interface{}{
		"compartment_id":                           Representation{repType: Required, create: `${var.compartment_id}`},
		"job_configuration_details":                RepresentationGroup{Required, jobJobConfigurationDetailsRepresentation},
		"job_infrastructure_configuration_details": RepresentationGroup{Required, jobJobInfrastructureConfigurationDetailsRepresentation},
		"project_id":                               Representation{repType: Required, create: `${oci_datascience_project.test_project.id}`},
		"job_artifact":                             Representation{repType: Required, create: `../examples/datascience/job-artifact.py`},
		"artifact_content_length":                  Representation{repType: Required, create: `1380`}, // wc -c job-artifact.py
		"artifact_content_disposition":             Representation{repType: Required, create: `attachment; filename=job-artifact.py`},
		"lifecycle":                                RepresentationGroup{Required, ignoreMlJobDefinedTagsChangesRepresentation},
	}

	mlJobResourceDependencies = generateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", Required, Create, shapeDataSourceRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		generateResourceFromRepresentationMap("oci_datascience_project", "test_project", Required, Create, projectRepresentation) +
		DefinedTagsDependencies +
		generateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupRepresentation)
)

// issue-routing-tag: datascience/default
func TestDatascienceJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceJobResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_job.test_job"
	datasourceName := "data.oci_datascience_jobs.test_jobs"
	singularDatasourceName := "data.oci_datascience_job.test_job"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+mlJobResourceDependencies+
		generateResourceFromRepresentationMap("oci_datascience_job", "test_job", Optional, Create, mlJobRepresentation), "datascience", "job", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatascienceJobDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + mlJobResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_job", "test_job", Required, Create, mlJobRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "51"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.job_infrastructure_type", "STANDALONE"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.shape_name"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + mlJobResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + mlJobResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_job", "test_job", Optional, Create, mlJobRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "51"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.job_infrastructure_type", "STANDALONE"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.shape_name"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_details.0.enable_auto_log_creation", "true"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_details.0.enable_logging", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "job_log_configuration_details.0.log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + mlJobResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_job", "test_job", Optional, Create,
						representationCopyWithNewProperties(mlJobRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "51"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.job_infrastructure_type", "STANDALONE"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.shape_name"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_details.0.enable_auto_log_creation", "true"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_details.0.enable_logging", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "job_log_configuration_details.0.log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + mlJobResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_job", "test_job", Optional, Update, mlJobRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "52"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.job_infrastructure_type", "STANDALONE"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.shape_name", "VM.Standard2.4"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.subnet_id"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_details.0.enable_auto_log_creation", "true"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_details.0.enable_logging", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "job_log_configuration_details.0.log_group_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource - step 5
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_datascience_jobs", "test_jobs", Optional, Update, mlJobDataSourceRepresentation) +
					compartmentIdVariableStr + mlJobResourceDependencies +
					generateResourceFromRepresentationMap("oci_datascience_job", "test_job", Optional, Update, mlJobRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "created_by"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "jobs.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "jobs.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "jobs.0.created_by"),
					resource.TestCheckResourceAttr(datasourceName, "jobs.0.defined_tags.%", "3"),
					resource.TestCheckResourceAttr(datasourceName, "jobs.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "jobs.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "jobs.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "jobs.0.project_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "jobs.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "jobs.0.time_created"),
				),
			},
			// verify singular datasource - step 6
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_datascience_job", "test_job", Required, Create, mlJobSingularDataSourceRepresentation) +
					compartmentIdVariableStr + mlJobResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "job_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "52"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_infrastructure_configuration_details.0.job_infrastructure_type", "STANDALONE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_log_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_log_configuration_details.0.enable_auto_log_creation", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_log_configuration_details.0.enable_logging", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + mlJobResourceConfig,
			},
			// verify resource import - step 8
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"artifact_content_disposition",
					"artifact_content_length",
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
	client := testAccProvider.Meta().(*OracleClients).dataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_job" {
			noResourceFound = false
			request := oci_datascience.GetJobRequest{}

			tmp := rs.Primary.ID
			request.JobId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datascience")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatascienceJob") {
		resource.AddTestSweepers("DatascienceJob", &resource.Sweeper{
			Name:         "DatascienceJob",
			Dependencies: DependencyGraph["job"],
			F:            sweepDatascienceJobResource,
		})
	}
}

func sweepDatascienceJobResource(compartment string) error {
	dataScienceClient := GetTestClients(&schema.ResourceData{}).dataScienceClient()
	jobIds, err := getMlJobIds(compartment)
	if err != nil {
		return err
	}
	for _, jobId := range jobIds {
		if ok := SweeperDefaultResourceId[jobId]; !ok {
			deleteJobRequest := oci_datascience.DeleteJobRequest{}

			deleteJobRequest.JobId = &jobId

			deleteJobRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteJob(context.Background(), deleteJobRequest)
			if error != nil {
				fmt.Printf("Error deleting Job %s %s, It is possible that the resource is already deleted. Please verify manually \n", jobId, error)
				continue
			}
			waitTillCondition(testAccProvider, &jobId, mlJobSweepWaitCondition, time.Duration(3*time.Minute),
				mlJobSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getMlJobIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "JobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := GetTestClients(&schema.ResourceData{}).dataScienceClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "JobId", id)
	}
	return resourceIds, nil
}

func mlJobSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if jobResponse, ok := response.Response.(oci_datascience.GetJobResponse); ok {
		return jobResponse.LifecycleState != oci_datascience.JobLifecycleStateDeleted
	}
	return false
}

func mlJobSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataScienceClient().GetJob(context.Background(), oci_datascience.GetJobRequest{
		JobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
