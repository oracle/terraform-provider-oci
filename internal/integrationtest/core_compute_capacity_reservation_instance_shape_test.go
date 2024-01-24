// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreCoreComputeCapacityReservationInstanceShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	CoreComputeCapacityReservationInstanceShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeCapacityReservationInstanceShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReservationInstanceShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_capacity_reservation_instance_shapes.test_compute_capacity_reservation_instance_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instance_shapes", "test_compute_capacity_reservation_instance_shapes", acctest.Required, acctest.Create, CoreCoreComputeCapacityReservationInstanceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeCapacityReservationInstanceShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "compute_capacity_reservation_instance_shapes.0.instance_shape"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instance_shapes", "test_compute_capacity_reservation_instance_shapes", acctest.Optional, acctest.Create, CoreCoreComputeCapacityReservationInstanceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + CoreComputeCapacityReservationInstanceShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
