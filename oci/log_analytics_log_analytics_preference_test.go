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
	logAnalyticsPreferenceSingularDataSourceRepresentation = map[string]interface{}{
		"namespace": Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsPreferenceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsPreferenceResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_preferences_management.test_log_analytics_preferences_management"
	singularDatasourceName := "data.oci_log_analytics_log_analytics_preference.test_log_analytics_preference"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// Set preference to test data source
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsPreferencesManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_preferences_management", "test_log_analytics_preferences_management", Required, Create, logAnalyticsPreferencesManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsPreferencesManagementResourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_preference", "test_log_analytics_preference", Required, Create, logAnalyticsPreferenceSingularDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},

		// Delete the preference
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsPreferencesManagementResourceDependencies,
		},
	})
}
