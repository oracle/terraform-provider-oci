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
	dedicatedVmHostShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"availability_domain": Representation{RepType: Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"instance_shape_name": Representation{RepType: Optional, Create: `VM.Standard2.1`},
	}

	DedicatedVmHostShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_dedicated_vm_host_shapes.test_dedicated_vm_host_shapes"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host_shapes", "test_dedicated_vm_host_shapes", Optional, Create, dedicatedVmHostShapeDataSourceRepresentation) +
				compartmentIdVariableStr + DedicatedVmHostShapeResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
