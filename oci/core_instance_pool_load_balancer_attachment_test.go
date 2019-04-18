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
	instancePoolLoadBalancerAttachmentSingularDataSourceRepresentation = map[string]interface{}{
		"instance_pool_id":                          Representation{repType: Required, create: `${oci_core_instance_pool.test_instance_pool.id}`},
		"instance_pool_load_balancer_attachment_id": Representation{repType: Required, create: `${oci_core_instance_pool.test_instance_pool.load_balancers.0.id}`},
	}

	InstancePoolLoadBalancerAttachmentResourceConfig = InstancePoolResourceConfig
)

func TestCoreInstancePoolLoadBalancerAttachmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePoolLoadBalancerAttachmentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_instance_pool_load_balancer_attachment.test_instance_pool_load_balancer_attachment"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_core_instance_pool_load_balancer_attachment", "test_instance_pool_load_balancer_attachment", Required, Create, instancePoolLoadBalancerAttachmentSingularDataSourceRepresentation) +
					compartmentIdVariableStr + InstancePoolLoadBalancerAttachmentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pool_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "instance_pool_load_balancer_attachment_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "backend_set_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "port"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vnic_selection"),
				),
			},
		},
	})
}
