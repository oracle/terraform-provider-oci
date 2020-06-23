---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_catalog_private_endpoints"
sidebar_current: "docs-oci-datasource-datacatalog-catalog_private_endpoints"
description: |-
  Provides the list of Catalog Private Endpoints in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_catalog_private_endpoints
This data source provides the list of Catalog Private Endpoints in Oracle Cloud Infrastructure Data Catalog service.

Returns a list of all the catalog private endpoints in the specified compartment.


## Example Usage

```hcl
data "oci_datacatalog_catalog_private_endpoints" "test_catalog_private_endpoints" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.catalog_private_endpoint_display_name}"
	state = "${var.catalog_private_endpoint_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment where you want to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the specified lifecycle state. The value is case insensitive.


## Attributes Reference

The following attributes are exported:

* `catalog_private_endpoints` - The list of catalog_private_endpoints.

### CatalogPrivateEndpoint Reference

The following attributes are exported:

* `attached_catalogs` - The list of catalogs using the private reverse connection endpoint
* `compartment_id` - Identifier of the compartment this private endpoint belongs to
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Mutable name of the Private Reverse Connection Endpoint
* `dns_zones` - List of DNS zones to be used by the data assets to be harvested. Example: custpvtsubnet.oraclevcn.com for data asset: db.custpvtsubnet.oraclevcn.com 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
* `state` - The current state of the private endpoint resource.
* `subnet_id` - Subnet Identifier
* `time_created` - The time the private endpoint was created. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.
* `time_updated` - The time the private endpoint was updated. An [RFC3339](https://tools.ietf.org/html/rfc3339) formatted datetime string.

