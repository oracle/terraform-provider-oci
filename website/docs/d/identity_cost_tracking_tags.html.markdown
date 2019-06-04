---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_cost_tracking_tags"
sidebar_current: "docs-oci-datasource-identity-cost_tracking_tags"
description: |-
  Provides the list of Cost Tracking Tags in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_cost_tracking_tags
This data source provides the list of Cost Tracking Tags in Oracle Cloud Infrastructure Identity service.

Lists all the tags enabled for cost-tracking in the specified tenancy. For information about
cost-tracking tags, see [Using Cost-tracking Tags](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm#costs).


## Example Usage

```hcl
data "oci_identity_cost_tracking_tags" "test_cost_tracking_tags" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment (remember that the tenancy is simply the root compartment). 


## Attributes Reference

The following attributes are exported:

* `tags` - The list of tags.

### CostTrackingTag Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the tag definition.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`` 
* `description` - The description you assign to the tag.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the tag definition.
* `is_cost_tracking` - Indicates whether the tag is enabled for cost tracking. 
* `is_retired` - Indicates whether the tag is retired. See [Retiring Key Definitions and Namespace Definitions](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm#Retiring). 
* `name` - The name of the tag. The name must be unique across all tags in the namespace and can't be changed. 
* `state` - The tag's current state. After creating a tag, make sure its `lifecycleState` is ACTIVE before using it. After retiring a tag, make sure its `lifecycleState` is INACTIVE before using it.
* `tag_namespace_id` - The OCID of the namespace that contains the tag definition.
* `tag_namespace_name` - The name of the tag namespace that contains the tag definition. 
* `time_created` - Date and time the tag was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

