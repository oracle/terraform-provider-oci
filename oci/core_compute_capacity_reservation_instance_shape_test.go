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
	computeCapacityReservationInstanceShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"availability_domain": Representation{RepType: Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        Representation{RepType: Optional, Create: `displayName`},
	}

	ComputeCapacityReservationInstanceShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeCapacityReservationInstanceShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReservationInstanceShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_capacity_reservation_instance_shapes.test_compute_capacity_reservation_instance_shapes"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instance_shapes", "test_compute_capacity_reservation_instance_shapes", Required, Create, computeCapacityReservationInstanceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeCapacityReservationInstanceShapeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.0.instance_shape"),
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instance_shapes", "test_compute_capacity_reservation_instance_shapes", Optional, Create, computeCapacityReservationInstanceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeCapacityReservationInstanceShapeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),

				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.0.instance_shape"),
			),
		},
	})
}
