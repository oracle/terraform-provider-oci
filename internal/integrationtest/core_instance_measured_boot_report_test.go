// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreCoreInstanceMeasuredBootReportSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
	}

	CoreInstanceMeasuredInstanceWithPlatformConfigVMIntelRepresentation = acctest.RepresentationCopyWithNewProperties(CoreInstanceRepresentation, map[string]interface{}{
		"image":               acctest.Representation{RepType: acctest.Required, Create: `${var.InstanceImageOCIDShieldedCompatible[var.region]}`},
		"platform_config":     acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceVMIntelShieldedPlatformConfigRepresentation},
		"availability_domain": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
	})

	CoreInstanceMeasuredBootReportResourceConfig = utils.DefinedShieldedImageOCIDs +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreVcnRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
			"dns_label": acctest.Representation{RepType: acctest.Required, Create: `dnslabel`},
		})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, CoreInstanceMeasuredInstanceWithPlatformConfigVMIntelRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceMeasuredBootReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceMeasuredBootReportResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_instance_measured_boot_report.test_instance_measured_boot_report"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_core_instance_measured_boot_report", "test_instance_measured_boot_report", acctest.Required, acctest.Create, CoreCoreInstanceMeasuredBootReportSingularDataSourceRepresentation) +
					compartmentIdVariableStr + CoreInstanceMeasuredBootReportResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_policy_verification_successful"),
					resource.TestCheckResourceAttr(singularDatasourceName, "measurements.#", "1"),
				),
			},
		},
	})
}
