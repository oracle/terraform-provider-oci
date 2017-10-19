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
variable "region" {}


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
}

resource "oci_objectstorage_object" "t" {
  namespace = "${data.oci_objectstorage_namespace.t.namespace}"
  bucket = "${oci_objectstorage_bucket.t.name}"
  object = "index.html"
  content_language = "en-US"
  content_type = "text/html"
  content = "${file("index.html")}"
}

resource "oci_objectstorage_preauthrequest" "par" {
  namespace = "${data.oci_objectstorage_namespace.t.namespace}"
  bucket = "${oci_objectstorage_bucket.t.name}"
  object = "${oci_objectstorage_object.t.object}"
  name = "par"
  access_type = "ObjectRead" // ObjectRead, ObjectWrite, ObjectReadWrite, AnyObjectWrite
  time_expires = "2019-11-10T23:00:00Z"
}

// Note: this will only output the full, usable url the first time.
output "par_request_url" {
  value = "https://objectstorage.${var.region}.oraclecloud.com${oci_objectstorage_preauthrequest.par.access_uri}"
}
