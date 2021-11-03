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
	LogAnalyticsResourceCategoriesManagementRequiredOnlyResource = LogAnalyticsResourceCategoriesManagementResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", Required, Create, logAnalyticsResourceCategoriesManagementRepresentation)

	logAnalyticsResourceCategoriesManagementRepresentation = map[string]interface{}{
		"namespace":           Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`, Update: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"resource_categories": Representation{RepType: Required, Create: []string{`cat1`, `cat2`, `cat3`}, Update: []string{`cat1`, `cat4`, `cat5`}},
		"resource_id":         Representation{RepType: Required, Create: `resource1`, Update: `resource1`},
		"resource_type":       Representation{RepType: Required, Create: `DASHBOARD`, Update: `DASHBOARD`},
	}

	LogAnalyticsResourceCategoriesManagementResourceDependencies = "" +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsResourceCategoriesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsResourceCategoriesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_resource_categories_management.test_log_analytics_resource_categories_management"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsResourceCategoriesManagementResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", Optional, Create, logAnalyticsResourceCategoriesManagementRepresentation), "loganalytics", "logAnalyticsResourceCategoriesManagement", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsResourceCategoriesManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", Required, Create, logAnalyticsResourceCategoriesManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "resource_id", "resource1"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "DASHBOARD"),
				resource.TestCheckResourceAttr(resourceName, "resource_categories.#", "3"),
			),
		},

		// verify update
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsResourceCategoriesManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", Required, Update, logAnalyticsResourceCategoriesManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "resource_id", "resource1"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "DASHBOARD"),
				resource.TestCheckResourceAttr(resourceName, "resource_categories.#", "3"),
				//
			),
		},

		// verify delete
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsResourceCategoriesManagementResourceDependencies,
		},
	})
}
