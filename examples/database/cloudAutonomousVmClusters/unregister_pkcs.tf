# Example: Unregister PKCS on an externally-created Cloud Autonomous VM Cluster using Terraform (SPECIAL_UPDATE pattern)
#
# This pattern assumes the cluster was created outside Terraform (via console/other automation).
# - Import the resource by OCID.
# - Then use the unregister_pkcs_trigger and tde_key_store_type = "NONE" fields to invoke the unregisterPkcs SPECIAL_UPDATE action.

# variable "cloud_autonomous_vm_cluster_id" {
#   description = "OCID of an externally created Cloud Autonomous VM Cluster"
#   type        = string
# }
#
# variable "unregister_pkcs_trigger" {
#   description = "Change this value (increment by 1) to re-trigger the unregisterPkcs operation. See usage instructions."
#   type        = number
#   default     = 2
# }
#
# resource "oci_database_cloud_autonomous_vm_cluster" "external_pkcs_unregister" {
#   # Required fields as in register_pkcs.tf
#   display_name                    = var.cloud_autonomous_vm_cluster_display_name
#   compartment_id                  = var.compartment_ocid
#   cloud_exadata_infrastructure_id = var.cloud_exadata_infrastructure_id
#   subnet_id                       = var.subnet_id
#
#   # Fields for unregistering PKCS
#   tde_key_store_type      = "AWS"
#   unregister_pkcs_trigger = var.unregister_pkcs_trigger
# }

# USAGE WORKFLOW:
# 1. Provide the OCID for your existing Cloud Autonomous VM Cluster, e.g. in terraform.tfvars:
#      cloud_autonomous_vm_cluster_id = ""
#      cloud_autonomous_vm_cluster_display_name = ""
#      compartment_ocid = ""
#      cloud_exadata_infrastructure_id = ""
#      subnet_id = ""
# 2. Run: terraform init
# 3. IMPORT the cluster if not already done:
#      terraform import oci_database_cloud_autonomous_vm_cluster.external_pkcs_unregister <cloud_autonomous_vm_cluster_OCID>
# 4. To invoke unregisterPkcs, increment 'unregister_pkcs_trigger' and apply:
#      terraform apply
#    (Increment again for re-invocation.)
