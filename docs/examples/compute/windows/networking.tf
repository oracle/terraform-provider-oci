#############
# Networking
#############
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_virtual_network" "ExampleVCN" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_internet_gateway" "ExampleInternetGateway" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFExampleInternetGateway"
  vcn_id         = "${oci_core_virtual_network.ExampleVCN.id}"
}

resource "oci_core_route_table" "ExampleRouteTable" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.ExampleVCN.id}"
  display_name   = "TFExampleRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.ExampleInternetGateway.id}"
  }
}

# https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/accessinginstance.htm#one
resource "oci_core_security_list" "ExampleSecurityList" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.ExampleVCN.id}"
  display_name   = "TFExampleSecurityList"

  // allow inbound remote desktop traffic
  ingress_security_rules {
    protocol  = "6"         // tcp
    source    = "0.0.0.0/0"
    stateless = false

    tcp_options {
      // These values correspond to the destination port range.
      "min" = 3389
      "max" = 3389
    }
  }

  // allow inbound winrm traffic
  ingress_security_rules {
    protocol  = "6"         // tcp
    source    = "0.0.0.0/0"
    stateless = false

    tcp_options {
      // These values correspond to the destination port range.
      "min" = 5985
      "max" = 5986
    }
  }

  // allow all outbound traffic
  egress_security_rules {
    protocol    = "all"
    destination = "0.0.0.0/0"
    stateless   = false
  }
}

###########
# Subnet
###########
resource "oci_core_subnet" "ExampleSubnet" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.availability_domain - 1],"name")}"
  cidr_block          = "10.1.20.0/24"
  display_name        = "TFExampleSubnet"
  dns_label           = "tfexamplesubnet"
  security_list_ids   = ["${oci_core_security_list.ExampleSecurityList.id}"]
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.ExampleVCN.id}"
  route_table_id      = "${oci_core_route_table.ExampleRouteTable.id}"
  dhcp_options_id     = "${oci_core_virtual_network.ExampleVCN.default_dhcp_options_id}"
}
