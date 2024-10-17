
variable "compartment_id" {}

variable "defined_tags_value" {
  default = "defined_tags_value"
}

resource "oci_logging_log_group" "test_log_group" {
  compartment_id = var.compartment_id
  display_name = "logRuleTFExampleGroup"
}

resource "oci_logging_log" "test_log" {
  display_name = "logRuleTFExampleLog"
  log_group_id = oci_logging_log_group.test_log_group.id
  log_type = "CUSTOM"
}

resource "oci_logging_log_rule" "test_log_rule" {
  compartment_id = var.compartment_id
  custom_log_id = oci_logging_log.test_log.id
  description = "description2"
  display_name = "logRuleTFExampleName"
  freeform_tags = {
    "Department" = "Accounting"
  }
  log_rule_status = "DISABLED"
  operator = "GREATER"
  query = "search \"${var.compartment_id}\""
  query_recurrences = "FREQ=MINUTELY;INTERVAL=20"
  query_start_policy {
    start_policy_type = "ABSOLUTE_TIME_START_POLICY"
    time_to_start_query = "2207-01-02T15:04:05Z"
  }
  recommendation_text = "recommendationText2"
  severity = "HIGH"
  threshold = "10"
}
data "oci_logging_log_rules" "test_log_rules" {
  compartment_id = var.compartment_id
  display_name = "logRuleTFExampleName"
  filter {
    name = "id"
    values = [oci_logging_log_rule.test_log_rule.id]
  }
  frequency = "20"
  severity = "HIGH"
  state = "INACTIVE"
}