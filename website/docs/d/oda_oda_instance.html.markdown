---
subcategory: "Digital Assistant"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oda_oda_instance"
sidebar_current: "docs-oci-datasource-oda-oda_instance"
description: |-
  Provides details about a specific Oda Instance in Oracle Cloud Infrastructure Digital Assistant service
---

# Data Source: oci_oda_oda_instance
This data source provides details about a specific Oda Instance resource in Oracle Cloud Infrastructure Digital Assistant service.

Gets the specified Digital Assistant instance.

## Example Usage

```hcl
data "oci_oda_oda_instance" "test_oda_instance" {
	#Required
	oda_instance_id = oci_oda_oda_instance.test_oda_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `oda_instance_id` - (Required) Unique Digital Assistant instance identifier.


## Attributes Reference

The following attributes are exported:

* `attachment_ids` - A list of attachment identifiers for this instance (if any). Use GetOdaInstanceAttachment to get the details of the attachments.
* `attachment_types` - A list of attachment types for this instance (if any). Use attachmentIds to get the details of the attachments.
* `compartment_id` - Identifier of the compartment that the instance belongs to.
* `connector_url` - URL for the connector's endpoint.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the Digital Assistant instance.
* `display_name` - User-defined name for the Digital Assistant instance. Avoid entering confidential information. You can change this value. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Example: `{"bar-key": "value"}` 
* `id` - Unique immutable identifier that was assigned when the instance was created.
* `identity_app_console_url` - If isRoleBasedAccess is set to true, this property specifies the URL for the administration console used to manage the Identity Application instance Digital Assistant has created inside the user-specified identity domain.
* `identity_app_guid` - If isRoleBasedAccess is set to true, this property specifies the GUID of the Identity Application instance Digital Assistant has created inside the user-specified identity domain. This identity application instance may be used to host user roll mappings to grant access to this Digital Assistant instance for users within the identity domain.
* `identity_domain` - If isRoleBasedAccess is set to true, this property specifies the identity domain that is to be used to implement this type of authorzation. Digital Assistant will create an Identity Application instance and Application Roles within this identity domain. The caller may then perform and user roll mappings they like to grant access to users within the identity domain.
* `imported_package_ids` - A list of package ids imported into this instance (if any). Use GetImportedPackage to get the details of the imported packages.
* `imported_package_names` - A list of package names imported into this instance (if any). Use importedPackageIds field to get the details of the imported packages.
* `is_role_based_access` - Should this Digital Assistant instance use role-based authorization via an identity domain (true) or use the default policy-based authorization via IAM policies (false)
* `lifecycle_sub_state` - The current sub-state of the Digital Assistant instance.
* `restricted_operations` - A list of restricted operations (across all attachments) for this instance (if any). Use GetOdaInstanceAttachment to get the details of the attachments.
	* `operation_name` - Name of the restricted operation.
	* `restricting_service` - Name of the service restricting the operation.
* `shape_name` - Shape or size of the instance.
* `state` - The current state of the Digital Assistant instance.
* `state_message` - A message that describes the current state in more detail. For example, actionable information about an instance that's in the `FAILED` state. 
* `time_created` - When the Digital Assistant instance was created. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `time_updated` - When the Digital Assistance instance was last updated. A date-time string as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339), section 14.29.
* `web_app_url` - URL for the Digital Assistant web application that's associated with the instance.

