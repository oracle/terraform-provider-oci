# Gets a list of Availability Domains
data "oci_identity_availability_domains" "ADs" {
  compartment_id = "${var.tenancy_ocid}"
}

data "oci_core_instance_configuration" TFInstanceConfigurationDatasource {
  instance_configuration_id = "${oci_core_instance_configuration.TFInstanceConfiguration.id}"
}

data "oci_core_instance_configurations" TFInstanceConfigurationDatasources {
  compartment_id = "${var.compartment_ocid}"

  filter {
    name   = "id"
    values = ["${oci_core_instance_configuration.TFInstanceConfiguration.id}"]
  }
}

data "oci_core_instance_pool" "TFInstancePoolDatasource" {
  instance_pool_id = "${oci_core_instance_pool.TFInstancePool.id}"
}

data "oci_core_instance_pools" "TFInstancePoolDatasources" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFInstancePool"
  state          = "RUNNING"

  filter {
    name   = "id"
    values = ["${oci_core_instance_pool.TFInstancePool.id}"]
  }
}

data "oci_core_instance_pool_instances" "TFInstancePoolInstanceDatasources" {
  compartment_id   = "${var.compartment_ocid}"
  instance_pool_id = "${oci_core_instance_pool.TFInstancePool.id}"
  display_name     = "TFInstancePool"
}
