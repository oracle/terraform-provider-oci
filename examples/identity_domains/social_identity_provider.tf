// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "social_identity_provider_social_identity_provider_count" {
  default = 10
}

variable "social_identity_provider_social_identity_provider_filter" {
  default = "socialIdentityProviderFilter"
}

variable "social_identity_provider_access_token_url" {
  default = "https://something1.com/token"
}

variable "social_identity_provider_account_linking_enabled" {
  default = false
}

variable "social_identity_provider_admin_scope" {
  default = []
}

variable "social_identity_provider_authorization" {
  default = "authorization"
}

variable "social_identity_provider_authz_url" {
  default = "https://something1.com"
}

variable "social_identity_provider_auto_redirect_enabled" {
  default = false
}

variable "social_identity_provider_client_credential_in_payload" {
  default = false
}

variable "social_identity_provider_clock_skew_in_seconds" {
  default = 10
}

variable "social_identity_provider_compartment_ocid" {
  default = "compartmentOcid"
}

variable "social_identity_provider_consumer_key" {
  default = "consumerKey"
}

variable "social_identity_provider_consumer_secret" {
  default = "consumerSecret"
}

variable "social_identity_provider_delete_in_progress" {
  default = false
}

variable "social_identity_provider_description" {
  default = "description"
}

# variable "social_identity_provider_discovery_url" {
#   default = "discoveryUrl"
# }

variable "social_identity_provider_domain_ocid" {
  default = "domainOcid"
}

variable "social_identity_provider_enabled" {
  default = false
}

# variable "social_identity_provider_icon_url" {
#   default = "iconUrl"
# }

variable "social_identity_provider_id" {
  default = "id"
}

variable "social_identity_provider_id_attribute" {
  default = "idAttribute"
}

variable "social_identity_provider_idcs_created_by__ref" {
  default = "ref"
}

variable "social_identity_provider_idcs_created_by_display" {
  default = "display"
}

variable "social_identity_provider_idcs_created_by_ocid" {
  default = "ocid"
}

variable "social_identity_provider_idcs_created_by_type" {
  default = "User"
}

variable "social_identity_provider_idcs_created_by_value" {
  default = "value"
}

variable "social_identity_provider_idcs_endpoint" {
  default = "idcsEndpoint"
}

variable "social_identity_provider_idcs_last_modified_by_ref" {
  default = "ref"
}

variable "social_identity_provider_idcs_last_modified_by_display" {
  default = "display"
}

variable "social_identity_provider_idcs_last_modified_by_ocid" {
  default = "ocid"
}

variable "social_identity_provider_idcs_last_modified_by_type" {
  default = "User"
}

variable "social_identity_provider_idcs_last_modified_by_value" {
  default = "value"
}

variable "social_identity_provider_idcs_last_upgraded_in_release" {
  default = "idcsLastUpgradedInRelease"
}

variable "social_identity_provider_idcs_prevented_operations" {
  default = []
}

variable "social_identity_provider_jit_prov_assigned_groups_ref" {
  default = "ref"
}

variable "social_identity_provider_jit_prov_assigned_groups_display" {
  default = "display"
}

variable "social_identity_provider_jit_prov_assigned_groups_value" {
  default = "value"
}

variable "social_identity_provider_jit_prov_group_static_list_enabled" {
  default = false
}

variable "social_identity_provider_meta_created" {
  default = "created"
}

variable "social_identity_provider_meta_last_modified" {
  default = "lastModified"
}

variable "social_identity_provider_meta_location" {
  default = "location"
}

variable "social_identity_provider_meta_resource_type" {
  default = "resourceType"
}

variable "social_identity_provider_meta_version" {
  default = "version"
}

variable "social_identity_provider_name" {
  default = "name"
}

variable "social_identity_provider_ocid" {
  default = "ocid"
}

variable "social_identity_provider_profile_url" {
  default = "https://something.com/profileUrl1.png"
}

variable "social_identity_provider_redirect_url" {
  default = "https://redirectUrl1.com"
}

variable "social_identity_provider_refresh_token_url" {
  default = "https://refreshTokenUrl1.com"
}

variable "social_identity_provider_registration_enabled" {
  default = false
}

variable "social_identity_provider_relay_idp_param_mappings_relay_param_key" {
  default = "relayParamKey"
}

variable "social_identity_provider_relay_idp_param_mappings_relay_param_value" {
  default = "relayParamValue"
}


variable "social_identity_provider_schemas" {
  default = []
}

variable "social_identity_provider_scope" {
  default = []
}

variable "social_identity_provider_service_provider_name" {
  default = "Google"
}

variable "social_identity_provider_show_on_login" {
  default = false
}

variable "social_identity_provider_social_jit_provisioning_enabled" {
  default = false
}

variable "social_identity_provider_start_index" {
  default = 1
}

variable "social_identity_provider_status" {
  default = "created"
}

variable "social_identity_provider_tags_key" {
  default = "key"
}

variable "social_identity_provider_tags_value" {
  default = "value"
}

variable "social_identity_provider_tenancy_ocid" {
  default = "tenancyOcid"
}



# provider "oci" {
#   tenancy_ocid     = var.tenancy_ocid
#   user_ocid        = var.user_ocid
#   fingerprint      = var.fingerprint
#   private_key_path = var.private_key_path
#   region           = var.region
# }

resource "oci_identity_domains_social_identity_provider" "test_social_identity_provider" {
  #Required
  account_linking_enabled = var.social_identity_provider_account_linking_enabled
  consumer_key            = var.social_identity_provider_consumer_key
  consumer_secret         = var.social_identity_provider_consumer_secret
  enabled                 = var.social_identity_provider_enabled
  idcs_endpoint           = data.oci_identity_domain.test_domain.url
  name                    = var.social_identity_provider_name
  registration_enabled    = var.social_identity_provider_registration_enabled
  schemas                 = ["urn:ietf:params:scim:schemas:oracle:idcs:SocialIdentityProvider"]
  service_provider_name   = var.social_identity_provider_service_provider_name
  show_on_login           = var.social_identity_provider_show_on_login

  #Optional
  access_token_url             = var.social_identity_provider_access_token_url
  admin_scope                  = var.social_identity_provider_admin_scope
  authorization                = var.social_identity_provider_authorization
  authz_url                    = var.social_identity_provider_authz_url
  auto_redirect_enabled        = var.social_identity_provider_auto_redirect_enabled
  client_credential_in_payload = var.social_identity_provider_client_credential_in_payload
  clock_skew_in_seconds        = var.social_identity_provider_clock_skew_in_seconds
  description                  = var.social_identity_provider_description
#   discovery_url                = var.social_identity_provider_discovery_url
#   icon_url                     = var.social_identity_provider_icon_url
  id_attribute                 = var.social_identity_provider_id_attribute
#   jit_prov_assigned_groups {
#     #Required
#     value = var.social_identity_provider_jit_prov_assigned_groups_value
#   }
  jit_prov_group_static_list_enabled = var.social_identity_provider_jit_prov_group_static_list_enabled
  ocid                               = var.social_identity_provider_ocid
  profile_url                        = var.social_identity_provider_profile_url
  redirect_url                       = var.social_identity_provider_redirect_url
  refresh_token_url                  = var.social_identity_provider_refresh_token_url
  relay_idp_param_mappings {
    #Required
    relay_param_key = var.social_identity_provider_relay_idp_param_mappings_relay_param_key

    #Optional
    relay_param_value = var.social_identity_provider_relay_idp_param_mappings_relay_param_value
  }
  #resource_type_schema_version    = var.social_identity_provider_resource_type_schema_version
  scope                           = var.social_identity_provider_scope
  social_jit_provisioning_enabled = var.social_identity_provider_social_jit_provisioning_enabled
  status                          = var.social_identity_provider_status
  tags {
    #Required
    key   = var.social_identity_provider_tags_key
    value = var.social_identity_provider_tags_value
  }
}

data "oci_identity_domains_social_identity_providers" "test_social_identity_providers" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  social_identity_provider_count  = var.social_identity_provider_social_identity_provider_count
  #social_identity_provider_filter = var.social_identity_provider_social_identity_provider_filter
  social_identity_provider_filter = ""
  authorization                   = var.social_identity_provider_authorization
  # resource_type_schema_version    = var.social_identity_provider_resource_type_schema_version
  start_index                     = var.social_identity_provider_start_index
}