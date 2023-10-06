---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_cluster_attachment"
sidebar_current: "docs-oci-datasource-containerengine-cluster_attachment"
description: |-
  Provides details about a specific Cluster Attachment in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_cluster_attachment
This data source provides details about a specific Cluster Attachment resource in Oracle Cloud Infrastructure Container Engine service.

Gets a ClusterAttachment by identifier

## Example Usage

```hcl
data "oci_containerengine_cluster_attachment" "test_cluster_attachment" {
	#Required
	cluster_attachment_id = oci_containerengine_cluster_attachment.test_cluster_attachment.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_attachment_id` - (Required) unique ClusterAttachment identifier


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

