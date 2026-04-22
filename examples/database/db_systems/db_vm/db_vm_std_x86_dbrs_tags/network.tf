resource "oci_core_vcn" "test_vcn" {
  display_name   = "tfVcnForX86DBRSTagsExample"
  cidr_block     = "10.2.0.0/16"
  compartment_id = var.compartment_id
  dns_label      = "tfx86vcn"
}

resource "oci_core_route_table" "test_route_table" {
  display_name   = "tfRouteTable"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id

  route_rules {
    cidr_block        = "0.0.0.0/0"
    description       = "Internal traffic for OCI Services"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  display_name   = "tfInternetGateway"
  compartment_id = var.compartment_id
  enabled        = "true"
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_subnet" "test_subnet" {
  display_name      = "tfPublicSubnet"
  cidr_block        = "10.2.0.0/24"
  compartment_id    = var.compartment_id
  dhcp_options_id   = oci_core_vcn.test_vcn.default_dhcp_options_id
  dns_label         = "tfx86pub"
  route_table_id    = oci_core_route_table.test_route_table.id
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  vcn_id            = oci_core_vcn.test_vcn.id
}

resource "oci_core_security_list" "test_private_subnet_security_list" {
  display_name   = "tfRecoveryServiceSecurityList"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id

  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol    = "all"
  }

  ingress_security_rules {
    protocol    = "6"
    source      = "10.2.0.0/16"
    source_type = "CIDR_BLOCK"

    tcp_options {
      min = "8005"
      max = "8005"
    }
  }

  ingress_security_rules {
    protocol    = "6"
    source      = "10.2.0.0/16"
    source_type = "CIDR_BLOCK"

    tcp_options {
      min = "2484"
      max = "2484"
    }
  }
}

resource "oci_core_service_gateway" "test_service_gateway" {
  display_name   = "tfRecoveryServiceServiceGateway"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id

  services {
    service_id = data.oci_core_services.test_services.services.0.id
  }
}

resource "oci_core_route_table" "test_private_subnet_route_table" {
  display_name   = "tfRecoveryServicePrivateSubnetRouteTable"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id

  route_rules {
    description       = "Recovery Service traffic for OCI Services"
    destination       = data.oci_core_services.test_services.services[0].cidr_block
    destination_type  = "SERVICE_CIDR_BLOCK"
    network_entity_id = oci_core_service_gateway.test_service_gateway.id
  }
}

resource "oci_core_subnet" "test_private_subnet" {
  display_name               = "tfPrivateSubnet"
  cidr_block                 = "10.2.1.0/24"
  compartment_id             = var.compartment_id
  dhcp_options_id            = oci_core_vcn.test_vcn.default_dhcp_options_id
  dns_label                  = "tfx86priv"
  prohibit_public_ip_on_vnic = "true"
  route_table_id             = oci_core_route_table.test_private_subnet_route_table.id
  security_list_ids          = [oci_core_security_list.test_private_subnet_security_list.id]
  vcn_id                     = oci_core_vcn.test_vcn.id
}
