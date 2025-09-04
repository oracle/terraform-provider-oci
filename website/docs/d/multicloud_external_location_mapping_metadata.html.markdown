---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_external_location_mapping_metadata"
sidebar_current: "docs-oci-datasource-multicloud-external_location_mapping_metadata"
description: |-
  Provides the list of External Location Mapping Metadata in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_external_location_mapping_metadata
This data source provides the list of External Location Mapping Metadata in Oracle Cloud Infrastructure Multicloud service.

List externalLocation metadata from Oracle Cloud Infrastructure to the Cloud Service Provider for regions, Physical Availability Zones.

## Example Usage

```hcl
data "oci_multicloud_external_location_mapping_metadata" "test_external_location_mapping_metadata" {
  #Required
  compartment_id            = var.compartment_id
  subscription_service_name = var.subscription_service_name_list
  
  #Optional
  subscription_id			= var.subscription_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `subscription_service_name` - (Required) The subscription type values from [ORACLEDBATAZURE, ORACLEDBATGOOGLE, ORACLEDBATAWS]
* `subscription_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription in which to list resources.


## Attributes Reference

The following attributes are exported:

* `external_location_mapping_metadatum_summary_collection` - The list of ExternalLocationMappingMetadata.

### ExternalLocationMappingMetadata Reference

The following attributes are exported:

* `items` - List of ExternalLocationMappingMetadatumSummary
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `external_location` - External location for CSP Region, CSP-Physical-AZ
		* `csp_physical_az` - A mapping of Oracle Cloud Infrastructure site group name to CSP physical availability zone name
		* `csp_physical_az_display_name` - User friendly display name for cspPhysicalAZ
		* `csp_region` - CSP region corresponding to the given Oracle Cloud Infrastructure region
		* `csp_region_display_name` - CSP region display Name corresponding to the given Oracle Cloud Infrastructure region
		* `service_name` - The serviceName that externalLocation map object belongs to
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `oci_logical_ad` - Oracle Cloud Infrastructure logical ad name
	* `oci_physical_ad` - Oracle Cloud Infrastructure physical ad name
	* `oci_region` - Oracle Cloud Infrastructure region identifier https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}