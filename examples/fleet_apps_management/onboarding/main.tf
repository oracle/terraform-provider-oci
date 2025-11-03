#########################################################################################################
# OCI Fleet Apps Management - Onboarding Resource (Enhanced Documentation)
#########################################################################################################

variable "tenancy_ocid" {
  type        = string
  description = "The OCID of the tenancy. Example: ocid1.tenancy.oc1..<unique_id>"
}

variable "compartment_id" {
  type        = string
  description = "The OCID of the compartment where the Onboarding resource will be created."
}

variable "is_cost_tracking_tag_enabled" {
  type        = bool
  description = <<EOT
Specifies whether cost-tracking tags are enabled for this onboarding.
Possible values:
  - true  : Enables OCI cost tracking at the onboarding level.
  - false : Disables cost tracking.
EOT
  default     = false
}

variable "is_fams_tag_enabled" {
  type        = bool
  description = <<EOT
Specifies whether Fleet Apps Management System (FAMS) tags are enabled for this onboarding.
Possible values:
  - true  : Enables FAMS tagging for managed resources.
  - false : Disables automatic FAMS tagging.
EOT
  default     = false
}

variable "defined_tags" {
  type        = map(string)
  description = "Predefined tags scoped under namespaces. Example: { 'Oracle-Tags.CreatedBy' = 'user' }"
  default     = {}
}

variable "freeform_tags" {
  type        = map(string)
  description = "Simple key-value pairs applied without predefined namespaces. Example: { 'Department' = 'Finance' }"
  default     = {}
}

#########################################################################################################
# Resource Definition
#########################################################################################################

resource "oci_fleet_apps_management_onboarding" "test_onboarding" {
  compartment_id               = var.compartment_id
  defined_tags                 = { "Oracle-Tags.CreatedBy" = "user" }
  freeform_tags                = var.freeform_tags
  is_cost_tracking_tag_enabled = var.is_cost_tracking_tag_enabled
  is_fams_tag_enabled          = var.is_fams_tag_enabled
}

#########################################################################################################
# Data Source (List)
#########################################################################################################

data "oci_fleet_apps_management_onboardings" "test_onboardings" {
  compartment_id = var.tenancy_ocid

  state = "ACTIVE"  # Possible values: ACTIVE, INACTIVE, FAILED

  filter {
    name   = "id"
    values = [oci_fleet_apps_management_onboarding.test_onboarding.id]
  }
}

#########################################################################################################
# Computed Attributes (Returned in Data Source)
#########################################################################################################
# - id                         : OCID of the Onboarding resource.
# - compartment_id              : OCID of the compartment associated with the onboarding.
# - resource_region             : OCI region in which the onboarding resource resides (e.g., us-ashburn-1).
# - lifecycle_state             : Current lifecycle state of the onboarding (ACTIVE / INACTIVE / FAILED).
# - version                     : Current version of the onboarding configuration (e.g., 2.0.1).
# - applied_policies             : Object containing details of applied policies:
#       - id                    : Policy ID
#       - time_created          : RFC3339 timestamp when the policy was created.
#       - time_updated          : RFC3339 timestamp when the policy was last updated.
#       - statements            : List of IAM statements applied automatically.
#       - system_tags           : System-assigned tags under Oracle control.
# - is_fams_tag_enabled         : Whether FAMS tagging is enabled (true/false).
# - is_cost_tracking_tag_enabled: Whether cost tracking is enabled (true/false).
# - discovery_frequency         : Frequency of discovery scans (if applicable, else null).
# - time_created                : Timestamp when the onboarding was created (RFC3339 format).
# - time_updated                : Timestamp when the onboarding was last updated (RFC3339 format).
# - defined_tags                : Map of defined tags under namespaces.
# - freeform_tags               : Simple key-value tag pairs defined by the user.
# - system_tags                 : Tags applied automatically by Oracle Cloud.
#########################################################################################################
