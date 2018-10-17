---
layout: "oci"
page_title: "OCI: oci_core_app_catalog_listing_resource_version_agreement"
sidebar_current: "docs-oci-datasource-core-app_catalog_listing_resource_version_agreement"
description: |-
  Provides details about a specific AppCatalogListingResourceVersionAgreement
---

#oci_core_app_catalog_listing_resource_version_agreement
The `oci_core_app_catalog_listing_resource_version_agreement` resource creates AppCatalogListingResourceVersionAgreement for a particular resource version of a listing.

## Example Usage

```hcl
resource "oci_core_app_catalog_listing_resource_version_agreement" "test_app_catalog_listing_resource_version_agreement" {
	#Required
	listing_id = "${oci_core_listing.test_listing.id}"
	listing_resource_version = "${var.app_catalog_listing_resource_version_agreement_listing_resource_version}"
}
```

## Argument Reference

The following arguments are supported:

* `listing_id` - (Required) The OCID of the listing.
* `listing_resource_version` - (Required) Listing Resource Version.


## Attributes Reference

The following attributes are exported:

* `eula_link` - EULA link
* `listing_id` - The OCID of the listing associated with these agreements.
* `listing_resource_version` - Listing resource version associated with these agreements.
* `oracle_terms_of_use_link` - Oracle TOU link
* `signature` - A generated signature for this agreement retrieval operation which should be used in the create subscription call. 
* `time_retrieved` - Date and time the agreements were retrieved, in RFC3339 format. Example: `2018-03-20T12:32:53.532Z` 

