// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_marketplace_accepted_agreement", MarketplaceAcceptedAgreementDataSource())
	tfresource.RegisterDatasource("oci_marketplace_accepted_agreements", MarketplaceAcceptedAgreementsDataSource())
	tfresource.RegisterDatasource("oci_marketplace_categories", MarketplaceCategoriesDataSource())
	tfresource.RegisterDatasource("oci_marketplace_listing", MarketplaceListingDataSource())
	tfresource.RegisterDatasource("oci_marketplace_listing_package", MarketplaceListingPackageDataSource())
	tfresource.RegisterDatasource("oci_marketplace_listing_package_agreements", MarketplaceListingPackageAgreementsDataSource())
	tfresource.RegisterDatasource("oci_marketplace_listing_packages", MarketplaceListingPackagesDataSource())
	tfresource.RegisterDatasource("oci_marketplace_listing_taxes", MarketplaceListingTaxesDataSource())
	tfresource.RegisterDatasource("oci_marketplace_listings", MarketplaceListingsDataSource())
	tfresource.RegisterDatasource("oci_marketplace_publication", MarketplacePublicationDataSource())
	tfresource.RegisterDatasource("oci_marketplace_publication_package", MarketplacePublicationPackageDataSource())
	tfresource.RegisterDatasource("oci_marketplace_publication_packages", MarketplacePublicationPackagesDataSource())
	tfresource.RegisterDatasource("oci_marketplace_publications", MarketplacePublicationsDataSource())
	tfresource.RegisterDatasource("oci_marketplace_publishers", MarketplacePublishersDataSource())
}
