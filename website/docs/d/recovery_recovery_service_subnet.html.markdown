---
subcategory: "Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_recovery_recovery_service_subnet"
sidebar_current: "docs-oci-datasource-recovery-recovery_service_subnet"
description: |-
  Provides details about a specific Recovery Service Subnet in Oracle Cloud Infrastructure Recovery service
---

# Data Source: oci_recovery_recovery_service_subnet
This data source provides details about a specific Recovery Service Subnet resource in Oracle Cloud Infrastructure Recovery service.

Gets information about a specified recovery service subnet.

## Example Usage

```hcl
data "oci_recovery_recovery_service_subnet" "test_recovery_service_subnet" {
	#Required
	recovery_service_subnet_id = oci_recovery_recovery_service_subnet.test_recovery_service_subnet.id
}
```

## Argument Reference

The following arguments are supported:

* `recovery_service_subnet_id` - (Required) The recovery service subnet OCID.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment OCID.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `display_name` - A user-provided name for the recovery service subnet.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The recovery service subnet OCID.
* `lifecycle_details` - Detailed description about the current lifecycle state of the recovery service subnet. For example, it can be used to provide actionable information for a resource in a Failed state
* `nsg_ids` - A list of network security group (NSG) OCIDs that are associated with the Recovery Service subnet. You can specify a maximum of 5 unique OCIDs, which implies that you can associate a maximum of 5 NSGs to each Recovery Service subnet. Specify an empty array if you want to remove all the associated NSGs from a Recovery Service subnet. See [Network Security Groups](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/) for more information. 
* `state` - The current state of the recovery service subnet. 
* `subnet_id` - Deprecated. One of the subnets associated with the Recovery Service subnet. 
* `subnets` - A list of OCIDs of all the subnets associated with the Recovery Service subnet.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `time_created` - An RFC3339 formatted datetime string that indicates the last created time for a recovery service subnet. For example: '2020-05-22T21:10:29.600Z'. 
* `time_updated` - An RFC3339 formatted datetime string that indicates the last updated time for a recovery service subnet. For example: '2020-05-22T21:10:29.600Z'. 
* `vcn_id` - VCN Identifier.

