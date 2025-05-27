// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	WlmsWlsDomainServerBackupSingularDataSourceRepresentation = map[string]interface{}{
		"backup_id":     acctest.Representation{RepType: acctest.Required, Create: wlmsBackupId},
		"server_id":     acctest.Representation{RepType: acctest.Required, Create: wlmsServerId},
		"wls_domain_id": acctest.Representation{RepType: acctest.Required, Create: wlsDomainOcid},
	}

	WlmsWlsDomainServerBackupDataSourceRepresentation = map[string]interface{}{
		"server_id":     acctest.Representation{RepType: acctest.Required, Create: wlmsServerId},
		"wls_domain_id": acctest.Representation{RepType: acctest.Required, Create: wlsDomainOcid},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsWlsDomainServerBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsWlsDomainServerBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_wlms_wls_domain_server_backups.test_wls_domain_server_backups"
	singularDatasourceName := "data.oci_wlms_wls_domain_server_backup.test_wls_domain_server_backup"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_wls_domain_server_backups", "test_wls_domain_server_backups", acctest.Required, acctest.Create, WlmsWlsDomainServerBackupDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "backup_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_wls_domain_server_backup", "test_wls_domain_server_backup", acctest.Required, acctest.Create, WlmsWlsDomainServerBackupSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_location"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
	})
}
