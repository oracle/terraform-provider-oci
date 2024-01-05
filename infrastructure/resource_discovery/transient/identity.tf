// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_identity_user" "user1_rd" {
  name        = "tf-example-user"
  description = "user created by terraform"
}

resource "oci_identity_group" "identity_group_rd" {
  name        = "identityGroupRD"
  description = "group created by terraform"
}

resource "oci_identity_user_group_membership" "user-group-membership_rd" {
  compartment_id = "${var.tenancy_ocid}"
  user_id        = "${oci_identity_user.user1_rd.id}"
  group_id       = "${oci_identity_group.identity_group_rd.id}"
}

/*
 * Some more directives to show dynamic groups and policy for it
 */
resource "oci_identity_dynamic_group" "dynamic_group_rd" {
  compartment_id = "${var.tenancy_ocid}"
  name           = "tfExampleDynamicGroupRD"
  description    = "dynamic group created by terraform"
  matching_rule  = "instance.compartment.id = ${oci_identity_compartment.identity_compartment_rd.id}"
}

resource "oci_identity_compartment" "identity_compartment_rd" {
  name        = "tfExampleCompartmentRD"
  description = "compartment created by terraform"
}

data "oci_identity_compartments" "compartments1" {
  compartment_id = "${oci_identity_compartment.identity_compartment_rd.compartment_id}"

  filter {
    name   = "name"
    values = ["tf-example-compartment"]
  }
}

resource "oci_identity_api_key" "identity_api-key_rd" {
  user_id = "${oci_identity_user.user1_rd.id}"

  key_value = <<EOF
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtBLQAGmKJ7tpfzYJyqLG
ZDwHL51+d6T8Z00BnP9CFfzxZZZ48PcYSUHuTyCM8mR5JqYLyH6C8tZ/DKqwxUnc
ONgBytG3MM42bgxfHIhsZRj5rCz1oqWlSLuXvgww1kuqWnt6r+NtnXog439YsGTH
RotrTLTdEgOxH0EFP5uHUc9w/Uix7rWU7GB2ra060oeTB/hKpts5U70eI2EI6ec9
1sJdUIj7xNfBJeQQrz4CFUrkyzL06211CFvhmxH2hA9gBKOqC3rGL8XraHZBhGWn
mXlrQB7nNKsJrrv5fHwaPDrAY4iNP2W0q3LRpyNigJ6cgRuGJhHa82iHPmxgIx8m
fwIDAQAB
-----END PUBLIC KEY-----
EOF
}

# SwiftPassword has been deprecated. Use AuthToken instead.
resource "oci_identity_auth_token" "identity_auth_token_rd" {
  #Required
  user_id     = "${oci_identity_user.user1_rd.id}"
  description = "user auth token created by terraform"
}

resource "oci_identity_customer_secret_key" "identity_customer_secret_key_rd" {
  user_id      = "${oci_identity_user.user1_rd.id}"
  display_name = "tf-example-customer-secret-key"
}

resource "oci_identity_policy" "identity_policy_rd" {
  name           = "identityPolicyRD"
  description    = "policy created by terraform"
  compartment_id = "${data.oci_identity_compartments.compartments1.compartments.0.id}"

  statements = ["Allow group ${oci_identity_group.identity_group_rd.name} to read instances in compartment ${data.oci_identity_compartments.compartments1.compartments.0.name}",
    "Allow group ${oci_identity_group.identity_group_rd.name} to inspect instances in compartment ${data.oci_identity_compartments.compartments1.compartments.0.name}",
  ]

  version_date = "2020-02-02"
}

/*
 * This example file shows how to maintain authentication policy for the current tenancy.
 */
resource "oci_identity_authentication_policy" "identity_authentication_policy_rd" {
  #Required
  compartment_id = "${var.tenancy_ocid}"

  #Optional
  password_policy {
    #Optional
    is_lowercase_characters_required = "${var.authentication_policy_password_policy_is_lowercase_characters_required}"
    is_numeric_characters_required   = "${var.authentication_policy_password_policy_is_numeric_characters_required}"
    is_special_characters_required   = "${var.authentication_policy_password_policy_is_special_characters_required}"
    is_uppercase_characters_required = "${var.authentication_policy_password_policy_is_uppercase_characters_required}"
    is_username_containment_allowed  = "${var.authentication_policy_password_policy_is_username_containment_allowed}"
    minimum_password_length          = "${var.authentication_policy_password_policy_minimum_password_length}"
  }
}

resource "oci_identity_ui_password" "identity_ui_password_rd" {
  user_id = "${oci_identity_user.user1_rd.id}"
}

resource "oci_identity_smtp_credential" "identity_smtp_credential_rd" {
  description = "identitySmtpCredentialRD"
  user_id     = "${oci_identity_user.user1_rd.id}"
}

resource "oci_identity_network_source" "identity_network_source_rd" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  description    = "${var.network_source_description}"
  name           = "identityNetworkSourceRD"

  #Optional
  freeform_tags      = "${var.network_source_freeform_tags}"
  public_source_list = "${var.network_source_public_source_list}"
  services           = "${var.network_source_services}"

  virtual_source_list = {
    vcn_id    = "${oci_core_vcn.vcn2_rd.id}"
    ip_ranges = ["10.0.0.0/16"]
  }
}
