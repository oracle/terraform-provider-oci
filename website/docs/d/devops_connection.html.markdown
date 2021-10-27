---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_connection"
sidebar_current: "docs-oci-datasource-devops-connection"
description: |-
  Provides details about a specific Connection in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_connection
This data source provides details about a specific Connection resource in Oracle Cloud Infrastructure Devops service.

Gets a Connection by identifier

## Example Usage

```hcl
data "oci_devops_connection" "test_connection" {
	#Required
	connection_id = oci_devops_connection.test_connection.id
}
```

## Argument Reference

The following arguments are supported:

* `connection_id` - (Required) unique Connection identifier


## Attributes Reference

The following attributes are exported:

* `access_token` - OCID of personal access token saved in secret store
* `compartment_id` - Compartment Identifier
* `connection_type` - The type of connection.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `description` - Optional description about the connection
* `display_name` - Connection identifier which can be renamed and is not necessarily unique
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation
* `project_id` - Project Identifier
* `state` - The current state of the Connection.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the Connection was created. An RFC3339 formatted datetime string
* `time_updated` - The time at which the Connection was updated. An RFC3339 formatted datetime string

