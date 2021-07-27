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
	vmClusterUpdateSingularDataSourceRepresentation = map[string]interface{}{
		"update_id":     Representation{repType: Required, create: `{}`},
		"vm_cluster_id": Representation{repType: Required, create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	vmClusterUpdateDataSourceRepresentation = map[string]interface{}{
		"vm_cluster_id": Representation{repType: Required, create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"state":         Representation{repType: Optional, create: `AVAILABLE`},
		"update_type":   Representation{repType: Optional, create: `GI_UPGRADE`},
	}

	VmClusterUpdateResourceConfig = VmClusterNetworkValidatedResourceConfig +
		generateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseVmClusterUpdateResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterUpdateResource_basic")
	defer httpreplay.SaveScenario()
	if !strings.Contains(getEnvSettingWithBlankDefault("enabled_tests"), "TestDatabaseVmClusterUpdateResource_basic") {
		t.Skip("test not supported due to GI Update not supported in terraform which is pre-requisite for this test")
	}
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_vm_cluster_updates.test_vm_cluster_updates"
	singularDatasourceName := "data.oci_database_vm_cluster_update.test_vm_cluster_update"

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
					generateDataSourceFromRepresentationMap("oci_database_vm_cluster_updates", "test_vm_cluster_updates", Required, Create, vmClusterUpdateDataSourceRepresentation) +
					compartmentIdVariableStr + VmClusterUpdateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(datasourceName, "update_type", "GI_UPGRADE"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.#"),
					resource.TestCheckResourceAttr(datasourceName, "vm_cluster_updates.0.available_actions.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.description"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.last_action"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.time_released"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.update_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_updates.0.version"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_vm_cluster_update", "test_vm_cluster_update", Required, Create, vmClusterUpdateSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VmClusterUpdateResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "update_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "available_actions.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "last_action"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_released"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "update_type"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
				),
			},
		},
	})
}
