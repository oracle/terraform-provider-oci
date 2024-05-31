// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

// If using the objectStorage target
//variable "object_storage_bucket_name" {}

// If using the log analytics target
variable "log_analytics_log_group_id" {}

// If using the notification target
//variable "notification_topic_id" {}

// streaming cursor kind

variable "streaming_cursor_kind" {
  default = "LATEST"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "image" {
  default = ""
}

variable "queue_id" {
  default = ""
}

variable "function_id" {
  default = ""
}

variable "service_connector_target_log_source_identifier" {
  default = "logSourceIdentifier"
}

variable defined_tag_namespace_name {
  default = ""
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  #Required
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

  is_retired = false
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
}

resource "oci_functions_application" "test_application" {
  compartment_id = var.compartment_ocid
  display_name   = "displayName"
  subnet_ids     = [oci_core_subnet.test_subnet.id]
}

resource "oci_functions_function" "test_function" {
  application_id = oci_functions_application.test_application.id
  display_name   = "displayName"
  image          = var.image
  memory_in_mbs  = "128"
}

resource "oci_streaming_stream" "test_stream" {
  compartment_id = var.compartment_ocid
  name           = "mynewstream"
  partitions     = "1"
}

resource "oci_logging_log_group" "test_log_group" {
  compartment_id = var.compartment_ocid
  display_name   = "tempDisplayName"
}

resource "oci_logging_log" "test_log" {
  configuration {
    source {
      category    = "write"
      resource    = oci_objectstorage_bucket.test_bucket.name
      service     = "objectstorage"
      source_type = "OCISERVICE"
    }
  }

  display_name = "displayName"
  log_group_id = oci_logging_log_group.test_log_group.id
  log_type     = "SERVICE"
}

resource "oci_objectstorage_bucket" "test_bucket" {
  compartment_id = var.compartment_ocid
  name           = "newName"
  namespace      = data.oci_objectstorage_namespace.test_namespace.namespace
}

data "oci_objectstorage_namespace" "test_namespace" {
  compartment_id = var.compartment_ocid
}

resource "oci_sch_service_connector" "test_service_connector" {
  compartment_id = var.compartment_ocid
  defined_tags   = { "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue" }
  description    = "description2"
  display_name   = "displayName2"

  freeform_tags = {
    "Department" = "Accounting"
  }

  /* source {
     kind = "logging"

     log_sources {
       compartment_id = var.compartment_ocid
       log_group_id   = oci_logging_log_group.test_log_group.id
       log_id         = oci_logging_log.test_log.id
     }
   }
   */

  // If using streaming source
  source {
    kind = "Streaming"

    // Optional
    cursor {

      // Optional
      kind = var.streaming_cursor_kind
    }

    stream_id = oci_streaming_stream.test_stream.id
  }

  /*target {
    kind      = "streaming"
    stream_id = oci_streaming_stream.test_stream.id
  }*/

  // If using the objectStorage target
  /*target {
    kind                        = "objectStorage"
    bucket                      = var.object_storage_bucket_name

    // Optional
    batch_rollover_size_in_mbs" = "10"
    // Optional
    batch_rollover_time_in_ms"  = "80000"
  }*/

  // If using the log analytics target
  target {
    kind            = "loggingAnalytics"
    log_group_id    = var.log_analytics_log_group_id
    log_source_identifier = var.service_connector_target_log_source_identifier
  }

  // If using the notification target
  /*target {
    kind            		        = "notifications"
    topic_id                    = var.notification_topic_id
    enable_formatted_messaging	= "true"
  }*/

  // If using the monitoring target
  /*target {
    kind                        = "monitoring"
    compartment_id              = var.compartment_ocid
    metric                      = var.metric
    metric_namespace            = var.metric_namespace
    // Optional
    dimensions                  = var.dimensions
  }*/

  /*tasks {
    condition = "logContent='20'"
    kind      = "logRule"
  }*/

  // If using function task
  /*
  tasks {
    function_id = oci_functions_function.test_function.id
    kind        = "functions"
  }
  */

  state = "ACTIVE"
}

data "oci_sch_service_connector" "test_service_connector" {
  service_connector_id = oci_sch_service_connector.test_service_connector.id
}

output "oci_sch_service_connector_id" {
  value = [data.oci_sch_service_connector.test_service_connector.id]
}

resource "oci_sch_service_connector" "test_connector_plugins" {
  compartment_id = var.compartment_ocid
  display_name   = "My_Service_Connector"
  source {
    kind        = "plugin"
    plugin_name = "QueueSource"
    config_map = "{\"queueId\": \"${var.queue_id}\"}"
  }
  // If using the functions target
  target {
    kind        = "functions"
    function_id = oci_functions_function.test_function.id

    // Optional
    batch_size_in_kbs = "5000"
    // Optional
    batch_time_in_sec = "62"
  }
}
