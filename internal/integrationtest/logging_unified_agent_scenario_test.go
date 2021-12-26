// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	UnifiedAgentConfigurationLogTailRequiredOnlyResource = UnifiedAgentConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Required, acctest.Create, unifiedAgentConfigurationLogTailRepresentation)

	UnifiedAgentConfigurationLogTailResourceConfig = UnifiedAgentConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update, unifiedAgentConfigurationLogTailRepresentation)

	unifiedAgentConfigurationLogTailSingularDataSourceRepresentation = map[string]interface{}{
		"unified_agent_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_unified_agent_configuration.test_unified_agent_configuration.id}`},
	}

	unifiedAgentConfigurationLogTailDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"group_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_group.test_group.id}`},
		"is_compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"log_id":                       acctest.Representation{RepType: acctest.Optional, Create: `${oci_logging_log.test_log.id}`},
		"state":                        acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: unifiedAgentConfigurationLogTailDataSourceFilterRepresentation}}
	unifiedAgentConfigurationLogTailDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_logging_unified_agent_configuration.test_unified_agent_configuration.id}`}},
	}

	unifiedAgentConfigurationLogTailRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_enabled":            acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"service_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: unifiedAgentConfigurationServiceConfigurationLogTailRepresentation},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":           acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"group_association":     acctest.RepresentationGroup{RepType: acctest.Required, Group: unifiedAgentConfigurationGroupAssociationRepresentation},
	}
	unifiedAgentConfigurationServiceConfigurationLogTailRepresentation = map[string]interface{}{
		"configuration_type": acctest.Representation{RepType: acctest.Required, Create: `LOGGING`},
		"destination":        acctest.RepresentationGroup{RepType: acctest.Required, Group: unifiedAgentConfigurationServiceConfigurationDestinationRepresentation},
		"sources":            acctest.RepresentationGroup{RepType: acctest.Required, Group: unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation},
	}

	unifiedAgentConfigurationServiceConfigurationDestinationRepresentation = map[string]interface{}{
		"log_object_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_log.id}`},
	}
	unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation = map[string]interface{}{
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `LOG_TAIL`},
		"name":        acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"parser":      acctest.RepresentationGroup{RepType: acctest.Required, Group: unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation},
		"paths":       acctest.Representation{RepType: acctest.Optional, Create: []string{`paths`}},
	}

	unifiedAgentConfigurationServiceConfigurationSourcesJSONParserRepresentation = acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update,
		acctest.GetUpdatedRepresentationCopy("service_configuration", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("sources", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("parser", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
			"parser_type":               acctest.Representation{RepType: acctest.Required, Update: `JSON`},
			"field_time_key":            acctest.Representation{RepType: acctest.Optional, Update: `fieldTimeKey2`},
			"is_estimate_current_event": acctest.Representation{RepType: acctest.Optional, Update: `true`},
			"is_keep_time_key":          acctest.Representation{RepType: acctest.Optional, Update: `true`},
			"is_null_empty_string":      acctest.Representation{RepType: acctest.Optional, Update: `true`},
			"null_value_pattern":        acctest.Representation{RepType: acctest.Optional, Update: `nullValuePattern2`},
			"timeout_in_milliseconds":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
			"types":                     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"types": "types"}},
			"time_format":               acctest.Representation{RepType: acctest.Optional, Update: `timeFormat2`},
			"time_type":                 acctest.Representation{RepType: acctest.Optional, Update: `UNIXTIME`}})},
			unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
			unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	unifiedAgentConfigurationServiceConfigurationSourcesCSVParserRepresentation = acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update,
		acctest.GetUpdatedRepresentationCopy("service_configuration", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("sources", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("parser", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
			"parser_type": acctest.Representation{RepType: acctest.Required, Update: `CSV`},
			"delimiter":   acctest.Representation{RepType: acctest.Optional, Update: `delimiter`},
			"keys":        acctest.Representation{RepType: acctest.Optional, Update: []string{`key1`}}})},
			unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
			unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	unifiedAgentConfigurationServiceConfigurationSourcesGROKParserRepresentation = acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update,
		acctest.GetUpdatedRepresentationCopy("service_configuration", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("sources", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("parser", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
			"parser_type":      acctest.Representation{RepType: acctest.Required, Update: `GROK`},
			"grok_failure_key": acctest.Representation{RepType: acctest.Optional, Update: `grokFailureKey2`},
			"grok_name_key":    acctest.Representation{RepType: acctest.Optional, Update: `grokNameKey2`},
			"patterns":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: unifiedAgentConfigurationServiceConfigurationSourcesParserPatternsRepresentation},
		})},
			unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
			unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	unifiedAgentConfigurationServiceConfigurationSourcesMSGPACKParserRepresentation = acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update,
		acctest.GetUpdatedRepresentationCopy("service_configuration", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("sources", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("parser", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
			"parser_type":               acctest.Representation{RepType: acctest.Required, Update: `MSGPACK`},
			"field_time_key":            acctest.Representation{RepType: acctest.Optional, Update: `fieldTimeKey3`},
			"is_estimate_current_event": acctest.Representation{RepType: acctest.Optional, Update: `true`},
			"is_keep_time_key":          acctest.Representation{RepType: acctest.Optional, Update: `true`},
			"is_null_empty_string":      acctest.Representation{RepType: acctest.Optional, Update: `true`},
			"null_value_pattern":        acctest.Representation{RepType: acctest.Optional, Update: `nullValuePattern3`},
			"timeout_in_milliseconds":   acctest.Representation{RepType: acctest.Optional, Update: `11`},
			"types":                     acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"types": "types2"}},
		})},
			unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
			unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	//MULTILINE
	unifiedAgentConfigurationServiceConfigurationSourcesMULTILINEParserRepresentation = acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update,
		acctest.GetUpdatedRepresentationCopy("service_configuration", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("sources", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("parser", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
			"parser_type":      acctest.Representation{RepType: acctest.Required, Update: `MULTILINE`},
			"format":           acctest.Representation{RepType: acctest.Optional, Update: []string{`format2`}},
			"format_firstline": acctest.Representation{RepType: acctest.Optional, Update: `formatFirstline2`},
		})},
			unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
			unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	unifiedAgentConfigurationServiceConfigurationSourcesMULTILINEGROKParserRepresentation = acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update,
		acctest.GetUpdatedRepresentationCopy("service_configuration", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("sources", acctest.RepresentationGroup{RepType: acctest.Required, Group: acctest.GetUpdatedRepresentationCopy("parser", acctest.RepresentationGroup{RepType: acctest.Optional, Group: acctest.RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
			"parser_type":             acctest.Representation{RepType: acctest.Required, Update: `MULTILINE_GROK`},
			"grok_failure_key":        acctest.Representation{RepType: acctest.Optional, Update: `grokFailureKey2`},
			"grok_name_key":           acctest.Representation{RepType: acctest.Optional, Update: `grokNameKey2`},
			"multi_line_start_regexp": acctest.Representation{RepType: acctest.Optional, Create: `multiLineStartRegexp`, Update: `multiLineStartRegexp2`},
			"patterns":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: unifiedAgentConfigurationServiceConfigurationSourcesParserPatternsRepresentation},
		})},
			unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
			unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	unifiedAgentConfigurationServiceConfigurationSourcesParserPatternsRepresentation = map[string]interface{}{
		"field_time_format": acctest.Representation{RepType: acctest.Optional, Update: `fieldTimeFormat2`},
		"field_time_key":    acctest.Representation{RepType: acctest.Optional, Update: `fieldTimeKey2`},
		"field_time_zone":   acctest.Representation{RepType: acctest.Optional, Update: `fieldTimeZone2`},
		"name":              acctest.Representation{RepType: acctest.Optional, Update: `name2`},
		"pattern":           acctest.Representation{RepType: acctest.Optional, Update: `pattern2`},
	}
)

// issue-routing-tag: logging/default
func TestLoggingUnifiedAgentConfigurationLogTailResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingUnifiedAgentConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_logging_unified_agent_configuration.test_unified_agent_configuration"
	datasourceName := "data.oci_logging_unified_agent_configurations.test_unified_agent_configurations"
	singularDatasourceName := "data.oci_logging_unified_agent_configuration.test_unified_agent_configuration"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoggingUnifiedAgentConfigurationDestroy,
		Steps: []resource.TestStep{
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Create, unifiedAgentConfigurationLogTailRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//Uncomment configuration_state once bug is fixed
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
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.field_time_key", "fieldTimeKey"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_estimate_current_event", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_keep_time_key", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_null_empty_string", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.null_value_pattern", "nullValuePattern"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "AUDITD"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.timeout_in_milliseconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.types.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.paths.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),

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

			//verify Update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + UnifiedAgentConfigurationResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(unifiedAgentConfigurationLogTailRepresentation, map[string]interface{}{
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
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.field_time_key", "fieldTimeKey"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_estimate_current_event", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_keep_time_key", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_null_empty_string", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.null_value_pattern", "nullValuePattern"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "AUDITD"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.timeout_in_milliseconds", "10"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.types.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.paths.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies + unifiedAgentConfigurationServiceConfigurationSourcesJSONParserRepresentation,
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
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.paths.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.timeout_in_milliseconds", "11"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.types.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.paths.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.time_format", "timeFormat2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.time_type", "UNIXTIME"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_estimate_current_event", "true"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_keep_time_key", "true"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "JSON"),

					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_null_empty_string", "true"),
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

			// verify updates to parser type CSV
			{
				Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies + unifiedAgentConfigurationServiceConfigurationSourcesCSVParserRepresentation,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//Uncomment configuration_state once bug fixed
					//resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.keys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.delimiter", "delimiter"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "CSV"),
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

			// verify updates to parser type GROK
			{
				Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies + unifiedAgentConfigurationServiceConfigurationSourcesGROKParserRepresentation,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.grok_failure_key", "grokFailureKey2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.grok_name_key", "grokNameKey2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "GROK"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.0.field_time_format", "fieldTimeFormat2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.0.field_time_key", "fieldTimeKey2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.0.field_time_zone", "fieldTimeZone2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.0.name", "name2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.0.pattern", "pattern2"),
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

			//verify updates to parser type MSGPACK
			{
				Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies + unifiedAgentConfigurationServiceConfigurationSourcesMSGPACKParserRepresentation,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//Uncomment configuration_state once bug fixed
					//resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.field_time_key", "fieldTimeKey3"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_estimate_current_event", "true"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_keep_time_key", "true"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.is_null_empty_string", "true"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.null_value_pattern", "nullValuePattern3"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "MSGPACK"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.types.%", "1"),
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

			// verify updates to parser type MULTILINE
			{
				Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies + unifiedAgentConfigurationServiceConfigurationSourcesMULTILINEParserRepresentation,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//Uncomment configuration_state once bug fixed
					//resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.format.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.format_firstline", "formatFirstline2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "MULTILINE"),
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

			// verify updates to parser type MULTILINEGROK
			{
				Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies + unifiedAgentConfigurationServiceConfigurationSourcesMULTILINEGROKParserRepresentation,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.grok_failure_key", "grokFailureKey2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.grok_name_key", "grokNameKey2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.0.field_time_format", "fieldTimeFormat2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.0.field_time_key", "fieldTimeKey2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.0.field_time_zone", "fieldTimeZone2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.0.name", "name2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.patterns.0.pattern", "pattern2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.multi_line_start_regexp", "multiLineStartRegexp2"),
					resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.parser.0.parser_type", "MULTILINE_GROK"),
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

			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_logging_unified_agent_configurations", "test_unified_agent_configurations", acctest.Optional, acctest.Update, unifiedAgentConfigurationLogTailDataSourceRepresentation) +
					compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update, unifiedAgentConfigurationLogTailRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "group_id"),
					resource.TestCheckResourceAttr(datasourceName, "is_compartment_id_in_subtree", "false"),
					resource.TestCheckResourceAttrSet(datasourceName, "log_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(datasourceName, "unified_agent_configuration_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "unified_agent_configuration_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Required, acctest.Create, unifiedAgentConfigurationLogTailSingularDataSourceRepresentation) +
					compartmentIdVariableStr + UnifiedAgentConfigurationLogTailResourceConfig,
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
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.name", "name2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.parser.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.parser.0.field_time_key", "fieldTimeKey"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.parser.0.is_estimate_current_event", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.parser.0.is_keep_time_key", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.parser.0.is_null_empty_string", "false"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.parser.0.timeout_in_milliseconds", "10"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.parser.0.types.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.paths.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.source_type", "LOG_TAIL"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationLogTailResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
