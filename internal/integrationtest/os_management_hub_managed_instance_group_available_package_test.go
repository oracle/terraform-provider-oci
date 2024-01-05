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
	OsManagementHubManagedInstanceGroupAvailablePackageDataSourceRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: []string{`zsh`}},
		"display_name_contains":     acctest.Representation{RepType: acctest.Optional, Create: `zsh`},
		"is_latest":                 acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	OsManagementHubManagedInstanceGroupAvailablePackageResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceGroupAvailablePackageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceGroupAvailablePackageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_managed_instance_group_available_packages.test_managed_instance_group_available_packages"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubManagedInstanceGroupAvailablePackageResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_available_packages", "test_managed_instance_group_available_packages", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupAvailablePackageDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "zsh"),
				resource.TestCheckResourceAttr(datasourceName, "is_latest", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_available_package_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "managed_instance_group_available_package_collection.0.items.#", "1"),
			),
		},
	})
}
