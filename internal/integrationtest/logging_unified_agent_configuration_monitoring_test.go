// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	LoggingUnifiedAgentConfigurationMonitoringRequiredOnlyResource = LoggingUnifiedAgentConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Required, acctest.Create, LoggingUnifiedAgentConfigurationMonitoringRepresentation)

	LoggingUnifiedAgentConfigurationMonitoringResourceConfig = LoggingUnifiedAgentConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update, LoggingUnifiedAgentConfigurationMonitoringRepresentation)

	LoggingLoggingUnifiedAgentConfigurationMonitoringSingularDataSourceRepresentation = map[string]interface{}{
		"unified_agent_configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_unified_agent_configuration.test_unified_agent_configuration.id}`},
	}

	LoggingLoggingUnifiedAgentConfigurationMonitoringDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationMonitoringDataSourceFilterRepresentation},
	}
	LoggingUnifiedAgentConfigurationMonitoringDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_logging_unified_agent_configuration.test_unified_agent_configuration.id}`}},
	}

	LoggingUnifiedAgentConfigurationMonitoringRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_enabled":            acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
		"service_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationMonitoringServiceConfigurationRepresentation},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":           acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"group_association":     acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationMonitoringGroupAssociationRepresentation},
	}

	LoggingUnifiedAgentConfigurationMonitoringServiceConfigurationRepresentation = map[string]interface{}{
		"configuration_type":         acctest.Representation{RepType: acctest.Required, Create: `MONITORING`},
		"application_configurations": acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsRepresentation},
	}

	LoggingUnifiedAgentConfigurationMonitoringGroupAssociationRepresentation = map[string]interface{}{
		"group_list": acctest.Representation{RepType: acctest.Required, Create: []string{`ocid1.group.oc1..aaaaaaaa6qexvubb5vy7olbnuephj4ord3vkzmpljw6ajnipgckr2dyed37a`}},
	}

	LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsRepresentation = map[string]interface{}{
		"destination":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsDestinationRepresentation},
		"source_type":                        acctest.Representation{RepType: acctest.Required, Create: `KUBERNETES`},
		"source":                             acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsSourceRepresentation},
		"unified_agent_configuration_filter": acctest.RepresentationGroup{RepType: acctest.Optional, Group: LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsUnifiedAgentConfigurationFilterRepresentation},
	}
	LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsDestinationRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"metrics_namespace": acctest.Representation{RepType: acctest.Required, Create: `metricsNamespace`, Update: `metricsNamespace2`},
	}
	LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsSourceRepresentation = map[string]interface{}{
		"name":           acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"scrape_targets": acctest.RepresentationGroup{RepType: acctest.Required, Group: LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsSourceScrapeTargetsRepresentation},
	}

	LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsUnifiedAgentConfigurationFilterRepresentation = map[string]interface{}{
		"allow_list":  acctest.Representation{RepType: acctest.Optional, Create: []string{`allowList`}, Update: []string{`allowList2`}},
		"deny_list":   acctest.Representation{RepType: acctest.Optional, Create: []string{`denyList`}, Update: []string{`denyList2`}},
		"filter_type": acctest.Representation{RepType: acctest.Optional, Create: `KUBERNETES_FILTER`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
	}
	LoggingUnifiedAgentConfigurationServiceConfigurationApplicationConfigurationsSourceScrapeTargetsRepresentation = map[string]interface{}{
		"k8s_namespace":  acctest.Representation{RepType: acctest.Required, Create: `k8sNamespace`, Update: `k8sNamespace2`},
		"resource_group": acctest.Representation{RepType: acctest.Optional, Create: `resourceGroup`, Update: `resourceGroup2`},
		"resource_type":  acctest.Representation{RepType: acctest.Required, Create: `PODS`, Update: `ENDPOINTS`},
		"service_name":   acctest.Representation{RepType: acctest.Optional, Create: `serviceName`, Update: `serviceName2`},
	}

	LoggingUnifiedAgentConfigurationMonitoringResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_group", "test_group", acctest.Required, acctest.Create, LoggingUAIdentityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, LoggingLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, customLogRepresentation) // +
)

// issue-routing-tag: logging/default
func TestLoggingUnifiedAgentConfigurationMonitoringResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoggingUnifiedAgentConfigurationMonitoringResource_basic")
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LoggingUnifiedAgentConfigurationMonitoringResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Create, LoggingUnifiedAgentConfigurationMonitoringRepresentation), "logging", "unifiedAgentConfigurationMonitoring", t)

	acctest.ResourceTest(t, testAccCheckLoggingUnifiedAgentConfigurationMonitoringDestroy, []resource.TestStep{

		// Don't change below tests
		// 0. verify Create
		{
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationMonitoringResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Required, acctest.Create, LoggingUnifiedAgentConfigurationMonitoringRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "MONITORING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source_type", "KUBERNETES"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.destination.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.destination.0.metrics_namespace", "metricsNamespace"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.k8s_namespace", "k8sNamespace"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.resource_type", "PODS"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// 1. delete before next Create
		{
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationMonitoringResourceDependencies,
		},
		// 2. verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationMonitoringResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Create, LoggingUnifiedAgentConfigurationMonitoringRepresentation),
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
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "MONITORING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source_type", "KUBERNETES"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.destination.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.destination.0.metrics_namespace", "metricsNamespace"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.k8s_namespace", "k8sNamespace"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.resource_group", "resourceGroup"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.resource_type", "PODS"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.service_name", "serviceName"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.filter_type", "KUBERNETES_FILTER"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.allow_list.0", "allowList"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.deny_list.0", "denyList"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + LoggingUnifiedAgentConfigurationMonitoringResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(LoggingUnifiedAgentConfigurationMonitoringRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
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
			Config: config + compartmentIdVariableStr + LoggingUnifiedAgentConfigurationMonitoringResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update, LoggingUnifiedAgentConfigurationMonitoringRepresentation),
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
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "MONITORING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source_type", "KUBERNETES"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.destination.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.destination.0.metrics_namespace", "metricsNamespace2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.k8s_namespace", "k8sNamespace2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.resource_group", "resourceGroup2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.resource_type", "ENDPOINTS"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.service_name", "serviceName2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.filter_type", "KUBERNETES_FILTER"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.allow_list.0", "allowList2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.deny_list.0", "denyList2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.name", "name2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_logging_unified_agent_configurations", "test_unified_agent_configurations", acctest.Optional, acctest.Update, LoggingLoggingUnifiedAgentConfigurationMonitoringDataSourceRepresentation) +
				compartmentIdVariableStr + LoggingUnifiedAgentConfigurationMonitoringResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Optional, acctest.Update, LoggingUnifiedAgentConfigurationMonitoringRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_logging_unified_agent_configuration", "test_unified_agent_configuration", acctest.Required, acctest.Create, LoggingLoggingUnifiedAgentConfigurationMonitoringSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LoggingUnifiedAgentConfigurationMonitoringResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "unified_agent_configuration_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "configuration_state"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_association.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "group_association.0.group_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.configuration_type", "MONITORING"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source_type", "KUBERNETES"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.destination.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.destination.0.metrics_namespace", "metricsNamespace2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.k8s_namespace", "k8sNamespace2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.resource_group", "resourceGroup2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.resource_type", "ENDPOINTS"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.source.0.scrape_targets.0.service_name", "serviceName2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.filter_type", "KUBERNETES_FILTER"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.allow_list.0", "allowList2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.deny_list.0", "denyList2"),
				resource.TestCheckResourceAttr(resourceName, "service_configuration.0.application_configurations.0.unified_agent_configuration_filter.0.name", "name2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
			),
		},
		// 7. verify resource import
		{
			Config:                  config + LoggingUnifiedAgentConfigurationMonitoringRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckLoggingUnifiedAgentConfigurationMonitoringDestroy(s *terraform.State) error {
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
	if !acctest.InSweeperExcludeList("LoggingUnifiedAgentConfigurationMonitoring") {
		resource.AddTestSweepers("LoggingUnifiedAgentConfigurationMonitoring", &resource.Sweeper{
			Name:         "LoggingUnifiedAgentConfigurationMonitoring",
			Dependencies: acctest.DependencyGraph["unifiedAgentConfigurationMonitoring"],
			F:            sweepLoggingUnifiedAgentConfigurationMonitoringResource,
		})
	}
}

func sweepLoggingUnifiedAgentConfigurationMonitoringResource(compartment string) error {
	loggingManagementClient := acctest.GetTestClients(&schema.ResourceData{}).LoggingManagementClient()
	unifiedAgentConfigurationIds, err := getLoggingUnifiedAgentConfigurationMonitoringIds(compartment)
	if err != nil {
		return err
	}
	for _, unifiedAgentConfigurationId := range unifiedAgentConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[unifiedAgentConfigurationId]; !ok {
			deleteUnifiedAgentConfigurationRequest := oci_logging.DeleteUnifiedAgentConfigurationRequest{}

			deleteUnifiedAgentConfigurationRequest.UnifiedAgentConfigurationId = &unifiedAgentConfigurationId

			deleteUnifiedAgentConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "logging")
			_, err := loggingManagementClient.DeleteUnifiedAgentConfiguration(context.Background(), deleteUnifiedAgentConfigurationRequest)
			if err != nil {
				fmt.Printf("Error deleting unifiedAgentConfigurationMonitoring %s %s, It is possible that the resource is already deleted. Please verify manually \n", unifiedAgentConfigurationId, err)
				continue
			}
		}
	}
	return nil
}

func getLoggingUnifiedAgentConfigurationMonitoringIds(compartment string) ([]string, error) {
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
		return resourceIds, fmt.Errorf("Error getting unifiedAgentConfigurationMonitoring list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, unifiedAgentConfigurationMonitoring := range listUnifiedAgentConfigurationsResponse.Items {
		id := *unifiedAgentConfigurationMonitoring.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "UnifiedAgentConfigurationId", id)
	}
	return resourceIds, nil
}
