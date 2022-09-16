---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_application_vips"
sidebar_current: "docs-oci-datasource-database-application_vips"
description: |-
  Provides the list of Application Vips in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_application_vips
This data source provides the list of Application Vips in Oracle Cloud Infrastructure Database service.

Gets a list of application virtual IP (VIP) addresses on a cloud VM cluster.


## Example Usage

```hcl
data "oci_database_application_vips" "test_application_vips" {
	#Required
	cloud_vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id
	compartment_id = var.compartment_id

	#Optional
	state = var.application_vip_state
}
```

## Argument Reference

The following arguments are supported:

* `cloud_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud VM cluster associated with the application virtual IP (VIP) address.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `application_vips` - The list of application_vips.

### ApplicationVip Reference

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

