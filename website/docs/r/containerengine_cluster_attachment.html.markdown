---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_attachment"
sidebar_current: "docs-oci-resource-containerengine-cluster_attachment"
description: |-
  Provides the Cluster Attachment resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_cluster_attachment
This resource provides the Cluster Attachment resource in Oracle Cloud Infrastructure Container Engine service.

Creates a new ClusterAttachment.


## Example Usage

```hcl
resource "oci_containerengine_cluster_attachment" "test_cluster_attachment" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
	cluster_namespace_profile_id = oci_containerengine_cluster_namespace_profile.test_cluster_namespace_profile.id
	compartment_id = var.compartment_id
	display_name = var.cluster_attachment_display_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.cluster_attachment_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) OCID of the Cluster
* `cluster_namespace_profile_id` - (Required) OCID of the Cluster Namespace Profile
* `compartment_id` - (Required) (Updatable) OCID of compartment owning the Cluster Namespace.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) Description of the resource. It can be changed after creation.
* `display_name` - (Required) (Updatable) Name of the Cluster Namespace.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cluster_id` - OCID of the Cluster
* `cluster_namespace_profile_id` - OCID of the Cluster Namespace Profile
* `compartment_id` - OCID of compartment owning the Cluster Namespace.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of the resource. It can be changed after creation.
* `display_name` - Name of the Cluster Namespace.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `state` - The current state of the ClusterAttachment.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when this resource was created in an RFC3339 formatted datetime string.
* `time_updated` - The time when this resource was updated in an RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster Attachment
	* `update` - (Defaults to 20 minutes), when updating the Cluster Attachment
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster Attachment


## Import

ClusterAttachments can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_cluster_attachment.test_cluster_attachment "id"
```

