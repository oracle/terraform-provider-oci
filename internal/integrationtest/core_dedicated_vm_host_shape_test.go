// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	vmShapeName                 = `VM.Standard.E6.Flex`
	capacityConfigName          = `standard_e6_flex`
	capacityBinsSize            = `2`
	capacityConfigsSize         = `1`
	dvhHostShape                = `DVH.Standard.E6.256`
	isDefault                   = `true`
	isMemoryEncryptionSupported = `true`

	CoreCoreDedicatedVmHostShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"instance_shape_name": acctest.Representation{RepType: acctest.Optional, Create: vmShapeName},
	}

	coreDedicatedVmHostResourceDependencies = AvailabilityDomainConfig
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostShapeResource_basic(t *testing.T) {
	if err := httpreplay.SetScenario("TestCoreDedicatedVmHostShapeResource_basic"); err != nil {
		fmt.Printf("error occurred in httpreplay.SetScenario, %s", err)
		return
	}
	defer func() {
		if err := httpreplay.SaveScenario(); err != nil {
			fmt.Printf("error occurred in httpreplay.SaveScenario, %s", err)
			return
		}
	}()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_dedicated_vm_host_shapes.test_dedicated_vm_host_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host_shapes", "test_dedicated_vm_host_shapes",
				acctest.Optional, acctest.Create, CoreCoreDedicatedVmHostShapeDataSourceRepresentation) +
				compartmentIdVariableStr + coreDedicatedVmHostResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "instance_shape_name", vmShapeName),

				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.0.capacity_configs.#"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_shapes.0.capacity_configs.#", capacityConfigsSize),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_shapes.0.capacity_configs.0.capacity_bins.#", capacityBinsSize),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_shapes.0.capacity_configs.0.capacity_config_name", capacityConfigName),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_shapes.0.capacity_configs.0.is_default", isDefault),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_shapes.0.capacity_configs.0.supported_capabilities.0.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.0.capacity_configs.0.supported_capabilities.0.is_memory_encryption_supported"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_shapes.0.capacity_configs.0.supported_capabilities.0.is_memory_encryption_supported", isMemoryEncryptionSupported),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.0.dedicated_vm_host_shape"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_shapes.0.dedicated_vm_host_shape", dvhHostShape),
			),
		},
	})
}
