// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	MysqlMysqlBackupCopyRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup_copy", acctest.Required, acctest.Create, MysqlMysqlBackupCopyWithSourceDetailsRepresentation)

	MysqlMysqlBackupCopyWithSourceDetailsRepresentation = map[string]interface{}{
		"source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlMysqlBackupCopySourceDetailsRepresentation},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
	}

	MysqlMysqlBackupCopySourceDetailsRepresentation = map[string]interface{}{}
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlBackupResource_copy(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlBackupResource_crossRegionCopy")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceNameCopy := "oci_mysql_mysql_backup.test_mysql_backup_copy"
	datasourceNameCopy := "data.oci_mysql_mysql_backups.test_mysql_backups"

	if utils.GetEnvSettingWithBlankDefault("source_region") == "" {
		t.Skip("Skipping TestMysqlMysqlBackupResource_copy test because there is no source region specified")
	}

	err := createSourceBackupToCopy()
	if err != nil {
		t.Fatalf("Unable to Create source DB System and Backup to copy. Error: %v", err)
	}

	MysqlMysqlBackupCopySourceDetailsRepresentation = map[string]interface{}{
		"region":         acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("source_region")},
		"backup_id":      acctest.Representation{RepType: acctest.Required, Create: mysqlBackupId},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: compartmentId},
	}

	MysqlMysqlBackupCopyWithSourceDetailsRepresentation = acctest.GetUpdatedRepresentationCopy("source_details", acctest.RepresentationGroup{RepType: acctest.Required, Group: MysqlMysqlBackupCopySourceDetailsRepresentation}, MysqlMysqlBackupCopyWithSourceDetailsRepresentation)

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckMysqlMysqlBackupDestroy, []resource.TestStep{

		// verify backup copy
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup_copy", acctest.Required, acctest.Create, MysqlMysqlBackupCopyWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "db_system_id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "immediate_source_backup_id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "original_source_backup_id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "time_copy_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameCopy, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},

		// verify backup copy with optionals
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup_copy", acctest.Optional, acctest.Create, MysqlMysqlBackupCopyWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceNameCopy, "db_system_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "description", "description"),
				resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "immediate_source_backup_id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "original_source_backup_id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "time_copy_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceNameCopy, "id")
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup_copy", acctest.Optional, acctest.Update, MysqlMysqlBackupCopyWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNameCopy, "backup_type", "INCREMENTAL"),
				resource.TestCheckResourceAttr(resourceNameCopy, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "creation_type"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "db_system_id"),
				resource.TestCheckResourceAttr(resourceNameCopy, "description", "description2"),
				resource.TestCheckResourceAttr(resourceNameCopy, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "id"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "state"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "time_created"),
				resource.TestCheckResourceAttrSet(resourceNameCopy, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceNameCopy, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_backups", "test_mysql_backups", acctest.Optional, acctest.Update, map[string]interface{}{
					"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
					"backup_id":      acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_backup.test_mysql_backup_copy.id}`},
				}) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_backup", "test_mysql_backup_copy", acctest.Optional, acctest.Update, MysqlMysqlBackupCopyWithSourceDetailsRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backup_id"),
				resource.TestCheckResourceAttr(datasourceNameCopy, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceNameCopy, "backups.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.backup_size_in_gbs"),
				resource.TestCheckResourceAttr(datasourceNameCopy, "backups.0.backup_type", "INCREMENTAL"),
				resource.TestCheckResourceAttr(datasourceNameCopy, "backups.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.creation_type"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.data_storage_size_in_gb"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.db_system_id"),
				resource.TestCheckResourceAttr(datasourceNameCopy, "backups.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceNameCopy, "backups.0.display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.id"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.immediate_source_backup_id"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.mysql_version"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.original_source_backup_id"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.shape_name"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.state"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.time_copy_created"),
				resource.TestCheckResourceAttrSet(datasourceNameCopy, "backups.0.time_created"),
			),
		},

		// verify resource import
		{
			Config:            config + MysqlMysqlBackupCopyRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"source_details",
				"immediate_source_backup_id",
				"original_source_backup_id",
				"time_copy_created",
			},
			ResourceName: resourceNameCopy,
		},
	})
}

func createSourceBackupToCopy() error {
	sourceRegion := utils.GetEnvSettingWithBlankDefault("source_region")

	var err error

	mysqlDbSystemId, err = createDbSystemInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion)
	if err != nil {
		log.Printf("[WARN] failed to createDbSystemInRegion with the error %v", err)
		return err
	}

	mysqlBackupId, err = createBackupInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, &mysqlDbSystemId)
	if err != nil {
		log.Printf("[WARN] failed to createBackupInRegion with the error %v", err)
		return err
	}

	return nil
}

func deleteCopyBackupSource() {
	sourceRegion := utils.GetEnvSettingWithBlankDefault("source_region")

	var err error
	err = deleteBackupInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, mysqlBackupId)
	if err != nil {
		log.Printf("[WARN] failed to deleteBackupInRegion with error %v", err)
	}

	err = deleteDbSystemInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, mysqlDbSystemId)
	if err != nil {
		log.Printf("[WARN] failed to deleteDbSystemInRegion with error %v", err)
	}
}
