---
subcategory: "Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_analytics_analytics_instance_resource_group"
sidebar_current: "docs-oci-datasource-analytics-analytics_instance_resource_group"
description: |-
  Provides details about a specific Analytics Instance Resource Group in Oracle Cloud Infrastructure Analytics service
---

# Data Source: oci_analytics_analytics_instance_resource_group
This data source provides details about a specific Analytics Instance Resource Group resource in Oracle Cloud Infrastructure Analytics service.

Get details of a resource group for an instance


## Example Usage

```hcl
data "oci_analytics_analytics_instance_resource_group" "test_analytics_instance_resource_group" {
	#Required
	analytics_instance_id = oci_analytics_analytics_instance.test_analytics_instance.id
	analytics_instance_resource_group_id = oci_analytics_analytics_instance_resource_group.test_analytics_instance_resource_group.id
}
```

## Argument Reference

The following arguments are supported:

* `analytics_instance_id` - (Required) The OCID of the Analytics instance. 
* `analytics_instance_resource_group_id` - (Required) Specify unique id of a resource group within an Analytics instance. 


## Attributes Reference

The following attributes are exported:

* `capacity` - The capacity (in OCPU's) to be allocated for this resource.
* `description` - Optional description of the resource group
* `display_name` - Meaningful name of resource group for end user
* `id` - Unique identifier and name of resource group.  Must be unique within the instance
* `resource_name` - Meaningful name of resource group for end user

