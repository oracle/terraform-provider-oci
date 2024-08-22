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


resource "oci_data_safe_unset_user_assessment_baseline_management" "test_unset_user_assessment_baseline_management" {
    #Required
    user_assessment_id = oci_data_safe_set_user_assessment_baseline_management.test_set_user_assessment_baseline_management.user_assessment_id
    compartment_id = var.compartment_ocid
}