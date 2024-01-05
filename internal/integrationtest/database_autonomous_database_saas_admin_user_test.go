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
	DatabaseSaasAdminUserRepresentation = map[string]interface{}{
		"autonomous_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_database.test_autonomous_database.id}`},
		"password":               acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"duration":               acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `5`},
		"access_type":            acctest.Representation{RepType: acctest.Required, Create: `READ_WRITE`, Update: `READ_ONLY`},
	}
)

func TestDatabaseAutonomousDatabaseSaasAdminUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseSaasAdminUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	okvSecret = utils.GetEnvSettingWithBlankDefault("okv_secret")
	okvSecretVariableStr := fmt.Sprintf("variable \"okv_secret\" { default = \"%s\" }\n", okvSecret)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	adbResourceName := "oci_database_autonomous_database.test_autonomous_database"
	saasAdminUserResourceName := "oci_database_autonomous_database_saas_admin_user.test_saas_admin_user"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousDatabaseResourceDependencies, "database", "autonomousDatabase", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create with password
		{
			Config: config + compartmentIdVariableStr + okvSecretVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_saas_admin_user", "test_saas_admin_user", acctest.Required, acctest.Create, DatabaseSaasAdminUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(adbResourceName, "id"),
				resource.TestCheckResourceAttr(adbResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(adbResourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(adbResourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(adbResourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(adbResourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(saasAdminUserResourceName, "duration", "3"),
				resource.TestCheckResourceAttr(saasAdminUserResourceName, "access_type", "READ_WRITE"),
				resource.TestCheckResourceAttr(saasAdminUserResourceName, "password", "BEstrO0ng_#11"),
			),
		},
		// verify "Update" with secret (should delete and create again)
		{
			Config: config + compartmentIdVariableStr + okvSecretVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database_saas_admin_user", "test_saas_admin_user", acctest.Required, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(DatabaseSaasAdminUserRepresentation, map[string]interface{}{
						"secret_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.okv_secret}`},
						"secret_version_number": acctest.Representation{RepType: acctest.Required, Create: `1`},
					}), []string{"password"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(adbResourceName, "id"),
				resource.TestCheckResourceAttr(adbResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(adbResourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(adbResourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(adbResourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(adbResourceName, "db_workload", "OLTP"),
				resource.TestCheckResourceAttr(saasAdminUserResourceName, "duration", "5"),
				resource.TestCheckResourceAttr(saasAdminUserResourceName, "access_type", "READ_ONLY"),
				resource.TestCheckResourceAttr(saasAdminUserResourceName, "secret_id", okvSecret),
			),
		},
		// verify Destroy
		{
			Config: config + compartmentIdVariableStr + okvSecretVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Required, acctest.Create, DatabaseAutonomousDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(adbResourceName, "id"),
				resource.TestCheckResourceAttr(adbResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(adbResourceName, "cpu_core_count", "1"),
				resource.TestCheckResourceAttr(adbResourceName, "data_storage_size_in_tbs", "1"),
				resource.TestCheckResourceAttr(adbResourceName, "db_name", adbName),
				resource.TestCheckResourceAttr(adbResourceName, "db_workload", "OLTP"),
			),
		},
	})
}
