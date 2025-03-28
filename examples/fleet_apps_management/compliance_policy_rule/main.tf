
variable "tenancy_ocid" {
#   default = ""
}

variable "ssh_public_key" {
#   default = ""
}

variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" { 
    # default = "" 
}

variable "compliance_policy_id" { 
    # default = "" 
}

data "oci_fleet_apps_management_compliance_policies" "test_compliance_policies" {
  compartment_id = "${var.compartment_id}"
}

resource "oci_fleet_apps_management_compliance_policy_rule" "test_compliance_policy_rule" {
  compartment_id       = "${var.compartment_id}"
  compliance_policy_id = "${var.compliance_policy_id}"
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
  patch_type = ["Security"]
  product_version {
    is_applicable_for_all_higher_versions = "false"
    version                               = "9"
  }
  severity = ["LOW"]
}
