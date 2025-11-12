// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


#########################################################################################################
# OCI Fleet Apps Management - Fleet Resource (Enhanced Documentation)
#########################################################################################################

variable "tenancy_ocid" {
  type        = string
  description = "The OCID of the tenancy. Example: ocid1.tenancy.oc1..<unique_id>"
}

variable "compartment_id" {
  type        = string
  description = "The OCID of the compartment where the Fleet resource will be created."
}

variable "instance_id" {
  type        = string
  description = "OCID of a compute instance associated with fleet credentials."
}

variable "key_id" {
  type        = string
  description = "OCID of the Vault key used for credential encryption."
}

variable "vault_id" {
  type        = string
  description = "OCID of the Vault containing secrets for Fleet credentials."
}

variable "fams_user_id_secret" {
  type        = string
  description = "OCID of the secret storing the Fleet user ID."
}

variable "fams_user_password_secret" {
  type        = string
  description = "OCID of the secret storing the Fleet user password."
}

variable "notification_topic_id" {
  type        = string
  description = "OCID of the OCI Notifications topic used for Fleet event notifications."
}

variable "environment_type" {
  type        = string
  description = <<EOT
Defines the environment classification of the Fleet.
Possible values:
  - PROD  : Production environment.
  - STAGE : Staging or pre-production environment.
  - DEV   : Development or test environment.
EOT
  default     = "STAGE"
}

variable "fleet_type" {
  type        = string
  description = <<EOT
Specifies the type of Fleet managed.
Possible values:
  - GENERIC        : Generic fleet, not tied to a specific Oracle product.
  - ORACLE_FUSION  : Fleet managing Oracle Fusion Applications.
  - ORACLE_EBS     : Fleet managing Oracle E-Business Suite instances.
  - ORACLE_DB      : Fleet managing Oracle Database environments.
EOT
  default     = "GENERIC"
}

variable "credential_level" {
  type        = string
  description = <<EOT
Defines the credential scope.
Possible values:
  - FLEET  : Credential applies to the entire fleet.
  - TARGET : Credential applies to a specific target resource.
EOT
  default     = "FLEET"
}

variable "credential_type" {
  type        = string
  description = <<EOT
Specifies how the credential is stored or protected.
Possible values:
  - PLAIN_TEXT     : Direct text value stored in configuration.
  - VAULT_SECRET   : Reference to a secret in OCI Vault.
  - KMS_ENCRYPTED  : Credential encrypted using a Vault KMS key.
EOT
  default     = "PLAIN_TEXT"
}

variable "resource_selection_type" {
  type        = string
  description = <<EOT
Determines how resources are associated with the Fleet.
Possible values:
  - DYNAMIC : Automatically select resources based on defined rules.
  - STATIC  : Manually attach specific resources to the Fleet.
EOT
  default     = "DYNAMIC"
}

variable "match_condition" {
  type        = string
  description = <<EOT
Defines the rule evaluation condition for resource selection.
Possible values:
  - MATCH_ALL : All rules must be satisfied.
  - MATCH_ANY : Any rule may be satisfied.
EOT
  default     = "MATCH_ALL"
}

#########################################################################################################
# Resource Definition
#########################################################################################################

resource "oci_fleet_apps_management_fleet" "test_fleet" {
  compartment_id        = var.compartment_id
  display_name          = "sample-fleet"
  description           = "Sample Fleet demonstrating variable and resource setup."
  environment_type      = var.environment_type
  is_target_auto_confirm = true

  defined_tags = {
    "Oracle-Tags.CreatedBy" = "user"
  }

  freeform_tags = {
    "Department" = "Engineering"
  }

  details {
    fleet_type = var.fleet_type
  }

  #######################################################################################################
  # Fleet Credentials
  #######################################################################################################
  credentials {
    compartment_id = var.compartment_id
    display_name   = "fleet-credential"

    entity_specifics {
      credential_level = var.credential_level
      resource_id      = var.instance_id
      target           = "target-host"

      variables {
        name  = "username"
        value = "opc"
      }
    }

    user {
      credential_type = var.credential_type
      key_id          = var.key_id
      key_version     = "1"
      secret_id       = var.fams_user_id_secret
      secret_version  = "latest"
      value           = "opc-user"
      vault_id        = var.vault_id
    }

    password {
      credential_type = var.credential_type
      key_id          = var.key_id
      key_version     = "1"
      secret_id       = var.fams_user_password_secret
      secret_version  = "latest"
      value           = "password"
      vault_id        = var.vault_id
    }
  }

  #######################################################################################################
  # Notifications
  #######################################################################################################
  notification_preferences {
    compartment_id = var.compartment_id
    topic_id       = var.notification_topic_id

    preferences {
      on_job_start               = true
      on_job_success             = true
      on_job_failure             = true
      on_job_canceled            = true
      on_task_success            = true
      on_task_failure            = true
      on_task_pause              = true
      on_runbook_newer_version   = true
      on_topology_modification   = true
      on_job_schedule_change     = true
      on_resource_non_compliance = true

      upcoming_schedule {
        on_upcoming_schedule = true
        notify_before        = "P1D"  # ISO8601 duration (e.g. P1D = 1 day before)
      }
    }
  }

  #######################################################################################################
  # Dynamic Resource Selection Rules
  #######################################################################################################
  resource_selection {
    resource_selection_type = var.resource_selection_type

    rule_selection_criteria {
      match_condition = var.match_condition

      rules {
        basis                     = "inventoryProperties"
        compartment_id            = var.compartment_id
        compartment_id_in_subtree = true
        match_condition           = "ANY"

        conditions {
          attr_group = "Instance"
          attr_key   = "shape"
          attr_value = "VM.Standard.E4.Flex"
        }
      }
    }
  }
}

#########################################################################################################
# Data Source (Singular)
#########################################################################################################

data "oci_fleet_apps_management_fleet" "test_fleet" {
  fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
}

#########################################################################################################
# Data Source (List)
#########################################################################################################

data "oci_fleet_apps_management_fleets" "test_fleets" {
  compartment_id   = var.compartment_id
  display_name     = "sample-fleet"
  environment_type = var.environment_type
  fleet_type       = var.fleet_type
  state            = "ACTIVE"  # Possible values: ACTIVE, NEEDS_ATTENTION, INACTIVE

  filter {
    name   = "id"
    values = [oci_fleet_apps_management_fleet.test_fleet.id]
  }
}


#########################################################################################################
# Computed Attributes (Available in Data Source)
#########################################################################################################
# The following attributes are computed after resource creation or retrieved via data sources:
#
# - id                      : Unique OCID of the Fleet.
# - compartment_id           : OCID of the compartment where the Fleet resides.
# - resource_region          : OCI Region where the Fleet is hosted (e.g., us-ashburn-1).
# - display_name             : Display name of the Fleet.
# - description              : Description of the Fleet.
# - products                 : List of associated products or managed services (if any).
# - details                  : Object providing internal details about the Fleet configuration:
#       - fleet_type          : Type of fleet (e.g., GENERIC, ORACLE_FUSION, ORACLE_EBS, ORACLE_DB).
# - environment_type         : Environment classification (e.g., Stage, Prod, Dev).
# - resource_selection       : Object describing how resources are associated with the Fleet:
#       - resource_selection_type : STATIC or DYNAMIC.
#       - rule_selection_criteria :
#           - match_condition   : MATCH_ALL or MATCH_ANY.
#           - rules[] :
#               - basis                     : Basis for selection (e.g., inventoryProperties).
#               - compartment_id            : Parent compartment ID for evaluation.
#               - resource_compartment_id   : Target compartment for the resources.
#               - compartment_id_in_subtree : Boolean indicating recursive selection.
#               - match_condition           : Rule-level condition (ANY / ALL).
#               - conditions[] :
#                   - attr_group : Attribute group (e.g., Instance).
#                   - attr_key   : Attribute key (e.g., shape).
#                   - attr_value : Attribute value (e.g., VM.Standard.E4.Flex).
# - notification_preferences : List of subscribed notification settings:
#       - topic_id            : OCID of the associated Notification topic.
#       - compartment_id      : Compartment of the topic.
#       - preferences :
#           - on_job_start, on_job_success, on_job_failure, on_job_canceled,
#             on_job_schedule_change, on_resource_non_compliance,
#             on_runbook_newer_version, on_task_success, on_task_failure,
#             on_task_pause, on_topology_modification : Boolean flags.
#           - upcoming_schedule :
#               - on_upcoming_schedule : Boolean flag to enable pre-schedule notification.
#               - notify_before        : Duration before schedule (e.g., P1D = 1 day).
# - is_target_auto_confirm    : Indicates whether discovered targets are auto-confirmed (true/false).
# - parent_fleet_id           : Parent Fleet OCID if this is part of a Fleet hierarchy (else null).
# - lifecycle_state           : Current lifecycle state of the Fleet (e.g., ACTIVE, NEEDS_ATTENTION, INACTIVE).
# - lifecycle_details         : JSON-encoded diagnostic information about Fleet substate or errors.
# - time_created              : RFC3339 timestamp when the Fleet was created.
# - time_updated              : RFC3339 timestamp when the Fleet was last updated.
# - defined_tags              : System-assigned and user-defined tags within defined namespaces.
# - freeform_tags             : User-defined key-value pairs applied directly.
# - system_tags               : Tags applied automatically by Oracle Cloud Infrastructure.
#########################################################################################################
