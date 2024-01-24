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
	OpsiOpsiAwrHubAwrSnapshotSingularDataSourceRepresentation = map[string]interface{}{
		"awr_hub_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"awr_source_database_identifier": acctest.Representation{RepType: acctest.Required, Create: `12345`},
		"time_greater_than_or_equal_to":  acctest.Representation{RepType: acctest.Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":     acctest.Representation{RepType: acctest.Optional, Create: `timeLessThanOrEqualTo`},
	}

	OpsiOpsiAwrHubAwrSnapshotDataSourceRepresentation = map[string]interface{}{
		"awr_hub_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"awr_source_database_identifier": acctest.Representation{RepType: acctest.Required, Create: `12345`},
		"time_greater_than_or_equal_to":  acctest.Representation{RepType: acctest.Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":     acctest.Representation{RepType: acctest.Optional, Create: `timeLessThanOrEqualTo`},
	}

	OpsiAwrHubAwrSnapshotResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", acctest.Required, acctest.Create, OpsiAwrHubRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiAwrHubAwrSnapshotResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiAwrHubAwrSnapshotResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_opsi_awr_hub_awr_snapshots.test_awr_hub_awr_snapshots"
	singularDatasourceName := "data.oci_opsi_awr_hub_awr_snapshot.test_awr_hub_awr_snapshot"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			// Commenting out the Verification steps of the below two steps as the awr snapshot API will return no results
			// as there are no active sources configured.
			//Source configuration is a manual step and requires a user to login to each Oracle database and run SQL queries.

			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub_awr_snapshots", "test_awr_hub_awr_snapshots", acctest.Required, acctest.Create, OpsiOpsiAwrHubAwrSnapshotDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiAwrHubAwrSnapshotResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				/*resource.TestCheckResourceAttrSet(datasourceName, "awr_hub_id"),
				resource.TestCheckResourceAttr(datasourceName, "awr_source_database_identifier", "awrSourceDatabaseIdentifier"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_less_than_or_equal_to"),*/

				resource.TestCheckResourceAttrSet(datasourceName, "awr_snapshot_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "awr_snapshot_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub_awr_snapshot", "test_awr_hub_awr_snapshot", acctest.Required, acctest.Create, OpsiOpsiAwrHubAwrSnapshotSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OpsiAwrHubAwrSnapshotResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				/*resource.TestCheckResourceAttrSet(singularDatasourceName, "awr_hub_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "awr_source_database_identifier", "awrSourceDatabaseIdentifier"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_greater_than_or_equal_to", "timeGreaterThanOrEqualTo"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_less_than_or_equal_to", "timeLessThanOrEqualTo"),*/

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "0"),
			),
		},
	})
}
