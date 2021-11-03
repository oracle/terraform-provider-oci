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
	shapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"availability_domain": Representation{RepType: Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"image_id":            Representation{RepType: Optional, Create: `${oci_core_image.test_image.id}`},
	}

	shapeDataSourceRepresentationForFlexShape = RepresentationCopyWithNewProperties(shapeDataSourceRepresentation, map[string]interface{}{
		"filter": RepresentationGroup{Required, shapeDataSourceFilterRepresentationForFlexShape},
	})

	shapeDataSourceFilterRepresentationForFlexShape = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `name`},
		"values": Representation{RepType: Required, Create: []string{`VM.Standard.E3.Flex`}},
	}

	shapeResourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"image_id":       Representation{RepType: Required, Create: `${oci_core_image.test_image.id}`},
		"shape_name":     Representation{RepType: Required, Create: `VM.Standard.E2.1`},
	}

	shapeResourceRepresentationForFlexShape = GetUpdatedRepresentationCopy("shape_name", Representation{RepType: Required, Create: InstanceConfigurationVmShapeForFlex},
		RepresentationCopyWithNewProperties(shapeResourceRepresentation, map[string]interface{}{"shape_config": RepresentationGroup{Optional, instanceShapeConfigRepresentationForFlexShape}}))

	commonShapeResourceConfig = AvailabilityDomainConfig +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, RepresentationCopyWithNewProperties(VcnRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		})) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, RepresentationCopyWithNewProperties(SubnetRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		})) +
		GenerateResourceFromRepresentationMap("oci_core_image", "test_image", Required, Create, imageRepresentation)
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_shape_management.test_shape"
	datasourceName := "data.oci_core_shapes.test_shapes"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Add Compatible Image Shape
		{
			Config: config + compartmentIdVariableStr + commonShapeResourceConfig + OciImageIdsVariable +
				GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_shape_management", "test_shape", Required, Create, shapeResourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", Required, Create, shapeDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.gpus"),
				//resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.local_disk_description"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.is_live_migration_supported"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.local_disks"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.local_disks_total_size_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.max_vnic_attachments"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.memory_in_gbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.ocpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.processor_description"),
			),
		},

		// Delete before re-recreate
		{
			Config: config + compartmentIdVariableStr,
		},

		// ------------------ tests for E3.flex shape -------------------
		// verify Add Compatible Image Shape
		{
			Config: config + compartmentIdVariableStr + commonShapeResourceConfig + FlexVmImageIdsVariable +
				GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentationForFlexShape) +
				GenerateResourceFromRepresentationMap("oci_core_shape_management", "test_shape", Required, Create, shapeResourceRepresentationForFlexShape),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", Required, Create, shapeDataSourceRepresentationForFlexShape) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.ocpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.processor_description"),
			),
		},
	})
}
