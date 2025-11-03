// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


#########################################################################################################
# Variables for OCI Fleet Apps Management Catalog Item
#########################################################################################################

variable "tenancy_ocid" {
  type        = string
  description = "The OCID of the tenancy. Example: ocid1.tenancy.oc1..<unique_id>"
}

variable "compartment_id" {
  type        = string
  description = "The OCID of the compartment where the catalog item will be created."
}

variable "access_uri" {
  type        = string
  description = "Pre-authenticated request (PAR) URL for accessing the catalog object in Object Storage. Optional."
  default     = null
}

variable "bucket" {
  type        = string
  description = "The name of the Object Storage bucket containing the catalog ZIP object."
}

variable "config_source_type" {
  type        = string
  description = <<EOT
Type of the catalog source. Determines where the catalog is sourced from.
Possible values:
  - PAR_CATALOG_SOURCE : Catalog source is a pre-authenticated request URL.
  - OBJECT_STORAGE_BUCKET : Catalog source is an Object Storage bucket directly.
  - GIT : Catalog source is a Git repository.
  - MANUAL_UPLOAD : Catalog is manually uploaded through console or API.
EOT
  default     = "PAR_CATALOG_SOURCE"
}

variable "namespace" {
  type        = string
  description = "The Object Storage namespace name for the tenancy."
}

variable "object" {
  type        = string
  description = "The name of the ZIP file uploaded to Object Storage that represents the catalog (e.g., CreateObjectStorageTFCatalog.zip)."
}

variable "time_expires" {
  type        = string
  description = "Expiration timestamp for the catalog item validity. Format: YYYY-MM-DDThh:mm:ssZ (e.g., 2029-12-31T00:00:00Z). Optional."
  default     = null
}

variable "working_directory" {
  type        = string
  description = "Relative path within the ZIP archive that contains Terraform configuration files. Use null if files are in the root."
  default     = null
}

variable "defined_tags" {
  type        = map(string)
  description = "Defined tags for the resource, following OCI tag namespace conventions. Example: { 'Oracle-Tags.CreatedBy' = 'user' }"
  default     = {}
}

variable "description" {
  type        = string
  description = "Detailed description of the catalog item."
  default     = null
}

variable "display_name" {
  type        = string
  description = "Display name of the catalog item."
}

variable "freeform_tags" {
  type        = map(string)
  description = "Simple key-value pairs applied to the resource without predefined namespacing. Example: { 'Project' = 'TestApp' }"
  default     = {}
}

variable "package_type" {
  type        = string
  description = <<EOT
Type of package used in the catalog.
Possible values:
  - TF_PACKAGE : Catalog package containing Terraform configuration.
  - CONFIG_FILE : Catalog containing non-Terraform configuration files (e.g., JSON/YAML).
  - NON_TF_PACKAGE : Catalog package that does not conform to Terraform or config standards.
EOT
  default     = "TF_PACKAGE"
}

variable "short_description" {
  type        = string
  description = "Short summary or overview of the catalog item."
  default     = null
}

variable "time_released" {
  type        = string
  description = "Timestamp indicating when this catalog version was released. Format: YYYY-MM-DDThh:mm:ss.sssZ (e.g., 2025-10-27T00:00:00.000Z)."
}

variable "version_description" {
  type        = string
  description = "Version notes or changelog information for this catalog item (e.g., Initial Version, v1, etc.)."
  default     = "V1"
}

#########################################################################################################
# Resource Definition
#########################################################################################################

resource "oci_fleet_apps_management_catalog_item" "test_catalog_item" {
  compartment_id      = var.compartment_id
  config_source_type  = var.config_source_type
  description         = var.description
  display_name        = var.display_name
  defined_tags        = { "Oracle-Tags.CreatedBy" = "user" }
  freeform_tags       = { "Project" = "SampleCatalog" }
  package_type        = var.package_type
  short_description   = var.short_description
  time_released       = var.time_released
  version_description = var.version_description

  catalog_source_payload {
    access_uri         = var.access_uri
    bucket             = var.bucket
    config_source_type = var.config_source_type
    namespace          = var.namespace
    object             = var.object
    time_expires       = var.time_expires
    working_directory  = var.working_directory
  }
}


#########################################################################################################
# Singular Data Definition
#########################################################################################################
data "oci_fleet_apps_management_catalog_item" "test_catalog_item" {
  catalog_item_id = "${oci_fleet_apps_management_catalog_item.test_catalog_item.id}"
}


#########################################################################################################
# List Data Definition
#########################################################################################################
data "oci_fleet_apps_management_catalog_items" "test_catalog_items" {
  catalog_listing_version_criteria = "LIST_ALL_VERSIONS"
  compartment_id                   = "${var.compartment_id}"
  config_source_type               = var.config_source_type
  display_name                     = var.display_name
  filter {
    name   = "id"
    values = ["${oci_fleet_apps_management_catalog_item.test_catalog_item.id}"]
  }
  package_type             = var.package_type
  should_list_public_items = "false"
  state                    = "ACTIVE"
}


#########################################################################################################
# Computed Attributes (Available in Data Source)
#########################################################################################################
# The following attributes are automatically populated after creation or when retrieved via
# the data source `oci_fleet_apps_management_catalog_item`:
#
# - id                          : Unique OCID of the Catalog Item.
# - compartment_id               : OCID of the compartment where the Catalog Item resides.
# - config_source_type           : Type of catalog source configuration.
#       Possible values:
#         - PAR_CATALOG_SOURCE   : Source from a pre-authenticated request (PAR) URL.
#         - OBJECT_STORAGE_BUCKET: Source directly from Object Storage bucket.
#         - GIT                  : Source from a Git repository.
#         - MANUAL_UPLOAD        : Catalog manually uploaded via console or API.
# - should_list_public_items     : Boolean flag indicating whether public catalog items are visible.
# - description                  : Detailed description of the catalog item.
#
# - catalog_source_payload       : Object describing the catalog’s source configuration.
#       - config_source_type     : Same as top-level config_source_type.
#       - working_directory      : Directory path inside the ZIP archive (if applicable).
#       - namespace_name         : Object Storage namespace name.
#       - bucket_name            : Name of the Object Storage bucket.
#       - object_name            : Name of the object/ZIP file (e.g., ObjectStorageCatalog.zip).
#       - access_uri             : Pre-authenticated request (PAR) URL, if applicable.
#       - time_expires           : Expiration timestamp of the source (RFC3339 format).
#
# - catalog_result_payload       : Object describing result metadata generated during catalog validation.
#       - config_result_type     : Type of catalog result (e.g., PAR_RESULT_CONFIG).
#       - working_directory      : Path within the ZIP used during result processing.
#       - package_url            : URL of the validated Object Storage package.
#       - time_expires           : Expiration timestamp of the result (if applicable).
#
# - listing_id                   : Internal identifier or marketplace reference ID for the catalog.
# - listing_version              : Version string of the catalog listing (e.g., 1.0.0).
# - display_name                 : Display name of the catalog item.
# - version_description          : Version details or notes (e.g., V2, Initial Version).
# - short_description            : Short textual summary of the catalog.
# - package_type                 : Type of catalog package.
#       Possible values:
#         - TF_PACKAGE            : Catalog package with Terraform configuration.
#         - CONFIG_FILE           : Catalog package containing configuration files (non-Terraform).
#         - NON_TF_PACKAGE        : Other non-Terraform catalog package.
#
# - lifecycle_state              : Current lifecycle of the catalog item.
#       Possible values: ACTIVE / INACTIVE / FAILED / DELETED
# - lifecycle_details            : JSON-encoded diagnostic or validation information.
#
# - time_created                 : RFC3339 timestamp when the catalog item was created.
# - time_updated                 : RFC3339 timestamp when the catalog item was last updated.
# - time_released                : Timestamp when this catalog version was released.
# - time_backfill_last_checked   : Timestamp of the last schema backfill check.
# - time_last_checked            : Timestamp when the catalog was last validated.
#
# - is_item_locked               : Boolean indicating whether the catalog item is locked for modification.
# - freeform_tags                : Map of user-defined key-value tags.
#       Example: { "Department" = "Accounting" }
# - defined_tags                 : Map of defined tag namespaces and key-value pairs.
#       Example: { "Oracle-Tags" = { "CreatedBy" = "user", "CreatedOn" = "2025-10-29T12:41:15Z" } }
# - system_tags                  : Map of system-generated tags (if any).
#########################################################################################################
