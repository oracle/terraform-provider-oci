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
	oci_sch "github.com/oracle/oci-go-sdk/v65/sch"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	SchServiceConnectorRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorFunctionTargetRepresentation)

	// Dependency definition
	SchServiceConnectorResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, LoggingLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_update_log", acctest.Required, acctest.Update, acctest.GetUpdatedRepresentationCopy("configuration.source.category", acctest.Representation{RepType: acctest.Required, Create: `read`}, LoggingLogRepresentation)) +
		SchLoggingLogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, SchFunctionsFunctionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Required, acctest.Create, StreamingStreamRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)

	SchLoggingLogResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, LoggingLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_update_log_group", acctest.Required, acctest.Create, logGroupUpdateRepresentation)

	// source definitions
	SchServiceConnectorSourceLogSourcesRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"log_group_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_logging_log_group.test_log_group.id}`, Update: `${oci_logging_log_group.test_update_log_group.id}`},
		"log_id":         acctest.Representation{RepType: acctest.Optional, Create: `${oci_logging_log.test_log.id}`, Update: `${oci_logging_log.test_update_log.id}`},
	}

	SchServiceConnectorSourceRepresentation = map[string]interface{}{
		"kind":        acctest.Representation{RepType: acctest.Required, Create: `logging`},
		"log_sources": acctest.RepresentationGroup{RepType: acctest.Required, Group: SchServiceConnectorSourceLogSourcesRepresentation},
	}

	SchSchServiceConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `My_Service_Connector`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorDataSourceFilterRepresentation},
	}

	// task definitions
	SchServiceConnectorTasksRepresentation = map[string]interface{}{
		"condition": acctest.Representation{RepType: acctest.Required, Create: `data.action='REJECT'`, Update: `logContent='20'`},
		"kind":      acctest.Representation{RepType: acctest.Required, Create: `logRule`},
	}

	// target definitions
	functionTargetRepresentation = map[string]interface{}{
		"kind":        acctest.Representation{RepType: acctest.Required, Create: `functions`},
		"function_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_function.test_function.id}`},
	}

	// target definitions with batching details
	functionTargetBatchRepresentation = map[string]interface{}{
		"kind":              acctest.Representation{RepType: acctest.Required, Create: `functions`},
		"function_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_function.test_function.id}`},
		"batch_size_in_kbs": acctest.Representation{RepType: acctest.Optional, Create: `5000`},
		"batch_size_in_num": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"batch_time_in_sec": acctest.Representation{RepType: acctest.Optional, Create: `5`},
	}

	objectStorageTargetRepresentation = map[string]interface{}{
		"kind":                       acctest.Representation{RepType: acctest.Required, Create: `objectStorage`},
		"bucket":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object_name_prefix":         acctest.Representation{RepType: acctest.Optional, Create: `test_prefix`},
		"batch_rollover_size_in_mbs": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"batch_rollover_time_in_ms":  acctest.Representation{RepType: acctest.Optional, Create: `80000`},
	}

	logAnTargetRepresentation = map[string]interface{}{
		"kind":         acctest.Representation{RepType: acctest.Required, Create: `loggingAnalytics`},
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.logAn_log_group_ocid}`},
	}

	onsTargetRepresentation = map[string]interface{}{
		"kind":                       acctest.Representation{RepType: acctest.Required, Create: `notifications`},
		"topic_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
		"enable_formatted_messaging": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	monitoringTargetRepresentation = map[string]interface{}{
		"kind":             acctest.Representation{RepType: acctest.Required, Create: `monitoring`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"metric":           acctest.Representation{RepType: acctest.Required, Create: `metric`, Update: `metric1`},
		"metric_namespace": acctest.Representation{RepType: acctest.Required, Create: `metricnamespace`, Update: `metricnamespace_1`},
		"dimensions":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: serviceConnectorTargetStaticDimensionsRepresentation_0},
	}

	//Making dimensions required though it is optional in the API spec
	monitoringTargetStaticDimensionRepresentation = map[string]interface{}{
		"kind":             acctest.Representation{RepType: acctest.Required, Create: `monitoring`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"metric":           acctest.Representation{RepType: acctest.Required, Create: `static_metric_0`, Update: `metric_update_1`},
		"metric_namespace": acctest.Representation{RepType: acctest.Required, Create: `static_metricnamespace_0`, Update: `static_metricnamespace_1`},
		"dimensions":       acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorTargetStaticDimensionsRepresentation_0},
	}

	monitoringTargetJmesPathDimensionRepresentation = map[string]interface{}{
		"kind":             acctest.Representation{RepType: acctest.Required, Create: `monitoring`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"metric":           acctest.Representation{RepType: acctest.Required, Create: `jmespath_metric_0`, Update: `metric_update_1`},
		"metric_namespace": acctest.Representation{RepType: acctest.Required, Create: `jmespath_metricnamespace_0`, Update: `jmespath_metricnamespace_1`},
		"dimensions":       acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorTargetJmesPathDimensionsRepresentation_0},
	}

	monitoringTargetStaticAndJmesPathRepresentation = map[string]interface{}{
		"kind":             acctest.Representation{RepType: acctest.Required, Create: `monitoring`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"metric":           acctest.Representation{RepType: acctest.Required, Create: `metric`, Update: `metric_1`},
		"metric_namespace": acctest.Representation{RepType: acctest.Required, Create: `metricnamespace`, Update: `metricnamespace_1`},
		"dimensions": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: serviceConnectorTargetJmesPathDimensionsRepresentation_0},
			{RepType: acctest.Required, Group: serviceConnectorTargetJmesPathDimensionsRepresentation_1},
			{RepType: acctest.Required, Group: serviceConnectorTargetStaticDimensionsRepresentation_0},
			{RepType: acctest.Required, Group: serviceConnectorTargetStaticDimensionsRepresentation_1},
		},
	}

	// Create serviceConnector definitions
	serviceConnectorRepresentationNoTarget = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `My_Service_Connector`, Update: `displayName2`},
		"source":         acctest.RepresentationGroup{RepType: acctest.Required, Group: SchServiceConnectorSourceRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `My service connector description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"tasks":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: SchServiceConnectorTasksRepresentation},
	}

	serviceConnectorRepresentationMonitoringSource = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `My_Service_Connector`, Update: `displayName2`},
		"source":         acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorMonitoringSourceRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `My service connector description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"tasks":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: SchServiceConnectorTasksRepresentation},
	}

	serviceConnectorMonitoringSourceRepresentation = map[string]interface{}{
		"kind":               acctest.Representation{RepType: acctest.Required, Create: `monitoring`},
		"monitoring_sources": acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorSourceMonitoringSourcesRepresentation},
	}

	serviceConnectorSourceMonitoringSourcesRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorSourceMonitoringSourcesNamespaceDetailsRepresentation},
	}

	serviceConnectorSourceMonitoringSourcesNamespaceDetailsRepresentation = map[string]interface{}{
		"kind":       acctest.Representation{RepType: acctest.Required, Create: `selected`, Update: `selected`},
		"namespaces": acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespacesRepresentation},
	}

	serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespacesRepresentation = map[string]interface{}{
		"metrics":   acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespacesMetricsRepresentation},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `oci_computeagent`},
	}

	serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespacesMetricsRepresentation = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Required, Create: `all`, Update: `all`},
	}

	serviceConnectorRepresentationMonitoringSourceMultipleNamespaces = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `My_Service_Connector`, Update: `displayName2`},
		"source":         acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorMonitoringSourceNamespaceRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `My service connector description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"tasks":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: SchServiceConnectorTasksRepresentation},
	}

	serviceConnectorMonitoringSourceNamespaceRepresentation = map[string]interface{}{
		"kind":               acctest.Representation{RepType: acctest.Required, Create: `monitoring`},
		"monitoring_sources": acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorSourceMonitoringSourcesNamespaceRepresentation},
	}

	serviceConnectorSourceMonitoringSourcesNamespaceRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"namespace_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespaceRepresentation},
	}

	serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespaceRepresentation = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Required, Create: `selected`, Update: `selected`},
		"namespaces": []acctest.RepresentationGroup{
			{RepType: acctest.Required, Group: serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespacesRepresentation_0},
			{RepType: acctest.Required, Group: serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespacesRepresentation_1},
		},
	}

	serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespacesRepresentation_0 = map[string]interface{}{
		"metrics":   acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespacesMetricsRepresentation},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `oci_computeagent`},
	}

	serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespacesRepresentation_1 = map[string]interface{}{
		"metrics":   acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorSourceMonitoringSourcesNamespaceDetailsNamespacesMetricsRepresentation},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `oci_logging_analytics`},
	}

	// targets for logging as a source
	serviceConnectorFunctionTargetRepresentation      = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, functionTargetRepresentation)
	serviceConnectorObjectStorageTargetRepresentation = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, objectStorageTargetRepresentation)
	serviceConnectorLogAnTargetRepresentation         = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, logAnTargetRepresentation)
	serviceConnectorOnsTargetRepresentation           = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, onsTargetRepresentation)

	serviceConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"service_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_sch_service_connector.test_service_connector.id}`},
	}

	serviceConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_sch_service_connector.test_service_connector.id}`}},
	}

	// Update serviceConnector definitions
	SchServiceConnectorTargetRepresentation = map[string]interface{}{
		"kind":      acctest.Representation{RepType: acctest.Required, Create: `streaming`},
		"stream_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_streaming_stream.test_stream.id}`},
	}

	SchFunctionsFunctionRepresentation = map[string]interface{}{
		"application_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_application.test_application.id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `ExampleFunction`},
		"memory_in_mbs":  acctest.Representation{RepType: acctest.Required, Create: `128`, Update: `256`},
		"config":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"MY_FUNCTION_CONFIG": "ConfVal"}},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"image":          acctest.Representation{RepType: acctest.Required, Create: `${var.image}`, Update: `${var.image_for_update}`},
		"image_digest":   acctest.Representation{RepType: acctest.Optional, Create: `${var.image_digest}`, Update: `${var.image_digest_for_update}`},
	}

	serviceConnectorTargetStaticDimensionsRepresentation_0 = map[string]interface{}{
		"dimension_value": acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorTargetDimensionsStaticDimensionValueRepresentation_0},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `static_dimension_0`, Update: `static_dimension_update_1`},
	}

	serviceConnectorTargetStaticDimensionsRepresentation_1 = map[string]interface{}{
		"dimension_value": acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorTargetDimensionsStaticDimensionValueRepresentation_1},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `static_dimension_1`, Update: `static_dimension_update_2`},
	}

	serviceConnectorTargetJmesPathDimensionsRepresentation_0 = map[string]interface{}{
		"dimension_value": acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorTargetDimensionsJmesPathDimensionValueRepresentation_0},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `jmespath_dimension_0`, Update: `jmespath_dimension_update_1`},
	}

	serviceConnectorTargetJmesPathDimensionsRepresentation_1 = map[string]interface{}{
		"dimension_value": acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorTargetDimensionsJmesPathDimensionValueRepresentation_1},
		"name":            acctest.Representation{RepType: acctest.Required, Create: `jmespath_dimension_1`, Update: `jmespath_dimension_update_2`},
	}

	serviceConnectorTargetDimensionsStaticDimensionValueRepresentation_0 = map[string]interface{}{
		"kind":  acctest.Representation{RepType: acctest.Required, Create: `static`, Update: `static`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `static_value_0`, Update: `static_value_update_1`},
	}

	serviceConnectorTargetDimensionsStaticDimensionValueRepresentation_1 = map[string]interface{}{
		"kind":  acctest.Representation{RepType: acctest.Required, Create: `static`, Update: `static`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `static_value_1`, Update: `static_value_update_2`},
	}

	serviceConnectorTargetDimensionsJmesPathDimensionValueRepresentation_0 = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Required, Create: `jmesPath`, Update: `jmesPath`},
		"path": acctest.Representation{RepType: acctest.Required, Create: `logContent.data.compartmentId`, Update: `logContent.data.datacenterid`},
	}
	serviceConnectorTargetDimensionsJmesPathDimensionValueRepresentation_1 = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Required, Create: `jmesPath`, Update: `jmesPath`},
		"path": acctest.Representation{RepType: acctest.Required, Create: `logContent.data.namespace`, Update: `logContent.data.compartmentId`},
	}
	ServiceConnectorResourceConfig = SchServiceConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Update, serviceConnectorFunctionTargetRepresentation)
)

// issue-routing-tag: sch/default
func TestSchServiceConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestSchServiceConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	image := utils.GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	logAnLogGroupId := utils.GetEnvSettingWithBlankDefault("logAn_log_group_ocid")
	logAnLogGroupIdVariableStr := fmt.Sprintf("variable \"logAn_log_group_ocid\" { default = \"%s\" }\n", logAnLogGroupId)

	resourceName := "oci_sch_service_connector.test_service_connector"
	datasourceName := "data.oci_sch_service_connectors.test_service_connectors"
	singularDatasourceName := "data.oci_sch_service_connector.test_service_connector"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SchServiceConnectorResourceDependencies+imageVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Create, serviceConnectorObjectStorageTargetRepresentation), "sch", "serviceConnector", t)

	acctest.ResourceTest(t, testAccCheckSchServiceConnectorDestroy, []resource.TestStep{
		// verify Create with functions
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorFunctionTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "functions"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.function_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr,
		},

		// verify Create with objectstorage
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Create, serviceConnectorObjectStorageTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "objectStorage"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.bucket"),
				resource.TestCheckResourceAttr(resourceName, "target.0.batch_rollover_size_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "target.0.batch_rollover_time_in_ms", "80000"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr,
		},

		// verify Create with log analytics
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr + logAnLogGroupIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorLogAnTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "loggingAnalytics"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.log_group_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr,
		},

		// verify Create with ons
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Create, serviceConnectorOnsTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "My service connector description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "notifications"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.topic_id"),
				resource.TestCheckResourceAttr(resourceName, "target.0.enable_formatted_messaging", "true"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Create, serviceConnectorFunctionTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "My service connector description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "functions"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.function_id"),
				resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.condition", "data.action='REJECT'"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "logRule"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(serviceConnectorFunctionTargetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "My service connector description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "functions"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.function_id"),
				resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.condition", "data.action='REJECT'"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "logRule"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"target"}), map[string]interface{}{
						"target": acctest.RepresentationGroup{RepType: acctest.Required, Group: SchServiceConnectorTargetRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "streaming"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.stream_id"),
				resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.condition", "logContent='20'"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "logRule"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify stop service connector
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"target"}), map[string]interface{}{
						"target": acctest.RepresentationGroup{RepType: acctest.Required, Group: SchServiceConnectorTargetRepresentation},
						"state":  acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "streaming"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.stream_id"),
				resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.condition", "logContent='20'"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "logRule"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify start service connector
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"target"}), map[string]interface{}{
						"target": acctest.RepresentationGroup{RepType: acctest.Required, Group: SchServiceConnectorTargetRepresentation},
						"state":  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.log_sources.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "streaming"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.stream_id"),
				resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.condition", "logContent='20'"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "logRule"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_sch_service_connectors", "test_service_connectors", acctest.Optional, acctest.Update, SchSchServiceConnectorDataSourceRepresentation) +
				compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"target"}), map[string]interface{}{
						"target": acctest.RepresentationGroup{RepType: acctest.Required, Group: SchServiceConnectorTargetRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "service_connector_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "service_connector_collection.0.items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"target"}), map[string]interface{}{
						"target": acctest.RepresentationGroup{RepType: acctest.Required, Group: SchServiceConnectorTargetRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_connector_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.kind", "streaming"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.condition", "logContent='20'"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.kind", "logRule"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + SchServiceConnectorRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckSchServiceConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ServiceConnectorClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_sch_service_connector" {
			noResourceFound = false
			request := oci_sch.GetServiceConnectorRequest{}

			tmp := rs.Primary.ID
			request.ServiceConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "sch")

			response, err := client.GetServiceConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_sch.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("SchServiceConnector") {
		resource.AddTestSweepers("SchServiceConnector", &resource.Sweeper{
			Name:         "SchServiceConnector",
			Dependencies: acctest.DependencyGraph["serviceConnector"],
			F:            sweepSchServiceConnectorResource,
		})
	}
}

func sweepSchServiceConnectorResource(compartment string) error {
	serviceConnectorClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceConnectorClient()
	serviceConnectorIds, err := getSchServiceConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, serviceConnectorId := range serviceConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[serviceConnectorId]; !ok {
			deleteServiceConnectorRequest := oci_sch.DeleteServiceConnectorRequest{}

			deleteServiceConnectorRequest.ServiceConnectorId = &serviceConnectorId

			deleteServiceConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "sch")
			_, error := serviceConnectorClient.DeleteServiceConnector(context.Background(), deleteServiceConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting ServiceConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", serviceConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &serviceConnectorId, SchServiceConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				SchServiceConnectorSweepResponseFetchOperation, "sch", true)
		}
	}
	return nil
}

func getSchServiceConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ServiceConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	serviceConnectorClient := acctest.GetTestClients(&schema.ResourceData{}).ServiceConnectorClient()

	listServiceConnectorsRequest := oci_sch.ListServiceConnectorsRequest{}
	listServiceConnectorsRequest.CompartmentId = &compartmentId
	listServiceConnectorsRequest.LifecycleState = oci_sch.ListServiceConnectorsLifecycleStateActive
	listServiceConnectorsResponse, err := serviceConnectorClient.ListServiceConnectors(context.Background(), listServiceConnectorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ServiceConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, serviceConnector := range listServiceConnectorsResponse.Items {
		id := *serviceConnector.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ServiceConnectorId", id)
	}
	return resourceIds, nil
}

func SchServiceConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if serviceConnectorResponse, ok := response.Response.(oci_sch.GetServiceConnectorResponse); ok {
		return serviceConnectorResponse.LifecycleState != oci_sch.LifecycleStateDeleted
	}
	return false
}

func SchServiceConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ServiceConnectorClient().GetServiceConnector(context.Background(), oci_sch.GetServiceConnectorRequest{
		ServiceConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func createServiceConnectorRepresentation(sc map[string]interface{}, target map[string]interface{}) map[string]interface{} {
	serviceConnector := make(map[string]interface{})

	// Copy map and populate target
	for key, value := range sc {
		serviceConnector[key] = value
	}
	serviceConnector["target"] = acctest.RepresentationGroup{RepType: acctest.Required, Group: target}

	return serviceConnector
}
