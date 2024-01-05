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
	RecoveryprotectedDatabaseFetchConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"protected_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_recovery_protected_database.test_protected_database.id}`},
		"base64_encode_content": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"configuration_type":    acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
	}

	RecoveryProtectedDatabaseFetchConfigurationResourceConfig = RecoveryProtectedDatabaseRequiredOnlyResource
)

// issue-routing-tag: recovery/default
func TestRecoveryProtectedDatabaseFetchConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRecoveryProtectedDatabaseFetchConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_recovery_protected_database_fetch_configuration.test_protected_database_fetch_configuration"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_protected_database_fetch_configuration", "test_protected_database_fetch_configuration", acctest.Required, acctest.Create, RecoveryprotectedDatabaseFetchConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryProtectedDatabaseFetchConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "protected_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
			),
		},

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_recovery_protected_database_fetch_configuration", "test_protected_database_fetch_configuration", acctest.Optional, acctest.Create, RecoveryprotectedDatabaseFetchConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RecoveryProtectedDatabaseFetchConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_type", "ALL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "protected_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
			),
		},
	})
}
