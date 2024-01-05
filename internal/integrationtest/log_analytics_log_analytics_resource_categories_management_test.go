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
	LogAnalyticsLogAnalyticsResourceCategoriesManagementRequiredOnlyResource = LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsResourceCategoriesManagementRepresentation)

	LogAnalyticsLogAnalyticsResourceCategoriesManagementRepresentation = map[string]interface{}{
		"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`, Update: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"resource_categories": acctest.Representation{RepType: acctest.Required, Create: []string{`cat1`, `cat2`, `cat3`}, Update: []string{`cat1`, `cat4`, `cat5`}},
		"resource_id":         acctest.Representation{RepType: acctest.Required, Create: `resource1`, Update: `resource1`},
		"resource_type":       acctest.Representation{RepType: acctest.Required, Create: `DASHBOARD`, Update: `DASHBOARD`},
	}

	LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceDependencies = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsResourceCategoriesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsResourceCategoriesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_resource_categories_management.test_log_analytics_resource_categories_management"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsResourceCategoriesManagementRepresentation), "loganalytics", "logAnalyticsResourceCategoriesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsResourceCategoriesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "resource_id", "resource1"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "DASHBOARD"),
				resource.TestCheckResourceAttr(resourceName, "resource_categories.#", "3"),
			),
		},

		// verify update
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", acctest.Required, acctest.Update, LogAnalyticsLogAnalyticsResourceCategoriesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "resource_id", "resource1"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "DASHBOARD"),
				resource.TestCheckResourceAttr(resourceName, "resource_categories.#", "3"),
				//
			),
		},

		// verify delete
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsResourceCategoriesManagementResourceDependencies,
		},
	})
}
