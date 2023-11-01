// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "stack_mon_management_agent_id_resource1" {}
variable "stack_mon_hostname_resource1" {}
variable "publish_trigger" {default = true} // setting to true so that enable testing can happen on a published ME


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_stack_monitoring_metric_extension" "test_metric_extension_on_given_resource" {
  #Required
  compartment_id = var.compartment_ocid
  name = "ME_MetricExtensionEnableTerraformExample"
  resource_type = "host_linux"
  display_name = "OS File System Utilization"
  collection_recurrences = "FREQ=HOURLY;INTERVAL=6"
  metric_list {
    name = "MountPoint"
    is_dimension = true
    data_type = "STRING"
    is_hidden = false
  }
  metric_list {
    name = "FileSystemSize"
    is_dimension = false
    data_type = "NUMBER"
    is_hidden = true
  }
  metric_list {
    name = "FileSystemUsed"
    is_dimension = false
    data_type = "NUMBER"
    is_hidden = true
  }
  metric_list {
    name = "FileSystemUsage"
    display_name = "NumberOfUpInstances"
    is_dimension = false
    data_type = "NUMBER"
    is_hidden = false
    metric_category = "UTILIZATION"
    unit = "percent"
    compute_expression = "(FileSystemUsed / FileSystemSize) * 100"
  }
  query_properties {
    collection_method = "OS_COMMAND"
    command = "/bin/bash"
    script_details {
      name = "fileSystem.sh"
      content = "IyEvYmluL2Jhc2gKIyBDb3B5cmlnaHQgKGMpIDIwMjIsIE9yYWNsZSBhbmQvb3IgaXRzIGFmZmlsaWF0ZXMuIEFsbCByaWdodHMgcmVzZXJ2ZWQuCiMKIyBTdGFjayBNb25pdG9yaW5nIC8gSG9zdDogY29sbGVjdCBmaWxlc3lzdGVtIHN0YXRpc3RpY3MgZnJvbSBMaW51eCBob3N0cwojCiMgT3V0cHV0IGZvcm1hdDoKIwojIHJlc3VsdD1tb3VudHxzaXplfHVzZWQKCmV4ZWMgMTA+JjEKZXhlYyAxPiYyCgoKd2hpbGUgcmVhZCAtciBkZXYgc2l6ZSB1c2VkIGF2YWlsIHVzZWRwIG1vdW50IG90aGVyCmRvCiAgICBpZiBbWyAiJHtkZXZ9IiA9fiAvIF1dCiAgICB0aGVuCiAgICAgICAgaWYgWyAiJHt0b3R9IiA9PSAiMCIgXQogICAgICAgIHRoZW4KICAgICAgICAgICAgIyBQcmV2ZW50IGRldmlzaW9uIGJ5IHplcm8KICAgICAgICAgICAgdXNlZD0wCiAgICAgICAgICAgIHVzZWRwPTAKICAgICAgICBmaQoKICAgICAgICBwcmludGYgIm9jaV9yZXN1bHQ9JXN8JXN8JXNcbiIgIiR7bW91bnR9IiAiJHtzaXplfSIgIiR7dXNlZH0iID4mMTAKICAgIGZpCmRvbmUgPCA8KGRmIC1rIDI+L2Rldi9udWxsKQ=="
    }
    delimiter = "|"
    starts_with = "oci_result="
  }
  #Optional
  description = "Computes File System Utilization Percentage of various mount points"
  publish_trigger = var.publish_trigger

}

resource "oci_stack_monitoring_monitored_resource" "test_monitored_resource_enable_example" {
  #Required
  compartment_id = var.compartment_ocid
  name = "ResourceMetricExtEnableTerraformExample"
  type = "host"

  #Optional
  display_name = "ResourceMetricExtEnableTerraformExample"
  host_name = var.stack_mon_hostname_resource1
  management_agent_id = var.stack_mon_management_agent_id_resource1
  properties {
    name = "osName"
    value = "Linux"
  }
  properties {
    name = "osVersion"
    value = "7.0"
  }
  resource_time_zone = "en"
  lifecycle {
    ignore_changes = [
      credentials,
      properties,
      external_id,
      defined_tags]
  }
}

resource "oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management" "test_metric_extension_on_given_resource_management" {
  #Required
  metric_extension_id = oci_stack_monitoring_metric_extension.test_metric_extension_on_given_resource.id
  resource_ids = [
    oci_stack_monitoring_monitored_resource.test_monitored_resource_enable_example.id]
  enable_metric_extension_on_given_resources = true
}

