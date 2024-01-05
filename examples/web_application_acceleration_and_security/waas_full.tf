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

variable "certificate_display_name" {
  default = "tf_example_waas_certificate"
}

variable "waas_policy_display_name" {
  default = "tf_example_waas_policy_test"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_waas_certificate" "test_certificate" {
  #Required
  certificate_data = "-----BEGIN CERTIFICATE-----\nMIICljCCAX4CCQCEpaMjTCJ8WzANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJV\nUzAeFw0yMTAxMTkyMTI2MjRaFw0yNDAxMTkyMTI2MjRaMA0xCzAJBgNVBAYTAlVT\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo83kaUQXpCcSoEuRVFX3\njztWDNKtWpjNG240f0RpERI1NnZtHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19g\neIXR6TeavT+W5iRh4goK+N7gubYkSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCt\nCgd4MDlsvLv/YHCLvJL4JgRxKyevjlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc\n+Je9EC3MWWxd5jBwXu3vgIYRuGR4DPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJ\nKN0NXp5obaQToYqMsvAZyHoEyfCBDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t5\n8QIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAX1rxV2hai02Pb4Cf8U44zj+1aY6wV\nLvOMWiL3zl53up4/X7PDcmWcPM9UMVCGTISZD6A6IPvNlkvbtvYCzgjhtGxDmrj7\nwTRV5gO9j3bAhxBO7XgTmwmD/9hpykM58nbhLFnkGf+Taja8qsy0U8H74Tr9w1M8\n8E5kghgGzBElNquM8AUuDakC1JL4aLO/VDMxe/1BLtmBHLZy3XTzVycjP9ZFPh6h\nT+cWJcVOjQSYY2U75sDnKD2Sg1cmK54HauA6SPh4kAkpmxyLyDZZjPBQe2sLFmmS\naZSE+g16yMR9TVHo3pTpRkxJwDEH0LePwYXA4vUIK3HHS6zgLe0ody8g\n-----END CERTIFICATE-----"
  compartment_id   = var.compartment_ocid
  private_key_data = "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAo83kaUQXpCcSoEuRVFX3jztWDNKtWpjNG240f0RpERI1NnZt\nHH0qnZqfaWAQQa8kx3+W1LOeFbkkRnkJz19geIXR6TeavT+W5iRh4goK+N7gubYk\nSMa2shVf+XsoHKERSbhdhrtX+GqvKzAvplCtCgd4MDlsvLv/YHCLvJL4JgRxKyev\njlnE1rqrICJMCLbbZMrIKTzwb/K13hGrm6Bc+Je9EC3MWWxd5jBwXu3vgIYRuGR4\nDPg/yfMKPZr2xFDLpBsv5jaqULS9t6GwoEBJKN0NXp5obaQToYqMsvAZyHoEyfCB\nDka16Bm5hGF60FwqgUT3p/+qlBn61cAJe9t58QIDAQABAoIBADIyHuOPJTt9abzL\nS26vpVw0D6uAR/UyS/Ay9k1ltliv3rSg19DaHlwLjPwqnvCx7jBgTeVCYZhAkvgx\nkSsGDDcCsw+npXiG6wP9dC1jbHdVPUJLqZTPqB6sZCu8bM9RIE4Z/DcUY+HRN3qh\nmoh5wn0HSvJkNokjhx+TfY687uQfDMu0de4V2UPScZ7mboCu9HqK9qu0/krdTMH1\nrtnnFGEnx/Pe38YJl0fWxo8BHKHprwEvWX0MQzQeklnUtxREMuofSAOBe/I2DJGh\n1I94b6I66ypxuX0qAozT1MPbJGuaR+puyKawLNAQmZa9pgrrFK7e8PQUzrGVpVCp\nFtwx420CgYEA0uX/G0ycia0UTdkxkIsKIiLjs12LC0XmYjiWgkoL0PjiZzcPITn6\nvqqqGSz44HwtbrttZPm3Mo79yJ5xFiHCX0vFJykgy6cfS94imMgm8qIOS0bXjX7w\nxH2BOgp0H32LP/Zt7owcWJLEIQCjj0/4+Nvu0GskGVHlE8EYrXWf1E8CgYEAxtWk\nxBo52uNXL712VGDPNxprVGUpWSbkXpE+7wtRso5LnAnAj6dpmZsGe2zaYjRIt3Ls\nGnno4HUmwpQ5yXlHFpDUJvb2soXq3afnuAh5aVu6aKRQoG/5o3cD4pOupNbjDDNs\nTVLtTLIAIYDbph/j7pV/JnJ2WHcdk6WiVJoW/b8CgYAopLZzNmJ8jeR51D+fEYyU\nY5DqQj7Hn2L0zt8CoO6CCVToe03pI1lVYWKCk44rBQNkca51ZUKO9cum3BIDJ+Jj\npyCJmX1+geigIGEefIQ1AlIq464q0Knp1B4RZ25Vm0Y4v28UJ+BWmYI+sfbTaaAb\npZbyh5NfZc717aKp2x9ANQKBgHQpvOkUqVhIGVe6yLbjGCyJMstLjqyXHDRjhvEB\nG+nFWEcBK47Br+Adwdu57JwTD6ida3LMZlE8IDjtgBVE1VNJqahaACasNlrpDWdn\nDAeRn4Yi+TfCM4Zcsdhdj1qecGdgY5WJLTnxhEIOlkSnvPJWRMKhfKKSdKUdz4i9\nvVDhAoGAEHxfhFLVwdTa0RMUq3KYSXa5WqLANRn2e62Cc3eUWekcUjbIATRF5AIo\nm0WS+rURZWy1Fd6fGg8sRHock0+vxwqeP6OlyW4tJMhL33NrNbgyvkXlMMIX6bC4\nUq8aAew0B3j61UUsTqhHMhYwIS3GOIHx/O10wwINPnUMIVER3Wg=\n-----END RSA PRIVATE KEY-----"

  #Optional
  display_name                   = var.certificate_display_name
  is_trust_verification_disabled = true
}

data "oci_waas_certificates" "test_certificates" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_names = [var.certificate_display_name]
  ids           = [oci_waas_certificate.test_certificate.id]
}

resource "oci_waas_custom_protection_rule" "test_custom_protection_rule" {
  compartment_id = var.compartment_ocid
  display_name   = "tf_example_protection_rule"
  template       = "SecRule REQUEST_URI / \"phase:2,   t:none,   capture,   msg:'Custom (XSS) Attack. Matched Data: %%%{TX.0}   found within %%%{MATCHED_VAR_NAME}: %%%{MATCHED_VAR}',   id:{{id_1}},   ctl:ruleEngine={{mode}},   tag:'Custom',   severity:'2'\""

  #Optional
  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags = {"example-tag-namespace-all.example-tag", "originalValue"}
  description = "Tf example custom protection rule"

  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_waas_address_list" "test_address_list" {
  #Required
  addresses      = ["0.0.0.0/16", "192.168.0.0/20"]
  compartment_id = var.compartment_ocid
  display_name   = "tf-example-address-list"

  #To use defined_tags, set the values below to an existing tag namespace, refer to the identity example on how to create tag namespaces
  #defined_tags = {"example-tag-namespace-all.example-tag", "originalValue"}
  freeform_tags = {
    "Department" = "Finance"
  }
}

resource "oci_waas_waas_policy" "test_waas_policy" {
  #Required
  compartment_id = var.compartment_ocid
  domain         = "somethingnew66.oracle.com"

  #Optional
  additional_domains = ["somethingnew67.oracle.com", "somethingnew68.oracle.com"]
  display_name       = var.waas_policy_display_name

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
    certificate_id                = oci_waas_certificate.test_certificate.id
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
    tls_protocols                 = ["TLS_V1_3"]

    health_checks {
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

    load_balancing_method {
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

      response_header_manipulation {
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
      id     = oci_waas_custom_protection_rule.test_custom_protection_rule.id

      exclusions {
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

      criteria {
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
      address_lists = [oci_waas_address_list.test_address_list.id]
    }
  }
}

resource "oci_waas_protection_rule" "test_waas_protection_rule" {
  waas_policy_id = oci_waas_waas_policy.test_waas_policy.id
  key            = "933161"
  action         = "DETECT"

  exclusions {
    exclusions = ["example.com"]
    target     = "REQUEST_COOKIES"
  }
}

data "oci_waas_protection_rules" "test_waas_protection_rules" {
  waas_policy_id = oci_waas_waas_policy.test_waas_policy.id
  action         = ["DETECT"]
}

data "oci_waas_protection_rule" "test_waas_protection_rule" {
  waas_policy_id      = oci_waas_waas_policy.test_waas_policy.id
  protection_rule_key = oci_waas_protection_rule.test_waas_protection_rule.key
}

data "oci_waas_waas_policies" "test_waas_policies" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_names                         = [var.waas_policy_display_name]
  ids                                   = [oci_waas_waas_policy.test_waas_policy.id]
  states                                = ["ACTIVE"]
  time_created_greater_than_or_equal_to = "2018-01-01T00:00:00.000Z"
  time_created_less_than                = "2038-01-01T00:00:00.000Z"
}

data "oci_waas_address_lists" "test_address_lists" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  ids                                   = [oci_waas_address_list.test_address_list.id]
  names                                 = ["tf-example-address-list"]
  states                                = ["ACTIVE"]
  time_created_greater_than_or_equal_to = "2018-01-01T00:00:00.000Z"
  time_created_less_than                = "2038-01-01T00:00:00.000Z"
}

data "oci_waas_custom_protection_rules" "test_custom_protection_rules" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_names                         = ["tf_example_protection_rule"]
  ids                                   = [oci_waas_custom_protection_rule.test_custom_protection_rule.id]
  states                                = ["ACTIVE"]
  time_created_greater_than_or_equal_to = "2018-01-01T00:00:00.000Z"
  time_created_less_than                = "2038-01-01T00:00:00.000Z"
}

resource "oci_waas_http_redirect" "test_http_redirect" {
  #Required
  compartment_id = var.compartment_ocid
  domain         = "example5.com"

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
  compartment_id = var.compartment_ocid

  #Optional
  display_names                         = [oci_waas_http_redirect.test_http_redirect.display_name]
  ids                                   = [oci_waas_http_redirect.test_http_redirect.id]
  states                                = [oci_waas_http_redirect.test_http_redirect.state]
  time_created_greater_than_or_equal_to = "2018-01-01T00:00:00.000Z"
  time_created_less_than                = "2038-01-01T00:00:00.000Z"
}
