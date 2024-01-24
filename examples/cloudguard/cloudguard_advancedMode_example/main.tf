// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/**
 * @author Yogeesh Kapila on 2020-09-29
 */

//Common Variables required
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

//******  CG Enable ******
variable "cloud_guard_configuration_reporting_region" {
  default = "us-ashburn-1"
}
variable "cloud_guard_configuration_status" {
  default = "ENABLED"
}
variable "cloud_guard_configuration_self_manage_resources" {
  default = true
}
resource "oci_cloud_guard_cloud_guard_configuration" "enable_cloud_guard" {
  #Required
  compartment_id   = var.tenancy_ocid
  reporting_region = var.cloud_guard_configuration_reporting_region
  status           = var.cloud_guard_configuration_status
  self_manage_resources = var.cloud_guard_configuration_self_manage_resources  //**** ADVANCED MODE ON ****
}

//******  Create a new ManagedList for ResourceOcids ******
variable "managed_list_description" {
  default = "Bucket Resource Ocids in a High Security Zone which absolutely should not be public"
}
variable "managed_list_display_name" {
  default = "Bucket OCIDS in High Security Zone"
}
variable "managed_list_list_items" {
  default = ["namespace/test1", "namespace/test2", "namespace/test3"]
}
variable "managed_list_list_type" {
  default = "RESOURCE_OCID"
}
variable "managed_list_state" {
  default = "ACTIVE"
}
resource "oci_cloud_guard_managed_list" "buckets_resource_ocids" {
  compartment_id = var.compartment_ocid
  display_name   = var.managed_list_display_name
  description   = var.managed_list_description
  list_items     = var.managed_list_list_items
  list_type      = var.managed_list_list_type

  depends_on = [oci_cloud_guard_cloud_guard_configuration.enable_cloud_guard]

  //Added only because we are testing this in a splat-disabled env
  lifecycle {
    ignore_changes = [system_tags]
  }
}

//**********  Create Oracle Managed Responder Recipe **************
data "oci_cloud_guard_responder_recipes" "list_preloaded_responder_recipes" {
  compartment_id = var.tenancy_ocid
  resource_metadata_only = true
}

resource "oci_cloud_guard_responder_recipe" "oracle_responder_recipe" {
  compartment_id             = var.tenancy_ocid
  display_name               = data.oci_cloud_guard_responder_recipes.list_preloaded_responder_recipes.responder_recipe_collection.0.items.0.display_name
  source_responder_recipe_id = data.oci_cloud_guard_responder_recipes.list_preloaded_responder_recipes.responder_recipe_collection.0.items.0.id

  depends_on = [oci_cloud_guard_cloud_guard_configuration.enable_cloud_guard]

  lifecycle {
    ignore_changes = [system_tags]
  }
}


//**********  Clone Responder Recipe **************
variable "responder_recipe_description" {
  default = "Custom Responder Recipe"
}
variable "responder_recipe_display_name" {
  default = "TF Demo Responder Recipe"
}
variable "responder_recipe_responder_rules_details_is_enabled" {
  default = false
}
variable "responder_recipe_state" {
  default = "ACTIVE"
}

resource "oci_cloud_guard_responder_recipe" "cloned_responder_recipe" {
  compartment_id             = var.compartment_ocid
  description                = var.responder_recipe_description
  display_name               = var.responder_recipe_display_name
  source_responder_recipe_id = oci_cloud_guard_responder_recipe.oracle_responder_recipe.id
  responder_rules {
    details {
      is_enabled = var.responder_recipe_responder_rules_details_is_enabled
    }
    responder_rule_id = "TERMINATE_INSTANCE"
  }

  lifecycle {
    ignore_changes = [system_tags]
  }
}

//**********  Create Oracle Managed Detector Recipe ************** cloning config recipe only
data "oci_cloud_guard_detector_recipes" "list_preloaded_detector_recipes" {
  compartment_id = var.tenancy_ocid
  resource_metadata_only = true
  display_name = "OCI Configuration Detector Recipe"
}

resource "oci_cloud_guard_detector_recipe" "oracle_detector_recipe" {
  compartment_id             = var.tenancy_ocid
  display_name               = data.oci_cloud_guard_detector_recipes.list_preloaded_detector_recipes.detector_recipe_collection.0.items.0.display_name
  source_detector_recipe_id = data.oci_cloud_guard_detector_recipes.list_preloaded_detector_recipes.detector_recipe_collection.0.items.0.id

  depends_on = [oci_cloud_guard_cloud_guard_configuration.enable_cloud_guard]

  lifecycle {
    ignore_changes = [system_tags]
  }
}

//******  Clone a Detector Recipe with Overrides ******
variable "detector_recipe_description" {
  default = "Custom Detector Recipe"
}
variable "detector_recipe_display_name" {
  default = "TF Demo Detector Recipe"
}
variable "detector_recipe_detector_rules_details_risk_level" {
  default = "HIGH"
}
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
  default = ["hsz-lb-certs"]
}
variable "detector_recipe_state" {
  default = "ACTIVE"
}

resource "oci_cloud_guard_detector_recipe" "cloned_detector_recipe" {
  compartment_id            = var.compartment_ocid
  display_name              = var.detector_recipe_display_name
  source_detector_recipe_id = oci_cloud_guard_detector_recipe.oracle_detector_recipe.id
  description  = var.detector_recipe_description
  detector_rules {
    details {
      is_enabled = var.detector_recipe_detector_rules_details_is_enabled
      risk_level = var.detector_recipe_detector_rules_details_risk_level
      condition = var.detector_recipe_detector_rules_details_condition
      configurations {
        config_key = var.detector_recipe_detector_rules_details_configurations_config_key
        name       = var.detector_recipe_detector_rules_details_configurations_name
        data_type = var.detector_recipe_detector_rules_details_configurations_data_type
        value     = var.detector_recipe_detector_rules_details_configurations_value
      }
      labels = var.detector_recipe_detector_rules_details_labels
    }
    detector_rule_id = "LB_CERTIFICATE_EXPIRING_SOON"
  }

  lifecycle {
    ignore_changes = [system_tags]
  }
}

//******  Create a target with above recipes ******
variable "target_description" {
  default = "Custom Target for High Security Zone Compartment"
}
variable "target_display_name" {
  default = "TF Demo Target"
}
variable "target_state" {
  default = "ACTIVE"
}
variable "target_target_resource_type" {
  default = "COMPARTMENT"
}
variable "target_target_detector_recipes_detector_rules_details_condition_groups_condition" {
  default = "{\"kind\":\"SIMPLE\",\"parameter\":\"lbCertificateExpiringSoonFilter\",\"operator\":\"EQUALS\",\"value\":\"20\",\"valueType\":\"CUSTOM\"}"
}
variable "target_target_responder_recipes_responder_rules_details_mode" {
  default = "USERACTION"
}
resource "oci_cloud_guard_target" "test_target" {
  compartment_id       = var.compartment_ocid
  display_name         = var.target_display_name
  target_resource_id   = var.compartment_ocid
  target_resource_type = var.target_target_resource_type
  description   = var.target_description
  state         = var.target_state

  target_detector_recipes {
    detector_recipe_id = oci_cloud_guard_detector_recipe.cloned_detector_recipe.id
    detector_rules {
      details {
        condition_groups {
          compartment_id = var.compartment_ocid
          condition      = var.target_target_detector_recipes_detector_rules_details_condition_groups_condition
        }
      }
      detector_rule_id = "LB_CERTIFICATE_EXPIRING_SOON"
    }
  }

  target_responder_recipes {
    responder_recipe_id = oci_cloud_guard_responder_recipe.cloned_responder_recipe.id
    responder_rules {
      details {
        condition = "{\"kind\":\"COMPOSITE\",\"leftOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"resourceId\",\"operator\":\"NOT_IN\",\"value\":\"${oci_cloud_guard_managed_list.buckets_resource_ocids.id}\",\"valueType\":\"MANAGED\"},\"compositeOperator\":\"AND\",\"rightOperand\":{\"kind\":\"SIMPLE\",\"parameter\":\"region\",\"operator\":\"IN\",\"value\":\"us-ashburn-1\",\"valueType\":\"CUSTOM\"}}"
        mode = var.target_target_responder_recipes_responder_rules_details_mode
      }
      responder_rule_id = "MAKE_BUCKET_PRIVATE"
    }
  }

  lifecycle {
    ignore_changes = [system_tags]
  }
}