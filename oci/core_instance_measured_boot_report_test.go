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
	instanceMeasuredBootReportSingularDataSourceRepresentation = map[string]interface{}{
		"instance_id": Representation{RepType: Required, Create: `${oci_core_instance.test_instance.id}`},
	}

	instanceWithPlatformConfigVMIntelRepresentation = RepresentationCopyWithNewProperties(instanceRepresentation, map[string]interface{}{
		"image":               Representation{RepType: Required, Create: `${var.InstanceImageOCIDShieldedCompatible[var.region]}`},
		"platform_config":     RepresentationGroup{Required, instanceVMIntelShieldedPlatformConfigRepresentation},
		"availability_domain": Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.1.name}`},
	})

	InstanceMeasuredBootReportResourceConfig = DefinedShieldedImageOCIDs +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, RepresentationCopyWithNewProperties(VcnRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		})) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, RepresentationCopyWithNewProperties(SubnetRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		})) +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceWithPlatformConfigVMIntelRepresentation) +
		AvailabilityDomainConfig
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreInstanceMeasuredBootReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstanceMeasuredBootReportResource_basic")
	defer httpreplay.SaveScenario()

	provider := TestAccProvider
	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_instance_measured_boot_report.test_instance_measured_boot_report"

	SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { PreCheck() },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_core_instance_measured_boot_report", "test_instance_measured_boot_report", Required, Create, instanceMeasuredBootReportSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceMeasuredBootReportResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_policy_verification_successful"),
					resource.TestCheckResourceAttr(singularDatasourceName, "measurements.#", "1"),
				),
			},
		},
	})
}
