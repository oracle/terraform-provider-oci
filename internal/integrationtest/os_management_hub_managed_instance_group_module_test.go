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
	OsManagementHubManagedInstanceGroupModuleDataSourceRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: utils.GetEnvSettingWithBlankDefault("managed_instance_group_ocid")},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"name":                      acctest.Representation{RepType: acctest.Optional, Create: `php`},
		"name_contains":             acctest.Representation{RepType: acctest.Optional, Create: `php`},
		"stream_name":               acctest.Representation{RepType: acctest.Optional, Create: `7.2`},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceGroupModuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceGroupModuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_managed_instance_group_modules.test_managed_instance_group_modules"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_modules", "test_managed_instance_group_modules", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupModuleDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "php"),
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "php"),
				resource.TestCheckResourceAttrSet(datasourceName, "stream_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_module_collection.#"),
			),
		},
	})
}
