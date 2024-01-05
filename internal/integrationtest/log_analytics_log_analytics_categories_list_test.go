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
	LogAnalyticsLogAnalyticsLogAnalyticsCategoriesListSingularDataSourceRepresentation = map[string]interface{}{
		"namespace":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"category_display_text": acctest.Representation{RepType: acctest.Optional, Create: `Oracle`},
		"category_type":         acctest.Representation{RepType: acctest.Optional, Create: `VENDOR,PRODUCT`},
	}

	LogAnalyticsLogAnalyticsCategoriesListResourceConfig = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceSingularDataSourceRepresentation)
)

func TestLogAnalyticsLogAnalyticsCategoriesListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsCategoriesListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	requiredDatasourceName := "data.oci_log_analytics_log_analytics_categories_list.test_log_analytics_categories_list_required"
	optionalDatasourceName := "data.oci_log_analytics_log_analytics_categories_list.test_log_analytics_categories_list_optional"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify required input
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_categories_list", "test_log_analytics_categories_list_required", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsLogAnalyticsCategoriesListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsLogAnalyticsCategoriesListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "items.0.name"),
			),
		},
		// verify optionals
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_categories_list", "test_log_analytics_categories_list_optional", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsLogAnalyticsCategoriesListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsLogAnalyticsCategoriesListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "items.0.name"),
			),
		},
	})
}
