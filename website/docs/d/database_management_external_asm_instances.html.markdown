---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_asm_instances"
sidebar_current: "docs-oci-datasource-database_management-external_asm_instances"
description: |-
  Provides the list of External Asm Instances in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_asm_instances
This data source provides the list of External Asm Instances in Oracle Cloud Infrastructure Database Management service.

Lists the ASM instances in the specified external ASM.

## Example Usage

```hcl
data "oci_database_management_external_asm_instances" "test_external_asm_instances" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.external_asm_instance_display_name
	external_asm_id = oci_database_management_external_asm.test_external_asm.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.
* `external_asm_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM.


## Attributes Reference

The following attributes are exported:

* `external_asm_instance_collection` - The list of external_asm_instance_collection.

### ExternalAsmInstance Reference

The following attributes are exported:

* `adr_home_directory` - The Automatic Diagnostic Repository (ADR) home directory for the ASM instance.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `component_name` - The name of the external ASM instance.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the ASM instance. The name does not have to be unique.
* `external_asm_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM that the ASM instance belongs to.
* `external_db_node_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB node on which the ASM instance is running.
* `external_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system that the ASM instance is a part of.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `host_name` - The name of the host on which the ASM instance is running.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM instance.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the external ASM instance.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the external ASM instance was created.
* `time_updated` - The date and time the external ASM instance was last updated.

