// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

data "oci_identity_availability_domain" "ad" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 1
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

// Usage of singular instance datasources to show the public_ips, private_ips, and hostname_labels for the instances in the pool
data "oci_core_instance" "TFInstancePoolInstanceSingularDatasources" {
  count       = 2
  instance_id = "${lookup(data.oci_core_instance_pool_instances.TFInstancePoolInstanceDatasources.instances[count.index], "id")}"
}

data "oci_core_instance_pool_load_balancer_attachment" test_instance_pool_load_balancer_attachment {
  instance_pool_id                          = "${oci_core_instance_pool.TFInstancePool.id}"
  instance_pool_load_balancer_attachment_id = "${oci_core_instance_pool.TFInstancePool.load_balancers.0.id}"
}

output "Pooled instances private IPs" {
  value = ["${data.oci_core_instance.TFInstancePoolInstanceSingularDatasources.*.private_ip}"]
}

output "Pooled instances public IPs" {
  value = ["${data.oci_core_instance.TFInstancePoolInstanceSingularDatasources.*.public_ip}"]
}

output "Pooled instances hostname labels" {
  value = ["${data.oci_core_instance.TFInstancePoolInstanceSingularDatasources.*.hostname_label}"]
}

output "Load Balancer backend set name" {
  value = ["${data.oci_core_instance_pool_load_balancer_attachment.test_instance_pool_load_balancer_attachment.backend_set_name}"]
}
