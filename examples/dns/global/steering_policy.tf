// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * Provider config for dns sample
 */

variable "http_monitor_display_name" {
  default = "displayName"
}

variable "http_monitor_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "http_monitor_headers" {
  default = "headers"
}

variable "http_monitor_interval_in_seconds" {
  default = 10
}

variable "http_monitor_is_enabled" {
  default = true
}

variable "http_monitor_method" {
  default = "GET"
}

variable "http_monitor_path" {
  default = "/"
}

variable "http_monitor_port" {
  default = "443"
}

variable "http_monitor_protocol" {
  default = "HTTPS"
}

variable "http_monitor_targets" {
  default = ["www.oracle.com"]
}

variable "http_monitor_timeout_in_seconds" {
  default = 10
}

variable "http_monitor_vantage_point_names" {
  default = ["goo-chs"]
}

variable "steering_policy_answers_is_disabled" {
  default = false
}

variable "steering_policy_answers_name" {
  default = "name"
}

variable "steering_policy_answers_pool" {
  default = "pool"
}

variable "steering_policy_answers_rdata" {
  default = "192.0.2.1"
}

variable "steering_policy_answers_rtype" {
  default = "A"
}

variable "steering_policy_display_name" {
  default = "displayName"
}

variable "steering_policy_display_name_contains" {
  default = "displayNameContains"
}

variable "steering_policy_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}

variable "steering_policy_id" {
  default = "id"
}

variable "steering_policy_rules_cases_answer_data_answer_condition" {
  default = "answer.name == 'sampler'"
}

variable "steering_policy_rules_cases_answer_data_should_keep" {
  default = false
}

variable "steering_policy_rules_cases_answer_data_value" {
  default = 10
}

variable "steering_policy_rules_cases_case_condition" {
  default = "query.client.address in (subnet '198.51.100.0/24')"
}

variable "steering_policy_rules_cases_count" {
  default = 10
}

variable "steering_policy_rules_default_answer_data_answer_condition" {
  default = "answer.name == 'sampler'"
}

variable "steering_policy_rules_default_answer_data_should_keep" {
  default = false
}

variable "steering_policy_rules_default_answer_data_value" {
  default = 10
}

variable "steering_policy_rules_default_count" {
  default = 10
}

variable "steering_policy_rules_rule_type" {
  default = "PRIORITY"
}

variable "steering_policy_state" {
  default = "ACTIVE"
}

variable "steering_policy_template" {
  default = "CUSTOM"
}

variable "steering_policy_time_created_greater_than_or_equal_to" {
  default = "2018-01-01T00:00:00.000Z"
}

variable "steering_policy_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}

variable "steering_policy_ttl" {
  default = 10
}

resource "oci_health_checks_http_monitor" "test_http_monitor" {
  #Required
  compartment_id      = var.compartment_ocid
  display_name        = var.http_monitor_display_name
  interval_in_seconds = var.http_monitor_interval_in_seconds
  protocol            = var.http_monitor_protocol
  targets             = var.http_monitor_targets

  #Optional
  freeform_tags       = var.http_monitor_freeform_tags
  is_enabled          = var.http_monitor_is_enabled
  method              = var.http_monitor_method
  path                = var.http_monitor_path
  port                = var.http_monitor_port
  timeout_in_seconds  = var.http_monitor_timeout_in_seconds
  vantage_point_names = var.http_monitor_vantage_point_names
}

resource "oci_dns_steering_policy" "test_steering_policy" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.steering_policy_display_name
  template       = var.steering_policy_template

  #Optional
  answers {
    #Required
    name  = var.steering_policy_answers_name
    rdata = var.steering_policy_answers_rdata
    rtype = var.steering_policy_answers_rtype

    #Optional
    is_disabled = var.steering_policy_answers_is_disabled
    pool        = var.steering_policy_answers_pool
  }

  freeform_tags           = var.steering_policy_freeform_tags
  health_check_monitor_id = oci_health_checks_http_monitor.test_http_monitor.id

  rules {
    #Required
    rule_type = "FILTER"

    #Optional
    cases {
      #Optional
      answer_data {
        #Optional
        answer_condition = var.steering_policy_rules_cases_answer_data_answer_condition
        should_keep      = var.steering_policy_rules_cases_answer_data_should_keep
      }

      case_condition = var.steering_policy_rules_cases_case_condition
    }

    default_answer_data {
      #Optional
      answer_condition = var.steering_policy_rules_default_answer_data_answer_condition
      should_keep      = var.steering_policy_rules_default_answer_data_should_keep
    }
  }

  rules {
    #Required
    rule_type = "HEALTH"

    #Optional
    cases {
      #Optional
      case_condition = var.steering_policy_rules_cases_case_condition
    }
  }

  rules {
    #Required
    rule_type = "LIMIT"

    #Optional
    cases {
      case_condition = var.steering_policy_rules_cases_case_condition
      count          = var.steering_policy_rules_cases_count
    }

    default_count = var.steering_policy_rules_default_count
  }

  rules {
    #Required
    rule_type = "PRIORITY"

    #Optional
    cases {
      #Optional
      answer_data {
        #Optional
        answer_condition = var.steering_policy_rules_cases_answer_data_answer_condition
        value            = var.steering_policy_rules_cases_answer_data_value
      }

      case_condition = var.steering_policy_rules_cases_case_condition
    }

    default_answer_data {
      #Optional
      answer_condition = var.steering_policy_rules_default_answer_data_answer_condition
      value            = var.steering_policy_rules_default_answer_data_value
    }
  }

  rules {
    #Required
    rule_type = var.steering_policy_rules_rule_type

    #Optional
    cases {
      #Optional
      answer_data {
        #Optional
        answer_condition = var.steering_policy_rules_cases_answer_data_answer_condition
        value            = var.steering_policy_rules_cases_answer_data_value
      }

      case_condition = var.steering_policy_rules_cases_case_condition
    }

    default_answer_data {
      #Optional
      answer_condition = var.steering_policy_rules_default_answer_data_answer_condition
      value            = var.steering_policy_rules_default_answer_data_value
    }
  }

  ttl = var.steering_policy_ttl
}

data "oci_dns_steering_policies" "test_steering_policies" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.steering_policy_display_name

  #display_name_contains                 = var.steering_policy_display_name_contains
  health_check_monitor_id               = oci_health_checks_http_monitor.test_http_monitor.id
  id                                    = oci_dns_steering_policy.test_steering_policy.id
  state                                 = var.steering_policy_state
  template                              = var.steering_policy_template
  time_created_greater_than_or_equal_to = var.steering_policy_time_created_greater_than_or_equal_to
  time_created_less_than                = var.steering_policy_time_created_less_than
}

