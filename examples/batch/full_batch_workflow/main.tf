provider "oci" {
}

############################################################
# Variables
############################################################
variable "compartment_ocid" {
  description = "OCID of the compartment where Batch resources will be created."
  default     = ""
}

variable "display_name_suffix" {
  description = "Suffix appended to resource display names."
  default     = "demo"
}

variable "image_url" {
  description = "Container image URL used by the Batch task environment (for example, <region>.ocir.io/<namespace>/<repo>:tag)."
  type        = string
}

variable "task_env_working_directory" {
  description = "Working directory inside the container."
  default     = "/"
}

variable "task_env_local_mount_directory" {
  description = "Local directory path inside the container where the shared volume is mounted."
  default     = "/mnt/batch"
}

variable "task_env_volume_name" {
  description = "Logical name for the Batch task environment volume."
  default     = "nfs-volume"
}

variable "task_env_mount_target_fqdn" {
  description = "Mount target FQDN providing the shared file system (leave blank to skip volume configuration)."
  default     = ""
}

variable "task_env_mount_target_export_path" {
  description = "Export path on the mount target (leave blank to skip volume configuration)."
  default     = ""
}

variable "vcn_cidr" {
  default = "10.20.0.0/16"
}

variable "subnet_cidr" {
  default = "10.20.1.0/24"
}

variable "fleet_shape_name" {
  default = "VM.Standard.E5.Flex"
}

variable "fleet_ocpus" {
  default = 1
}

variable "fleet_memory_gbs" {
  default = 16
}

locals {
  name_suffix   = var.display_name_suffix != "" ? var.display_name_suffix : "demo"
  configure_nfs = var.task_env_mount_target_fqdn != "" && var.task_env_mount_target_export_path != ""
}

############################################################
# Networking prerequisites
############################################################
resource "oci_core_vcn" "batch_vcn" {
  cidr_block     = var.vcn_cidr
  compartment_id = var.compartment_ocid
  display_name   = "batch-vcn-${local.name_suffix}"
  dns_label      = "batchvcn"
}

resource "oci_core_network_security_group" "batch_nsg" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.batch_vcn.id
  display_name   = "batch-nsg-${local.name_suffix}"
}

resource "oci_core_subnet" "batch_subnet" {
  cidr_block                 = var.subnet_cidr
  compartment_id             = var.compartment_ocid
  vcn_id                     = oci_core_vcn.batch_vcn.id
  display_name               = "batch-subnet-${local.name_suffix}"
  prohibit_public_ip_on_vnic = true
  dns_label                  = "batchsubnet"
  route_table_id             = oci_core_vcn.batch_vcn.default_route_table_id
  security_list_ids          = [oci_core_vcn.batch_vcn.default_security_list_id]
}

############################################################
# Batch Context
############################################################
resource "oci_batch_batch_context" "batch_context" {
  compartment_id = var.compartment_ocid
  display_name   = "batch-context-${local.name_suffix}"
  description    = "Terraform example Batch Context."

  entitlements = {
    entitlementA = 1
  }

  network {
    subnet_id = oci_core_subnet.batch_subnet.id
  }

  fleets {
    name                 = "fleet-${local.name_suffix}"
    type                 = "SERVICE_MANAGED_FLEET"
    max_concurrent_tasks = 1

    shape {
      shape_name    = var.fleet_shape_name
      ocpus         = var.fleet_ocpus
      memory_in_gbs = var.fleet_memory_gbs
    }
  }

  job_priority_configurations {
    tag_key       = "priority"
    tag_namespace = "example"
    values        = { high = 1 }
    weight        = 10
  }

  timeouts {
    create = "120m"
    delete = "120m"
  }
}

############################################################
# Batch Job Pool bound to the context
############################################################
resource "oci_batch_batch_job_pool" "job_pool" {
  compartment_id   = var.compartment_ocid
  batch_context_id = oci_batch_batch_context.batch_context.id
  display_name     = "batch-job-pool-${local.name_suffix}"

  timeouts {
    create = "45m"
    delete = "45m"
  }
}

############################################################
# Batch Task Environment
############################################################
resource "oci_batch_batch_task_environment" "task_environment" {
  compartment_id    = var.compartment_ocid
  display_name      = "batch-env-${local.name_suffix}"
  image_url         = var.image_url
  working_directory = var.task_env_working_directory

  dynamic "volumes" {
    for_each = local.configure_nfs ? [1] : []
    content {
      name                       = var.task_env_volume_name
      type                       = "NFS"
      mount_target_fqdn          = var.task_env_mount_target_fqdn
      mount_target_export_path   = var.task_env_mount_target_export_path
      local_mount_directory_path = var.task_env_local_mount_directory
    }
  }

  timeouts {
    create = "30m"
    delete = "30m"
  }
}

############################################################
# Batch Task Profile (minimum capacity requirements)
############################################################
resource "oci_batch_batch_task_profile" "task_profile" {
  compartment_id    = var.compartment_ocid
  display_name      = "batch-profile-${local.name_suffix}"
  description       = "Task profile referencing the demo fleet sizing constraints."
  min_ocpus         = var.fleet_ocpus
  min_memory_in_gbs = var.fleet_memory_gbs
}

############################################################
# Outputs
############################################################
output "batch_context_id" {
  description = "OCID of the created Batch Context."
  value       = oci_batch_batch_context.batch_context.id
}

output "batch_job_pool_id" {
  description = "OCID of the created Batch Job Pool."
  value       = oci_batch_batch_job_pool.job_pool.id
}

output "task_environment_id" {
  description = "OCID of the created Batch Task Environment."
  value       = oci_batch_batch_task_environment.task_environment.id
}

output "task_profile_id" {
  description = "OCID of the created Batch Task Profile."
  value       = oci_batch_batch_task_profile.task_profile.id
}

