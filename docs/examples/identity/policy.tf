/*
 * This example file shows how to create a policy pertaining to a compartment and group. It is important to
 * note the use of interpolation syntax in the statement field, this ensures Terraform creates the group and 
 * compartment prior to creating the policy.  
 */

resource "oci_identity_policy" "policy1" {
  name = "tf-example-policy"
  description = "policy created by terraform"
  compartment_id = "${var.tenancy_ocid}"
  statements = ["Allow group ${oci_identity_group.group1.name} to read instances in compartment ${oci_identity_compartment.compartment1.name}",
                "Allow group ${oci_identity_group.group1.name} to inspect instances in compartment ${oci_identity_compartment.compartment1.name}"]
}

data "oci_identity_policies" "policies1" {
  compartment_id = "${oci_identity_policy.policy1.compartment_id}"

  filter {
    name = "name"
    values = ["tf-example-policy"]
  }
}

output "policy" {
  value = "${data.oci_identity_policies.policies1.policies}"
}
