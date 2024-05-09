// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubManagedInstanceInstalledWindowsUpdateDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `Windows Malicious Software Removal Tool x64 - v5.122 (KB890830)`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `Software`},
		"name":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`2c6db99f-8d39-4580-8474-31c45fb79525`}},
	}

	OsManagementHubManagedInstanceInstalledWindowsUpdateResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceWindowsRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceInstalledWindowsUpdateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceInstalledWindowsUpdateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_managed_instance_installed_windows_updates.test_managed_instance_installed_windows_updates"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify optional datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_installed_windows_updates", "test_managed_instance_installed_windows_updates", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceInstalledWindowsUpdateDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubManagedInstanceInstalledWindowsUpdateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Windows Malicious Software Removal Tool x64 - v5.122 (KB890830)"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "Software"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "name.#", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "installed_windows_update_collection.#"),
			),
		},
		// verify required datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_installed_windows_updates", "test_managed_instance_installed_windows_updates", acctest.Required, acctest.Create, OsManagementHubManagedInstanceInstalledWindowsUpdateDataSourceRepresentation) +
				compartmentIdVariableStr + OsManagementHubManagedInstanceInstalledWindowsUpdateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "installed_windows_update_collection.#"),
			),
		},
	})
}
