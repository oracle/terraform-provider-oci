// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example file shows how to read and output the object storage namespace and namespace_metadata.
 */

data "oci_objectstorage_namespace" "ns" {
  #Optional
  compartment_id = var.compartment_ocid
}

output "namespace" {
  value = data.oci_objectstorage_namespace.ns.namespace
}

resource "oci_objectstorage_namespace_metadata" "namespace-metadata1" {
  namespace                    = data.oci_objectstorage_namespace.ns.namespace
  default_s3compartment_id     = var.compartment_ocid
  default_swift_compartment_id = var.compartment_ocid
}

data "oci_objectstorage_namespace_metadata" "namespace-metadata1" {
  namespace = data.oci_objectstorage_namespace.ns.namespace
}

output "namespace-metadata" {
  value = <<EOF

  namespace = ${data.oci_objectstorage_namespace_metadata.namespace-metadata1.namespace}
  default_s3compartment_id = ${data.oci_objectstorage_namespace_metadata.namespace-metadata1.default_s3compartment_id}
  default_swift_compartment_id = ${data.oci_objectstorage_namespace_metadata.namespace-metadata1.default_swift_compartment_id}
EOF

}

