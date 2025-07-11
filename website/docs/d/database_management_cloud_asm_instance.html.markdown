---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_asm_instance"
sidebar_current: "docs-oci-datasource-database_management-cloud_asm_instance"
description: |-
  Provides details about a specific Cloud Asm Instance in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_asm_instance
This data source provides details about a specific Cloud Asm Instance resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the cloud ASM instance specified by `cloudAsmInstanceId`.


## Example Usage

```hcl
data "oci_database_management_cloud_asm_instance" "test_cloud_asm_instance" {
	#Required
	cloud_asm_instance_id = oci_database_management_cloud_asm_instance.test_cloud_asm_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_asm_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM instance.


## Attributes Reference

The following attributes are exported:

* `adr_home_directory` - The Automatic Diagnostic Repository (ADR) home directory for the ASM instance.
* `cloud_asm_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM that the ASM instance belongs to.
* `cloud_db_node_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB node on which the ASM instance is running.
* `cloud_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the ASM instance is a part of.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the cloud ASM instance.
* `dbaas_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in DBaas service.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the ASM instance. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `host_name` - The name of the host on which the ASM instance is running.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM instance.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the cloud ASM instance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cloud ASM instance was created.
* `time_updated` - The date and time the cloud ASM instance was last updated.

