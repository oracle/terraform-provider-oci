---
subcategory: "Data Connectivity"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_connectivity_registry_type"
sidebar_current: "docs-oci-datasource-data_connectivity-registry_type"
description: |-
  Provides details about a specific Registry Type in Oracle Cloud Infrastructure Data Connectivity service
---

# Data Source: oci_data_connectivity_registry_type
This data source provides details about a specific Registry Type resource in Oracle Cloud Infrastructure Data Connectivity service.

This endpoint retrieves dataAsset and connection attributes from DataAssetRegistry.


## Example Usage

```hcl
data "oci_data_connectivity_registry_type" "test_registry_type" {
	#Required
	registry_id = oci_data_connectivity_registry.test_registry.id
	type_key = var.registry_type_type_key

	#Optional
	fields = var.registry_type_fields
}
```

## Argument Reference

The following arguments are supported:

* `fields` - (Optional) Specifies the fields to get for an object.
* `registry_id` - (Required) The registry OCID.
* `type_key` - (Required) Key of the a specific type.


## Attributes Reference

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

