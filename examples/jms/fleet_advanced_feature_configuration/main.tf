// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "fleet_id" {
  default = "example-fleet-id"
}

variable "analytic_bucket_name" {
  default= "example-analytic-bucket-name"
}

variable "analytic_namespace" {
  default= "example-analytic-namespace"
}

variable "log_group_ocid" {
  default = "example-log-group-id"
}

variable "inventory_log_ocid" {
  default = "example-inventory-log-id"
}
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_jms_fleet_advanced_feature_configuration" "example_fleet_advanced_feature_configuration" {
  #Required
  fleet_id = var.fleet_id

  #Optional
  advanced_usage_tracking {
    is_enabled = "false"
  }
  analytic_bucket_name = var.analytic_bucket_name
  analytic_namespace = var.analytic_namespace
  crypto_event_analysis {
    is_enabled = "false"
    summarized_events_log {
      log_group_id = var.log_group_ocid
      log_id = var.inventory_log_ocid
    }
  }
  java_migration_analysis {
    is_enabled = "false"
  }
  jfr_recording {
    is_enabled = "false"
  }
  lcm {
    is_enabled = "false"
    post_installation_actions {
      add_logging_handler = "false"
      disabled_tls_versions = ["TLS_1_0"]
      global_logging_level = "ALL"
      minimum_key_size_settings {
        certpath {
          key_size = "2048"
          name = "RSA"
        }
        jar {
          key_size = "2048"
          name = "RSA"
        }
        tls {
          key_size = "2048"
          name = "RSA"
        }
      }
      proxies {
        ftp_proxy_host = "ftpProxyHost"
        ftp_proxy_port = "10"
        http_proxy_host = "httpProxyHost"
        http_proxy_port = "10"
        https_proxy_host = "httpsProxyHost"
        https_proxy_port = "10"
        socks_proxy_host = "socksProxyHost"
        socks_proxy_port = "10"
        use_system_proxies = "false"
      }
      should_replace_certificates_operating_system = "false"
    }
  }
  performance_tuning_analysis {
    is_enabled = "false"
  }
}

data "oci_jms_fleet_advanced_feature_configuration" "test_fleet_advanced_feature_configuration" {
	#Required
	fleet_id = var.fleet_id
}