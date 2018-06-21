// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

const (
	displayName  = "-tf-public-ip"
	displayName2 = displayName + "-updated"
	privateIpId  = "private_ips.0.id"
	privateIpId2 = "private_ips.1.id"

	PublicIpRequiredOnlyResource = PublicIpResourceDependencies + `
resource "oci_core_public_ip" "test_public_ip" {
	#Required
	compartment_id = "${var.compartment_id}"
	lifetime = "${var.public_ip_lifetime}"
}
`

	PublicIpResourceConfig = PublicIpResourceDependencies + `
resource "oci_core_public_ip" "test_public_ip" {
	#Required
	compartment_id = "${var.compartment_id}"
	lifetime = "${var.public_ip_lifetime}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.public_ip_defined_tags_value}")}"
	display_name = "${var.public_ip_display_name}"
	freeform_tags = "${var.public_ip_freeform_tags}"
	private_ip_id = "${data.oci_core_private_ips.test_private_ips.` + privateIpId + `}"
}
`
	PublicIpUnassignedResourceConfig = PublicIpResourceDependencies + `
resource "oci_core_public_ip" "test_public_ip" {
	#Required
	compartment_id = "${var.compartment_id}"
	lifetime = "${var.public_ip_lifetime}"

	#Optional
	defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.public_ip_defined_tags_value}")}"
	display_name = "${var.public_ip_display_name}"
	freeform_tags = "${var.public_ip_freeform_tags}"
}
`
	PublicIpPropertyVariables = `
variable "public_ip_defined_tags_value" { default = "value" }
variable "public_ip_display_name" { default = "-tf-public-ip" }
variable "public_ip_freeform_tags" { default = {"Department"= "Finance"} }
variable "public_ip_lifetime" { default = "RESERVED" }
variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeRegion) + `" }
`
	PublicIpResourceDependencies = `
	variable "InstanceImageOCID" {
		type = "map"
		default = {
			// See https://docs.us-phoenix-1.oraclecloud.com/images/
			// Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
			us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
			us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
			eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
			uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
		}
	}
	data "oci_identity_availability_domains" "test_availability_domains" {
		compartment_id = "${var.compartment_id}"
	}
	resource "oci_core_virtual_network" "test_vcn" {
		cidr_block      = "10.0.0.0/16"
		compartment_id  = "${var.compartment_id}"
		display_name    = "-tf-vcn"
		dns_label       = "testvcn"
	}
	resource "oci_core_subnet" "test_subnet" {
		availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.test_vcn.id}"
		route_table_id      = "${oci_core_virtual_network.test_vcn.default_route_table_id}"
		security_list_ids   = ["${oci_core_virtual_network.test_vcn.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.test_vcn.default_dhcp_options_id}"
		dns_label           = "testsubnet"
	}
	resource "oci_core_instance" "test_instance" {
		availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
		compartment_id      = "${var.compartment_id}"
		display_name        = "-tf-instance"
		image               = "${var.InstanceImageOCID[var.region]}"
		shape               = "VM.Standard1.8"
		create_vnic_details {
			assign_public_ip = false
			subnet_id        = "${oci_core_subnet.test_subnet.id}"
			hostname_label   = "testinstance"
			display_name     = "-tf-instance-vnic"
		}
		metadata {
			ssh_authorized_keys = "${var.ssh_public_key}"
		}
		timeouts {
			create = "15m"
		}
	}
	resource "oci_core_vnic_attachment" "test_vnic_attachments" {
		instance_id  = "${oci_core_instance.test_instance.id}"
		display_name = "-tf-vnic-attachment-2"
		create_vnic_details {
			assign_public_ip = false
			subnet_id        = "${oci_core_subnet.test_subnet.id}"
			display_name     = "-tf-vnic-2"
			hostname_label   = "testinstance2"
		}
	}
	data "oci_core_private_ips" "test_private_ips" {
		subnet_id        = "${oci_core_subnet.test_subnet.id}"
	}`
)

func TestCorePublicIpResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig() + DefinedTagsDependencies

	compartmentId := getRequiredEnvSetting("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_public_ip.test_public_ip"
	datasourceName := "data.oci_core_public_ips.test_public_ips"
	sDatasourceNameById := "data.oci_core_public_ip.test_oci_core_public_ip_by_id"
	sDatasourceNameByIp := "data.oci_core_public_ip.test_oci_core_public_ip_by_ip"
	sDatasourceNameByPrivateIpId := "data.oci_core_public_ip.test_oci_core_public_ip_by_private_ip_id"

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
				Config: config + PublicIpPropertyVariables + compartmentIdVariableStr + PublicIpRequiredOnlyResource + `
					data "oci_core_public_ip" "test_oci_core_public_ip_by_id" {
						id = "${oci_core_public_ip.test_public_ip.id}"
					}

					data "oci_core_public_ip" "test_oci_core_public_ip_by_ip" {
						ip_address = "${oci_core_public_ip.test_public_ip.ip_address}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
					resource.TestCheckNoResourceAttr(resourceName, "private_ip_id"),
					resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAvailable)),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},

					// check oci_core_public_ip by id
					resource.TestCheckResourceAttr(sDatasourceNameById, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(sDatasourceNameById, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
					resource.TestCheckResourceAttrSet(sDatasourceNameById, "id"),
					resource.TestCheckResourceAttrSet(sDatasourceNameById, "ip_address"),
					resource.TestCheckNoResourceAttr(sDatasourceNameById, "private_ip_id"),
					resource.TestCheckResourceAttrSet(sDatasourceNameById, "display_name"),
					resource.TestCheckNoResourceAttr(sDatasourceNameById, "availability_domain"),
					resource.TestCheckResourceAttrSet(sDatasourceNameById, "time_created"),
					resource.TestCheckResourceAttr(sDatasourceNameById, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckResourceAttr(sDatasourceNameById, "state", string(oci_core.PublicIpLifecycleStateAvailable)),

					// check oci_core_public_ip by public ip
					resource.TestCheckResourceAttr(sDatasourceNameByIp, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(sDatasourceNameByIp, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
					resource.TestCheckResourceAttrSet(sDatasourceNameByIp, "id"),
					resource.TestCheckResourceAttrSet(sDatasourceNameByIp, "ip_address"),
					resource.TestCheckNoResourceAttr(sDatasourceNameByIp, "private_ip_id"),
					resource.TestCheckResourceAttrSet(sDatasourceNameByIp, "display_name"),
					resource.TestCheckNoResourceAttr(sDatasourceNameByIp, "availability_domain"),
					resource.TestCheckResourceAttrSet(sDatasourceNameByIp, "time_created"),
					resource.TestCheckResourceAttr(sDatasourceNameByIp, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckResourceAttr(sDatasourceNameByIp, "state", string(oci_core.PublicIpLifecycleStateAvailable)),
				),
			},

			// delete before next create
			{
				Config: config + PublicIpPropertyVariables + compartmentIdVariableStr + PublicIpResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + PublicIpPropertyVariables + compartmentIdVariableStr + PublicIpResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "-tf-public-ip"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "lifetime", "RESERVED"),
					TestCheckResourceAttributesEqual(resourceName, "private_ip_id", "data.oci_core_private_ips.test_private_ips", privateIpId),
					resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAssigned)),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters (partial update)
			{
				Config: config + PublicIpPropertyVariables + compartmentIdVariableStr + PublicIpResourceDependencies + `
					resource "oci_core_public_ip" "test_public_ip" {
						#Required
						compartment_id = "${var.compartment_id}"
						lifetime = "${var.public_ip_lifetime}"

						#Optional
						display_name = "` + displayName2 + `"
						defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.public_ip_defined_tags_value}")}"
						freeform_tags = "${var.public_ip_freeform_tags}"
						private_ip_id = "${data.oci_core_private_ips.test_private_ips.` + privateIpId + `}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName2),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
					TestCheckResourceAttributesEqual(resourceName, "private_ip_id", "data.oci_core_private_ips.test_private_ips", privateIpId),
					resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAssigned)),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates to updatable parameters (full update)
			{
				Config: config + compartmentIdVariableStr + PublicIpResourceDependencies + `

					variable "public_ip_defined_tags_value" { default = "updatedValue" }
					variable "public_ip_display_name" { default = "-tf-public-ip-updated" }
					variable "public_ip_freeform_tags" { default = {"Department"= "Accounting"} }
					variable "public_ip_lifetime" { default = "RESERVED" }
					variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeRegion) + `" }

					resource "oci_core_public_ip" "test_public_ip" {
						#Required
						compartment_id = "${var.compartment_id}"
						lifetime = "${var.public_ip_lifetime}"

						#Optional
						display_name = "${var.public_ip_display_name}"
						defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.public_ip_defined_tags_value}")}"
						freeform_tags = "${var.public_ip_freeform_tags}"
						private_ip_id = "${data.oci_core_private_ips.test_private_ips.` + privateIpId2 + `}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "-tf-public-ip-updated"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "lifetime", "RESERVED"),
					TestCheckResourceAttributesEqual(resourceName, "private_ip_id", "data.oci_core_private_ips.test_private_ips", privateIpId2),
					resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAssigned)),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify updates to updatable parameters (unassign private ip id)
			{
				Config: config + PublicIpPropertyVariables + compartmentIdVariableStr + PublicIpUnassignedResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
					resource.TestCheckResourceAttr(resourceName, "private_ip_id", ""), // Still defined, but now empty.
					resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAvailable)),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config + `
					variable "public_ip_defined_tags_value" { default = "updatedValue" }
					variable "public_ip_display_name" { default = "-tf-public-ip-updated" }
					variable "public_ip_freeform_tags" { default = {"Department"= "Accounting"} }
					variable "public_ip_lifetime" { default = "RESERVED" }
					variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeRegion) + `" }

					resource "oci_core_public_ip" "test_public_ip2" {
						#Required
						compartment_id = "${var.compartment_id}"
						lifetime = "${var.public_ip_lifetime}"

						#Optional
						display_name = "` + displayName2 + `"
						defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.public_ip_defined_tags_value}")}"
						freeform_tags = "${var.public_ip_freeform_tags}"
						private_ip_id = "${data.oci_core_private_ips.test_private_ips.` + privateIpId2 + `}"
					}

					data "oci_core_public_ips" "test_public_ips" {
						#Required
						compartment_id = "${var.compartment_id}"
						scope = "${var.public_ip_scope}"

						filter {
							name = "id"
							values = ["${oci_core_public_ip.test_public_ip2.id}"]
						}
					}

					data "oci_core_public_ip" "test_oci_core_public_ip_by_private_ip_id" {
						private_ip_id = "${oci_core_public_ip.test_public_ip2.private_ip_id}"
					}
					` + compartmentIdVariableStr + PublicIpResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckResourceAttr(datasourceName, "public_ips.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "public_ips.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "public_ips.0.display_name", "-tf-public-ip-updated"),
					resource.TestCheckResourceAttr(datasourceName, "public_ips.0.lifetime", string(oci_core.PublicIpLifetimeReserved)),
					resource.TestCheckResourceAttrSet(datasourceName, "public_ips.0.private_ip_id"),

					// check oci_core_public_ip by private ip id
					resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
					resource.TestCheckResourceAttrSet(sDatasourceNameByPrivateIpId, "id"),
					resource.TestCheckResourceAttrSet(sDatasourceNameByPrivateIpId, "ip_address"),
					resource.TestCheckResourceAttrSet(sDatasourceNameByPrivateIpId, "private_ip_id"),
					resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "defined_tags.%", "1"),
					resource.TestCheckResourceAttrSet(sDatasourceNameByPrivateIpId, "display_name"),
					resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "freeform_tags.%", "1"),
					resource.TestCheckNoResourceAttr(sDatasourceNameByPrivateIpId, "availability_domain"),
					resource.TestCheckResourceAttrSet(sDatasourceNameByPrivateIpId, "time_created"),
					resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckResourceAttr(sDatasourceNameByPrivateIpId, "state", string(oci_core.PublicIpLifecycleStateAssigned)),
				),
			},
			// Test client-side filtering.
			{
				Config: config + `
					variable "public_ip_defined_tags_value" { default = "value" }
					variable "public_ip_display_name" { default = "-tf-public-ip" }
					variable "public_ip_freeform_tags" { default = {"Department"= "Finance"} }
					variable "public_ip_lifetime" { default = "` + string(oci_core.PublicIpLifetimeEphemeral) + `" }
					variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeAvailabilityDomain) + `" }

					data "oci_core_public_ips" "test_public_ips" {
						#Required
						compartment_id = "${var.compartment_id}"
						scope = "${var.public_ip_scope}"

						#Optional
						availability_domain = "${data.oci_core_private_ips.test_private_ips.private_ips.0.availability_domain}"
						
						filter {
							name = "lifetime"
							values = ["` + string(oci_core.PublicIpLifetimeReserved) + `"]
						}
					}` + compartmentIdVariableStr + PublicIpResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "scope", string(oci_core.PublicIpScopeAvailabilityDomain)),
					resource.TestCheckResourceAttr(datasourceName, "public_ips.#", "0"),
				),
			},
		},
	})
}
