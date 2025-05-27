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
	WlmsWlsDomainServerSingularDataSourceRepresentation = map[string]interface{}{
		"server_id":     acctest.Representation{RepType: acctest.Required, Create: wlmsServerId},
		"wls_domain_id": acctest.Representation{RepType: acctest.Required, Create: wlsDomainOcid},
	}

	WlmsWlsDomainServerDataSourceRepresentation = map[string]interface{}{
		"wls_domain_id": acctest.Representation{RepType: acctest.Required, Create: wlsDomainOcid},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: wlmsServerName},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsWlsDomainServerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsWlsDomainServerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_wlms_wls_domain_servers.test_wls_domain_servers"
	singularDatasourceName := "data.oci_wlms_wls_domain_server.test_wls_domain_server"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_wls_domain_servers", "test_wls_domain_servers", acctest.Required, acctest.Create, WlmsWlsDomainServerDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "server_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_wls_domain_server", "test_wls_domain_server", acctest.Required, acctest.Create, WlmsWlsDomainServerSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "server_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "wls_domain_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_admin"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jdk_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jdk_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "latest_patches_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "middleware_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "middleware_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_readiness_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "restart_order"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "weblogic_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "wls_domain_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "wls_domain_path"),
			),
		},
	})
}
