// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	OsManagementHubManagementStationMirrorDataSourceRepresentation = map[string]interface{}{
		"management_station_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_os_management_hub_management_station.test_management_station.id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayNameContains`},
		"mirror_states":         acctest.Representation{RepType: acctest.Optional, Create: []string{`mirrorStates`}},
	}
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubManagementStationMirrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsManagementHubManagementStationMirrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_os_management_hub_management_station_mirrors.test_management_station_mirrors"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_management_station_mirrors", "test_management_station_mirrors", acctest.Required, acctest.Create, OsManagementHubManagementStationMirrorDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_management_station", "test_management_station", acctest.Required, acctest.Create, OsManagementHubManagementStationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "management_station_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "mirrors_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "mirrors_collection.0.items.#"),
			),
		},
	})
}
