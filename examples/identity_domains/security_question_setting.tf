// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "security_question_setting_authorization" {
  default = "authorization"
}

variable "security_question_setting_id" {
  default = "SecurityQuestionSettings"
}

variable "security_question_setting_max_field_length" {
  default = 10
}

variable "security_question_setting_meta_last_modified" {
  default = "lastModified"
}

variable "security_question_setting_min_answer_length" {
  default = "6"
}

variable "security_question_setting_num_questions_to_ans" {
  default = "2"
}

variable "security_question_setting_num_questions_to_setup" {
  default = "5"
}

variable "security_question_setting_tags_key" {
  default = "key"
}

variable "security_question_setting_tags_value" {
  default = "value"
}

resource "oci_identity_domains_security_question_setting" "test_security_question_setting" {
  #Required
  idcs_endpoint                = data.oci_identity_domain.test_domain.url
  max_field_length             = var.security_question_setting_max_field_length
  min_answer_length            = var.security_question_setting_min_answer_length
  num_questions_to_ans         = var.security_question_setting_num_questions_to_ans
  num_questions_to_setup       = var.security_question_setting_num_questions_to_setup
  schemas                      = ["urn:ietf:params:scim:schemas:oracle:idcs:SecurityQuestionSettings"]
  security_question_setting_id = var.security_question_setting_id

  #Optional
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.security_question_setting_authorization
  external_id                  = "externalId"
  tags {
    #Required
    key   = var.security_question_setting_tags_key
    value = var.security_question_setting_tags_value
  }
}

data "oci_identity_domains_security_question_settings" "test_security_question_settings" {
  #Required
  idcs_endpoint = data.oci_identity_domain.test_domain.url

  #Optional
  attribute_sets               = ["all"]
  attributes                   = ""
  authorization                = var.security_question_setting_authorization
}

