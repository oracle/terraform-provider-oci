// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	wlmsBackupId = utils.GetEnvSettingWithBlankDefault("wlms_backup_id")

	WlmsWlsDomainServerBackupContentSingularDataSourceRepresentation = map[string]interface{}{
		"backup_id":     acctest.Representation{RepType: acctest.Required, Create: wlmsBackupId},
		"server_id":     acctest.Representation{RepType: acctest.Required, Create: wlmsServerId},
		"wls_domain_id": acctest.Representation{RepType: acctest.Required, Create: wlsDomainOcid},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsWlsDomainServerBackupContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsWlsDomainServerBackupContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	singularDatasourceName := "data.oci_wlms_wls_domain_server_backup_content.test_wls_domain_server_backup_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_wls_domain_server_backup_content", "test_wls_domain_server_backup_content", acctest.Required, acctest.Create, WlmsWlsDomainServerBackupContentSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "middleware.#"),
			),
		},
	})
}
