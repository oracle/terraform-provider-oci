// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "compartment_ocid" {}

variable "private_key_data" {}

variable "dns_secret" {}

variable "function_image" {}

variable "function_image_digest" {}

variable "instance_image_ocid" {
  type = "map"

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"

    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

variable "db_edition" {
  default = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
}

variable "db_admin_password" {}

variable "db_version" {
  default = "12.1.0.2"
}

variable "db_system_shape" {
  default = "Exadata.Quarter1.84"
}

variable "n_character_set" {
  default = "AL16UTF16"
}

variable "character_set" {
  default = "AL32UTF8"
}

variable "db_workload" {
  default = "OLTP"
}

/* apigateway */
variable "gateway_display_name" {
  default = "apigatewayGatewayRd"
}

variable "gateway_endpoint_type" {
  default = "PUBLIC"
}

variable "gateway_state" {
  default = "ACTIVE"
}

variable "deployment_defined_tags_value" {
  default = "value"
}

variable "deployment_display_name" {
  default = "apigatewayDeploymentRd"
}

variable "deployment_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "application_state" {
  default = "AVAILABLE"
}

variable "config" {
  default = {
    "MY_FUNCTION_CONFIG" = "ConfVal"
  }
}

variable "function_memory_in_mbs" {
  default = 128
}

variable "function_timeout_in_seconds" {
  default = 30
}

variable "deployment_path_prefix" {
  default = "/v1"
}

variable "deployment_specification_logging_policies_access_log_is_enabled" {
  default = false
}

variable "deployment_specification_logging_policies_execution_log_is_enabled" {
  default = false
}

variable "deployment_specification_logging_policies_execution_log_log_level" {
  default = "INFO"
}

variable "deployment_specification_request_policies_authentication_is_anonymous_access_allowed" {
  default = false
}

variable "deployment_specification_request_policies_authentication_token_header" {
  default = "Authorization"
}

variable "deployment_specification_request_policies_authentication_token_query_param" {
  default = "tokenQueryParam"
}

variable "deployment_specification_request_policies_authentication_type" {
  default = "CUSTOM_AUTHENTICATION"
}

variable "deployment_specification_request_policies_cors_allowed_headers" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_allowed_methods" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_allowed_origins" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_exposed_headers" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_is_allow_credentials_enabled" {
  default = false
}

variable "deployment_specification_request_policies_cors_max_age_in_seconds" {
  default = "600"
}

variable "deployment_specification_request_policies_rate_limiting_rate_in_requests_per_second" {
  default = 10
}

variable "deployment_specification_request_policies_rate_limiting_rate_key" {
  default = "CLIENT_IP"
}

variable "deployment_specification_routes_backend_body" {
  default = "body"
}

variable "deployment_specification_routes_backend_connect_timeout_in_seconds" {
  default = 1.0
}

variable "deployment_specification_routes_backend_headers_name" {
  default = "name"
}

variable "deployment_specification_routes_backend_headers_value" {
  default = "value"
}

variable "deployment_specification_routes_backend_is_ssl_verify_disabled" {
  default = false
}

variable "deployment_specification_routes_backend_read_timeout_in_seconds" {
  default = 1.0
}

variable "deployment_specification_routes_backend_send_timeout_in_seconds" {
  default = 1.0
}

variable "deployment_specification_routes_backend_status" {
  default = 10
}

variable "deployment_specification_routes_backend_type" {
  default = "HTTP_BACKEND"
}

variable "deployment_specification_routes_backend_url" {
  default = "https://api.weather.gov"
}

variable "deployment_specification_routes_logging_policies_access_log_is_enabled" {
  default = false
}

variable "deployment_specification_routes_logging_policies_execution_log_is_enabled" {
  default = false
}

variable "deployment_specification_routes_logging_policies_execution_log_log_level" {
  default = "INFO"
}

variable "deployment_specification_routes_methods" {
  default = ["GET"]
}

variable "deployment_specification_routes_path" {
  default = "/hello"
}

variable "deployment_specification_routes_request_policies_authorization_allowed_scope" {
  default = []
}

variable "deployment_specification_routes_request_policies_authorization_type" {
  default = "AUTHENTICATION_ONLY"
}

variable "deployment_specification_routes_request_policies_cors_allowed_headers" {
  default = ["*"]
}

variable "deployment_specification_routes_request_policies_cors_allowed_methods" {
  default = ["GET"]
}

variable "deployment_specification_routes_request_policies_cors_allowed_origins" {
  default = ["*"]
}

variable "deployment_specification_routes_request_policies_cors_exposed_headers" {
  default = ["*"]
}

variable "deployment_specification_routes_request_policies_cors_is_allow_credentials_enabled" {
  default = false
}

variable "deployment_specification_routes_request_policies_cors_max_age_in_seconds" {
  default = "600"
}

variable "deployment_state" {
  default = "ACTIVE"
}

/* email */
variable "sender_email_address" {
  default = "JohnSmith@example.com"
}

variable "suppression_email_address" {
  default = "JohnSmith@example.com"
}

/* health_checks */
variable "http_monitor_defined_tags_value" {
  default = "value"
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

variable "ping_monitor_defined_tags_value" {
  default = "value"
}

variable "ping_monitor_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "ping_monitor_interval_in_seconds" {
  default = 10
}

variable "ping_monitor_is_enabled" {
  default = false
}

variable "ping_monitor_port" {
  default = 80
}

variable "ping_monitor_protocol" {
  default = "TCP"
}

variable "ping_monitor_targets" {
  default = ["www.oracle.com"]
}

variable "ping_monitor_timeout_in_seconds" {
  default = 10
}

variable "ping_monitor_vantage_point_names" {
  default = ["goo-chs"]
}

/* Nosql */

variable "table_ddl_statement" {
  default = "CREATE TABLE IF NOT EXISTS test_table(id INTEGER, name STRING, age STRING, info JSON, PRIMARY KEY(SHARD(id)))"
}

variable "index_keys_column_name" {
  default = "name"
}

/* Monitoring */

variable "alarm_body" {
  default = "High CPU utilization reached"
}

variable "alarm_compartment_id_in_subtree" {
  default = false
}

variable "alarm_defined_tags_value" {
  default = "value"
}

variable "alarm_destinations" {
  default = []
}

variable "alarm_display_name" {
  default = "High CPU Utilization"
}

variable "alarm_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "alarm_is_enabled" {
  default = false
}

variable "alarm_metric_compartment_id_in_subtree" {
  default = false
}

variable "alarm_namespace" {
  default = "oci_computeagent"
}

variable "alarm_pending_duration" {
  default = "PT5M"
}

variable "alarm_query" {
  default = "CpuUtilization[10m].percentile(0.9) < 85"
}

variable "alarm_repeat_notification_duration" {
  default = "PT2H"
}

variable "alarm_resolution" {
  default = "1m"
}
variable "is_notifications_per_metric_dimension_enabled"{
  default = false
}

variable "alarm_resource_group" {
  default = "resourceGroup"
}

variable "alarm_severity" {
  default = "WARNING"
}

variable "alarm_state" {
  default = "ACTIVE"
}

variable "alarm_suppression_description" {
  default = "System Maintenance"
}

variable "alarm_suppression_time_suppress_from" {
  default = "2029-02-01T18:00:00.000Z"
}

variable "alarm_suppression_time_suppress_until" {
  default = "2029-02-01T19:00:00.000Z"
}

/* Waas */

variable "certificate_display_name" {
  default = "tf_example_waas_certificate_rd"
}

variable "waas_policy_display_name" {
  default = "tf_example_waas_policy_rd"
}

/* dataflow */

variable "application_driver_shape" {
  default = "VM.Standard2.1"
}

variable "application_executor_shape" {
  default = "VM.Standard2.1"
}

variable "application_language" {
  default = "PYTHON"
}

variable "application_num_executors" {
  default = 1
}

variable "application_spark_version" {
  default = "2.4.4"
}

variable "application_file_uri" {
  default = "oci://StreamingArchiverTestBucket@dxterraformdev/dataflowTestFile.py"
}

variable "application_archive_uri" {
  default = "oci://StreamingArchiverTestBucket@dxterraformdev/dataflowTestFile.py.zip"
}

variable "application_warehouse_bucket_uri" {
  default = "oci://dataflow-logs@dxterraformdev/"
}

variable "application_logs_bucket_uri" {
  default = "oci://dataflow-logs@dxterraformdev/"
}

/* dns */

variable "http_monitor_display_name_dns_rd" {
  default = "httpMonitorDisplayNameDnsRD"
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
  default = "dnsSteeringPolicyRD"
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

variable "dns_tsig_key_name" {
  default = "test_tsig_key-name_rd"
}

/* dns attachment */
variable "steering_policy_attachment_display_name" {
  default = "Test-Steering-Policy-Attachment"
}

variable "steering_policy_attachment_state" {
  default = "ACTIVE"
}

variable "steering_policy_attachment_time_created_greater_than_or_equal_to" {
  default = "2018-01-01T00:00:00.000Z"
}

variable "steering_policy_attachment_time_created_less_than" {
  default = "2038-01-01T00:00:00.000Z"
}

/* waas */
variable "waas_policy_domain" {
  default = "testdomainforrd.oracle.com"
}

variable "waas_http_redirect_domain" {
  default = "testdomainforrd3.oracle.com"
}

variable "waas_http_redirect_host" {
  default = "testdomainforrd2.oracle.com"
}

/* tagging */
variable "tag_namespace_name" {
  default = "tagNamespaceRD"
}
