// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "compartment_ocid" {}

variable "instance_image_ocid" {
  type = "map"

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"

    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

variable "db_edition" {
  default = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
}

variable "db_admin_password" {
  default = "BEstrO0ng_#12"
}

variable "db_version" {
  default = "12.1.0.2"
}

variable "db_system_shape" {
  default = "Exadata.Quarter1.84"
}

variable "n_character_set" {
  default = "AL16UTF16"
}

variable "character_set" {
  default = "AL32UTF8"
}

variable "db_workload" {
  default = "OLTP"
}

/* Nosql */

variable "table_ddl_statement" {
  default = "CREATE TABLE IF NOT EXISTS test_table(id INTEGER, name STRING, age STRING, info JSON, PRIMARY KEY(SHARD(id)))"
}

variable "index_keys_column_name" {
  default = "name"
}

/* Monitoring */

variable "alarm_body" {
  default = "High CPU utilization reached"
}

variable "alarm_compartment_id_in_subtree" {
  default = false
}

variable "alarm_defined_tags_value" {
  default = "value"
}

variable "alarm_destinations" {
  default = []
}

variable "alarm_display_name" {
  default = "High CPU Utilization"
}

variable "alarm_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "alarm_is_enabled" {
  default = false
}

variable "alarm_metric_compartment_id_in_subtree" {
  default = false
}

variable "alarm_namespace" {
  default = "oci_computeagent"
}

variable "alarm_pending_duration" {
  default = "PT5M"
}

variable "alarm_query" {
  default = "CpuUtilization[10m].percentile(0.9) < 85"
}

variable "alarm_repeat_notification_duration" {
  default = "PT2H"
}

variable "alarm_resolution" {
  default = "1m"
}

variable "alarm_resource_group" {
  default = "resourceGroup"
}

variable "alarm_severity" {
  default = "WARNING"
}

variable "alarm_state" {
  default = "ACTIVE"
}

variable "alarm_suppression_description" {
  default = "System Maintenance"
}

variable "alarm_suppression_time_suppress_from" {
  default = "2029-02-01T18:00:00.000Z"
}

variable "alarm_suppression_time_suppress_until" {
  default = "2029-02-01T19:00:00.000Z"
}

/* Waas */

variable "certificate_display_name" {
  default = "tf_example_waas_certificate_rd"
}

variable "waas_policy_display_name" {
  default = "tf_example_waas_policy_rd"
}