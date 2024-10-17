# to run the tests, replace the user, fingerprint and private key path
# run commands in this folder:
# terraform init
# terraform plan
# terraform apply to create resources
# terraform destroy to remove resources
locals {
  region = "us-phoenix-1"
  tenancy_ocid = "ocid1.tenancy.oc1..aaaaaaaa4s2hncj4oaulmf5tz4yfeska6fya4gkd5jsg3fmlgq7pprgr7wiq"
  user_ocid = "ocid1.user.oc1..aaaaaaaark6yo7jgevogxohlgerphpr6lreunmmsovjdkhmujnuj2urix5aq"
  fingerprint = "16:9a:cf:f4:78:3f:ba:fd:67:fc:74:30:72:e8:e7:11"
  private_key_path = "/Users/zhenyao/.oci/oci_api_key.pem"
}

module "identity" {
  source = "./identity"
  compartment_id = var.compartment_ocid
}

module "log_group" {
  source = "./log_group"
  compartment_id = var.compartment_ocid
}

module "log" {
  source = "./log"
  test_log_group_id = module.log_group.test_log_group_id
}

module "log_saved_search" {
  source = "./log_saved_search"
  compartment_id = var.compartment_ocid
}

module "log_agent_configuration" {
  source = "./log_agent_configuration"
  compartment_id = var.compartment_ocid
  test_log_id = module.log.test_log_id
  test_log_group_id = module.log_group.test_log_group_id
}