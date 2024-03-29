// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example file shows how to access information on the current tenancy.
 */

data "oci_identity_tenancy" "tenancy" {
  tenancy_id = var.tenancy_ocid
}

output "tenancy" {
  value = <<EOF

name            = ${data.oci_identity_tenancy.tenancy.name}
description     = ${data.oci_identity_tenancy.tenancy.description}
tenancy_id      = ${data.oci_identity_tenancy.tenancy.tenancy_id}
home_region_key = ${data.oci_identity_tenancy.tenancy.home_region_key}
EOF

}

