data "oci_core_fast_connect_provider_services" "fast_connect_provider_services" {
  #Required
  compartment_id = "${var.compartment_ocid}"
}

output "fast_connect_provider_services" {
  value = "${data.oci_core_fast_connect_provider_services.fast_connect_provider_services.fast_connect_provider_services}"
}

data "oci_core_fast_connect_provider_services" "fast_connect_provider_services_private_layer2" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  filter {
    name   = "type"
    values = ["LAYER2"]
  }

  filter {
    name   = "private_peering_bgp_management"
    values = ["CUSTOMER_MANAGED"]
  }

  filter {
    name   = "supported_virtual_circuit_types"
    values = ["${var.virtual_circuit_type_private}"]
  }

  filter {
    name   = "public_peering_bgp_management"
    values = ["ORACLE_MANAGED"]
  }
}

output "fast_connect_provider_services_layer2" {
  value = "${data.oci_core_fast_connect_provider_services.fast_connect_provider_services_private_layer2.fast_connect_provider_services}"
}

data "oci_core_fast_connect_provider_services" "fast_connect_provider_services_public_layer3" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  filter {
    name   = "supported_virtual_circuit_types"
    values = ["${var.virtual_circuit_type_public}"]
  }

  filter {
    name   = "type"
    values = ["LAYER3"]
  }

  filter {
    name   = "private_peering_bgp_management"
    values = ["PROVIDER_MANAGED"]
  }

  filter {
    name   = "public_peering_bgp_management"
    values = ["ORACLE_MANAGED"]
  }
}

output "fast_connect_provider_services_public_layer3" {
  value = "${data.oci_core_fast_connect_provider_services.fast_connect_provider_services_public_layer3.fast_connect_provider_services}"
}
