// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

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

resource "oci_waas_waas_policy" "test_waas_policy" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  domain         = "somethingnew42.oracle.com"

  #Optional
  additional_domains = ["somethingnew1.oracle.com", "somethingnew2.oracle.com"]
  display_name       = "${var.waas_policy_display_name}"

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
    certificate_id   = "${oci_waas_certificate.test_certificate.id}"
    is_https_enabled = true
    is_https_forced  = true
  }

  timeouts {
    create = "60m"
    delete = "60m"
    update = "60m"
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
    }

    address_rate_limiting {
      #Required
      is_enabled = true

      #Optional
      allowed_rate_per_address      = 10
      block_response_code           = 403
      max_delayed_count_per_address = 10
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

    origin = "primary"

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
      addresses = ["192.168.127.127", "192.168.127.128"]
      name      = "whitelist_name"
    }
  }
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
