---
subcategory: "Cloud Migrations"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_migrations_migration_plan_available_shapes"
sidebar_current: "docs-oci-datasource-cloud_migrations-migration_plan_available_shapes"
description: |-
  Provides the list of Migration Plan Available Shapes in Oracle Cloud Infrastructure Cloud Migrations service
---

# Data Source: oci_cloud_migrations_migration_plan_available_shapes
This data source provides the list of Migration Plan Available Shapes in Oracle Cloud Infrastructure Cloud Migrations service.

List of shapes by parameters.

## Example Usage

```hcl
data "oci_cloud_migrations_migration_plan_available_shapes" "test_migration_plan_available_shapes" {
	#Required
	migration_plan_id = oci_cloud_migrations_migration_plan.test_migration_plan.id

	#Optional
	availability_domain = var.migration_plan_available_shape_availability_domain
	compartment_id = var.compartment_id
	dvh_host_id = oci_cloud_migrations_dvh_host.test_dvh_host.id
	reserved_capacity_id = oci_cloud_migrations_reserved_capacity.test_reserved_capacity.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The availability domain in which to list resources.
* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `dvh_host_id` - (Optional) The ID of the Dvh in which to list resources.
* `migration_plan_id` - (Required) Unique migration plan identifier
* `reserved_capacity_id` - (Optional) The reserved capacity ID for which to list resources.


## Attributes Reference

The following attributes are exported:

* `available_shapes_collection` - The list of available_shapes_collection.

### MigrationPlanAvailableShape Reference

The following attributes are exported:

* `items` - Available shapes list.
    * `availability_domain` - Availability domain of the shape.
    * `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
    * `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. It exists only for cross-compatibility. Example: `{"bar-key": "value"}`
    * `gpu_description` - Description of the GPUs.
    * `gpus` - Number of GPUs.
    * `local_disk_description` - Description of local disks.
    * `local_disks` - Number of local disks.
    * `local_disks_total_size_in_gbs` - Total size of local disks for shape.
    * `max_vnic_attachments` - Maximum number of virtual network interfaces that can be attached.
    * `memory_in_gbs` - Amount of memory for the shape.
    * `min_total_baseline_ocpus_required` - Minimum CPUs required.
    * `networking_bandwidth_in_gbps` - Shape bandwidth.
    * `ocpus` - Number of CPUs.
    * `pagination_token` - Shape name and availability domain.  Used for pagination.
    * `processor_description` - Description of the processor.
    * `shape` - Name of the shape.
    * `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
