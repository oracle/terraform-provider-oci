variable "bucket_name" {
  default = "testBucketName_1"
}

variable "bucket_namespace" {
  default = "ansh8lvru1zp"
}

data "oci_identity_tenancy" "test_tenancy" {
  #Required
  tenancy_id = var.tenancy_ocid
}

data "oci_objectstorage_namespace" "test_namespace" {

  #Optional
  compartment_id = var.compartment_id
}

data "oci_objectstorage_bucket" "test_bucket" {
  namespace = data.oci_objectstorage_namespace.test_namespace.namespace
  name      = var.bucket_name
}