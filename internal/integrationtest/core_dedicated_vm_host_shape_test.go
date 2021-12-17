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
	dedicatedVmHostShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"instance_shape_name": acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard2.1`},
	}

	DedicatedVmHostShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_dedicated_vm_host_shapes.test_dedicated_vm_host_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host_shapes", "test_dedicated_vm_host_shapes", acctest.Optional, acctest.Create, dedicatedVmHostShapeDataSourceRepresentation) +
				compartmentIdVariableStr + DedicatedVmHostShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "instance_shape_name", "VM.Standard2.1"),

				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_shapes.0.dedicated_vm_host_shape"),
			),
		},
	})
}
