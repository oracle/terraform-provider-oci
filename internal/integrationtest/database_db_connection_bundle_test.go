// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	DatabaseDbConnectionBundleSingularDataSourceRepresentation = map[string]interface{}{
		"db_connection_bundle_id": acctest.Representation{RepType: acctest.Required, Create: `${var.db_connection_bundle_id}`},
	}

	DatabaseDbConnectionBundleDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"associated_resource_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.associated_resource_id}`},
		"db_connection_bundle_type": acctest.Representation{RepType: acctest.Optional, Create: `TLS`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}
)

// issue-routing-tag: database/default
func TestDatabaseDbConnectionBundleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbConnectionBundleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	dbConnectionBundleId := utils.GetEnvSettingWithBlankDefault("db_connection_bundle_ocid")
	dbConnectionBundleIdVariableStr := fmt.Sprintf("variable \"db_connection_bundle_id\" { default = \"%s\" }\n", dbConnectionBundleId)

	associatedResourceId := utils.GetEnvSettingWithBlankDefault("associated_resource_ocid")
	associatedResourceIdVariableStr := fmt.Sprintf("variable \"associated_resource_id\" { default = \"%s\" }\n", associatedResourceId)

	datasourceName := "data.oci_database_db_connection_bundles.test_db_connection_bundles"
	singularDatasourceName := "data.oci_database_db_connection_bundle.test_db_connection_bundle"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_connection_bundles", "test_db_connection_bundles", acctest.Optional, acctest.Create, DatabaseDbConnectionBundleDataSourceRepresentation) +
				compartmentIdVariableStr + associatedResourceIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "associated_resource_id", associatedResourceId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_connection_bundle_type", "TLS"),
				resource.TestCheckNoResourceAttr(datasourceName, "display_name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "db_connection_bundles.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "db_connection_bundles.0.associated_resource_details.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "db_connection_bundles.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "db_connection_bundles.0.db_connection_bundle_type", "TLS"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_connection_bundles.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_connection_bundles.0.is_protected"),
				resource.TestCheckResourceAttr(datasourceName, "db_connection_bundles.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_connection_bundles.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "db_connection_bundles.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_connection_bundle", "test_db_connection_bundle", acctest.Required, acctest.Create, DatabaseDbConnectionBundleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + dbConnectionBundleIdVariableStr + associatedResourceIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "db_connection_bundle_id", dbConnectionBundleId),

				resource.TestCheckResourceAttr(singularDatasourceName, "associated_resource_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_connection_bundle_type", "TLS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "id", dbConnectionBundleId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_protected"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
