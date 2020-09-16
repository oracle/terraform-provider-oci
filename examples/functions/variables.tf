// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
  default = {
    "MY_FUNCTION_CONFIG" = "ConfVal"
  }
}

##### Docker image ######
# To use functions as a service, you need docker images of your functions in OCI registry which can be accessed by you
#
# You can learn about pushing images to OCI registry here:
# https://docs.cloud.oracle.com/iaas/Content/Registry/Tasks/registrypushingimagesusingthedockercli.htm
#

variable "application_state" {
  default = "AVAILABLE"
}

variable "function_image" {
}

variable "function_image_digest" {
}

variable "function_memory_in_mbs" {
  default = 128
}

variable "function_timeout_in_seconds" {
  default = 30
}

variable "invoke_function_body" {
}

variable "invoke_function_body_source_path" {
}

