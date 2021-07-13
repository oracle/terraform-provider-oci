// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "apm_domain_description" {
  default = "description"
}

variable "apm_domain_display_name" {
  default = "displayName"
}

variable "apm_domain_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "apm_domain_is_free_tier" {
  default = false
}

variable "apm_domain_state" {
  default = "ACTIVE"
}

variable "data_key_data_key_type" {
  default = "PRIVATE"
}

variable "monitor_configuration_config_type" {
  default = "SCRIPTED_BROWSER_CONFIG"
}

variable "monitor_configuration_is_certificate_validation_enabled" {
  default = false
}

variable "monitor_configuration_is_failure_retried" {
  default = false
}

variable "monitor_configuration_is_redirection_enabled" {
  default = false
}

variable "monitor_display_name" {
  default = "displayName"
}

variable "monitor_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "monitor_monitor_type" {
  default = "SCRIPTED_BROWSER"
}

variable "monitor_repeat_interval_in_seconds" {
  default = 600
}

variable "monitor_script_parameters_param_name" {
  default = "testName"
}

variable "monitor_script_parameters_param_value" {
  default = "myTest1"
}

variable "monitor_status" {
  default = "ENABLED"
}

variable "monitor_target" {
  default = "https://console.us-ashburn-1.oraclecloud.com"
}

variable "monitor_timeout_in_seconds" {
  default = 60
}

variable "monitor_tag_name" {
  default =  "tagName"
}

variable "monitor_tag_value" {
  default =  "tagValue"
}

variable "monitor_configuration_network_configuration_number_of_hops" {
  default = 10
}

variable "monitor_configuration_network_configuration_probe_mode" {
  default = "SACK"
}

variable "monitor_configuration_network_configuration_probe_per_hop" {
  default = 10
}

variable "monitor_configuration_network_configuration_protocol" {
  default = "TCP"
}

variable "monitor_configuration_network_configuration_transmission_rate" {
  default = 10
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apm_synthetics_monitor" "test_monitor" {
  #Required
  apm_domain_id              = oci_apm_apm_domain.test_apm_domain.id
  display_name               = var.monitor_display_name
  monitor_type               = var.monitor_monitor_type
  repeat_interval_in_seconds = var.monitor_repeat_interval_in_seconds
  vantage_points   = ["OraclePublic-us-ashburn-1"]

  #Optional
  configuration {

    #Optional
    config_type                       = var.monitor_configuration_config_type
    is_certificate_validation_enabled = var.monitor_configuration_is_certificate_validation_enabled
    is_failure_retried                = var.monitor_configuration_is_failure_retried
    is_redirection_enabled            = var.monitor_configuration_is_redirection_enabled

    #Optional
    network_configuration {
      number_of_hops           = var.monitor_configuration_network_configuration_number_of_hops
      probe_mode               = var.monitor_configuration_network_configuration_probe_mode
      probe_per_hop            = var.monitor_configuration_network_configuration_probe_per_hop
      protocol                 = var.monitor_configuration_network_configuration_protocol
      transmission_rate        = var.monitor_configuration_network_configuration_transmission_rate
    }
  }
  freeform_tags = var.monitor_freeform_tags
  script_id     = oci_apm_synthetics_script.test_script.id
  script_parameters {
    #Required
    param_name  = var.monitor_script_parameters_param_name
    param_value = var.monitor_script_parameters_param_value
  }
  status             = var.monitor_status
  target             = var.monitor_target
  timeout_in_seconds = var.monitor_timeout_in_seconds
}

data "oci_apm_synthetics_monitors" "test_monitors" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  display_name = var.monitor_display_name
  monitor_type = var.monitor_monitor_type
  script_id    = oci_apm_synthetics_script.test_script.id
  status       = var.monitor_status
}


resource "oci_apm_apm_domain" "test_apm_domain" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.apm_domain_display_name

  #Optional
  description   = var.apm_domain_description
  freeform_tags = var.apm_domain_freeform_tags
  is_free_tier  = var.apm_domain_is_free_tier
}

variable "script_content" {
  default = "{ \"id\":\"f672ea8c-9508-483e-a123-878920eee73c\", \"version\":\"2.0\", \"name\":\"Sample Project\", \"url\":\"https://console.us-ashburn-1.oraclecloud.com\", \"tests\": [  { \"id\":\"b4522766-e382-40c2-ab01-452cf62e1cec\", \"name\":\"<ORAP><ON>testName</ON><OV>myTest</OV><OS>false</OS></ORAP>\", \"commands\":[ { \"id\":\"d1bc2093-bb61-4919-a554-38ef2653ac02\", \"comment\":\"comment\", \"command\":\"open\", \"target\":\"/\", \"targets\":[[\"css=td.bodytext\",\"css\"]], \"value\":\"xyz\"  } ] } ], \"suites\": [ { \"id\":\"a86b2934-7aa3-4838-b389-93c8aea2af05\",  \"name\":\"Default Suite\",  \"persistSession\":false, \"parallel\":false, \"timeout\":600,  \"tests\":  [  \"b4522766-e382-40c2-ab01-452cf62e1cec\" ] } ], \"urls\": [ \"https://console.us-ashburn-1.oraclecloud.com/\"  ], \"plugins\":[\"xxx\"] }"
}

variable "script_content_type" {
  default = "SIDE"
}

variable "script_display_name" {
  default = "displayName"
}

variable "script_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "script_parameters_is_secret" {
  default = false
}

variable "script_parameters_param_name" {
  default = "testName"
}

variable "script_parameters_param_value" {
  default = "myTest1"
}



resource "oci_apm_synthetics_script" "test_script" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  content       = var.script_content
  content_type  = var.script_content_type
  display_name  = var.script_display_name

  #Optional
  freeform_tags     = var.script_freeform_tags
  parameters {
    #Required
    param_name = var.script_parameters_param_name

    #Optional
    is_secret   = var.script_parameters_is_secret
    param_value = var.script_parameters_param_value
  }
}

data "oci_apm_synthetics_scripts" "test_scripts" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  content_type = var.script_content_type
  display_name = var.script_display_name
}

