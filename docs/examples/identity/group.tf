/*
 * This example file shows how to create a group and add a user to it. 
 */

resource "oci_identity_group" "group1" {
  name        = "tf-example-group"
  description = "group created by terraform"
}

resource "oci_identity_user_group_membership" "user-group-mem1" {
  compartment_id = "${var.tenancy_ocid}"
  user_id        = "${oci_identity_user.user1.id}"
  group_id       = "${oci_identity_group.group1.id}"
}

data "oci_identity_groups" "groups1" {
  compartment_id = "${oci_identity_group.group1.compartment_id}"

  filter {
    name   = "name"
    values = ["tf-example-group"]
  }
}

output "groups" {
  value = "${data.oci_identity_groups.groups1.groups}"
}

/*
 * Some more directives to show dynamic groups and policy for it
 */
resource "oci_identity_dynamic_group" "dynamic-group-1" {
  compartment_id = "${var.tenancy_ocid}"
  name           = "tf-example-dynamic-group"
  description    = "dynamic group created by terraform"
  matching_rule  = "instance.compartment.id = ${oci_identity_compartment.compartment1.id}"
}

data "oci_identity_dynamic_groups" "dynamic-groups-1" {
  compartment_id = "${oci_identity_dynamic_group.dynamic-group-1.compartment_id}"

  filter {
    name   = "id"
    values = ["${oci_identity_dynamic_group.dynamic-group-1.id}"]
  }
}

output "dynamicGroups" {
  value = "${data.oci_identity_dynamic_groups.dynamic-groups-1.dynamic_groups}"
}
