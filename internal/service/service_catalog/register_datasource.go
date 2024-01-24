// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_catalog

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_service_catalog_private_application", ServiceCatalogPrivateApplicationDataSource())
	tfresource.RegisterDatasource("oci_service_catalog_private_application_package", ServiceCatalogPrivateApplicationPackageDataSource())
	tfresource.RegisterDatasource("oci_service_catalog_private_application_packages", ServiceCatalogPrivateApplicationPackagesDataSource())
	tfresource.RegisterDatasource("oci_service_catalog_private_applications", ServiceCatalogPrivateApplicationsDataSource())
	tfresource.RegisterDatasource("oci_service_catalog_service_catalog", ServiceCatalogServiceCatalogDataSource())
	tfresource.RegisterDatasource("oci_service_catalog_service_catalog_association", ServiceCatalogServiceCatalogAssociationDataSource())
	tfresource.RegisterDatasource("oci_service_catalog_service_catalog_associations", ServiceCatalogServiceCatalogAssociationsDataSource())
	tfresource.RegisterDatasource("oci_service_catalog_service_catalogs", ServiceCatalogServiceCatalogsDataSource())
}
