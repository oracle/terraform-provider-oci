```

//resource not exposed to user through Terraform, but generated.
//Hence TF team suggested to keep the file commented as codeGen patch build fails if file not present
---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_resource_ports"
sidebar_current: "docs-oci-datasource-cloud_guard-resource_ports"
description: |-
  Provides the list of Resource Ports in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_resource_ports
This data source provides the list of Resource Ports in Oracle Cloud Infrastructure Cloud Guard service.

Returns the list of open ports associated with the resourceId where resource is an instance

## Example Usage

```hcl
data "oci_cloud_guard_resource_ports" "test_resource_ports" {
	#Required
	resource_id = oci_cloud_guard_resource.test_resource.id

	#Optional
	open_port = var.resource_port_open_port
}


## Argument Reference

The following arguments are supported:

* `open_port` - (Optional) open port associated with the resource.
* `resource_id` - (Required) CloudGuard resource OCID


## Attributes Reference

The following attributes are exported:

* `resource_port_collection` - The list of resource_port_collection.

### ResourcePort Reference

The following attributes are exported:

* `items` - List of CloudGuardResourcePortSummary
    * `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
    * `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

      Avoid entering confidential information.
    * `port_number` - The open port number
    * `process` - Process running on the open port
    * `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
    * `type` - Type of port
```