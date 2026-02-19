// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
  description = "The OCID of the tenancy."
}

variable "region" {
  description = "The region to connect to."
}

variable "bds_instance_id" {
  description = "The OCID of the BDS instance for which to reset the password."
}

variable "service" {
  description = "The service for which to reset the password (Can be AMBARI, RANGER, HUE, or JUPYTERHUB)."
}

variable "secret_id" {
  description = "The OCID of the secret required for resetting passwords for certain services (Can be only RANGER, HUE, JUPYTERHUB). Optional. You may pass null value '' for this variable in case of AMBARI password reset."
}

variable "current_cluster_admin_password" {
  description = "The current Base64-encoded cluster admin password required for resetting passwords for certain services (Can be only RANGER, HUE, JUPYTERHUB). Optional. You may pass null value '' for this variable in case of AMBARI password reset"
}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  region       = var.region
}

resource "oci_bds_bds_cluster_admin_password_reset_action" "reset" {
  bds_instance_id                 = var.bds_instance_id
  service                         = var.service
  # Don't change the logic below
  secret_id                       = var.secret_id != "" && var.service != "AMBARI" ? var.secret_id : null
  current_cluster_admin_password  = var.current_cluster_admin_password != "" && var.service != "AMBARI" ? var.current_cluster_admin_password : null
  reset_trigger                   = timestamp()
}

output "new_generated_password" {
  value       = oci_bds_bds_cluster_admin_password_reset_action.reset.cluster_admin_password
  sensitive   = true
  description = "The new Base64-encoded admin password for the specified service. Use 'terraform output new_admin_password' to view it. For subsequent resets, update 'current_cluster_admin_password' with this password if supported by the provider."
}