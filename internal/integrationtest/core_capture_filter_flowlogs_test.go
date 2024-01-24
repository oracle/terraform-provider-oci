// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreCaptureFilterFlowLogsRequiredOnlyResource = CoreCaptureFilterFlowLogsResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Required, acctest.Create, CoreCaptureFilterFlowLogsRepresentation)

	CoreCaptureFilterFlowLogsResourceConfig = CoreCaptureFilterFlowLogsResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Update, CoreCaptureFilterFlowLogsRepresentation)

	CoreCoreCaptureFilterFlowLogsSingularDataSourceRepresentation = map[string]interface{}{
		"capture_filter_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_capture_filter.test_capture_filter.id}`},
	}

	CoreCoreCaptureFilterFlowLogsDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MyCaptureFilter`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreCaptureFilterFlowLogsDataSourceFilterRepresentation}}
	CoreCaptureFilterFlowLogsDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_capture_filter.test_capture_filter.id}`}},
	}

	CoreCaptureFilterFlowLogsRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		// maybe this should be required now, and in VTAP file, also check if it creates VTAP capture filter if filter_type is null (or make that optional there)
		"filter_type":                   acctest.Representation{RepType: acctest.Required, Create: `FLOWLOG`},
		"flow_log_capture_filter_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreCaptureFilterFlowLogCaptureFilterRulesRepresentation},
		"defined_tags":                  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `MyCaptureFilter`, Update: `displayName2`},
		"freeform_tags":                 acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	CoreCaptureFilterFlowLogCaptureFilterRulesRepresentation = map[string]interface{}{
		"flow_log_type":    acctest.Representation{RepType: acctest.Required, Create: `ALL`, Update: `REJECT`},
		"is_enabled":       acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"priority":         acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"rule_action":      acctest.Representation{RepType: acctest.Required, Create: `INCLUDE`, Update: `EXCLUDE`},
		"sampling_rate":    acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"destination_cidr": acctest.Representation{RepType: acctest.Optional, Create: `10.3.0.0/16`, Update: `10.4.0.0/16`},
		"protocol":         acctest.Representation{RepType: acctest.Optional, Create: `6`},
		"source_cidr":      acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.1/32`, Update: `10.0.0.2/32`},
		"tcp_options":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCaptureFilterFlowLogCaptureFilterRulesTcpOptionsRepresentation},
	}

	CoreCaptureFilterFlowLogCaptureFilterRulesTcpOptionsRepresentation = map[string]interface{}{
		"destination_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCaptureFilterFlowLogCaptureFilterRulesTcpOptionsDestinationPortRangeRepresentation},
		"source_port_range":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCaptureFilterFlowLogCaptureFilterRulesTcpOptionsSourcePortRangeRepresentation},
	}

	CoreCaptureFilterFlowLogCaptureFilterRulesTcpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	CoreCaptureFilterFlowLogCaptureFilterRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	// These are needed only if the protocol is set to UDP/ICMP. If its set to TCP then these fields will be null
	//CoreCaptureFilterFlowLogCaptureFilterRulesUdpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
	//	"max": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	//	"min": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	//}
	//
	//CoreCaptureFilterFlowLogCaptureFilterRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
	//	"max": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	//	"min": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	//}
	//
	//CoreCaptureFilterFlowLogCaptureFilterRulesUdpOptionsRepresentation = map[string]interface{}{
	//	"destination_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCaptureFilterFlowLogCaptureFilterRulesUdpOptionsDestinationPortRangeRepresentation},
	//	"source_port_range":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreCaptureFilterFlowLogCaptureFilterRulesUdpOptionsSourcePortRangeRepresentation},
	//}
	//
	//CoreCaptureFilterFlowLogCaptureFilterRulesIcmpOptionsRepresentation = map[string]interface{}{
	//	"type": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	//	"code": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	//}

	CoreCaptureFilterFlowLogsResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreCaptureFilterFlowLogsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCaptureFilterFlowLogsResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_core_capture_filter.test_capture_filter"
	datasourceName := "data.oci_core_capture_filters.test_capture_filters"
	singularDatasourceName := "data.oci_core_capture_filter.test_capture_filter"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreCaptureFilterFlowLogsResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Create, CoreCaptureFilterFlowLogsRepresentation), "core", "captureFilter", t)

	acctest.ResourceTest(t, testAccCheckCoreCaptureFilterFlowLogsDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreCaptureFilterFlowLogsResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Required, acctest.Create, CoreCaptureFilterFlowLogsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "filter_type", "FLOWLOG"),

				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.flow_log_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.rule_action", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.sampling_rate", "10"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.priority", "2"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CoreCaptureFilterFlowLogsResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreCaptureFilterFlowLogsResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Create, CoreCaptureFilterFlowLogsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyCaptureFilter"),
				resource.TestCheckResourceAttr(resourceName, "filter_type", "FLOWLOG"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.destination_cidr", "10.3.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.flow_log_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.priority", "2"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.protocol", "6"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.rule_action", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.sampling_rate", "10"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.source_cidr", "10.0.0.1/32"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.0.max", "10"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.0.min", "10"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.0.max", "10"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.0.min", "10"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CoreCaptureFilterFlowLogsResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreCaptureFilterFlowLogsRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyCaptureFilter"),
				resource.TestCheckResourceAttr(resourceName, "filter_type", "FLOWLOG"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.destination_cidr", "10.3.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.flow_log_type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.priority", "2"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.protocol", "6"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.rule_action", "INCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.sampling_rate", "10"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.source_cidr", "10.0.0.1/32"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.0.max", "10"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.0.min", "10"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.0.max", "10"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.0.min", "10"),

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
			Config: config + compartmentIdVariableStr + CoreCaptureFilterFlowLogsResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Update, CoreCaptureFilterFlowLogsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "filter_type", "FLOWLOG"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.destination_cidr", "10.4.0.0/16"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.flow_log_type", "REJECT"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.priority", "3"),

				// protocol cannot be updated
				//resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.protocol", "protocol2"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.rule_action", "EXCLUDE"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.sampling_rate", "11"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.source_cidr", "10.0.0.2/32"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.0.max", "11"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.0.min", "11"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.0.max", "11"),
				resource.TestCheckResourceAttr(resourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.0.min", "11"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_capture_filters", "test_capture_filters", acctest.Optional, acctest.Update, CoreCoreCaptureFilterFlowLogsDataSourceRepresentation) +
				compartmentIdVariableStr + CoreCaptureFilterFlowLogsResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Optional, acctest.Update, CoreCaptureFilterFlowLogsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "capture_filters.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.filter_type", "FLOWLOG"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "capture_filters.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "capture_filters.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "capture_filters.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.destination_cidr", "10.4.0.0/16"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.flow_log_type", "REJECT"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.priority", "3"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.protocol", "6"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.rule_action", "EXCLUDE"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.sampling_rate", "11"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.source_cidr", "10.0.0.2/32"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.tcp_options.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.0.max", "11"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.0.min", "11"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.0.max", "11"),
				resource.TestCheckResourceAttr(datasourceName, "capture_filters.0.flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.0.min", "11"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_capture_filter", "test_capture_filter", acctest.Required, acctest.Create, CoreCoreCaptureFilterFlowLogsSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CoreCaptureFilterFlowLogsResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "capture_filter_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "filter_type", "FLOWLOG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.destination_cidr", "10.4.0.0/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.flow_log_type", "REJECT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.priority", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.protocol", "6"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.rule_action", "EXCLUDE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.sampling_rate", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.source_cidr", "10.0.0.2/32"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.tcp_options.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.0.max", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.tcp_options.0.destination_port_range.0.min", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.0.max", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "flow_log_capture_filter_rules.0.tcp_options.0.source_port_range.0.min", "11"),
			),
		},
		// verify resource import
		{
			Config:                  config + CoreCaptureFilterFlowLogsRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCoreCaptureFilterFlowLogsDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).VirtualNetworkClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_core_capture_filter" {
			noResourceFound = false
			request := oci_core.GetCaptureFilterRequest{}

			tmp := rs.Primary.ID
			request.CaptureFilterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "core")

			response, err := client.GetCaptureFilter(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_core.CaptureFilterLifecycleStateTerminated): true,
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
