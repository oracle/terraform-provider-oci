// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_objectstorage_object" "object_rd" {
  namespace           = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket              = "${oci_objectstorage_bucket.bucket_rd.name}"
  object              = "index.html"
  content_language    = "en-US"
  content_type        = "text/html"
  content             = "${file("resources/index.html")}"
  content_disposition = "attachment; filename=\"filename.html\""
}

data "oci_objectstorage_namespace" "ns" {
  #Optional
  compartment_id = "${var.compartment_ocid}"
}

resource "oci_objectstorage_bucket" "bucket_rd" {
  compartment_id = "${var.compartment_ocid}"
  namespace      = "${data.oci_objectstorage_namespace.ns.namespace}"
  name           = "tf-example-bucket-RD"
  access_type    = "NoPublicAccess"
}
