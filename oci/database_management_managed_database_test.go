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
	managedDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": Representation{repType: Required, create: `${var.tenancy_ocid}testManagedDatabase0`},
	}

	managedDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":    Representation{repType: Required, create: `${var.compartment_id}`},
		"deployment_type":   Representation{repType: Optional, create: `ONPREMISE`},
		"id":                Representation{repType: Optional, create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"management_option": Representation{repType: Optional, create: `BASIC`},
		"name":              Representation{repType: Optional, create: `name`},
	}

	ManagedDatabaseResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_databases.test_managed_databases"
	singularDatasourceName := "data.oci_database_management_managed_database.test_managed_database"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", Required, Create, managedDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_collection.0.items.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_collection.0.items.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_collection.0.items.0.deployment_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_collection.0.items.0.management_option"),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_database_management_managed_database", "test_managed_database", Required, Create, managedDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_sub_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_cluster"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_groups.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_option"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
