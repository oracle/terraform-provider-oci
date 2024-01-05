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
	OsManagementHubManagedInstanceGroupAvailableModuleDataSourceRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"name_contains":             acctest.Representation{RepType: acctest.Optional, Create: `nameContains`},
	}

	OsManagementHubManagedInstanceGroupAvailableModuleResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceGroupAvailableModuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceGroupAvailableModuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_managed_instance_group_available_modules.test_managed_instance_group_available_modules"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubManagedInstanceGroupAvailableModuleResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_available_modules", "test_managed_instance_group_available_modules", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupAvailableModuleDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_available_module_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_available_module_collection.0.items.#"),
			),
		},
	})
}
