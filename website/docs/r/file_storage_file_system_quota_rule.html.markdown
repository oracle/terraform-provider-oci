---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_file_system_quota_rule"
sidebar_current: "docs-oci-resource-file_storage-file_system_quota_rule"
description: |-
  Provides the File System Quota Rule resource in Oracle Cloud Infrastructure File Storage service
---

# oci_file_storage_file_system_quota_rule
This resource provides the File System Quota Rule resource in Oracle Cloud Infrastructure File Storage service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/filestorage/latest/FileSystemQuotaRule

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Create an FS level, user or group quota rule given the `fileSystemId`, `principalId`, `principalType` and
`isHardQuota` parameters.


## Example Usage

```hcl
resource "oci_file_storage_file_system_quota_rule" "test_file_system_quota_rule" {
	#Required
	file_system_id = oci_file_storage_file_system.test_file_system.id
	is_hard_quota = var.file_system_quota_rule_is_hard_quota
	principal_type = var.file_system_quota_rule_principal_type
	quota_limit_in_gigabytes = var.file_system_quota_rule_quota_limit_in_gigabytes

	#Optional
	display_name = var.file_system_quota_rule_display_name
	principal_id = oci_file_storage_principal.test_principal.id
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information. Example: `UserXYZ's quota` 
* `file_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
* `is_hard_quota` - (Required) The flag is an identifier to tell whether the quota rule will be enforced. If `isHardQuota` is true, the quota rule will be enforced so the write will be blocked if usage exceeds the hard quota limit. If `isHardQuota` is false, usage can exceed the soft quota limit. An alarm or notification will be sent to the customer, if the specific usage exceeds. 
* `principal_id` - (Optional) An identifier for the owner of this usage and quota rule. Unix-like operating systems use this integer value to identify a user or group to manage access control. 
* `principal_type` - (Required) The type of the owner of this quota rule and usage. 
* `quota_limit_in_gigabytes` - (Required) (Updatable) The value of the quota rule. The unit is Gigabyte. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information. Example: `UserXYZ's quota` 
* `file_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file System.
* `id` - The identifier of the quota rule. It is the base64 encoded string of the tuple <principalId, principalType, isHardQuota>.
* `is_hard_quota` - The flag is an identifier to tell whether the quota rule will be enforced. If `isHardQuota` is false, the quota rule will be enforced so the usage cannot exceed the hard quota limit. If `isHardQuota` is true, usage can exceed the soft quota limit. An alarm or notification will be sent to the customer, if the specific usage exceeds. 
* `principal_id` - An identifier for the owner of this usage and quota rule. Unix-like operating systems use this integer value to identify a user or group to manage access control. 
* `principal_type` - The type of the owner of this quota rule and usage. 
* `quota_limit_in_gigabytes` - The value of the quota rule. The unit is Gigabyte. 
* `time_created` - The date and time the quota rule was started, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the quota rule was last updated, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the File System Quota Rule
	* `update` - (Defaults to 20 minutes), when updating the File System Quota Rule
	* `delete` - (Defaults to 20 minutes), when destroying the File System Quota Rule


## Import

FileSystemQuotaRules can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_file_system_quota_rule.test_file_system_quota_rule "fileSystems/{fileSystemId}/quotaRules/{quotaRuleId}" 
```

