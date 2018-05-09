// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	PrivateIpRequiredOnlyResource = PrivateIpResourceDependencies + `
resource "oci_core_private_ip" "test_private_ip" {
	#Required
	vnic_id = "${oci_core_vnic.test_vnic.id}"
}
`

	PrivateIpResourceConfig = PrivateIpResourceDependencies + `
resource "oci_core_private_ip" "test_private_ip" {
	#Required
	vnic_id = "${oci_core_vnic.test_vnic.id}"

	#Optional
	display_name = "${var.private_ip_display_name}"
	hostname_label = "${var.private_ip_hostname_label}"
	ip_address = "${var.private_ip_ip_address}"
}
`
	PrivateIpPropertyVariables = `
variable "private_ip_display_name" { default = "displayName" }
variable "private_ip_hostname_label" { default = "hostnameLabel" }
variable "private_ip_ip_address" { default = "ipAddress" }
variable "private_ip_subnet_id" { default = "subnetId" }

`
	PrivateIpResourceDependencies = "" // Uncomment once defined: VnicPropertyVariables + VnicResourceConfig
)

func TestCorePrivateIpResource_basic(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_private_ip.test_private_ip"
	datasourceName := "data.oci_core_private_ips.test_private_ips"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config + PrivateIpPropertyVariables + compartmentIdVariableStr + PrivateIpRequiredOnlyResource,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + PrivateIpResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + PrivateIpPropertyVariables + compartmentIdVariableStr + PrivateIpResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + `
variable "private_ip_display_name" { default = "displayName2" }
variable "private_ip_hostname_label" { default = "hostnameLabel2" }
variable "private_ip_ip_address" { default = "ipAddress" }
variable "private_ip_subnet_id" { default = "subnetId" }

                ` + compartmentIdVariableStr + PrivateIpResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates to Force New parameters.
			{
				Config: config + `
variable "private_ip_display_name" { default = "displayName2" }
variable "private_ip_hostname_label" { default = "hostnameLabel2" }
variable "private_ip_ip_address" { default = "ipAddress2" }
variable "private_ip_subnet_id" { default = "subnetId2" }

                ` + compartmentIdVariableStr2 + PrivateIpResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress2"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated but it wasn't.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
variable "private_ip_display_name" { default = "displayName2" }
variable "private_ip_hostname_label" { default = "hostnameLabel2" }
variable "private_ip_ip_address" { default = "ipAddress2" }
variable "private_ip_subnet_id" { default = "subnetId2" }

data "oci_core_private_ips" "test_private_ips" {

	#Optional
	ip_address = "${var.private_ip_ip_address}"
	subnet_id = "${var.private_ip_subnet_id}"
	vnic_id = "${oci_core_vnic.test_vnic.id}"

    filter {
    	name = "id"
    	values = ["${oci_core_private_ip.test_private_ip.id}"]
    }
}
                ` + compartmentIdVariableStr2 + PrivateIpResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "ip_address", "ipAddress2"),
					resource.TestCheckResourceAttr(datasourceName, "subnet_id", "subnetId2"),
					resource.TestCheckResourceAttrSet(datasourceName, "vnic_id"),

					resource.TestCheckResourceAttr(datasourceName, "private_ips.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.hostname_label", "hostnameLabel2"),
					resource.TestCheckResourceAttr(datasourceName, "private_ips.0.ip_address", "ipAddress2"),
					resource.TestCheckResourceAttrSet(datasourceName, "private_ips.0.vnic_id"),
				),
			},
		},
	})
}

func TestCorePrivateIpResource_forcenew(t *testing.T) {
	t.Skip("Skipping generated test for now as it has not been worked on.")
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_private_ip.test_private_ip"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + PrivateIpPropertyVariables + compartmentIdVariableStr + PrivateIpResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
variable "private_ip_display_name" { default = "displayName" }
variable "private_ip_hostname_label" { default = "hostnameLabel" }
variable "private_ip_ip_address" { default = "ipAddress2" }
variable "private_ip_subnet_id" { default = "subnetId" }
				` + compartmentIdVariableStr + PrivateIpResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "hostname_label", "hostnameLabel"),
					resource.TestCheckResourceAttr(resourceName, "ip_address", "ipAddress2"),
					resource.TestCheckResourceAttrSet(resourceName, "vnic_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter IpAddress but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
