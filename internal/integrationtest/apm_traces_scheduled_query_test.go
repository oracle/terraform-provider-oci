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
	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApmTracesScheduledQueryRequiredOnlyResource = ApmTracesScheduledQueryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_traces_scheduled_query", "test_scheduled_query", acctest.Required, acctest.Create, ApmTracesScheduledQueryRepresentation)

	ApmTracesScheduledQueryResourceConfig = ApmTracesScheduledQueryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apm_traces_scheduled_query", "test_scheduled_query", acctest.Optional, acctest.Update, ApmTracesScheduledQueryRepresentation)

	ApmTracesScheduledQuerySingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.apm_domain_id}`},
		"scheduled_query_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apm_traces_scheduled_query.test_scheduled_query.id}`},
	}

	ApmTracesScheduledQueryDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.apm_domain_id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"filter":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmTracesScheduledQueryDataSourceFilterRepresentation}}
	ApmTracesScheduledQueryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apm_traces_scheduled_query.test_scheduled_query.id}`}},
	}

	ApmTracesScheduledQueryRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.apm_domain_id}`},
		"defined_tags": acctest.Representation{RepType: acctest.Optional,
			Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`,
			Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsDifferencesRepresentation},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"opc_dry_run":                 acctest.Representation{RepType: acctest.Optional, Create: `opcDryRun`, Update: `opcDryRun2`},
		"scheduled_query_description": acctest.Representation{RepType: acctest.Optional, Create: `scheduledQueryDescription`, Update: `scheduledQueryDescription2`},
		"scheduled_query_maximum_runtime_in_seconds": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"scheduled_query_name":                       acctest.Representation{RepType: acctest.Required, Create: `TestScheduledQuery1`, Update: `TestScheduledQuery2`},
		"scheduled_query_processing_configuration":   acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmTracesScheduledQueryScheduledQueryProcessingConfigurationRepresentation},
		"scheduled_query_processing_sub_type":        acctest.Representation{RepType: acctest.Required, Create: `CUSTOM_METRIC`},
		"scheduled_query_processing_type":            acctest.Representation{RepType: acctest.Required, Create: `EXPORT`},
		"scheduled_query_retention_criteria":         acctest.Representation{RepType: acctest.Required, Create: `UPDATE`},
		"scheduled_query_retention_period_in_ms":     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `30`},
		"scheduled_query_schedule":                   acctest.Representation{RepType: acctest.Required, Create: `SCHEDULE STARTING AFTER 2025-06-20T21:20:00Z EVERY 5 MINUTES`},
		"scheduled_query_text":                       acctest.Representation{RepType: acctest.Required, Create: `SHOW SPANS time_bucket_start(1,apmdbInsertTime2) AS metricStartTime,1 MINUTE AS metricDuration,serviceName,operationName,count(*) AS metricValue WHERE apmdbinserttime2>=TimeTruncate(now(),'minute') - 5 MINUTES GROUP BY time_bucket_start(1,apmdbInsertTime2),serviceName,operationName FIRST 10000 ROWS BETWEEN now() - 2 HOURS AND now()`},
	}
	ApmTracesScheduledQueryScheduledQueryProcessingConfigurationRepresentation = map[string]interface{}{
		"custom_metric": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApmTracesScheduledQueryScheduledQueryProcessingConfigurationCustomMetricRepresentation},
		/*"object_storage": acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmTracesScheduledQueryScheduledQueryProcessingConfigurationObjectStorageRepresentation},
		"streaming":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: ApmTracesScheduledQueryScheduledQueryProcessingConfigurationStreamingRepresentation},*/
	}
	ApmTracesScheduledQueryScheduledQueryProcessingConfigurationCustomMetricRepresentation = map[string]interface{}{
		"name":                         acctest.Representation{RepType: acctest.Required, Create: `name`},
		"compartment":                  acctest.Representation{RepType: acctest.Optional, Create: `compartment`},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"is_anomaly_detection_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_metric_published":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"namespace":                    acctest.Representation{RepType: acctest.Optional, Create: `namespace`},
		"resource_group":               acctest.Representation{RepType: acctest.Optional, Create: `resourceGroup`},
		"unit":                         acctest.Representation{RepType: acctest.Optional, Create: `unit`},
	}
	ApmTracesScheduledQueryScheduledQueryProcessingConfigurationObjectStorageRepresentation = map[string]interface{}{
		"bucket":             acctest.Representation{RepType: acctest.Optional, Create: `bucket`, Update: `bucket2`},
		"name_space":         acctest.Representation{RepType: acctest.Optional, Create: `nameSpace`, Update: `nameSpace2`},
		"object_name_prefix": acctest.Representation{RepType: acctest.Optional, Create: `objectNamePrefix`, Update: `objectNamePrefix2`},
	}
	ApmTracesScheduledQueryScheduledQueryProcessingConfigurationStreamingRepresentation = map[string]interface{}{
		"stream_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_streaming_stream.test_stream.id}`},
	}

	ApmTracesScheduledQueryResourceDependencies = DefinedTagsDependencies /*+
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, NamespaceSingularDataSourceRepresentation)
	acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Required, acctest.Create, StreamingStreamRepresentation)*/
)

// issue-routing-tag: apm_traces/default
func TestApmTracesScheduledQueryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmTracesScheduledQueryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	apmDomainId := utils.GetEnvSettingWithBlankDefault("apm_domain_id")
	if apmDomainId == "" {
		t.Skip("Set apm_domain_id to run this test")
	}
	apmDomainIdVariableStr := fmt.Sprintf("variable \"apm_domain_id\" { default = \"%s\" }\n", apmDomainId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_apm_traces_scheduled_query.test_scheduled_query"
	datasourceName := "data.oci_apm_traces_scheduled_queries.test_scheduled_queries"
	singularDatasourceName := "data.oci_apm_traces_scheduled_query.test_scheduled_query"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApmTracesScheduledQueryResourceDependencies+apmDomainIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_apm_traces_scheduled_query", "test_scheduled_query", acctest.Optional, acctest.Create, ApmTracesScheduledQueryRepresentation), "apmtraces", "scheduledQuery", t)

	acctest.ResourceTest(t, testAccCheckApmTracesScheduledQueryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ApmTracesScheduledQueryResourceDependencies + apmDomainIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_traces_scheduled_query", "test_scheduled_query", acctest.Required, acctest.Create, ApmTracesScheduledQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ApmTracesScheduledQueryResourceDependencies + apmDomainIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApmTracesScheduledQueryResourceDependencies + apmDomainIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_traces_scheduled_query", "test_scheduled_query", acctest.Optional, acctest.Create, ApmTracesScheduledQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "opc_dry_run", "opcDryRun"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_description", "scheduledQueryDescription"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_maximum_runtime_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_name", "TestScheduledQuery1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.compartment", "compartment"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.is_anomaly_detection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.is_metric_published", "false"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.resource_group", "resourceGroup"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.unit", "unit"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_sub_type", "CUSTOM_METRIC"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_type", "EXPORT"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_retention_criteria", "UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_retention_period_in_ms", "10"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_schedule", "SCHEDULE STARTING AFTER 2025-06-20T21:20:00Z EVERY 5 MINUTES"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_text", "SHOW SPANS time_bucket_start(1,apmdbInsertTime2) AS metricStartTime,1 MINUTE AS metricDuration,serviceName,operationName,count(*) AS metricValue WHERE apmdbinserttime2>=TimeTruncate(now(),'minute') - 5 MINUTES GROUP BY time_bucket_start(1,apmdbInsertTime2),serviceName,operationName FIRST 10000 ROWS BETWEEN now() - 2 HOURS AND now()"),

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

		// verify updates to updatable parameters (ScheduledQueryName, ScheduledQueryDescription , ScheduledQueryMaxRuntimeInSeconds)
		{
			Config: config + compartmentIdVariableStr + ApmTracesScheduledQueryResourceDependencies + apmDomainIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_traces_scheduled_query", "test_scheduled_query", acctest.Optional, acctest.Update, ApmTracesScheduledQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "opc_dry_run", "opcDryRun2"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_description", "scheduledQueryDescription2"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_maximum_runtime_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_name", "TestScheduledQuery2"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.compartment", "compartment"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.is_anomaly_detection_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.is_metric_published", "false"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.resource_group", "resourceGroup"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_configuration.0.custom_metric.0.unit", "unit"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_sub_type", "CUSTOM_METRIC"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_processing_type", "EXPORT"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_retention_criteria", "UPDATE"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_retention_period_in_ms", "30"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_schedule", "SCHEDULE STARTING AFTER 2025-06-20T21:20:00Z EVERY 5 MINUTES"),
				resource.TestCheckResourceAttr(resourceName, "scheduled_query_text", "SHOW SPANS time_bucket_start(1,apmdbInsertTime2) AS metricStartTime,1 MINUTE AS metricDuration,serviceName,operationName,count(*) AS metricValue WHERE apmdbinserttime2>=TimeTruncate(now(),'minute') - 5 MINUTES GROUP BY time_bucket_start(1,apmdbInsertTime2),serviceName,operationName FIRST 10000 ROWS BETWEEN now() - 2 HOURS AND now()"),

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
			Config: config + apmDomainIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_traces_scheduled_queries", "test_scheduled_queries", acctest.Optional, acctest.Update, ApmTracesScheduledQueryDataSourceRepresentation) +
				compartmentIdVariableStr + ApmTracesScheduledQueryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apm_traces_scheduled_query", "test_scheduled_query", acctest.Optional, acctest.Update, ApmTracesScheduledQueryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),

				resource.TestCheckResourceAttr(datasourceName, "scheduled_query_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "scheduled_query_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config + apmDomainIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_traces_scheduled_query", "test_scheduled_query", acctest.Required, acctest.Create, ApmTracesScheduledQuerySingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApmTracesScheduledQueryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_query_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_description", "scheduledQueryDescription2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_query_instances"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_maximum_runtime_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_name", "TestScheduledQuery2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "scheduled_query_next_run_in_ms"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.custom_metric.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.custom_metric.0.compartment", "compartment"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.custom_metric.0.description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.custom_metric.0.is_anomaly_detection_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.custom_metric.0.is_metric_published", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.custom_metric.0.name", "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.custom_metric.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.custom_metric.0.resource_group", "resourceGroup"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.custom_metric.0.unit", "unit"),
				/*resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.object_storage.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.object_storage.0.bucket", "bucket2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.object_storage.0.name_space", "nameSpace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.object_storage.0.object_name_prefix", "objectNamePrefix2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_configuration.0.streaming.#", "1"),*/
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_sub_type", "CUSTOM_METRIC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_processing_type", "EXPORT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_retention_criteria", "UPDATE"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_retention_period_in_ms", "30"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_schedule", "SCHEDULE STARTING AFTER 2025-06-20T21:20:00Z EVERY 5 MINUTES"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scheduled_query_text", "SHOW SPANS time_bucket_start(1,apmdbInsertTime2) AS metricStartTime,1 MINUTE AS metricDuration,serviceName,operationName,count(*) AS metricValue WHERE apmdbinserttime2>=TimeTruncate(now(),'minute') - 5 MINUTES GROUP BY time_bucket_start(1,apmdbInsertTime2),serviceName,operationName FIRST 10000 ROWS BETWEEN now() - 2 HOURS AND now()"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
			),
		},
		// verify resource import
		{
			Config:            config + ApmTracesScheduledQueryRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"apm_domain_id",
				"opc_dry_run",
				"scheduled_query_retention_period_in_ms", //ignore as it does not apply to this case
				"scheduled_query_next_run_in_ms",         //ignore as it after each run, it calculates the value for next run
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckApmTracesScheduledQueryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ScheduledQueryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apm_traces_scheduled_query" {
			noResourceFound = false
			request := oci_apm_traces.GetScheduledQueryRequest{}

			if value, ok := rs.Primary.Attributes["apm_domain_id"]; ok {
				request.ApmDomainId = &value
			}

			tmp := rs.Primary.ID
			request.ScheduledQueryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_traces")

			response, err := client.GetScheduledQuery(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_apm_traces.LifecycleStatesDeleted): true,
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
	if !acctest.InSweeperExcludeList("ApmTracesScheduledQuery") {
		resource.AddTestSweepers("ApmTracesScheduledQuery", &resource.Sweeper{
			Name:         "ApmTracesScheduledQuery",
			Dependencies: acctest.DependencyGraph["scheduledQuery"],
			F:            sweepApmTracesScheduledQueryResource,
		})
	}
}

func sweepApmTracesScheduledQueryResource(compartment string) error {
	scheduledQueryClient := acctest.GetTestClients(&schema.ResourceData{}).ScheduledQueryClient()
	scheduledQueryIds, err := getApmTracesScheduledQueryIds(compartment)
	if err != nil {
		return err
	}
	for _, scheduledQueryId := range scheduledQueryIds {
		if ok := acctest.SweeperDefaultResourceId[scheduledQueryId]; !ok {
			deleteScheduledQueryRequest := oci_apm_traces.DeleteScheduledQueryRequest{}

			deleteScheduledQueryRequest.ScheduledQueryId = &scheduledQueryId

			deleteScheduledQueryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apm_traces")
			_, error := scheduledQueryClient.DeleteScheduledQuery(context.Background(), deleteScheduledQueryRequest)
			if error != nil {
				fmt.Printf("Error deleting ScheduledQuery %s %s, It is possible that the resource is already deleted. Please verify manually \n", scheduledQueryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &scheduledQueryId, ApmTracesScheduledQuerySweepWaitCondition, time.Duration(3*time.Minute),
				ApmTracesScheduledQuerySweepResponseFetchOperation, "apm_traces", true)
		}
	}
	return nil
}

func getApmTracesScheduledQueryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ScheduledQueryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	scheduledQueryClient := acctest.GetTestClients(&schema.ResourceData{}).ScheduledQueryClient()

	listScheduledQueriesRequest := oci_apm_traces.ListScheduledQueriesRequest{}
	//listScheduledQueriesRequest.CompartmentId = &compartmentId

	apmDomainIds, error := getApmDomainIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting apmDomainId required for ScheduledQuery resource requests \n")
	}
	for _, apmDomainId := range apmDomainIds {
		listScheduledQueriesRequest.ApmDomainId = &apmDomainId

		listScheduledQueriesResponse, err := scheduledQueryClient.ListScheduledQueries(context.Background(), listScheduledQueriesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting ScheduledQuery list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, scheduledQuery := range listScheduledQueriesResponse.Items {
			id := *scheduledQuery.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ScheduledQueryId", id)
		}

	}
	return resourceIds, nil
}

func ApmTracesScheduledQuerySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if scheduledQueryResponse, ok := response.Response.(oci_apm_traces.GetScheduledQueryResponse); ok {
		return scheduledQueryResponse.LifecycleState != oci_apm_traces.LifecycleStatesDeleted
	}
	return false
}

func ApmTracesScheduledQuerySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ScheduledQueryClient().GetScheduledQuery(context.Background(), oci_apm_traces.GetScheduledQueryRequest{
		ScheduledQueryId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
