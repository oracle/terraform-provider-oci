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
		"compartment_id":      Representation{repType: Required, create: `${var.compartment_id}`},
		"availability_domain": Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"image_id":            Representation{repType: Optional, create: `${oci_core_image.test_image.id}`},
	}

	shapeDataSourceRepresentationForFlexShape = representationCopyWithNewProperties(shapeDataSourceRepresentation, map[string]interface{}{
		"filter": RepresentationGroup{Required, shapeDataSourceFilterRepresentationForFlexShape},
	})

	shapeDataSourceFilterRepresentationForFlexShape = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`VM.Standard.E3.Flex`}},
	}

	shapeResourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"image_id":       Representation{repType: Required, create: `${oci_core_image.test_image.id}`},
		"shape_name":     Representation{repType: Required, create: `VM.Standard.E2.1`},
	}

	shapeResourceRepresentationForFlexShape = getUpdatedRepresentationCopy("shape_name", Representation{repType: Required, create: InstanceConfigurationVmShapeForFlex},
		representationCopyWithNewProperties(shapeResourceRepresentation, map[string]interface{}{"shape_config": RepresentationGroup{Optional, instanceShapeConfigRepresentationForFlexShape}}))

	commonShapeResourceConfig = AvailabilityDomainConfig +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, representationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, representationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": Representation{repType: Required, create: `dnslabel`},
		})) +
		generateResourceFromRepresentationMap("oci_core_image", "test_image", Required, Create, imageRepresentation)
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_shape_management.test_shape"
	datasourceName := "data.oci_core_shapes.test_shapes"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Add Compatible Image Shape
		{
			Config: config + compartmentIdVariableStr + commonShapeResourceConfig + OciImageIdsVariable +
				generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
				generateResourceFromRepresentationMap("oci_core_shape_management", "test_shape", Required, Create, shapeResourceRepresentation),
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
				generateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", Required, Create, shapeDataSourceRepresentation) +
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
				generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentationForFlexShape) +
				generateResourceFromRepresentationMap("oci_core_shape_management", "test_shape", Required, Create, shapeResourceRepresentationForFlexShape),
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
				generateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", Required, Create, shapeDataSourceRepresentationForFlexShape) +
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
