// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubManagedInstanceGroupManagedInstanceDataSourceRepresentation = map[string]interface{}{
		"managed_instance_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: []string{`displayName`}},
		"display_name_contains":     acctest.Representation{RepType: acctest.Optional, Create: `displayNameContains`},
		"managed_instance_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
	}

	OsManagementHubManagedInstanceGroupManagedInstanceResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance_group", "test_managed_instance_group", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubManagedInstanceRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceGroupManagedInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceGroupManagedInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_managed_instance_group_managed_instances.test_managed_instance_group_managed_instances"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_managed_instances", "test_managed_instance_group_managed_instances", acctest.Required, acctest.Create, OsManagementHubManagedInstanceGroupManagedInstanceDataSourceRepresentation) +
				OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
				compartmentIdVariableStr + OsManagementHubManagedInstanceGroupManagedInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_collection.0.items.#"),
			),
		},

		// verify datasource optional
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_group_managed_instances", "test_managed_instance_group_managed_instances", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceGroupManagedInstanceDataSourceRepresentation) +
				OsManagementHubVendorSoftwareSourceOl8BaseosLatestX8664Config +
				compartmentIdVariableStr + OsManagementHubManagedInstanceGroupManagedInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayNameContains"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_collection.0.items.#"),
			),
		},
	})
}
