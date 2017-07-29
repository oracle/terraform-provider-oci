// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"regexp"
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// This test used to dynamically select the image and shape from the available options, but some images and sizes are /much/ slower to create than others. These are known to be better than the worst case, but feel free to optimize more!
const (
	OracleLinux73ImageID = "ocid1.image.oc1.phx.aaaaaaaa6uwtn7h3hogd5zlwd35eeqbndurkayshzvrfx5usqn6cwxd5vdqq" // DisplayName: Oracle-Linux-7.3-2017.05.23-0
	FastestImageID       = OracleLinux73ImageID
	SmallestShapeName    = "VM.Standard1.1"
)

func TestResourceCoreInstanceCreate(t *testing.T) {
	client := GetTestProvider()

	provider := Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		},
	)

	providers := map[string]terraform.ResourceProvider{
		"baremetal": provider,
	}

	// config := instanceConfig + `
	config := `
data "baremetal_identity_availability_domains" "ADs" {
  compartment_id = "${var.compartment_id}"
}

resource "baremetal_core_virtual_network" "t" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = "${var.compartment_id}"
  display_name   = "network_name"
}

resource "baremetal_core_subnet" "s" {
  availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${baremetal_core_virtual_network.t.id}"
  security_list_ids   = ["${baremetal_core_virtual_network.t.default_security_list_id}"]
  route_table_id      = "${baremetal_core_virtual_network.t.default_route_table_id}"
  dhcp_options_id     = "${baremetal_core_virtual_network.t.default_dhcp_options_id}"
  cidr_block          = "10.0.2.0/24"
}

resource "baremetal_core_instance" "t" {
  availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
  compartment_id      = "${var.compartment_id}"
  image               = "` + FastestImageID + `"
  shape               = "` + SmallestShapeName + `"
  // TODO: replace with create_vnic_details.subnet_id once https://github.com/MustWin/baremetal-sdk-go/pull/159 is resolved
  subnet_id           = "${baremetal_core_subnet.s.id}"
  // create_vnic_details {
  //   subnet_id        = "${baremetal_core_subnet.s.id}"
  // }
}
`
	config += testProviderConfig()

	resourceName := "baremetal_core_instance.t"
	resource.UnitTest(t, resource.TestCase{
		Providers: providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config,
				Check: resource.ComposeAggregateTestCheckFunc(
					// TODO: create_vnic_details
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "create_vnic_details.0.display_name"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.private_ip", "10.0.2.2"),
					resource.TestCheckResourceAttr(resourceName, "create_vnic_details.0.hostname_label", ""),
					// TODO: create_vnic_details.hostname_label
					resource.TestMatchResourceAttr(resourceName, "availability_domain", regexp.MustCompile("uGEq:PHX-AD-1")),
					resource.TestMatchResourceAttr(resourceName, "compartment_id", regexp.MustCompile("ocid1\\.compartment\\.oc1\\..*")),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					// resource.TestMatchResourceAttr(resourceName, "image", FastestImageID), // TODO: investigate why this does not match the image specified
					resource.TestMatchResourceAttr(resourceName, "image", regexp.MustCompile("ocid1\\.image\\.oc1\\..*")),

					resource.TestCheckResourceAttr(resourceName, "ipxe_script", ""),
					resource.TestCheckNoResourceAttr(resourceName, "metadata"), // TODO: investigate why this isn't empty or if I'm checking it wrong
					resource.TestCheckResourceAttr(resourceName, "region", "phx"),
					resource.TestCheckResourceAttr(resourceName, "shape", SmallestShapeName),
					resource.TestCheckResourceAttr(resourceName, "state", baremetal.ResourceRunning),

					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip"),
					resource.TestCheckResourceAttrSet(resourceName, "public_ip"),

					// Deprecated
					resource.TestCheckNoResourceAttr(resourceName, "subnet_id"),
					resource.TestCheckNoResourceAttr(resourceName, "hostname_label"),
					resource.TestCheckNoResourceAttr(resourceName, "create_vnic_details.subnet_id"),
					resource.TestCheckNoResourceAttr(resourceName, "create_vnic_details.hostname_label"),
				),
			},
		},
	})
}

func TestResourceCoreInstanceTerminate(t *testing.T) {
	client := GetTestProvider()
	provider := Provider(
		func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		},
	)
	config := `
data "baremetal_identity_availability_domains" "ADs" {
  compartment_id = "${var.compartment_id}"
}

resource "baremetal_core_virtual_network" "t" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = "${var.compartment_id}"
  display_name   = "network_name"
}

resource "baremetal_core_subnet" "s" {
  availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${baremetal_core_virtual_network.t.id}"
  security_list_ids   = ["${baremetal_core_virtual_network.t.default_security_list_id}"]
  route_table_id      = "${baremetal_core_virtual_network.t.default_route_table_id}"
  dhcp_options_id     = "${baremetal_core_virtual_network.t.default_dhcp_options_id}"
  cidr_block          = "10.0.2.0/24"
}

resource "baremetal_core_instance" "t" {
  availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
  compartment_id      = "${var.compartment_id}"
  metadata            = {}
  image               = "` + FastestImageID + `"
  shape               = "` + SmallestShapeName + `"
  // TODO: replace with create_vnic_details.subnet_id once https://github.com/MustWin/baremetal-sdk-go/pull/159 is resolved
  subnet_id           = "${baremetal_core_subnet.s.id}"

  create_vnic_details {
    assign_public_ip = true
    display_name     = "test-display-name"
    private_ip       = "10.0.2.2"
    // subnet_id           = "${baremetal_core_subnet.s.id}"
  }
}
	`
	config += testProviderConfig()
	resource.UnitTest(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"baremetal": provider,
		},
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config,
			},
			{
				Config:  config,
				Destroy: true,
			},
		},
	})

}
