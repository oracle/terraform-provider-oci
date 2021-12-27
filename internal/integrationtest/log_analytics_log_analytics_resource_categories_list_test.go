// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	logAnalyticsResourceCategoriesListSingularDataSourceRepresentation = map[string]interface{}{
		"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"resource_categories": acctest.Representation{RepType: acctest.Optional, Create: `cat1,cat2`},
	}

	LogAnalyticsResourceCategoriesListResourceConfig = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsResourceCategoriesListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsResourceCategoriesListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_resource_categories_management.test_log_analytics_resource_categories_management"

	requiredDatasourceName := "data.oci_log_analytics_log_analytics_resource_categories_list.test_log_analytics_resource_categories_list_required"
	optionalDatasourceName := "data.oci_log_analytics_log_analytics_resource_categories_list.test_log_analytics_resource_categories_list_optional"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// Set categories to test data source
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsResourceCategoriesListResourceConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", acctest.Required, acctest.Create, logAnalyticsResourceCategoriesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "resource_id", "resource1"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "DASHBOARD"),
				resource.TestCheckResourceAttr(resourceName, "resource_categories.#", "3"),
			),
		},
		// verify required input
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", acctest.Required, acctest.Create, logAnalyticsResourceCategoriesManagementRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_list", "test_log_analytics_resource_categories_list_required", acctest.Required, acctest.Create, logAnalyticsResourceCategoriesListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsResourceCategoriesListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "items.0.resource_id"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "items.0.resource_type"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "items.0.is_system"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "items.0.category_name"),
			),
		},
		// verify optionals
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_management", "test_log_analytics_resource_categories_management", acctest.Required, acctest.Create, logAnalyticsResourceCategoriesManagementRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_resource_categories_list", "test_log_analytics_resource_categories_list_optional", acctest.Optional, acctest.Create, logAnalyticsResourceCategoriesListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsResourceCategoriesListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "resource_categories", "cat1,cat2"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "items.#", "2"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "items.0.resource_id"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "items.0.resource_type"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "items.0.is_system"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "items.0.category_name"),
			),
		},
		// verify delete
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsResourceCategoriesListResourceConfig,
		},
	})
}
