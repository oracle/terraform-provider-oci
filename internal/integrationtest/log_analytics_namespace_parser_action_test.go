// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LogAnalyticsNamespaceParserActionDataSourceRepresentation = map[string]interface{}{
		"namespace":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"action_display_text": acctest.Representation{RepType: acctest.Optional, Create: `Decode`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `BASE64_DECODE_AND_UNZIP`},
	}

	LogAnalyticsNamespaceParserActionResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespaceParserActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespaceParserActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_log_analytics_namespace_parser_actions.test_namespace_parser_actions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_parser_actions", "test_namespace_parser_actions", acctest.Required, acctest.Create, LogAnalyticsNamespaceParserActionDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceParserActionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(datasourceName, "parser_action_summary_collection.#"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_parser_actions", "test_namespace_parser_actions", acctest.Optional, acctest.Create, LogAnalyticsNamespaceParserActionDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespaceParserActionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "action_display_text", "Decode"),
				resource.TestCheckResourceAttr(datasourceName, "name", "BASE64_DECODE_AND_UNZIP"),
				resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
				resource.TestCheckResourceAttr(datasourceName, "parser_action_summary_collection.#", "1"),
			),
		},
	})
}
