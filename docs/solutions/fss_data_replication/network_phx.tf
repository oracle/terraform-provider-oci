resource "oci_core_virtual_network" "my_vcn_phx" {
  provider       = "oci.phx"
  cidr_block     = "${var.my_vcn_cidr_phx}"
  dns_label      = "fssdemophx"
  compartment_id = "${var.compartment_id}"
  display_name   = "fssdemo_phx"
  dns_label      = "fssdemophx"
}

resource "oci_core_drg" "drg_phx" {
  provider       = "oci.phx"
  compartment_id = "${var.compartment_id}"
  display_name   = "drg_phx"
}

resource "oci_core_drg_attachment" "drg_attch_phx" {
  provider     = "oci.phx"
  drg_id       = "${oci_core_drg.drg_phx.id}"
  vcn_id       = "${oci_core_virtual_network.my_vcn_phx.id}"
  display_name = "drg_attch_phx"
}

resource "oci_core_remote_peering_connection" "remote_peering_phx" {
  provider         = "oci.phx"
  compartment_id   = "${var.compartment_id}"
  drg_id           = "${oci_core_drg.drg_phx.id}"
  peer_id          = "${oci_core_remote_peering_connection.remote_peering_iad.id}"
  peer_region_name = "us-ashburn-1"
  display_name     = "remote_peering_phx"
}

resource "oci_core_subnet" "subnet_phx_ad1" {
  provider            = "oci.phx"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_phx.availability_domains[0],"name")}"
  cidr_block          = "${var.subnet_cidr_phx_ad1}"
  display_name        = "subnet_phx_ad1"
  dns_label           = "subnetphxad1"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_virtual_network.my_vcn_phx.id}"
  security_list_ids   = ["${oci_core_security_list.my_security_list_phx.id}"]
}

resource "oci_core_internet_gateway" "igw_phx" {
  provider       = "oci.phx"
  compartment_id = "${var.compartment_id}"
  display_name   = "igw_phx"
  vcn_id         = "${oci_core_virtual_network.my_vcn_phx.id}"
}

resource "oci_core_default_route_table" "default_route_table_phx" {
  provider                   = "oci.phx"
  manage_default_resource_id = "${oci_core_virtual_network.my_vcn_phx.default_route_table_id}"
  display_name               = "fssdemo_default_route_table_phx"

  route_rules {
    cidr_block        = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.igw_phx.id}"
  }

  route_rules {
    cidr_block        = "${var.subnet_cidr_iad_ad1}"
    network_entity_id = "${oci_core_drg.drg_phx.id}"
  }
}

# Protocols are specified as protocol numbers.
# http://www.iana.org/assignments/protocol_numbers/protocol_numbers.xhtml

resource "oci_core_security_list" "my_security_list_phx" {
  provider       = "oci.phx"
  compartment_id = "${var.compartment_id}"
  display_name   = "my_security_list_phx"
  vcn_id         = "${oci_core_virtual_network.my_vcn_phx.id}"

  // Allow all outbound requests
  egress_security_rules = [
    {
      destination = "0.0.0.0/0"
      protocol    = "all"
    },
  ]

  // See https://docs.us-phoenix-1.oraclecloud.com/Content/File/Tasks/creatingfilesystems.htm.
  // Specific security list rules are required to allow mount targets to work properly.
  ingress_security_rules = [
    {
      protocol = "all"
      source   = "${var.my_vcn_cidr_phx}"
    },
    {
      protocol = "all"
      source   = "${var.my_vcn_cidr_iad}"
    },
    {
      // allow inbound ssh traffic
      protocol  = "6"         // tcp
      source    = "0.0.0.0/0"
      stateless = false

      tcp_options {
        "min" = 22
        "max" = 22
      }
    },
  ]
}
