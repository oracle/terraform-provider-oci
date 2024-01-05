// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_waas_certificate" "certificate_rd" {
  #Required
  certificate_data = "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"
  compartment_id   = "${var.compartment_ocid}"
  private_key_data = "${var.private_key_data}"

  #Optional
  display_name                   = "${var.certificate_display_name}"
  is_trust_verification_disabled = true
}

resource "oci_waas_custom_protection_rule" "custom_protection_rule_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "tf_example_protection_rule_rd"
  template       = "SecRule REQUEST_URI / \"phase:2,   t:none,   capture,   msg:'Custom (XSS) Attack. Matched Data: %%{TX.0}   found within %%{MATCHED_VAR_NAME}: %%{MATCHED_VAR}',   id:{{id_1}},   ctl:ruleEngine={{mode}},   tag:'Custom',   severity:'2'\""

  #Optional
  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
  description = "Tf example custom protection rule"

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_waas_address_list" "address_list_rd" {
  #Required
  addresses      = ["0.0.0.0/16", "192.168.0.0/20"]
  compartment_id = "${var.compartment_ocid}"
  display_name   = "tf-example-address-list"

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_waas_waas_policy" "waas_policy_rd" {
  compartment_id = "${var.compartment_ocid}"
  domain         = "${var.waas_policy_domain}"
}

resource "oci_waas_http_redirect" "http_redirect_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  domain         = "${var.waas_http_redirect_domain}"

  target {
    #Required
    host     = "${var.waas_http_redirect_host}"
    path     = "/test{path}"
    protocol = "HTTP"
    query    = "{query}"

    #Optional
    port = "8080"
  }
}
