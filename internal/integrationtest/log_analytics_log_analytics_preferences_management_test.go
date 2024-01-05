// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	LogAnalyticsLogAnalyticsPreferencesManagementRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`, Update: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"items":     []acctest.RepresentationGroup{{RepType: acctest.Required, Group: LogAnalyticsLogAnalyticsPreferencesManagementItemsRepresentation}},
	}
	LogAnalyticsLogAnalyticsPreferencesManagementItemsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `DEFAULT_HOMEPAGE`, Update: `DEFAULT_HOMEPAGE`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value1`, Update: `value2`},
	}

	LogAnalyticsLogAnalyticsPreferencesManagementResourceDependencies = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsPreferencesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsPreferencesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_preferences_management.test_log_analytics_preferences_management"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsLogAnalyticsPreferencesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", acctest.Required, acctest.Update, LogAnalyticsLogAnalyticsPreferencesManagementRepresentation), "loganalytics", "logAnalyticsPreferencesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsPreferencesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsPreferencesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
			),
		},

		// verify update
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsPreferencesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", acctest.Required, acctest.Update, LogAnalyticsLogAnalyticsPreferencesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
			),
		},

		// verify delete
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsPreferencesManagementResourceDependencies,
		},
	})
}
