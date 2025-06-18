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
	WlmsManagedInstanceServerSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: wlmsManagedInstanceOcid},
		"server_id":           acctest.Representation{RepType: acctest.Required, Create: wlmsServerId},
	}

	WlmsManagedInstanceServerDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: wlmsManagedInstanceOcid},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: wlmsServerName},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsManagedInstanceServerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsManagedInstanceServerResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_wlms_managed_instance_servers.test_managed_instance_servers"
	singularDatasourceName := "data.oci_wlms_managed_instance_server.test_managed_instance_server"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_managed_instance_servers", "test_managed_instance_servers", acctest.Required, acctest.Create, WlmsManagedInstanceServerDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "server_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_managed_instance_server", "test_managed_instance_server", acctest.Required, acctest.Create, WlmsManagedInstanceServerSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(singularDatasourceName, "host_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_admin"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jdk_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "jdk_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "latest_patches_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "middleware_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "middleware_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_readiness_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "restart_order"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "weblogic_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "wls_domain_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "wls_domain_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "wls_domain_path"),
			),
		},
	})
}
