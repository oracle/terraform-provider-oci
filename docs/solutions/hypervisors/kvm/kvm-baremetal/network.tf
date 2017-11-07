resource "oci_core_virtual_network" "kvm-vcn" {
  cidr_block     = "${var.vcn_cidr_block}"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "${var.prefix}-vcn"
  dns_label      = "kvmvcn"
}

resource "oci_core_internet_gateway" "kvm-ig" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "${var.prefix}-ig"
  vcn_id         = "${oci_core_virtual_network.kvm-vcn.id}"
}

resource "oci_core_route_table" "kvm-rt" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.kvm-vcn.id}"
  display_name   = "${var.prefix}-kvm-rt"

  route_rules {
    cidr_block        = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.kvm-ig.id}"
  }
}

resource "oci_core_subnet" "kvm-host-subnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads.availability_domains[var.availability_domain - 1],"name")}"
  cidr_block          = "${var.kvm_host_subnet_cidr_block}"
  display_name        = "kvm-host-subnet"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.kvm-vcn.id}"
  route_table_id      = "${oci_core_route_table.kvm-rt.id}"
  security_list_ids   = ["${oci_core_security_list.kvm-host-security_list.id}"]
  dhcp_options_id     = "${oci_core_virtual_network.kvm-vcn.default_dhcp_options_id}"
  dns_label           = "kvmhostsubnet"
}

# Protocols are specified as protocol numbers.
# http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
resource "oci_core_security_list" "kvm-host-security_list" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.kvm-vcn.id}"
  display_name   = "kvm-mgmt-security-list"

  // allow outbound tcp traffic on all ports
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "6"
  }

  // allow inbound http (port 443) traffic
  ingress_security_rules {
    protocol  = "6"         // tcp
    source    = "0.0.0.0/0"
    stateless = false

    tcp_options {
      "min" = 443
      "max" = 443
    }
  }

  // allow inbound traffic to port 5901 (vnc)
  ingress_security_rules {
    protocol  = "6"         // tcp
    source    = "0.0.0.0/0"
    stateless = false

    tcp_options {
      "min" = 5901
      "max" = 5901
    }
  }

  // allow inbound ssh traffic
  ingress_security_rules {
    protocol  = "6"         // tcp
    source    = "0.0.0.0/0"
    stateless = false

    tcp_options {
      "min" = 22
      "max" = 22
    }
  }

  // allow inbound icmp traffic of a specific type
  ingress_security_rules {
    protocol  = 1
    source    = "0.0.0.0/0"
    stateless = true

    icmp_options {
      "type" = 3
      "code" = 4
    }
  }
}
