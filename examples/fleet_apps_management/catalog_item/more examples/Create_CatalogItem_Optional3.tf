
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }
variable "compartment_id_for_update" { default = "ocid1.compartment.oc1.." }

resource "oci_fleet_apps_management_catalog_item" "test_catalog_item" {
  catalog_source_payload {
    access_uri         = ""
    bucket             = "test-catalog-bucket"
    config_source_type = "PAR_CATALOG_SOURCE"
    namespace          = "namespace"
    object             = "ObjectStorageCatalog.zip"
    time_expires       = "2029-12-31T00:00:00Z"
    working_directory  = "workingDirectory"
  }
  compartment_id     = "${var.compartment_id_for_update}"
  config_source_type = "PAR_CATALOG_SOURCE"
  defined_tags       = "${map("Oracle-Tags.CreatedBy", "value")}"
  description        = "description"
  display_name       = "displayName"
  freeform_tags = {
    "bar-key" = "value"
  }
  package_type        = "TF_PACKAGE"
  short_description   = "shortDescription"
  time_released       = "2025-10-27T00:00:00.000Z"
  version_description = "V1"
}
