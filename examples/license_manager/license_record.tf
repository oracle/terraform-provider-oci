variable "license_record_display_name" {
  default = "License Record"
}

variable "license_record_expiration_date" {
  default = "2199-06-30T23:59:59.000Z"
}

variable "license_record_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "license_record_is_perpetual" {
  default = false
}

variable "license_record_is_unlimited" {
  default = false
}

variable "license_record_license_count" {
  default = 10
}

variable "license_record_support_end_date" {
  default = "2199-06-30T23:59:59.000Z"
}

variable "product_id" {
  default = "1234"
}

resource "oci_license_manager_license_record" "test_license_record" {
  #Required
  display_name       = var.license_record_display_name
  is_perpetual       = var.license_record_is_perpetual
  is_unlimited       = var.license_record_is_unlimited
  product_license_id = oci_license_manager_product_license.test_product_license.id

  #Optional
  expiration_date  = var.license_record_expiration_date
  freeform_tags    = var.license_record_freeform_tags
  license_count    = var.license_record_license_count
  product_id       = var.product_id
  support_end_date = var.license_record_support_end_date
}

data "oci_license_manager_license_records" "test_license_records" {
  #Required
  product_license_id = oci_license_manager_product_license.test_product_license.id
}