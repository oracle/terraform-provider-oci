/*
 * This example creates a private load balancer on one subnet in a single AD.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}


provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

data "baremetal_identity_availability_domains" "ADs" {
    compartment_id = "${var.tenancy_ocid}"
}


/* Network */

resource "baremetal_core_virtual_network" "vcn1" {
    cidr_block = "10.1.0.0/16"
    compartment_id = "${var.compartment_ocid}"
    display_name = "vcn1"
    dns_label = "vcn1"
}

resource "baremetal_core_subnet" "subnet1" {
    availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.1.20.0/24"
    display_name = "subnet1"
    dns_label = "subnet1"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${baremetal_core_virtual_network.vcn1.id}"
    security_list_ids = ["${baremetal_core_virtual_network.vcn1.default_security_list_id}"]
    route_table_id = "${baremetal_core_virtual_network.vcn1.default_route_table_id}"
    dhcp_options_id = "${baremetal_core_virtual_network.vcn1.default_dhcp_options_id}"
    prohibit_public_ip_on_vnic = true

    provisioner "local-exec" {
        command = "sleep 5"
    }
}


/* Load Balancer */

resource "baremetal_load_balancer" "lb1" {
  shape          = "100Mbps"
  compartment_id = "${var.compartment_ocid}"
  subnet_ids     = [
    "${baremetal_core_subnet.subnet1.id}"
  ]
  display_name = "lb1"
  is_private = true
}

resource "baremetal_load_balancer_backendset" "lb-bes1" {
  name             = "lb-bes1"
  load_balancer_id = "${baremetal_load_balancer.lb1.id}"
  policy           = "ROUND_ROBIN"

  health_checker {
    port     = "80"
    protocol = "HTTP"
    response_body_regex = ".*"
    url_path = "/"
  }

  session_persistence_configuration {
    cookie_name = "lb-session1"
    disable_fallback = true
  }
}

output "lb_private_ip" {
  value = ["${baremetal_load_balancer.lb1.ip_addresses}"]
}
