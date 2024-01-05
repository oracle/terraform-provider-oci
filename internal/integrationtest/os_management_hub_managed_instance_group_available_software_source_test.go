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
	OsManagementHubManagedInstanceGroupAvailableSoftwareSourceDataSourceRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: []string{`ol8_appstream-x86_64`}},
		"display_name_contains":     acctest.Representation{RepType: acctest.Optional, Create: `ol8_appstream-x86_64`},
	}

	OsManagementHubManagedInstanceGroupAvailableSoftwareSourceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceGroupAvailableSoftwareSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceGroupAvailableSoftwareSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_managed_instance_group_available_software_sources.test_managed_instance_group_available_software_sources"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config + OsManagementHubManagedInstanceGroupAvailableSoftwareSourceResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_available_software_sources", "test_managed_instance_group_available_software_sources", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupAvailableSoftwareSourceDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "ol8_appstream-x86_64"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "available_software_source_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "available_software_source_collection.0.items.#", "1"),
			),
		},
	})
}
