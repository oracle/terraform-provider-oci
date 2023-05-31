---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_storage_server_iorm_plan"
sidebar_current: "docs-oci-datasource-database_management-external_exadata_storage_server_iorm_plan"
description: |-
  Provides details about a specific External Exadata Storage Server Iorm Plan in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_exadata_storage_server_iorm_plan
This data source provides details about a specific External Exadata Storage Server Iorm Plan resource in Oracle Cloud Infrastructure Database Management service.

Get the IORM plan from the specific Exadata storage server.


## Example Usage

```hcl
data "oci_database_management_external_exadata_storage_server_iorm_plan" "test_external_exadata_storage_server_iorm_plan" {
	#Required
	external_exadata_storage_server_id = oci_database_management_external_exadata_storage_server.test_external_exadata_storage_server.id
}
```

## Argument Reference

The following arguments are supported:

* `external_exadata_storage_server_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata storage server.


## Attributes Reference

The following attributes are exported:

* `db_plan` - The resource allocation directives must all use the share attribute, or they must all use the level and allocation attributes. If you use the share attribute to allocate I/O resources, then the database plan can have a maximum of 1024 directives. If you use the level and allocation attributes to allocate I/O resources, then the database plan can have a maximum of 32 directives. Only one directive is allowed for each database name and each profile name. 
	* `items` - A list of DatabasePlanDirectives.
		* `allocation` - The resource allocation as a percentage (0-100) within the level. 
		* `asm_cluster` - Starting with Oracle Exadata System Software release 19.1.0, you can use the asmcluster attribute to distinguish between databases with the same name running in different Oracle ASM clusters. 
		* `flash_cache_limit` - Defines a soft limit for space usage in Exadata Smart Flash Cache. If the cache is not full, the limit can be exceeded. You specify the value for flashcachelimit in bytes. You can also use the suffixes M (megabytes), G (gigabytes), or T (terabytes) to specify larger values. For example, 300M, 150G, or 1T. The value for flashcachelimit must be at least 4 MB. The flashcachelimit and flashcachesize attributes cannot be specified in the same directive. The value for flashcachelimit cannot be smaller than flashcachemin, if it is specified. 
		* `flash_cache_min` - Specifies a minimum guaranteed space allocation in Exadata Smart Flash Cache. You specify the value for flashcachemin in bytes. You can also use the suffixes M (megabytes), G (gigabytes), or T (terabytes) to specify larger values. For example, 300M, 150G, or 1T. The value for flashcachemin must be at least 4 MB. In any plan, the sum of all flashcachemin values cannot exceed the size of Exadata Smart Flash Cache. If flashcachelimit is specified, then the value for flashcachemin cannot exceed flashcachelimit. If flashcachesize is specified, then the value for flashcachemin cannot exceed flashcachesize. 
		* `flash_cache_size` - Defines a hard limit for space usage in Exadata Smart Flash Cache. The limit cannot be exceeded, even if the cache is not full. In an IORM plan, if the size of Exadata Smart Flash Cache can accommodate all of the flashcachemin and flashcachesize allocations, then each flashcachesize definition represents a guaranteed space allocation. However, starting with Oracle Exadata System Software release 19.2.0 you can use the flashcachesize attribute to over-provision space in Exadata Smart Flash Cache. Consequently, if the size of Exadata Smart Flash Cache cannot accommodate all of the flashcachemin and flashcachesize allocations, then only flashcachemin is guaranteed. 
		* `is_flash_cache_on` - Controls use of Exadata Smart Flash Cache by a database. This ensures that cache space is reserved for mission-critical databases. flashcache=off is invalid in a directive that contains the flashcachemin, flashcachelimit, or flashcachesize attributes. 
		* `is_flash_log_on` - Controls use of Exadata Smart Flash Log by a database. This ensures that Exadata Smart Flash Log is reserved for mission-critical databases. 
		* `is_pmem_cache_on` - Controls use of the persistent memory (PMEM) cache by a database. This ensures that cache space is reserved for mission-critical databases. pmemcache=off is invalid in a directive that contains the pmemcachemin, pmemcachelimit, or pmemcachesize attributes. 
		* `is_pmem_log_on` - Controls use of persistent memory logging (PMEM log) by a database. This ensures that PMEM log is reserved for mission-critical databases. 
		* `level` - The allocation level. Valid values are from 1 to 8. Resources are allocated to level 1 first, and then remaining resources are allocated to level 2, and so on. 
		* `name` - The name of a database or a profile.
		* `pmem_cache_limit` - Defines a soft limit for space usage in the persistent memory (PMEM) cache. If the cache is not full, the limit can be exceeded. You specify the value for pmemcachelimit in bytes. You can also use the suffixes M (megabytes), G (gigabytes), or T (terabytes) to specify larger values. For example, 300M, 150G, or 1T. The value for pmemcachelimit must be at least 4 MB. The pmemcachelimit and pmemcachesize attributes cannot be specified in the same directive. The value for pmemcachelimit cannot be smaller than pmemcachemin, if it is specified. 
		* `pmem_cache_min` - Specifies a minimum guaranteed space allocation in the persistent memory (PMEM) cache. 
		* `pmem_cache_size` - Defines a hard limit for space usage in the persistent memory (PMEM) cache. The limit cannot be exceeded, even if the cache is not full. In an IORM plan, if the size of the PMEM cache can accommodate all of the pmemcachemin and pmemcachesize allocations, then each pmemcachesize definition represents a guaranteed space allocation. However, you can use the pmemcachesize attribute to over-provision space in the PMEM cache. Consequently, if the PMEM cache size cannot accommodate all of the pmemcachemin and pmemcachesize allocations, then only pmemcachemin is guaranteed. 
		* `role` - Enables you to specify different plan directives based on the Oracle Data Guard database role. 
		* `share` - The relative priority of a database in the database plan. A higher share value implies higher priority and more access to the I/O resources. Use either share or (level, allocation). All plan directives in a database plan should use the same setting. Share-based resource allocation is the recommended method for a database plan. 
		* `type` - Enables you to create a profile or template, to ease management and configuration of resource plans in environments with many databases.
			* type=database: Specifies a directive that applies to a specific database. If type in not specified, then the directive defaults to the database type.
			* type=profile: Specifies a directive that applies to a profile rather than a specific database.

			To associate a database with an IORM profile, you must set the database initialization parameter db_performance_profile to the value of the profile name. Databases that map to a profile inherit the settings specified in the profile. 
* `plan_objective` - The objective of the IORM plan.
* `plan_status` - The status of the IORM plan.

