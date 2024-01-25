---
subcategory: "Opa"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opa_opa_instance"
sidebar_current: "docs-oci-resource-opa-opa_instance"
description: |-
  Provides the Opa Instance resource in Oracle Cloud Infrastructure Opa service
---

# oci_opa_opa_instance
This resource provides the Opa Instance resource in Oracle Cloud Infrastructure Opa service.

Creates a new OpaInstance.


## Example Usage

```hcl
resource "oci_opa_opa_instance" "test_opa_instance" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.opa_instance_display_name
	shape_name = oci_core_shape.test_shape.name

	#Optional
	consumption_model = var.opa_instance_consumption_model
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.opa_instance_description
	freeform_tags = {"bar-key"= "value"}
	idcs_at = var.opa_instance_idcs_at
	is_breakglass_enabled = var.opa_instance_is_breakglass_enabled
	metering_type = var.opa_instance_metering_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier
* `consumption_model` - (Optional) Parameter specifying which entitlement to use for billing purposes
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Description of the Oracle Process Automation instance.
* `display_name` - (Required) (Updatable) OpaInstance Identifier. User-friendly name for the instance. Avoid entering confidential information. You can change this value anytime.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `idcs_at` - (Optional) IDCS Authentication token. This is required for all realms with IDCS. This property is optional, as it is not required for non-IDCS realms.
* `is_breakglass_enabled` - (Optional) indicates if breakGlass is enabled for the opa instance.
* `metering_type` - (Optional) MeteringType Identifier
* `shape_name` - (Required) Shape of the instance.
* `state` - (Optional) (Updatable) The target state for the Opa Instance. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `attachments` - A list of associated attachments to other services 
	* `is_implicit` - 
		* If role == `PARENT`, the attached instance was created by this service instance 
		* If role == `CHILD`, this instance was created from attached instance on behalf of a user 
	* `target_id` - The OCID of the target instance (which could be any other Oracle Cloud Infrastructure PaaS/SaaS resource), to which this instance is attached.
	* `target_instance_url` - The dataplane instance URL of the attached instance
	* `target_role` - The role of the target attachment. 
		* `PARENT` - The target instance is the parent of this attachment. 
		* `CHILD` - The target instance is the child of this attachment. 
	* `target_service_type` - The type of the target instance, such as "FUSION".
* `compartment_id` - Compartment Identifier
* `consumption_model` - The entitlement used for billing purposes
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the Process Automation instance.
* `display_name` - OpaInstance Identifier, can be renamed
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `identity_app_display_name` - This property specifies the name of the Identity Application instance OPA has created inside the user-specified identity domain. This identity application instance may be used to host user roll mappings to grant access to this OPA instance for users within the identity domain.
* `identity_app_guid` - This property specifies the GUID of the Identity Application instance OPA has created inside the user-specified identity domain. This identity application instance may be used to host user role mappings to grant access to this OPA instance for users within the identity domain.
* `identity_app_opc_service_instance_guid` - This property specifies the OPC Service Instance GUID of the Identity Application instance OPA has created inside the user-specified identity domain. This identity application instance may be used to host user roll mappings to grant access to this OPA instance for users within the identity domain.
* `identity_domain_url` - This property specifies the domain url of the Identity Application instance OPA has created inside the user-specified identity domain. This identity application instance may be used to host user roll mappings to grant access to this OPA instance for users within the identity domain.
* `instance_url` - OPA Instance URL
* `is_breakglass_enabled` - indicates if breakGlass is enabled for the opa instance.
* `metering_type` - MeteringType Identifier
* `shape_name` - Shape of the instance.
* `state` - The current state of the OpaInstance.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when OpaInstance was created. An RFC3339 formatted datetime string
* `time_updated` - The time the OpaInstance was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Opa Instance
	* `update` - (Defaults to 20 minutes), when updating the Opa Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Opa Instance


## Import

OpaInstances can be imported using the `id`, e.g.

```
$ terraform import oci_opa_opa_instance.test_opa_instance "id"
```

