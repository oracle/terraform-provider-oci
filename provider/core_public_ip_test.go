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
	compartment_id = "${var.test_compartment_id}"
	lifetime = "${var.public_ip_lifetime}"
}
`

	PublicIpResourceConfig = PublicIpResourceDependencies + `
resource "oci_core_public_ip" "test_public_ip" {
	#Required
	compartment_id = "${var.test_compartment_id}"
	lifetime = "${var.public_ip_lifetime}"

	#Optional
	display_name = "${var.public_ip_display_name}"
	private_ip_id = "${data.oci_core_private_ips.test_private_ips.` + privateIpId + `}"
}
`
	PublicIpUnassignedResourceConfig = PublicIpResourceDependencies + `
resource "oci_core_public_ip" "test_public_ip" {
	#Required
	compartment_id = "${var.test_compartment_id}"
	lifetime = "${var.public_ip_lifetime}"

	#Optional
	display_name = "${var.public_ip_display_name}"
}
`
	PublicIpPropertyVariables = `
variable "public_ip_display_name" { default = "` + displayName + `" }
variable "public_ip_lifetime" { default = "` + string(oci_core.PublicIpLifetimeReserved) + `" }
variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeRegion) + `" }
`
	// TODO: Replace dependencies' config with auto-generated ones. This requires changing the
	//   "compartment_id" var to "test_compartment_id".
	PublicIpResourceDependencies = `
	variable "InstanceImageOCID" {
		type = "map"
		default = {
			// Oracle-provided image "Oracle-Linux-7.4-2017.12.18-0"
			us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaasc56hnpnx7swoyd2fw5gyvbn3kcdmqc2guiiuvnztl2erth62xnq"
			us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaaxrqeombwty6jyqgk3fraczdd63bv66xgfsqka4ktr7c57awr3p5a"
			eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaayxmzu6n5hsntq4wlffpb4h6qh6z3uskpbm5v3v4egqlqvwicfbyq"
		}
	}
	data "oci_identity_availability_domains" "test_availability_domains" {
		compartment_id = "${var.test_compartment_id}"
	}
	resource "oci_core_virtual_network" "test_vcn" {
		cidr_block      = "10.0.0.0/16"
		compartment_id  = "${var.test_compartment_id}"
		display_name    = "-tf-vcn"
		dns_label       = "testvcn"
	}
	resource "oci_core_subnet" "test_subnet" {
		availability_domain = "${lookup(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0],"name")}"
		cidr_block          = "10.0.1.0/24"
		display_name        = "-tf-subnet"
		compartment_id      = "${var.test_compartment_id}"
		vcn_id              = "${oci_core_virtual_network.test_vcn.id}"
		route_table_id      = "${oci_core_virtual_network.test_vcn.default_route_table_id}"
		security_list_ids   = ["${oci_core_virtual_network.test_vcn.default_security_list_id}"]
		dhcp_options_id     = "${oci_core_virtual_network.test_vcn.default_dhcp_options_id}"
		dns_label           = "testsubnet"
	}
	resource "oci_core_instance" "test_instance" {
		availability_domain = "${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}"
		compartment_id      = "${var.test_compartment_id}"
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
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"test_compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_public_ip.test_public_ip"
	datasourceName := "data.oci_core_public_ips.test_public_ips"

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
				Config:            config + PublicIpPropertyVariables + compartmentIdVariableStr + PublicIpRequiredOnlyResource,
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
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
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
						compartment_id = "${var.test_compartment_id}"
						lifetime = "${var.public_ip_lifetime}"

						#Optional
						display_name = "` + displayName2 + `"
						private_ip_id = "${data.oci_core_private_ips.test_private_ips.` + privateIpId + `}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName2),
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
				Config: config + PublicIpPropertyVariables + compartmentIdVariableStr + PublicIpResourceDependencies + `
					resource "oci_core_public_ip" "test_public_ip" {
						#Required
						compartment_id = "${var.test_compartment_id}"
						lifetime = "${var.public_ip_lifetime}"

						#Optional
						display_name = "${var.public_ip_display_name}"
						private_ip_id = "${data.oci_core_private_ips.test_private_ips.` + privateIpId2 + `}"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
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
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
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
			// verify updates to Force New parameters.
			{
				// Don't change compartment IDs for this test. Public IPs cannot reference a private IP in another compartment.
				Config: config + `
					variable "public_ip_display_name" { default = "` + displayName + `" }
					variable "public_ip_lifetime" { default = "` + string(oci_core.PublicIpLifetimeEphemeral) + `" }
					variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeAvailabilityDomain) + `" }
				` + compartmentIdVariableStr + PublicIpResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeEphemeral)),
					TestCheckResourceAttributesEqual(resourceName, "private_ip_id", "data.oci_core_private_ips.test_private_ips", privateIpId),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_address"),
					resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeAvailabilityDomain)),
					resource.TestCheckResourceAttr(resourceName, "state", string(oci_core.PublicIpLifecycleStateAssigned)),

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
					variable "public_ip_display_name" { default = "` + displayName + `" }
					variable "public_ip_lifetime" { default = "` + string(oci_core.PublicIpLifetimeEphemeral) + `" }
					variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeAvailabilityDomain) + `" }

					resource "oci_core_public_ip" "test_public_ip2" {
						#Required
						compartment_id = "${var.test_compartment_id}"
						lifetime = "${var.public_ip_lifetime}"

						#Optional
						display_name = "` + displayName2 + `"
						private_ip_id = "${data.oci_core_private_ips.test_private_ips.` + privateIpId2 + `}"
					}

					data "oci_core_public_ips" "test_public_ips" {
						#Required
						compartment_id = "${var.test_compartment_id}"
						scope = "${var.public_ip_scope}"

						#Optional
						availability_domain = "${data.oci_core_private_ips.test_private_ips.private_ips.0.availability_domain}"
						
						filter {
							name = "id"
							values = ["${oci_core_public_ip.test_public_ip.id}"]
						}
					}` + compartmentIdVariableStr + PublicIpResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "scope", string(oci_core.PublicIpScopeAvailabilityDomain)),
					resource.TestCheckResourceAttr(datasourceName, "public_ips.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "public_ips.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "public_ips.0.display_name", displayName),
					resource.TestCheckResourceAttr(datasourceName, "public_ips.0.lifetime", string(oci_core.PublicIpLifetimeEphemeral)),
					resource.TestCheckResourceAttrSet(datasourceName, "public_ips.0.private_ip_id"),
				),
			},
			// Test client-side filtering.
			{
				Config: config + `
					variable "public_ip_display_name" { default = "` + displayName + `" }
					variable "public_ip_lifetime" { default = "` + string(oci_core.PublicIpLifetimeEphemeral) + `" }
					variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeAvailabilityDomain) + `" }

					data "oci_core_public_ips" "test_public_ips" {
						#Required
						compartment_id = "${var.test_compartment_id}"
						scope = "${var.public_ip_scope}"

						#Optional
						availability_domain = "${data.oci_core_private_ips.test_private_ips.private_ips.0.availability_domain}"
						
						filter {
							name = "lifetime"
							values = ["` + string(oci_core.PublicIpLifetimeReserved) + `"]
						}
					}` + compartmentIdVariableStr + PublicIpResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "scope", string(oci_core.PublicIpScopeAvailabilityDomain)),
					resource.TestCheckResourceAttr(datasourceName, "public_ips.#", "0"),
				),
			},
		},
	})
}

func TestCorePublicIpResource_forcenew(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getRequiredEnvSetting("compartment_id_for_create")
	compartmentIdVariableStr := fmt.Sprintf("variable \"test_compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"test_compartment_id\" { default = \"%s\" }\n", compartmentId2)

	resourceName := "oci_core_public_ip.test_public_ip"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create with optionals
			{
				Config: config + `
					variable "public_ip_display_name" { default = "` + displayName + `" }
					variable "public_ip_lifetime" { default = "` + string(oci_core.PublicIpLifetimeReserved) + `" }
					variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeRegion) + `" }
				` + compartmentIdVariableStr + PublicIpUnassignedResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
					resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckNoResourceAttr(resourceName, "private_ip_id"),
					resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// force new tests, test that changing a parameter would result in creation of a new resource.

			{
				Config: config + `
					variable "public_ip_display_name" { default = "` + displayName + `" }
					variable "public_ip_lifetime" { default = "` + string(oci_core.PublicIpLifetimeReserved) + `" }
					variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeRegion) + `" }
				` + compartmentIdVariableStr2 + PublicIpUnassignedResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeReserved)),
					resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeRegion)),
					resource.TestCheckNoResourceAttr(resourceName, "private_ip_id"),
					resource.TestCheckNoResourceAttr(resourceName, "availability_domain"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter CompartmentId but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},

			{
				Config: config + `
					variable "public_ip_display_name" { default = "` + displayName + `" }
					variable "public_ip_lifetime" { default = "` + string(oci_core.PublicIpLifetimeEphemeral) + `" }
					variable "public_ip_scope" { default = "` + string(oci_core.PublicIpScopeAvailabilityDomain) + `" }
				` + compartmentIdVariableStr2 + PublicIpResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId2),
					resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
					resource.TestCheckResourceAttr(resourceName, "lifetime", string(oci_core.PublicIpLifetimeEphemeral)),
					resource.TestCheckResourceAttr(resourceName, "scope", string(oci_core.PublicIpScopeAvailabilityDomain)),
					resource.TestCheckResourceAttrSet(resourceName, "private_ip_id"),
					resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource was expected to be recreated when updating parameter Lifetime but the id did not change.")
						}
						resId = resId2
						return err
					},
				),
			},
		},
	})
}
