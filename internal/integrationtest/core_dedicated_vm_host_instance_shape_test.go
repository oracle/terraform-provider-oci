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
	dedicatedVmHostInstanceShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":     acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"dedicated_vm_host_shape": acctest.Representation{RepType: acctest.Optional, Create: `DVH.Standard2.52`},
	}

	DedicatedVmHostInstanceShapeResourceConfig = AvailabilityDomainConfig
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostInstanceShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostInstanceShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_dedicated_vm_host_instance_shapes.test_dedicated_vm_host_instance_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_host_instance_shapes", "test_dedicated_vm_host_instance_shapes", acctest.Required, acctest.Create, dedicatedVmHostInstanceShapeDataSourceRepresentation) +
				compartmentIdVariableStr + DedicatedVmHostInstanceShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instance_shapes.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instance_shapes.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instance_shapes.0.instance_shape_name"),
			),
		},
	})
}
