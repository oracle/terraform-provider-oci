// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	OpsiOpsiAwrHubAwrSourcesSummarySingularDataSourceRepresentation = map[string]interface{}{
		"awr_hub_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	OpsiOpsiAwrHubAwrSourcesSummaryDataSourceRepresentation = map[string]interface{}{
		"awr_hub_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	OpsiAwrHubAwrSourcesSummaryResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Required, acctest.Create, OpsiAwrHubRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiAwrHubAwrSourcesSummaryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiAwrHubAwrSourcesSummaryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_opsi_awr_hub_awr_sources_summary.test_awr_hub_awr_sources_summary"
	singularDatasourceName := "data.oci_opsi_awr_hub_awr_sources_summary.test_awr_hub_awr_sources_summary"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			// Commenting out the Verification steps of the below two steps as the awr snapshot API will return no results
			// as there are no active sources configured.
			//Source configuration is a manual step and requires a user to login to each Oracle database and run SQL queries.

			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub_awr_sources_summary", "test_awr_hub_awr_sources_summary", acctest.Required, acctest.Create, OpsiOpsiAwrHubAwrSourcesSummaryDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiAwrHubAwrSourcesSummaryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub_awr_sources_summary", "test_awr_hub_awr_sources_summary", acctest.Required, acctest.Create, OpsiOpsiAwrHubAwrSourcesSummarySingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiAwrHubAwrSourcesSummaryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				/*resource.TestCheckResourceAttrSet(singularDatasourceName, "awr_hub_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),*/

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "0"),
			),
		},
	})
}
