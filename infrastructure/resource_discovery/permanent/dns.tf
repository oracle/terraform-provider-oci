// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_health_checks_http_monitor" "test_http_monitor_dns_rd" {
  #Required
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "${var.http_monitor_display_name_dns_rd}"
  interval_in_seconds = "${var.http_monitor_interval_in_seconds}"
  protocol            = "${var.http_monitor_protocol}"
  targets             = "${var.http_monitor_targets}"

  #Optional
  freeform_tags       = "${var.http_monitor_freeform_tags}"
  is_enabled          = "${var.http_monitor_is_enabled}"
  method              = "${var.http_monitor_method}"
  path                = "${var.http_monitor_path}"
  port                = "${var.http_monitor_port}"
  timeout_in_seconds  = "${var.http_monitor_timeout_in_seconds}"
  vantage_point_names = "${var.http_monitor_vantage_point_names}"
}

resource "oci_dns_steering_policy" "test_dns_steering_policy_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  display_name   = "${var.steering_policy_display_name}"
  template       = "${var.steering_policy_template}"

  #Optional
  answers {
    #Required
    name  = "${var.steering_policy_answers_name}"
    rdata = "${var.steering_policy_answers_rdata}"
    rtype = "${var.steering_policy_answers_rtype}"

    #Optional
    is_disabled = "${var.steering_policy_answers_is_disabled}"
    pool        = "${var.steering_policy_answers_pool}"
  }

  freeform_tags           = "${var.steering_policy_freeform_tags}"
  health_check_monitor_id = "${oci_health_checks_http_monitor.test_http_monitor_dns_rd.id}"

  rules {
    #Required
    rule_type = "FILTER"

    #Optional
    cases {
      #Optional
      answer_data {
        #Optional
        answer_condition = "${var.steering_policy_rules_cases_answer_data_answer_condition}"
        should_keep      = "${var.steering_policy_rules_cases_answer_data_should_keep}"
      }

      case_condition = "${var.steering_policy_rules_cases_case_condition}"
    }

    default_answer_data {
      #Optional
      answer_condition = "${var.steering_policy_rules_default_answer_data_answer_condition}"
      should_keep      = "${var.steering_policy_rules_default_answer_data_should_keep}"
    }
  }

  rules {
    #Required
    rule_type = "HEALTH"

    #Optional
    cases {
      #Optional
      case_condition = "${var.steering_policy_rules_cases_case_condition}"
    }
  }

  rules {
    #Required
    rule_type = "LIMIT"

    #Optional
    cases {
      case_condition = "${var.steering_policy_rules_cases_case_condition}"
      count          = "${var.steering_policy_rules_cases_count}"
    }

    default_count = "${var.steering_policy_rules_default_count}"
  }

  rules {
    #Required
    rule_type = "PRIORITY"

    #Optional
    cases {
      #Optional
      answer_data {
        #Optional
        answer_condition = "${var.steering_policy_rules_cases_answer_data_answer_condition}"
        value            = "${var.steering_policy_rules_cases_answer_data_value}"
      }

      case_condition = "${var.steering_policy_rules_cases_case_condition}"
    }

    default_answer_data {
      #Optional
      answer_condition = "${var.steering_policy_rules_default_answer_data_answer_condition}"
      value            = "${var.steering_policy_rules_default_answer_data_value}"
    }
  }

  rules {
    #Required
    rule_type = "${var.steering_policy_rules_rule_type}"

    #Optional
    cases {
      #Optional
      answer_data {
        #Optional
        answer_condition = "${var.steering_policy_rules_cases_answer_data_answer_condition}"
        value            = "${var.steering_policy_rules_cases_answer_data_value}"
      }

      case_condition = "${var.steering_policy_rules_cases_case_condition}"
    }

    default_answer_data {
      #Optional
      answer_condition = "${var.steering_policy_rules_default_answer_data_answer_condition}"
      value            = "${var.steering_policy_rules_default_answer_data_value}"
    }
  }

  ttl = "${var.steering_policy_ttl}"
}

resource "oci_dns_steering_policy_attachment" "test_steering_policy_attachment" {
  #Required
  domain_name        = "${oci_dns_record.record-a.domain}"
  steering_policy_id = "${oci_dns_steering_policy.test_dns_steering_policy_rd.id}"
  zone_id            = "${oci_dns_zone.zone1.id}"

  #Optional
  display_name = "${var.steering_policy_attachment_display_name}"
}

data "oci_dns_steering_policy_attachments" "test_steering_policy_attachments" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "${oci_dns_steering_policy_attachment.test_steering_policy_attachment.display_name}"
  domain       = "${oci_dns_steering_policy_attachment.test_steering_policy_attachment.domain_name}"

  #domain_contains                       = "${oci_dns_steering_policy_attachment.test_steering_policy_attachment.domain_name}"
  id                                    = "${oci_dns_steering_policy_attachment.test_steering_policy_attachment.id}"
  state                                 = "${var.steering_policy_attachment_state}"
  steering_policy_id                    = "${oci_dns_steering_policy.test_dns_steering_policy_rd.id}"
  time_created_greater_than_or_equal_to = "${var.steering_policy_attachment_time_created_greater_than_or_equal_to}"
  time_created_less_than                = "${var.steering_policy_attachment_time_created_less_than}"
  zone_id                               = "${oci_dns_zone.zone2.id}"
}

resource "oci_dns_record" "record-a" {
  zone_name_or_id = "${oci_dns_zone.zone1.name}"
  domain          = "${oci_dns_zone.zone1.name}"
  rtype           = "A"
  rdata           = "192.168.0.1"
  ttl             = 3600
}

resource "oci_dns_zone" "zone1" {
  compartment_id = "${var.compartment_ocid}"
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example-primary.oci-dns1"
  zone_type      = "PRIMARY"
}

resource "oci_dns_zone" "zone2" {
  compartment_id = "${var.compartment_ocid}"
  name           = "${data.oci_identity_tenancy.tenancy.name}-${random_string.random_prefix.result}-tf-example-secondary.oci-dns2"
  zone_type      = "SECONDARY"

  external_masters {
    address     = "77.64.12.1"
    tsig_key_id = "${oci_dns_tsig_key.test_tsig_key.id}"
  }

  external_masters {
    address     = "77.64.12.2"
    tsig_key_id = "${oci_dns_tsig_key.test_tsig_key.id}"
  }
}

data "oci_identity_tenancy" "tenancy" {
  tenancy_id = "${var.tenancy_ocid}"
}

resource "oci_dns_tsig_key" "test_tsig_key" {
  algorithm      = "hmac-sha1"
  compartment_id = "${var.compartment_ocid}"
  name           = "${var.dns_tsig_key_name}"
  secret         = "${var.dns_secret}"
}

resource "random_string" "random_prefix" {
  length  = 4
  number  = false
  special = false
}
