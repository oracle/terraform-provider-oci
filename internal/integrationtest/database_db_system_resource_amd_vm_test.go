// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	DbSystemAmdRepresentation = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"database_edition":        acctest.Representation{RepType: acctest.Required, Create: `ENTERPRISE_EDITION`},
		"disk_redundancy":         acctest.Representation{RepType: acctest.Required, Create: `NORMAL`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"cpu_core_count":          acctest.Representation{RepType: acctest.Required, Create: `2`},
		"ssh_public_keys":         acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCBDM0G21Tc6IOp6H5fwUVhVcxDxbwRwb9I53lXDdfqytw/pRAfXxDAzlw1jMEWofoVxTVDyqxcEg5yg4ImKFYHIDrZuU9eHv5SoHYJvI9r+Dqm9z52MmEyoTuC4dUyOs79V0oER5vLcjoMQIqmGSKMSlIMoFV2d+AV//RhJSpRPWGQ6lAVPYAiaVk3EzYacayetk1ZCEnMGPV0OV1UWqovm3aAGDozs7+9Isq44HEMyJwdBTYmBu3F8OA8gss2xkwaBgK3EQjCJIRBgczDwioT7RF5WG3IkwKsDTl2bV0p5f5SeX0U8SGHnni9uNoc9wPAWaleZr3Jcp1yIcRFR9YV`}},
		"domain":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.subnet_domain_name}`},
		"hostname":                acctest.Representation{RepType: acctest.Required, Create: `myOracleDB`},
		"data_storage_size_in_gb": acctest.Representation{RepType: acctest.Required, Create: `256`},
		"license_model":           acctest.Representation{RepType: acctest.Required, Create: `LICENSE_INCLUDED`},
		"node_count":              acctest.Representation{RepType: acctest.Required, Create: `1`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `tfDbSystemAmd`},
		"db_home":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemAmdDbHomeGroup},
		"kms_key_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"kms_key_version_id":      acctest.Representation{RepType: acctest.Required, Update: `${var.kms_key_version_id}`},
	}

	DbSystemAmdDbHomeGroup = map[string]interface{}{
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `tfDbHome`},
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemAmdDatabaseGroup},
	}

	DbSystemAmdDatabaseGroup = map[string]interface{}{
		"admin_password":   acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: nil},
		"db_name":          acctest.Representation{RepType: acctest.Required, Create: `tfDb`},
		"character_set":    acctest.Representation{RepType: acctest.Required, Create: `AL32UTF8`},
		"ncharacter_set":   acctest.Representation{RepType: acctest.Required, Create: `AL16UTF16`},
		"db_workload":      acctest.Representation{RepType: acctest.Required, Create: `OLTP`},
		"pdb_name":         acctest.Representation{RepType: acctest.Required, Create: `tfPdb`},
		"db_backup_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemAmdDbBackupConfigGroup},
		"kms_key_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"vault_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.vault_id}`},
	}

	DbSystemAmdDbBackupConfigGroup = map[string]interface{}{
		"auto_backup_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
	}
)

func TestResourceDatabaseDBSystemAmdVM(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDBSystemAmdVM")
	defer httpreplay.SaveScenario()

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsKeyVersionId := utils.GetEnvSettingWithBlankDefault("kms_key_version_id")
	kmsKeyVersionIdVariableStr := fmt.Sprintf("variable \"kms_key_version_id\" { default = \"%s\" }\n", kmsKeyVersionId)

	vaultId := utils.GetEnvSettingWithBlankDefault("vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	resourceName := "oci_database_db_system.test_amd_db_system"

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify create
		{
			Config: ResourceDatabaseBaseConfig + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_amd_db_system", acctest.Optional, acctest.Create, DbSystemAmdRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "storage_volume_performance_mode", "HIGH_PERFORMANCE"),
				resource.TestCheckResourceAttr(resourceName, "disk_redundancy", "NORMAL"),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard.E4.Flex"),
				resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
				resource.TestCheckResourceAttr(resourceName, "ssh_public_keys.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_gb", "256"),
				resource.TestCheckResourceAttr(resourceName, "license_model", "LICENSE_INCLUDED"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home.0.db_version"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home.0.display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "kms_key_id"),
			),
		},
		// verify update
		{
			Config: ResourceDatabaseBaseConfig + kmsKeyIdVariableStr + vaultIdVariableStr + kmsKeyVersionIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_amd_db_system", acctest.Optional, acctest.Update, DbSystemAmdRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// DB System Resource tests
				resource.TestCheckResourceAttr(resourceName, "kms_key_version_id", kmsKeyVersionId),
			),
		},
	})
}
