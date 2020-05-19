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
	managedInstanceSingularDataSourceRepresentation = map[string]interface{}{
		"managed_instance_id": Representation{repType: Required, create: `${var.managed_instance_id}`},
	}

	managedInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `osms-instance`},
		"os_family":      Representation{repType: Optional, create: `ALL`},
	}

	ManagedInstanceResourceConfig = ""
)

func TestOsmanagementManagedInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOsmanagementManagedInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedInstanceId := getEnvSettingWithBlankDefault("managed_instance_id")
	managedInstanceIdVariableStr := fmt.Sprintf("variable \"managed_instance_id\" { default = \"%s\" }\n", managedInstanceId)

	datasourceName := "data.oci_osmanagement_managed_instances.test_managed_instances"
	singularDatasourceName := "data.oci_osmanagement_managed_instance.test_managed_instance"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_osmanagement_managed_instances", "test_managed_instances", Required, Create, managedInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedInstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_osmanagement_managed_instance", "test_managed_instance", Required, Create, managedInstanceSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ManagedInstanceResourceConfig + managedInstanceIdVariableStr,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_instance_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_reboot_required"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "last_boot"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "last_checkin"),
					resource.TestCheckResourceAttr(singularDatasourceName, "managed_instance_groups.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_family"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_kernel_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "os_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "updates_available"),
				),
			},
		},
	})
}
