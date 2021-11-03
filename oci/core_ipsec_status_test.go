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
	ipSecConnectionDeviceStatusSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": Representation{RepType: Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	IpSecConnectionDeviceStatusResourceConfig = IpSecConnectionRequiredOnlyResource
)

// issue-routing-tag: core/default
func TestCoreIpSecConnectionDeviceStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpSecConnectionDeviceStatusResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_ipsec_status.test_ip_sec_connection_device_status"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_ipsec_status", "test_ip_sec_connection_device_status", Required, Create, ipSecConnectionDeviceStatusSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IpSecConnectionDeviceStatusResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ipsec_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tunnels.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnels.0.ip_address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "tunnels.0.state", "DOWN"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnels.0.time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnels.0.time_state_modified"),
			),
		},
	})
}
