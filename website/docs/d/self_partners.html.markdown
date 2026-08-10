---
subcategory: "Self"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_self_partners"
sidebar_current: "docs-oci-datasource-self-partners"
description: |-
  Provides the list of Partners in Oracle Cloud Infrastructure Self service
---

# Data Source: oci_self_partners
This data source provides the list of Partners in Oracle Cloud Infrastructure Self service.

Lists marketplace publisher partner info for a compartment.


## Example Usage

```hcl
data "oci_self_partners" "test_partners" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.partner_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given name.


## Attributes Reference

The following attributes are exported:

* `partner_collection` - The list of partner_collection.

### Partner Reference

The following attributes are exported:

* `items` - The list of marketplace publisher partners.
	* `compartment_id` - The unique identifier of the compartment of partner. 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - The display name of marketplace publisher partner.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `id` - The unique identifier of the marketplace publisher partner. 
	* `state` - The current lifecycle state of the marketplace publisher partner.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 

