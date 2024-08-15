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
	DesktopsDesktopPoolVolumeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"desktop_pool_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_desktops_desktop_pool.test_desktop_pool.id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		//"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_desktops_desktop_pool_volume.test_desktop_pool_volume.id}`},
		"state": acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	DesktopsDesktopPoolVolumeResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Required, acctest.Create, DesktopsDesktopPoolRepresentation) +
		DesktopsDesktopPoolResourceDependencies
)

// issue-routing-tag: desktops/default
func TestDesktopsDesktopPoolVolumeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDesktopsDesktopPoolVolumeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_desktops_desktop_pool_volumes.test_desktop_pool_volumes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_desktops_desktop_pool_volumes", "test_desktop_pool_volumes", acctest.Optional, acctest.Create, DesktopsDesktopPoolVolumeDataSourceRepresentation) +
				compartmentIdVariableStr + DesktopsDesktopPoolVolumeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				//resource.TestCheckResourceAttr(datasourceName, "id", "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_volume_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "desktop_pool_volume_collection.0.items.#", "0"),
			),
		},
	})
}
