variable "product_license_display_name" {
  default = "Oracle Database Enterprise Edition"
}

variable "product_license_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "product_license_is_vendor_oracle" {
  default = true
}

variable "product_license_license_unit" {
  default = "NAMED_USER_PLUS"
}

variable "product_license_vendor_name" {
  default = "Oracle"
}

resource "oci_license_manager_product_license" "test_product_license" {
  #Required
  compartment_id   = var.tenancy_ocid
  display_name     = var.product_license_display_name
  is_vendor_oracle = var.product_license_is_vendor_oracle
  license_unit     = var.product_license_license_unit

  #Optional
  freeform_tags = var.product_license_freeform_tags
  images {
    #Required
    listing_id      = data.oci_marketplace_listing.test_listing.listing_id
    package_version = data.oci_marketplace_listing.test_listing.default_package_version
  }
  vendor_name = var.product_license_vendor_name
}

data "oci_marketplace_listing" "test_listing" {
#  listing_id     = data.oci_marketplace_listings.test_listings.listings[0].id
  listing_id = "101747862"
}

data "oci_marketplace_listings" "test_listings" {
  category       = ["Analytics"]
  compartment_id = var.tenancy_ocid
}

data "oci_license_manager_product_licenses" "test_product_licenses" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  is_compartment_id_in_subtree = var.is_compartment_id_in_subtree
}