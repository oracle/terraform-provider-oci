data "oci_core_virtual_circuit_bandwidth_shapes" "virtual_circuit_bandwidth_shapes" {
  #Required
  provider_service_id = "${data.oci_core_fast_connect_provider_services.fast_connect_provider_services.fast_connect_provider_services.0.id}"
}

output "virtual_circuit_bandwidth_shapes" {
  value = "${data.oci_core_virtual_circuit_bandwidth_shapes.virtual_circuit_bandwidth_shapes.virtual_circuit_bandwidth_shapes}"
}
