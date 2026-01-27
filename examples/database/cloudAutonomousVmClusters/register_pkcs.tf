# Example: Register PKCS on an externally-created Cloud Autonomous VM Cluster using Terraform (import pattern)
#
# **This pattern assumes the cluster was created outside Terraform (via console/other automation).**
# - Import the resource by OCID.
# - Then use the register_pkcs_trigger and tde_key_store_type fields to invoke the registerPkcs SPECIAL_UPDATE action.

variable "cloud_autonomous_vm_cluster_id" {
   description = "OCID of an externally created Cloud Autonomous VM Cluster"
   type        = string
 }

 variable "tde_key_store_type" {
   description = "The keystore type for the PKCS11 driver (e.g., AWS, ORACLE)."
   type        = string
   default     = "AWS"
 }

 variable "register_pkcs_trigger" {
   description = "Change this value (increment by 1) to re-trigger the registerPkcs operation. See usage instructions."
   type        = number
   default     = 1
 }

 variable "unregister_pkcs_trigger" {
   description = "Change this value (increment by 1) to re-trigger the unregisterPkcs operation. See usage instructions."
   type        = number
   default     = 1
 }

 resource "oci_database_cloud_autonomous_vm_cluster" "external_pkcs" {
   # Required fields for import validation
   display_name                    = var.cloud_autonomous_vm_cluster_display_name
   compartment_id                  = var.compartment_ocid
   cloud_exadata_infrastructure_id = var.cloud_exadata_infrastructure_id
   subnet_id                       = var.subnet_id

   # Fields used for registerPkcs SPECIAL_UPDATE call

   tde_key_store_type    = var.tde_key_store_type
#    register_pkcs_trigger = var.register_pkcs_trigger
   unregister_pkcs_trigger = var.unregister_pkcs_trigger
 }

# USAGE WORKFLOW:
# 1. Provide the OCID for your existing Cloud Autonomous VM Cluster using one of:
#      - Create a terraform.tfvars file with content:
#           cloud_autonomous_vm_cluster_id = ""
#        (Recommended)
#      - OR pass on the CLI:
#           terraform apply -var="cloud_autonomous_vm_cluster_id="
#      - OR set environment variable:
#           export TF_VAR_cloud_autonomous_vm_cluster_id=""
# 2. Run: terraform init
# 3. IMPORT the cluster into Terraform state:
#      terraform import oci_database_cloud_autonomous_vm_cluster.external_pkcs <cloud_autonomous_vm_cluster_OCID>
# 4. To invoke registerPkcs, increment 'register_pkcs_trigger' and apply:
#      terraform apply
#    (You can re-invoke by incrementing again.)
# 5. Only the SPECIAL_UPDATE action is affected; unrelated fields will remain as managed outside of Terraform.

# Example terraform.tfvars for variable assignment:
# cloud_autonomous_vm_cluster_id = ""
# tde_key_store_type             = "AWS"
# register_pkcs_trigger          = 1
