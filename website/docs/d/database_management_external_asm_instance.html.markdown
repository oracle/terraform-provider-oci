---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_asm_instance"
sidebar_current: "docs-oci-datasource-database_management-external_asm_instance"
description: |-
  Provides details about a specific External Asm Instance in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_asm_instance
This data source provides details about a specific External Asm Instance resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the external ASM instance specified by `externalAsmInstanceId`.


## Example Usage

```hcl
data "oci_database_management_external_asm_instance" "test_external_asm_instance" {
	#Required
	external_asm_instance_id = oci_database_management_external_asm_instance.test_external_asm_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `external_asm_instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external ASM instance.


## Attributes Reference

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
* `time_created` - The date and time the external ASM instance was created.
* `time_updated` - The date and time the external ASM instance was last updated.

