/*
 * This example file shows how to create a user, add an api key, and output their password. 
 */

resource "oci_identity_user" "user1" {
  name = "user1"
  description = "user1 created by terraform"
}

resource "oci_identity_ui_password" "password1" {
  user_id = "${oci_identity_user.user1.id}"
}

resource "oci_identity_api_key" "api-key1" {
  user_id = "${oci_identity_user.user1.id}"
  key_value = 
<<EOF
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

data "oci_identity_users" "users1" {
  compartment_id = "${oci_identity_user.user1.compartment_id}"
  filter {
    name = "name"
    values = ["user1"]
  }
}

output "users1" {
  value = "${data.oci_identity_users.users1.users}"
}

output "user-password" {
  sensitive = false
  value = "${oci_identity_ui_password.password1.password}"
}

output "user-api-key" {
  value = "${oci_identity_api_key.api-key1.key_value}"
}
