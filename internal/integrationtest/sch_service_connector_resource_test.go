// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	SchServiceConnectorResourceDependenciesForPrivateStream = acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, LoggingLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_update_log", acctest.Required, acctest.Update, acctest.GetUpdatedRepresentationCopy("configuration.source.category", acctest.Representation{RepType: acctest.Required, Create: `read`}, LoggingLogRepresentation)) +
		SchLoggingLogResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Required, Create: `true`}, "prohibit_internet_ingress": acctest.Representation{RepType: acctest.Required, Create: `true`}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_application", "test_application", acctest.Required, acctest.Create, FunctionsApplicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_functions_function", "test_function", acctest.Required, acctest.Create, SchFunctionsFunctionRepresentation) +
		KeyResourceDependencyConfig + kmsKeyIdCreateVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", acctest.Optional, acctest.Create, streampoolidRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", acctest.Optional, acctest.Create, StreamingStreamPoolRepresentation)
	// streaming as a source definition
	serviceConnectorStreamingSourceCursorRepresentation = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Optional, Create: `LATEST`, Update: `TRIM_HORIZON`},
	}

	serviceConnectorStreamingSourceRepresentation = map[string]interface{}{
		"kind":      acctest.Representation{RepType: acctest.Required, Create: `streaming`},
		"cursor":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: serviceConnectorStreamingSourceCursorRepresentation},
		"stream_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_streaming_stream.test_stream.id}`},
	}

	// function as a task
	serviceConnectorFunctionTasksRepresentation = map[string]interface{}{
		"kind":              acctest.Representation{RepType: acctest.Required, Create: `function`},
		"batch_size_in_kbs": acctest.Representation{RepType: acctest.Required, Create: `60`},
		"batch_time_in_sec": acctest.Representation{RepType: acctest.Required, Create: `60`},
		"function_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_functions_function.test_function.id}`},
	}

	// Create serviceConnector definitions
	serviceConnectorRepresentationNoTargetStreamingSource = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `My_Service_Connector`, Update: `displayName2`},
		"source":         acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorStreamingSourceRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `My service connector description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"tasks":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: SchServiceConnectorTasksRepresentation},
	}

	serviceConnectorRepresentationNoTargetNoTasksStreamingSource = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `My_Service_Connector`, Update: `displayName2`},
		"source":         acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorStreamingSourceRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `My service connector description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	serviceConnectorRepresentationNoTargetStreamingSourceFunctionTask = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `My_Service_Connector`, Update: `displayName2`},
		"source":         acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorStreamingSourceRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `My service connector description`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"tasks":          acctest.RepresentationGroup{RepType: acctest.Required, Group: serviceConnectorFunctionTasksRepresentation},
	}

	logAnalyticsTargetRepresentation = map[string]interface{}{
		"kind":                  acctest.Representation{RepType: acctest.Required, Create: `loggingAnalytics`},
		"log_group_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.logAn_log_group_ocid}`},
		"log_source_identifier": acctest.Representation{RepType: acctest.Required, Create: `${var.logAn_log_source_name}`, Update: `LinuxSyslogSource`},
	}

	// targets for streaming as a source
	serviceConnectorFunctionTargetStreamingSourceRepresentation             = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTargetStreamingSource, functionTargetRepresentation)
	serviceConnectorFunctionTargetStreamingSourceFunctionTaskRepresentation = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTargetStreamingSourceFunctionTask, functionTargetRepresentation)
	serviceConnectorLogAnTargetStreamingSourceRepresentation                = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTargetNoTasksStreamingSource, logAnalyticsTargetRepresentation)

	updatedServiceConnectorFunctionTasksRepresentation = map[string]interface{}{
		"kind":              acctest.Representation{RepType: acctest.Optional, Update: `function`},
		"batch_size_in_kbs": acctest.Representation{RepType: acctest.Optional, Update: `60`},
		"batch_time_in_sec": acctest.Representation{RepType: acctest.Optional, Update: `60`},
		"function_id":       acctest.Representation{RepType: acctest.Optional, Update: `${oci_functions_function.test_function.id}`},
	}

	updatedServiceConnectorStreamingSourceRepresentation = map[string]interface{}{
		"kind":      acctest.Representation{RepType: acctest.Optional, Update: `streaming`},
		"cursor":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: serviceConnectorStreamingSourceCursorRepresentation},
		"stream_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_streaming_stream.test_stream.id}`},
	}

	Schserviceconnectorsourcelogsourcesrepresentation2 = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.tenancy_ocid}`},
		"log_group_id":   acctest.Representation{RepType: acctest.Optional, Create: `_Audit`},
	}

	Schserviceconnectorsourcerepresentation2 = map[string]interface{}{
		"kind":        acctest.Representation{RepType: acctest.Required, Create: `logging`},
		"log_sources": acctest.RepresentationGroup{RepType: acctest.Required, Group: Schserviceconnectorsourcelogsourcesrepresentation2},
	}

	serviceConnectorRepresentationNoTarget2 = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `My_Service_Connector`, Update: `displayName2`},
		"source":         acctest.RepresentationGroup{RepType: acctest.Required, Group: Schserviceconnectorsourcerepresentation2},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `My service connector description`, Update: `description2`},
	}

	logAnalyticsTargetRepresentationForLogSource = map[string]interface{}{
		"kind":         acctest.Representation{RepType: acctest.Required, Create: `loggingAnalytics`},
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.logAn_log_group_ocid}`},
	}

	serviceConnectorLogAnTargetLoggingSourceRepresentation2 = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget2, logAnalyticsTargetRepresentationForLogSource)
)

// issue-routing-tag: sch/default
func TestSchServiceConnectorResource_streamingAnalytics(t *testing.T) {
	httpreplay.SetScenario("TestSchServiceConnectorResource_streamingAnalytics")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	image := utils.GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	logAnLogGroupId := utils.GetEnvSettingWithBlankDefault("logAn_log_group_ocid")
	logAnLogGroupIdVariableStr := fmt.Sprintf("variable \"logAn_log_group_ocid\" { default = \"%s\" }\n", logAnLogGroupId)

	logAnLogSourceName := utils.GetEnvSettingWithBlankDefault("logAn_log_source_name")
	logAnLogSourceNameVariableStr := fmt.Sprintf("variable \"logAn_log_source_name\" { default = \"%s\" }\n", logAnLogSourceName)

	resourceName := "oci_sch_service_connector.test_service_connector"
	singularDatasourceName := "data.oci_sch_service_connector.test_service_connector"

	var resId, resId2 string
	acctest.ResourceTest(t, testAccCheckSchServiceConnectorDestroy, []resource.TestStep{
		// verify private streaming as a source with functions target
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependenciesForPrivateStream + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorFunctionTargetStreamingSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "streaming"),
				resource.TestCheckResourceAttr(resourceName, "source.0.cursor.0.kind", "LATEST"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.stream_id"),
				resource.TestCheckResourceAttr(resourceName, "source.0.private_endpoint_metadata.#", "1"),
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
			Config: config + compartmentIdVariableStr + imageVariableStr,
		},

		// verify streaming as a source with functions target
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorFunctionTargetStreamingSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "lifecycle_details", ""),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "streaming"),
				resource.TestCheckResourceAttr(resourceName, "source.0.cursor.0.kind", "LATEST"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.stream_id"),
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

		// verify streaming as a source with functions task and functions target
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorFunctionTargetStreamingSourceFunctionTaskRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "lifecycle_details", ""),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "streaming"),
				resource.TestCheckResourceAttr(resourceName, "source.0.cursor.0.kind", "LATEST"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.stream_id"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "functions"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.function_id"),
				resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tasks.0.function_id"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.batch_size_in_kbs", "60"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.batch_time_in_sec", "60"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "function"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(serviceConnectorFunctionTargetStreamingSourceFunctionTaskRepresentation, []string{"target"}), map[string]interface{}{
						"source": acctest.RepresentationGroup{RepType: acctest.Optional, Group: serviceConnectorStreamingSourceRepresentation},
						"target": acctest.RepresentationGroup{RepType: acctest.Required, Group: SchServiceConnectorTargetRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "lifecycle_details", ""),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "streaming"),
				resource.TestCheckResourceAttr(resourceName, "source.0.cursor.0.kind", "TRIM_HORIZON"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.stream_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "streaming"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.stream_id"),
				resource.TestCheckResourceAttr(resourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "tasks.0.function_id"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.batch_size_in_kbs", "60"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.batch_time_in_sec", "60"),
				resource.TestCheckResourceAttr(resourceName, "tasks.0.kind", "function"),
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

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(serviceConnectorFunctionTargetRepresentation, []string{"source", "task", "target"}), map[string]interface{}{
						"source": acctest.RepresentationGroup{RepType: acctest.Optional, Group: updatedServiceConnectorStreamingSourceRepresentation},
						"tasks":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: updatedServiceConnectorFunctionTasksRepresentation},
						"target": acctest.RepresentationGroup{RepType: acctest.Required, Group: SchServiceConnectorTargetRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "service_connector_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lifecycle_details", ""),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.kind", "streaming"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source.0.cursor.0.kind", "TRIM_HORIZON"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source.0.stream_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target.0.kind", "streaming"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tasks.0.function_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.batch_size_in_kbs", "60"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.batch_time_in_sec", "60"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tasks.0.kind", "function"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr,
		},

		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr + logAnLogGroupIdVariableStr + logAnLogSourceNameVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorLogAnTargetStreamingSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "streaming"),
				resource.TestCheckResourceAttr(resourceName, "source.0.cursor.0.kind", "LATEST"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.stream_id"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "loggingAnalytics"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.log_source_identifier"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr + logAnLogGroupIdVariableStr + logAnLogSourceNameVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(serviceConnectorLogAnTargetStreamingSourceRepresentation, []string{"target"}), map[string]interface{}{
						"target": acctest.RepresentationGroup{RepType: acctest.Required, Group: logAnalyticsTargetRepresentation},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "streaming"),
				resource.TestCheckResourceAttr(resourceName, "source.0.cursor.0.kind", "TRIM_HORIZON"),
				resource.TestCheckResourceAttrSet(resourceName, "source.0.stream_id"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "loggingAnalytics"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.log_group_id"),
				resource.TestCheckResourceAttr(resourceName, "target.0.log_source_identifier", "LinuxSyslogSource"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
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

/*
Test for validating update operation bug tracked by https://jira.oci.oraclecorp.com/browse/OCH-1877
TL;DR - 'logSourceIdentifier' attribute must only be supplied in the request if the value was explicitly updated
*/
func TestSchServiceConnectorResource_LogSrc_LogAnTarget(t *testing.T) {
	httpreplay.SetScenario("TestSchServiceConnectorResource_LogSrc_LogAnTarget")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	logAnLogGroupId := utils.GetEnvSettingWithBlankDefault("logAn_log_group_ocid")
	logAnLogGroupIdVariableStr := fmt.Sprintf("variable \"logAn_log_group_ocid\" { default = \"%s\" }\n", logAnLogGroupId)

	resourceName := "oci_sch_service_connector.test_service_connector"
	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckSchServiceConnectorDestroy, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + logAnLogGroupIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorLogAnTargetLoggingSourceRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		{
			Config: config + compartmentIdVariableStr + logAnLogGroupIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Update, serviceConnectorLogAnTargetLoggingSourceRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
	})
}
