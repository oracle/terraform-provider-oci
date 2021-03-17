// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	vmClusterPatchSingularDataSourceRepresentation = map[string]interface{}{
		"patch_id":      Representation{repType: Required, create: `{}`},
		"vm_cluster_id": Representation{repType: Required, create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	vmClusterPatchDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": Representation{repType: Required, create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	VmClusterPatchResourceConfig = VmClusterNetworkValidatedResourceConfig +
		generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRepresentation)
)

func TestDatabaseVmClusterPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterPatchResource_basic")
	defer httpreplay.SaveScenario()
	if !strings.Contains(getEnvSettingWithBlankDefault("enabled_tests"), "TestDatabaseVmClusterPatchResource_basic") {
		t.Skip("test not supported due to GI Patching not supported in terraform which is pre-requisite for this test")
	}
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_vm_cluster_patches.test_vm_cluster_patches"
	singularDatasourceName := "data.oci_database_vm_cluster_patch.test_vm_cluster_patch"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_vm_cluster_patches", "test_vm_cluster_patches", Required, Create, vmClusterPatchDataSourceRepresentation) +
					compartmentIdVariableStr + VmClusterPatchResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "patches.#"),
					resource.TestCheckResourceAttr(datasourceName, "patches.0.available_actions.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.last_action"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.time_released"),
					resource.TestCheckResourceAttrSet(datasourceName, "patches.0.version"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_vm_cluster_patch", "test_vm_cluster_patch", Required, Create, vmClusterPatchSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VmClusterPatchResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "patch_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "available_actions.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "last_action"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_released"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
				),
			},
		},
	})
}
