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

This endpoint retrieves dataAsset and connection attributes from DataAssetRegistry


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
* `registry_id` - (Required) The registry Ocid.
* `type_key` - (Required) key of the a specefic Type.


## Attributes Reference

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

