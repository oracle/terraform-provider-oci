// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# Gets a list of Availability Domains
data "oci_identity_availability_domain" "AD" {
  compartment_id = var.tenancy_ocid
  ad_number      = "1"
}

data "oci_core_instance_configuration" "TFInstanceConfigurationDatasource" {
  instance_configuration_id = oci_core_instance_configuration.TFInstanceConfiguration.id
}

data "oci_core_instance_configurations" "TFInstanceConfigurationDatasources" {
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_core_instance_configuration.TFInstanceConfiguration.id]
  }
}

data "oci_core_instance_pool" "TFInstancePoolDatasource" {
  instance_pool_id = oci_core_instance_pool.TFInstancePool.id
}

data "oci_core_instance_pools" "TFInstancePoolDatasources" {
  compartment_id = var.compartment_ocid
  display_name   = "TFInstancePool"
  state          = "RUNNING"

  filter {
    name   = "id"
    values = [oci_core_instance_pool.TFInstancePool.id]
  }
}

data "oci_core_instance_pool_instances" "TFInstancePoolInstanceDatasources" {
  compartment_id   = var.compartment_ocid
  instance_pool_id = oci_core_instance_pool.TFInstancePool.id
  display_name     = "TFInstancePool"
}

data "oci_autoscaling_auto_scaling_configuration" "TFAutoScalingConfigurationDatasource" {
  auto_scaling_configuration_id = oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfiguration.id
}

data "oci_autoscaling_auto_scaling_configuration" "TFAutoScalingConfigurationScheduledPolicyDatasource" {
  auto_scaling_configuration_id = oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfigurationScheduledPolicy.id
}

data "oci_autoscaling_auto_scaling_configuration" "TFAutoScalingConfigurationScheduledPolicyResourceActionDatasource" {
  auto_scaling_configuration_id = oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfigurationScheduledPolicyResourceAction.id
}

data "oci_autoscaling_auto_scaling_configurations" "TFAutoScalingConfigurationDatasources" {
  compartment_id = var.compartment_ocid
  display_name   = "TFAutoScalingConfiguration"

  filter {
    name   = "id"
    values = [oci_autoscaling_auto_scaling_configuration.TFAutoScalingConfiguration.id]
  }
}

// Usage of singular instance datasources to show the public_ips, private_ips, and hostname_labels for the instances in the pool
data "oci_core_instance" "TFInstancePoolInstanceSingularDatasources" {
  count       = 2
  instance_id = data.oci_core_instance_pool_instances.TFInstancePoolInstanceDatasources.instances[count.index]["id"]
}

output "Pooled_instances_private_IPs" {
  value = [data.oci_core_instance.TFInstancePoolInstanceSingularDatasources.*.private_ip]
}

output "Pooled_instances_public_IPs" {
  value = [data.oci_core_instance.TFInstancePoolInstanceSingularDatasources.*.public_ip]
}

output "Pooled_instances_hostname_labels" {
  value = [data.oci_core_instance.TFInstancePoolInstanceSingularDatasources.*.hostname_label]
}

