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
	DatascienceJobRunRequiredOnlyResource = DatascienceJobRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_job_run", "test_job_run", acctest.Required, acctest.Create, DatascienceJobRunRepresentation)

	DatascienceJobRunResourceConfig = DatascienceJobRunResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_job_run", "test_job_run", acctest.Optional, acctest.Update, DatascienceJobRunRepresentation)

	DatascienceDatascienceJobRunSingularDataSourceRepresentation = map[string]interface{}{
		"job_run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_job_run.test_job_run.id}`},
	}

	DatascienceDatascienceJobRunDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"created_by":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_job_run.test_job_run.created_by}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_job_run.test_job_run.id}`},
		"job_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_job.test_job.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `SUCCEEDED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobRunDataSourceFilterRepresentation},
	}

	DatascienceJobRunDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_job_run.test_job_run.id}`}},
	}

	DatascienceJobRunRepresentation = map[string]interface{}{
		"compartment_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"job_id":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_job.test_job.id}`},
		"project_id":                         acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"asynchronous":                       acctest.Representation{RepType: acctest.Required, Create: `false`},
		"job_configuration_override_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobRunJobConfigurationOverrideDetailsRepresentation},
		"job_environment_configuration_override_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceJobRunJobEnvironmentConfigurationOverrideDetailsRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreJobRunDefinedTagsChangesRepresentation},
	}
	DatascienceJobRunJobConfigurationOverrideDetailsRepresentation = map[string]interface{}{
		"job_type":                   acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"command_line_arguments":     acctest.Representation{RepType: acctest.Optional, Create: `commandLineArguments`},
		"environment_variables":      acctest.Representation{RepType: acctest.Required, Create: map[string]string{"environmentVariables": "environmentVariables"}},
		"maximum_runtime_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DatascienceJobRunJobEnvironmentConfigurationOverrideDetailsRepresentation = map[string]interface{}{
		"image":                acctest.Representation{RepType: acctest.Required, Create: `iad.ocir.io/ociodscdev/byod_hello_wrld:1.0`},
		"job_environment_type": acctest.Representation{RepType: acctest.Required, Create: `OCIR_CONTAINER`},
		"cmd":                  acctest.Representation{RepType: acctest.Optional, Create: []string{``}},
		"entrypoint":           acctest.Representation{RepType: acctest.Optional, Create: []string{``}},
		"image_digest":         acctest.Representation{RepType: acctest.Optional, Create: ``},
		"image_signature_id":   acctest.Representation{RepType: acctest.Optional, Create: ``},
	}
	ignoreJobRunDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	DatascienceJobRunResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", acctest.Required, acctest.Create, CoreCoreShapeDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Required, acctest.Create, mlJobWithArtifactNoLogging) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceDatascienceJobShapeDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceJobRunResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceJobRunResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_datascience_job_run.test_job_run"
	datasourceName := "data.oci_datascience_job_runs.test_job_runs"
	singularDatasourceName := "data.oci_datascience_job_run.test_job_run"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceJobRunResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_job_run", "test_job_run", acctest.Optional, acctest.Create, DatascienceJobRunRepresentation), "datascience", "jobRun", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatascienceJobRunDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + DatascienceJobRunResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job_run", "test_job_run", acctest.Required, acctest.Create, DatascienceJobRunRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "job_id"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + DatascienceJobRunResourceDependencies,
			},

			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + DatascienceJobRunResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job_run", "test_job_run", acctest.Optional, acctest.Create, DatascienceJobRunRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_override_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_override_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_override_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_override_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_override_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.cmd.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.entrypoint.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.image", "iad.ocir.io/ociodscdev/byod_hello_wrld:1.0"),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.image_digest", ""),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.image_signature_id", ""),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.job_environment_type", "OCIR_CONTAINER"),
					resource.TestCheckResourceAttrSet(resourceName, "job_id"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_override_details.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_accepted"),

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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + DatascienceJobRunResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job_run", "test_job_run", acctest.Optional, acctest.Update, DatascienceJobRunRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_override_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_override_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_override_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_override_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_override_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.cmd.#", ""),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.entrypoint.#", ""),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.image", "iad.ocir.io/ociodscdev/byod_hello_wrld:1.0"),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.image_digest", ""),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.image_signature_id", ""),
					resource.TestCheckResourceAttr(resourceName, "job_environment_configuration_override_details.0.job_environment_type", "OCIR_CONTAINER"),
					resource.TestCheckResourceAttrSet(resourceName, "job_id"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_log_configuration_override_details.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_accepted"),

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
					acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_job_runs", "test_job_runs", acctest.Optional, acctest.Update, DatascienceDatascienceJobRunDataSourceRepresentation) +
					compartmentIdVariableStr + DatascienceJobRunResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job_run", "test_job_run", acctest.Optional, acctest.Update, DatascienceJobRunRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "created_by"),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),

					resource.TestCheckResourceAttr(datasourceName, "job_runs.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "job_runs.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "job_runs.0.created_by"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_runs.0.defined_tags.%"),
					resource.TestCheckResourceAttr(datasourceName, "job_runs.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "job_runs.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_runs.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_runs.0.job_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_runs.0.project_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_runs.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "job_runs.0.time_accepted"),
				),
			},

			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_job_run", "test_job_run", acctest.Required, acctest.Create, DatascienceDatascienceJobRunSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DatascienceJobRunResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "job_run_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_override_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_override_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_override_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_override_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_override_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_log_configuration_override_details.#", "0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_environment_configuration_override_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_environment_configuration_override_details.0.cmd.#", ""),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_environment_configuration_override_details.0.entrypoint.#", ""),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_environment_configuration_override_details.0.image", "iad.ocir.io/ociodscdev/byod_hello_wrld:1.0"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_environment_configuration_override_details.0.image_digest", ""),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_environment_configuration_override_details.0.job_environment_type", "OCIR_CONTAINER"),
					resource.TestCheckResourceAttr(singularDatasourceName, "log_details.#", "0"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
				),
			},
			// verify resource import
			{
				Config:            config + DatascienceJobRunRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"asynchronous",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckDatascienceJobRunDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_job_run" {
			noResourceFound = false
			request := oci_datascience.GetJobRunRequest{}

			tmp := rs.Primary.ID
			request.JobRunId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetJobRun(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.JobRunLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceJobRun") {
		resource.AddTestSweepers("DatascienceJobRun", &resource.Sweeper{
			Name:         "DatascienceJobRun",
			Dependencies: acctest.DependencyGraph["jobRun"],
			F:            sweepDatascienceJobRunResource,
		})
	}
}

func sweepDatascienceJobRunResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	jobRunIds, err := getDatascienceJobRunIds(compartment)
	if err != nil {
		return err
	}
	for _, jobRunId := range jobRunIds {
		if ok := acctest.SweeperDefaultResourceId[jobRunId]; !ok {
			deleteJobRunRequest := oci_datascience.DeleteJobRunRequest{}

			deleteJobRunRequest.JobRunId = &jobRunId

			deleteJobRunRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteJobRun(context.Background(), deleteJobRunRequest)
			if error != nil {
				fmt.Printf("Error deleting JobRun %s %s, It is possible that the resource is already deleted. Please verify manually \n", jobRunId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &jobRunId, DatascienceJobRunSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceJobRunSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceJobRunIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "JobRunId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listJobRunsRequest := oci_datascience.ListJobRunsRequest{}
	listJobRunsRequest.CompartmentId = &compartmentId
	listJobRunsRequest.LifecycleState = oci_datascience.ListJobRunsLifecycleStateSucceeded
	listJobRunsResponse, err := dataScienceClient.ListJobRuns(context.Background(), listJobRunsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting JobRun list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, jobRun := range listJobRunsResponse.Items {
		id := *jobRun.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "JobRunId", id)
	}
	return resourceIds, nil
}

func DatascienceJobRunSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if jobRunResponse, ok := response.Response.(oci_datascience.GetJobRunResponse); ok {
		return jobRunResponse.LifecycleState != oci_datascience.JobRunLifecycleStateDeleted
	}
	return false
}

func DatascienceJobRunSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetJobRun(context.Background(), oci_datascience.GetJobRunRequest{
		JobRunId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
