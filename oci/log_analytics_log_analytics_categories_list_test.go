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
	logAnalyticsCategoriesListSingularDataSourceRepresentation = map[string]interface{}{
		"namespace":             Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"category_display_text": Representation{RepType: Optional, Create: `Oracle`},
		"category_type":         Representation{RepType: Optional, Create: `VENDOR,PRODUCT`},
	}

	LogAnalyticsCategoriesListResourceConfig = "" +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

func TestLogAnalyticsLogAnalyticsCategoriesListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsCategoriesListResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	requiredDatasourceName := "data.oci_log_analytics_log_analytics_categories_list.test_log_analytics_categories_list_required"
	optionalDatasourceName := "data.oci_log_analytics_log_analytics_categories_list.test_log_analytics_categories_list_optional"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify required input
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_categories_list", "test_log_analytics_categories_list_required", Required, Create, logAnalyticsCategoriesListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsCategoriesListResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "items.0.name"),
			),
		},
		// verify optionals
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_categories_list", "test_log_analytics_categories_list_optional", Optional, Create, logAnalyticsCategoriesListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsCategoriesListResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "items.0.name"),
			),
		},
	})
}
