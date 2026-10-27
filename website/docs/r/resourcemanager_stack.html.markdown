---
subcategory: "Resource Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_resourcemanager_stack"
sidebar_current: "docs-oci-resource-resourcemanager-stack"
description: |-
  Provides the Stack resource in Oracle Cloud Infrastructure Resource Manager service
---

# oci_resourcemanager_stack

This resource provides the Stack resource in Oracle Cloud Infrastructure Resource Manager service.

Creates, updates, and deletes a stack.
For more information, see
[Managing Stacks](https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/stacks.htm).

The current implementation supports ZIP-upload configuration sources only.

Because Resource Manager does not return ZIP-upload contents from the standard `GetStack` API, this resource downloads the stack configuration archive during refresh and import and stores it in Terraform state. That preserves config fidelity, but large stack archives can noticeably increase state size.

## Example Usage

```hcl
resource "oci_resourcemanager_stack" "test_stack" {
  compartment_id = var.compartment_id
  display_name   = "example-stack"
  description    = "Stack managed by Terraform"

  config_source {
    config_source_type     = "ZIP_UPLOAD"
    zip_file_base64encoded = filebase64("${path.module}/stack.zip")
    working_directory      = "env/dev"
  }

  terraform_version = "1.5.x"
  variables = {
    compartment_ocid = var.compartment_id
  }

  freeform_tags = {
    Department = "Finance"
  }
}
```

For a runnable example layout, see `examples/resourcemanager/stack_resource`.

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this stack.
* `config_source` - (Required) The Terraform configuration source for the stack. Only ZIP-upload stacks are currently supported by this resource.
  * `config_source_type` - (Required) Configuration source type. `ZIP_UPLOAD` is currently supported.
  * `zip_file_base64encoded` - (Required) Base64-encoded ZIP archive containing the Terraform configuration.
  * `working_directory` - (Optional) File path to the directory from which Terraform runs. If not specified, the root directory is used.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `description` - (Optional) (Updatable) General description of the stack.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `terraform_version` - (Optional) (Updatable) The version of Terraform to use with the stack. Example: `1.5.x`
* `variables` - (Optional) (Updatable) Terraform variables associated with this resource. Maximum number of variables supported is 250. The maximum size of each variable, including both name and value, is 8192 bytes. Example: `{"CompartmentId": "compartment-id-value"}`

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Unique identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for the compartment where the stack is located.
* `config_source` - Location of the Terraform configuration.
  * `config_source_type` - Configuration source type.
  * `working_directory` - File path to the directory from which Terraform runs. If not specified, the root directory is used.
  * `zip_file_base64encoded` - For ZIP-upload stacks, the ZIP archive content fetched from the stack configuration download API and stored in Terraform state as base64. This is what allows refresh and import to preserve the stack configuration in state.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `description` - General description of the stack.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `id` - Unique identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for the stack.
* `stack_drift_status` - Drift status of the stack. Drift refers to differences between the actual (current) state of the stack and the expected (defined) state of the stack.
* `state` - The current lifecycle state of the stack.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}`
* `terraform_version` - The version of Terraform specified for the stack. Example: `1.5.x`
* `time_created` - The date and time at which the stack was created. Format is defined by RFC3339. Example: `2020-01-25T21:10:29.600Z`
* `time_drift_last_checked` - The date and time when the drift detection was last executed. Format is defined by RFC3339. Example: `2020-01-25T21:10:29.600Z`
* `variables` - Terraform variables associated with this resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
  * `create` - (Defaults to 20 minutes), when creating the Stack
  * `update` - (Defaults to 20 minutes), when updating the Stack
  * `delete` - (Defaults to 20 minutes), when destroying the Stack

## Import

Stacks can be imported using the `id`, e.g.

```bash
$ terraform import oci_resourcemanager_stack.test_stack "ocid1.ormstack.oc1..exampleuniqueID"
```
