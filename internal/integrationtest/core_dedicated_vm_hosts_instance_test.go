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
	vmShape = `VM.Standard.E5.Flex`

	CoreDedicatedVmHostsInstanceDataSourceRepresentation = map[string]interface{}{
		"availability_domain": acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		// The DVH OCID used here is Confidential, therefore, we need to check for is_memory_encryption_enabled in the tests
		"dedicated_vm_host_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.dedicated_vm_host_id}`},
		"is_memory_encryption_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}
)

// issue-routing-tag: core/default
func TestCoreDedicatedVmHostsInstanceResource_basic(t *testing.T) {
	if err := httpreplay.SetScenario("TestCoreDedicatedVmHostsInstanceResource_basic"); err != nil {
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

	availabilityDomain := utils.GetEnvSettingWithBlankDefault("availability_domain")
	availabilityDomainVariableStr := fmt.Sprintf("variable \"availability_domain\" { default = \"%s\" }\n", availabilityDomain)

	dedicatedVmHostId := utils.GetEnvSettingWithBlankDefault("dedicated_vm_host_id")
	dedicatedVmHostIdVariableStr := fmt.Sprintf("variable \"dedicated_vm_host_id\" { default = \"%s\" }\n", dedicatedVmHostId)

	datasourceName := "data.oci_core_dedicated_vm_hosts_instances.test_dedicated_vm_hosts_instances"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + availabilityDomainVariableStr + dedicatedVmHostIdVariableStr,
		},
		{
			Config: config + availabilityDomainVariableStr + dedicatedVmHostIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_dedicated_vm_hosts_instances",
					"test_dedicated_vm_hosts_instances", acctest.Required, acctest.Create,
					CoreDedicatedVmHostsInstanceDataSourceRepresentation) +
				compartmentIdVariableStr + coreDedicatedVmHostResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.instance_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.is_memory_encryption_enabled"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_instances.0.is_memory_encryption_enabled", isMemoryEncryptionEnabled),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.shape"),
				resource.TestCheckResourceAttr(datasourceName, "dedicated_vm_host_instances.0.shape", vmShape),
				resource.TestCheckResourceAttrSet(datasourceName, "dedicated_vm_host_instances.0.time_created"),
			),
		},
	})
}
