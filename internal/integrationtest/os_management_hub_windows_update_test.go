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
	OsManagementHubWindowsUpdateSingularDataSourceRepresentation = map[string]interface{}{
		"windows_update_id": acctest.Representation{RepType: acctest.Required, Create: `f30b67d6-f749-4a7b-a1e9-d84039c8b35c`},
	}

	OsManagementHubWindowsUpdateDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"classification_type":   acctest.Representation{RepType: acctest.Optional, Create: []string{`SECURITY`}},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `Defender`},
		"name":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`f30b67d6-f749-4a7b-a1e9-d84039c8b35c`}},
	}

	OsManagementHubWindowsUpdateToInstallDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: []string{utils.GetEnvSettingWithBlankDefault("osmh_windows_update_to_install")}},
	}

	OsManagementHubWindowsUpdateResourceConfig = ""
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubWindowsUpdateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubWindowsUpdateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_windows_updates.test_windows_updates"
	singularDatasourceName := "data.oci_os_management_hub_windows_update.test_windows_update"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_windows_updates", "test_windows_updates", acctest.Required, acctest.Create, OsManagementHubWindowsUpdateDataSourceRepresentation) +
				OsManagementHubWindowsUpdateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "windows_update_collection.#"),
			),
		},
		// verify datasource with optional query parameters
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_windows_updates", "test_windows_updates", acctest.Optional, acctest.Create, OsManagementHubWindowsUpdateDataSourceRepresentation) +
				OsManagementHubWindowsUpdateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "classification_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "Defender"),
				resource.TestCheckResourceAttr(datasourceName, "name.#", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "windows_update_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_windows_update", "test_windows_update", acctest.Required, acctest.Create, OsManagementHubWindowsUpdateSingularDataSourceRepresentation) +
				OsManagementHubWindowsUpdateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "windows_update_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "installable"),
				resource.TestCheckResourceAttr(singularDatasourceName, "installation_requirements.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_reboot_required_for_installation"),
				resource.TestCheckResourceAttr(singularDatasourceName, "kb_article_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_bytes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "update_type"),
			),
		},
	})
}
