---
subcategory: "Waf"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waf_protection_capability_group_tags"
sidebar_current: "docs-oci-datasource-waf-protection_capability_group_tags"
description: |-
  Provides the list of Protection Capability Group Tags in Oracle Cloud Infrastructure Waf service
---

# Data Source: oci_waf_protection_capability_group_tags
This data source provides the list of Protection Capability Group Tags in Oracle Cloud Infrastructure Waf service.

Lists of available group tags filtered by query parameters.


## Example Usage

```hcl
data "oci_waf_protection_capability_group_tags" "test_protection_capability_group_tags" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.protection_capability_group_tag_name
	type = var.protection_capability_group_tag_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `name` - (Optional) A filter to return only resources that match the entire name given.
* `type` - (Optional) A filter to return only resources that matches given type.


## Attributes Reference

The following attributes are exported:

* `protection_capability_group_tag_collection` - The list of protection_capability_group_tag_collection.

### ProtectionCapabilityGroupTag Reference

The following attributes are exported:

* `items` - List of protection capabilities group tags.
	* `name` - Unique name of protection capability group tag.

