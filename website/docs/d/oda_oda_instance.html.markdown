---
subcategory: "Oda"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_instance"
sidebar_current: "docs-oci-datasource-oda-oda_instance"
description: |-
  Provides details about a specific Oda Instance in Oracle Cloud Infrastructure Oda service
---

# Data Source: oci_oda_oda_instance
This data source provides details about a specific Oda Instance resource in Oracle Cloud Infrastructure Oda service.

Gets the specified Digital Assistant instance.

## Example Usage

```hcl
data "oci_oda_oda_instance" "test_oda_instance" {
	#Required
	oda_instance_id = "${oci_oda_oda_instance.test_oda_instance.id}"
}
```

## Argument Reference

The following arguments are supported:

* `oda_instance_id` - (Required) Unique Digital Assistant instance identifier.


## Attributes Reference

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

