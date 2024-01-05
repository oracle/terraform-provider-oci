// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to manage Log Analytics Scheduled Tasks
 */


variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}
variable "saved_search_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

# lookup namespace corresponding to this tenancy
data "oci_objectstorage_namespace" "test_namespace" {
  compartment_id = var.tenancy_ocid
}

# create a purge scheduled task
resource "oci_log_analytics_namespace_scheduled_task" "test_namespace_scheduled_task" {
  compartment_id = var.compartment_ocid
  namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  kind = "STANDARD"
  task_type = "PURGE"
  display_name = "tfPurgeTask1"

  action {
    compartment_id_in_subtree = "false"
    data_type = "LOG"
    purge_compartment_id = var.compartment_ocid
    purge_duration = "-P1D"
    query_string = "fake_query"
    type = "PURGE"
  }

  // only one schedule allowed for purge tasks
  schedules {
    schedule {
        type = "FIXED_FREQUENCY"
        misfire_policy = "RETRY_ONCE"
        recurring_interval = "P1D"
        repeat_count = "4"
    }
  }

  freeform_tags = {
    "Department" = "Accounting"
  }
}

resource "oci_log_analytics_namespace_scheduled_task" "test_namespace_scheduled_task_metric" {
    compartment_id = var.compartment_ocid
    namespace = data.oci_objectstorage_namespace.test_namespace.namespace
    kind = "STANDARD"
    task_type = "SAVED_SEARCH"
    display_name = "tfPurgeTask2"

    action {
        metric_extraction {
          compartment_id = var.compartment_ocid
          namespace = "test_scheduled_task_metrics"
          metric_name = "count"
        }
        type = "STREAM"
        saved_search_id = var.saved_search_id
    }

    schedules {
      schedule {
         type = "FIXED_FREQUENCY"
         misfire_policy = "RETRY_ONCE"
         recurring_interval = "PT5M"
         repeat_count = "10"
       }
    }
} 

# look up using the scheduled tasks data source
data "oci_log_analytics_namespace_scheduled_tasks" "test_namespace_scheduled_tasks" {
  compartment_id = var.compartment_ocid
  display_name = "tfPurgeTask3"
  filter {
    name = "id"
    values = ["${oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task.id}"]
  }
  namespace = "${data.oci_objectstorage_namespace.test_namespace.namespace}"
  task_type = "PURGE"
}

# look up using the scheduled task singular data source to look a scheduled task
# with its composite ID which includes the namespace and scheduled_task_id
data "oci_log_analytics_namespace_scheduled_task" "test_namespace_scheduled_task" {
  namespace = "${data.oci_objectstorage_namespace.test_namespace.namespace}"
  scheduled_task_id = "${oci_log_analytics_namespace_scheduled_task.test_namespace_scheduled_task.scheduled_task_id}"
}

