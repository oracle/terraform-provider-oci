// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_outbound_connector" "my_ldap_outbound_connector" {
  #Required
  availability_domain     = data.oci_identity_availability_domain.ad.name
  bind_distinguished_name = var.ldap_outbound_connector_bind_distinguished_name
  compartment_id          = var.compartment_ocid
  connector_type          = "LDAPBIND"
  endpoints {
    #Required
    hostname = var.ldap_outbound_connector_endpoints_hostname
    port     = var.ldap_outbound_connector_endpoints_port
  }
 locks {
    #Required
    type = var.locks_type

    #Optional
    message = var.locks_message
  }
  is_lock_override = var.is_lock_override
  #Optional
  #defined_tags           = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.outbound_connector_defined_tags_value)
  display_name            = var.ldap_outbound_connector_display_name
  #freeform_tags          = var.outbound_connector_freeform_tags
  password_secret_id      = oci_vault_secret.krb_ldap_pwd_secret.id
  password_secret_version = var.ldap_outbound_connector_password_secret_version
}