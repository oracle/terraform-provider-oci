// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

/*
 * This example shows how to manage a bucket with a replication policy
 */

variable "tenancy_ocid" {}

variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

variable "region" {
  default = "us-ashburn-1"
}

//the object can be created from the object data in the other region
locals {
  source_region = "${var.region}"
}

provider "oci" {
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

resource "oci_objectstorage_bucket" "bucket1" {
  compartment_id = "${var.compartment_ocid}"
  namespace      = "${data.oci_objectstorage_namespace.ns.namespace}"
  name           = "tf-example-source-bucket"
  access_type    = "NoPublicAccess"
}

resource "oci_objectstorage_bucket" "bucket2" {
  compartment_id = "${var.compartment_ocid}"
  namespace      = "${data.oci_objectstorage_namespace.ns.namespace}"
  name           = "tf-example-destination-bucket"
  access_type    = "NoPublicAccess"
}

data "oci_objectstorage_bucket_summaries" "buckets" {
  compartment_id = "${var.compartment_ocid}"
  namespace      = "${data.oci_objectstorage_namespace.ns.namespace}"

  filter {
    name   = "name"
    values = ["${oci_objectstorage_bucket.bucket1.name}", "${oci_objectstorage_bucket.bucket2.name}"]
  }
}

output buckets {
  value = "${data.oci_objectstorage_bucket_summaries.buckets.bucket_summaries}"
}

/*
 * This example file shows how to read and output the object storage namespace and namespace_metadata.
 */

data "oci_objectstorage_namespace" "ns" {
  #Optional
  compartment_id = "${var.compartment_ocid}"
}

output namespace {
  value = "${data.oci_objectstorage_namespace.ns.namespace}"
}

resource "oci_objectstorage_namespace_metadata" "namespace-metadata1" {
  namespace                    = "${data.oci_objectstorage_namespace.ns.namespace}"
  default_s3compartment_id     = "${var.compartment_ocid}"
  default_swift_compartment_id = "${var.compartment_ocid}"
}

data oci_objectstorage_namespace_metadata namespace-metadata1 {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
}

output namespace-metadata {
  value = <<EOF

  namespace = ${data.oci_objectstorage_namespace_metadata.namespace-metadata1.namespace}
  default_s3compartment_id = ${data.oci_objectstorage_namespace_metadata.namespace-metadata1.default_s3compartment_id}
  default_swift_compartment_id = ${data.oci_objectstorage_namespace_metadata.namespace-metadata1.default_swift_compartment_id}
EOF
}

resource "oci_objectstorage_object" "object1" {
  namespace           = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket              = "${oci_objectstorage_bucket.bucket1.name}"
  object              = "index.html"
  content_language    = "en-US"
  content_type        = "text/html"
  content             = "${file("index.html")}"
  content_disposition = "attachment; filename=\"filename.html\""
}

data "oci_objectstorage_objects" "objects1" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket    = "${oci_objectstorage_bucket.bucket1.name}"
}

data "oci_objectstorage_object" "object" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket    = "${oci_objectstorage_bucket.bucket1.name}"
  object    = "index.html"
}

output objects {
  value = "${data.oci_objectstorage_objects.objects1.objects}"
}

resource "oci_objectstorage_replication_policy" "bucket_rp" {
  namespace                           = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket                              = "${oci_objectstorage_bucket.bucket1.name}"
  name                                = "rpOnBucket"
  destination_region_name             = "${var.region}"
  delete_object_in_destination_bucket = "ACCEPT"
  destination_bucket_name             = "${oci_objectstorage_bucket.bucket2.name}"
}
