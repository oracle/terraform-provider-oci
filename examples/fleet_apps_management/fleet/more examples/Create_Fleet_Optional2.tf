
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }
variable "compartment_id_for_update" { default = "ocid1.compartment.oc1.." }
variable "tenancy_id" { default = "ocid1.tenancy.oc1.." }
variable "test_instance_id" { default = "ocid1.instance.oc1." }
variable "compatible_product" { default = "ocid1.famsplatformconfiguration.oc1." }
variable "key_id" { default = "ocid1.key.oc1." }
variable "vault_id" { default = "ocid1.vault.oc1." }
variable "oci_ons_notification_topic" { default = "ocid1.onstopic.oc1." }
variable "oci_kms_vault" { default = "ocid1.vault.oc1." }
variable "fams_user_id" { default = "ocid1.vaultsecret.oc1." }
variable "fams_user_password" { default = "ocid1.vaultsecret.oc1." }

resource "oci_fleet_apps_management_fleet" "test_fleet" {
  compartment_id = "${var.compartment_id_for_update}"
  credentials {
    compartment_id = "${var.compartment_id}"
    display_name   = "tersi-testing-credential"
    entity_specifics {
      credential_level = "FLEET"
      resource_id      = "${var.test_instance_id}"
      target           = "target"
      variables {
        name  = "name"
        value = "value"
      }
    }
    password {
      credential_type = "PLAIN_TEXT"
      key_id          = "${var.key_id}"
      key_version     = "keyVersion"
      secret_id       = "${var.fams_user_password}"
      secret_version  = "secretVersion"
      value           = "value"
      vault_id        = "${var.vault_id}"
    }
    user {
      credential_type = "PLAIN_TEXT"
      key_id          = "${var.key_id}"
      key_version     = "keyVersion"
      secret_id       = "${var.fams_user_id}"
      secret_version  = "secretVersion"
      value           = "value"
      vault_id        = "${var.vault_id}"
    }
  }
  defined_tags = "${map("Oracle-Tags.CreatedBy", "value")}"
  description  = "description"
  details {
    fleet_type = "GENERIC"
  }
  display_name     = "displayName"
  environment_type = "Stage"
  freeform_tags = {
    "bar-key" = "value"
  }
  is_target_auto_confirm = "false"
  notification_preferences {
    compartment_id = "${var.compartment_id}"
    preferences {
      on_job_canceled            = "false"
      on_job_failure             = "false"
      on_job_schedule_change     = "false"
      on_job_start               = "false"
      on_job_success             = "false"
      on_resource_non_compliance = "false"
      on_runbook_newer_version   = "false"
      on_task_failure            = "false"
      on_task_pause              = "false"
      on_task_success            = "false"
      on_topology_modification   = "false"
      upcoming_schedule {
        notify_before        = "notifyBefore"
        on_upcoming_schedule = "false"
      }
    }
    topic_id = "${var.oci_ons_notification_topic}"
  }
  resource_selection {
    resource_selection_type = "DYNAMIC"
    rule_selection_criteria {
      match_condition = "ANY"
      rules {
        basis                     = "inventoryProperties"
        compartment_id            = "${var.tenancy_id}"
        compartment_id_in_subtree = "false"
        conditions {
          attr_group = "Instance"
          attr_key   = "displayName"
          attr_value = "attrValue1"
        }
        match_condition         = "MATCH_ALL"
        resource_compartment_id = "${var.compartment_id}"
      }
    }
  }
}
