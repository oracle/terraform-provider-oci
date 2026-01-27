// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	OsManagementHubManagedInstanceSnapDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_managed_instance.test_managed_instance.id}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `oracle-cloud-agent`},
		"name_contains":       acctest.Representation{RepType: acctest.Optional, Create: `oracle-cloud-agent`},
	}

	OsManagementHubManagedInstanceSnapResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_managed_instance", "test_managed_instance", acctest.Required, acctest.Create, OsManagementHubUbuntuManagedInstanceRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagedInstanceSnapResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagedInstanceSnapResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_os_management_hub_managed_instance_snaps.test_managed_instance_snaps"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_snaps", "test_managed_instance_snaps", acctest.Required, acctest.Create, OsManagementHubManagedInstanceSnapDataSourceRepresentation) +
				OsManagementHubManagedInstanceSnapResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "snap_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "snap_collection.0.items.#"),
			),
		},

		// verify datasource with optional
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_managed_instance_snaps", "test_managed_instance_snaps", acctest.Optional, acctest.Create, OsManagementHubManagedInstanceSnapDataSourceRepresentation) +
				OsManagementHubManagedInstanceSnapResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_instance_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "oracle-cloud-agent"),
				resource.TestCheckResourceAttr(datasourceName, "name_contains", "oracle-cloud-agent"),
				resource.TestCheckResourceAttrSet(datasourceName, "snap_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "snap_collection.0.items.#"),
			),
		},
	})
}
