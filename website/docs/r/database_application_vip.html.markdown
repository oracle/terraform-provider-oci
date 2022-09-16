---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_application_vip"
sidebar_current: "docs-oci-resource-database-application_vip"
description: |-
  Provides the Application Vip resource in Oracle Cloud Infrastructure Database service
---

# oci_database_application_vip
This resource provides the Application Vip resource in Oracle Cloud Infrastructure Database service.

Creates a new application virtual IP (VIP) address in the specified cloud VM cluster based on the request parameters you provide.


## Example Usage

```hcl
resource "oci_database_application_vip" "test_application_vip" {
	#Required
	cloud_vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id
	hostname_label = var.application_vip_hostname_label
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	db_node_id = oci_database_db_node.test_db_node.id
	ip_address = var.application_vip_ip_address
}
```

## Argument Reference

The following arguments are supported:

* `cloud_vm_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud VM cluster associated with the application virtual IP (VIP) address.
* `db_node_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB node associated with the application virtual IP (VIP) address.
* `hostname_label` - (Required) The hostname of the application virtual IP (VIP) address.
* `ip_address` - (Optional) The application virtual IP (VIP) address.
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet associated with the application virtual IP (VIP) address.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Application Vip
	* `delete` - (Defaults to 20 minutes), when destroying the Application Vip


## Import

ApplicationVips can be imported using the `id`, e.g.

```
$ terraform import oci_database_application_vip.test_application_vip "id"
```

