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
	"github.com/oracle/terraform-provider-oci/internal/service/log_analytics"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	LogAnalyticsNamespaceScheduledTaskRequiredOnlyResource = LogAnalyticsNamespaceScheduledTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Required, acctest.Create, LogAnalyticsPurgeTaskRepresentation)

	LogAnalyticsNamespaceScheduledTaskResourceConfig = LogAnalyticsNamespaceScheduledTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Optional, acctest.Update, LogAnalyticsPurgeTaskRepresentation)

	LogAnalyticsNamespaceScheduledTaskMetricExtractionResourceConfig = LogAnalyticsNamespaceScheduledTaskResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Optional, acctest.Update, LogAnalyticsNamespaceScheduledTaskMetricExtractionRepresentation)

	LogAnalyticsLogAnalyticsNamespaceScheduledTaskSingularDataSourceRepresentation = map[string]interface{}{
		"namespace":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"scheduled_task_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task.scheduled_task_id}`},
	}

	namespaceScheduledTaskDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `tfPurgeTask1`, Update: `tfPurgeTask3`},
		"task_type":      acctest.Representation{RepType: acctest.Optional, Create: `PURGE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsNamespaceScheduledTaskDataSourceFilterRepresentation}}

	LogAnalyticsNamespaceScheduledTaskDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task.id}`}},
	}

	LogAnalyticsPurgeTaskRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kind":           acctest.Representation{RepType: acctest.Required, Create: `STANDARD`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"task_type":      acctest.Representation{RepType: acctest.Required, Create: `PURGE`},
		"action":         acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsPurgeActionRepresentation},
		"schedules":      acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsSchedulesRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `tfPurgeTask1`, Update: `tfPurgeTask3`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	LogAnalyticsPurgeActionRepresentation = map[string]interface{}{
		"type":                      acctest.Representation{RepType: acctest.Required, Create: `PURGE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `false`},
		"data_type":                 acctest.Representation{RepType: acctest.Required, Create: `LOG`},
		"purge_compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"purge_duration":            acctest.Representation{RepType: acctest.Required, Create: `-P7D`},
		"query_string":              acctest.Representation{RepType: acctest.Required, Create: `fake_query`},
	}

	namespaceScheduledTaskMetricExtractionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `tfMetricsExtractionTask`},
		"task_type":      acctest.Representation{RepType: acctest.Optional, Create: `SAVED_SEARCH`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsNamespaceScheduledTaskDataSourceFilterRepresentation}}

	LogAnalyticsNamespaceScheduledTaskMetricExtractionRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"kind":            acctest.Representation{RepType: acctest.Required, Create: `STANDARD`},
		"namespace":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"task_type":       acctest.Representation{RepType: acctest.Required, Create: `SAVED_SEARCH`},
		"action":          acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsNamespaceScheduledTaskActionMetricExtractionRepresentation},
		"schedules":       acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsSchedulesRepresentation},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `tfMetricsExtractionTask`, Update: `tfMetricsExtractionTask2`},
		"saved_search_id": acctest.Representation{RepType: acctest.Required, Create: `ocid1.managementsavedsearch.oc1..aaaaaaaalkfzmfv4467wkqgugxlkgiaiflyqm6edvlf7vjqnxyzsjmgqtvzq`},
	}

	LogAnalyticsNamespaceScheduledTaskActionMetricExtractionRepresentation = map[string]interface{}{
		"type":              acctest.Representation{RepType: acctest.Required, Create: `STREAM`},
		"metric_extraction": acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsMetricsExtractionRepresentation},
		"saved_search_id":   acctest.Representation{RepType: acctest.Required, Create: `ocid1.managementsavedsearch.oc1..aaaaaaaalkfzmfv4467wkqgugxlkgiaiflyqm6edvlf7vjqnxyzsjmgqtvzq`},
	}

	LogAnalyticsMetricsExtractionRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"metric_name":    acctest.Representation{RepType: acctest.Required, Create: `metricName`},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"resource_group": acctest.Representation{RepType: acctest.Required, Create: `resourceGroup`},
	}

	LogAnalyticsSchedulesRepresentation = map[string]interface{}{
		"schedule": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: fixedSchedulesRepresentation}},
	}

	/*
		    TODO: use for acceleration tasks. for purge, only one schedule is allowed.
			schedulesUpdatedRepresentation = map[string]interface{}{
				"schedule": []acctest.RepresentationGroup{{Required, fixedSchedulesRepresentation}, {Required, cronSchedulesRepresentation}},
			}
	*/

	fixedSchedulesRepresentation = map[string]interface{}{
		"type":               acctest.Representation{RepType: acctest.Required, Create: `FIXED_FREQUENCY`, Update: `FIXED_FREQUENCY`},
		"recurring_interval": acctest.Representation{RepType: acctest.Required, Create: `P1D`, Update: `P2D`},
		"repeat_count":       acctest.Representation{RepType: acctest.Required, Create: `4`, Update: `6`},
		"misfire_policy":     acctest.Representation{RepType: acctest.Required, Create: `RETRY_ONCE`, Update: `RETRY_INDEFINITELY`},
	}
	cronSchedulesRepresentation = map[string]interface{}{
		"type":           acctest.Representation{RepType: acctest.Required, Create: `CRON`, Update: `CRON`},
		"expression":     acctest.Representation{RepType: acctest.Required, Create: `0 0 * ? * 2,3,4,5,6`, Update: `0 0 * ? * 2,3`},
		"misfire_policy": acctest.Representation{RepType: acctest.Required, Create: `RETRY_INDEFINITELY`, Update: `RETRY_ONCE`},
		"time_zone":      acctest.Representation{RepType: acctest.Required, Create: `UTC`, Update: `UTC`},
	}

	LogAnalyticsNamespaceScheduledTaskResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

func TestLogAnalyticsNamespaceScheduledTaskResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceScheduledTaskResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task"
	datasourceName := "data.oci_log_analytics_namespace_scheduled_tasks.test_namespace_scheduled_tasks"
	singularDatasourceName := "data.oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task"
	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLogAnalyticsNamespaceScheduledTaskDestroy,
		Steps: []resource.TestStep{
			// verify creation of purge task
			{
				Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceScheduledTaskResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Required, acctest.Create, LogAnalyticsPurgeTaskRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "kind", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tfPurgeTask1"),
					resource.TestCheckResourceAttr(resourceName, "task_type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.data_type", "LOG"),
					resource.TestCheckResourceAttrSet(resourceName, "action.0.purge_compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "action.0.purge_duration", "-P7D"),
					resource.TestCheckResourceAttr(resourceName, "action.0.query_string", "fake_query"),
					resource.TestCheckResourceAttr(resourceName, "action.0.type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_ONCE",
						"recurring_interval": "P1D",
						"repeat_count":       "4",
					}, []string{}),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify update to the compartment (the compartment will be switched back in the next step) of purge task
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LogAnalyticsNamespaceScheduledTaskResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Required, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(LogAnalyticsPurgeTaskRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "kind", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tfPurgeTask1"),
					resource.TestCheckResourceAttr(resourceName, "task_type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.data_type", "LOG"),
					resource.TestCheckResourceAttrSet(resourceName, "action.0.purge_compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "action.0.purge_duration", "-P7D"),
					resource.TestCheckResourceAttr(resourceName, "action.0.query_string", "fake_query"),
					resource.TestCheckResourceAttr(resourceName, "action.0.type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_ONCE",
						"recurring_interval": "P1D",
						"repeat_count":       "4",
					}, []string{}),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LogAnalyticsNamespaceScheduledTaskResourceDependencies,
			},
			//verify creation of scheduled task with metric Extraction
			{
				Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceScheduledTaskResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Required, acctest.Create, LogAnalyticsNamespaceScheduledTaskMetricExtractionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "kind", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tfMetricsExtractionTask"),
					resource.TestCheckResourceAttr(resourceName, "task_type", "SAVED_SEARCH"),
					resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.type", "STREAM"),
					resource.TestCheckResourceAttr(resourceName, "action.0.metric_extraction.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.metric_extraction.0.metric_name", "metricName"),
					resource.TestCheckResourceAttr(resourceName, "action.0.metric_extraction.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "action.0.metric_extraction.0.resource_group", "resourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_ONCE",
						"recurring_interval": "P1D",
						"repeat_count":       "4",
					}, []string{}),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId == "" {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_tasks", "test_namespace_scheduled_tasks", acctest.Optional, acctest.Create, namespaceScheduledTaskMetricExtractionDataSourceRepresentation) +
					compartmentIdVariableStr + LogAnalyticsNamespaceScheduledTaskResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Optional, acctest.Create, LogAnalyticsNamespaceScheduledTaskMetricExtractionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "tfMetricsExtractionTask"),
					resource.TestCheckResourceAttr(datasourceName, "task_type", "SAVED_SEARCH"),
					resource.TestCheckResourceAttr(datasourceName, "scheduled_task_collection.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceScheduledTaskSingularDataSourceRepresentation) +
					compartmentIdVariableStr + LogAnalyticsNamespaceScheduledTaskMetricExtractionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_task_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "num_occurrences"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "tfMetricsExtractionTask2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "task_status"),
					resource.TestCheckResourceAttr(singularDatasourceName, "task_type", "SAVED_SEARCH"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.0.type", "STREAM"),
					resource.TestCheckResourceAttr(resourceName, "action.0.metric_extraction.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.metric_extraction.0.metric_name", "metricName"),
					resource.TestCheckResourceAttr(resourceName, "action.0.metric_extraction.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "action.0.metric_extraction.0.resource_group", "resourceGroup"),
					resource.TestCheckResourceAttr(singularDatasourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.schedule.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_INDEFINITELY",
						"recurring_interval": "P2D",
						"repeat_count":       "6",
					}, []string{}),
				),
			},
			// verify resource import
			{
				Config:            config + LogAnalyticsNamespaceScheduledTaskMetricExtractionResourceConfig,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: namespaceScheduledTaskImportId,
				ImportStateVerifyIgnore: []string{
					"kind",
					"namespace",
					"saved_search_id",
				},
				ResourceName: resourceName,
			},
			// delete before next create
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LogAnalyticsNamespaceScheduledTaskResourceDependencies,
			},
			// verify creation of purge task with optionals
			{
				Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceScheduledTaskResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Optional, acctest.Create, LogAnalyticsPurgeTaskRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "kind", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tfPurgeTask1"),
					resource.TestCheckResourceAttr(resourceName, "task_type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.data_type", "LOG"),
					resource.TestCheckResourceAttrSet(resourceName, "action.0.purge_compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "action.0.purge_duration", "-P7D"),
					resource.TestCheckResourceAttr(resourceName, "action.0.query_string", "fake_query"),
					resource.TestCheckResourceAttr(resourceName, "action.0.type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_ONCE",
						"recurring_interval": "P1D",
						"repeat_count":       "4",
					}, []string{}),

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
			// verify updates to updatable parameters of purge task
			{
				Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceScheduledTaskResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Optional, acctest.Update, LogAnalyticsPurgeTaskRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "kind", "STANDARD"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "tfPurgeTask3"),
					resource.TestCheckResourceAttr(resourceName, "task_type", "PURGE"),
					// TODO: add check for defined_tags value change
					// TODO: add check for freeform tags change
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "action.0.data_type", "LOG"),
					resource.TestCheckResourceAttrSet(resourceName, "action.0.purge_compartment_id"),
					resource.TestCheckResourceAttr(resourceName, "action.0.purge_duration", "-P7D"),
					resource.TestCheckResourceAttr(resourceName, "action.0.query_string", "fake_query"),
					resource.TestCheckResourceAttr(resourceName, "action.0.type", "PURGE"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "schedules.0.schedule.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_INDEFINITELY",
						"recurring_interval": "P2D",
						"repeat_count":       "6",
					}, []string{}),

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
					acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_tasks", "test_namespace_scheduled_tasks", acctest.Optional, acctest.Update, namespaceScheduledTaskDataSourceRepresentation) +
					compartmentIdVariableStr + LogAnalyticsNamespaceScheduledTaskResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Optional, acctest.Update, LogAnalyticsPurgeTaskRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "tfPurgeTask3"),
					resource.TestCheckResourceAttr(datasourceName, "task_type", "PURGE"),

					resource.TestCheckResourceAttr(datasourceName, "scheduled_task_collection.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_scheduled_task", "test_namespace_scheduled_task", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceScheduledTaskSingularDataSourceRepresentation) +
					compartmentIdVariableStr + LogAnalyticsNamespaceScheduledTaskResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_task_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "num_occurrences"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "tfPurgeTask3"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "task_status"),
					resource.TestCheckResourceAttr(singularDatasourceName, "task_type", "PURGE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.0.data_type", "LOG"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "action.0.purge_compartment_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.0.purge_duration", "-P7D"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.0.query_string", "fake_query"),
					resource.TestCheckResourceAttr(singularDatasourceName, "action.0.type", "PURGE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "schedules.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "schedules.0.schedule.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "schedules.0.schedule", map[string]string{
						"type":               "FIXED_FREQUENCY",
						"misfire_policy":     "RETRY_INDEFINITELY",
						"recurring_interval": "P2D",
						"repeat_count":       "6",
					}, []string{}),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				),
			},
			// verify resource import
			{
				Config:            config + LogAnalyticsNamespaceScheduledTaskResourceConfig,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: namespaceScheduledTaskImportId,
				ImportStateVerifyIgnore: []string{
					"kind",
					"namespace",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckLogAnalyticsNamespaceScheduledTaskDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LogAnalyticsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_log_analytics_namespace_scheduled_task" {
			noResourceFound = false
			request := oci_log_analytics.GetScheduledTaskRequest{}

			if namespace, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &namespace
			}

			if scheduledTaskId, ok := rs.Primary.Attributes["scheduled_task_id"]; ok {
				request.ScheduledTaskId = &scheduledTaskId
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")

			response, err := client.GetScheduledTask(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_log_analytics.ScheduledTaskLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("LogAnalyticsNamespaceScheduledTask") {
		resource.AddTestSweepers("LogAnalyticsNamespaceScheduledTask", &resource.Sweeper{
			Name:         "LogAnalyticsNamespaceScheduledTask",
			Dependencies: acctest.DependencyGraph["namespaceScheduledTask"],
			F:            sweepLogAnalyticsNamespaceScheduledTaskResource,
		})
	}
}

func sweepLogAnalyticsNamespaceScheduledTaskResource(compartment string) error {
	logAnalyticsClient := acctest.GetTestClients(&schema.ResourceData{}).LogAnalyticsClient()
	namespaceScheduledTaskIds, err := getLogAnalyticsNamespaceScheduledTaskIds(compartment)
	if err != nil {
		return err
	}
	for _, namespaceScheduledTaskId := range namespaceScheduledTaskIds {
		if ok := acctest.SweeperDefaultResourceId[namespaceScheduledTaskId]; !ok {
			deleteScheduledTaskRequest := oci_log_analytics.DeleteScheduledTaskRequest{}

			deleteScheduledTaskRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "log_analytics")
			_, error := logAnalyticsClient.DeleteScheduledTask(context.Background(), deleteScheduledTaskRequest)
			if error != nil {
				fmt.Printf("Error deleting NamespaceScheduledTask %s %s, It is possible that the resource is already deleted. Please verify manually \n", namespaceScheduledTaskId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &namespaceScheduledTaskId, LogAnalyticsNamespaceScheduledTaskSweepWaitCondition, time.Duration(3*time.Minute),
				LogAnalyticsNamespaceScheduledTaskSweepResponseFetchOperation, "log_analytics", true)
		}
	}
	return nil
}

func getLogAnalyticsNamespaceScheduledTaskIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NamespaceScheduledTaskId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	logAnalyticsClient := acctest.GetTestClients(&schema.ResourceData{}).LogAnalyticsClient()

	listScheduledTasksRequest := oci_log_analytics.ListScheduledTasksRequest{}
	listScheduledTasksRequest.CompartmentId = &compartmentId

	namespaces, error := getNamespaces(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting namespace required for NamespaceScheduledTask resource requests \n")
	}
	for _, namespace := range namespaces {
		listScheduledTasksRequest.NamespaceName = &namespace

		listScheduledTasksResponse, err := logAnalyticsClient.ListScheduledTasks(context.Background(), listScheduledTasksRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NamespaceScheduledTask list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, namespaceScheduledTask := range listScheduledTasksResponse.Items {
			id := *namespaceScheduledTask.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NamespaceScheduledTaskId", id)
		}

	}
	return resourceIds, nil
}

func LogAnalyticsNamespaceScheduledTaskSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if namespaceScheduledTaskResponse, ok := response.Response.(oci_log_analytics.GetScheduledTaskResponse); ok {
		return namespaceScheduledTaskResponse.GetLifecycleState() != oci_log_analytics.ScheduledTaskLifecycleStateDeleted
	}
	return false
}

func LogAnalyticsNamespaceScheduledTaskSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.LogAnalyticsClient().GetScheduledTask(context.Background(), oci_log_analytics.GetScheduledTaskRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}

func namespaceScheduledTaskImportId(state *terraform.State) (string, error) {
	for _, rs := range state.RootModule().Resources {
		if rs.Type == "oci_log_analytics_namespace_scheduled_task" {
			return log_analytics.GetNamespaceScheduledTaskCompositeId(rs.Primary.Attributes["namespace"], rs.Primary.Attributes["scheduled_task_id"]), nil
			//return rs.Primary.Attributes["ID"], nil
		}
	}

	return "", fmt.Errorf("unable to create import id as no resource of type oci_log_analytics_namespace_scheduled_task found in state")
}
