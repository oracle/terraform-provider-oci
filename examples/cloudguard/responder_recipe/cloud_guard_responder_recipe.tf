// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "responder_recipe_access_level" {
  default = "ACCESSIBLE"
}

variable "responder_recipe_compartment_id_in_subtree" {
  default = true
}
//Refer to the note in managed_list.tf for the above two variables

variable "responder_recipe_defined_tags_value" {
  default = "value"
}

variable "responder_recipe_description" {
  default = "description"
}

//Has to be unique
variable "responder_recipe_display_name" {
  default = "displayName"
}

variable "responder_recipe_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}

variable "responder_recipe_responder_rules_details_is_enabled" {
  default = false
}

//Acceptable values come from LifecycleStateEnum
variable "responder_recipe_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

//Plural Data Source representation
data "oci_cloud_guard_responder_recipes" "test_responder_recipes" {
  #Required
  compartment_id = "${var.tenancy_ocid}"

  #Optional
  state                     = "${var.responder_recipe_state}"
}

resource "oci_cloud_guard_responder_recipe" "test_responder_recipe" {
  #Required
  compartment_id             = "${var.compartment_id}"
  description                = "${var.responder_recipe_description}"
  display_name               = "${var.responder_recipe_display_name}"
  /*
  When CloudGuard is Enabled, an Oracle Managed Responder Recipe is made available having all the default rules in their default state.
  If an user needs to make its own recipe with customizations to the rules, it needs to clone an `ORACLE MANAGED RESPONDER RECIPE`.
  In order to clone an `ORACLE MANAGED RESPONDER RECIPE` (Create new Responder Recipe) the source responder recipe id needs to be the ocid of
  `ORACLE MANAGED RESPONDER RECIPE`.
  In order to achieve that, we will list all the responder recipes and then pick the first item (Oracle Managed Recipe) in the collection for cloning.

  However if user wants to create and manage the ORACLE MANAGED Entity itself, it will have to enable the cloudguard with
  selfManageResources flag set to true. (More details in cloud_guard_configuration.tf)
  If user chooses to create its own ORACLE Managed Entity, below value should have id of the ORACLE
  MANAGED Entity which is always available
  */
  source_responder_recipe_id = "${data.oci_cloud_guard_responder_recipes.test_responder_recipes.responder_recipe_collection.0.items.0.id}"

  #Optional

  responder_rules {
    #Required
    details {
      #Required
      is_enabled = "${var.responder_recipe_responder_rules_details_is_enabled}"
    }
    //For testing purposes we will pick up an existing responder rule id
    responder_rule_id = "MAKE_BUCKET_PRIVATE"
  }
}