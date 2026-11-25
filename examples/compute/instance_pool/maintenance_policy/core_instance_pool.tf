terraform {
  required_providers {
    oci = {
      source = "oracle/oci"
    }
  }
}

variable "compartment_ocid" {
  type = string
}

variable "instance_configuration_id" {
  type        = string
  description = "Existing instance configuration used by the pool."
}

variable "availability_domain" {
  type = string
}

variable "primary_subnet_id" {
  type = string
}

variable "pool_size" {
  type    = number
  default = 2
}

variable "is_full_refresh" {
  type    = bool
  default = false
}

variable "max_healthy_percent" {
  type    = number
  default = 110
}

resource "oci_core_instance_pool" "maintenance_policy" {
  compartment_id            = var.compartment_ocid
  instance_configuration_id = var.instance_configuration_id
  display_name              = "instance-pool-maintenance-policy-example"
  size                      = var.pool_size

  placement_configurations {
    availability_domain = var.availability_domain
    primary_subnet_id   = var.primary_subnet_id
  }

  instance_maintenance_policy {
    launch_before_terminating {
      is_full_refresh     = var.is_full_refresh
      max_healthy_percent = var.max_healthy_percent
    }
  }
}

data "oci_core_instance_pool" "maintenance_policy" {
  instance_pool_id = oci_core_instance_pool.maintenance_policy.id
}

output "instance_maintenance_policy" {
  value = data.oci_core_instance_pool.maintenance_policy.instance_maintenance_policy
}
