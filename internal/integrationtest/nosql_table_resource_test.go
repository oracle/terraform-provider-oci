// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	OnDemandTableResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Required, acctest.Create, onDemandTableRepresentation)

	onDemandTableDdlStatement = "CREATE TABLE IF NOT EXISTS test_ondemand(id INTEGER, name STRING, age STRING, PRIMARY KEY(SHARD(id)))"

	onDemandTableRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ddl_statement":  acctest.Representation{RepType: acctest.Required, Create: onDemandTableDdlStatement},
		"name":           acctest.Representation{RepType: acctest.Required, Create: "test_ondemand"},
		"table_limits":   acctest.RepresentationGroup{RepType: acctest.Required, Group: onDemandTableLimitsRepresentation},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreOnDemandTableLimitsReadWrite},
	}
	onDemandTableLimitsRepresentation = map[string]interface{}{
		"max_read_units":     acctest.Representation{RepType: acctest.Required, Create: `0`},
		"max_write_units":    acctest.Representation{RepType: acctest.Required, Create: `0`},
		"max_storage_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `5`},
		"capacity_mode":      acctest.Representation{RepType: acctest.Required, Create: `ON_DEMAND`},
	}
	ignoreOnDemandTableLimitsReadWrite = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`table_limits[0].max_read_units`, `table_limits[0].max_write_units`}},
	}

	onDemandNoLfcTableLimitsRepresentation = acctest.RepresentationCopyWithRemovedProperties(onDemandTableRepresentation, []string{"lifecycle"})
	onDemandToPrevisonTableRepresentation  = acctest.RepresentationCopyWithNewProperties(
		onDemandNoLfcTableLimitsRepresentation,
		map[string]interface{}{
			"table_limits": acctest.RepresentationGroup{RepType: acctest.Required, Group: tableTableLimitsRepresentation},
		},
	)

	onDemandTableDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `test_ondemand`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: onDemandTableDataSourceFilterRepresentation}}
	onDemandTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_nosql_table.test_ondemand.id}`}},
	}

	onDemandTableSingularDataSourceRepresentation = map[string]interface{}{
		"table_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_nosql_table.test_ondemand.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}
)

// issue-routing-tag: nosql/default
func TestNosqlTableResource_test(t *testing.T) {
	httpreplay.SetScenario("TestNosqlTableResource_test")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_nosql_table.test_ondemand"
	datasourceName := "data.oci_nosql_tables.test_tables"
	singularDatasourceName := "data.oci_nosql_table.test_ondemand"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckNosqlTableDestroy,
		Steps: []resource.TestStep{

			// verify create auto scaling table
			{
				Config: config + compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Required, acctest.Create, onDemandTableRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "ddl_statement", onDemandTableDdlStatement),
					resource.TestCheckResourceAttr(resourceName, "name", "test_ondemand"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "table_limits.0.max_read_units"),
					resource.TestCheckResourceAttrSet(resourceName, "table_limits.0.max_write_units"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_storage_in_gbs", "5"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.0.capacity_mode", "ON_DEMAND"),
				),
			},

			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_tables", "test_tables", acctest.Optional, acctest.Create, onDemandTableDataSourceRepresentation) +
					compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Required, acctest.Create, onDemandTableRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "name", "test_ondemand"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "table_collection.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "table_collection.0.id"),
				),
			},

			// verify singular datasource: test_ondemand
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Required, acctest.Create, onDemandTableSingularDataSourceRepresentation) +
					compartmentIdVariableStr + OnDemandTableResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "table_name_or_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "test_ondemand"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ddl_statement", onDemandTableDdlStatement),
					resource.TestCheckResourceAttr(singularDatasourceName, "schema.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "table_limits.0.max_read_units"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "table_limits.0.max_write_units"),
					resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.max_storage_in_gbs", "5"),
					resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.capacity_mode", "ON_DEMAND"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},

			// verify update table limits to PREVISIONED
			{
				Config: config + compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Optional, acctest.Create, onDemandToPrevisonTableRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "ddl_statement", onDemandTableDdlStatement),
					resource.TestCheckResourceAttr(resourceName, "name", "test_ondemand"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_read_units", "10"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_storage_in_gbs", "10"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_write_units", "10"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.0.capacity_mode", "PROVISIONED"),
				),
			},
		},
	})
}
