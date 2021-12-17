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
	shapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"image_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_core_image.test_image.id}`},
	}

	shapeDataSourceRepresentationForFlexShape = acctest.RepresentationCopyWithNewProperties(shapeDataSourceRepresentation, map[string]interface{}{
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: shapeDataSourceFilterRepresentationForFlexShape},
	})

	shapeDataSourceFilterRepresentationForFlexShape = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`VM.Standard.E3.Flex`}},
	}

	shapeResourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"image_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_core_image.test_image.id}`},
		"shape_name":     acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E2.1`},
	}

	shapeResourceRepresentationForFlexShape = acctest.GetUpdatedRepresentationCopy("shape_name", acctest.Representation{RepType: acctest.Required, Create: InstanceConfigurationVmShapeForFlex},
		acctest.RepresentationCopyWithNewProperties(shapeResourceRepresentation, map[string]interface{}{"shape_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: instanceShapeConfigRepresentationForFlexShape}}))

	commonShapeResourceConfig = AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_image", "test_image", acctest.Required, acctest.Create, imageRepresentation)
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
			Config: config + compartmentIdVariableStr + commonShapeResourceConfig + utils.OciImageIdsVariable +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_shape_management", "test_shape", acctest.Required, acctest.Create, shapeResourceRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", acctest.Required, acctest.Create, shapeDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + commonShapeResourceConfig + utils.FlexVmImageIdsVariable +
				acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, instanceRepresentationForFlexShape) +
				acctest.GenerateResourceFromRepresentationMap("oci_core_shape_management", "test_shape", acctest.Required, acctest.Create, shapeResourceRepresentationForFlexShape),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", acctest.Required, acctest.Create, shapeDataSourceRepresentationForFlexShape) +
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
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.networking_bandwidth_in_gbps"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.ocpus"),
				resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.processor_description"),
			),
		},
	})
}
