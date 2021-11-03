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
	logAnalyticsUnprocessedDataBucketManagementRepresentation = map[string]interface{}{
		"bucket":     Representation{RepType: Required, Create: `dummy_bucket`, Update: `udb_tf`},
		"namespace":  Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"is_enabled": Representation{RepType: Required, Create: `false`, Update: `true`},
	}

	LogAnalyticsUnprocessedDataBucketManagementResourceDependencies = "" +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_log_analytics_log_analytics_unprocessed_data_bucket_management.test_log_analytics_unprocessed_data_bucket_management"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+LogAnalyticsUnprocessedDataBucketManagementResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_unprocessed_data_bucket_management", "test_log_analytics_unprocessed_data_bucket_management", Required, Create, logAnalyticsUnprocessedDataBucketManagementRepresentation), "loganalytics", "logAnalyticsUnprocessedDataBucketManagement", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsUnprocessedDataBucketManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_unprocessed_data_bucket_management", "test_log_analytics_unprocessed_data_bucket_management", Required, Create, logAnalyticsUnprocessedDataBucketManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", "dummy_bucket"),
				resource.TestCheckResourceAttrSet(resourceName, "namespace"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
			),
		},

		// verify update
		{
			Config: config + compartmentIdVariableStr + LogAnalyticsUnprocessedDataBucketManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_log_analytics_log_analytics_unprocessed_data_bucket_management", "test_log_analytics_unprocessed_data_bucket_management", Required, Update, logAnalyticsUnprocessedDataBucketManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
