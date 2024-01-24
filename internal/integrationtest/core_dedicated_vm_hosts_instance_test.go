// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreCoreDedicatedVmHostsInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dedicated_vm_host_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"availability_domain":  acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
	}

	CoreDedicatedVmHostsInstanceResourceConfig = CoreInstanceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, acctest.GetUpdatedRepresentationCopy("dedicated_vm_host_id",
			acctest.Representation{RepType: acctest.Required, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`}, CoreInstanceRepresentation))
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostsInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostsInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_dedicated_vm_hosts_instances.test_dedicated_vm_hosts_instances"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + CoreDedicatedVmHostsInstanceResourceConfig,
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_hosts_instances", "test_dedicated_vm_hosts_instances", acctest.Required, acctest.Create, CoreCoreDedicatedVmHostsInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + CoreDedicatedVmHostsInstanceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.shape"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.time_created"),
			),
		},
	})
}
