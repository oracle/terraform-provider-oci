// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreComputeCapacityReportRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_availabilities": acctest.RepresentationGroup{RepType: acctest.Required,
			Group: CoreComputeCapacityReportShapeAvailabilitiesRepresentation},
	}
	CoreComputeCapacityReportShapeAvailabilitiesRepresentation = map[string]interface{}{
		"instance_shape": acctest.Representation{RepType: acctest.Required, Create: `VM.Standard1.8`},
		"fault_domain":   acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
	}

	CoreComputeCapacityReportRepresentationWithShapeConfig = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_availabilities": acctest.RepresentationGroup{RepType: acctest.Required,
			Group: CoreComputeCapacityReportShapeAvailabilitiesRepresentationWithShapeConfig},
	}
	CoreComputeCapacityReportShapeAvailabilitiesRepresentationWithShapeConfig = map[string]interface{}{
		"instance_shape": acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"fault_domain":   acctest.Representation{RepType: acctest.Optional, Create: `FAULT-DOMAIN-1`},
		"instance_shape_config": acctest.RepresentationGroup{RepType: acctest.Required,
			Group: CoreComputeCapacityReportShapeAvailabilitiesInstanceShapeConfigRepresentationWithShapeConfig},
	}
	CoreComputeCapacityReportShapeAvailabilitiesInstanceShapeConfigRepresentationWithShapeConfig = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `16`},
		"nvmes":         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"ocpus":         acctest.Representation{RepType: acctest.Required, Create: `1`},
	}

	CoreComputeCapacityReportResourceDependencies = AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeCapacityReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_compute_capacity_report.test_compute_capacity_report"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreComputeCapacityReportResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_report", "test_compute_capacity_report", acctest.Required, acctest.Create, CoreComputeCapacityReportRepresentation), "core", "computeCapacityReport", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreComputeCapacityReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_report", "test_compute_capacity_report", acctest.Required, acctest.Create, CoreComputeCapacityReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape_availabilities.0.instance_shape", "VM.Standard1.8"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}

func TestCoreComputeCapacityReportResource_withShapeConfig(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReportResource_withShapeConfig")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_compute_capacity_report.test_compute_capacity_report"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreComputeCapacityReportResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_report", "test_compute_capacity_report", acctest.Required, acctest.Create, CoreComputeCapacityReportRepresentationWithShapeConfig), "core", "computeCapacityReport", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CoreComputeCapacityReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_compute_capacity_report", "test_compute_capacity_report", acctest.Required, acctest.Create, CoreComputeCapacityReportRepresentationWithShapeConfig),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape_availabilities.0.instance_shape", "VM.Standard.E4.Flex"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
