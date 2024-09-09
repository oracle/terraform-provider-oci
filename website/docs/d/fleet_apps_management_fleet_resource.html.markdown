---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_resource"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleet_resource"
description: |-
  Provides details about a specific Fleet Resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleet_resource
This data source provides details about a specific Fleet Resource resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a FleetResource by identifier

## Example Usage

```hcl
data "oci_fleet_apps_management_fleet_resource" "test_fleet_resource" {
	#Required
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
	fleet_resource_id = oci_cloud_guard_resource.test_resource.id
}
```

## Argument Reference

The following arguments are supported:

* `fleet_id` - (Required) unique Fleet identifier
* `fleet_resource_id` - (Required) unique FleetResource identifier


## Attributes Reference

The following attributes are exported:

* `application_type` - Application Type associated with the resource when resource type is fleet.Will only be returned for ENVIRONMENT fleets that are part of a GROUP Fleet.  
* `compartment` - Resource Compartment
* `compartment_id` - OCID of the compartment to which the resource belongs to.
* `compliance_state` - Compliance State of the Resource
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `environment_type` - Environment Type associated with the Fleet when resource type is fleet.Will only be returned for ENVIRONMENT fleets that are part of a GROUP Fleet. 
* `id` - The unique id of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `product` - Product associated with the resource when resource type is fleet.Will only be returned for PRODUCT fleets that are part of a GROUP Fleet
* `product_count` - Count of products within the resource.
* `resource_id` - The OCID of the resource.
* `resource_region` - Associated region
* `resource_type` - Type of the Resource.
* `state` - The current state of the FleetResource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_count` - Count of targets  within the resource.
* `tenancy_id` - OCID of the tenancy to which the resource belongs to.
* `tenancy_name` - Resource Tenancy Name
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

