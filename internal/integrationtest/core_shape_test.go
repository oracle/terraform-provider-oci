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
	CoreCoreShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"image_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_image.test_image.id}`},
	}

	CoreCoreShapeDataSourceRepresentationForFlexShape = acctest.RepresentationCopyWithNewProperties(CoreCoreShapeDataSourceRepresentation, map[string]interface{}{
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreCoreShapeDataSourceFilterRepresentationForFlexShape},
	})

	CoreCoreShapeDataSourceFilterRepresentationForFlexShape = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`VM.Standard.E3.Flex`}},
	}

	CoreCoreShapeResourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"image_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_image.test_image.id}`},
		"shape_name":     acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E2.1`},
	}

	CoreCoreShapeResourceRepresentationForFlexShape = acctest.GetUpdatedRepresentationCopy("shape_name", acctest.Representation{RepType: acctest.Required, Create: InstanceConfigurationVmShapeForFlex},
		acctest.RepresentationCopyWithNewProperties(CoreCoreShapeResourceRepresentation, map[string]interface{}{"shape_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentationForFlexShape}}))

	CoreShapeResourceConfig = AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_image", "test_image", acctest.Required, acctest.Create, CoreImageRepresentation)
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_shape_management.test_shape"
	datasourceName := "data.oci_core_shapes.test_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Add Compatible Image Shape
		{
			Config: config + compartmentIdVariableStr + CoreShapeResourceConfig + utils.OciImageIdsVariable +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_shape_management", "test_shape", acctest.Required, acctest.Create, CoreCoreShapeResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "image_id"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "VM.Standard.E2.1"),
			),
		},
		// verify Delete Compatible Image Shape
		{
			Config: config + compartmentIdVariableStr,
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", acctest.Required, acctest.Create, CoreCoreShapeDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.billing_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.gpus"),
				//resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.local_disk_description"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.is_billed_for_stopped_instance"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.is_flexible"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.is_live_migration_supported"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.is_subcore"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.local_disks"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.local_disks_total_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.max_vnic_attachments"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.ocpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.processor_description"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.quota_names.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.recommended_alternatives.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.resize_compatible_shapes.#"),
			),
		},

		// Delete before re-recreate
		{
			Config: config + compartmentIdVariableStr,
		},

		// ------------------ tests for E3.flex shape -------------------
		// verify Add Compatible Image Shape
		{
			Config: config + compartmentIdVariableStr + CoreShapeResourceConfig + utils.FlexVmImageIdsVariable +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentationForFlexShape) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_shape_management", "test_shape", acctest.Required, acctest.Create, CoreCoreShapeResourceRepresentationForFlexShape),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "image_id"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "VM.Standard.E3.Flex"),
			),
		},

		// verify Delete Compatible Image Shape
		{
			Config: config + compartmentIdVariableStr,
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", acctest.Required, acctest.Create, CoreCoreShapeDataSourceRepresentationForFlexShape) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "shapes.#", "3"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.baseline_ocpu_utilizations.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.gpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.local_disks"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.local_disks_total_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.max_vnic_attachments"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.memory_in_gbs"),
				resource.TestCheckResourceAttr(datasourceName, "shapes.0.memory_options.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.min_total_baseline_ocpus_required"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.network_ports"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.ocpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.processor_description"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.rdma_bandwidth_in_gbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.rdma_ports"),
			),
		},
	})
}
