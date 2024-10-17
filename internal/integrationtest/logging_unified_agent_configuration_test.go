// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_logging "github.com/oracle/oci-go-sdk/v65/logging"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	LoggingUnifiedAgentConfigurationRequiredOnlyResource = LoggingUnifiedAgentConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Required, acctest.Create, LoggingUnifiedAgentConfigurationRepresentation)

	LoggingUnifiedAgentConfigurationResourceConfig = LoggingUnifiedAgentConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update, LoggingUnifiedAgentConfigurationRepresentation)

	LoggingLoggingUnifiedAgentConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"unified_agent_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_unified_agent_configuration.test_unified_agent_configuration.id}`},
	}

	LoggingLoggingUnifiedAgentConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationDataSourceFilterRepresentation},
	}

	LoggingUnifiedAgentConfigurationRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_enabled":            acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"service_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationRepresentation},
		"description":           acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"group_association":     acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationGroupAssociationRepresentation},
	}

	LoggingUnifiedAgentConfigurationServiceConfigurationRepresentation = map[string]interface{}{
		"configuration_type":                 acctest.Representation{RepType: acctest.Required, Create: `LOGGING`},
		"destination":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationDestinationRepresentation},
		"sources":                            acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationSourcesRepresentation},
		"unified_agent_configuration_filter": acctest.RepresentationGroup{RepType: acctest.Optional, Group: LoggingUnifiedAgentConfigurationServiceConfigurationUnifiedAgentConfigurationFilterRepresentation},
	}
	LoggingUnifiedAgentConfigurationServiceConfigurationUnifiedAgentConfigurationFilterRepresentation = map[string]interface{}{
		"filter_type": acctest.Representation{RepType: acctest.Required, Create: `GREP_FILTER`},
		"allow_list":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: LoggingUnifiedAgentConfigurationServiceConfigurationUnifiedAgentConfigurationFilterAllowListRepresentation},
		"deny_list":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: LoggingUnifiedAgentConfigurationServiceConfigurationUnifiedAgentConfigurationFilterDenyListRepresentation},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
	}
	LoggingUnifiedAgentConfigurationServiceConfigurationUnifiedAgentConfigurationFilterAllowListRepresentation = map[string]interface{}{
		"key":     acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"pattern": acctest.Representation{RepType: acctest.Optional, Create: `pattern`, Update: `pattern2`},
	}
	LoggingUnifiedAgentConfigurationServiceConfigurationUnifiedAgentConfigurationFilterDenyListRepresentation = map[string]interface{}{
		"key":     acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"pattern": acctest.Representation{RepType: acctest.Optional, Create: `pattern`, Update: `pattern2`},
	}
	LoggingUnifiedAgentConfigurationGroupAssociationRepresentation = map[string]interface{}{
		//"group_list": acctest.Representation{RepType: acctest.Required, Create: []string{`ocid1.group.oc1..aaaaaaaafxpft7qucqbpsygm555uzxipqslne7d5meupykscq57q32jfiifa`}}, // Update: []string{`${oci_identity_group.test_group.id}`, `ocid1.Group.oc1..aaaaaaaa5rvs7zjwdk3zdmysm7x7wcxyanbllutswe4xbl7ng4stohtg3sla`}},
		"group_list": acctest.Representation{RepType: acctest.Required, Create: []string{`ocid1.group.oc1..testid`}}, // Update: []string{`${oci_identity_group.test_group.id}`, `ocid1.Group.oc1..aaaaaaaa5rvs7zjwdk3zdmysm7x7wcxyanbllutswe4xbl7ng4stohtg3sla`}},
	}

	LoggingUnifiedAgentConfigurationServiceConfigurationDestinationRepresentation = map[string]interface{}{
		"log_object_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_log.id}`},
		"operational_metrics_configuration": acctest.RepresentationGroup{RepType: acctest.Optional, Group: LoggingUnifiedAgentConfigurationServiceConfigurationDestinationOperationalMetricsConfigurationRepresentation},
	}

	LoggingUnifiedAgentConfigurationServiceConfigurationDestinationOperationalMetricsConfigurationRepresentation = map[string]interface{}{
		"destination": acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationDestinationOperationalMetricsConfigurationDestinationRepresentation},
		"source":      acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationDestinationOperationalMetricsConfigurationSourceRepresentation},
	}

	// operational metrics dest
	LoggingUnifiedAgentConfigurationServiceConfigurationDestinationOperationalMetricsConfigurationDestinationRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	// operational metrics source
	LoggingUnifiedAgentConfigurationServiceConfigurationDestinationOperationalMetricsConfigurationSourceRepresentation = map[string]interface{}{
		"type":         acctest.Representation{RepType: acctest.Required, Create: `UMA_METRICS`},
		"metrics":      acctest.Representation{RepType: acctest.Optional, Create: []string{`RestartMetric`}, Update: []string{`EmitRecords`}},
		"record_input": acctest.RepresentationGroup{RepType: acctest.Optional, Group: LoggingUnifiedAgentConfigurationServiceConfigurationDestinationOperationalMetricsConfigurationSourceRecordInputRepresentation},
	}

	LoggingUnifiedAgentConfigurationServiceConfigurationDestinationOperationalMetricsConfigurationSourceRecordInputRepresentation = map[string]interface{}{
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `namespace`, Update: `namespace2`},
		"resource_group": acctest.Representation{RepType: acctest.Optional, Create: `resourceGroup`, Update: `resourceGroup2`},
	}

	LoggingUnifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation = map[string]interface{}{
		"parser_type":             acctest.Representation{RepType: acctest.Required, Create: `AUDITD`},
		"field_time_key":          acctest.Representation{RepType: acctest.Optional, Create: `fieldTimeKey`},
		"is_keep_time_key":        acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_null_empty_string":    acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"null_value_pattern":      acctest.Representation{RepType: acctest.Optional, Create: `nullValuePattern`},
		"timeout_in_milliseconds": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"types":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"types": "types"}},
	}

	LoggingUnifiedAgentConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_logging_unified_agent_configuration.test_unified_agent_configuration.id}`}},
	}

	LoggingUAIdentityGroupRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}`},
		"description":    acctest.Representation{RepType: acctest.Required, Create: `Group for Logging UA`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: `LoggingAgentIdentityGroup`},
	}

	// need to add policy for creating groups
	LoggingUnifiedAgentConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, LoggingLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, customLogRepresentation) // +

	LoggingUnifiedAgentConfigurationServiceConfigurationSourcesRepresentation = map[string]interface{}{
		"advanced_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsSourcesAdvancedOptionsRepresentation},
		"source_type":      acctest.Representation{RepType: acctest.Required, Create: `LOG_TAIL`},
		"paths":            acctest.Representation{RepType: acctest.Required, Create: []string{`/var/log/*`}},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `name`},
		"parser":           acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation},
	}
	LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsSourcesAdvancedOptionsRepresentation = map[string]interface{}{
		"is_read_from_head": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	// Add new test configs here
	// CRI parser configs
	LoggingUnifiedAgentConfigurationServiceConfigurationSourcesParserNestedParserRepresentation = map[string]interface{}{
		"time_format":    acctest.Representation{RepType: acctest.Optional, Create: `%Y-%m-%dT%H:%M:%S.%L%z`, Update: `%Y-%m-%d %H:%M:%S.%L%z`},
		"field_time_key": acctest.Representation{RepType: acctest.Optional, Create: `time`, Update: `time1`},
	}

	LoggingUnifiedAgentConfigurationServiceConfigurationSourcesCriParserRepresentation = map[string]interface{}{
		"parser_type":         acctest.Representation{RepType: acctest.Required, Create: `CRI`},
		"is_merge_cri_fields": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"nested_parser":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: LoggingUnifiedAgentConfigurationServiceConfigurationSourcesParserNestedParserRepresentation},
	}

	LoggingUnifiedAgentConfigurationCriRepresentation = acctest.GetUpdatedRepresentationCopy(
		"service_configuration.sources.parser",
		acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationSourcesCriParserRepresentation},
		LoggingUnifiedAgentConfigurationRepresentation)
)

// issue-routing-tag: logging/default
func TestLoggingUnifiedAgentConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingUnifiedAgentConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_logging_unified_agent_configuration.test_unified_agent_configuration"
	datasourceName := "data.oci_logging_unified_agent_configurations.test_unified_agent_configurations"
	singularDatasourceName := "data.oci_logging_unified_agent_configuration.test_unified_agent_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LoggingUnifiedAgentConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Create, LoggingUnifiedAgentConfigurationRepresentation), "logging", "unifiedAgentConfiguration", t)

	acctest.ResourceTest(t, testAccCheckLoggingUnifiedAgentConfigurationDestroy, []resource.TestStep{

		// Add new tests here
		// CRI parser test required
		{
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Required, acctest.Create, LoggingUnifiedAgentConfigurationCriRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "LOGGING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "service_configuration.0.destination.0.log_object_id"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "CRI"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// CRI parser test optional
		{
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Create, LoggingUnifiedAgentConfigurationCriRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "LOGGING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "service_configuration.0.destination.0.log_object_id"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "CRI"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_merge_cri_fields", "true"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.nested_parser.0.time_format", "%Y-%m-%dT%H:%M:%S.%L%z"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.nested_parser.0.field_time_key", "time"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.nested_parser.0.is_keep_time_key", "false"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// CRI parser test optional update
		{
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update, LoggingUnifiedAgentConfigurationCriRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "LOGGING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "service_configuration.0.destination.0.log_object_id"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "CRI"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_merge_cri_fields", "false"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.nested_parser.0.time_format", "%Y-%m-%d %H:%M:%S.%L%z"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.nested_parser.0.field_time_key", "time1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.nested_parser.0.is_keep_time_key", "false"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Don't change below tests
		// 0. verify Create
		{
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Required, acctest.Create, LoggingUnifiedAgentConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "LOGGING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "service_configuration.0.destination.0.log_object_id"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// 1. delete before next Create
		{
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationResourceDependencies,
		},
		// 2. verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Create, LoggingUnifiedAgentConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_association.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_association.0.group_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "LOGGING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "service_configuration.0.destination.0.log_object_id"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.destination.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.metrics.0", "RestartMetric"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.0.resource_group", "resourceGroup"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.type", "UMA_METRICS"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
				//resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "AUDITD"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.advanced_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.advanced_options.0.is_read_from_head", "false"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.0.pattern", "pattern"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.0.pattern", "pattern"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.filter_type", "GREP_FILTER"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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

		// 3. verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LoggingUnifiedAgentConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LoggingUnifiedAgentConfigurationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_association.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_association.0.group_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "LOGGING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "service_configuration.0.destination.0.log_object_id"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.destination.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.metrics.0", "RestartMetric"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.0.resource_group", "resourceGroup"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.type", "UMA_METRICS"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.0.pattern", "pattern"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.0.pattern", "pattern"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.filter_type", "GREP_FILTER"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.advanced_options.0.is_read_from_head", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// 4. verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update, LoggingUnifiedAgentConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_association.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_association.0.group_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "LOGGING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "service_configuration.0.destination.0.log_object_id"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.destination.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.metrics.0", "EmitRecords"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.0.resource_group", "resourceGroup2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.type", "UMA_METRICS"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.0.pattern", "pattern2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.0.pattern", "pattern2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.filter_type", "GREP_FILTER"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.unified_agent_configuration_filter.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.advanced_options.0.is_read_from_head", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// 5. verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_logging_unified_agent_configurations", "test_unified_agent_configurations", acctest.Optional, acctest.Update, LoggingLoggingUnifiedAgentConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + LoggingUnifiedAgentConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update, LoggingUnifiedAgentConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(datasourceName, "unified_agent_configuration_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "unified_agent_configuration_collection.0.items.#", "1"),
			),
		},

		// 6. verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Required, acctest.Create, LoggingLoggingUnifiedAgentConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LoggingUnifiedAgentConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "unified_agent_configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "configuration_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "group_association.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "group_association.0.group_list.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.configuration_type", "LOGGING"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.destination.0.operational_metrics_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.destination.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.destination.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.metrics.0", "EmitRecords"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.record_input.0.resource_group", "resourceGroup2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.destination.0.operational_metrics_configuration.0.source.0.type", "UMA_METRICS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.advanced_options.0.is_read_from_head", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.unified_agent_configuration_filter.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.unified_agent_configuration_filter.0.allow_list.0.pattern", "pattern2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.unified_agent_configuration_filter.0.deny_list.0.pattern", "pattern2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.unified_agent_configuration_filter.0.filter_type", "GREP_FILTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.unified_agent_configuration_filter.0.name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
			),
		},
		// 7. verify resource import
		{
			Config:                  config + LoggingUnifiedAgentConfigurationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLoggingUnifiedAgentConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LoggingManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_logging_unified_agent_configuration" {
			noResourceFound = false
			request := oci_logging.GetUnifiedAgentConfigurationRequest{}

			tmp := rs.Primary.ID
			request.UnifiedAgentConfigurationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "logging")

			_, err := client.GetUnifiedAgentConfiguration(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("LoggingUnifiedAgentConfiguration") {
		resource.AddTestSweepers("LoggingUnifiedAgentConfiguration", &resource.Sweeper{
			Name:         "LoggingUnifiedAgentConfiguration",
			Dependencies: acctest.DependencyGraph["unifiedAgentConfiguration"],
			F:            sweepLoggingUnifiedAgentConfigurationResource,
		})
	}
}

func sweepLoggingUnifiedAgentConfigurationResource(compartment string) error {
	loggingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).LoggingManagementClient()
	unifiedAgentConfigurationIds, err := getLoggingUnifiedAgentConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, unifiedAgentConfigurationId := range unifiedAgentConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[unifiedAgentConfigurationId]; !ok {
			deleteUnifiedAgentConfigurationRequest := oci_logging.DeleteUnifiedAgentConfigurationRequest{}

			deleteUnifiedAgentConfigurationRequest.UnifiedAgentConfigurationId = &unifiedAgentConfigurationId

			deleteUnifiedAgentConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "logging")
			_, error := loggingManagementClient.DeleteUnifiedAgentConfiguration(context.Background(), deleteUnifiedAgentConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting UnifiedAgentConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", unifiedAgentConfigurationId, error)
				continue
			}
		}
	}
	return nil
}

func getLoggingUnifiedAgentConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "UnifiedAgentConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loggingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).LoggingManagementClient()

	listUnifiedAgentConfigurationsRequest := oci_logging.ListUnifiedAgentConfigurationsRequest{}
	listUnifiedAgentConfigurationsRequest.CompartmentId = &compartmentId
	listUnifiedAgentConfigurationsResponse, err := loggingManagementClient.ListUnifiedAgentConfigurations(context.Background(), listUnifiedAgentConfigurationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting UnifiedAgentConfiguration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, unifiedAgentConfiguration := range listUnifiedAgentConfigurationsResponse.Items {
		id := *unifiedAgentConfiguration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "UnifiedAgentConfigurationId", id)
	}
	return resourceIds, nil
}
