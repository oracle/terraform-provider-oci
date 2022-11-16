variable "bucket_name" {
  default = "testBucketName"
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

resource "oci_objectstorage_bucket" "test_bucket" {
  #Required
  compartment_id = var.compartment_id
  name           = var.bucket_name
  namespace      = data.oci_objectstorage_namespace.test_namespace.namespace
}
