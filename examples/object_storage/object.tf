// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/* This example demonstrates object store object management. It uses Terraforms built-in `file` function to upload a file.
 * It also demonstrates the use of `local-exec` to download the object content using a preauthenticated request.
 *
 * WARNING: This should only be used with small files. The file helper does stringification so large files
 * may cause terraform to slow, become unresponsive or exceed allowed memory usage.
 */

resource "oci_objectstorage_object" "object1" {
  namespace           = data.oci_objectstorage_namespace.ns.namespace
  bucket              = oci_objectstorage_bucket.bucket1.name
  object              = "index.html"
  content_language    = "en-US"
  content_type        = "text/html"
  content             = file("index.html")
  content_disposition = "attachment; filename=\"filename.html\""
  storage_tier        = "InfrequentAccess"
  opc_sse_kms_key_id  = var.kms_key_ocid
}

resource "oci_objectstorage_object" "object_with_content_md5" {
  namespace           = data.oci_objectstorage_namespace.ns.namespace
  bucket              = oci_objectstorage_bucket.bucket1.name
  object              = "index.html"
  content_language    = "en-US"
  content_type        = "text/html"
  content             = file("index.html")
  content_disposition = "attachment; filename=\"filename.html\""
  storage_tier        = "InfrequentAccess"
  opc_sse_kms_key_id  = var.kms_key_ocid

  # we also support the base64 type of the content_md5, for example: G5uedo16+mhIOKSt4h/L2g==
  content_md5 = filemd5("index.html")
}

resource "oci_objectstorage_object" "source_object" {
  namespace           = data.oci_objectstorage_namespace.ns.namespace
  bucket              = oci_objectstorage_bucket.bucket1.name
  object              = "same_index.html"
  content_language    = "en-US"
  content_type        = "text/html"
  source              = oci_objectstorage_object.object1.object
  content_disposition = "attachment; filename=\"filename.html\""
  storage_tier        = "InfrequentAccess"
  opc_sse_kms_key_id  = var.kms_key_ocid
}

resource "oci_objectstorage_object" "source_uri_object" {
  namespace           = data.oci_objectstorage_namespace.ns.namespace
  bucket              = oci_objectstorage_bucket.bucket1.name
  object              = "copy_index.html"
  content_language    = "en-US"
  content_type        = "text/html"
  content_disposition = "attachment; filename=\"filename.html\""
  storage_tier        = "InfrequentAccess"
  opc_sse_kms_key_id  = var.kms_key_ocid

  source_uri_details {
    region    = local.source_region
    namespace = data.oci_objectstorage_namespace.ns.namespace
    bucket    = oci_objectstorage_bucket.bucket1.name
    object    = oci_objectstorage_object.object1.object
  }
}

resource "oci_objectstorage_object" "source_uri_object_from_version" {
  namespace           = data.oci_objectstorage_namespace.ns.namespace
  bucket              = oci_objectstorage_bucket.bucket_with_versioning.name
  object              = "copy_from_version_index.html"
  content_language    = "en-US"
  content_type        = "text/html"
  content_disposition = "attachment; filename=\"filename.html\""

  source_uri_details {
    region            = local.source_region
    namespace         = data.oci_objectstorage_namespace.ns.namespace
    bucket            = oci_objectstorage_bucket.bucket1.name
    object            = oci_objectstorage_object.object1.object
    source_version_id = oci_objectstorage_object.object1.version_id
  }

  delete_all_object_versions = true
}

data "oci_objectstorage_object_head" "object_head1" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  bucket    = oci_objectstorage_bucket.bucket1.name
  object    = oci_objectstorage_object.object1.object
}

data "oci_objectstorage_object_head" "source_object_head" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  bucket    = oci_objectstorage_bucket.bucket1.name
  object    = oci_objectstorage_object.source_object.object
}

data "oci_objectstorage_objects" "objects1" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  bucket    = oci_objectstorage_bucket.bucket1.name
}

data "oci_objectstorage_object" "object" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
  bucket    = oci_objectstorage_bucket.bucket1.name
  object    = oci_objectstorage_object.object1.object
}

output "object_data" {
  value = <<EOF

  content = ${data.oci_objectstorage_object.object.content}
  content-length = ${data.oci_objectstorage_object.object.content_length}
  content-type = ${data.oci_objectstorage_object.object.content_type}
EOF

}

output "object-head-data" {
  value = <<EOF

  object = ${data.oci_objectstorage_object_head.object_head1.object}
  content-length = ${data.oci_objectstorage_object_head.object_head1.content_length}
  content-type = ${data.oci_objectstorage_object_head.object_head1.content_type}
  storage-tier = ${data.oci_objectstorage_object_head.object_head1.storage_tier}
EOF

}

output "object-source-head-data" {
  value = <<EOF

  object = ${data.oci_objectstorage_object_head.source_object_head.object}
  content-length = ${data.oci_objectstorage_object_head.source_object_head.content_length}
  content-type = ${data.oci_objectstorage_object_head.source_object_head.content_type}
  storage-tier = ${data.oci_objectstorage_object_head.source_object_head.storage_tier}
EOF

}

output "objects" {
  value = data.oci_objectstorage_objects.objects1.objects
}

# example to download object content to local file using object PAR with local-exec

resource "null_resource" "download_object_content" {
  # using command
  # using command
  provisioner "local-exec" {
    command = "curl -o object-content https://objectstorage.${var.region}.oraclecloud.com${oci_objectstorage_preauthrequest.object_par.access_uri}"
  }

  # using script
  # using script
  provisioner "local-exec" {
    command = "sh get-content.sh https://objectstorage.${var.region}.oraclecloud.com${oci_objectstorage_preauthrequest.object_par.access_uri} object-content-using-script"
  }

  provisioner "local-exec" {
    when    = destroy
    command = "rm -rf object-content content/"
  }
}

