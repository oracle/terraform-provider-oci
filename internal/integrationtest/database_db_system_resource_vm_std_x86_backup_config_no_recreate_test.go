// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DbSystemVmStdx86BackupConfigNoRecreateCreateBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type":                      acctest.Representation{RepType: acctest.Optional, Create: `DBRS`},
		"dbrs_policy_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.protection_policy_id}`},
		"is_zero_data_loss_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	DbSystemVmStdx86BackupConfigNoRecreateCreateBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":        acctest.Representation{RepType: acctest.Required, Create: `true`},
		"auto_full_backup_day":       acctest.Representation{RepType: acctest.Optional, Create: `SUNDAY`},
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DbSystemVmStdx86BackupConfigNoRecreateCreateBackupDestinationDetailsRepresentation},
	}

	DbSystemVmStdx86BackupConfigNoRecreateUpdateBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type":                      acctest.Representation{RepType: acctest.Optional, Create: `DBRS`},
		"dbrs_policy_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.protection_policy_id}`},
		"is_zero_data_loss_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
	}

	DbSystemVmStdx86BackupConfigNoRecreateUpdateBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":        acctest.Representation{RepType: acctest.Required, Create: `true`},
		"auto_full_backup_day":       acctest.Representation{RepType: acctest.Optional, Create: `SUNDAY`},
		"backup_deletion_policy":     acctest.Representation{RepType: acctest.Optional, Create: `DELETE_IMMEDIATELY`, Update: `DELETE_AFTER_RETENTION_PERIOD`},
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DbSystemVmStdx86BackupConfigNoRecreateUpdateBackupDestinationDetailsRepresentation},
	}

	DbSystemVmStdx86BackupConfigNoRecreateCreateRepresentation = acctest.GetUpdatedRepresentationCopy(
		"db_home.database.db_backup_config",
		acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemVmStdx86BackupConfigNoRecreateCreateBackupConfigRepresentation},
		acctest.RepresentationCopyWithNewProperties(DbSystemVmStdx86Representation, map[string]interface{}{
			"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		}),
	)

	DbSystemVmStdx86BackupConfigNoRecreateUpdateRepresentation = acctest.GetUpdatedRepresentationCopy(
		"db_home.database.db_backup_config",
		acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemVmStdx86BackupConfigNoRecreateUpdateBackupConfigRepresentation},
		acctest.RepresentationCopyWithNewProperties(DbSystemVmStdx86BackupConfigNoRecreateCreateRepresentation, map[string]interface{}{
			"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Admin"}},
		}),
	)
)

func TestResourceDatabaseDBSystemVMStdx86BackupConfigNoRecreateUpdate(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDBSystemVMStdx86BackupConfigNoRecreateUpdate")
	defer httpreplay.SaveScenario()

	const recoveryServiceSubnetWaitDuration = time.Duration(2 * time.Minute)

	protectionPolicyId := utils.GetEnvSettingWithBlankDefault("protection_policy_id")
	protectionPolicyIdVariableStr := fmt.Sprintf("variable \"protection_policy_id\" { default = \"%s\" }\n", protectionPolicyId)

	resourceName := "oci_database_db_system.test_vm_std_x86_db_system"

	var resId, resId2 string

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: ResourceDatabaseDbrsBaseConfig + protectionPolicyIdVariableStr,
			Check:  acctest.ComposeAggregateTestCheckFuncWrapper(),
		},
		{
			PreConfig: func() {
				time.Sleep(recoveryServiceSubnetWaitDuration)
			},
			Config: ResourceDatabaseDbrsBaseConfig + protectionPolicyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_vm_std_x86_db_system", acctest.Optional, acctest.Create, DbSystemVmStdx86BackupConfigNoRecreateCreateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "8"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.auto_full_backup_day", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.dbrs_policy_id", protectionPolicyId),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.type", "DBRS"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.is_zero_data_loss_enabled", "true"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		{
			Config: ResourceDatabaseDbrsBaseConfig + protectionPolicyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_vm_std_x86_db_system", acctest.Optional, acctest.Update, DbSystemVmStdx86BackupConfigNoRecreateUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Admin"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.auto_full_backup_day", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_deletion_policy", "DELETE_AFTER_RETENTION_PERIOD"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.dbrs_policy_id", protectionPolicyId),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.type", "DBRS"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.is_zero_data_loss_enabled", "false"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if err != nil {
						return err
					}
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return nil
				},
			),
		},
	})
}
