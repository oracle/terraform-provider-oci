#########################################################################################################
# Variables for OCI Apps Management Property
#########################################################################################################

variable "tenancy_ocid" {
  type        = string
  description = "The OCID of the tenancy. Example: ocid1.tenancy.oc1..<unique_id>"
}

variable "compartment_id" {
  type        = string
  description = "The OCID of the compartment where the Property will be created."
}

variable "display_name" {
  type        = string
  description = "The display name of the Property."
  default     = "Property1"
}

variable "values" {
  type        = list(string)
  description = "List of allowed or default values for this property."
  default     = ["values"]
}

#########################################################################################################
# Enumerations
#########################################################################################################

variable "selection" {
  type        = string
  description = <<EOT
Defines the selection mode for the Property — how values can be chosen or assigned.

Possible values:
  - SINGLE_CHOICE : Allows exactly one value to be selected.
  - MULTI_CHOICE  : Allows multiple values to be selected.
  - INPUT_TEXT     : Allows user-entered, unrestricted input.
EOT
  default = "SINGLE_CHOICE"
}

variable "value_type" {
  type        = string
  description = <<EOT
Defines the data type of the Property values.

Possible values:
  - STRING : Property values are simple text strings.
  - NUMERIC: Property values are numeric (integer or float).
EOT
  default = "STRING"
}

variable "scope" {
  type        = string
  description = <<EOT
Defines the functional scope where this property applies.

Possible values:
  - TAXONOMY 
  - PLATFORM_CONFIG
EOT
  default = "TAXONOMY"
}

variable "type" {
  type        = string
  description = <<EOT
Specifies whether the property is Oracle-provided or user-defined.

Possible values:
  - USER_DEFINED
  - SYSTEM_DEFINED
  - ORACLE_DEFINED
EOT
  default = "USER_DEFINED"
}

variable "state" {
  type        = string
  description = <<EOT
Defines the lifecycle state of the Property.

Possible values:
  - ACTIVE
  - FAILED 
  - DELETED
  - UPDATING
EOT
  default = "ACTIVE"
}

#########################################################################################################
# Resource Definition
#########################################################################################################

resource "oci_fleet_apps_management_property" "test_property" {
  compartment_id = var.compartment_id
  display_name   = var.display_name
  selection      = var.selection
  value_type     = var.value_type
  values         = var.values

  defined_tags = {
    "Oracle-Tags.CreatedBy" = "value"
  }
}

#########################################################################################################
# Data Source (Singular)
#########################################################################################################

data "oci_fleet_apps_management_property" "test_property" {
  property_id = oci_fleet_apps_management_property.test_property.id
}

#########################################################################################################
# Data Source (List)
#########################################################################################################

data "oci_fleet_apps_management_properties" "test_properties" {
  compartment_id = var.compartment_id
  display_name   = var.display_name
  scope          = var.scope
  state          = var.state
  type           = var.type

  filter {
    name   = "id"
    values = [oci_fleet_apps_management_property.test_property.id]
  }
}

#########################################################################################################
# Computed Attributes (Returned by Data Source)
#########################################################################################################

# id                : Unique OCID of the Property.
# compartment_id     : The compartment OCID where the property is defined.
# display_name       : The name of the Property as displayed in the console.
# description        : Optional text describing the purpose of the property.
# selection          : SINGLE_CHOICE / MULTI_CHOICE / FREE_FORM.
# value_type         : STRING / NUMERIC
# values             : List of valid or default values for the property.
# scope              : TAXONOMY / PLATFORM_CONFIG.
# type               : USER_DEFINED / SYSTEM_DEFINED/ ORACLE_DEFINED.
# state              : ACTIVE / FAILED / DELETED/ UPDATING.
# lifecycle_details  : Additional diagnostic or status information.
# defined_tags       : Namespaced system or user-defined tags.
# freeform_tags      : Simple user-defined key-value pairs.
# system_tags        : Oracle-managed metadata tags.
#########################################################################################################
