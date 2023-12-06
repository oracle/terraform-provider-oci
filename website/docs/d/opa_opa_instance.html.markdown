---
subcategory: "Opa"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opa_opa_instance"
sidebar_current: "docs-oci-datasource-opa-opa_instance"
description: |-
  Provides details about a specific Opa Instance in Oracle Cloud Infrastructure Opa service
---

# Data Source: oci_opa_opa_instance
This data source provides details about a specific Opa Instance resource in Oracle Cloud Infrastructure Opa service.

Gets a OpaInstance by identifier

## Example Usage

```hcl
data "oci_opa_opa_instance" "test_opa_instance" {
	#Required
	opa_instance_id = oci_opa_opa_instance.test_opa_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `opa_instance_id` - (Required) unique OpaInstance identifier


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

