// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	UnifiedAgentConfigurationLogTailRequiredOnlyResource = UnifiedAgentConfigurationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Required, Create, unifiedAgentConfigurationLogTailRepresentation)

	UnifiedAgentConfigurationLogTailResourceConfig = UnifiedAgentConfigurationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update, unifiedAgentConfigurationLogTailRepresentation)

	unifiedAgentConfigurationLogTailSingularDataSourceRepresentation = map[string]interface{}{
		"unified_agent_configuration_id": Representation{RepType: Required, Create: `${oci_logging_unified_agent_configuration.test_unified_agent_configuration.id}`},
	}

	unifiedAgentConfigurationLogTailDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":                 Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"group_id":                     Representation{RepType: Optional, Create: `${oci_identity_group.test_group.id}`},
		"is_compartment_id_in_subtree": Representation{RepType: Optional, Create: `false`},
		"log_id":                       Representation{RepType: Optional, Create: `${oci_logging_log.test_log.id}`},
		"state":                        Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":                       RepresentationGroup{Required, unifiedAgentConfigurationLogTailDataSourceFilterRepresentation}}
	unifiedAgentConfigurationLogTailDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_logging_unified_agent_configuration.test_unified_agent_configuration.id}`}},
	}

	unifiedAgentConfigurationLogTailRepresentation = map[string]interface{}{
		"compartment_id":        Representation{RepType: Required, Create: `${var.compartment_id}`},
		"is_enabled":            Representation{RepType: Required, Create: `true`, Update: `false`},
		"service_configuration": RepresentationGroup{Required, unifiedAgentConfigurationServiceConfigurationLogTailRepresentation},
		"defined_tags":          Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":           Representation{RepType: Required, Create: `description`, Update: `description2`},
		"display_name":          Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":         Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"group_association":     RepresentationGroup{Required, unifiedAgentConfigurationGroupAssociationRepresentation},
	}
	unifiedAgentConfigurationServiceConfigurationLogTailRepresentation = map[string]interface{}{
		"configuration_type": Representation{RepType: Required, Create: `LOGGING`},
		"destination":        RepresentationGroup{Required, unifiedAgentConfigurationServiceConfigurationDestinationRepresentation},
		"sources":            RepresentationGroup{Required, unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation},
	}

	unifiedAgentConfigurationServiceConfigurationDestinationRepresentation = map[string]interface{}{
		"log_object_id": Representation{RepType: Required, Create: `${oci_logging_log.test_log.id}`},
	}
	unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation = map[string]interface{}{
		"source_type": Representation{RepType: Required, Create: `LOG_TAIL`},
		"name":        Representation{RepType: Required, Create: `name`, Update: `name2`},
		"parser":      RepresentationGroup{Required, unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation},
		"paths":       Representation{RepType: Optional, Create: []string{`paths`}},
	}

	unifiedAgentConfigurationServiceConfigurationSourcesJSONParserRepresentation = GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update,
		GetUpdatedRepresentationCopy("service_configuration", RepresentationGroup{Required,
			GetUpdatedRepresentationCopy("sources", RepresentationGroup{Required,
				GetUpdatedRepresentationCopy("parser", RepresentationGroup{Optional,
					RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
						"parser_type":               Representation{RepType: Required, Update: `JSON`},
						"field_time_key":            Representation{RepType: Optional, Update: `fieldTimeKey2`},
						"is_estimate_current_event": Representation{RepType: Optional, Update: `true`},
						"is_keep_time_key":          Representation{RepType: Optional, Update: `true`},
						"is_null_empty_string":      Representation{RepType: Optional, Update: `true`},
						"null_value_pattern":        Representation{RepType: Optional, Update: `nullValuePattern2`},
						"timeout_in_milliseconds":   Representation{RepType: Optional, Create: `10`, Update: `11`},
						"types":                     Representation{RepType: Optional, Create: map[string]string{"types": "types"}},
						"time_format":               Representation{RepType: Optional, Update: `timeFormat2`},
						"time_type":                 Representation{RepType: Optional, Update: `UNIXTIME`}})},
					unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
				unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	unifiedAgentConfigurationServiceConfigurationSourcesCSVParserRepresentation = GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update,
		GetUpdatedRepresentationCopy("service_configuration", RepresentationGroup{Required,
			GetUpdatedRepresentationCopy("sources", RepresentationGroup{Required,
				GetUpdatedRepresentationCopy("parser", RepresentationGroup{Optional,
					RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
						"parser_type": Representation{RepType: Required, Update: `CSV`},
						"delimiter":   Representation{RepType: Optional, Update: `delimiter`},
						"keys":        Representation{RepType: Optional, Update: []string{`key1`}}})},
					unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
				unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	unifiedAgentConfigurationServiceConfigurationSourcesGROKParserRepresentation = GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update,
		GetUpdatedRepresentationCopy("service_configuration", RepresentationGroup{Required,
			GetUpdatedRepresentationCopy("sources", RepresentationGroup{Required,
				GetUpdatedRepresentationCopy("parser", RepresentationGroup{Optional,
					RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
						"parser_type":      Representation{RepType: Required, Update: `GROK`},
						"grok_failure_key": Representation{RepType: Optional, Update: `grokFailureKey2`},
						"grok_name_key":    Representation{RepType: Optional, Update: `grokNameKey2`},
						"patterns":         RepresentationGroup{Optional, unifiedAgentConfigurationServiceConfigurationSourcesParserPatternsRepresentation},
					})},
					unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
				unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	unifiedAgentConfigurationServiceConfigurationSourcesMSGPACKParserRepresentation = GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update,
		GetUpdatedRepresentationCopy("service_configuration", RepresentationGroup{Required,
			GetUpdatedRepresentationCopy("sources", RepresentationGroup{Required,
				GetUpdatedRepresentationCopy("parser", RepresentationGroup{Optional,
					RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
						"parser_type":               Representation{RepType: Required, Update: `MSGPACK`},
						"field_time_key":            Representation{RepType: Optional, Update: `fieldTimeKey3`},
						"is_estimate_current_event": Representation{RepType: Optional, Update: `true`},
						"is_keep_time_key":          Representation{RepType: Optional, Update: `true`},
						"is_null_empty_string":      Representation{RepType: Optional, Update: `true`},
						"null_value_pattern":        Representation{RepType: Optional, Update: `nullValuePattern3`},
						"timeout_in_milliseconds":   Representation{RepType: Optional, Update: `11`},
						"types":                     Representation{RepType: Optional, Update: map[string]string{"types": "types2"}},
					})},
					unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
				unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	//MULTILINE
	unifiedAgentConfigurationServiceConfigurationSourcesMULTILINEParserRepresentation = GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update,
		GetUpdatedRepresentationCopy("service_configuration", RepresentationGroup{Required,
			GetUpdatedRepresentationCopy("sources", RepresentationGroup{Required,
				GetUpdatedRepresentationCopy("parser", RepresentationGroup{Optional,
					RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
						"parser_type":      Representation{RepType: Required, Update: `MULTILINE`},
						"format":           Representation{RepType: Optional, Update: []string{`format2`}},
						"format_firstline": Representation{RepType: Optional, Update: `formatFirstline2`},
					})},
					unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
				unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	unifiedAgentConfigurationServiceConfigurationSourcesMULTILINEGROKParserRepresentation = GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update,
		GetUpdatedRepresentationCopy("service_configuration", RepresentationGroup{Required,
			GetUpdatedRepresentationCopy("sources", RepresentationGroup{Required,
				GetUpdatedRepresentationCopy("parser", RepresentationGroup{Optional,
					RepresentationCopyWithNewProperties(unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation, map[string]interface{}{
						"parser_type":             Representation{RepType: Required, Update: `MULTILINE_GROK`},
						"grok_failure_key":        Representation{RepType: Optional, Update: `grokFailureKey2`},
						"grok_name_key":           Representation{RepType: Optional, Update: `grokNameKey2`},
						"multi_line_start_regexp": Representation{RepType: Optional, Create: `multiLineStartRegexp`, Update: `multiLineStartRegexp2`},
						"patterns":                RepresentationGroup{Optional, unifiedAgentConfigurationServiceConfigurationSourcesParserPatternsRepresentation},
					})},
					unifiedAgentConfigurationServiceConfigurationSourcesLogTailRepresentation)},
				unifiedAgentConfigurationServiceConfigurationLogTailRepresentation)},
			unifiedAgentConfigurationLogTailRepresentation))

	unifiedAgentConfigurationServiceConfigurationSourcesParserPatternsRepresentation = map[string]interface{}{
		"field_time_format": Representation{RepType: Optional, Update: `fieldTimeFormat2`},
		"field_time_key":    Representation{RepType: Optional, Update: `fieldTimeKey2`},
		"field_time_zone":   Representation{RepType: Optional, Update: `fieldTimeZone2`},
		"name":              Representation{RepType: Optional, Update: `name2`},
		"pattern":           Representation{RepType: Optional, Update: `pattern2`},
	}
)

// issue-routing-tag: logging/default
func TestLoggingUnifiedAgentConfigurationLogTailResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingUnifiedAgentConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_logging_unified_agent_configuration.test_unified_agent_configuration"
	datasourceName := "data.oci_logging_unified_agent_configurations.test_unified_agent_configurations"
	singularDatasourceName := "data.oci_logging_unified_agent_configuration.test_unified_agent_configuration"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoggingUnifiedAgentConfigurationDestroy,
		Steps: []resource.TestStep{
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Create, unifiedAgentConfigurationLogTailRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					//Uncomment configuration_state once bug is fixed
					resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId, err = FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
							if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
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
					GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Create,
						RepresentationCopyWithNewProperties(unifiedAgentConfigurationLogTailRepresentation, map[string]interface{}{
							"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = FromInstanceState(s, resourceName, "id")
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
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
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
						resId2, err = FromInstanceState(s, resourceName, "id")
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
				Check: ComposeAggregateTestCheckFuncWrapper(
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
						resId2, err = FromInstanceState(s, resourceName, "id")
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
				Check: ComposeAggregateTestCheckFuncWrapper(
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
						resId2, err = FromInstanceState(s, resourceName, "id")
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
				Check: ComposeAggregateTestCheckFuncWrapper(
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
						resId2, err = FromInstanceState(s, resourceName, "id")
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
				Check: ComposeAggregateTestCheckFuncWrapper(
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
						resId2, err = FromInstanceState(s, resourceName, "id")
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
				Check: ComposeAggregateTestCheckFuncWrapper(
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
						resId2, err = FromInstanceState(s, resourceName, "id")
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
					GenerateDataSourceFromRepresentationMap("oci_logging_unified_agent_configurations", "test_unified_agent_configurations", Optional, Update, unifiedAgentConfigurationLogTailDataSourceRepresentation) +
					compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update, unifiedAgentConfigurationLogTailRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					GenerateDataSourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Required, Create, unifiedAgentConfigurationLogTailSingularDataSourceRepresentation) +
					compartmentIdVariableStr + UnifiedAgentConfigurationLogTailResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "unified_agent_configuration_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "configuration_state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
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
