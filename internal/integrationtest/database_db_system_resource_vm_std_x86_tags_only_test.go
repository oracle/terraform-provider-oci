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
	DbSystemVmStdx86FreeformTagsOnlyBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type":                      acctest.Representation{RepType: acctest.Optional, Create: `DBRS`},
		"dbrs_policy_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.protection_policy_id}`},
		"is_zero_data_loss_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	DbSystemVmStdx86FreeformTagsOnlyBackupConfigRepresentation = map[string]interface{}{
		"auto_backup_enabled":        acctest.Representation{RepType: acctest.Required, Create: `true`},
		"auto_full_backup_day":       acctest.Representation{RepType: acctest.Optional, Create: `SUNDAY`},
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DbSystemVmStdx86FreeformTagsOnlyBackupDestinationDetailsRepresentation},
	}

	DbSystemVmStdx86FreeformTagsOnlyCreateRepresentation = acctest.GetUpdatedRepresentationCopy(
		"db_home.database.db_backup_config",
		acctest.RepresentationGroup{RepType: acctest.Required, Group: DbSystemVmStdx86FreeformTagsOnlyBackupConfigRepresentation},
		acctest.RepresentationCopyWithNewProperties(DbSystemVmStdx86Representation, map[string]interface{}{
			"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		}),
	)
	DbSystemVmStdx86FreeformTagsOnlyUpdateRepresentation = acctest.RepresentationCopyWithNewProperties(
		DbSystemVmStdx86FreeformTagsOnlyCreateRepresentation,
		map[string]interface{}{
			"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Admin"}},
		},
	)
)

func TestResourceDatabaseDBSystemVMStdx86FreeformTagsOnlyUpdate(t *testing.T) {
	httpreplay.SetScenario("TestResourceDatabaseDBSystemVMStdx86FreeformTagsOnlyUpdate")
	defer httpreplay.SaveScenario()

	const recoveryServiceSubnetWaitDuration = time.Duration(2 * time.Minute)

	backupConfigAttributes := []string{
		"db_home.0.database.0.db_backup_config.0.auto_backup_enabled",
		"db_home.0.database.0.db_backup_config.0.auto_backup_window",
		"db_home.0.database.0.db_backup_config.0.auto_full_backup_day",
		"db_home.0.database.0.db_backup_config.0.backup_destination_details.#",
		"db_home.0.database.0.db_backup_config.0.backup_destination_details.0.dbrs_policy_id",
		"db_home.0.database.0.db_backup_config.0.backup_destination_details.0.type",
		"db_home.0.database.0.db_backup_config.0.backup_destination_details.0.is_zero_data_loss_enabled",
	}
	backupConfigBeforeTagUpdate := map[string]string{}

	protectionPolicyId := utils.GetEnvSettingWithBlankDefault("protection_policy_id")
	protectionPolicyIdVariableStr := fmt.Sprintf("variable \"protection_policy_id\" { default = \"%s\" }\n", protectionPolicyId)

	resourceName := "oci_database_db_system.test_vm_std_x86_db_system"

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
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_vm_std_x86_db_system", acctest.Optional, acctest.Create, DbSystemVmStdx86FreeformTagsOnlyCreateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "8"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.auto_backup_window", ""),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.auto_full_backup_day", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.dbrs_policy_id", protectionPolicyId),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.type", "DBRS"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.is_zero_data_loss_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				testAccCaptureResourceAttributes(resourceName, backupConfigAttributes, backupConfigBeforeTagUpdate),
			),
		},
		{
			Config: ResourceDatabaseDbrsBaseConfig + protectionPolicyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_system", "test_vm_std_x86_db_system", acctest.Optional, acctest.Update, DbSystemVmStdx86FreeformTagsOnlyUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "compute_count", "4"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Admin"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.auto_backup_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.auto_backup_window", ""),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.auto_full_backup_day", "SUNDAY"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.dbrs_policy_id", protectionPolicyId),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.type", "DBRS"),
				resource.TestCheckResourceAttr(resourceName, "db_home.0.database.0.db_backup_config.0.backup_destination_details.0.is_zero_data_loss_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				testAccCheckResourceAttributesUnchanged(resourceName, backupConfigAttributes, backupConfigBeforeTagUpdate),
			),
		},
	})
}

func testAccCaptureResourceAttributes(name string, keys []string, values map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, key := range keys {
			value, err := acctest.FromInstanceState(s, name, key)
			if err != nil {
				return err
			}

			values[key] = value
		}

		return nil
	}
}

func testAccCheckResourceAttributesUnchanged(name string, keys []string, expectedValues map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, key := range keys {
			expectedValue, ok := expectedValues[key]
			if !ok {
				return fmt.Errorf("missing expected value for attribute %q", key)
			}

			currentValue, err := acctest.FromInstanceState(s, name, key)
			if err != nil {
				return err
			}

			if currentValue != expectedValue {
				return fmt.Errorf("attribute %q changed after freeform tags update: expected %q, got %q", key, expectedValue, currentValue)
			}
		}

		return nil
	}
}
