/* This example demonstrates object store object management. It uses Terraforms built-in `file` function to upload a file.
 * 
 * WARNING: This should only be used with small files. The file helper does stringification so large files
 * may cause terraform to slow, become unresponsive or exceed allowed memory usage.
 */

resource "oci_objectstorage_object" "object1" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket = "${oci_objectstorage_bucket.bucket1.name}"
  object = "index.html"
  content_language = "en-US"
  content_type = "text/html"
  content = "${file("index.html")}"
}

data "oci_objectstorage_object_head" "object-head1" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket = "${oci_objectstorage_bucket.bucket1.name}"
  object = "${oci_objectstorage_object.object1.object}"
}

data "oci_objectstorage_objects" "objects1" {
  namespace = "${data.oci_objectstorage_namespace.ns.namespace}"
  bucket = "${oci_objectstorage_bucket.bucket1.name}"
}

output object-head-data {
  value = <<EOF

  object = ${data.oci_objectstorage_object_head.object-head1.object}
  content-length = ${data.oci_objectstorage_object_head.object-head1.content-length}
  content-type = ${data.oci_objectstorage_object_head.object-head1.content-type}
EOF
}

output objects {
  value = "${data.oci_objectstorage_objects.objects1.objects}"
}
