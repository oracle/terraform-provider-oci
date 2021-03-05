// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	dedicatedVmHostInstanceShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain":     Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"dedicated_vm_host_shape": Representation{repType: Optional, create: `DVH.Standard2.52`},
	}

	DedicatedVmHostInstanceShapeResourceConfig = AvailabilityDomainConfig
)

func TestCoreDedicatedVmHostInstanceShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostInstanceShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_dedicated_vm_host_instance_shapes.test_dedicated_vm_host_instance_shapes"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host_instance_shapes", "test_dedicated_vm_host_instance_shapes", Required, Create, dedicatedVmHostInstanceShapeDataSourceRepresentation) +
					compartmentIdVariableStr + DedicatedVmHostInstanceShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instance_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instance_shapes.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instance_shapes.0.instance_shape_name"),
				),
			},
		},
	})
}
