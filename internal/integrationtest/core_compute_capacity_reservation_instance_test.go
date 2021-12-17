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
	computeCapacityReservationInstanceDataSourceRepresentation = map[string]interface{}{
		"capacity_reservation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"availability_domain":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	capacityReservationInstanceRepresentation = map[string]interface{}{
		"availability_domain":     acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape":                   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"capacity_reservation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"subnet_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"image":                   acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCID[var.region]}`},
	}
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeCapacityReservationInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReservationInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_capacity_reservation_instances.test_compute_capacity_reservation_instances"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				utils.OciImageIdsVariable +
				ComputeCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Required, acctest.Create, computeCapacityReservationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, capacityReservationInstanceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instances", "test_compute_capacity_reservation_instances", acctest.Required, acctest.Create, computeCapacityReservationInstanceDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				utils.OciImageIdsVariable +
				ComputeCapacityReservationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", acctest.Required, acctest.Create, computeCapacityReservationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, capacityReservationInstanceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instances", "test_compute_capacity_reservation_instances", acctest.Optional, acctest.Create, computeCapacityReservationInstanceDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	})
}
