/* This example demonstrates limited support for object store file upload using terraforms `file` 
 * helper method.
 * 
 * WARNING: This is not an appropriate solution for large files, this should only be used with small files.
 * The file helper does stringification so large files may cause terraform to slow, become unresponsive or 
 * exceed allowed memory usage.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" { default = "us-phoenix-1" }


provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

data "oci_objectstorage_namespace" "t" {
}

resource "oci_objectstorage_bucket" "t" {
  compartment_id = "${var.compartment_ocid}"
  namespace = "${data.oci_objectstorage_namespace.t.namespace}"
  name = "-tf-bucket"
  access_type="ObjectRead"
}

resource "oci_objectstorage_object" "t" {
  namespace = "${data.oci_objectstorage_namespace.t.namespace}"
  bucket = "${oci_objectstorage_bucket.t.name}"
  object = "-tf-object"
  content = "${file("data.txt")}"
  content_type = "text/plain"
  metadata = {
    "version" = "1"
  }
}
