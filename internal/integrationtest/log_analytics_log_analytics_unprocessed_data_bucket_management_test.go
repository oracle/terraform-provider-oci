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
	logAnalyticsUnprocessedDataBucketManagementRepresentation = map[string]interface{}{
		"bucket":     acctest.Representation{RepType: acctest.Required, Create: `dummy_bucket`, Update: `udb_tf`},
		"namespace":  acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"is_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
	}

	LogAnalyticsUnprocessedDataBucketManagementResourceDependencies = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_unprocessed_data_bucket_management.test_log_analytics_unprocessed_data_bucket_management"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsUnprocessedDataBucketManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_unprocessed_data_bucket_management", "test_log_analytics_unprocessed_data_bucket_management", acctest.Required, acctest.Create, logAnalyticsUnprocessedDataBucketManagementRepresentation), "loganalytics", "logAnalyticsUnprocessedDataBucketManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsUnprocessedDataBucketManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_unprocessed_data_bucket_management", "test_log_analytics_unprocessed_data_bucket_management", acctest.Required, acctest.Create, logAnalyticsUnprocessedDataBucketManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", "dummy_bucket"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
			),
		},

		// verify update
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsUnprocessedDataBucketManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_unprocessed_data_bucket_management", "test_log_analytics_unprocessed_data_bucket_management", acctest.Required, acctest.Update, logAnalyticsUnprocessedDataBucketManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", "udb_tf"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
			),
		},

		// verify delete
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsUnprocessedDataBucketManagementResourceDependencies,
		},
	})
}
