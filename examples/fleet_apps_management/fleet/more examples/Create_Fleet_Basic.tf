
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }


resource "oci_fleet_apps_management_fleet" "test_fleet" {
  compartment_id = "${var.compartment_id}"
  details {
    fleet_type = "GENERIC"
  }
  display_name           = "displayName"
  is_target_auto_confirm = "false"
  resource_selection {
    resource_selection_type = "DYNAMIC"
    rule_selection_criteria {
      match_condition = "ANY"
      rules {
        basis          = "inventoryProperties"
        compartment_id = "${var.tenancy_id}"
        conditions {
          attr_group = "Instance"
          attr_key   = "displayName"
          attr_value = "attrValue1"
        }
        resource_compartment_id = "${var.compartment_id}"
      }
    }
  }
}
