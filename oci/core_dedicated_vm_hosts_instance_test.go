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
	dedicatedVmHostsInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       Representation{RepType: Required, Create: `${var.compartment_id}`},
		"dedicated_vm_host_id": Representation{RepType: Required, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"availability_domain":  Representation{RepType: Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
	}

	DedicatedVmHostsInstanceResourceConfig = InstanceResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, GetUpdatedRepresentationCopy("dedicated_vm_host_id",
			Representation{RepType: Required, Create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`}, instanceRepresentation))
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostsInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostsInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_dedicated_vm_hosts_instances.test_dedicated_vm_hosts_instances"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + DedicatedVmHostsInstanceResourceConfig,
		},
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_hosts_instances", "test_dedicated_vm_hosts_instances", Required, Create, dedicatedVmHostsInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + DedicatedVmHostsInstanceResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
