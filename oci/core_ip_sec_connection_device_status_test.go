// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	ipSecConnectionDeviceStatusSingularDataSourceRepresentation = map[string]interface{}{
		"ipsec_id": Representation{repType: Required, create: `${oci_core_ipsec.test_ip_sec_connection.id}`},
	}

	IpSecConnectionDeviceStatusResourceConfig = IpSecConnectionRequiredOnlyResource
)

func TestCoreIpSecConnectionDeviceStatusResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_ipsec_status.test_ip_sec_connection_device_status"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_ipsec_status", "test_ip_sec_connection_device_status", Required, Create, ipSecConnectionDeviceStatusSingularDataSourceRepresentation) +
					compartmentIdVariableStr + IpSecConnectionDeviceStatusResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}
