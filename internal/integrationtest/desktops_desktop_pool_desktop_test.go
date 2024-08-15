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
	DesktopsDesktopPoolDesktopDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"desktop_pool_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_desktops_desktop_pool.test_desktop_pool.id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		//"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		//"id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_desktops_desktop_pool_desktop.test_desktop_pool_desktop.id}`},
	}

	DesktopsDesktopPoolDesktopResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_desktops_desktop_pool", "test_desktop_pool", acctest.Required, acctest.Create, DesktopsDesktopPoolRepresentation) +
		DesktopsDesktopPoolResourceDependencies
)

// issue-routing-tag: desktops/default
func TestDesktopsDesktopPoolDesktopResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDesktopsDesktopPoolDesktopResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_desktops_desktop_pool_desktops.test_desktop_pool_desktops"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_desktops_desktop_pool_desktops", "test_desktop_pool_desktops", acctest.Optional, acctest.Create, DesktopsDesktopPoolDesktopDataSourceRepresentation) +
				compartmentIdVariableStr + DesktopsDesktopPoolDesktopResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_id"),
				//resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_desktop_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "desktop_pool_desktop_collection.0.items.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "desktop_pool_desktop_collection.0.items.0.user_name", ""),
				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_desktop_collection.0.items.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_desktop_collection.0.items.0.instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "desktop_pool_desktop_collection.0.items.0.desktop_id"),
			),
		},
	})
}
