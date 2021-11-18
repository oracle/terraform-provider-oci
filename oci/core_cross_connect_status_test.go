// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	crossConnectStatusSingularDataSourceRepresentation = map[string]interface{}{
		"cross_connect_id": Representation{RepType: Required, Create: `${oci_core_cross_connect.test_cross_connect.id}`},
	}

	CrossConnectStatusResourceConfig = GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_locations", "test_cross_connect_locations", Required, Create, crossConnectLocationDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_cross_connect", "test_cross_connect", Required, Create, crossConnectRepresentation)
)

// issue-routing-tag: core/default
func TestCoreCrossConnectStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreCrossConnectStatusResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_cross_connect_status.test_cross_connect_status"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_cross_connect_status", "test_cross_connect_status", Required, Create, crossConnectStatusSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CrossConnectStatusResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cross_connect_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "encryption_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "interface_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "light_level_ind_bm"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "light_level_indicator"),
			),
		},
	})
}
