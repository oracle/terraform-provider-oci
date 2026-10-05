// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci_product_catalog

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_oci_product_catalog_internal_admin_product", OciProductCatalogInternalAdminProductResource())
	tfresource.RegisterResource("oci_oci_product_catalog_internal_product", OciProductCatalogInternalProductResource())
}
