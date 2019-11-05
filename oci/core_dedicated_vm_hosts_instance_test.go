// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	dedicatedVmHostsInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":       Representation{repType: Required, create: `${var.compartment_id}`},
		"dedicated_vm_host_id": Representation{repType: Required, create: `${oci_core_dedicated_vm_host.test_dedicated_vm_host.id}`},
		"availability_domain":  Representation{repType: Optional, create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
	}

	DedicatedVmHostsInstanceResourceConfig = InstanceResourceDependencies +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Optional, Create, instanceRepresentation)
)

func TestCoreDedicatedVmHostsInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreDedicatedVmHostsInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_dedicated_vm_hosts_instances.test_dedicated_vm_hosts_instances"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + compartmentIdVariableStr + DedicatedVmHostsInstanceResourceConfig,
			},
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_dedicated_vm_hosts_instances", "test_dedicated_vm_hosts_instances", Required, Create, dedicatedVmHostsInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + DedicatedVmHostsInstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}
