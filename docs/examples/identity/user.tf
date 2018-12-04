/*
 * This example file shows how to create a user, add an api key, define auth tokens and customer secret keys.
 */

resource "oci_identity_user" "user1" {
  name           = "tf-example-user"
  description    = "user created by terraform"
  compartment_id = "${var.tenancy_ocid}"
}

// Use the "user2" to have non-default values of capabilities without corresponding authentication resources being actually created
resource "oci_identity_user" "user2" {
  name           = "tf-example-user2"
  description    = "user2 created by terraform"
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_identity_user_capabilities_management" "user2-capabilities-management" {
  user_id                      = "${oci_identity_user.user2.id}"
  can_use_api_keys             = "false"
  can_use_auth_tokens          = "false"
  can_use_console_password     = "false"
  can_use_customer_secret_keys = "false"
  can_use_smtp_credentials     = "false"
}

data "oci_identity_users" "users1" {
  compartment_id = "${oci_identity_user.user1.compartment_id}"

  filter {
    name   = "name"
    values = ["tf-example-user"]
  }
}

data "oci_identity_users" "users2" {
  compartment_id = "${oci_identity_user.user1.compartment_id}"

  filter {
    name   = "id"
    values = ["${oci_identity_user.user2.id}"]
  }
}

output "users1" {
  value = "${data.oci_identity_users.users1.users}"
}

output "users2" {
  value = "${data.oci_identity_users.users2.users}"
}

resource "oci_identity_ui_password" "password1" {
  user_id = "${oci_identity_user.user1.id}"
}

output "user-password" {
  sensitive = false
  value     = "${oci_identity_ui_password.password1.password}"
}

resource "oci_identity_api_key" "api-key1" {
  user_id = "${oci_identity_user.user1.id}"

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

output "user-api-key" {
  value = "${oci_identity_api_key.api-key1.key_value}"
}

# SwiftPassword has been deprecated. Use AuthToken instead.
resource "oci_identity_auth_token" "auth-token1" {
  #Required
  user_id     = "${oci_identity_user.user1.id}"
  description = "user auth token created by terraform"
}

output "auth-token" {
  value = "${oci_identity_auth_token.auth-token1.token}"
}

resource "oci_identity_customer_secret_key" "customer-secret-key1" {
  user_id      = "${oci_identity_user.user1.id}"
  display_name = "tf-example-customer-secret-key"
}

data "oci_identity_customer_secret_keys" "customer-secret-keys1" {
  user_id = "${oci_identity_customer_secret_key.customer-secret-key1.user_id}"
}

output "customer-secret-key" {
  value = [
    "${oci_identity_customer_secret_key.customer-secret-key1.key}",
    "${data.oci_identity_customer_secret_keys.customer-secret-keys1.customer_secret_keys}",
  ]
}

resource "oci_identity_smtp_credential" "smtp-credential-1" {
  description = "tf-example-smtp-credential"
  user_id     = "${oci_identity_user.user1.id}"
}

data "oci_identity_smtp_credentials" "smtp-credentials-1" {
  user_id = "${oci_identity_smtp_credential.smtp-credential-1.user_id}"
}

output "smtp-credential" {
  value = [
    "${oci_identity_smtp_credential.smtp-credential-1.password}",
    "${data.oci_identity_smtp_credentials.smtp-credentials-1.smtp_credentials}",
  ]
}
