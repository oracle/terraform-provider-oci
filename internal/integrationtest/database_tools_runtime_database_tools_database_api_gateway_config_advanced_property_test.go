// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertyDataSourceRepresentation = map[string]interface{}{}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertyResourceConfig = ""
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_tools_runtime_database_tools_database_api_gateway_config_advanced_properties.test_database_tools_database_api_gateway_config_advanced_properties"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_advanced_properties", "test_database_tools_database_api_gateway_config_advanced_properties", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertyDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_database_api_gateway_config_advanced_property_summary_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_database_api_gateway_config_advanced_property_summary_collection.0.items.#"),
			),
		},
	})
}
