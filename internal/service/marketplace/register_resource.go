// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_marketplace_accepted_agreement", MarketplaceAcceptedAgreementResource())
	tfresource.RegisterResource("oci_marketplace_publication", MarketplacePublicationResource())
	tfresource.RegisterResource("oci_marketplace_listing_package_agreement", MarketplaceListingPackageAgreementResource())
}
