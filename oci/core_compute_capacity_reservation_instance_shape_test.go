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
	computeCapacityReservationInstanceShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        Representation{repType: Optional, create: `displayName`},
	}

	ComputeCapacityReservationInstanceShapeResourceConfig = AvailabilityDomainConfig
)

func TestCoreComputeCapacityReservationInstanceShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReservationInstanceShapeResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_capacity_reservation_instance_shapes.test_compute_capacity_reservation_instance_shapes"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instance_shapes", "test_compute_capacity_reservation_instance_shapes", Required, Create, computeCapacityReservationInstanceShapeDataSourceRepresentation) +
					compartmentIdVariableStr + ComputeCapacityReservationInstanceShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.0.instance_shape"),
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instance_shapes", "test_compute_capacity_reservation_instance_shapes", Optional, Create, computeCapacityReservationInstanceShapeDataSourceRepresentation) +
					compartmentIdVariableStr + ComputeCapacityReservationInstanceShapeResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),

					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.0.instance_shape"),
				),
			},
		},
	})
}
