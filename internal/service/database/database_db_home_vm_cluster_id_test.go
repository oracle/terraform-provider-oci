// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	common "github.com/oracle/oci-go-sdk/v65/common"
	db "github.com/oracle/oci-go-sdk/v65/database"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// A DB home created with source VM_CLUSTER_NEW may also set database.vm_cluster_id.
// The value is ForceNew, so it must be read back from the Database model on refresh.
func TestDatabaseDbHomeSetDataKeepsNestedVmClusterId(t *testing.T) {
	const vmClusterId = "ocid1.vmcluster.oc1..example"
	config := map[string]interface{}{
		"source": "VM_CLUSTER_NEW", "vm_cluster_id": vmClusterId, "db_version": "23.26.3.0.0",
		"database": []interface{}{map[string]interface{}{"admin_password": "TestOnlyPassword1!", "db_name": "TEST", "vm_cluster_id": vmClusterId}},
	}
	resource := DatabaseDbHomeResource()
	seed := schema.TestResourceDataRaw(t, resource.Schema, config)
	seed.SetId("H1")
	data := resource.Data(seed.State())
	crud := &DatabaseDbHomeResourceCrud{
		BaseCrud: tfresource.BaseCrud{D: data},
		Res:      &db.DbHome{Id: common.String("H1"), VmClusterId: common.String(vmClusterId), DbVersion: common.String("23.26.3.0.0")},
		Database: &db.Database{Id: common.String("D1"), DbName: common.String("TEST"), VmClusterId: common.String(vmClusterId)},
	}
	if err := crud.SetData(); err != nil {
		t.Fatalf("SetData: %v", err)
	}
	state := data.State()
	if got := state.Attributes["database.0.vm_cluster_id"]; got != vmClusterId {
		t.Errorf("state database.0.vm_cluster_id = %q; want %s", got, vmClusterId)
	}

	diff, err := resource.Diff(context.Background(), state, terraform.NewResourceConfigRaw(config), nil)
	if err != nil {
		t.Fatalf("Diff: %v", err)
	}
	if diff != nil {
		for key, attr := range diff.Attributes {
			if attr.RequiresNew {
				t.Errorf("unchanged configuration replaces the DB home: %s %q => %q", key, attr.Old, attr.New)
			}
		}
	}
}
