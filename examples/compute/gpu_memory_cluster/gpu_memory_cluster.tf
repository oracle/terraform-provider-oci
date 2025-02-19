provider "oci" {
  auth                = var.auth
  config_file_profile = var.config_file_profile
  region              = var.region
}

// variables
variable "auth" {}
variable "region" {}
variable "config_file_profile" {}
variable "compartment_ocid" {}
variable "tenancy_ocid" {}
variable "gb200_image_id" {}
variable "compute_gpu_memory_cluster_size" {
  default = 18
}

// dependent data
data "oci_identity_availability_domain" "ad" {
  compartment_id = var.compartment_ocid
  ad_number      = 1
}

resource "oci_core_compute_capacity_topology" "test_compute_capacity_topology" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id = var.compartment_ocid
  display_name = "TestDedicatedCapacityTopology"
  freeform_tags = {
    "department" = "Internal"
  }
  capacity_source {
    capacity_type = "DEDICATED"
    compartment_id = var.tenancy_ocid
  }
}

output "output_compute_capacity_topology" {
  value = oci_core_compute_capacity_topology.test_compute_capacity_topology
}

// required resources
resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "TestVcn"
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = lower(
    data.oci_identity_availability_domain.ad.name,
  )
  cidr_block        = "10.0.1.0/24"
  compartment_id    = var.compartment_ocid
  vcn_id            = oci_core_vcn.test_vcn.id
  display_name      = "TestSubnet"
}

resource "oci_core_network_security_group" "test_network_security_group" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "TestNetworkSecurityGroup"
}

resource "oci_core_instance_configuration" "test_instance_configuration" {
  compartment_id = var.compartment_ocid
  display_name   = "TestInstanceConfiguration"

  instance_details {
    instance_type = "compute"

    launch_details {
      // Since sufficient capacity available only in AD3
      availability_domain = data.oci_identity_availability_domain.ad.name
      compartment_id      = var.compartment_ocid
      shape = "BM.GPU.GB200.4"

      metadata = {}

      source_details {
        image_id                = var.gb200_image_id
        source_type             = "image"
      }

      create_vnic_details {
        assign_public_ip       = "false"
        nsg_ids                = [oci_core_network_security_group.test_network_security_group.id]
        subnet_id              = oci_core_subnet.test_subnet.id
        assign_private_dns_record  = "true"
      }
    }
  }
}

output "output_instance_configuration" {
  value = oci_core_instance_configuration.test_instance_configuration
}

resource "oci_core_compute_cluster" "test_compute_cluster" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestComputeCluster"
}

output "output_compute_cluster" {
  value = oci_core_compute_cluster.test_compute_cluster
}

resource "oci_identity_tag_namespace" "test_tag_namespace" {
  compartment_id = var.compartment_ocid
  description    = "test tag namespace"
  name           = "test-tag-namespace-all"
}

resource "oci_identity_tag" "tag" {
  description      = "test tag"
  name             = "test-tag"
  tag_namespace_id = oci_identity_tag_namespace.test_tag_namespace.id
}

// our new data sources & resources
data "oci_core_compute_gpu_memory_fabrics" "all_available_memory_fabrics" {
  compartment_id = var.compartment_ocid
  availability_domain = data.oci_identity_availability_domain.ad.name
  compute_gpu_memory_fabric_health = "HEALTHY"
  compute_gpu_memory_fabric_lifecycle_state = "AVAILABLE"

  depends_on = [oci_core_compute_capacity_topology.test_compute_capacity_topology]
}

output "all_available_gpu_memory_fabrics" {
  value = data.oci_core_compute_gpu_memory_fabrics.all_available_memory_fabrics
}

resource "oci_core_compute_gpu_memory_cluster" "test_compute_gpu_memory_cluster" {
  #Required
  availability_domain       = data.oci_identity_availability_domain.ad.name
  compartment_id            = var.compartment_ocid
  compute_cluster_id        = oci_core_compute_cluster.test_compute_cluster.id
  instance_configuration_id = oci_core_instance_configuration.test_instance_configuration.id

  #Optional
  defined_tags         = {
    "${oci_identity_tag_namespace.test_tag_namespace.name}.${oci_identity_tag.tag.name}" = "TestGMC-tag"
  }
  display_name         = "TestGMC"
  freeform_tags        = { "department" = "Internal" }
  gpu_memory_fabric_id = data.oci_core_compute_gpu_memory_fabrics.all_available_memory_fabrics.compute_gpu_memory_fabric_collection[0].items[0].compute_gpu_memory_fabric_id
  size                 = var.compute_gpu_memory_cluster_size

  depends_on = [oci_core_compute_capacity_topology.test_compute_capacity_topology]
}

output "output_get_gpu_memory_cluster" {
  value = oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster
}

data "oci_core_compute_gpu_memory_clusters" "test_compute_gpu_memory_clusters" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  availability_domain           = data.oci_identity_availability_domain.ad.name
}

output "output_list_gpu_memory_clusters" {
  value = data.oci_core_compute_gpu_memory_clusters.test_compute_gpu_memory_clusters
}

// related gmc cluster instances data source
data "oci_core_compute_gpu_memory_cluster_instances" "test_compute_gpu_memory_cluster_instances" {
  compute_gpu_memory_cluster_id = oci_core_compute_gpu_memory_cluster.test_compute_gpu_memory_cluster.id
}

output "list_gpu_memory_cluster_instances" {
  value = data.oci_core_compute_gpu_memory_cluster_instances.test_compute_gpu_memory_cluster_instances
}
