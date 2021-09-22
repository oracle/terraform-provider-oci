// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

const zipFile = "../examples/log_analytics/files/TFSource1.zip"

var (
	LogAnalyticsImportCustomContentRequiredOnlyResource = LogAnalyticsImportCustomContentResourceDependencies +
		generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_import_custom_content", "test_log_analytics_import_custom_content", Required, Create, logAnalyticsImportCustomContentRepresentation)

	logAnalyticsImportCustomContentRepresentation = map[string]interface{}{
		"import_custom_content_file": Representation{repType: Required, create: zipFile},
		"namespace":                  Representation{repType: Required, create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"is_overwrite":               Representation{repType: Optional, create: `true`},
	}

	LogAnalyticsImportCustomContentResourceDependencies = "" +
		generateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsImportCustomContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsImportCustomContentResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_import_custom_content.test_log_analytics_import_custom_content"

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+LogAnalyticsImportCustomContentResourceDependencies+
		generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_import_custom_content", "test_log_analytics_import_custom_content", Optional, Create, logAnalyticsImportCustomContentRepresentation), "loganalytics", "logAnalyticsImportCustomContent", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsImportCustomContentResourceDependencies +
				generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_import_custom_content", "test_log_analytics_import_custom_content", Required, Create, logAnalyticsImportCustomContentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "import_custom_content_file", zipFile),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsImportCustomContentResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsImportCustomContentResourceDependencies +
				generateResourceFromRepresentationMap("oci_log_analytics_log_analytics_import_custom_content", "test_log_analytics_import_custom_content", Optional, Create, logAnalyticsImportCustomContentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "import_custom_content_file", zipFile),
				resource.TestCheckResourceAttr(resourceName, "is_overwrite", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
			),
		},
	})
}
