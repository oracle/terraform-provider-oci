// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "certificate_display_name" {
  default = "tf_example_waas_certificate"
}

variable "waas_policy_display_name" {
  default = "tf_example_waas_policy"
}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_waas_certificate" "test_certificate" {
  #Required
  certificate_data = "-----BEGIN CERTIFICATE-----\nMIIC9jCCAd4CCQD2rPUVJETHGzANBgkqhkiG9w0BAQsFADA9MQswCQYDVQQGEwJV\nUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBAoMBk9yYWNs\nZTAeFw0xOTAxMTcyMjU4MDVaFw0yMTAxMTYyMjU4MDVaMD0xCzAJBgNVBAYTAlVT\nMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGT3JhY2xl\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA30+wt7OlUB/YpmWbTRkx\nnLG0lKWiV+oupNKj8luXmC5jvOFTUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU\n+DWVV2So2B/obYxpiiyWF2tcF/cYi1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oC\nMQ2985/MTdCXONgnbmePU64GrJwfvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOO\njLKRM68KXC5us4879IrSA77NQr1KwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6\nytM66P/1CTpk1YpbI4gqiG0HBbuXG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc\n2wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAuI53m8Va6EafDi6GQdQrzNNQFCAVQ\nxIABAB0uaSYCs3H+pqTktHzOrOluSUEogXRl0UU5/OuvxAz4idA4cfBdId4i7AcY\nqZsBjA/xqH/rxR3pcgfaGyxQzrUsJFf0ZwnzqYJs7fUvuatHJYi/cRBxrKR2+4Oj\nlUbb9TSmezlzHK5CaD5XzN+lZqbsSvN3OQbOryJCbtjZVQFGZ1SmL6OLrwpbBKuP\nn2ob+gaP57YSzO3zk1NDXMlQPHRsdSOqocyKx8y+7J0g6MqPvBzIe+wI3QW85MQY\nj1/IHmj84LNGp7pHCyiYx/oI+00gRch04H2pJv0TP3sAQ37gplBwDrUo\n-----END CERTIFICATE-----"
  compartment_id   = "${var.compartment_ocid}"
  private_key_data = "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA30+wt7OlUB/YpmWbTRkxnLG0lKWiV+oupNKj8luXmC5jvOFT\nUejt1pQhpA47nCqywlOAfk2N8hJWTyJZUmKU+DWVV2So2B/obYxpiiyWF2tcF/cY\ni1kBYeAIu5JkVFwDe4ITK/oQUFEhIn3Qg/oCMQ2985/MTdCXONgnbmePU64GrJwf\nvOeJcQB3VIL1BBfISj4pPw5708qTRv5MJBOOjLKRM68KXC5us4879IrSA77NQr1K\nwjGnQlykyCgGvvgwgrUTd5c/dH8EKrZVcFi6ytM66P/1CTpk1YpbI4gqiG0HBbuX\nG4JRIjyzW4GT4JXeSjgvrkIYL8k/M4Az1WEc2wIDAQABAoIBAGQznukfG/uS/qTT\njNcQifl0p8HXfLwUIa/lsJkMTj6D+k8DkF59tVMGjv3NQSQ26JVX4J1L8XiAj+fc\nUtYr1Ap4CLX5PeYUkzesvKK6lPKXQvCh+Ip2eq9PVrvL2WcdDpb5695cy7suXD7c\n05aUtS0LrINH3eXAxkpEe5UHtQFni5YLrCLEXd+SSA3OKdCB+23HRELc1iCTvqjK\n5AtR916uHTBhtREHRMvWIdu4InRUsedlJhaJOLJ8G8r64JUtfm3wLUK1U8HFOsd0\nLAx9ZURU6cXl4osTWiy1vigGaM8Xuish2HkOLNYZADDUiDBB3SshmW5IDAJ5XTn5\nqVrszRECgYEA79j1y+WLTyV7yz7XkWk3OqoQXG4b2JfKItJI1M95UwllzQ8U/krM\n+QZjP3NTtB9i1YoHyaEfic103hV9Fkgz8jvKS5ocLGJulpN4CgqbHN6v9EJ3dqTk\no6X8mpx2eP2E0ngRekFyC/OCp0Zhe2KR9PXhijMa5eB2LTeCMIS/tzkCgYEA7lmk\nIdVjcpfqY7UFJ2R8zqPJHOne2+llrl9vzo6N5kx4DzAg7MP6XO9MekOvfmD1X1Lm\nFckXWFEF+0TlN5YvCTR/+OmVufYM3xp4GBT8RZdLFbyI4+xpAAeSC4SeM0ZkC9Jt\nrKqCS24+Kqy/+qSqtkxiPLQrXSdCSfCUlmn0ALMCgYBB7SLy3q+CG82BOk7Km18g\n8un4XhOtX1uiYqa+SCETH/wpd0HP/AOHV6gkIrEZS59BDuXBGFaw7BZ5jPKLE2Gj\n7adXTI797Dh1jydpqyyjrNo0i6iGpiBqkw9x+Bvged7ucy5qql6MxmxdSk01Owzf\nhk5uTEnScfZJy34vk+2WkQKBgBXx5uy+iuN4HTqE5i6UT/FunwusdLpmqNf/LXol\nIed8TumHEuD5wklgNvhi1vuZzb2zEkAbPa0B+L0DwN73UulUDhxK1WBDyTeZZklB\nVWDK5zzfGPNzRs+b4tRwp2gtKPT1sOde45QyWELxmNNo6dbS/ZB9Pijbfnz0S5n1\ns2OFAoGBAJUohI1+d2hKlkSUzpCorQarDe8lFVEbGMu0kX0JNDI7QU+H8vDp9NOl\nGqLm3sCVBYypT8sWfchgZpcVaLRLQCQtWy4+CbMN6DT3j/uBWeDpayU5Gvqt0/no\nvwqbG6b0NEYLRPLEdsS/c8TV9mMlvb0EW+GXfmkpTrTNt3hyXniu\n-----END RSA PRIVATE KEY-----"

  #Optional
  display_name                   = "${var.certificate_display_name}"
  is_trust_verification_disabled = true
}

data "oci_waas_certificates" "test_certificates" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_names = ["${var.certificate_display_name}"]
  ids           = ["${oci_waas_certificate.test_certificate.id}"]
}

resource "oci_waas_custom_protection_rule" "test_custom_protection_rule" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "tf_example_protection_rule"
  template       = "SecRule REQUEST_URI / \"phase:2,   t:none,   capture,   msg:'Custom (XSS) Attack. Matched Data: %%{TX.0}   found within %%{MATCHED_VAR_NAME}: %%{MATCHED_VAR}',   id:{{id_1}},   ctl:ruleEngine={{mode}},   tag:'Custom',   severity:'2'\""

  #Optional
  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
  description = "Tf example custom protection rule"

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_waas_address_list" "test_address_list" {
  #Required
  addresses      = ["0.0.0.0/16", "192.168.0.0/20"]
  compartment_id = "${var.compartment_ocid}"
  display_name   = "tf-example-address-list"

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_waas_waas_policy" "test_waas_policy" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  domain         = "somethingnew54.oracle.com"

  #Optional
  additional_domains = ["somethingnew55.oracle.com", "somethingnew56.oracle.com"]
  display_name       = "${var.waas_policy_display_name}"

  origin_groups {
    label = "originGroups1"

    origin_group {
      origin = "primary"
      weight = "1"
    }

    origin_group {
      origin = "secondary"
      weight = "2"
    }
  }

  origin_groups {
    label = "originGroups2"

    origin_group {
      origin = "primary"
      weight = "1"
    }

    origin_group {
      origin = "secondary"
      weight = "2"
    }
  }

  origins {
    #Required
    label = "primary"
    uri   = "192.168.0.1"

    #Optional
    custom_headers {
      #Required
      name  = "name1"
      value = "value1"
    }

    custom_headers {
      #Required
      name  = "name2"
      value = "value2"
    }

    http_port  = "80"
    https_port = "443"
  }

  origins {
    #Required
    label = "secondary"
    uri   = "192.168.0.2"

    #Optional
    custom_headers {
      #Required
      name  = "name3"
      value = "value3"
    }

    custom_headers {
      #Required
      name  = "name4"
      value = "value4"
    }

    http_port  = "8080"
    https_port = "8443"
  }

  policy_config {
    #Optional
    certificate_id                = "${oci_waas_certificate.test_certificate.id}"
    cipher_group                  = "DEFAULT"
    client_address_header         = "X_FORWARDED_FOR"
    is_behind_cdn                 = true
    is_cache_control_respected    = true
    is_https_enabled              = true
    is_https_forced               = true
    is_origin_compression_enabled = true
    is_response_buffering_enabled = true
    is_sni_enabled                = true
    websocket_path_prefixes       = ["/url1"]
    tls_protocols                 = ["TLS_V1_1"]

    health_checks = {
      is_enabled                     = true
      is_response_text_check_enabled = true
      expected_response_code_group   = ["2XX"]
      expected_response_text         = "testResponseText"

      headers = {
        "Host"       = "oracle.com"
        "User-Agent" = "Oracle-TerraformProvider"
      }

      method              = "GET"
      path                = "/testpath"
      timeout_in_seconds  = 10
      unhealthy_threshold = 5
      interval_in_seconds = 10
      healthy_threshold   = 2
    }

    load_balancing_method = {
      method                     = "STICKY_COOKIE"
      name                       = "name2"
      domain                     = "example.com"
      expiration_time_in_seconds = 10
    }
  }

  timeouts {
    create = "120m"
    delete = "120m"
    update = "120m"
  }

  waf_config {
    #Optional
    access_rules {
      #Required
      action = "ALLOW"

      criteria {
        #Required
        condition = "URL_IS"
        value     = "/public"
      }

      name = "tf_example_access_rule"

      #Optional
      block_action                 = "SET_RESPONSE_CODE"
      block_error_page_code        = 403
      block_error_page_description = "blockErrorPageDescription"
      block_error_page_message     = "blockErrorPageMessage"
      block_response_code          = 403
      bypass_challenges            = ["JS_CHALLENGE"]
      redirect_response_code       = "FOUND"
      redirect_url                 = "http://192.168.0.3"
      captcha_footer               = "captchaFooter"
      captcha_header               = "captchaHeader"
      captcha_submit_label         = "captchaSubmitLabel"
      captcha_title                = "captchaTitle"

      response_header_manipulation = {
        #Required
        action = "EXTEND_HTTP_RESPONSE_HEADER"
        header = "header"

        #Optional
        value = "value"
      }
    }

    address_rate_limiting {
      #Required
      is_enabled = true

      #Optional
      allowed_rate_per_address      = 10
      block_response_code           = 403
      max_delayed_count_per_address = 10
    }

    caching_rules {
      #Required
      action = "CACHE"

      criteria {
        #Required
        condition = "URL_IS"
        value     = "/public"
      }

      name = "name"

      #Optional
      caching_duration          = "PT1S"
      client_caching_duration   = "PT1S"
      is_client_caching_enabled = false
      key                       = "key"
    }

    captchas {
      #Required
      failure_message               = "message"
      session_expiration_in_seconds = 10
      submit_label                  = "label"
      title                         = "title"
      url                           = "url"

      #Optional
      footer_text = "footer_text"
      header_text = "header_text"
    }

    custom_protection_rules {
      #Optional
      action = "DETECT"
      id     = "${oci_waas_custom_protection_rule.test_custom_protection_rule.id}"

      exclusions = {
        exclusions = ["example.com"]
        target     = "REQUEST_COOKIES"
      }
    }

    device_fingerprint_challenge {
      #Required
      is_enabled = true

      #Optional
      action                       = "DETECT"
      action_expiration_in_seconds = 10

      challenge_settings {
        #Optional
        block_action                 = "SET_RESPONSE_CODE"
        block_error_page_code        = 403
        block_error_page_description = "blockErrorPageDescription"
        block_error_page_message     = "blockErrorPageMessage"
        block_response_code          = 403
        captcha_footer               = "captchaFooter"
        captcha_header               = "captchaHeader"
        captcha_submit_label         = "captchaSubmitLabel"
        captcha_title                = "captchaTitle"
      }

      failure_threshold                       = 10
      failure_threshold_expiration_in_seconds = 10
      max_address_count                       = 10
      max_address_count_expiration_in_seconds = 10
    }

    human_interaction_challenge {
      #Required
      is_enabled = true

      #Optional
      action                       = "DETECT"
      action_expiration_in_seconds = 10

      challenge_settings {
        #Optional
        block_action                 = "SET_RESPONSE_CODE"
        block_error_page_code        = 403
        block_error_page_description = "blockErrorPageDescription"
        block_error_page_message     = "blockErrorPageMessage"
        block_response_code          = 403
        captcha_footer               = "captchaFooter"
        captcha_header               = "captchaHeader"
        captcha_submit_label         = "captchaSubmitLabel"
        captcha_title                = "captchaTitle"
      }

      failure_threshold                       = 10
      failure_threshold_expiration_in_seconds = 10
      interaction_threshold                   = 10
      recording_period_in_seconds             = 10

      set_http_header {
        #Required
        name  = "hc_name1"
        value = "hc_value1"
      }
    }

    js_challenge {
      #Required
      is_enabled = true

      #Optional
      action                       = "DETECT"
      action_expiration_in_seconds = 10
      are_redirects_challenged     = true
      is_nat_enabled               = true

      criteria = {
        #Required
        condition = "URL_IS"
        value     = "/public"

        #Optional
        is_case_sensitive = true
      }

      challenge_settings {
        #Optional
        block_action                 = "SET_RESPONSE_CODE"
        block_error_page_code        = 403
        block_error_page_description = "blockErrorPageDescription"
        block_error_page_message     = "blockErrorPageMessage"
        block_response_code          = 403
        captcha_footer               = "captchaFooter"
        captcha_header               = "captchaHeader"
        captcha_submit_label         = "captchaSubmitLabel"
        captcha_title                = "captchaTitle"
      }

      failure_threshold = 10
    }

    origin        = "primary"
    origin_groups = ["originGroups1"]

    protection_settings {
      #Optional
      allowed_http_methods               = ["OPTIONS", "HEAD"]
      block_action                       = "SET_RESPONSE_CODE"
      block_error_page_code              = 403
      block_error_page_description       = "blockErrorPageDescription"
      block_error_page_message           = "blockErrorPageMessage"
      block_response_code                = 403
      is_response_inspected              = false
      max_argument_count                 = 10
      max_name_length_per_argument       = 10
      max_response_size_in_ki_b          = 10
      max_total_name_length_of_arguments = 10
      media_types                        = ["application/plain", "application/json"]
      recommendations_period_in_days     = 10
    }

    whitelists {
      #Required
      name = "whitelist_name"

      #Optional
      addresses     = ["192.168.127.127", "192.168.127.128"]
      address_lists = ["${oci_waas_address_list.test_address_list.id}`}"]
    }
  }
}

resource "oci_waas_protection_rule" "test_waas_protection_rule" {
  waas_policy_id = "${oci_waas_waas_policy.test_waas_policy.id}"
  key            = "933161"
  action         = "DETECT"

  exclusions = {
    exclusions = ["example.com"]
    target     = "REQUEST_COOKIES"
  }
}

data "oci_waas_protection_rules" "test_waas_protection_rules" {
  waas_policy_id = "${oci_waas_waas_policy.test_waas_policy.id}"
  action         = ["DETECT"]
}

data "oci_waas_protection_rule" "test_waas_protection_rule" {
  waas_policy_id      = "${oci_waas_waas_policy.test_waas_policy.id}"
  protection_rule_key = "${oci_waas_protection_rule.test_waas_protection_rule.key}"
}

data "oci_waas_waas_policies" "test_waas_policies" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_names                         = ["${var.waas_policy_display_name}"]
  ids                                   = ["${oci_waas_waas_policy.test_waas_policy.id}"]
  states                                = ["ACTIVE"]
  time_created_greater_than_or_equal_to = "2018-01-01T00:00:00.000Z"
  time_created_less_than                = "2038-01-01T00:00:00.000Z"
}

data "oci_waas_address_lists" "test_address_lists" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  ids                                   = ["${oci_waas_address_list.test_address_list.id}`}"]
  names                                 = ["tf-example-address-list"]
  states                                = ["ACTIVE"]
  time_created_greater_than_or_equal_to = "2018-01-01T00:00:00.000Z"
  time_created_less_than                = "2038-01-01T00:00:00.000Z"
}

data "oci_waas_custom_protection_rules" "test_custom_protection_rules" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_names                         = ["tf_example_protection_rule"]
  ids                                   = ["${oci_waas_custom_protection_rule.test_custom_protection_rule.id}"]
  states                                = ["ACTIVE"]
  time_created_greater_than_or_equal_to = "2018-01-01T00:00:00.000Z"
  time_created_less_than                = "2038-01-01T00:00:00.000Z"
}

resource "oci_waas_http_redirect" "test_http_redirect" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  domain         = "example3.com"

  target {
    #Required
    host     = "example4.com"
    path     = "/test{path}"
    protocol = "HTTP"
    query    = "{query}"

    #Optional
    port = "8080"
  }

  #Optional
  display_name = "displayName"

  freeform_tags = {
    "Department" = "Finance"
  }

  response_code = 301
}

data "oci_waas_http_redirects" "test_http_redirects" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_names                         = ["${oci_waas_http_redirect.test_http_redirect.display_name}"]
  ids                                   = ["${oci_waas_http_redirect.test_http_redirect.id}"]
  states                                = ["${oci_waas_http_redirect.test_http_redirect.state}"]
  time_created_greater_than_or_equal_to = "2018-01-01T00:00:00.000Z"
  time_created_less_than                = "2038-01-01T00:00:00.000Z"
}
