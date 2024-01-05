// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "target_access_level" {
  default = "ACCESSIBLE"
}

variable "target_compartment_id_in_subtree" {
  default = true
}
//Refer to the note in managed_list.tf for the above two variables

variable "target_defined_tags_value" {
  default = "value"
}

variable "target_description" {
  default = "description"
}

//Has to be unique
variable "target_display_name" {
  default = "displayName"
}

variable "target_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}

//Acceptable values come from LifecycleStateEnum
variable "target_state" {
  default = "ACTIVE"
}

/*
The configuration and condition Objects are dependent on the specific detector rule id and the same for
responder rule id.
Hence for testing purposes we are going to hardcode a detector rule id and a responder rule id each having
a valid condition and configuration.
*/
//Below is an example of a composite condition, refer to detector_recipe.tf for a sample of SIMPLE condition
variable "target_target_detector_recipes_detector_rules_details_condition_groups_condition" {
  default = "{\"kind\":\"COMPOSITE\",\"leftOperand\":{\"kind\" :\"COMPOSITE\",\"leftOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"12\",\"valueType\":\"CUSTOM\"},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"NOT_EQUALS\",\"value\":\"12\",\"valueType\":\"CUSTOM\"}},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"10\",\"valueType\":\"CUSTOM\"}}"
}

/*
Takes acceptable values from TargetResourceTypeEnum,but as of now only allowed and acceptable value is COMPARTMENT
*/
variable "target_target_resource_type" {
  default = "COMPARTMENT"
}

variable "target_target_responder_recipes_responder_rules_details_condition" {
  default = "{\"kind\":\"SIMPLE\",\"parameter\":\"resourceName\",\"operator\":\"EQUALS\",\"value\":\"namespace/bucket1\",\"valueType\":\"CUSTOM\"}"
}

variable "target_target_responder_recipes_responder_rules_details_configurations_config_key" {
  default = "autoBackupWindowConfig"
}

variable "target_target_responder_recipes_responder_rules_details_configurations_name" {
  default = "Backup time window (Slot)"
}

variable "target_target_responder_recipes_responder_rules_details_configurations_value" {
  default = "10"
}

//Acceptable values come from ResponderModeTypesEnum
variable "target_target_responder_recipes_responder_rules_details_mode" {
  default = "USERACTION"
}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

/*
We got the two below data sources in order to get the detector/responder recipe id to attach to target.
For more explanation refer to detector_recipe.tf or responder_recipe.tf
*/
data "oci_cloud_guard_detector_recipes" "test_detector_recipes" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  #Optional
  state          = "ACTIVE"
  //Adding this to make sure the detector rule id is compatible with the returned detector recipes
  display_name   = "OCI Configuration Detector Recipe"
}

data "oci_cloud_guard_responder_recipes" "test_responder_recipes" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  #Optional
  state          = "ACTIVE"
}



resource "oci_cloud_guard_target" "test_target" {
  #Required
  compartment_id       = "${var.compartment_id}"
  display_name         = "${var.target_display_name}"
  //For now target resource id has to be equal to comaprtment id
  target_resource_id   = "${var.compartment_id}"
  target_resource_type = "${var.target_target_resource_type}"

  #Optional
  description   = "${var.target_description}"
  state         = "${var.target_state}"

  target_detector_recipes {
    #Required
    detector_recipe_id = "${data.oci_cloud_guard_detector_recipes.test_detector_recipes.detector_recipe_collection.0.items.0.id}"

    #Optional
    detector_rules {
      #Required
      details {
        #Optional
        condition_groups {
          #Required
          compartment_id = "${var.compartment_id}"
          condition      = "${var.target_target_detector_recipes_detector_rules_details_condition_groups_condition}"
        }
      }

      detector_rule_id = "LB_CERTIFICATE_EXPIRING_SOON"
    }
  }

  target_responder_recipes {
    #Required
    responder_recipe_id = "${data.oci_cloud_guard_responder_recipes.test_responder_recipes.responder_recipe_collection.0.items.0.id}"

    #Optional
    responder_rules {
      #Required
      details {
        #Optional
        condition = "${var.target_target_responder_recipes_responder_rules_details_condition}"

        configurations {
          #Required
          config_key = "${var.target_target_responder_recipes_responder_rules_details_configurations_config_key}"
          name       = "${var.target_target_responder_recipes_responder_rules_details_configurations_name}"
          value      = "${var.target_target_responder_recipes_responder_rules_details_configurations_value}"
        }

        mode = "${var.target_target_responder_recipes_responder_rules_details_mode}"
      }

      responder_rule_id = "ENABLE_DB_BACKUP"
    }
  }
}

data "oci_cloud_guard_targets" "test_targets" {
  #Required
  compartment_id = "${var.compartment_id}"

  #Optional
  access_level              = "${var.target_access_level}"
  compartment_id_in_subtree = "${var.target_compartment_id_in_subtree}"
  display_name              = "${var.target_display_name}"
  state                     = "${var.target_state}"
}