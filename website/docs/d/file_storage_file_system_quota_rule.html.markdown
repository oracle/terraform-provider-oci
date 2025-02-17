---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_file_system_quota_rule"
sidebar_current: "docs-oci-datasource-file_storage-file_system_quota_rule"
description: |-
  Provides details about a specific File System Quota Rule in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_file_system_quota_rule
This data source provides details about a specific File System Quota Rule resource in Oracle Cloud Infrastructure File Storage service.

Get an FS level, user or group quota rule given the `fileSystemId` and `quotaRuleId` parameters.


## Example Usage

```hcl
data "oci_file_storage_file_system_quota_rule" "test_file_system_quota_rule" {
	#Required
	file_system_id = oci_file_storage_file_system.test_file_system.id
	quota_rule_id = oci_events_rule.test_rule.id
}
```

## Argument Reference

The following arguments are supported:

* `file_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
* `quota_rule_id` - (Required) The identifier of the quota rule. It is the base64 encoded string of the tuple <principalId, principalType, isHardQuota>.


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

