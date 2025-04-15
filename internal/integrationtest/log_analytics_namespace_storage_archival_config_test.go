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
	LogAnalyticsNamespaceStorageArchivalConfigResourceConfig = LogAnalyticsNamespaceStorageArchivalConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_storage_archival_config", "test_namespace_storage_archival_config", acctest.Optional, acctest.Update, LogAnalyticsNamespaceStorageArchivalConfigRepresentation)

	LogAnalyticsNamespaceStorageArchivalConfigSingularDataSourceRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
	}

	LogAnalyticsNamespaceStorageArchivalConfigRepresentation = map[string]interface{}{
		"archiving_configuration": acctest.RepresentationGroup{RepType: acctest.Required, Group: LogAnalyticsNamespaceStorageArchivalConfigArchivingConfigurationRepresentation},
		"namespace":               acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_log_analytics_namespaces.test_namespaces.namespace_collection[0].items[0], "namespace")}`},
	}

	LogAnalyticsNamespaceStorageArchivalConfigArchivingConfigurationRepresentation = map[string]interface{}{
		"active_storage_duration":   acctest.Representation{RepType: acctest.Required, Create: `P45D`, Update: `P90D`},
		"archival_storage_duration": acctest.Representation{RepType: acctest.Required, Create: `P180D`, Update: `-1`},
	}

	LogAnalyticsNamespaceStorageArchivalConfigResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespaces", "test_namespaces", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceStorageArchivalConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceStorageArchivalConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_namespace_storage_archival_config.test_namespace_storage_archival_config"

	singularDatasourceName := "data.oci_log_analytics_namespace_storage_archival_config.test_namespace_storage_archival_config"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsNamespaceStorageArchivalConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_storage_archival_config", "test_namespace_storage_archival_config", acctest.Required, acctest.Create, LogAnalyticsNamespaceStorageArchivalConfigRepresentation), "loganalytics", "namespaceStorageArchivalConfig", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceStorageArchivalConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_storage_archival_config", "test_namespace_storage_archival_config", acctest.Required, acctest.Create, LogAnalyticsNamespaceStorageArchivalConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "archiving_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "archiving_configuration.0.active_storage_duration", "P45D"),
				resource.TestCheckResourceAttr(resourceName, "archiving_configuration.0.archival_storage_duration", "P180D"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

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
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceStorageArchivalConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_namespace_storage_archival_config", "test_namespace_storage_archival_config", acctest.Optional, acctest.Update, LogAnalyticsNamespaceStorageArchivalConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "archiving_configuration.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "archiving_configuration.0.active_storage_duration", "P90D"),
				resource.TestCheckResourceAttr(resourceName, "archiving_configuration.0.archival_storage_duration", "-1"),
				resource.TestCheckResourceAttrSet(resourceName, "is_archiving_enabled"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_storage_archival_config", "test_namespace_storage_archival_config", acctest.Required, acctest.Create, LogAnalyticsNamespaceStorageArchivalConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceStorageArchivalConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),

				resource.TestCheckResourceAttr(singularDatasourceName, "archiving_configuration.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "archiving_configuration.0.active_storage_duration", "P90D"),
				resource.TestCheckResourceAttr(singularDatasourceName, "archiving_configuration.0.archival_storage_duration", "-1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_archiving_enabled"),
			),
		},
		// verify resource import
		{
			Config:                  config + LogAnalyticsNamespaceStorageArchivalConfigResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
