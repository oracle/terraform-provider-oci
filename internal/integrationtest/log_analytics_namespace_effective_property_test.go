// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LogAnalyticsLogAnalyticsNamespaceEffectivePropertyDataSourceRepresentation = map[string]interface{}{
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"source_name": acctest.Representation{RepType: acctest.Optional, Create: `LinuxSyslogSource`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `management_agent.database_sql.badsql_retry`},
	}

	LogAnalyticsLogAnalyticsPreferencesManagementCollectionPropertyRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"items":     []acctest.RepresentationGroup{{RepType: acctest.Required, Group: LogAnalyticsLogAnalyticsPreferencesManagementCollectionPropertyItemsRepresentation}},
	}

	LogAnalyticsLogAnalyticsPreferencesManagementCollectionPropertyItemsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `management_agent.database_sql.badsql_retry`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `PT25M`},
	}

	LogAnalyticsNamespaceEffectivePropertyResourceConfig = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceEffectivePropertyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceEffectivePropertyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	preferencesResourceName := "oci_log_analytics_log_analytics_preferences_management.test_log_analytics_preferences_management"

	requiredDatasourceName := "data.oci_log_analytics_namespace_effective_properties.test_effective_properties_required"
	optionalDatasourceName := "data.oci_log_analytics_namespace_effective_properties.test_effective_properties_optional"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Set a tenant level preference for collection property
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsPreferencesManagementCollectionPropertyRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceEffectivePropertyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(preferencesResourceName, "namespace"),
				resource.TestCheckResourceAttr(preferencesResourceName, "items.#", "1"),
			),
		},
		// verify datasource with required parameters
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsPreferencesManagementCollectionPropertyRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_effective_properties", "test_effective_properties_required", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceEffectivePropertyDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceEffectivePropertyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "effective_property_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "effective_property_collection.0.items.0.value"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "effective_property_collection.0.items.0.effective_level"),
			),
		},
		// verify datasource with optional parameters
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsPreferencesManagementCollectionPropertyRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_effective_properties", "test_effective_properties_optional", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsNamespaceEffectivePropertyDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceEffectivePropertyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "effective_property_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "effective_property_collection.0.items.0.name", "management_agent.database_sql.badsql_retry"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "effective_property_collection.0.items.0.value", "PT25M"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "effective_property_collection.0.items.0.effective_level", "TENANT"),
			),
		},
		// Delete the preference for collection property
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsNamespaceEffectivePropertyResourceConfig,
		},
	})
}
