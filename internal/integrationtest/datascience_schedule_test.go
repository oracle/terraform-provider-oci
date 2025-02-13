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
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatascienceScheduleRequiredOnlyResource = DatascienceScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_schedule", "test_schedule", acctest.Required, acctest.Create, DatascienceScheduleRepresentation)

	DatascienceScheduleResourceConfig = DatascienceScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_schedule", "test_schedule", acctest.Optional, acctest.Update, DatascienceScheduleRepresentation)

	DatascienceScheduleSingularDataSourceRepresentation = map[string]interface{}{
		"schedule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_schedule.test_schedule.id}`},
	}

	DatascienceScheduleDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_schedule.test_schedule.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.schedule_project.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceScheduleDataSourceFilterRepresentation}}

	DatascienceScheduleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_schedule.test_schedule.id}`}},
	}

	DatascienceScheduleRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.schedule_project.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"action":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceScheduleActionRepresentation},
		"trigger":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceScheduleCronTriggerRepresentation},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		//"log_details":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceScheduleLogDetailsRepresentation},
	}

	//DatascienceScheduleLogDetailsRepresentation = map[string]interface{}{
	//	"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
	//	"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_log.id}`},
	//}

	DatascienceScheduleActionRepresentation = map[string]interface{}{
		"action_type":    acctest.Representation{RepType: acctest.Required, Create: `HTTP`},
		"action_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceCreateJobRunScheduleActionDetailsRepresentation},
	}

	DatascienceScheduleCronTriggerRepresentation = map[string]interface{}{
		"trigger_type":    acctest.Representation{RepType: acctest.Required, Create: `CRON`, Update: `CRON`},
		"time_start":      acctest.Representation{RepType: acctest.Optional, Create: `2025-07-21T16:11:29Z`, Update: `2026-07-21T16:11:29Z`},
		"time_end":        acctest.Representation{RepType: acctest.Optional, Create: `2027-07-21T16:11:29Z`, Update: `2028-07-21T16:11:29Z`},
		"cron_expression": acctest.Representation{RepType: acctest.Required, Create: `10 * * * *`, Update: `11 * * * *`},
	}

	DatascienceCreateJobRunScheduleActionDetailsRepresentation = map[string]interface{}{
		"http_action_type":       acctest.Representation{RepType: acctest.Required, Create: `CREATE_JOB_RUN`, Update: `CREATE_JOB_RUN`},
		"create_job_run_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceScheduleActionActionDetailsCreateJobRunDetailsRepresentation},
	}
	DatascienceScheduleActionActionDetailsCreateJobRunDetailsRepresentation = map[string]interface{}{
		"project_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.schedule_project.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"job_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_job.schedule_job.id}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatascienceScheduleJobShapeConfigurationDetailsRepresentation = map[string]interface{}{
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `8`},
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `256`},
	}

	DatascienceScheduleJobRepresentation = map[string]interface{}{
		"compartment_id":                           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"job_configuration_details":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobJobConfigurationDetailsRepresentation},
		"job_infrastructure_configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobJobInfrastructureConfigurationDetailsRepresentation},
		"job_environment_configuration_details":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatascienceJobJobEnvironmentConfigurationDetailsRepresentation},
		"project_id":                               acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.schedule_project.id}`},
		"job_artifact":                             acctest.Representation{RepType: acctest.Required, Create: `../../examples/datascience/job-artifact.py`},
		"artifact_content_length":                  acctest.Representation{RepType: acctest.Required, Create: `1380`}, // wc -c job-artifact.py
		"artifact_content_disposition":             acctest.Representation{RepType: acctest.Required, Create: `attachment; filename=job-artifact.py`},
		"description":                              acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name":                             acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                            acctest.Representation{RepType: acctest.Required, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"delete_related_job_runs":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"lifecycle":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMlJobDefinedTagsChangesRepresentation},
	}

	DatascienceScheduleResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "schedule_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "schedule_job", acctest.Required, acctest.Create, DatascienceScheduleJobRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: datascience/default
func TestDatascienceScheduleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_schedule.test_schedule"
	datasourceName := "data.oci_datascience_schedules.test_schedules"
	singularDatasourceName := "data.oci_datascience_schedule.test_schedule"

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatascienceScheduleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_schedule", "test_schedule", acctest.Optional, acctest.Create, DatascienceScheduleRepresentation), "datascience", "schedule", t)

	acctest.ResourceTest(t, testAccCheckDatascienceScheduleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatascienceScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_schedule", "test_schedule", acctest.Required, acctest.Create, DatascienceScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "action.0.action_details.0.create_job_run_details.0.job_id"),
				resource.TestCheckResourceAttrSet(resourceName, "action.0.action_details.0.create_job_run_details.0.project_id"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.http_action_type", "CREATE_JOB_RUN"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_type", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "trigger.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.trigger_type", "CRON"),
				resource.TestCheckResourceAttrSet(resourceName, "trigger.0.time_start"),
				//resource.TestCheckResourceAttrSet(resourceName, "trigger.0.time_end"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.cron_expression", "10 * * * *"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatascienceScheduleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatascienceScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_schedule", "test_schedule", acctest.Optional, acctest.Create, DatascienceScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				//check schedule details
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				//check schedule.action details
				resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_type", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "action.0.action_details.0.create_job_run_details.0.project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "action.0.action_details.0.create_job_run_details.0.job_id"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.0.freeform_tags.%", "1"),

				//check schedule.trigger details
				resource.TestCheckResourceAttr(resourceName, "trigger.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.trigger_type", "CRON"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.cron_expression", "10 * * * *"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.time_start", "2025-07-21T16:11:29Z"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.time_end", "2027-07-21T16:11:29Z"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatascienceScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_schedule", "test_schedule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatascienceScheduleRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				//check schedule
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				//check schedule.trigger
				resource.TestCheckResourceAttr(resourceName, "trigger.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.trigger_type", "CRON"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.cron_expression", "10 * * * *"),
				resource.TestCheckResourceAttrSet(resourceName, "trigger.0.time_start"),
				resource.TestCheckResourceAttrSet(resourceName, "trigger.0.time_end"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.time_start", "2025-07-21T16:11:29Z"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.time_end", "2027-07-21T16:11:29Z"),

				//check schedule.action
				resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_type", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.http_action_type", "CREATE_JOB_RUN"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.0.compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "action.0.action_details.0.create_job_run_details.0.project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "action.0.action_details.0.create_job_run_details.0.job_id"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.0.display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.0.freeform_tags.%", "1"),

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
			Config: config + compartmentIdVariableStr + DatascienceScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_schedule", "test_schedule", acctest.Optional, acctest.Update, DatascienceScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

				//check schedule.trigger
				resource.TestCheckResourceAttr(resourceName, "trigger.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.trigger_type", "CRON"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.cron_expression", "11 * * * *"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.time_start", "2026-07-21T16:11:29Z"),
				resource.TestCheckResourceAttr(resourceName, "trigger.0.time_end", "2028-07-21T16:11:29Z"),

				//check schedule.action
				resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_type", "HTTP"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.http_action_type", "CREATE_JOB_RUN"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "action.0.action_details.0.create_job_run_details.0.project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "action.0.action_details.0.create_job_run_details.0.job_id"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "action.0.action_details.0.create_job_run_details.0.freeform_tags.%", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_schedules", "test_schedules", acctest.Optional, acctest.Update, DatascienceScheduleDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_schedule", "test_schedule", acctest.Optional, acctest.Update, DatascienceScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "schedules.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "schedules.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "schedules.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "schedules.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "schedules.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "schedules.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "schedules.0.project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "schedules.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "schedules.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "schedules.0.time_updated"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_schedule", "test_schedule", acctest.Required, acctest.Create, DatascienceScheduleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatascienceScheduleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "schedule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_next_scheduled_run"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),

				resource.TestCheckResourceAttr(singularDatasourceName, "trigger.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trigger.0.cron_expression", "11 * * * *"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trigger.0.time_start", "2026-07-21T16:11:29Z"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trigger.0.time_end", "2028-07-21T16:11:29Z"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trigger.0.trigger_type", "CRON"),

				resource.TestCheckResourceAttr(singularDatasourceName, "action.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action.0.action_type", "HTTP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action.0.action_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action.0.action_details.0.create_job_run_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action.0.action_details.0.create_job_run_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "action.0.action_details.0.create_job_run_details.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action.0.action_details.0.create_job_run_details.0.freeform_tags.%", "1"),
			),
		},

		// verify resource import
		{
			Config:                  config + DatascienceScheduleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatascienceScheduleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_schedule" {
			noResourceFound = false
			request := oci_datascience.GetScheduleRequest{}

			tmp := rs.Primary.ID
			request.ScheduleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

			response, err := client.GetSchedule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ScheduleLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatascienceSchedule") {
		resource.AddTestSweepers("DatascienceSchedule", &resource.Sweeper{
			Name:         "DatascienceSchedule",
			Dependencies: acctest.DependencyGraph["schedule"],
			F:            sweepDatascienceScheduleResource,
		})
	}
}

func sweepDatascienceScheduleResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	scheduleIds, err := getDatascienceScheduleIds(compartment)
	if err != nil {
		return err
	}
	for _, scheduleId := range scheduleIds {
		if ok := acctest.SweeperDefaultResourceId[scheduleId]; !ok {
			deleteScheduleRequest := oci_datascience.DeleteScheduleRequest{}

			deleteScheduleRequest.ScheduleId = &scheduleId

			deleteScheduleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteSchedule(context.Background(), deleteScheduleRequest)
			if error != nil {
				fmt.Printf("Error deleting Schedule %s %s, It is possible that the resource is already deleted. Please verify manually \n", scheduleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &scheduleId, DatascienceScheduleSweepWaitCondition, time.Duration(3*time.Minute),
				DatascienceScheduleSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getDatascienceScheduleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ScheduleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

	listSchedulesRequest := oci_datascience.ListSchedulesRequest{}
	listSchedulesRequest.CompartmentId = &compartmentId
	listSchedulesRequest.LifecycleState = oci_datascience.ListSchedulesLifecycleStateActive
	listSchedulesResponse, err := dataScienceClient.ListSchedules(context.Background(), listSchedulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Schedule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, schedule := range listSchedulesResponse.Items {
		id := *schedule.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ScheduleId", id)
	}
	return resourceIds, nil
}

func DatascienceScheduleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if scheduleResponse, ok := response.Response.(oci_datascience.GetScheduleResponse); ok {
		return scheduleResponse.LifecycleState != oci_datascience.ScheduleLifecycleStateDeleted
	}
	return false
}

func DatascienceScheduleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetSchedule(context.Background(), oci_datascience.GetScheduleRequest{
		ScheduleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
