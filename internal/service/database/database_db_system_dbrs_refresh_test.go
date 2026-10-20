// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	common "github.com/oracle/oci-go-sdk/v65/common"
	db "github.com/oracle/oci-go-sdk/v65/database"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TestDatabaseDbSystemSetDataReportsNestedWriteError(t *testing.T) {
	resource := DatabaseDbSystemResource()
	// Simulate an API/schema mismatch to ensure refresh does not silently succeed.
	backupSchema := resource.Schema["db_home"].Elem.(*schema.Resource).Schema["database"].Elem.(*schema.Resource).Schema["db_backup_config"].Elem.(*schema.Resource).Schema["backup_destination_details"].Elem.(*schema.Resource).Schema
	delete(backupSchema, "vpc_user")
	crud := &DatabaseDbSystemResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: resource.Data(nil)},
		Res:      &db.DbSystem{Id: common.String("S1")},
		DbHome:   &db.DbHome{Id: common.String("H1"), DbVersion: common.String("23.26.3.0.0")},
		Database: &db.Database{DbBackupConfig: &db.DbBackupConfig{BackupDestinationDetails: []db.BackupDestinationDetails{{VpcUser: common.String("test-vpc-user")}}}},
	}
	if err := crud.SetData(); err == nil || !strings.Contains(err.Error(), "vpc_user") {
		t.Fatalf("expected nested state-write error mentioning vpc_user, got %v", err)
	}
}

// Exercise the complete nested state write using post-upgrade API models.
// The DB home and database identifiers are unchanged by the upgrade.
func TestDatabaseDbSystemSetDataAfterUpgradeWithBackupDestination(t *testing.T) {
	for _, destinationType := range []string{"OBJECT_STORE", "DBRS"} {
		t.Run(destinationType, func(t *testing.T) {
			resource := DatabaseDbSystemResource()
			seed := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
				"db_home": []interface{}{map[string]interface{}{
					"id": "H1", "db_version": "19.32.0.0.0", "db_home_location": "/oracle/19/dbhome_1",
					"database": []interface{}{map[string]interface{}{"id": "D1", "db_name": "TEST", "admin_password": "TestOnlyPassword1!"}},
				}},
			})
			seed.SetId("S1")
			data := resource.Data(seed.State())
			destination := db.BackupDestinationDetails{Type: db.BackupDestinationDetailsTypeEnum(destinationType)}
			if destinationType == "DBRS" {
				destination.VpcUser = common.String("test-vpc-user")
			}
			crud := &DatabaseDbSystemResourceCrud{
				BaseCrud: tfresource.BaseCrud{D: data},
				Res:      &db.DbSystem{Id: common.String("S1"), Version: common.String("23.26.3.0.0")},
				DbHome:   &db.DbHome{Id: common.String("H1"), DbVersion: common.String("23.26.3.0.0"), DbHomeLocation: common.String("/oracle/23/dbhome_2")},
				Database: &db.Database{Id: common.String("D1"), DbName: common.String("TEST"), DbBackupConfig: &db.DbBackupConfig{BackupDestinationDetails: []db.BackupDestinationDetails{destination}}},
			}
			if err := crud.SetData(); err != nil {
				t.Fatalf("SetData: %v", err)
			}
			expected := map[string]string{
				"version": "23.26.3.0.0", "db_home.0.id": "H1", "db_home.0.db_version": "23.26.3.0.0",
				"db_home.0.db_home_location": "/oracle/23/dbhome_2", "db_home.0.database.0.id": "D1",
				"db_home.0.database.0.db_backup_config.0.backup_destination_details.0.type": destinationType,
			}
			if destinationType == "DBRS" {
				expected["db_home.0.database.0.db_backup_config.0.backup_destination_details.0.vpc_user"] = "test-vpc-user"
			}
			for key, want := range expected {
				if got := data.Get(key); got != want {
					t.Errorf("%s = %v; want %s", key, got, want)
				}
			}
		})
	}
}
