// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci_product_catalog

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_oci_product_catalog_internal_admin_product", OciProductCatalogInternalAdminProductDataSource())
	tfresource.RegisterDatasource("oci_oci_product_catalog_internal_admin_products", OciProductCatalogInternalAdminProductsDataSource())
	tfresource.RegisterDatasource("oci_oci_product_catalog_internal_product", OciProductCatalogInternalProductDataSource())
	tfresource.RegisterDatasource("oci_oci_product_catalog_internal_products", OciProductCatalogInternalProductsDataSource())
	tfresource.RegisterDatasource("oci_oci_product_catalog_product", OciProductCatalogProductDataSource())
	tfresource.RegisterDatasource("oci_oci_product_catalog_products", OciProductCatalogProductsDataSource())
}
