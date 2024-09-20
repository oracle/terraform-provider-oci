---
subcategory: "Zpr"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_zpr_configuration"
sidebar_current: "docs-oci-datasource-zpr-configuration"
description: |-
  Provides details about a specific Configuration in Oracle Cloud Infrastructure Zpr service
---

# Data Source: oci_zpr_configuration
This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Zpr service.

Retrieves the ZPR configuration details for the root compartment (the root compartment is the tenancy).
Returns ZPR configuration for root compartment (the root compartment is the tenancy).


## Example Usage

```hcl
data "oci_zpr_configuration" "test_configuration" {
	#Required
	compartment_id = var.tenancy_ocid
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy into which ZPR will be onboarded. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ZprConfiguration.
* `lifecycle_details` - A message that describes the current state of ZPR in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `state` - The current state of ZPR in the tenancy.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that ZPR was onboarded to the tenancy, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time that ZPR was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `zpr_status` - The enabled or disabled status of ZPR in tenancy.

