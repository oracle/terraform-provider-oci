resource "oci_core_virtual_network" "vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleVCNDBSystem"
  dns_label      = "tfexvcndbsys"
}

resource "oci_core_subnet" "subnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  cidr_block          = "10.1.20.0/24"
  display_name        = "TFExampleSubnetDBSystem"
  dns_label           = "tfexsubdbsys"
  security_list_ids   = ["${oci_core_security_list.ExampleSecurityList.id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.vcn.id}"
  route_table_id      = "${oci_core_route_table.route_table.id}"
  dhcp_options_id     = "${oci_core_virtual_network.vcn.default_dhcp_options_id}"
}

resource "oci_core_subnet" "subnet_backup" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  cidr_block          = "10.1.1.0/24"
  display_name        = "TFExampleSubnetDBSystemBackup"
  dns_label           = "tfexsubdbsysbp"
  security_list_ids   = ["${oci_core_security_list.ExampleSecurityList.id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.vcn.id}"
  route_table_id      = "${oci_core_route_table.route_table_backup.id}"
  dhcp_options_id     = "${oci_core_virtual_network.vcn.default_dhcp_options_id}"
}

resource "oci_core_internet_gateway" "internet_gateway" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleIGDBSystem"
  vcn_id         = "${oci_core_virtual_network.vcn.id}"
}

resource "oci_core_route_table" "route_table" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.vcn.id}"
  display_name   = "TFExampleRouteTableDBSystem"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.internet_gateway.id}"
  }
}

resource "oci_core_route_table" "route_table_backup" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.vcn.id}"
  display_name   = "TFExampleRouteTableDBSystemBackup"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.internet_gateway.id}"
  }
}

resource "oci_core_security_list" "ExampleSecurityList" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.vcn.id}"
  display_name   = "TFExampleSecurityList"

  // allow outbound tcp traffic on all ports
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "6"
  }

  // allow outbound udp traffic on a port range
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "17"        // udp
    stateless   = true
  }

  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "1"
    stateless   = true
  }

  // allow inbound ssh traffic from a specific port
  ingress_security_rules {
    protocol  = "6"         // tcp
    source    = "0.0.0.0/0"
    stateless = false
  }

  // allow inbound icmp traffic of a specific type
  ingress_security_rules {
    protocol  = 1
    source    = "0.0.0.0/0"
    stateless = true
  }
}
