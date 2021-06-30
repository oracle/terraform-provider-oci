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
	computeCapacityReservationInstanceDataSourceRepresentation = map[string]interface{}{
		"capacity_reservation_id": Representation{repType: Required, create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"availability_domain":     Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          Representation{repType: Optional, create: `${var.compartment_id}`},
	}

	capacityReservationInstanceRepresentation = map[string]interface{}{
		"availability_domain":     Representation{repType: Required, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          Representation{repType: Required, create: `${var.compartment_id}`},
		"shape":                   Representation{repType: Required, create: `VM.Standard2.1`},
		"capacity_reservation_id": Representation{repType: Required, create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"subnet_id":               Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"image":                   Representation{repType: Required, create: `${var.InstanceImageOCID[var.region]}`},
	}
)

func TestCoreComputeCapacityReservationInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReservationInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_capacity_reservation_instances.test_compute_capacity_reservation_instances"

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
					OciImageIdsVariable +
					ComputeCapacityReservationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
					generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
					generateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationRepresentation) +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, capacityReservationInstanceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instances", "test_compute_capacity_reservation_instances", Required, Create, computeCapacityReservationInstanceDataSourceRepresentation) +
					compartmentIdVariableStr,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.fault_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.shape"),
					resource.TestCheckResourceAttr(datasourceName, "capacity_reservation_instances.0.shape_config.#", "1"),
				),
			},
			{
				Config: config +
					OciImageIdsVariable +
					ComputeCapacityReservationResourceDependencies +
					generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
					generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
					generateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationRepresentation) +
					generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, capacityReservationInstanceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instances", "test_compute_capacity_reservation_instances", Optional, Create, computeCapacityReservationInstanceDataSourceRepresentation) +
					compartmentIdVariableStr,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_id"),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.fault_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.shape"),
					resource.TestCheckResourceAttr(datasourceName, "capacity_reservation_instances.0.shape_config.#", "1"),
				),
			},
		},
	})
}
