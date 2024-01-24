// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	// targets for Monitoring with logging as a source
	//Monitoring Target without dimensions
	serviceConnectorMonitoringTargetLoggingSourceRepresentation = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, monitoringTargetRepresentation)
	//Monitoring Target with dimensions representation
	serviceConnectorMonitoringTargetStaticDimLoggingSourceRepresentation         = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, monitoringTargetStaticDimensionRepresentation)
	serviceConnectorMonitoringTargetJmesPathLoggingSourceRepresentation          = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, monitoringTargetJmesPathDimensionRepresentation)
	serviceConnectorMonitoringTargetStaticAndJmesPathLoggingSourceRepresentation = createServiceConnectorRepresentation(serviceConnectorRepresentationNoTarget, monitoringTargetStaticAndJmesPathRepresentation)
	//Monitoring Source representation
	serviceConnectorMonitoringSourceObjectStorageTargetRepresentation = createServiceConnectorRepresentation(serviceConnectorRepresentationMonitoringSource, objectStorageTargetRepresentation)
	//Monitoring source with multiple namespaces representations
	serviceConnectorMonitoringSourceMultipleNameSpacesObjectStorageTargetRepresentation = createServiceConnectorRepresentation(serviceConnectorRepresentationMonitoringSourceMultipleNamespaces, objectStorageTargetRepresentation)
)

// issue-routing-tag: sch/default
func TestSchServiceConnectorResource_monitoring(t *testing.T) {
	httpreplay.SetScenario("TestSchServiceConnectorResource_monitoring")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_sch_service_connector.test_service_connector"

	image := utils.GetEnvSettingWithBlankDefault("image")
	imageVariableStr := fmt.Sprintf("variable \"image\" { default = \"%s\" }\n", image)

	var resId string

	acctest.ResourceTest(t, testAccCheckSchServiceConnectorDestroy, []resource.TestStep{
		//  verify logging as source with monitoring as target
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorMonitoringTargetLoggingSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "monitoring"),
				resource.TestCheckResourceAttr(resourceName, "target.0.metric", "metric"),
				resource.TestCheckResourceAttr(resourceName, "target.0.metric_namespace", "metricnamespace"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					_ = resId
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr,
		},

		// verify logging as source and monitoring with dimensions as target
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorMonitoringTargetJmesPathLoggingSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "monitoring"),
				resource.TestCheckResourceAttr(resourceName, "target.0.metric", "jmespath_metric_0"),
				resource.TestCheckResourceAttr(resourceName, "target.0.metric_namespace", "jmespath_metricnamespace_0"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.dimension_value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.dimension_value.0.kind", "jmesPath"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.dimension_value.0.path", "logContent.data.compartmentId"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.name", "jmespath_dimension_0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					_ = resId
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr,
		},

		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorMonitoringTargetStaticDimLoggingSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "monitoring"),
				resource.TestCheckResourceAttr(resourceName, "target.0.metric", "static_metric_0"),
				resource.TestCheckResourceAttr(resourceName, "target.0.metric_namespace", "static_metricnamespace_0"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.dimension_value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.dimension_value.0.kind", "static"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.dimension_value.0.value", "static_value_0"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.name", "static_dimension_0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					_ = resId
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr,
		},
		//Testing create connector with monitoring target for 4 dimensions -> 2 jmespath and 2 static
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorMonitoringTargetStaticAndJmesPathLoggingSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "logging"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.log_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "monitoring"),
				resource.TestCheckResourceAttr(resourceName, "target.0.metric", "metric"),
				resource.TestCheckResourceAttr(resourceName, "target.0.metric_namespace", "metricnamespace"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.#", "4"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.dimension_value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.dimension_value.0.kind", "jmesPath"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.dimension_value.0.path", "logContent.data.compartmentId"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.0.name", "jmespath_dimension_0"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.1.dimension_value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.1.dimension_value.0.kind", "jmesPath"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.1.dimension_value.0.path", "logContent.data.namespace"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.1.name", "jmespath_dimension_1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.2.dimension_value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.2.dimension_value.0.kind", "static"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.2.dimension_value.0.value", "static_value_0"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.2.name", "static_dimension_0"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.3.dimension_value.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.3.dimension_value.0.kind", "static"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.3.dimension_value.0.value", "static_value_1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.dimensions.3.name", "static_dimension_1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					_ = resId
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr,
		},

		//  verify monitoring as source with object storage as target
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorMonitoringSourceObjectStorageTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "monitoring"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.0.kind", "selected"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.0.namespaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.0.namespaces.0.namespace", "oci_computeagent"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.0.namespaces.0.metrics.0.kind", "all"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "objectStorage"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.bucket"),
				resource.TestCheckResourceAttr(resourceName, "target.0.batch_rollover_size_in_mbs", "100"),
				resource.TestCheckResourceAttr(resourceName, "target.0.batch_rollover_time_in_ms", "420000"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					_ = resId
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr,
		},

		//  verify monitoring as source with object storage as target with multiple namespaces
		{
			Config: config + compartmentIdVariableStr + SchServiceConnectorResourceDependencies + imageVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_sch_service_connector", "test_service_connector", acctest.Required, acctest.Create, serviceConnectorMonitoringSourceMultipleNameSpacesObjectStorageTargetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "My_Service_Connector"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.kind", "monitoring"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.0.kind", "selected"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.0.namespaces.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.0.namespaces.0.namespace", "oci_computeagent"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.0.namespaces.0.metrics.0.kind", "all"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.0.namespaces.1.namespace", "oci_logging_analytics"),
				resource.TestCheckResourceAttr(resourceName, "source.0.monitoring_sources.0.namespace_details.0.namespaces.1.metrics.0.kind", "all"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.kind", "objectStorage"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.bucket"),
				resource.TestCheckResourceAttr(resourceName, "target.0.batch_rollover_size_in_mbs", "100"),
				resource.TestCheckResourceAttr(resourceName, "target.0.batch_rollover_time_in_ms", "420000"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					_ = resId
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
