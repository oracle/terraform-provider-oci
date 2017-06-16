// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

func TestIsStatefulResource(t *testing.T) {
	var sr crud.StatefulResource
	sr = &InstanceResourceCrud{}
	if sr == nil {
		t.Fail()
	}
}

func TestDataSourceCoreInstanceCreate(t *testing.T) {
	client := GetTestProvider()

	provider := Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		},
	)

	providers := map[string]terraform.ResourceProvider{
		"baremetal": provider,
	}

	config := instanceConfig + `
	data "baremetal_core_instances" "s" {
		compartment_id = "${var.compartment_id}"
		availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
	}`

	config += testProviderConfig()

	resourceName := "baremetal_core_instance.t"
	resource.UnitTest(t, resource.TestCase{
		Providers: providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "instance_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "state", baremetal.ResourceRunning),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "public_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
					resource.TestCheckResourceAttrSet("data.baremetal_core_instances.s", "instances.#"),
				),
			},
		},
	})
}
