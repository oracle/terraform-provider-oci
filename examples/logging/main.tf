# to run the tests, replace the user, fingerprint and private key path
# run commands in this folder:
# terraform init
# terraform plan
# terraform apply to create resources
# terraform destroy to remove resources
locals {
  region = "us-phoenix-1"
  tenancy_ocid = "ocid1.tenancy.oc1..aaaaaaaa4s2hncj4oaulmf5tz4yfeska6fya4gkd5jsg3fmlgq7pprgr7wiq"
  user_ocid = "ocid1.user.oc1..aaaaaaaa5mxx5f6ltt5w6soekhzetsymtgntbjijmikyc4kqpbau2xfwnsva"
  fingerprint = "1e:84:e8:12:08:55:af:2e:51:5e:2a:57:41:ab:fd:c9"
  private_key_path = "/Users/shxi/.oci/oci_api_key.pem"
}
module "identity" {
  source = "./identity"
  compartment_id = var.compartment_ocid
}

module "log_group" {
  source = "./log_group"
  compartment_id = var.compartment_ocid
  tag2_name = module.identity.tag2_name
  tag_namespace1_name = module.identity.tag_namespace1_name
}

module "log" {
  source = "./log"
  test_log_group_id = module.log_group.test_log_group_id
  tag2_name = module.identity.tag2_name
  tag_namespace1_name = module.identity.tag_namespace1_name
}

module "log_saved_search" {
  source = "./log_saved_search"
  compartment_id = var.compartment_ocid
  tag1_name = module.identity.tag1_name
  tag_namespace1_name = module.identity.tag_namespace1_name
}

module "log_agent_configuration" {
  source = "./log_agent_configuration"
  compartment_id = var.compartment_ocid
  tag1_name = module.identity.tag1_name
  tag_namespace1_name = module.identity.tag_namespace1_name
  test_log_id = module.log.test_log_id
  test_log_group_id = module.log_group.test_log_group_id
}
