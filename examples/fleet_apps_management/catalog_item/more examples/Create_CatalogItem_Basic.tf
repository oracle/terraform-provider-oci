
variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1.."
}

variable "compartment_id" { default = "ocid1.compartment.oc1.." }

resource "oci_fleet_apps_management_catalog_item" "test_catalog_item" {
  catalog_source_payload {
    bucket             = "test-catalog-bucket"
    config_source_type = "PAR_CATALOG_SOURCE"
    namespace          = "namespace"
    object             = "ObjectStorageCatalog.zip"
  }
  compartment_id      = "${var.compartment_id}"
  config_source_type  = "PAR_CATALOG_SOURCE"
  description         = "description"
  display_name        = "displayName"
  package_type        = "TF_PACKAGE"
  short_description   = "shortDescription"
  version_description = "V1"
}
