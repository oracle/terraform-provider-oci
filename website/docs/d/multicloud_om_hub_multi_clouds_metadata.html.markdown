---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_om_hub_multi_clouds_metadata"
sidebar_current: "docs-oci-datasource-multicloud-om_hub_multi_clouds_metadata"
description: |-
  Provides the list of compartments under a root compartment in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_om_hub_multi_clouds_metadata
This data source provides information about the Multicloud base compartment for a given tenancy Id.
A Multicloud base compartment is an Oracle Cloud Infrastructure compartment that maps to a subscription in a Cloud Service Provider (such as Azure, AWS, or Google Cloud).

Gets a list of multicloud metadata with multicloud base compartment and subscription across Cloud Service Providers.

## Example Usage

```hcl
data "oci_multicloud_om_hub_multi_clouds_metadata" "test_om_hub_multi_clouds_metadata" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the root compartment in which to list resources.

## Attributes Reference

The following attributes are exported:

* `multi_cloud_metadata_collection` - The list of MultiCloudMetadataCollection.

### MultiCloudMetadataCollection Reference

The following attributes are exported:

* `items` - List of MultiCloudMetadataSummary
    * `compartment_id` - MultiCloud base compartment OCID associated with subscriptionId.
    * `time_created` - The date and time the multicloud compartment was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
    * `subscription_id` - Oracle Cloud Infrastructure subscriptionId.
    * `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
    * `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
    * `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`