// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BdsBdsCapacityReportRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_availabilities": acctest.RepresentationGroup{RepType: acctest.Required, Group: BdsBdsCapacityReportStandardShapeAvailabilitiesRepresentation},
	}

	BdsBdsCapacityReportStandardShapeAvailabilitiesRepresentation = map[string]interface{}{
		"shape": acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.2`},
	}
	BdsBdsCapacityReportFlexShapeAvailabilitiesRepresentation = map[string]interface{}{
		"shape":        acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E4.Flex`},
		"shape_config": acctest.RepresentationGroup{RepType: acctest.Optional, Group: BdsBdsCapacityReportShapeAvailabilitiesShapeConfigRepresentation},
	}
	BdsBdsCapacityReportShapeAvailabilitiesShapeConfigRepresentation = map[string]interface{}{
		"memory_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `120`},
		"ocpus":         acctest.Representation{RepType: acctest.Optional, Create: `8`},
	}

	BdsBdsCapacityReportResourceDependencies = ""
)

// issue-routing-tag: bds/default
func TestBdsBdsCapacityReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBdsBdsCapacityReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_bds_bds_capacity_report.test_bds_capacity_report"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BdsBdsCapacityReportResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_report", "test_bds_capacity_report", acctest.Required, acctest.Create, BdsBdsCapacityReportRepresentation), "bds", "bdsCapacityReport", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BdsBdsCapacityReportResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_bds_bds_capacity_report", "test_bds_capacity_report", acctest.Required, acctest.Create, BdsBdsCapacityReportRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape_availabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "shape_availabilities.0.shape", "VM.Standard2.2"),

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
