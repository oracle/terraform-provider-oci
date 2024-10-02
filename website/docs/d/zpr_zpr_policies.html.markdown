---
subcategory: "Zpr"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_zpr_zpr_policies"
sidebar_current: "docs-oci-datasource-zpr-zpr_policies"
description: |-
  Provides the list of Zpr Policies in Oracle Cloud Infrastructure Zpr service
---

# Data Source: oci_zpr_zpr_policies
This data source provides the list of Zpr Policies in Oracle Cloud Infrastructure Zpr service.

Gets a list of ZprPolicies.


## Example Usage

```hcl
data "oci_zpr_zpr_policies" "test_zpr_policies" {
	#Required
	compartment_id = var.tenancy_ocid

	#Optional
	name = var.zpr_policy_name
	state = var.zpr_policy_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `zpr_policies` - The list of zpr_policies.

### ZprPolicy Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description you assign to the ZprPolicy during creation. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `lifecycle_details` - A message that describes the current state of the ZprPolicy in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `name` - The name you assign to the ZprPolicy during creation. The name must be unique across all ZPL policies in the tenancy.
* `state` - The current state of the ZprPolicy.
* `statements` - An array of ZprPolicy statements (up to 25 statements per ZprPolicy) written in the Zero Trust Packet Routing Policy Language.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the ZprPolicy was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the ZprPolicy was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

