// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ipSecConnectionDeviceStatusSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	IpSecConnectionDeviceStatusResourceConfig = IpSecConnectionRequiredOnlyResource
)

// issue-routing-tag: core/default
func TestCoreIpSecConnectionDeviceStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreIpSecConnectionDeviceStatusResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_ipsec_status.test_ip_sec_connection_device_status"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_ipsec_status", "test_ip_sec_connection_device_status", acctest.Required, acctest.Create, ipSecConnectionDeviceStatusSingularDataSourceRepresentation) +
				compartmentIdVariableStr + IpSecConnectionDeviceStatusResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
