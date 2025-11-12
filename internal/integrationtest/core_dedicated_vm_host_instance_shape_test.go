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
	CoreCoreDedicatedVmHostInstanceShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"dedicated_vm_host_shape": acctest.Representation{RepType: acctest.Optional, Create: `DVH.Standard.E4.128`},
	}

	CoreDedicatedVmHostInstanceShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostInstanceShapeResource_basic(t *testing.T) {
	if err := httpreplay.SetScenario("TestCoreDedicatedVmHostInstanceShapeResource_basic"); err != nil {
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

	datasourceName := "data.oci_core_dedicated_vm_host_instance_shapes.test_dedicated_vm_host_instance_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host_instance_shapes", "test_dedicated_vm_host_instance_shapes", acctest.Required, acctest.Create, CoreCoreDedicatedVmHostInstanceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDedicatedVmHostInstanceShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instance_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instance_shapes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instance_shapes.0.instance_shape_name"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_instance_shapes.0.instance_shape_name", `VM.Standard.E4.Flex`),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_instance_shapes.0.supported_capabilities.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_instance_shapes.0.supported_capabilities.0.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instance_shapes.0.supported_capabilities.0.is_memory_encryption_supported"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_instance_shapes.0.supported_capabilities.0.is_memory_encryption_supported", isMemoryEncryptionSupported),
			),
		},
	})
}
