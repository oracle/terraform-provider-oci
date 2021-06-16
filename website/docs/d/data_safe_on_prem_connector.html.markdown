---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_on_prem_connector"
sidebar_current: "docs-oci-datasource-data_safe-on_prem_connector"
description: |-
  Provides details about a specific On Prem Connector in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_on_prem_connector
This data source provides details about a specific On Prem Connector resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified on-premises connector.

## Example Usage

```hcl
data "oci_data_safe_on_prem_connector" "test_on_prem_connector" {
	#Required
	on_prem_connector_id = oci_data_safe_on_prem_connector.test_on_prem_connector.id
}
```

## Argument Reference

The following arguments are supported:

* `on_prem_connector_id` - (Required) The OCID of the on-premises connector.


## Attributes Reference

The following attributes are exported:

* `available_version` - Latest available version of the on-premises connector.
* `compartment_id` - The OCID of the compartment that contains the on-premises connector.
* `created_version` - Created version of the on-premises connector.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the on-premises connector.
* `display_name` - The display name of the on-premises connector.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the on-premises connector.
* `lifecycle_details` - Details about the current state of the on-premises connector.
* `state` - The current state of the on-premises connector.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the on-premises connector was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

