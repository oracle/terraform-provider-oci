---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_application_vip"
sidebar_current: "docs-oci-datasource-database-application_vip"
description: |-
  Provides details about a specific Application Vip in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_application_vip
This data source provides details about a specific Application Vip resource in Oracle Cloud Infrastructure Database service.

Gets information about a specified application virtual IP (VIP) address.

## Example Usage

```hcl
data "oci_database_application_vip" "test_application_vip" {
	#Required
	application_vip_id = oci_database_application_vip.test_application_vip.id
}
```

## Argument Reference

The following arguments are supported:

* `application_vip_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application virtual IP (VIP) address.


## Attributes Reference

The following attributes are exported:

* `cloud_vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud VM cluster associated with the application virtual IP (VIP) address.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - The hostname of the application virtual IP (VIP) address.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application virtual IP (VIP) address.
* `ip_address` - The application virtual IP (VIP) address.
* `lifecycle_details` - Additional information about the current lifecycle state of the application virtual IP (VIP) address.
* `state` - The current lifecycle state of the application virtual IP (VIP) address.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the application virtual IP (VIP) address.
* `time_assigned` - The date and time when the create operation for the application virtual IP (VIP) address completed.

