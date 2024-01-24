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

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

resource "oci_data_labeling_service_dataset" "test_dataset" {
  annotation_format = "BOUNDING_BOX"
  compartment_id    = var.compartment_ocid

  dataset_format_details {
    format_type = "IMAGE"
  }

  dataset_source_details {
    bucket      = oci_objectstorage_bucket.test_bucket.name
    namespace   = data.oci_objectstorage_namespace.ns.namespace
    source_type = "OBJECT_STORAGE"
    prefix      = ""
  }

  label_set {
    items {
      name = "label1"
    }
    items {
      name = "label2"
    }
  }

  description  = "Example Dataset"
  display_name = "Example Dataset"
}