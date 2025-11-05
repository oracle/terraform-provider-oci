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

variable "compartment_id" { }

variable "stack_display_name" {
  default = "testtfstack"
}

variable "stack_freeform_tags" {
  default = { "Division" = "DEV" }
}

variable "stack_adb_is_public" {
  default = false
}

variable "password_secret_id" { }
variable "dataflow_file_uri" { }
variable "stack_subnet_id" { }

resource "oci_dif_stack" "test_stack" {
  #Required
  compartment_id  = var.compartment_id
  display_name    = var.stack_display_name

  services = [
    "OBJECTSTORAGE", "ADB", "DATAFLOW"
  ]

  stack_templates = [
    "DATALAKE", "DATATRANSFORMATION"
  ]

  objectstorage {
    instance_id       = "testLogBucket1"
    object_versioning = "ENABLED"
    storage_tier      = "STANDARD"
  }

  adb {
    instance_id                 = "testadb1"
    db_workload                 = "DW"
    ecpu                        = 2
    data_storage_size_in_tbs    = 1
    admin_password_id          = var.password_secret_id
    is_public                  = var.stack_adb_is_public
    subnet_id                  = var.stack_subnet_id
    db_version                 = "19c"
  }

  dataflow {
    instance_id = "testApp"
    log_bucket_instance_id = "testLogBucket1"
    num_executors = "3"
    spark_version = "3.5.0"
    connections {
      connection_details {
        dif_dependencies {
          service_instance_id = "testadb1"
          service_type = "ADB"
        }
        domain_names = ["custpvtsubnet.oraclevcn.com"]
      }
      subnet_id = var.stack_subnet_id
    }
    driver_shape = "VM.Standard.E5.Flex"
    driver_shape_config {
      memory_in_gbs = "16"
      ocpus = "1"
    }
    executor_shape = "VM.Standard.E5.Flex"
    executor_shape_config {
      memory_in_gbs = "16"
      ocpus = "1"
    }
    execute = var.dataflow_file_uri
  }
  objectstorage {
    instance_id       = "testBucket2"
    object_versioning = "DISABLED"
    storage_tier      = "STANDARD"
  }

  subnet_id = var.stack_subnet_id

  deploy_artifacts_trigger = 1
  add_service_trigger = 1

  //ignore so that this does not cause perpetual diff
  lifecycle {
    ignore_changes = [
      # Ignore service-generated metadata and tags
      defined_tags,
      freeform_tags,
      system_tags,
      service_details,
      time_created,
      time_updated,
    ]
  }
}