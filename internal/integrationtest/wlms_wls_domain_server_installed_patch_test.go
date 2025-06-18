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
	WlmsWlsDomainServerInstalledPatchDataSourceRepresentation = map[string]interface{}{
		"server_id":     acctest.Representation{RepType: acctest.Required, Create: wlmsServerId},
		"wls_domain_id": acctest.Representation{RepType: acctest.Required, Create: wlsDomainOcid},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsWlsDomainServerInstalledPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsWlsDomainServerInstalledPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_wlms_wls_domain_server_installed_patches.test_wls_domain_server_installed_patches"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_wls_domain_server_installed_patches", "test_wls_domain_server_installed_patches", acctest.Required, acctest.Create, WlmsWlsDomainServerInstalledPatchDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "installed_patch_collection.#"),
			),
		},
	})
}
