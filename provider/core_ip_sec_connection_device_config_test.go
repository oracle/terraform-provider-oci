// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	IpSecConnectionDeviceConfigResourceConfig = IpSecConnectionDeviceConfigResourceDependencies + `

`
	IpSecConnectionDeviceConfigPropertyVariables = `

`
	IpSecConnectionDeviceConfigResourceDependencies = IpSecConnectionPropertyVariables + IpSecConnectionRequiredOnlyResource
)

func TestCoreIpSecConnectionDeviceConfigResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_ipsec_config.test_ip_sec_connection_device_config"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + `

data "oci_core_ipsec_config" "test_ip_sec_connection_device_config" {
	#Required
	ipsec_id = "${oci_core_ipsec.test_ip_sec_connection.id}"
}
                ` + compartmentIdVariableStr + IpSecConnectionDeviceConfigResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "ipsec_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "tunnels.#", "2"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnels.0.ip_address"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnels.0.shared_secret"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnels.0.time_created"),
				),
			},
		},
	})
}
