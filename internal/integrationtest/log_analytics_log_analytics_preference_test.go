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
	LogAnalyticsLogAnalyticsLogAnalyticsPreferenceSingularDataSourceRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsPreferenceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsPreferenceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_preferences_management.test_log_analytics_preferences_management"
	singularDatasourceName := "data.oci_log_analytics_log_analytics_preference.test_log_analytics_preference"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Set preference to test data source
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsPreferencesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsPreferencesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsPreferencesManagementResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_preference", "test_log_analytics_preference", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsLogAnalyticsPreferenceSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},

		// Delete the preference
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsPreferencesManagementResourceDependencies,
		},
	})
}
