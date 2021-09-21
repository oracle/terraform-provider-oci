---
subcategory: "Waf"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waf_protection_capability_group_tag"
sidebar_current: "docs-oci-datasource-waf-protection_capability_group_tag"
description: |-
  Provides details about a specific Protection Capability Group Tag in Oracle Cloud Infrastructure Waf service
---

# Data Source: oci_waf_protection_capability_group_tag
This data source provides details about a specific Protection Capability Group Tag resource in Oracle Cloud Infrastructure Waf service.

Lists of available group tags filtered by query parameters.


## Example Usage

```hcl
data "oci_waf_protection_capability_group_tag" "test_protection_capability_group_tag" {
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

* `items` - List of protection capabilities group tags.
	* `name` - Unique name of protection capability group tag.

