// For named credential, we need a vault and secrets to setup
// In macs tenancies, S1 tests will have created secrets in the macs s1_c2 compartment, using vault in root compartment

data "oci_identity_compartments" "compartments" {
  compartment_id = var.tenancy_ocid
  compartment_id_in_subtree = true
  access_level   = "ANY"
  name          = "macs_test_s1_c1"
}

data "oci_vault_secrets" "find_secrets" {
  compartment_id = lookup(data.oci_identity_compartments.compartments.compartments[0], "id")
}

output "secrets" {
  value=data.oci_vault_secrets.find_secrets
}

output "comp" {
  value = data.oci_identity_compartments.compartments
}

resource "oci_management_agent_named_credential" "test_named_credential" {
  management_agent_id =  data.oci_management_agent_management_agents.find_agent.management_agents[0].id
  name = "nc_from_example_tf"
  description = "desc"
  properties {
    #Required
    name = "DBUserName"
    value = data.oci_vault_secrets.find_secrets.secrets[0].id
    value_category = "SECRET_IDENTIFIER"
  }
  properties {
    #Required
    name = "DBPassword"
    value = data.oci_vault_secrets.find_secrets.secrets[0].id
    value_category = "SECRET_IDENTIFIER"
  }
  type = "DBCREDS"

}

data "oci_management_agent_named_credentials" "test_named_credentials" {
  management_agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id
  name = ["name"]
  state = ["ACTIVE"]
  type = ["DBCREDS"]
}

data "oci_management_agent_management_agent_named_credentials_metadata" "test_nc_metadata" {
  management_agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id
  compartment_id = var.compartment_ocid
}