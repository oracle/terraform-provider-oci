---
subcategory: "Data Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datacatalog_catalog_private_endpoint"
sidebar_current: "docs-oci-datasource-datacatalog-catalog_private_endpoint"
description: |-
  Provides details about a specific Catalog Private Endpoint in Oracle Cloud Infrastructure Data Catalog service
---

# Data Source: oci_datacatalog_catalog_private_endpoint
This data source provides details about a specific Catalog Private Endpoint resource in Oracle Cloud Infrastructure Data Catalog service.

Gets a specific private reverse connection by identifier.

## Example Usage

```hcl
data "oci_datacatalog_catalog_private_endpoint" "test_catalog_private_endpoint" {
	#Required
	catalog_private_endpoint_id = oci_datacatalog_catalog_private_endpoint.test_catalog_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `catalog_private_endpoint_id` - (Required) Unique private reverse connection identifier.


## Attributes Reference

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

