// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_logging "github.com/oracle/oci-go-sdk/v53/logging"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	UnifiedAgentConfigurationRequiredOnlyResource = UnifiedAgentConfigurationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Required, Create, unifiedAgentConfigurationRepresentation)

	UnifiedAgentConfigurationResourceConfig = UnifiedAgentConfigurationResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update, unifiedAgentConfigurationRepresentation)

	unifiedAgentConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"unified_agent_configuration_id": Representation{RepType: Required, Create: `${oci_logging_unified_agent_configuration.test_unified_agent_configuration.id}`},
	}

	unifiedAgentConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":               Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":                 Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"group_id":                     Representation{RepType: Optional, Create: `${oci_identity_group.test_group.id}`},
		"is_compartment_id_in_subtree": Representation{RepType: Optional, Create: `false`},
		"log_id":                       Representation{RepType: Optional, Create: `${oci_logging_log.test_log.id}`},
		"state":                        Representation{RepType: Optional, Create: `AVAILABLE`},
		"filter":                       RepresentationGroup{Required, unifiedAgentConfigurationDataSourceFilterRepresentation},
	}

	unifiedAgentConfigurationRepresentation = map[string]interface{}{
		"compartment_id":        Representation{RepType: Required, Create: `${var.compartment_id}`},
		"is_enabled":            Representation{RepType: Required, Create: `true`, Update: `false`},
		"service_configuration": RepresentationGroup{Required, unifiedAgentConfigurationServiceConfigurationRepresentation},
		"defined_tags":          Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":           Representation{RepType: Required, Create: `description`, Update: `description2`},
		"display_name":          Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":         Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"group_association":     RepresentationGroup{Required, unifiedAgentConfigurationGroupAssociationRepresentation},
	}

	unifiedAgentConfigurationServiceConfigurationRepresentation = map[string]interface{}{
		"configuration_type": Representation{RepType: Required, Create: `LOGGING`},
		"destination":        RepresentationGroup{Required, unifiedAgentConfigurationServiceConfigurationDestinationRepresentation},
		"sources":            RepresentationGroup{Required, unifiedAgentConfigurationServiceConfigurationSourcesRepresentation},
	}

	unifiedAgentConfigurationGroupAssociationRepresentation = map[string]interface{}{
		"group_list": Representation{RepType: Required, Create: []string{`${oci_identity_group.test_group.id}`}}, // Update: []string{`${oci_identity_group.test_group.id}`, `ocid1.Group.oc1..aaaaaaaa5rvs7zjwdk3zdmysm7x7wcxyanbllutswe4xbl7ng4stohtg3sla`}},
	}

	unifiedAgentConfigurationServiceConfigurationSourcesParserRepresentation = map[string]interface{}{
		"parser_type":               Representation{RepType: Required, Create: `AUDITD`},
		"field_time_key":            Representation{RepType: Optional, Create: `fieldTimeKey`},
		"is_estimate_current_event": Representation{RepType: Optional, Create: `false`},
		"is_keep_time_key":          Representation{RepType: Optional, Create: `false`},
		"is_null_empty_string":      Representation{RepType: Optional, Create: `false`},
		"null_value_pattern":        Representation{RepType: Optional, Create: `nullValuePattern`},
		"timeout_in_milliseconds":   Representation{RepType: Optional, Create: `10`},
		"types":                     Representation{RepType: Optional, Create: map[string]string{"types": "types"}},
	}

	unifiedAgentConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_logging_unified_agent_configuration.test_unified_agent_configuration.id}`}},
	}

	UnifiedAgentConfigurationResourceDependencies = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", Required, Create,
			GetUpdatedRepresentationCopy("name", Representation{RepType: Required, Create: `LoggingAgentIdentityGroup`}, groupRepresentation)) +
		GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", Required, Create, logGroupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", Required, Create, customLogRepresentation) +
		GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Required, Create, objectRepresentation)

	unifiedAgentConfigurationServiceConfigurationSourcesRepresentation = map[string]interface{}{
		"source_type": Representation{RepType: Required, Create: `WINDOWS_EVENT_LOG`},
		"channels":    Representation{RepType: Required, Create: []string{`Security`}, Update: []string{`Security`, `Application`}},
		"name":        Representation{RepType: Required, Create: `name`, Update: `name2`},
	}
)

// issue-routing-tag: logging/default
func TestLoggingUnifiedAgentConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingUnifiedAgentConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_logging_unified_agent_configuration.test_unified_agent_configuration"
	datasourceName := "data.oci_logging_unified_agent_configurations.test_unified_agent_configurations"
	singularDatasourceName := "data.oci_logging_unified_agent_configuration.test_unified_agent_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+UnifiedAgentConfigurationResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Create, unifiedAgentConfigurationRepresentation), "logging", "unifiedAgentConfiguration", t)

	ResourceTest(t, testAccCheckLoggingUnifiedAgentConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Required, Create, unifiedAgentConfigurationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "LOGGING"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Create, unifiedAgentConfigurationRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.channels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "WINDOWS_EVENT_LOG"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + UnifiedAgentConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Create,
					RepresentationCopyWithNewProperties(unifiedAgentConfigurationRepresentation, map[string]interface{}{
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
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.channels.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "WINDOWS_EVENT_LOG"),
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
			Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update, unifiedAgentConfigurationRepresentation),
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
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.channels.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.sources.0.source_type", "WINDOWS_EVENT_LOG"),
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
				GenerateDataSourceFromRepresentationMap("oci_logging_unified_agent_configurations", "test_unified_agent_configurations", Optional, Update, unifiedAgentConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + UnifiedAgentConfigurationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Optional, Update, unifiedAgentConfigurationRepresentation),
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
				GenerateDataSourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", Required, Create, unifiedAgentConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + UnifiedAgentConfigurationResourceConfig,
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
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.channels.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_configuration.0.sources.0.source_type", "WINDOWS_EVENT_LOG"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + UnifiedAgentConfigurationResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLoggingUnifiedAgentConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loggingManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_logging_unified_agent_configuration" {
			noResourceFound = false
			request := oci_logging.GetUnifiedAgentConfigurationRequest{}

			tmp := rs.Primary.ID
			request.UnifiedAgentConfigurationId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "logging")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("LoggingUnifiedAgentConfiguration") {
		resource.AddTestSweepers("LoggingUnifiedAgentConfiguration", &resource.Sweeper{
			Name:         "LoggingUnifiedAgentConfiguration",
			Dependencies: DependencyGraph["unifiedAgentConfiguration"],
			F:            sweepLoggingUnifiedAgentConfigurationResource,
		})
	}
}

func sweepLoggingUnifiedAgentConfigurationResource(compartment string) error {
	loggingManagementClient := GetTestClients(&schema.ResourceData{}).loggingManagementClient()
	unifiedAgentConfigurationIds, err := getUnifiedAgentConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, unifiedAgentConfigurationId := range unifiedAgentConfigurationIds {
		if ok := SweeperDefaultResourceId[unifiedAgentConfigurationId]; !ok {
			deleteUnifiedAgentConfigurationRequest := oci_logging.DeleteUnifiedAgentConfigurationRequest{}

			deleteUnifiedAgentConfigurationRequest.UnifiedAgentConfigurationId = &unifiedAgentConfigurationId

			deleteUnifiedAgentConfigurationRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "logging")
			_, error := loggingManagementClient.DeleteUnifiedAgentConfiguration(context.Background(), deleteUnifiedAgentConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting UnifiedAgentConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", unifiedAgentConfigurationId, error)
				continue
			}
		}
	}
	return nil
}

func getUnifiedAgentConfigurationIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "UnifiedAgentConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loggingManagementClient := GetTestClients(&schema.ResourceData{}).loggingManagementClient()

	listUnifiedAgentConfigurationsRequest := oci_logging.ListUnifiedAgentConfigurationsRequest{}
	listUnifiedAgentConfigurationsRequest.CompartmentId = &compartmentId
	listUnifiedAgentConfigurationsResponse, err := loggingManagementClient.ListUnifiedAgentConfigurations(context.Background(), listUnifiedAgentConfigurationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting UnifiedAgentConfiguration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, unifiedAgentConfiguration := range listUnifiedAgentConfigurationsResponse.Items {
		id := *unifiedAgentConfiguration.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "UnifiedAgentConfigurationId", id)
	}
	return resourceIds, nil
}
