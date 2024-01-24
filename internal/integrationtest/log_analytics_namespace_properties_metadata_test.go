// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	LogAnalyticsLogAnalyticsNamespacePropertiesMetadataDataSourceRepresentation = map[string]interface{}{
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"level":     acctest.Representation{RepType: acctest.Optional, Create: `SOURCE`},
		"name":      acctest.Representation{RepType: acctest.Optional, Create: `management_agent.database_sql.badsql_retry`},
	}

	LogAnalyticsNamespacePropertiesMetadataResourceConfig = "" +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsNamespacePropertiesMetadataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsNamespacePropertiesMetadataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	requiredDatasourceName := "data.oci_log_analytics_namespace_properties_metadata.test_properties_metadata_required"
	optionalDatasourceName := "data.oci_log_analytics_namespace_properties_metadata.test_properties_metadata_optional"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource with required parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_properties_metadata", "test_properties_metadata_required", acctest.Required, acctest.Create, LogAnalyticsLogAnalyticsNamespacePropertiesMetadataDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespacePropertiesMetadataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "property_metadata_summary_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(requiredDatasourceName, "property_metadata_summary_collection.0.items.0.display_name"),
			),
		},
		// verify datasource with optional parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_log_analytics_namespace_properties_metadata", "test_properties_metadata_optional", acctest.Optional, acctest.Create, LogAnalyticsLogAnalyticsNamespacePropertiesMetadataDataSourceRepresentation) +
				compartmentIdVariableStr + LogAnalyticsNamespacePropertiesMetadataResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(optionalDatasourceName, "namespace"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "property_metadata_summary_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "property_metadata_summary_collection.0.items.0.name", "management_agent.database_sql.badsql_retry"),
				resource.TestCheckResourceAttr(optionalDatasourceName, "property_metadata_summary_collection.0.items.0.display_name", "management_agent.database_sql.badsql_retry"),
			),
		},
	})
}
