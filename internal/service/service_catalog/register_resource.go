// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_catalog

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_service_catalog_private_application", ServiceCatalogPrivateApplicationResource())
	tfresource.RegisterResource("oci_service_catalog_service_catalog", ServiceCatalogServiceCatalogResource())
	tfresource.RegisterResource("oci_service_catalog_service_catalog_association", ServiceCatalogServiceCatalogAssociationResource())
}
