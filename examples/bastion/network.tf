resource "oci_core_vcn" "test_bastion_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
  dns_label      = "testvcn"
}

resource "oci_core_service_gateway" "test_bastion_service_gateway" {
  compartment_id = var.compartment_ocid
  display_name   = "sgw"

  services {
    service_id = data.oci_core_services.test_bastion_services.services[1]["id"]
  }

  vcn_id = oci_core_vcn.test_bastion_vcn.id
}

resource "oci_core_default_route_table" "bastion_default_route_table" {
  manage_default_resource_id = oci_core_vcn.test_bastion_vcn.default_route_table_id
  display_name               = "DefaultRouteTable"

  route_rules {
    destination       = lookup(data.oci_core_services.test_bastion_services.services[1], "cidr_block")
    destination_type  = "SERVICE_CIDR_BLOCK"
    network_entity_id = oci_core_service_gateway.test_bastion_service_gateway.id
  }
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domain.bastion_ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "TestSubnet"
  dns_label           = "testsubnet"
  security_list_ids   = [oci_core_vcn.test_bastion_vcn.default_security_list_id]
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_bastion_vcn.id
  route_table_id      = oci_core_vcn.test_bastion_vcn.default_route_table_id
  dhcp_options_id     = oci_core_vcn.test_bastion_vcn.default_dhcp_options_id
}