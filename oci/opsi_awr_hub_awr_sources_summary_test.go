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
	awrHubAwrSourcesSummarySingularDataSourceRepresentation = map[string]interface{}{
		"awr_hub_id":     Representation{RepType: Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"name":           Representation{RepType: Optional, Create: `name`},
	}

	awrHubAwrSourcesSummaryDataSourceRepresentation = map[string]interface{}{
		"awr_hub_id":     Representation{RepType: Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"name":           Representation{RepType: Optional, Create: `name`},
	}

	AwrHubAwrSourcesSummaryResourceConfig = GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", Required, Create, awrHubRepresentation) +
		GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Required, Create, operationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiAwrHubAwrSourcesSummaryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiAwrHubAwrSourcesSummaryResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_opsi_awr_hub_awr_sources_summary.test_awr_hub_awr_sources_summary"
	singularDatasourceName := "data.oci_opsi_awr_hub_awr_sources_summary.test_awr_hub_awr_sources_summary"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			// Commenting out the Verification steps of the below two steps as the awr snapshot API will return no results
			// as there are no active sources configured.
			//Source configuration is a manual step and requires a user to login to each Oracle database and run SQL queries.

			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub_awr_sources_summary", "test_awr_hub_awr_sources_summary", Required, Create, awrHubAwrSourcesSummaryDataSourceRepresentation) +
				compartmentIdVariableStr + AwrHubAwrSourcesSummaryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				/*resource.TestCheckResourceAttrSet(datasourceName, "awr_hub_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),

				resource.TestCheckResourceAttrSet(datasourceName, "summarize_awr_sources_summaries_collection.#"),*/
				resource.TestCheckResourceAttr(datasourceName, "summarize_awr_sources_summaries_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub_awr_sources_summary", "test_awr_hub_awr_sources_summary", Required, Create, awrHubAwrSourcesSummarySingularDataSourceRepresentation) +
				compartmentIdVariableStr + AwrHubAwrSourcesSummaryResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				/*resource.TestCheckResourceAttrSet(singularDatasourceName, "awr_hub_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),*/

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "0"),
			),
		},
	})
}
