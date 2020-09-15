---
subcategory: "Resource Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resourcemanager_stack_tf_state"
sidebar_current: "docs-oci-datasource-resourcemanager-stack_tf_state"
description: |-
  Provides details about a specific Stack Tf State in Oracle Cloud Infrastructure Resource Manager service
---

# Data Source: oci_resourcemanager_stack_tf_state
This data source provides details about a specific Stack Tf State resource in Oracle Cloud Infrastructure Resource Manager service.

Returns the Terraform state for the specified stack.

## Example Usage

```hcl
data "oci_resourcemanager_stack_tf_state" "test_stack_tf_state" {
	#Required
	stack_id = oci_resourcemanager_stack.test_stack.id
}
```

## Argument Reference

The following arguments are supported:

* `stack_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stack.


## Attributes Reference

The following attributes are exported:

* `local_path` - The path and filename (relative to where Terraform is executing) to write the external statefile to. 

