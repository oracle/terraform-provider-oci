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

variable "compartment_ocid" {
}

variable "config" {
  type = map(string) 
  default = {
    "MEDIA_WORKFLOW_ID" = "someworkflowid"
  }
}

##### Application Performance Monitoring ######
# To use APM tracing with functions as a service, you need to have created an APM domain.
#
# You can learn how to create an APM domain to begins tracing your functions here:
# https://docs.oracle.com/en-us/iaas/application-performance-monitoring/doc/create-apm-domain.html
#

variable "application_trace_config" {
  type = object({
    domain_id = string
    is_enabled = bool
  })
  default = {
    domain_id = ""
    is_enabled = false
  }
}

variable "syslog_url" {
  default = ""
}

variable "application_image_policy_config_is_policy_enabled" {
  default = false
}

##### Docker image ######
# To use functions as a service, you need docker images of your functions in OCI registry which can be accessed by you
#
# You can learn about pushing images to OCI registry here:
# https://docs.cloud.oracle.com/iaas/Content/Registry/Tasks/registrypushingimagesusingthedockercli.htm
#

variable "application_state" {
  default = "ACTIVE"
}

variable "function_image" {
}

variable "function_image_digest" {
}

variable "dry_run" {
  default = "false"
}

variable "function_trace_config" {
  type = object({
    is_enabled = bool
  })
  default = {
    is_enabled = false
  }
}

variable "function_memory_in_mbs" {
  default = 128
}

variable "function_timeout_in_seconds" {
  default = 30
}

variable "kms_key_ocid" {
}

variable "pbf_listing_name" {
}

variable "pbf_listing_id" {
}

variable "pbf_listing_version_id" {
}

variable "pbf_trigger_name" {
}

variable "application_shape" {
  default = "GENERIC_X86"
}