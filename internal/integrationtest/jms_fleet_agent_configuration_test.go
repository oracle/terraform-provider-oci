// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	JmsFleetAgentConfigurationRepresentation = map[string]interface{}{
		"fleet_id":                                           acctest.Representation{RepType: acctest.Required, Create: JmsFleetId},
		"agent_polling_interval_in_minutes":                  acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"is_capturing_ip_address_and_fqdn_enabled":           acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_collecting_managed_instance_metrics_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_collecting_usernames_enabled":                    acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_libraries_scan_enabled":                          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"java_usage_tracker_processing_frequency_in_minutes": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"jre_scan_frequency_in_minutes":                      acctest.Representation{RepType: acctest.Optional, Create: `180`, Update: `181`},
		"linux_configuration": acctest.RepresentationGroup{
			RepType: acctest.Optional,
			Group: map[string]interface{}{
				"exclude_paths": acctest.Representation{RepType: acctest.Required, Create: []string{`/user/private1`}, Update: []string{`/user/private2`}},
				"include_paths": acctest.Representation{RepType: acctest.Required, Create: []string{`/user`}, Update: []string{`/user`, `/opt`}},
			}},
		"mac_os_configuration": acctest.RepresentationGroup{
			RepType: acctest.Optional,
			Group: map[string]interface{}{
				"exclude_paths": acctest.Representation{RepType: acctest.Required, Create: []string{`/home/private1`}, Update: []string{`/home/private2`}},
				"include_paths": acctest.Representation{RepType: acctest.Required, Create: []string{`/home`}, Update: []string{`/home`, `/library`}},
			}},
		"windows_configuration": acctest.RepresentationGroup{
			RepType: acctest.Optional,
			Group: map[string]interface{}{
				"exclude_paths": acctest.Representation{RepType: acctest.Required, Create: []string{`c:\\windows\\private1`}, Update: []string{`c:\\windows\\private2`}},
				"include_paths": acctest.Representation{RepType: acctest.Required, Create: []string{`c:\\windows`}, Update: []string{`c:\\windows`, `d:\\data`}},
			}},
		"work_request_validity_period_in_days": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	JmsFleetAgentConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: JmsFleetId},
	}
)

// issue-routing-tag: jms/default
func TestJmsFleetAgentConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetAgentConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	resourceName := "oci_jms_fleet_agent_configuration.test_fleet_agent_configuration"
	singularDatasourceName := "data.oci_jms_fleet_agent_configuration.test_fleet_agent_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties.
	// This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+
		acctest.GenerateResourceFromRepresentationMap(
			"oci_jms_fleet_agent_configuration",
			"test_fleet_agent_configuration",
			acctest.Optional,
			acctest.Create,
			JmsFleetAgentConfigurationRepresentation,
		),
		"jms",
		"fleetAgentConfiguration",
		t,
	)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_fleet_agent_configuration",
					"test_fleet_agent_configuration",
					acctest.Required,
					acctest.Create,
					JmsFleetAgentConfigurationRepresentation,
				),
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
			Config: config,
		},

		// verify Create with optionals
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_fleet_agent_configuration",
					"test_fleet_agent_configuration",
					acctest.Optional,
					acctest.Create,
					JmsFleetAgentConfigurationRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_polling_interval_in_minutes", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_capturing_ip_address_and_fqdn_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_collecting_managed_instance_metrics_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_collecting_usernames_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_libraries_scan_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "java_usage_tracker_processing_frequency_in_minutes", "10"),
				resource.TestCheckResourceAttr(resourceName, "jre_scan_frequency_in_minutes", "180"),
				resource.TestCheckResourceAttr(resourceName, "linux_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "linux_configuration.0.exclude_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "linux_configuration.0.include_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mac_os_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mac_os_configuration.0.exclude_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mac_os_configuration.0.include_paths.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_last_modified"),
				resource.TestCheckResourceAttr(resourceName, "windows_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "windows_configuration.0.exclude_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "windows_configuration.0.include_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "work_request_validity_period_in_days", "10"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &JmsCompartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_fleet_agent_configuration",
					"test_fleet_agent_configuration",
					acctest.Optional,
					acctest.Update,
					JmsFleetAgentConfigurationRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "agent_polling_interval_in_minutes", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "fleet_id"),
				resource.TestCheckResourceAttr(resourceName, "is_capturing_ip_address_and_fqdn_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_collecting_managed_instance_metrics_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_collecting_usernames_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_libraries_scan_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "java_usage_tracker_processing_frequency_in_minutes", "11"),
				resource.TestCheckResourceAttr(resourceName, "jre_scan_frequency_in_minutes", "181"),
				resource.TestCheckResourceAttr(resourceName, "linux_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "linux_configuration.0.exclude_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "linux_configuration.0.include_paths.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "mac_os_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mac_os_configuration.0.exclude_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mac_os_configuration.0.include_paths.#", "2"),
				resource.TestCheckResourceAttrSet(resourceName, "time_last_modified"),
				resource.TestCheckResourceAttr(resourceName, "windows_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "windows_configuration.0.exclude_paths.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "windows_configuration.0.include_paths.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "work_request_validity_period_in_days", "11"),

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
					"oci_jms_fleet_agent_configuration",
					"test_fleet_agent_configuration",
					acctest.Optional,
					acctest.Update,
					JmsFleetAgentConfigurationRepresentation,
				) +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_fleet_agent_configuration",
					"test_fleet_agent_configuration",
					acctest.Required,
					acctest.Create,
					JmsFleetAgentConfigurationSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "agent_polling_interval_in_minutes", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_capturing_ip_address_and_fqdn_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_collecting_managed_instance_metrics_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_collecting_usernames_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_libraries_scan_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "java_usage_tracker_processing_frequency_in_minutes", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "jre_scan_frequency_in_minutes", "181"),
				resource.TestCheckResourceAttr(singularDatasourceName, "linux_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "linux_configuration.0.exclude_paths.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "linux_configuration.0.include_paths.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mac_os_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mac_os_configuration.0.exclude_paths.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mac_os_configuration.0.include_paths.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
				resource.TestCheckResourceAttr(singularDatasourceName, "windows_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "windows_configuration.0.exclude_paths.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "windows_configuration.0.include_paths.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "work_request_validity_period_in_days", "11"),
			),
		},

		// verify resource import
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_jms_fleet_agent_configuration",
					"test_fleet_agent_configuration",
					acctest.Required,
					acctest.Create,
					JmsFleetAgentConfigurationRepresentation,
				),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
