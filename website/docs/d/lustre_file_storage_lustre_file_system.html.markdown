---
subcategory: "Lustre File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_lustre_file_storage_lustre_file_system"
sidebar_current: "docs-oci-datasource-lustre_file_storage-lustre_file_system"
description: |-
  Provides details about a specific Lustre File System in Oracle Cloud Infrastructure Lustre File Storage service
---

# Data Source: oci_lustre_file_storage_lustre_file_system
This data source provides details about a specific Lustre File System resource in Oracle Cloud Infrastructure Lustre File Storage service.

Gets information about a Lustre file system.

## Example Usage

```hcl
data "oci_lustre_file_storage_lustre_file_system" "test_lustre_file_system" {
	#Required
	lustre_file_system_id = oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id
}
```

## Argument Reference

The following arguments are supported:

* `lustre_file_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Lustre file system.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the file system is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `capacity_in_gbs` - Capacity of the Lustre file system in GB.
* `cluster_placement_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group in which the Lustre file system exists.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Lustre file system.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My Lustre file system` 
* `file_system_description` - Short description of the Lustre file system. Avoid entering confidential information. 
* `file_system_name` - The Lustre file system name. This is used in mount commands and other aspects of the client command line interface. The default file system name is 'lustre'. The file system name is limited to 8 characters. Allowed characters are lower and upper case English letters, numbers, and '_'. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Lustre file system.
* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key used to encrypt the encryption keys associated with this file system. 
* `lifecycle_details` - A message that describes the current state of the Lustre file system in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `lnet` - Type of network used by clients to mount the file system.   Example: `tcp` 
* `maintenance_window` - The preferred day and time to perform maintenance.
	* `day_of_week` - Day of the week when the maintainence window starts. 
	* `time_start` - The time to start the maintenance window. The format is 'HH:MM', 'HH:MM' represents the time in UTC.   Example: `22:00` 
* `major_version` - Major version of Lustre running in the Lustre file system.  Example: `2.15` 
* `management_service_address` - The IPv4 address of MGS (Lustre Management Service) used by clients to mount the file system. For example '10.0.0.4'.
* `nsg_ids` - A list of Network Security Group [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this lustre file system. A maximum of 5 is allowed. Setting this to an empty array after the list is created removes the lustre file system from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). 
* `performance_tier` - The Lustre file system performance tier. A value of `MBPS_PER_TB_125` represents 125 megabytes per second per terabyte.
* `root_squash_configuration` - An administrative feature that allows you to restrict root level access from clients that try to access your Lustre file system as root.
	* `client_exceptions` - A list of NIDs allowed with this lustre file system not to be squashed. A maximum of 10 is allowed. 
	* `identity_squash` - Used when clients accessing the Lustre file system have their UID and GID remapped to `squashUid` and `squashGid`. If `ROOT`, only the root user and group (UID/GID 0) are remapped; if `NONE`, no remapping is done. If unspecified, defaults to `NONE`. 
	* `squash_gid` - The GID value to remap to when squashing a client GID. See `identitySquash` for more details. If unspecified, defaults to `65534`. 
	* `squash_uid` - The UID value to remap to when squashing a client UID. See `identitySquash` for more details. If unspecified, defaults to `65534`. 
* `state` - The current state of the Lustre file system.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the Lustre file system is in.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_billing_cycle_end` - The date and time that the current billing cycle for the file system will end, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format. After the current cycle ends, this date is updated automatically to the next timestamp, which is 30 days later. File systems deleted earlier than this time will still incur charges until the billing cycle ends.  Example: `2016-08-25T21:10:29.600Z` 
* `time_created` - The date and time the Lustre file system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2024-04-25T21:10:29.600Z` 
* `time_updated` - The date and time the Lustre file system was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2024-04-25T21:10:29.600Z` 

