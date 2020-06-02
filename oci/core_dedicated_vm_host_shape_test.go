// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	dedicatedVmHostShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"instance_shape_name": Representation{repType: Optional, create: `VM.Standard2.1`},
	}

	DedicatedVmHostShapeResourceConfig = AvailabilityDomainConfig
)

func TestCoreDedicatedVmHostShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_dedicated_vm_host_shapes.test_dedicated_vm_host_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host_shapes", "test_dedicated_vm_host_shapes", Optional, Create, dedicatedVmHostShapeDataSourceRepresentation) +
					compartmentIdVariableStr + DedicatedVmHostShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "instance_shape_name", "VM.Standard2.1"),

					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.0.dedicated_vm_host_shape"),
				),
			},
		},
	})
}
