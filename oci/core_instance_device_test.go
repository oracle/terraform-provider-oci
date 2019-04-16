// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	instanceDeviceDataSourceRepresentation = map[string]interface{}{
		"instance_id":  Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"is_available": Representation{repType: Optional, create: `true`},
		"name":         Representation{repType: Optional, create: `/dev/oracleoci/oraclevdb`},
	}

	InstanceDeviceResourceConfig = InstanceRequiredOnlyResource
)

func TestCoreInstanceDeviceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceDeviceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_instance_devices.test_instance_devices"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_instance_devices", "test_instance_devices", Optional, Create, instanceDeviceDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceDeviceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "instance_id"),
					resource.TestCheckResourceAttr(datasourceName, "is_available", "true"),
					resource.TestCheckResourceAttr(datasourceName, "name", "/dev/oracleoci/oraclevdb"),

					resource.TestCheckResourceAttrSet(datasourceName, "devices.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "devices.0.is_available"),
					resource.TestCheckResourceAttrSet(datasourceName, "devices.0.name"),
				),
			},
		},
	})
}
