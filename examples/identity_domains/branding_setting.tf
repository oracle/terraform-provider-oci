// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "branding_setting_authorization" {
  default = "authorization"
}


data "oci_identity_domains_branding_settings" "test_branding_settings" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url
  
  #Optional
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.branding_setting_authorization
  #use the latest version if not provided
  # resource_type_schema_version = var.branding_setting_resource_type_schema_version
}

