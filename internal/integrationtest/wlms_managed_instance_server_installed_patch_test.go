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
	wlmsServerId                                                    = utils.GetEnvSettingWithBlankDefault("wlms_server_id")
	WlmsManagedInstanceServerInstalledPatchDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: wlmsManagedInstanceOcid},
		"server_id":           acctest.Representation{RepType: acctest.Required, Create: wlmsServerId},
	}
)

// issue-routing-tag: wlms/default
func TestWlmsManagedInstanceServerInstalledPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWlmsManagedInstanceServerInstalledPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_wlms_managed_instance_server_installed_patches.test_managed_instance_server_installed_patches"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_wlms_managed_instance_server_installed_patches", "test_managed_instance_server_installed_patches", acctest.Required, acctest.Create, WlmsManagedInstanceServerInstalledPatchDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "installed_patch_collection.#"),
			),
		},
	})
}
