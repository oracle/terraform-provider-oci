// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {

}

variable "user_ocid" {
}

variable "fingerprint" {

}

variable "private_key_path" {
}

variable "region" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_limits_quota" "test_quota" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "Quotas for VCN"
  name           = "TestQuotas"
  statements     = ["Set vcn quotas to 0 in tenancy"]

  depends_on = [oci_identity_tag.tag1]

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  }

  freeform_tags = {
    "Department" = "Finance"
  }
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_limits_quotas" "test_quotas" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  name  = "TestQuotas"
  state = "ACTIVE"
}

data "oci_limits_services" "test_services" {
  #Required
  compartment_id = var.tenancy_ocid

  filter {
    name   = "name"
    values = ["compute"]
  }
}

data "oci_limits_limit_definitions" "test_limit_definitions" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  name         = var.limit_definition_name
  service_name = data.oci_limits_services.test_services.services[0].name
}

data "oci_limits_resource_availability" "test_resource_availability" {
  #Required
  compartment_id = var.tenancy_ocid
  limit_name     = var.limit_definition_name
  service_name   = data.oci_limits_services.test_services.services[0].name

  #Optional
  #specify this parameter depending upon the limit and service
  availability_domain = data.oci_identity_availability_domain.ad.name
}

data "oci_limits_limit_values" "test_limit_values" {
  #Required
  compartment_id = var.tenancy_ocid
  service_name   = data.oci_limits_services.test_services.services[0].name

  #Optional
  availability_domain = data.oci_identity_availability_domain.ad.name
  name                = var.limit_definition_name
  scope_type          = "AD"
}


## quota lock example

resource "oci_limits_quota" "test_quota_lock" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "Quotas for VCN"
  name           = "TestQuotaLocks"
  statements     = ["Set vcn quotas to 0 in tenancy"]
  depends_on = [oci_identity_tag.tag1] 
  #Optional
  locks { 
        type = "FULL"
        #Optional
        message  = "lock testing" 
        #Optional
        related_resource_id = "some resource id"
  } 
}

data "oci_limits_quotas" "test_quotas_lock" {
  #Required
  compartment_id = var.tenancy_ocid 
  depends_on = [oci_limits_quota.test_quota_lock]
  #Optional  
  name  = "TestQuotaLocks"
  state = "ACTIVE"
}

