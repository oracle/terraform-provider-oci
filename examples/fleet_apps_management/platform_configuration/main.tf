#########################################################################################################
# Variables for OCI Fleet Apps Management Platform Configuration
#########################################################################################################

variable "tenancy_ocid" {
  type        = string
  description = "The OCID of the tenancy. Example: ocid1.tenancy.oc1..<unique_id>"
}

variable "compartment_id" {
  type        = string
  description = "The OCID of the compartment where the Platform Configuration will be created."
}

variable "display_name" {
  type        = string
  description = "Display name for the Platform Configuration."
}

variable "description" {
  type        = string
  description = "Optional description for the Platform Configuration."
  default     = null
}

variable "config_category" {
  type        = string
  description = <<EOT
Defines the functional category of this Platform Configuration. 
Determines how the configuration is used within the Fleet Apps ecosystem.

Possible values:
  - PRODUCT                 : Configuration for an individual product (e.g., Oracle Linux, WebLogic, Java).
  - PRODUCT_STACK            : Defines a grouping of multiple related products operating together.
  - PRODUCT_STACK_AS_PRODUCT : Represents a stack that can itself be used as a single product reference.
  - PATCH_TYPE              : Represents patching metadata for a product or stack.
  - CREDENTIAL              : Stores credential or authentication configurations.
  - ENVIRONMENT             : Defines environmental settings or global configurations.
EOT
  default = "PRODUCT"
}

variable "type" {
  type        = string
  description = <<EOT
Specifies the source of this Platform Configuration.

Possible values:
  - USER_DEFINED   : User-created and manageable configuration.
  - SYSTEM_DEFINED : Oracle-provided, read-only configuration.
EOT
  default = "USER_DEFINED"
}

variable "state" {
  type        = string
  description = <<EOT
Defines the current lifecycle state of the Platform Configuration.

Possible values:
  - ACTIVE   : Configuration is available and in active use.
  - INACTIVE : Configuration is temporarily disabled.
  - DELETED  : Configuration has been removed.
EOT
  default = "ACTIVE"
}

#########################################################################################################
# Identifiers for Relationships
#########################################################################################################

variable "compatible_product_id" {
  type        = string
  description = "OCID of the compatible product Platform Configuration."
}

variable "credential_id" {
  type        = string
  description = "OCID of the credential Platform Configuration."
}

variable "patch_type_id" {
  type        = string
  description = "OCID of the patch type Platform Configuration."
}

#########################################################################################################
# Resource Definition
#########################################################################################################

resource "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
  compartment_id = var.compartment_id
  display_name   = var.display_name
  description    = var.description

  config_category_details {
    config_category = var.config_category

    compatible_products {
      id = var.compatible_product_id
    }

    credentials {
      id = var.credential_id
    }

    patch_types {
      id = var.patch_type_id
    }

    components = ["component_name"]
    versions   = ["1"]
  }

  defined_tags = {
    "Oracle-Tags.CreatedBy" = "user"
  }

  freeform_tags = {
    "Department" = "Accounting"
  }
}

#########################################################################################################
# Data Source (Singular)
#########################################################################################################

data "oci_fleet_apps_management_platform_configuration" "test_platform_configuration" {
  platform_configuration_id = oci_fleet_apps_management_platform_configuration.test_platform_configuration.id
}

#########################################################################################################
# Data Source (List)
#########################################################################################################

data "oci_fleet_apps_management_platform_configurations" "test_platform_configurations" {
  compartment_id            = var.compartment_id
  compartment_id_in_subtree = false
  config_category           = var.config_category
  display_name              = var.display_name
  type                      = var.type
  state                     = var.state

  filter {
    name   = "id"
    values = [oci_fleet_apps_management_platform_configuration.test_platform_configuration.id]
  }
}

#########################################################################################################
# Computed Attributes (Returned by Data Source)
#########################################################################################################

# id                         : Unique OCID of the Platform Configuration.
# compartment_id              : The compartment OCID where the resource exists.
# display_name                : The user-defined name of the configuration.
# description                 : Text description of the configuration purpose.
# resource_region             : Region where the configuration exists (e.g., us-ashburn-1).
# type                        : USER_DEFINED or SYSTEM_DEFINED.
# config_category_details     : Nested details for configuration type:
#    - config_category         : One of PRODUCT / PRODUCT_STACK / PRODUCT_STACK_AS_PRODUCT / PATCH_TYPE / CREDENTIAL / ENVIRONMENT.
#    - products                : List of product references (only for PRODUCT_STACK types).
#    - sub_category_details    : Nested sub-category information for PRODUCT_STACK_AS_PRODUCT.
#    - compatible_products     : Related product configurations by OCID.
#    - credentials             : Credential configurations associated.
#    - components              : Component names in the configuration.
#    - patch_types             : List of patch type OCIDs and names.
#    - versions                : Supported version identifiers.
# lifecycle_state             : ACTIVE / INACTIVE / DELETED.
# lifecycle_details           : Additional diagnostic or status message.
# time_created                : RFC3339 timestamp of creation.
# time_updated                : RFC3339 timestamp of last update.
# defined_tags                : Oracle-defined tags in namespaced format.
# freeform_tags               : User-defined key-value pairs.
# system_tags                 : System-assigned metadata tags.
#########################################################################################################
