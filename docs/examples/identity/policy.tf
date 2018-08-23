/*
 * This example file shows how to create a policy pertaining to a compartment and group. It is important to
 * note the use of interpolation syntax in the statement field, this ensures Terraform creates the group and 
 * compartment prior to creating the policy.  
 */

resource "oci_identity_policy" "policy1" {
  name           = "tf-example-policy"
  description    = "policy created by terraform"
  compartment_id = "${data.oci_identity_compartments.compartments1.compartments.0.id}"

  statements = ["Allow group ${oci_identity_group.group1.name} to read instances in compartment ${data.oci_identity_compartments.compartments1.compartments.0.name}",
    "Allow group ${oci_identity_group.group1.name} to inspect instances in compartment ${data.oci_identity_compartments.compartments1.compartments.0.name}",
  ]
}

data "oci_identity_policies" "policies1" {
  compartment_id = "${data.oci_identity_compartments.compartments1.compartments.0.id}"

  filter {
    name   = "name"
    values = ["tf-example-policy"]
  }
}

output "policy" {
  value = "${data.oci_identity_policies.policies1.policies}"
}

/*
 * Policies for dynamic groups
 */
resource "oci_identity_policy" "dynamic-policy-1" {
  name           = "tf-example-dynamic-policy"
  description    = "dynamic policy created by terraform"
  compartment_id = "${data.oci_identity_compartments.compartments1.compartments.0.id}"

  statements = ["Allow dynamic-group ${oci_identity_dynamic_group.dynamic-group-1.name} to read instances in compartment ${data.oci_identity_compartments.compartments1.compartments.0.name}",
    "Allow dynamic-group ${oci_identity_dynamic_group.dynamic-group-1.name} to inspect instances in compartment ${data.oci_identity_compartments.compartments1.compartments.0.name}",
  ]
}

data "oci_identity_policies" "dynamic-policies-1" {
  compartment_id = "${data.oci_identity_compartments.compartments1.compartments.0.id}"

  filter {
    name   = "id"
    values = ["${oci_identity_policy.dynamic-policy-1.id}"]
  }
}

output "dynamicPolicies" {
  value = "${data.oci_identity_policies.dynamic-policies-1.policies}"
}
