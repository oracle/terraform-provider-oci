// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseDbSystemPatchDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_system.test_db_system.id}`},
	}

	// 1. Main Db System Resource Representation: Start
	DbSystemResourceBaseRepresentation = map[string]interface{}{
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDbSystem`},
		"database_edition":        acctest.Representation{RepType: acctest.Optional, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Optional, Create: `NORMAL`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Optional, Create: `4`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Optional, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`, Update: `BRING_YOUR_OWN_LICENSE`},
		"node_count":              acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"fault_domains":           acctest.Representation{RepType: acctest.Optional, Create: []string{`FAULT-DOMAIN-1`}},
		"domain":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`}},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `tfOracleDb`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemBaseDbHomeGroup},
	}

	DbSystemBaseDbHomeGroup = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `tfDbHome`},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemBaseDatabaseGroup},
	}

	DbSystemBaseDatabaseGroup = map[string]interface{}{
		"db_name":            acctest.Representation{RepType: acctest.Optional, Create: `tfDb`},
		"pdb_name":           acctest.Representation{RepType: acctest.Optional, Create: `tfPdb`},
		"character_set":      acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"ncharacter_set":     acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"db_workload":        acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"kms_key_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_id}`},
		"kms_key_version_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.kms_key_version_id}`},
		"vault_id":           acctest.Representation{RepType: acctest.Optional, Create: `${var.vault_id}`},
		"admin_password":     acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
	}
)

// issue-routing-tag: database/default
func TestDatabaseDbSystemPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbSystemPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.BaseDBProviderTestConfig()

	datasourceName := "data.oci_database_db_system_patches.test_db_system_patches"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_system_patches", "test_db_system_patches", acctest.Optional, acctest.Create, DatabaseDatabaseDbSystemPatchDataSourceRepresentation) +
				DbSystemBaseConfig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_db_system", acctest.Optional, acctest.Create, DbSystemResourceBaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.time_released"),
				resource.TestCheckResourceAttrSet(datasourceName, "patches.0.version"),
			),
		},
	})
}
