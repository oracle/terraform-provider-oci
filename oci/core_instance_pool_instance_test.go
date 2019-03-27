// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	instancePoolInstanceDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
		"instance_pool_id": Representation{repType: Required, create: `${oci_core_instance_pool.test_instance_pool.id}`},
		"display_name":     Representation{repType: Optional, create: `displayName`},
	}

	InstancePoolInstanceResourceConfig = InstancePoolResourceConfig
)

func TestCoreInstancePoolInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolInstanceResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_instance_pool_instances.test_instance_pool_instances"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_instance_pool_instances", "test_instance_pool_instances", Optional, Create, instancePoolInstanceDataSourceRepresentation) +
					compartmentIdVariableStr + InstancePoolInstanceResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(datasourceName, "instance_pool_id"),

					resource.TestCheckResourceAttrSet(datasourceName, "instances.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.instance_configuration_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "instances.0.time_created"),
				),
			},
		},
	})
}
