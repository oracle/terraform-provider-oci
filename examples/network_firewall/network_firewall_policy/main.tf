// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
  This example shows how to add multiple lists at a time while creating a firewall policy.
*/
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "network_firewall_policy_decryption_profiles_are_certificate_extensions_restricted" {
  default = false
}

variable "network_firewall_policy_decryption_profiles_is_auto_include_alt_name" {
  default = false
}

variable "network_firewall_policy_decryption_profiles_is_expired_certificate_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profiles_is_out_of_capacity_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profiles_is_revocation_status_timeout_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profiles_is_unknown_revocation_status_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profiles_is_unsupported_cipher_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profiles_is_unsupported_version_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profiles_is_untrusted_issuer_blocked" {
  default = false
}

variable "network_firewall_policy_decryption_profiles_type" {
  default = "SSL_INBOUND_INSPECTION"
}

//Application can of type TCP, UDP, or ICMP or ICMP6
variable "network_firewall_policy_application_lists_type" {
  default = "ICMP"
}

variable "network_firewall_policy_application_lists_icmp_type" {

}

variable "network_firewall_policy_application_lists_icmp_code" {

}

variable "network_firewall_policy_application_lists_minimum_port" {

}

variable "network_firewall_policy_application_lists_maximum_port" {

}

variable "network_firewall_policy_url_lists_type" {
  default = "SIMPLE"
}

variable "network_firewall_policy_url_lists_pattern" {
}

variable "network_firewall_policy_decryption_rules_action" {
  default = "DECRYPT"
}

variable "kms_key_ocid" {}

variable "network_firewall_policy_decryption_rules_condition_destinations" {
  default = ["hr_source"]
}

variable "network_firewall_policy_decryption_rules_condition_sources" {
  default = ["hr_source"]
}

variable "network_firewall_policy_decryption_rules_decryption_profile" {
}

variable "network_firewall_policy_decryption_rules_name" {
}

variable "network_firewall_policy_decryption_rules_secret" {
}

variable "network_firewall_policy_defined_tags_value" {
  default = "value"
}

variable "network_firewall_policy_display_name" {
  default = "displayName"
}

variable "network_firewall_policy_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "network_firewall_policy_id" {
  default = ""
}

variable "network_firewall_policy_mapped_secrets_type" {
  default = "SSL_INBOUND_INSPECTION"
}

variable "network_firewall_policy_mapped_secrets_version_number" {
  default = 10
}

variable "network_firewall_policy_security_rules_action" {
}

variable "network_firewall_policy_security_rules_condition_applications" {
  default = ["app-1"]
}

variable "network_firewall_policy_security_rules_condition_destinations" {
  default = ["hr_source"]
}

variable "network_firewall_policy_security_rules_condition_sources" {
  default = ["hr_source"]
}

variable "network_firewall_policy_security_rules_condition_urls" {
  default = ["hr"]
}

variable "network_firewall_policy_security_rules_inspection" {
  default = "INTRUSION_DETECTION"
}

variable "network_firewall_policy_security_rules_name" {
  default = "hr_access"
}

variable "network_firewall_policy_state" {
  default = "ACTIVE"
}

variable  "network_firewall_policy_ip_address_lists_name" {
}

variable  "network_firewall_policy_ip_address_lists_value" {
  default=["10.180.1.0/24", "10.180.2.0/25"]
}

variable "network_firewall_policy_application_list_name" {
}

variable "network_firewall_policy_decryption_profiles_key" {
}

variable "network_firewall_policy_mapped_secrets_key" {
}

variable "network_firewall_policy_url_lists_key" {
}

variable "kms_vault_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_vault_secret" "test_secret" {
  #Required
  compartment_id = var.compartment_id
  secret_content {
    #Required
    content_type = "BASE64"

    #Optional
    content = "PHZhcj4mbHQ7YmFzZTY0X2VuY29kZWRfc2VjcmV0X2NvbnRlbnRzJmd0OzwvdmFyPg=="
    name    = "name"
    stage   = "CURRENT"
  }
  key_id = var.kms_key_ocid
  secret_name = "TFsample1"
  vault_id    = var.kms_vault_ocid
}

resource "oci_network_firewall_network_firewall_policy" "test_network_firewall_policy" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  //application_lists = var.network_firewall_policy_application_lists
  application_lists {
    #Required
    application_list_name          = var.network_firewall_policy_application_list_name
    application_values {
      type                           = var.network_firewall_policy_application_lists_type
      #Optional
      icmp_type                      = var.network_firewall_policy_application_lists_icmp_type
      icmp_code                      = var.network_firewall_policy_application_lists_icmp_code
      minimum_port                   = var.network_firewall_policy_application_lists_minimum_port
      maximum_port                   = var.network_firewall_policy_application_lists_maximum_port
    }
  }
  #Optional
  application_lists {
    #Required
    application_list_name            = var.network_firewall_policy_application_list_name
    application_values {
      type                           = var.network_firewall_policy_application_lists_type
      #Optional
      icmp_type                   = var.network_firewall_policy_application_lists_icmp_type
      icmp_code                   = var.network_firewall_policy_application_lists_icmp_code
    }
    application_values {
      type                           = var.network_firewall_policy_application_lists_type
      #Optional
      minimum_port                   = var.network_firewall_policy_application_lists_minimum_port
      maximum_port                   = var.network_firewall_policy_application_lists_maximum_port
    }
  }
  decryption_profiles {
    #Required
    is_out_of_capacity_blocked     = var.network_firewall_policy_decryption_profiles_is_out_of_capacity_blocked
    is_unsupported_cipher_blocked  = var.network_firewall_policy_decryption_profiles_is_unsupported_cipher_blocked
    is_unsupported_version_blocked = var.network_firewall_policy_decryption_profiles_is_unsupported_version_blocked
    type                           = var.network_firewall_policy_decryption_profiles_type
    key                            = var.network_firewall_policy_decryption_profiles_key

    #Optional
    are_certificate_extensions_restricted = var.network_firewall_policy_decryption_profiles_are_certificate_extensions_restricted
    is_auto_include_alt_name              = var.network_firewall_policy_decryption_profiles_is_auto_include_alt_name
    is_expired_certificate_blocked        = var.network_firewall_policy_decryption_profiles_is_expired_certificate_blocked
    is_revocation_status_timeout_blocked  = var.network_firewall_policy_decryption_profiles_is_revocation_status_timeout_blocked
    is_unknown_revocation_status_blocked  = var.network_firewall_policy_decryption_profiles_is_unknown_revocation_status_blocked
    is_untrusted_issuer_blocked           = var.network_firewall_policy_decryption_profiles_is_untrusted_issuer_blocked
  }
  decryption_rules {
    #Required
    action = var.network_firewall_policy_decryption_rules_action
    condition {

      #Optional
      destinations = var.network_firewall_policy_decryption_rules_condition_destinations
      sources      = var.network_firewall_policy_decryption_rules_condition_sources
    }
    name = var.network_firewall_policy_decryption_rules_name

    #Optional
    decryption_profile = var.network_firewall_policy_decryption_rules_decryption_profile
    secret             = var.network_firewall_policy_decryption_rules_secret
  }
  #defined_tags = {
   # "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"
  #}
  display_name     = var.network_firewall_policy_display_name
  freeform_tags    = var.network_firewall_policy_freeform_tags
  ip_address_lists {
    ip_address_list_name = var.network_firewall_policy_ip_address_lists_name
    ip_address_list_value = var.network_firewall_policy_ip_address_lists_value
  }
  mapped_secrets {
    #Required
    type            = var.network_firewall_policy_mapped_secrets_type
    key             = var.network_firewall_policy_mapped_secrets_key
    vault_secret_id = oci_vault_secret.test_secret.id
    version_number  = var.network_firewall_policy_mapped_secrets_version_number
  }
  security_rules {
    #Required
    action = var.network_firewall_policy_security_rules_action
    condition {

      #Optional
      applications = var.network_firewall_policy_security_rules_condition_applications
      destinations = var.network_firewall_policy_security_rules_condition_destinations
      sources      = var.network_firewall_policy_security_rules_condition_sources
      urls         = var.network_firewall_policy_security_rules_condition_urls
    }
    name = var.network_firewall_policy_security_rules_name

    #Optional
    inspection = var.network_firewall_policy_security_rules_inspection
  }
  url_lists {
    url_list_name                = var.network_firewall_policy_url_lists_key
    url_list_values {
      type                           = var.network_firewall_policy_url_lists_type
      pattern                        = var.network_firewall_policy_url_lists_pattern
    }
  }
  url_lists {
    url_list_name                = var.network_firewall_policy_url_lists_key
    url_list_values {
      type                           = var.network_firewall_policy_url_lists_type
      pattern                        = var.network_firewall_policy_url_lists_pattern
    }
    url_list_values {
      type                           = var.network_firewall_policy_url_lists_type
      pattern                        = var.network_firewall_policy_url_lists_pattern
    }
  }
}

data "oci_network_firewall_network_firewall_policies" "test_network_firewall_policies" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.network_firewall_policy_display_name
  id           = var.network_firewall_policy_id
  state        = var.network_firewall_policy_state
}
