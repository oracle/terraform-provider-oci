// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_mount_target" "my_mount_target_1" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  subnet_id           = oci_core_subnet.my_subnet.id

  #Optional
  display_name = var.mount_target_1_display_name
  # defined_tags = {
  #   "example-tag-namespace-all.example-tag" = "value"
  # }

  freeform_tags = {
    "Department" = "Finance"
  }

  nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
}

resource "oci_file_storage_mount_target" "my_mount_target_2" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  subnet_id           = oci_core_subnet.my_subnet.id

  #Optional
  display_name = var.mount_target_2_display_name
  # defined_tags = {
  #   "example-tag-namespace-all.example-tag" = "value"
  # }

  freeform_tags = {
    "Department" = "Accounting"
  }

  nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
}

resource "oci_file_storage_mount_target" "my_krb_mount_target" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  subnet_id           = oci_core_subnet.my_subnet.id

  #Optional
  # defined_tags   = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.mount_target_defined_tags_value)
  display_name   = var.krb_mount_target_display_name
  # freeform_tags  = {
  #  "Department" = "Accounting"
  # }
  hostname_label = var.krb_mount_target_hostname_label
  idmap_type     = "LDAP"
  kerberos {
    #Required
    kerberos_realm = var.krb_mount_target_kerberos_kerberos_realm

    #Optional
    backup_key_tab_secret_version  = var.krb_mount_target_kerberos_backup_key_tab_secret_version
    current_key_tab_secret_version = var.krb_mount_target_kerberos_current_key_tab_secret_version
    is_kerberos_enabled            = var.krb_mount_target_krb_enabled
    key_tab_secret_id              = oci_vault_secret.krb_keytab_secret.id
  }
  ldap_idmap {
    #Required
    group_search_base = var.krb_mount_target_group_name
    user_search_base  = var.krb_mount_target_user_name

    #Optional
    cache_lifetime_seconds          = var.krb_mount_target_ldap_idmap_cache_lifetime_seconds
    cache_refresh_interval_seconds  = var.krb_mount_target_ldap_idmap_cache_refresh_interval_seconds
    negative_cache_lifetime_seconds = var.krb_mount_target_ldap_idmap_negative_cache_lifetime_seconds
    outbound_connector1id           = oci_file_storage_outbound_connector.my_ldap_outbound_connector.id
    # outbound_connector2id         = oci_file_storage_outbound_connector.test_outbound_connector2.id
    schema_type                     = "RFC2307"
  }
  nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
}

# Use export_set.tf config to update the size for a mount target
