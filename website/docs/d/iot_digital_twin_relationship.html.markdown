---
subcategory: "Iot"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_iot_digital_twin_relationship"
sidebar_current: "docs-oci-datasource-iot-digital_twin_relationship"
description: |-
  Provides details about a specific Digital Twin Relationship in Oracle Cloud Infrastructure Iot service
---

# Data Source: oci_iot_digital_twin_relationship
This data source provides details about a specific Digital Twin Relationship resource in Oracle Cloud Infrastructure Iot service.

Retrieves the digital twin relationship identified by the specified OCID.

## Example Usage

```hcl
data "oci_iot_digital_twin_relationship" "test_digital_twin_relationship" {
	#Required
	digital_twin_relationship_id = oci_iot_digital_twin_relationship.test_digital_twin_relationship.id
}
```

## Argument Reference

The following arguments are supported:

* `digital_twin_relationship_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of digital twin relationship.


## Attributes Reference

The following attributes are exported:

* `content` - The value(s) of the relationship properties defined in the source digital twin model.
* `content_path` - Its the name of the relationship that links two digital twin instances. Here, it is the relationship name of the source digital twin model. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the resource. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
* `iot_domain_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
* `source_digital_twin_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of source digital twin instance. 
* `state` - The current state of the digital twin relationship.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_digital_twin_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of target digital twin instance.
* `time_created` - The date and time when the resource was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time when the resource was last updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

