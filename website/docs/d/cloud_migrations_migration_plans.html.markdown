---
subcategory: "Cloud Migrations"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_migrations_migration_plans"
sidebar_current: "docs-oci-datasource-cloud_migrations-migration_plans"
description: |-
  Provides the list of Migration Plans in Oracle Cloud Infrastructure Cloud Migrations service
---

# Data Source: oci_cloud_migrations_migration_plans
This data source provides the list of Migration Plans in Oracle Cloud Infrastructure Cloud Migrations service.

Returns a list of migration plans.


## Example Usage

```hcl
data "oci_cloud_migrations_migration_plans" "test_migration_plans" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.migration_plan_display_name
	migration_id = oci_cloud_migrations_migration.test_migration.id
	migration_plan_id = oci_cloud_migrations_migration_plan.test_migration_plan.id
	state = var.migration_plan_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire given display name.
* `migration_id` - (Optional) Unique migration identifier
* `migration_plan_id` - (Optional) Unique migration plan identifier
* `state` - (Optional) The current state of the migration plan.


## Attributes Reference

The following attributes are exported:

* `migration_plan_collection` - The list of migration_plan_collection.

### MigrationPlan Reference

The following attributes are exported:

* `calculated_limits` - Limits of the resources that are needed for migration. Example: {"BlockVolume": 2, "VCN": 1}
* `compartment_id` - The OCID of the compartment containing the migration plan.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}` 
* `id` - The unique Oracle ID (OCID) that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
* `migration_id` - The OCID of the associated migration.
* `migration_plan_stats` - Status of the migration plan.
	* `time_updated` - The time when the migration plan was calculated. An RFC3339 formatted datetime string.
	* `total_estimated_cost` - Cost estimation description
		* `compute` - Cost estimation for compute
			* `gpu_count` - Total number of GPU
			* `gpu_per_hour` - GPU per hour
			* `gpu_per_hour_by_subscription` - GPU per hour by subscription
			* `memory_amount_gb` - Total usage of memory
			* `memory_gb_per_hour` - Gigabyte per hour
			* `memory_gb_per_hour_by_subscription` - Gigabyte per hour by subscription
			* `ocpu_count` - Total number of OCPUs
			* `ocpu_per_hour` - OCPU per hour
			* `ocpu_per_hour_by_subscription` - OCPU per hour by subscription
			* `total_per_hour` - Total per hour
			* `total_per_hour_by_subscription` - Total usage per hour by subscription
		* `currency_code` - Currency code in the ISO format.
		* `os_image` - Cost estimation for the OS image.
			* `total_per_hour` - Total price per hour
			* `total_per_hour_by_subscription` - Total price per hour by subscription
		* `storage` - Cost estimation for storage
			* `total_gb_per_month` - Gigabyte storage capacity per month.
			* `total_gb_per_month_by_subscription` - Gigabyte storage capacity per month by subscription.
			* `volumes` - Volume estimation
				* `capacity_gb` - Gigabyte storage capacity
				* `description` - Volume description
				* `total_gb_per_month` - Gigabyte storage capacity per month.
				* `total_gb_per_month_by_subscription` - Gigabyte storage capacity per month by subscription
		* `subscription_id` - Subscription ID
		* `total_estimation_per_month` - Total estimation per month
		* `total_estimation_per_month_by_subscription` - Total estimation per month by subscription.
	* `vm_count` - The total count of VMs in migration
* `reference_to_rms_stack` - OCID of the referenced ORM job.
* `source_migration_plan_id` - Source migraiton plan ID to be cloned.
* `state` - The current state of the migration plan.
* `strategies` - List of strategies for the resources to be migrated.
	* `adjustment_multiplier` - The real resource usage is multiplied to this number before making any recommendation.
	* `metric_time_window` - The current state of the migration plan.
	* `metric_type` - The current state of the migration plan.
	* `percentile` - Percentile value
	* `resource_type` - The type of resource.
	* `strategy_type` - The type of strategy used for migration.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_environments` - List of target environments.
	* `availability_domain` - Availability Domain of the VM configuration.
	* `dedicated_vm_host` - OCID of the dedicated VM configuration host.
	* `fault_domain` - Fault domain of the VM configuration.
	* `ms_license` - Microsoft license for the VM configuration.
	* `preferred_shape_type` - Preferred VM shape type provided by the customer.
	* `subnet` - OCID of the VM configuration subnet.
	* `target_compartment_id` - Target compartment identifier
	* `target_environment_type` - The type of target environment.
	* `vcn` - OCID of the VM configuration VCN.
* `time_created` - The time when the migration plan was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the migration plan was updated. An RFC3339 formatted datetime string.

