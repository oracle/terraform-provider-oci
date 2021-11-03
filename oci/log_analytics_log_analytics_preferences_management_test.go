// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	logAnalyticsPreferencesManagementRepresentation = map[string]interface{}{
		"namespace": Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`, Update: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"items":     []RepresentationGroup{{Required, logAnalyticsPreferencesManagementItemsRepresentation}},
	}
	logAnalyticsPreferencesManagementItemsRepresentation = map[string]interface{}{
		"name":  Representation{RepType: Required, Create: `DEFAULT_HOMEPAGE`, Update: `DEFAULT_HOMEPAGE`},
		"value": Representation{RepType: Required, Create: `value1`, Update: `value2`},
	}

	LogAnalyticsPreferencesManagementResourceDependencies = "" +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsPreferencesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsPreferencesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_preferences_management.test_log_analytics_preferences_management"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsPreferencesManagementResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", Required, Update, logAnalyticsPreferencesManagementRepresentation), "loganalytics", "logAnalyticsPreferencesManagement", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsPreferencesManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", Required, Create, logAnalyticsPreferencesManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
			),
		},

		// verify update
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsPreferencesManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", Required, Update, logAnalyticsPreferencesManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
			),
		},

		// verify delete
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsPreferencesManagementResourceDependencies,
		},
	})
}
