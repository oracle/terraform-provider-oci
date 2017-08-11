// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceCoreInstance(t *testing.T) {
	config := instanceConfig + `
	data "baremetal_core_instances" "s" {
      		compartment_id = "${var.compartment_id}"
      		availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
    	}`

	config += testProviderConfig()
	resource.UnitTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("baremetal_core_instance.t", "availability_domain"),
					resource.TestCheckResourceAttr("baremetal_core_instance.t", "display_name", "instance_name"),
					resource.TestCheckResourceAttrSet("baremetal_core_instance.t", "id"),
					resource.TestCheckResourceAttr("baremetal_core_instance.t", "state", baremetal.ResourceRunning),
					resource.TestCheckResourceAttrSet("baremetal_core_instance.t", "time_created"),
					resource.TestCheckResourceAttrSet("baremetal_core_instance.t", "public_ip"),
					resource.TestCheckResourceAttrSet("baremetal_core_instance.t", "private_ip"),
					resource.TestCheckResourceAttrSet("data.baremetal_core_instances.s", "instances.#"),
				),
			},
		},
	})
}
