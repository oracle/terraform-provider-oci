// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	JmsFleetAdvancedFeatureConfigurationRequiredOnlyResource = JmsFleetAdvancedFeatureConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet_advanced_feature_configuration", "test_fleet_advanced_feature_configuration", acctest.Required, acctest.Create, JmsFleetAdvancedFeatureConfigurationRepresentation)

	JmsFleetAdvancedFeatureConfigurationResourceConfig = JmsFleetAdvancedFeatureConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet_advanced_feature_configuration", "test_fleet_advanced_feature_configuration", acctest.Optional, acctest.Update, JmsFleetAdvancedFeatureConfigurationRepresentation)

	// before running tests, ensure to set up environment variables used below
	JmsFleetAdvancedFeatureConfigurationCompartmentId  = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	JmsFleetAdvancedFeatureConfigurationLogGroupId     = utils.GetEnvSettingWithBlankDefault("fleet_log_group_ocid")
	JmsFleetAdvancedFeatureConfigurationInventoryLogId = utils.GetEnvSettingWithBlankDefault("fleet_inventory_log_ocid")
	JmsFleetAdvancedFeatureConfigurationOperationLogId = utils.GetEnvSettingWithBlankDefault("fleet_operation_log_ocid")

	JmsFleetAdvancedFeatureConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
	}

	JmsFleetAdvancedFeatureConfigurationRepresentation = map[string]interface{}{
		"fleet_id":                    acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
		"advanced_usage_tracking":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationAdvancedUsageTrackingRepresentation},
		"analytic_bucket_name":        acctest.Representation{RepType: acctest.Optional, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"analytic_namespace":          acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"crypto_event_analysis":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationCryptoEventAnalysisRepresentation},
		"java_migration_analysis":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationJavaMigrationAnalysisRepresentation},
		"jfr_recording":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationJfrRecordingRepresentation},
		"lcm":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationLcmRepresentation},
		"performance_tuning_analysis": acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationPerformanceTuningAnalysisRepresentation},
	}
	JmsFleetAdvancedFeatureConfigurationAdvancedUsageTrackingRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	JmsFleetAdvancedFeatureConfigurationCryptoEventAnalysisRepresentation = map[string]interface{}{
		"is_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"summarized_events_log": acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationCryptoEventAnalysisSummarizedEventsLogRepresentation},
	}
	JmsFleetAdvancedFeatureConfigurationJavaMigrationAnalysisRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	JmsFleetAdvancedFeatureConfigurationJfrRecordingRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	JmsFleetAdvancedFeatureConfigurationLcmRepresentation = map[string]interface{}{
		"is_enabled":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"post_installation_actions": acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsRepresentation},
	}
	JmsFleetAdvancedFeatureConfigurationPerformanceTuningAnalysisRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	JmsFleetAdvancedFeatureConfigurationCryptoEventAnalysisSummarizedEventsLogRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log_group.test_log_group.id}`},
		"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_logging_log.test_log.id}`},
	}
	JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsRepresentation = map[string]interface{}{
		"add_logging_handler":                          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"disabled_tls_versions":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`TLS_1_0`}, Update: []string{`TLS_1_1`}},
		"global_logging_level":                         acctest.Representation{RepType: acctest.Optional, Create: `ALL`, Update: `SEVERE`},
		"minimum_key_size_settings":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsMinimumKeySizeSettingsRepresentation},
		"proxies":                                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsProxiesRepresentation},
		"should_replace_certificates_operating_system": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsMinimumKeySizeSettingsRepresentation = map[string]interface{}{
		"certpath": acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsMinimumKeySizeSettingsCertpathRepresentation},
		"jar":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsMinimumKeySizeSettingsJarRepresentation},
		"tls":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsMinimumKeySizeSettingsTlsRepresentation},
	}
	JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsProxiesRepresentation = map[string]interface{}{
		"ftp_proxy_host":     acctest.Representation{RepType: acctest.Optional, Create: `ftpProxyHost`, Update: `ftpProxyHost2`},
		"ftp_proxy_port":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"http_proxy_host":    acctest.Representation{RepType: acctest.Optional, Create: `httpProxyHost`, Update: `httpProxyHost2`},
		"http_proxy_port":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"https_proxy_host":   acctest.Representation{RepType: acctest.Optional, Create: `httpsProxyHost`, Update: `httpsProxyHost2`},
		"https_proxy_port":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"socks_proxy_host":   acctest.Representation{RepType: acctest.Optional, Create: `socksProxyHost`, Update: `socksProxyHost2`},
		"socks_proxy_port":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"use_system_proxies": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsMinimumKeySizeSettingsCertpathRepresentation = map[string]interface{}{
		"key_size": acctest.Representation{RepType: acctest.Optional, Create: `2048`, Update: `2048`},
		"name":     acctest.Representation{RepType: acctest.Optional, Create: `RSA`, Update: `DSA`},
	}
	JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsMinimumKeySizeSettingsJarRepresentation = map[string]interface{}{
		"key_size": acctest.Representation{RepType: acctest.Optional, Create: `2048`, Update: `2048`},
		"name":     acctest.Representation{RepType: acctest.Optional, Create: `RSA`, Update: `DSA`},
	}
	JmsFleetAdvancedFeatureConfigurationLcmPostInstallationActionsMinimumKeySizeSettingsTlsRepresentation = map[string]interface{}{
		"key_size": acctest.Representation{RepType: acctest.Optional, Create: `2048`, Update: `2048`},
		"name":     acctest.Representation{RepType: acctest.Optional, Create: `RSA`, Update: `DSA`},
	}
	JmsFleetAdvancedFeatureConfigurationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Required, acctest.Create, JmsFleetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_log_group", acctest.Required, acctest.Create, LoggingLogGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_logging_log", "test_log", acctest.Required, acctest.Create, LoggingLogRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: jms/default
func TestJmsFleetAdvancedFeatureConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetAdvancedFeatureConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_jms_fleet_advanced_feature_configuration.test_fleet_advanced_feature_configuration"

	singularDatasourceName := "data.oci_jms_fleet_advanced_feature_configuration.test_fleet_advanced_feature_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+JmsFleetAdvancedFeatureConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet_advanced_feature_configuration", "test_fleet_advanced_feature_configuration", acctest.Optional, acctest.Create, JmsFleetAdvancedFeatureConfigurationRepresentation), "jms", "fleetAdvancedFeatureConfiguration", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + JmsFleetAdvancedFeatureConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet_advanced_feature_configuration", "test_fleet_advanced_feature_configuration", acctest.Required, acctest.Create, JmsFleetAdvancedFeatureConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + JmsFleetAdvancedFeatureConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + JmsFleetAdvancedFeatureConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet_advanced_feature_configuration", "test_fleet_advanced_feature_configuration", acctest.Optional, acctest.Create, JmsFleetAdvancedFeatureConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_usage_tracking.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "advanced_usage_tracking.0.is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "analytic_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "analytic_namespace"),
				resource.TestCheckResourceAttr(resourceName, "crypto_event_analysis.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "crypto_event_analysis.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "crypto_event_analysis.0.summarized_events_log.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_event_analysis.0.summarized_events_log.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_event_analysis.0.summarized_events_log.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "java_migration_analysis.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "java_migration_analysis.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "jfr_recording.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "jfr_recording.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "lcm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.add_logging_handler", "false"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.disabled_tls_versions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.global_logging_level", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.certpath.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.certpath.0.key_size", "2048"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.certpath.0.name", "RSA"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.jar.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.jar.0.key_size", "2048"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.jar.0.name", "RSA"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.tls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.tls.0.key_size", "2048"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.tls.0.name", "RSA"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.ftp_proxy_host", "ftpProxyHost"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.ftp_proxy_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.http_proxy_host", "httpProxyHost"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.http_proxy_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.https_proxy_host", "httpsProxyHost"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.https_proxy_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.socks_proxy_host", "socksProxyHost"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.socks_proxy_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.use_system_proxies", "false"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.should_replace_certificates_operating_system", "false"),
				resource.TestCheckResourceAttr(resourceName, "performance_tuning_analysis.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "performance_tuning_analysis.0.is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "time_last_modified"),

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
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + JmsFleetAdvancedFeatureConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet_advanced_feature_configuration", "test_fleet_advanced_feature_configuration", acctest.Optional, acctest.Update, JmsFleetAdvancedFeatureConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_usage_tracking.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "advanced_usage_tracking.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "analytic_bucket_name"),
				resource.TestCheckResourceAttrSet(resourceName, "analytic_namespace"),
				resource.TestCheckResourceAttr(resourceName, "crypto_event_analysis.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "crypto_event_analysis.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "crypto_event_analysis.0.summarized_events_log.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_event_analysis.0.summarized_events_log.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "crypto_event_analysis.0.summarized_events_log.0.log_id"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "java_migration_analysis.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "java_migration_analysis.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "jfr_recording.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "jfr_recording.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "lcm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.add_logging_handler", "true"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.disabled_tls_versions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.global_logging_level", "SEVERE"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.certpath.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.certpath.0.key_size", "2048"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.certpath.0.name", "DSA"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.jar.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.jar.0.key_size", "2048"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.jar.0.name", "DSA"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.tls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.tls.0.key_size", "2048"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.tls.0.name", "DSA"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.ftp_proxy_host", "ftpProxyHost2"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.ftp_proxy_port", "11"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.http_proxy_host", "httpProxyHost2"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.http_proxy_port", "11"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.https_proxy_host", "httpsProxyHost2"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.https_proxy_port", "11"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.socks_proxy_host", "socksProxyHost2"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.socks_proxy_port", "11"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.proxies.0.use_system_proxies", "true"),
				resource.TestCheckResourceAttr(resourceName, "lcm.0.post_installation_actions.0.should_replace_certificates_operating_system", "true"),
				resource.TestCheckResourceAttr(resourceName, "performance_tuning_analysis.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "performance_tuning_analysis.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "time_last_modified"),

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
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_fleet_advanced_feature_configuration",
					"test_fleet_advanced_feature_configuration",
					acctest.Required,
					acctest.Create,
					JmsFleetAdvancedFeatureConfigurationRepresentation,
				) +
				compartmentIdVariableStr +
				JmsFleetAdvancedFeatureConfigurationResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_advanced_feature_configuration",
					"test_fleet_advanced_feature_configuration",
					acctest.Optional,
					acctest.Create,
					JmsFleetAdvancedFeatureConfigurationSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "advanced_usage_tracking.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "advanced_usage_tracking.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "analytic_namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "crypto_event_analysis.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "crypto_event_analysis.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "crypto_event_analysis.0.summarized_events_log.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "java_migration_analysis.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "java_migration_analysis.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jfr_recording.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jfr_recording.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.add_logging_handler", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.disabled_tls_versions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.global_logging_level", "SEVERE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.certpath.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.certpath.0.key_size", "2048"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.certpath.0.name", "DSA"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.jar.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.jar.0.key_size", "2048"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.jar.0.name", "DSA"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.tls.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.tls.0.key_size", "2048"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.minimum_key_size_settings.0.tls.0.name", "DSA"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.proxies.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.proxies.0.ftp_proxy_host", "ftpProxyHost2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.proxies.0.ftp_proxy_port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.proxies.0.http_proxy_host", "httpProxyHost2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.proxies.0.http_proxy_port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.proxies.0.https_proxy_host", "httpsProxyHost2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.proxies.0.https_proxy_port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.proxies.0.socks_proxy_host", "socksProxyHost2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.proxies.0.socks_proxy_port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.proxies.0.use_system_proxies", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "lcm.0.post_installation_actions.0.should_replace_certificates_operating_system", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "performance_tuning_analysis.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "performance_tuning_analysis.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
			),
		},
		// verify resource import
		{
			Config:                  config + JmsFleetAdvancedFeatureConfigurationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
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
