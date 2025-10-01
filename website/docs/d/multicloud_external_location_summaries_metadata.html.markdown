---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_external_location_summaries_metadata"
sidebar_current: "docs-oci-datasource-multicloud-external_location_summaries_metadata"
description: |-
  Provides the list of External Location Summaries Metadata in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_external_location_summaries_metadata
This data source provides the list of External Location Summaries Metadata in Oracle Cloud Infrastructure Multicloud service.

List externalLocationSummary metadata from Oracle Cloud Infrastructure Region to the Cloud Service Provider region across all regions.

## Example Usage

```hcl
data "oci_multicloud_external_location_summaries_metadata" "test_external_location_summaries_metadata" {
	#Required
	compartment_id              = var.compartment_id
	subscription_service_name   = var.subscription_service_name

	#Optional
	entity_type                 = var.entity_type
	subscription_id             = var.subscription_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `subscription_service_name` - (Required) The subscription service name values from [ORACLEDBATAZURE, ORACLEDBATGOOGLE, ORACLEDBATAWS]
* `entity_type` - (Optional) The resource type query (i.e. dbsystem, instance etc.)
* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription in which to list resources.


## Attributes Reference

The following attributes are exported:

* `external_location_summaries_metadatum_summary_collection` - The list of ExternalLocationSummariesMetadata.

### ExternalLocationSummariesMetadata Reference

The following attributes are exported:

* `items` - List of ExternalLocationSummariesMetadatumSummary
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `external_location` - External location for CSP Region
		* `csp_region` - CSP region corresponding to the given Oracle Cloud Infrastructure region
		* `csp_region_display_name` - CSP region display Name corresponding to the given Oracle Cloud Infrastructure region
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `oci_region` - Oracle Cloud Infrastructure region identifier https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`