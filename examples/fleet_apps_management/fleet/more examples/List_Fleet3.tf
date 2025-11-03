
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

data "oci_fleet_apps_management_fleets" "test_fleets" {
  application_type = "applicationType"
  compartment_id   = "${var.compartment_id}"
  display_name     = "displayName2"
  environment_type = "environmentType2"
  filter {
    name   = "id"
    values = ["${oci_fleet_apps_management_fleet.test_fleet.id}"]
  }
  fleet_type = "GENERIC"
  id         = "${oci_fleet_apps_management_fleet.test_fleet.id}"
  state      = "NEEDS_ATTENTION"
}
variable "compartment_id" { default = "ocid1.compartment.oc1.." }
variable "tenancy_id" { default = "ocid1.tenancy.oc1.." }
variable "test_instance_id" { default = "ocid1.instance.oc1." }
variable "key_id" { default = "ocid1.key.oc1." }
variable "vault_id" { default = "ocid1.vault.oc1." }
variable "oci_ons_notification_topic" { default = "ocid1.onstopic.oc1." }
variable "oci_kms_vault" { default = "ocid1.vault.oc1." }
variable "fams_user_id" { default = "ocid1.vaultsecret.oc1." }
variable "fams_user_password" { default = "ocid1.vaultsecret.oc1." }

resource "oci_fleet_apps_management_fleet" "test_fleet" {
  compartment_id = "${var.compartment_id}"
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
  defined_tags = "${map("Oracle-Tags.CreatedBy", "updatedValue")}"
  description  = "description2"
  details {
    fleet_type = "GENERIC"
  }
  display_name     = "displayName2"
  environment_type = "Stage"
  freeform_tags = {
    "Department" = "Accounting"
  }
  is_target_auto_confirm = "true"
  notification_preferences {
    compartment_id = "${var.compartment_id}"
    preferences {
      on_job_canceled            = "true"
      on_job_failure             = "true"
      on_job_schedule_change     = "true"
      on_job_start               = "true"
      on_job_success             = "true"
      on_resource_non_compliance = "true"
      on_runbook_newer_version   = "true"
      on_task_failure            = "true"
      on_task_pause              = "true"
      on_task_success            = "true"
      on_topology_modification   = "true"
      upcoming_schedule {
        notify_before        = "notifyBefore2"
        on_upcoming_schedule = "true"
      }
    }
    topic_id = "${var.oci_ons_notification_topic}"
  }
  resource_selection {
    resource_selection_type = "DYNAMIC"
    rule_selection_criteria {
      match_condition = "MATCH_ALL"
      rules {
        basis                     = "inventoryProperties"
        compartment_id            = "${var.tenancy_id}"
        compartment_id_in_subtree = "true"
        conditions {
          attr_group = "Instance"
          attr_key   = "shape"
          attr_value = "VM.Standard.E4.Flex"
        }
        match_condition         = "ANY"
        resource_compartment_id = "${var.compartment_id}"
      }
    }
  }
}
