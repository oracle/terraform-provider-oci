// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "apm_domain_id" {}

variable "scheduled_query_defined_tags_value" {
  default = "value"
}

variable "scheduled_query_display_name" {
  default = "displayName"
}

variable "scheduled_query_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "scheduled_query_opc_dry_run" {
  default = "opcDryRun"
}

variable "scheduled_query_scheduled_query_description" {
  default = "scheduledQueryDescription"
}

variable "scheduled_query_scheduled_query_maximum_runtime_in_seconds" {
  default = 10
}

variable "scheduled_query_scheduled_query_processing_configuration_custom_metric_compartment" {
  default = "compartment"
}

variable "scheduled_query_scheduled_query_processing_configuration_custom_metric_description" {
  default = "description"
}

variable "scheduled_query_scheduled_query_processing_configuration_custom_metric_is_anomaly_detection_enabled" {
  default = false
}

variable "scheduled_query_scheduled_query_processing_configuration_custom_metric_is_metric_published" {
  default = false
}

variable "scheduled_query_scheduled_query_processing_configuration_custom_metric_name" {
  default = "name"
}

variable "scheduled_query_scheduled_query_processing_configuration_custom_metric_namespace" {
  default = "namespace"
}

variable "scheduled_query_scheduled_query_processing_configuration_custom_metric_resource_group" {
  default = "resourceGroup"
}

variable "scheduled_query_scheduled_query_processing_configuration_custom_metric_unit" {
  default = "unit"
}

variable "scheduled_query_scheduled_query_processing_configuration_object_storage_bucket" {
  default = "bucket"
}

variable "scheduled_query_scheduled_query_processing_configuration_object_storage_name_space" {
  default = "nameSpace"
}

variable "scheduled_query_scheduled_query_processing_configuration_object_storage_object_name_prefix" {
  default = "objectNamePrefix"
}

variable "scheduled_query_scheduled_query_processing_sub_type" {
  default = "OBJECT_STORAGE"
}

variable "scheduled_query_scheduled_query_processing_type" {
  default = "EXPORT"
}

variable "scheduled_query_scheduled_query_retention_criteria" {
  default = "UPDATE"
}

variable "scheduled_query_scheduled_query_retention_period_in_ms" {
  default = 10
}

variable "scheduled_query_scheduled_query_schedule" {
  default = "SCHEDULE STARTING AFTER 2025-06-20T21:20:00Z EVERY 60 MINUTES"
}

variable "scheduled_query_scheduled_query_text" {
  default = "SHOW SPANS time_bucket_start(1,apmdbInsertTime2) AS metricStartTime,1 MINUTE AS metricDuration,serviceName,operationName,count(*) AS metricValue WHERE apmdbinserttime2>=TimeTruncate(now(),'minute') - 5 MINUTES GROUP BY time_bucket_start(1,apmdbInsertTime2),serviceName,operationName FIRST 10000 ROWS BETWEEN now() - 2 HOURS AND now()"
}

variable "scheduled_query_scheduled_query_name" {
  default = "scheduledQueryName"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apm_traces_scheduled_query" "test_scheduled_query" {
  #Required
  apm_domain_id = var.apm_domain_id

  #Optional
  #defined_tags                               = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.scheduled_query_defined_tags_value)
  freeform_tags                              = var.scheduled_query_freeform_tags
  opc_dry_run                                = var.scheduled_query_opc_dry_run
  scheduled_query_description                = var.scheduled_query_scheduled_query_description
  scheduled_query_maximum_runtime_in_seconds = var.scheduled_query_scheduled_query_maximum_runtime_in_seconds
  scheduled_query_name                       = var.scheduled_query_scheduled_query_name
  scheduled_query_processing_configuration {

    #Optional
    custom_metric {
      #Required
      name = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_name

      #Optional
      compartment                  = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_compartment
      description                  = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_description
      is_anomaly_detection_enabled = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_is_anomaly_detection_enabled
      is_metric_published          = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_is_metric_published
      namespace                    = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_namespace
      resource_group               = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_resource_group
      unit                         = var.scheduled_query_scheduled_query_processing_configuration_custom_metric_unit
    }
    object_storage {

      #Optional
      bucket             = var.scheduled_query_scheduled_query_processing_configuration_object_storage_bucket
      name_space         = var.scheduled_query_scheduled_query_processing_configuration_object_storage_name_space
      object_name_prefix = var.scheduled_query_scheduled_query_processing_configuration_object_storage_object_name_prefix
    }
    streaming {

      #Optional
      #stream_id = oci_streaming_stream.test_stream.id
    }
  }
  scheduled_query_processing_sub_type    = var.scheduled_query_scheduled_query_processing_sub_type
  scheduled_query_processing_type        = var.scheduled_query_scheduled_query_processing_type
  scheduled_query_retention_criteria     = var.scheduled_query_scheduled_query_retention_criteria
  scheduled_query_retention_period_in_ms = var.scheduled_query_scheduled_query_retention_period_in_ms
  scheduled_query_schedule               = var.scheduled_query_scheduled_query_schedule
  scheduled_query_text                   = var.scheduled_query_scheduled_query_text
}

data "oci_apm_traces_scheduled_queries" "test_scheduled_queries" {
  #Required
  apm_domain_id = var.apm_domain_id

  #Optional
  display_name = var.scheduled_query_display_name
}
