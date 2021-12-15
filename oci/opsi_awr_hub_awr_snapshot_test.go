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
	awrHubAwrSnapshotSingularDataSourceRepresentation = map[string]interface{}{
		"awr_hub_id":                     Representation{RepType: Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"awr_source_database_identifier": Representation{RepType: Required, Create: `12345`},
		"time_greater_than_or_equal_to":  Representation{RepType: Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":     Representation{RepType: Optional, Create: `timeLessThanOrEqualTo`},
	}

	awrHubAwrSnapshotDataSourceRepresentation = map[string]interface{}{
		"awr_hub_id":                     Representation{RepType: Required, Create: `${oci_opsi_awr_hub.test_awr_hub.id}`},
		"awr_source_database_identifier": Representation{RepType: Required, Create: `12345`},
		"time_greater_than_or_equal_to":  Representation{RepType: Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":     Representation{RepType: Optional, Create: `timeLessThanOrEqualTo`},
	}

	AwrHubAwrSnapshotResourceConfig = GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_opsi_awr_hub", "test_awr_hub", Required, Create, awrHubRepresentation) +
		GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Required, Create, operationsInsightsWarehouseRepresentation)
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiAwrHubAwrSnapshotResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiAwrHubAwrSnapshotResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_opsi_awr_hub_awr_snapshots.test_awr_hub_awr_snapshots"
	singularDatasourceName := "data.oci_opsi_awr_hub_awr_snapshot.test_awr_hub_awr_snapshot"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			// Commenting out the Verification steps of the below two steps as the awr snapshot API will return no results
			// as there are no active sources configured.
			//Source configuration is a manual step and requires a user to login to each Oracle database and run SQL queries.

			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub_awr_snapshots", "test_awr_hub_awr_snapshots", Required, Create, awrHubAwrSnapshotDataSourceRepresentation) +
				compartmentIdVariableStr + AwrHubAwrSnapshotResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_opsi_awr_hub_awr_snapshot", "test_awr_hub_awr_snapshot", Required, Create, awrHubAwrSnapshotSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AwrHubAwrSnapshotResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				/*resource.TestCheckResourceAttrSet(singularDatasourceName, "awr_hub_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "awr_source_database_identifier", "awrSourceDatabaseIdentifier"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_greater_than_or_equal_to", "timeGreaterThanOrEqualTo"),
				resource.TestCheckResourceAttr(singularDatasourceName, "time_less_than_or_equal_to", "timeLessThanOrEqualTo"),*/

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "0"),
			),
		},
	})
}
