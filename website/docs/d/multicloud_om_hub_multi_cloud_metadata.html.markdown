---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_om_hub_multi_cloud_metadata"
sidebar_current: "docs-oci-datasource-multicloud-om_hub_multi_cloud_metadata"
description: |-
  Provides details about a specific OmHub MultiCloud base compartment in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_om_hub_multi_clouds_metadata
This data source provides details about a specific OmHub MultiCloud base compartment in Oracle Cloud Infrastructure Multicloud service.

Gets information about multicloud base compartment

## Example Usage

```hcl
data "oci_multicloud_om_hub_multi_cloud_metadata" "test_om_hub_multi_cloud_metadata" {
	#Required
	compartment_id  = var.compartment_id
	subscription_id = var.subscription_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `subscription_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure subscription.


## Attributes Reference

The following attributes are exported:

* `base_compartment_id` - MultiCloud base compartment OCID associated with subscriptionId.
* `base_subscription_id` - Oracle Cloud Infrastructure subscriptionId.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the multicloud compartment was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`