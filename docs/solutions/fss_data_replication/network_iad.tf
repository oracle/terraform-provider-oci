resource "oci_core_virtual_network" "my_vcn_iad" {
  provider       = "oci.iad"
  cidr_block     = "${var.my_vcn_cidr_iad}"
  dns_label      = "fssdemoiad"
  compartment_id = "${var.compartment_id}"
  display_name   = "fssdemo_iad"
  dns_label      = "fssdemoiad"
}

resource "oci_core_drg" "drg_iad" {
  provider       = "oci.iad"
  compartment_id = "${var.compartment_id}"
  display_name   = "drg_iad"
}

resource "oci_core_drg_attachment" "drg_attch_iad" {
  provider     = "oci.iad"
  drg_id       = "${oci_core_drg.drg_iad.id}"
  vcn_id       = "${oci_core_virtual_network.my_vcn_iad.id}"
  display_name = "drg_attch_iad"
}

resource "oci_core_remote_peering_connection" "remote_peering_iad" {
  provider       = "oci.iad"
  compartment_id = "${var.compartment_id}"
  drg_id         = "${oci_core_drg.drg_iad.id}"
  display_name   = "remote_peering_iad"
}

resource "oci_core_subnet" "subnet_iad_ad1" {
  provider            = "oci.iad"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[0],"name")}"
  cidr_block          = "${var.subnet_cidr_iad_ad1}"
  display_name        = "mysubnetAD1_iad"
  dns_label           = "subnetiadad1"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_virtual_network.my_vcn_iad.id}"
  security_list_ids   = ["${oci_core_security_list.my_security_list_iad.id}"]
}

resource "oci_core_subnet" "subnet_iad_ad2" {
  provider            = "oci.iad"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ads_iad.availability_domains[1],"name")}"
  cidr_block          = "${var.subnet_cidr_iad_ad2}"
  display_name        = "mysubnetAD2_iad"
  dns_label           = "subnetiadad2"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_virtual_network.my_vcn_iad.id}"
  security_list_ids   = ["${oci_core_security_list.my_security_list_iad.id}"]
}

resource "oci_core_internet_gateway" "igw_iad" {
  provider       = "oci.iad"
  compartment_id = "${var.compartment_id}"
  display_name   = "igw_iad"
  vcn_id         = "${oci_core_virtual_network.my_vcn_iad.id}"
}

resource "oci_core_default_route_table" "default_route_table_iad" {
  provider                   = "oci.iad"
  manage_default_resource_id = "${oci_core_virtual_network.my_vcn_iad.default_route_table_id}"
  display_name               = "fssdemo_default_route_table_iad"

  route_rules {
    cidr_block        = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.igw_iad.id}"
  }

  route_rules {
    cidr_block        = "${var.subnet_cidr_phx_ad1}"
    network_entity_id = "${oci_core_drg.drg_iad.id}"
  }
}

# Protocols are specified as protocol numbers.
# http://www.iana.org/assignments/protocol_numbers/protocol_numbers.xhtml

resource "oci_core_security_list" "my_security_list_iad" {
  provider       = "oci.iad"
  compartment_id = "${var.compartment_id}"
  display_name   = "my_security_list_iad"
  vcn_id         = "${oci_core_virtual_network.my_vcn_iad.id}"

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
    {
      protocol = "all"
      source   = "${var.my_vcn_cidr_phx}"
    },
  ]
}
