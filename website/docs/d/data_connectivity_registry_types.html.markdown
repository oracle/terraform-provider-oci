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

This endpoint retrieves list of all the supported connector types


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
* `registry_id` - (Required) The registry Ocid.
* `type` - (Optional) Type of the object to filter the results with.


## Attributes Reference

The following attributes are exported:

* `types_summary_collection` - The list of types_summary_collection.

### RegistryType Reference

The following attributes are exported:

* `connection_attributes` - Map of connectionType as key and List of attributes as value
* `data_asset_attributes` - list of attributes for the dataAsset
	* `attribute_type` - Attribute type details
	* `is_base64encoded` - True if Attribute is encoded.
	* `is_generated` - True if Attribute is generated.
	* `is_mandatory` - True if Attribute is mandatory.
	* `is_sensitive` - True if Attribute is sensitive.
	* `name` - The name of of the Attribute.
	* `valid_key_list` - List of valid key list

