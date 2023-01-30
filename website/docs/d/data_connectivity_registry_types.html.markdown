---
subcategory: "Data Connectivity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_connectivity_registry_types"
sidebar_current: "docs-oci-datasource-data_connectivity-registry_types"
description: |-
  Provides the list of Registry Types in Oracle Cloud Infrastructure Data Connectivity service
---

# Data Source: oci_data_connectivity_registry_types
This data source provides the list of Registry Types in Oracle Cloud Infrastructure Data Connectivity service.

This endpoint retrieves a list of all the supported connector types.


## Example Usage

```hcl
data "oci_data_connectivity_registry_types" "test_registry_types" {
	#Required
	registry_id = oci_data_connectivity_registry.test_registry.id

	#Optional
	name = var.registry_type_name
	type = var.registry_type_type
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Used to filter by the name of the object.
* `registry_id` - (Required) The registry OCID.
* `type` - (Optional) Type of the object to filter the results with.


## Attributes Reference

The following attributes are exported:

* `types_summary_collection` - The list of types_summary_collection.

### RegistryType Reference

The following attributes are exported:

* `connection_attributes` - Mapping the connectionType as the key to the list of attributes as the value.
* `data_asset_attributes` - The list of attributes of the data asset.
	* `attribute_type` - The attribute type details.
	* `is_base64encoded` - True if attribute is encoded.
	* `is_generated` - True if attribute is generated.
	* `is_mandatory` - True if attribute is mandatory.
	* `is_sensitive` - True if attribute is sensitive.
	* `name` - The name of of the attribute.
	* `valid_key_list` - The list of valid keys.

