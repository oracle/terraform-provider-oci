variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
  default = "us-ashburn-1"
}

variable "cluster_id" {
}

provider "oci" {
  region           = var.region
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

data "oci_containerengine_cluster" "test_cluster" {
    #Required
    # replace the cluster id with a cluster which is pending lb decommission
    cluster_id = var.cluster_id
}

resource "oci_containerengine_cluster_public_api_endpoint_decommission_manager" "decommission_manager"{
    cluster_id = data.oci_containerengine_cluster.test_cluster.id
    is_public_api_endpoint_decommissioned = true
    rollback_deadline_delay = "P1D"
}

data "oci_containerengine_cluster_public_api_endpoint_decommission_status" "getStatus" {
  cluster_id = data.oci_containerengine_cluster.test_cluster.id
}

output "status" {
  value = data.oci_containerengine_cluster_public_api_endpoint_decommission_status.getStatus.status
}

output "time_decommission_rollback_deadline" {
  value = data.oci_containerengine_cluster_public_api_endpoint_decommission_status.getStatus.time_decommission_rollback_deadline
}
