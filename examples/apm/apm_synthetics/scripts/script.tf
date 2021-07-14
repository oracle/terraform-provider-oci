// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "script_content" {
  default = "{ \"id\":\"f672ea8c-9508-483e-a123-878920eee73c\", \"version\":\"2.0\", \"name\":\"Sample Project\", \"url\":\"https://console.us-ashburn-1.oraclecloud.com\", \"tests\": [  { \"id\":\"b4522766-e382-40c2-ab01-452cf62e1cec\", \"name\":\"<ORAP><ON>testName</ON><OV>myTest</OV><OS>false</OS></ORAP>\", \"commands\":[ { \"id\":\"d1bc2093-bb61-4919-a554-38ef2653ac02\", \"comment\":\"comment\", \"command\":\"open\", \"target\":\"/\", \"targets\":[[\"css=td.bodytext\",\"css\"]], \"value\":\"xyz\"  } ] } ], \"suites\": [ { \"id\":\"a86b2934-7aa3-4838-b389-93c8aea2af05\",  \"name\":\"Default Suite\",  \"persistSession\":false, \"parallel\":false, \"timeout\":600,  \"tests\":  [  \"b4522766-e382-40c2-ab01-452cf62e1cec\" ] } ], \"urls\": [ \"https://console.us-ashburn-1.oraclecloud.com/\"  ], \"plugins\":[\"xxx\"] }"
}

variable "script_content_file_name" {
  default = "contentFileName"
}

variable "script_content_type" {
  default = "SIDE"
}

variable "script_defined_tags_value" {
  default = "value"
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
  default = "paramName"
}

variable "script_parameters_param_value" {
  default = "paramValue"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apm_synthetics_script" "test_script" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  content       = var.script_content
  content_type  = var.script_content_type
  display_name  = var.script_display_name

  #Optional
  content_file_name = var.script_content_file_name
  defined_tags      = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.script_defined_tags_value)
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

