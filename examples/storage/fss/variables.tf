// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

# Refer https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingkeypairs.htm on how to setup SSH key pairs for compute instances
variable "ssh_public_key" {
}

variable "ssh_private_key" {
}

variable "my_vcn-cidr" {
  default = "10.0.0.0/16"
}

variable "my_subnet_cidr" {
  default = "10.0.1.0/24"
}

variable "file_system_1_display_name" {
  default = "my_fs_1"
}

variable "file_system_2_display_name" {
  default = "my_fs_2"
}

variable "file_system_clone_display_name" {
	default= "my_fs_clone"
}

variable "file_system_simple_display_name" {
	default= "my_fs_simple"
}

variable "file_system_with_snapshot_policy_display_name" {
  default = "my_fs_with_snapshot_policy"
}

variable "mount_target_1_display_name" {
  default = "my_mount_target_1"
}

variable "mount_target_2_display_name" {
  default = "my_mount_target_2"
}

variable "export_path_fs1_mt1" {
  default = "/myfsspaths/fs1/path1"
}

variable "export_path_fs1_mt2" {
  default = "/myfsspaths/fs1/path2"
}

variable "export_path_fs2_mt1" {
  default = "/myfsspaths/fs2/path1"
}

variable "snapshot_name" {
  default = "20180320_daily"
}

variable "snapshot_name_clone" {
  default = "snapshot_clone"
}

variable "export_set_name_1" {
  default = "export set for mount target 1"
}

variable "export_set_name_2" {
  default = "export set for mount target 2"
}

variable "max_byte" {
  default = 23843202333
}

variable "max_files" {
  default = 223442
}

variable "export_read_write_access_source" {
  default = "10.0.0.0/8"
}

variable "export_read_only_access_source" {
  default = "0.0.0.0/0"
}

variable "instance_image_ocid" {
  type = map(string)

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.05.09-1"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaazregkysspxnktw35k4r5vzwurxk6myu44umqthjeakbkvxvxdlkq"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaa6ybn2lkqp2ejhijhehf5i65spqh3igt53iyvncyjmo7uhm5235ca"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaayodsld656eh5stds5mo4hrmwuhk2ugin4eyfpgoiiskqfxll6a4a"
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaaozjbzisykoybkppaiwviyfzusjzokq7jzwxi7nvwdiopk7ligoia"
  }
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

locals {
  mount_target_1_ip_address = data.oci_core_private_ips.ip_mount_target1.private_ips[0]["ip_address"]
}

variable "filesystem_snapshot_policy_display_name" {
  default = "media-policy-1"
}

variable "filesystem_snapshot_policy_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "filesystem_snapshot_policy_id" {
  default = "id"
}

variable "filesystem_snapshot_policy_policy_prefix" {
  default = "mp1"
}

variable "filesystem_snapshot_policy_schedules_day_of_month" {
  default = 10
}

variable "filesystem_snapshot_policy_schedules_hour_of_day" {
  default = 10
}

variable "filesystem_snapshot_policy_schedules_month" {
  default = "JANUARY"
}

variable "filesystem_snapshot_policy_schedules_day_of_week" {
  default = "MONDAY"
}

variable "filesystem_snapshot_policy_schedules_retention_duration_in_seconds" {
  default = 7200
}

variable "filesystem_snapshot_policy_schedules_schedule_prefix" {
  default = "schedulePrefix"
}

variable "filesystem_snapshot_policy_schedules_time_zone" {
  default = "UTC"
}

variable "filesystem_snapshot_policy_state" {
  default = "ACTIVE"
}


variable "krb_mount_target_display_name" {
  default = "my_krb_mount_target"
}

variable "krb_mount_target_hostname_label" {
  default = "hostnamelabel"
}

variable "krb_mount_target_kerberos_kerberos_realm" {
  default = "kerberos.realm.com"
}

variable "krb_mount_target_kerberos_backup_key_tab_secret_version" {
  default = 0
}

variable "krb_mount_target_kerberos_current_key_tab_secret_version" {
  default = 1
}

variable "krb_mount_target_krb_enabled" {
  default = "true"
}

variable "krb_mount_target_group_name" {
  default = "group_name"
}

variable "krb_mount_target_user_name" {
  default = "user_name"
}

variable "krb_mount_target_ldap_idmap_cache_lifetime_seconds" {
  default = 300
}

variable "krb_mount_target_ldap_idmap_cache_refresh_interval_seconds" {
  default = 300
}

variable "krb_mount_target_ldap_idmap_negative_cache_lifetime_seconds" {
  default = 300
}

variable "ldap_outbound_connector_display_name" {
  default = "my_ldap_outbound_connector"
}

variable "ldap_outbound_connector_bind_distinguished_name" {
  default = "bindDistinguishedName"
}

variable "ldap_outbound_connector_endpoints_hostname" {
  default = "hostname"
}

variable "ldap_outbound_connector_endpoints_port" {
  default = 1080
}

variable "ldap_outbound_connector_password_secret_version" {
  default = 1
}

variable "krb_vault_display_name" {
  default = "my_krb_vault"
}

variable "krb_vault_type" {
  default = "DEFAULT"
}

variable "krb_key_display_name" {
  default = "my_krb_key"
}

variable "krb_key_shape_algorithm" {
  default = "AES"
}

variable "krb_key_shape_length" {
  default = "16"
}

variable "krb_keytab_content" {
  default = "BQIAAAClAAIAI0FEMkNBTkFSWS5QSFhERVZQQ0FOUy5PUkFDTEVWQ04uQ09NAANuZnMARmtlcmJlcm9zLWFwaS1jYW5hcnktbW91bnQtdGFyZ2V0LTEuYWQyY2FuYXJ5LnBoeGRldnBjYW5zLm9yYWNsZXZjbi5jb20AAAABYgMUPgIAEgAgIvKmyzN+v/xsEQpwSzwxfFCEwtbV5ozYkk8VAmx9NhQAAAAC"
}

variable "krb_ldap_pwd_content" {
  default = "dGVzdHB3ZAo="
}

variable "export_path_kfs_kmt" {
  default = "/myfsspaths/kfs/path1"
}

variable "krb_file_system" {
  default = "my_krb_file_system"
}

variable "krb_export_export_options_allowed_auth" {
  default = ["KRB5"]
}

variable "krb_export_export_options_is_anonymous_access_allowed" {
  default = "true"
}

variable "krb_export_is_idmap_groups_for_sys_auth" {
  default = "false"
}
