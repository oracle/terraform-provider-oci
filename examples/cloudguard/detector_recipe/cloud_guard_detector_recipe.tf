// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "detector_recipe_access_level" {
  default = "ACCESSIBLE"
}

variable "detector_recipe_compartment_id_in_subtree" {
  default = true
}
//Refer to the note in managed_list.tf for the above two variables

variable "detector_recipe_defined_tags_value" {
  default = "value"
}

variable "detector_recipe_description" {
  default = "description"
}

/*
The configuration and condition Objects are dependent on the rule id.
Hence for testing purposes we are going to hardcode a rule id and
corresponding valid condition and configuration
*/
variable "detector_recipe_detector_rules_details_condition" {
  default = "{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"10\",\"valueType\":\"CUSTOM\"}"
}

variable "detector_recipe_detector_rules_details_configurations_config_key" {
  default = "lbCertificateExpiringSoonConfig"
}

variable "detector_recipe_detector_rules_details_configurations_data_type" {
  default = "int"
}

variable "detector_recipe_detector_rules_details_configurations_name" {
  default = "Days before expiring"
}

variable "detector_recipe_detector_rules_details_configurations_value" {
  default = "30"
}

variable "detector_recipe_detector_rules_details_is_enabled" {
  default = true
}

variable "detector_recipe_detector_rules_details_labels" {
  default = []
}

//Acceptable values come from RiskLevelEnum
variable "detector_recipe_detector_rules_details_risk_level" {
  default = "HIGH"
}

//Has to be unique
variable "detector_recipe_display_name" {
  default = "displayName"
}

variable "detector_recipe_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}

//Acceptable values come from LifecycleStateEnum
variable "detector_recipe_state" {
  default = "ACTIVE"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  //version             = "5.39.0"
  /*tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"*/
}

data "oci_cloud_guard_detector_recipes" "test_detector_recipes" {
  #Required
  compartment_id = "${var.tenancy_ocid}"

  #Optional
  state        = "${var.detector_recipe_state}"
  # Adding this to make sure the detector_rule_id we use later on is valid against the returned recipes
  display_name = "OCI Configuration Detector Recipe"
}

resource "oci_cloud_guard_detector_recipe" "test_detector_recipe" {
  #Required
  compartment_id            = "${var.compartment_id}"
  display_name              = "${var.detector_recipe_display_name}"
  /*
 When CloudGuard is Enabled, an Oracle Managed Detector Recipe is made available having all the default rules in their default state.
 If an user needs to make its own recipe with customizations to the rules, it needs to clone an `ORACLE MANAGED DETECTOR RECIPE`.
 In order to clone an `ORACLE MANAGED DETECTOR RECIPE` (Create new Detector Recipe) the source responder recipe id needs to be the ocid of
 `ORACLE MANAGED DETECTOR RECIPE`.
 In order to achieve that, we will list all the detector recipes and then pick the first item (Oracle Managed Detector Recipe) in the collection for cloning.

  However if user wants to create and manage the ORACLE MANAGED Entity itself, it will have to enable the cloudguard with
  selfManageResources flag set to true. (More details in cloud_guard_configuration.tf)
  If user chooses to create its own ORACLE Managed Entity, below value should have id of the ORACLE
  MANAGED Entity which is always available
 */
  source_detector_recipe_id = "${data.oci_cloud_guard_detector_recipes.test_detector_recipes.detector_recipe_collection.0.items.0.id}"

  #Optional
  description  = "${var.detector_recipe_description}"

  detector_rules {
    #Required
    details {
      #Required
      is_enabled = "${var.detector_recipe_detector_rules_details_is_enabled}"
      risk_level = "${var.detector_recipe_detector_rules_details_risk_level}"

      #Optional
      condition = "${var.detector_recipe_detector_rules_details_condition}"

      configurations {
        #Required
        config_key = "${var.detector_recipe_detector_rules_details_configurations_config_key}"
        name       = "${var.detector_recipe_detector_rules_details_configurations_name}"

        #Optional
        data_type = "${var.detector_recipe_detector_rules_details_configurations_data_type}"
        value     = "${var.detector_recipe_detector_rules_details_configurations_value}"
      }

      labels = "${var.detector_recipe_detector_rules_details_labels}"
    }
    // Make sure the detector rule id is valid for the detector recipe being cloned.
    detector_rule_id = "LB_CERTIFICATE_EXPIRING_SOON"
  }

}