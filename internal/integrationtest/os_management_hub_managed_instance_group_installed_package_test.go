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
	OsManagementHubManagedInstanceGroupInstalledPackageDataSourceRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("managed_instance_group_ocid")},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: []string{`ed`}},
		"display_name_contains":     acctest.Representation{RepType: acctest.Optional, Create: `ed`},
		"time_install_date_end":     acctest.Representation{RepType: acctest.Optional, Create: `2006-01-02T15:04:05Z`},
		"time_install_date_start":   acctest.Representation{RepType: acctest.Optional, Create: `2006-01-02T15:04:05Z`},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceGroupInstalledPackageResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceGroupInstalledPackageResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_managed_instance_group_installed_packages.test_managed_instance_group_installed_packages"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_installed_packages", "test_managed_instance_group_installed_packages", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupInstalledPackageDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "ed"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_install_date_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_install_date_start"),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_installed_package_collection.#"),
			),
		},
	})
}
