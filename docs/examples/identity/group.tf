/*
 * This example file shows how to create a group and add a user to it. 
 */

resource "oci_identity_group" "group1" {
  name = "tf-example-group"
  description = "group created by terraform"
}

resource "oci_identity_user_group_membership" "user-group-mem1" {
  compartment_id = "${var.tenancy_ocid}"
  user_id = "${oci_identity_user.user1.id}"
  group_id = "${oci_identity_group.group1.id}"
}

data "oci_identity_groups" "groups1" {
  compartment_id = "${oci_identity_group.group1.compartment_id}"

  filter {
    name = "name"
    values = ["tf-example-group"]
  }
}

output "groups" {
  value = "${data.oci_identity_groups.groups1.groups}"
}
