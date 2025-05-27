
variable "tenancy_ocid" {  }

variable "private_key_path" { }
variable "user_ocid" { }
variable "fingerprint" { }

variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" {}

variable "compliance_policy_id" {}

variable "patch_type_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_fleet_apps_management_compliance_policies" "test_compliance_policies" {
  compartment_id = var.compartment_id
}

resource "oci_fleet_apps_management_compliance_policy_rule" "test_compliance_policy_rule" {
  compliance_policy_id = var.compliance_policy_id
  display_name         = "displayName"
  freeform_tags = {
    "Department" = "Finance"
  }
  grace_period = "gracePeriod2"
  patch_selection {
    days_since_release = "0"
    patch_level        = "LATEST"
    selection_type     = "PATCH_LEVEL"
  }
  patch_type_id   = [ var.patch_type_id ]
  product_version {
    is_applicable_for_all_higher_versions = "false"
    version                               = "9"
  }
  severity = ["LOW"]
}
