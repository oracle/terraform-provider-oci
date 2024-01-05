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

const zipFile = "../examples/log_analytics/import_custom_content/files/TFSource1.zip"

var (
	LogAnalyticsLogAnalyticsImportCustomContentRequiredOnlyResource = LogAnalyticsLogAnalyticsImportCustomContentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_import_custom_content", "test_log_analytics_import_custom_content", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsImportCustomContentRepresentation)

	LogAnalyticsLogAnalyticsImportCustomContentRepresentation = map[string]interface{}{
		"import_custom_content_file": acctest.Representation{RepType: acctest.Required, Create: zipFile},
		"namespace":                  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"is_overwrite":               acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"expect":                     acctest.Representation{RepType: acctest.Optional, Create: `100-Continue`},
	}

	LogAnalyticsLogAnalyticsImportCustomContentResourceDependencies = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsImportCustomContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsImportCustomContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_import_custom_content.test_log_analytics_import_custom_content"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsLogAnalyticsImportCustomContentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_import_custom_content", "test_log_analytics_import_custom_content", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsImportCustomContentRepresentation), "loganalytics", "logAnalyticsImportCustomContent", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsImportCustomContentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_import_custom_content", "test_log_analytics_import_custom_content", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsImportCustomContentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "import_custom_content_file", zipFile),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsImportCustomContentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsLogAnalyticsImportCustomContentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_import_custom_content", "test_log_analytics_import_custom_content", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsImportCustomContentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "import_custom_content_file", zipFile),
				resource.TestCheckResourceAttr(resourceName, "is_overwrite", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "expect", "100-Continue"),
			),
		},
	})
}
