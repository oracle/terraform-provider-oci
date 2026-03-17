---
subcategory: "Analytics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_analytics_analytics_instance_resource_groups"
sidebar_current: "docs-oci-datasource-analytics-analytics_instance_resource_groups"
description: |-
  Provides the list of Analytics Instance Resource Groups in Oracle Cloud Infrastructure Analytics service
---

# Data Source: oci_analytics_analytics_instance_resource_groups
This data source provides the list of Analytics Instance Resource Groups in Oracle Cloud Infrastructure Analytics service.

List resource groups associated with an instance.


## Example Usage

```hcl
data "oci_analytics_analytics_instance_resource_groups" "test_analytics_instance_resource_groups" {
	#Required
	analytics_instance_id = oci_analytics_analytics_instance.test_analytics_instance.id

	#Optional
	name = var.analytics_instance_resource_group_name
}
```

## Argument Reference

The following arguments are supported:

* `analytics_instance_id` - (Required) The OCID of the Analytics instance. 
* `name` - (Optional) A filter to return only resources that match the given name exactly. 


## Attributes Reference

The following attributes are exported:

* `instance_resource_groups` - The list of instance_resource_groups.

### AnalyticsInstanceResourceGroup Reference

The following attributes are exported:

* `capacity` - The capacity (in OCPU's) to be allocated for this resource.
* `description` - Optional description of the resource group
* `display_name` - Meaningful name of resource group for end user
* `id` - Unique identifier and name of resource group.  Must be unique within the instance
* `resource_name` - Meaningful name of resource group for end user

