// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "opensearch_cluster_backup_policy_frequency_in_hours" {
  default = 10
}

variable "opensearch_cluster_backup_policy_is_enabled" {
  default = false
}

variable "opensearch_cluster_backup_policy_retention_in_days" {
  default = 10
}

variable "opensearch_cluster_certificate_config_cluster_certificate_mode" {
  default = "OCI_CERTIFICATES_SERVICE"
}

variable "opensearch_cluster_certificate_config_dashboard_certificate_mode" {
  default = "OCI_CERTIFICATES_SERVICE"
}

variable "opensearch_cluster_data_node_count" {
  default = 10
}

variable "opensearch_cluster_data_node_host_bare_metal_shape" {
  default = "dataNodeHostBareMetalShape"
}

variable "opensearch_cluster_data_node_host_memory_gb" {
  default = 10
}

variable "opensearch_cluster_data_node_host_ocpu_count" {
  default = 10
}

variable "opensearch_cluster_data_node_host_shape" {
  default = "dataNodeHostShape"
}

variable "opensearch_cluster_data_node_host_type" {
  default = "FLEX"
}

variable "opensearch_cluster_data_node_storage_gb" {
  default = 10
}

variable "opensearch_cluster_defined_tags_value" {
  default = "value"
}

variable "opensearch_cluster_display_name" {
  default = "displayName"
}

variable "opensearch_cluster_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "opensearch_cluster_id" {
  default = "id"
}

variable "opensearch_cluster_inbound_cluster_ids" {
  default = []
}

variable "opensearch_cluster_load_balancer_config_load_balancer_max_bandwidth_in_mbps" {
  default = 10
}

variable "opensearch_cluster_load_balancer_config_load_balancer_min_bandwidth_in_mbps" {
  default = 10
}

variable "opensearch_cluster_load_balancer_config_load_balancer_service_type" {
  default = "LOAD_BALANCER"
}

variable "opensearch_cluster_maintenance_details_notification_email_ids" {
  default = []
}

variable "opensearch_cluster_master_node_count" {
  default = 10
}

variable "opensearch_cluster_master_node_host_bare_metal_shape" {
  default = "masterNodeHostBareMetalShape"
}

variable "opensearch_cluster_master_node_host_memory_gb" {
  default = 10
}

variable "opensearch_cluster_master_node_host_ocpu_count" {
  default = 10
}

variable "opensearch_cluster_master_node_host_shape" {
  default = "masterNodeHostShape"
}

variable "opensearch_cluster_master_node_host_type" {
  default = "FLEX"
}

variable "opensearch_cluster_ml_node_count" {
  default = 10
}

variable "opensearch_cluster_ml_node_host_memory_gb" {
  default = 10
}

variable "opensearch_cluster_ml_node_host_ocpu_count" {
  default = 10
}

variable "opensearch_cluster_ml_node_host_shape" {
  default = "mlNodeHostShape"
}

variable "opensearch_cluster_ml_node_host_type" {
  default = "FLEX"
}

variable "opensearch_cluster_ml_node_storage_gb" {
  default = 10
}

variable "opensearch_cluster_opendashboard_node_count" {
  default = 10
}

variable "opensearch_cluster_opendashboard_node_host_memory_gb" {
  default = 10
}

variable "opensearch_cluster_opendashboard_node_host_ocpu_count" {
  default = 10
}

variable "opensearch_cluster_opendashboard_node_host_shape" {
  default = "opendashboardNodeHostShape"
}

variable "opensearch_cluster_outbound_cluster_config_is_enabled" {
  default = false
}

variable "opensearch_cluster_outbound_cluster_config_outbound_clusters_display_name" {
  default = "displayName"
}

variable "opensearch_cluster_outbound_cluster_config_outbound_clusters_is_skip_unavailable" {
  default = false
}

variable "opensearch_cluster_outbound_cluster_config_outbound_clusters_mode" {
  default = "SEARCH_ONLY"
}

variable "opensearch_cluster_outbound_cluster_config_outbound_clusters_ping_schedule" {
  default = "pingSchedule"
}

variable "opensearch_cluster_reverse_connection_endpoint_customer_ips" {
  default = []
}

variable "opensearch_cluster_search_node_count" {
  default = 10
}

variable "opensearch_cluster_search_node_host_memory_gb" {
  default = 10
}

variable "opensearch_cluster_search_node_host_ocpu_count" {
  default = 10
}

variable "opensearch_cluster_search_node_host_shape" {
  default = "searchNodeHostShape"
}

variable "opensearch_cluster_search_node_host_type" {
  default = "FLEX"
}

variable "opensearch_cluster_search_node_storage_gb" {
  default = 10
}

variable "opensearch_cluster_security_attributes" {
  default = { "Oracle-ZPR" = { "MaxEgressCount" = { "value" = "42", "mode" = "enforce" } } }
}

variable "opensearch_cluster_security_master_user_password_hash" {
  default = "securityMasterUserPasswordHash"
}

variable "opensearch_cluster_security_mode" {
  default = "DISABLED"
}

variable "opensearch_cluster_security_saml_config_admin_backend_role" {
  default = "adminBackendRole"
}

variable "opensearch_cluster_security_saml_config_idp_metadata_content" {
  default = "idpMetadataContent"
}

variable "opensearch_cluster_security_saml_config_is_enabled" {
  default = false
}

variable "opensearch_cluster_security_saml_config_opendashboard_url" {
  default = "opendashboardUrl"
}

variable "opensearch_cluster_security_saml_config_roles_key" {
  default = "rolesKey"
}

variable "opensearch_cluster_security_saml_config_subject_key" {
  default = "subjectKey"
}

variable "opensearch_cluster_software_version" {
  default = "softwareVersion"
}

variable "opensearch_cluster_state" {
  default = "AVAILABLE"
}

variable "opensearch_cluster_system_tags" {
  default = "value"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_opensearch_opensearch_cluster" "test_opensearch_cluster" {
  #Required
  compartment_id                     = var.compartment_id
  data_node_count                    = var.opensearch_cluster_data_node_count
  data_node_host_memory_gb           = var.opensearch_cluster_data_node_host_memory_gb
  data_node_host_ocpu_count          = var.opensearch_cluster_data_node_host_ocpu_count
  data_node_host_type                = var.opensearch_cluster_data_node_host_type
  data_node_storage_gb               = var.opensearch_cluster_data_node_storage_gb
  display_name                       = var.opensearch_cluster_display_name
  master_node_count                  = var.opensearch_cluster_master_node_count
  master_node_host_memory_gb         = var.opensearch_cluster_master_node_host_memory_gb
  master_node_host_ocpu_count        = var.opensearch_cluster_master_node_host_ocpu_count
  master_node_host_type              = var.opensearch_cluster_master_node_host_type
  opendashboard_node_count           = var.opensearch_cluster_opendashboard_node_count
  opendashboard_node_host_memory_gb  = var.opensearch_cluster_opendashboard_node_host_memory_gb
  opendashboard_node_host_ocpu_count = var.opensearch_cluster_opendashboard_node_host_ocpu_count
  software_version                   = var.opensearch_cluster_software_version
  subnet_compartment_id              = oci_identity_compartment.test_compartment.id
  subnet_id                          = oci_core_subnet.test_subnet.id
  vcn_compartment_id                 = oci_identity_compartment.test_compartment.id
  vcn_id                             = oci_core_vcn.test_vcn.id

  #Optional
  backup_policy {

    #Optional
    frequency_in_hours = var.opensearch_cluster_backup_policy_frequency_in_hours
    is_enabled         = var.opensearch_cluster_backup_policy_is_enabled
    retention_in_days  = var.opensearch_cluster_backup_policy_retention_in_days
  }
  certificate_config {

    #Optional
    cluster_certificate_mode             = var.opensearch_cluster_certificate_config_cluster_certificate_mode
    dashboard_certificate_mode           = var.opensearch_cluster_certificate_config_dashboard_certificate_mode
    open_search_api_certificate_id       = oci_apigateway_certificate.test_certificate.id
    open_search_dashboard_certificate_id = oci_apigateway_certificate.test_certificate.id
  }
  data_node_host_bare_metal_shape = var.opensearch_cluster_data_node_host_bare_metal_shape
  data_node_host_shape            = var.opensearch_cluster_data_node_host_shape
  defined_tags                    = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.opensearch_cluster_defined_tags_value)
  freeform_tags                   = var.opensearch_cluster_freeform_tags
  inbound_cluster_ids             = var.opensearch_cluster_inbound_cluster_ids
  load_balancer_config {
    #Required
    load_balancer_service_type = var.opensearch_cluster_load_balancer_config_load_balancer_service_type

    #Optional
    load_balancer_max_bandwidth_in_mbps = var.opensearch_cluster_load_balancer_config_load_balancer_max_bandwidth_in_mbps
    load_balancer_min_bandwidth_in_mbps = var.opensearch_cluster_load_balancer_config_load_balancer_min_bandwidth_in_mbps
  }
  maintenance_details {

    #Optional
    notification_email_ids = var.opensearch_cluster_maintenance_details_notification_email_ids
  }
  master_node_host_bare_metal_shape = var.opensearch_cluster_master_node_host_bare_metal_shape
  master_node_host_shape            = var.opensearch_cluster_master_node_host_shape
  ml_node_count                     = var.opensearch_cluster_ml_node_count
  ml_node_host_memory_gb            = var.opensearch_cluster_ml_node_host_memory_gb
  ml_node_host_ocpu_count           = var.opensearch_cluster_ml_node_host_ocpu_count
  ml_node_host_shape                = var.opensearch_cluster_ml_node_host_shape
  ml_node_host_type                 = var.opensearch_cluster_ml_node_host_type
  ml_node_storage_gb                = var.opensearch_cluster_ml_node_storage_gb
  nsg_id                            = oci_opensearch_nsg.test_nsg.id
  opendashboard_node_host_shape     = var.opensearch_cluster_opendashboard_node_host_shape
  outbound_cluster_config {
    #Required
    is_enabled = var.opensearch_cluster_outbound_cluster_config_is_enabled
    outbound_clusters {
      #Required
      display_name    = var.opensearch_cluster_outbound_cluster_config_outbound_clusters_display_name
      seed_cluster_id = oci_opensearch_opensearch_cluster.test_opensearch_cluster.id

      #Optional
      is_skip_unavailable = var.opensearch_cluster_outbound_cluster_config_outbound_clusters_is_skip_unavailable
      mode                = var.opensearch_cluster_outbound_cluster_config_outbound_clusters_mode
      ping_schedule       = var.opensearch_cluster_outbound_cluster_config_outbound_clusters_ping_schedule
    }
  }
  reverse_connection_endpoint_customer_ips = var.opensearch_cluster_reverse_connection_endpoint_customer_ips
  search_node_count                        = var.opensearch_cluster_search_node_count
  search_node_host_memory_gb               = var.opensearch_cluster_search_node_host_memory_gb
  search_node_host_ocpu_count              = var.opensearch_cluster_search_node_host_ocpu_count
  search_node_host_shape                   = var.opensearch_cluster_search_node_host_shape
  search_node_host_type                    = var.opensearch_cluster_search_node_host_type
  search_node_storage_gb                   = var.opensearch_cluster_search_node_storage_gb
  security_attributes                      = var.opensearch_cluster_security_attributes
  security_master_user_name                = oci_identity_user.test_user.name
  security_master_user_password_hash       = var.opensearch_cluster_security_master_user_password_hash
  security_mode                            = var.opensearch_cluster_security_mode
  security_saml_config {
    #Required
    idp_entity_id        = oci_opensearch_idp_entity.test_idp_entity.id
    idp_metadata_content = var.opensearch_cluster_security_saml_config_idp_metadata_content
    is_enabled           = var.opensearch_cluster_security_saml_config_is_enabled

    #Optional
    admin_backend_role = var.opensearch_cluster_security_saml_config_admin_backend_role
    opendashboard_url  = var.opensearch_cluster_security_saml_config_opendashboard_url
    roles_key          = var.opensearch_cluster_security_saml_config_roles_key
    subject_key        = var.opensearch_cluster_security_saml_config_subject_key
  }
  system_tags = var.opensearch_cluster_system_tags
}

data "oci_opensearch_opensearch_clusters" "test_opensearch_clusters" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.opensearch_cluster_display_name
  id           = var.opensearch_cluster_id
  state        = var.opensearch_cluster_state
}

