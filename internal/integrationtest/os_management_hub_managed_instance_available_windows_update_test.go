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
	OsManagementHubManagedInstanceAvailableWindowsUpdateDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
		"classification_type":   acctest.Representation{RepType: acctest.Optional, Create: []string{`SECURITY`}},
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `Security Intelligence Update for Microsoft Defender Antivirus - KB2267602 (Version 1.407.558.0) - Current Channel (Broad)`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `Defender`},
		"is_installable":        acctest.Representation{RepType: acctest.Optional, Create: `INSTALLABLE`},
		"name":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`eee01ae0-0280-4071-9060-ffecf726e03b`}},
	}

	OsManagementHubManagedInstanceAvailableWindowsUpdateResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceWindowsRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceAvailableWindowsUpdateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceAvailableWindowsUpdateResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_managed_instance_available_windows_updates.test_managed_instance_available_windows_updates"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_available_windows_updates", "test_managed_instance_available_windows_updates", acctest.Required, acctest.Create, OsManagementHubManagedInstanceAvailableWindowsUpdateDataSourceRepresentation) +
				OsManagementHubManagedInstanceAvailableWindowsUpdateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "available_windows_update_collection.#"),
			),
		},
		// verify optional datasource
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_available_windows_updates", "test_managed_instance_available_windows_updates", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceAvailableWindowsUpdateDataSourceRepresentation) +
				OsManagementHubManagedInstanceAvailableWindowsUpdateResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "classification_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Security Intelligence Update for Microsoft Defender Antivirus - KB2267602 (Version 1.407.558.0) - Current Channel (Broad)"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "Defender"),
				resource.TestCheckResourceAttr(datasourceName, "is_installable", "INSTALLABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "name.#", "1"),

				resource.TestCheckResourceAttrSet(datasourceName, "available_windows_update_collection.#"),
			),
		},
	})
}
