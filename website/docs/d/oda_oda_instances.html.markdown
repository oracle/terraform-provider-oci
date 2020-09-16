---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_instances"
sidebar_current: "docs-oci-datasource-oda-oda_instances"
description: |-
  Provides the list of Oda Instances in Oracle Cloud Infrastructure Digital Assistant service
---

# Data Source: oci_oda_oda_instances
This data source provides the list of Oda Instances in Oracle Cloud Infrastructure Digital Assistant service.

Returns a page of Digital Assistant instances that belong to the specified
compartment.

If the `opc-next-page` header appears in the response, then
there are more items to retrieve. To get the next page in the subsequent
GET request, include the header's value as the `page` query parameter.


## Example Usage

```hcl
data "oci_oda_oda_instances" "test_oda_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.oda_instance_display_name
	state = var.oda_instance_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) List the Digital Assistant instances that belong to this compartment.
* `display_name` - (Optional) List only the information for the Digital Assistant instance with this user-friendly name. These names don't have to be unique and may change.  Example: `My new resource` 
* `state` - (Optional) List only the Digital Assistant instances that are in this lifecycle state.


## Attributes Reference

The following attributes are exported:

* `oda_instances` - The list of oda_instances.

### OdaInstance Reference

The following attributes are exported:

* `compartment_id` - Identifier of the compartment that the instance belongs to.
* `connector_url` - URL for the connector's endpoint.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the Digital Assistant instance.
* `display_name` - User-defined name for the Digital Assistant instance. Avoid entering confidential information. You can change this value. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique immutable identifier that was assigned when the instance was created.
* `lifecycle_sub_state` - The current sub-state of the Digital Assistant instance.
* `shape_name` - Shape or size of the instance.
* `state` - The current state of the Digital Assistant instance.
* `state_message` - A message that describes the current state in more detail. For example, actionable information about an instance that's in the `FAILED` state. 
* `time_created` - When the Digital Assistant instance was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `time_updated` - When the Digital Assistance instance was last updated. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `web_app_url` - URL for the Digital Assistant web application that's associated with the instance.

