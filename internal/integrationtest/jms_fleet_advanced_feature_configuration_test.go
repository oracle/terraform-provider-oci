// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// before running tests, ensure to set up environment variables used below
	JmsFleetAdvancedFeatureConfigurationCompartmentId  = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	JmsFleetAdvancedFeatureConfigurationLogGroupId     = utils.GetEnvSettingWithBlankDefault("fleet_log_group_ocid")
	JmsFleetAdvancedFeatureConfigurationInventoryLogId = utils.GetEnvSettingWithBlankDefault("fleet_inventory_log_ocid")
	JmsFleetAdvancedFeatureConfigurationOperationLogId = utils.GetEnvSettingWithBlankDefault("fleet_operation_log_ocid")

	JmsFleetAdvancedFeatureConfigurationResourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: JmsFleetAdvancedFeatureConfigurationCompartmentId},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `Created Fleet for Advanced Feature Configuration`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Created Fleet for Advanced Feature Configuration`},
		"inventory_log": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{
			"log_group_id": acctest.Representation{
				RepType: acctest.Required,
				Create:  JmsFleetAdvancedFeatureConfigurationLogGroupId,
				Update:  JmsFleetAdvancedFeatureConfigurationLogGroupId,
			},
			"log_id": acctest.Representation{
				RepType: acctest.Required,
				Create:  JmsFleetAdvancedFeatureConfigurationInventoryLogId,
				Update:  JmsFleetAdvancedFeatureConfigurationInventoryLogId,
			},
		}},
		"operation_log": acctest.RepresentationGroup{RepType: acctest.Optional, Group: map[string]interface{}{
			"log_group_id": acctest.Representation{
				RepType: acctest.Required,
				Create:  JmsFleetAdvancedFeatureConfigurationLogGroupId,
				Update:  JmsFleetAdvancedFeatureConfigurationLogGroupId,
			},
			"log_id": acctest.Representation{
				RepType: acctest.Required,
				Create:  JmsFleetAdvancedFeatureConfigurationOperationLogId,
				Update:  JmsFleetAdvancedFeatureConfigurationOperationLogId,
			},
		}},
	}

	JmsFleetAdvancedFeatureConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetAdvancedFeatureConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetAdvancedFeatureConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	singularDatasourceName := "data.oci_jms_fleet_advanced_feature_configuration.test_fleet_advanced_feature_configuration"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_fleet",
					"test_fleet",
					acctest.Required,
					acctest.Create,
					JmsFleetAdvancedFeatureConfigurationResourceRepresentation,
				) +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_advanced_feature_configuration",
					"test_fleet_advanced_feature_configuration",
					acctest.Optional,
					acctest.Create,
					JmsFleetAdvancedFeatureConfigurationSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "advanced_usage_tracking.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "advanced_usage_tracking.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "crypto_event_analysis.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "crypto_event_analysis.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jfr_recording.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jfr_recording.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.is_enabled", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lcm.0.post_installation_actions.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lcm.0.post_installation_actions.0.disabled_tls_versions.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.certpath.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.jar.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.tls.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "lcm.0.post_installation_actions.0.should_replace_certificates_operating_system"),
				resource.TestCheckResourceAttr(singularDatasourceName, "java_migration_analysis.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "performance_tuning_analysis.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
			),
		},
	})
}

// clean up Fleet resource after test
func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("JmsFleetAdvancedFeatureConfiguration") {
		resource.AddTestSweepers("JmsFleetAdvancedFeatureConfiguration", &resource.Sweeper{
			Name:         "JmsFleetAdvancedFeatureConfiguration",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepJmsFleetResource,
		})
	}
}
