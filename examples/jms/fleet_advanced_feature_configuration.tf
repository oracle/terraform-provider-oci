// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_jms_fleet" "example_fleet_with_advanced_features" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "Example Fleet with Advanced Features"
  inventory_log {
    log_group_id = var.fleet_log_group_ocid
    log_id       = var.fleet_inventory_log_ocid
  }

  #Optional
  description                  = "Example Fleet with Advanced Features created by Terraform"
  freeform_tags                = var.fleet_freeform_tags
  operation_log {
    log_group_id = var.fleet_log_group_ocid
    log_id       = var.fleet_operation_log_ocid
  }

  # Create the Tag namespace in OCI before enabling
  # See user guide: https://docs.oracle.com/en-us/iaas/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm
  # defined_tags  = var.fleet_defined_tags
}

resource "oci_jms_fleet_advanced_feature_configuration" "example_fleet_advanced_feature_configuration" {
  #Required
  fleet_id = oci_jms_fleet.example_fleet_with_advanced_features.id

  #Optional
  advanced_usage_tracking {
    is_enabled = "true"
  }
  analytic_bucket_name = var.analytic_bucket_name
  analytic_namespace = var.analytic_bucket_namespace
  crypto_event_analysis {
    is_enabled = "true"
    summarized_events_log {
      log_group_id = var.fleet_log_group_ocid
      log_id = var.crypto_event_log_ocid
    }
  }
  java_migration_analysis {
    is_enabled = "true"
  }
  jfr_recording {
    is_enabled = "true"
  }
  lcm {
    is_enabled = "true"
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
        ftp_proxy_host = "example-ftp-proxy-host"
        ftp_proxy_port = "10"
        http_proxy_host = "example-http-proxy-host"
        http_proxy_port = "10"
        https_proxy_host = "example-https-proxy-host"
        https_proxy_port = "10"
        socks_proxy_host = "example-socks-proxy-host"
        socks_proxy_port = "10"
        use_system_proxies = "false"
      }
      should_replace_certificates_operating_system = "false"
    }
  }
  performance_tuning_analysis {
    is_enabled = "true"
  }
}

data "oci_jms_fleet_advanced_feature_configuration" "test_fleet_advanced_feature_configuration" {
	#Required
	fleet_id = oci_jms_fleet.example_fleet_with_advanced_features.id
}