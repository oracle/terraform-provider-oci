// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
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
		"max_storage_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `5`, Update: `6`},
		"capacity_mode":      acctest.Representation{RepType: acctest.Required, Create: `ON_DEMAND`},
	}
	ignoreOnDemandTableLimitsReadWrite = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`table_limits[0].max_read_units`, `table_limits[0].max_write_units`}},
	}

	onDemandNoLfcTableRepresentation      = acctest.RepresentationCopyWithRemovedProperties(onDemandTableRepresentation, []string{"lifecycle"})
	onDemandToPrevisonTableRepresentation = acctest.RepresentationCopyWithNewProperties(
		onDemandNoLfcTableRepresentation,
		map[string]interface{}{
			"table_limits": acctest.RepresentationGroup{RepType: acctest.Required, Group: NosqlTableTableLimitsRepresentation},
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

	ChildTableResourceConfig = ChildTableResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_child", acctest.Required, acctest.Create, childTableRepresentation)

	childTableDdlStatement = "CREATE TABLE IF NOT EXISTS test_table.test_child(idc INTEGER, cname STRING, PRIMARY KEY(idc))"

	childTableRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"ddl_statement":  acctest.Representation{RepType: acctest.Required, Create: childTableDdlStatement},
		"name":           acctest.Representation{RepType: acctest.Required, Create: "test_table.test_child"},
		"depends_on":     acctest.Representation{RepType: acctest.Required, Create: []string{"oci_nosql_table.test_table"}},
	}
	childTableSingularDataSourceRepresentation = map[string]interface{}{
		"table_name_or_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_nosql_table.test_child.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	childTableDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `test_table.test_child`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: childTableDataSourceFilterRepresentation},
	}
	childTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_nosql_table.test_child.id}`}},
	}

	ChildTableResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_table", acctest.Required, acctest.Create, NosqlTableRepresentation)
)

// issue-routing-tag: nosql/default
func TestNosqlTableResource_test(t *testing.T) {
	httpreplay.SetScenario("TestNosqlTableResource_test")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	ondemandResourceName := "oci_nosql_table.test_ondemand"
	ondemandDatasourceName := "data.oci_nosql_tables.test_tables"
	ondemandSingularDatasourceName := "data.oci_nosql_table.test_ondemand"

	childResourceName := "oci_nosql_table.test_child"
	childDataResourceName := "data.oci_nosql_tables.test_child"
	singularChildDatasourceName := "data.oci_nosql_table.test_child"

	var timeUpdated0, timeUpdated1 string
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckNosqlTableDestroy,
		Steps: []resource.TestStep{

			// verify create ondemand table
			{
				Config: config + compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Required, acctest.Create, onDemandTableRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(ondemandResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(ondemandResourceName, "ddl_statement", onDemandTableDdlStatement),
					resource.TestCheckResourceAttr(ondemandResourceName, "name", "test_ondemand"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.#", "1"),
					resource.TestCheckResourceAttrSet(ondemandResourceName, "table_limits.0.max_read_units"),
					resource.TestCheckResourceAttrSet(ondemandResourceName, "table_limits.0.max_write_units"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.0.max_storage_in_gbs", "5"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.0.capacity_mode", "ON_DEMAND"),
					func(s *terraform.State) (err error) {
						timeUpdated0, err = acctest.FromInstanceState(s, ondemandResourceName, "time_updated")
						return err
					},
				),
			},

			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_tables", "test_tables", acctest.Optional, acctest.Create, onDemandTableDataSourceRepresentation) +
					compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Required, acctest.Create, onDemandTableRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(ondemandDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(ondemandDatasourceName, "name", "test_ondemand"),
					resource.TestCheckResourceAttr(ondemandDatasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(ondemandDatasourceName, "table_collection.#", "1"),
					resource.TestCheckResourceAttrSet(ondemandDatasourceName, "table_collection.0.id"),
				),
			},

			// verify singular datasource: test_ondemand
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Required, acctest.Create, onDemandTableSingularDataSourceRepresentation) +
					compartmentIdVariableStr + OnDemandTableResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(ondemandSingularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(ondemandSingularDatasourceName, "table_name_or_id"),
					resource.TestCheckResourceAttrSet(ondemandSingularDatasourceName, "id"),
					resource.TestCheckResourceAttr(ondemandSingularDatasourceName, "name", "test_ondemand"),
					resource.TestCheckResourceAttr(ondemandSingularDatasourceName, "ddl_statement", onDemandTableDdlStatement),
					resource.TestCheckResourceAttr(ondemandSingularDatasourceName, "schema.#", "1"),
					resource.TestCheckResourceAttr(ondemandSingularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(ondemandSingularDatasourceName, "table_limits.#", "1"),
					resource.TestCheckResourceAttrSet(ondemandSingularDatasourceName, "table_limits.0.max_read_units"),
					resource.TestCheckResourceAttrSet(ondemandSingularDatasourceName, "table_limits.0.max_write_units"),
					resource.TestCheckResourceAttr(ondemandSingularDatasourceName, "table_limits.0.max_storage_in_gbs", "5"),
					resource.TestCheckResourceAttr(ondemandSingularDatasourceName, "table_limits.0.capacity_mode", "ON_DEMAND"),
					resource.TestCheckResourceAttrSet(ondemandSingularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(ondemandSingularDatasourceName, "time_updated"),
				),
			},

			// verify no change on reapplying
			{
				Config: config + compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Optional, acctest.Create, onDemandNoLfcTableRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(ondemandResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(ondemandResourceName, "ddl_statement", onDemandTableDdlStatement),
					resource.TestCheckResourceAttr(ondemandResourceName, "name", "test_ondemand"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.#", "1"),
					resource.TestCheckResourceAttrSet(ondemandResourceName, "table_limits.0.max_read_units"),
					resource.TestCheckResourceAttrSet(ondemandResourceName, "table_limits.0.max_write_units"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.0.max_storage_in_gbs", "5"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.0.capacity_mode", "ON_DEMAND"),
					func(s *terraform.State) (err error) {
						timeUpdated1, err = acctest.FromInstanceState(s, ondemandResourceName, "time_updated")
						if timeUpdated0 != timeUpdated1 {
							return fmt.Errorf("Resource updated when it was supposed to be no change.")
						}
						return err
					},
				),
			},

			// update max_storage_in_gbs of TableLimits
			{
				Config: config + compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Optional, acctest.Update, onDemandNoLfcTableRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(ondemandResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(ondemandResourceName, "ddl_statement", onDemandTableDdlStatement),
					resource.TestCheckResourceAttr(ondemandResourceName, "name", "test_ondemand"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.#", "1"),
					resource.TestCheckResourceAttrSet(ondemandResourceName, "table_limits.0.max_read_units"),
					resource.TestCheckResourceAttrSet(ondemandResourceName, "table_limits.0.max_write_units"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.0.max_storage_in_gbs", "6"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.0.capacity_mode", "ON_DEMAND"),
				),
			},

			// verify update table limits to PREVISIONED
			{
				Config: config + compartmentIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_ondemand", acctest.Optional, acctest.Create, onDemandToPrevisonTableRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(ondemandResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(ondemandResourceName, "ddl_statement", onDemandTableDdlStatement),
					resource.TestCheckResourceAttr(ondemandResourceName, "name", "test_ondemand"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.#", "1"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.0.max_read_units", "10"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.0.max_storage_in_gbs", "10"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.0.max_write_units", "10"),
					resource.TestCheckResourceAttr(ondemandResourceName, "table_limits.0.capacity_mode", "PROVISIONED"),
				),
			},

			// verify create table: child table table "test_table.test_child"
			{
				Config: config + compartmentIdVariableStr + ChildTableResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_nosql_table", "test_child", acctest.Required, acctest.Create, childTableRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(childResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(childResourceName, "ddl_statement", childTableDdlStatement),
					resource.TestCheckResourceAttr(childResourceName, "name", "test_table.test_child"),
					resource.TestCheckNoResourceAttr(childResourceName, "table_limits"),
				),
			},

			// verify datasource: test_child
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_tables", "test_child", acctest.Optional, acctest.Create, childTableDataSourceRepresentation) +
					compartmentIdVariableStr + ChildTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(childDataResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(childDataResourceName, "name", "test_table.test_child"),
					resource.TestCheckResourceAttr(childDataResourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(childDataResourceName, "table_collection.#", "1"),
					resource.TestCheckResourceAttrSet(childDataResourceName, "table_collection.0.id"),
					resource.TestCheckResourceAttr(childDataResourceName, "table_collection.0.name", "test_table.test_child"),
					resource.TestCheckResourceAttr(childDataResourceName, "table_collection.0.state", "ACTIVE"),
					resource.TestCheckNoResourceAttr(childDataResourceName, "table_limits"),
				),
			},

			// verify singular datasource: test_child
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_nosql_table", "test_child", acctest.Required, acctest.Create, childTableSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ChildTableResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularChildDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularChildDatasourceName, "table_name_or_id"),
					resource.TestCheckResourceAttrSet(singularChildDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularChildDatasourceName, "name", "test_table.test_child"),
					resource.TestCheckResourceAttr(singularChildDatasourceName, "ddl_statement", childTableDdlStatement),
					resource.TestCheckResourceAttr(singularChildDatasourceName, "schema.#", "1"),
					resource.TestCheckResourceAttr(singularChildDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularChildDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularChildDatasourceName, "time_updated"),
					resource.TestCheckNoResourceAttr(singularChildDatasourceName, "table_limits"),
				),
			},

			// remove the child table resource
			{
				Config: config + compartmentIdVariableStr + ChildTableResourceConfig,
			},
		},
	})
}
