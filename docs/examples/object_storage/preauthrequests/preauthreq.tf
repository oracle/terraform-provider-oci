#     ___  ____     _    ____ _     _____
#    / _ \|  _ \   / \  / ___| |   | ____|
#   | | | | |_) | / _ \| |   | |   |  _|
#   | |_| |  _ < / ___ | |___| |___| |___
#    \___/|_| \_/_/   \_\____|_____|_____|

## PAR(preauthenticated requests) management.
## This example creates two preauthenticated requests, one on a bucket and the second on an object
## Export the environmental variables used for configuration.
## Provide the script with an object and bucket present in your namespace

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

variable "bucket_name" {
  default = "-tf-bucket"
}

variable "object_name" {
  default = "-tf-object"
}

provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}

data "oci_objectstorage_namespace" "ns" {
}

resource "oci_objectstorage_bucket" "ExampleBucket" {
  compartment_id = "${var.compartment_ocid}"
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  name = "${var.bucket_name}"
}

resource "oci_objectstorage_object" "ExampleObject" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket = "${oci_objectstorage_bucket.ExampleBucket.name}"
  object = "${var.object_name}"
  content = "This is some example text content!"
}

resource "oci_objectstorage_preauthrequest" "parOnBucket" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket = "${oci_objectstorage_bucket.ExampleBucket.name}"
  name = "parOnBucket"
  access_type = "AnyObjectWrite" //Other configurations accepted are ObjectWrite, ObjectReadWrite
  time_expires = "2019-11-10T23:00:00Z"
}

resource "oci_objectstorage_preauthrequest" "parOnObject" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket = "${oci_objectstorage_bucket.ExampleBucket.name}"
  object = "${oci_objectstorage_object.ExampleObject.object}"
  name = "parOnObject"
  access_type = "ObjectRead" //Other configurations accepted are ObjectWrite, ObjectReadWrite
  time_expires = "2019-11-10T23:00:00Z"
}

// Note: this will only output the full, usable url the first time.
output "par_request_url" {
  value = "https://objectstorage.${var.region}.oraclecloud.com${oci_objectstorage_preauthrequest.parOnObject.access_uri}"
}
