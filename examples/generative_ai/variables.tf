// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "hosting_cluster_display_name" {
  default = "hosting_cluster"
}

variable "hosting_cluster_description" {
  default = "this is a hosting cluster"
}

variable "hosting_cluster_shape" {
  default = "SMALL_COHERE"
}

variable "hosting_cluster_unit_count" {
  default = "1"
}

variable "fine_tuning_cluster_display_name" {
  default = "fine_tuning_cluster"
}

variable "fine_tuning_cluster_description" {
  default = "this is a fine tuning cluster"
}

variable "fine_tuning_cluster_shape" {
  default = "SMALL_COHERE"
}

variable "fine_tuning_cluster_unit_count" {
  default = "2"
}

variable "test_endpoint_display_name" {
  default = "test_endpoint"
}

variable "test_endpoint_description" {
  default = "test endpoint"
}

variable "test_model_display_name" {
  default = "test_model"
}

variable "test_model_description" {
  default = "test model"
}

variable "test_model_vendor" {
  default = "test_vendor"
}

variable "test_model_version" {
  default = "1.1"
}

variable "test_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}