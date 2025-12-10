---
subcategory: "Multicloud"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_multicloud_external_locations_metadata"
sidebar_current: "docs-oci-datasource-multicloud-external_locations_metadata"
description: |-
  Provides the list of External Locations Metadata in Oracle Cloud Infrastructure Multicloud service
---

# Data Source: oci_multicloud_external_locations_metadata
This data source provides the list of External Locations Metadata in Oracle Cloud Infrastructure Multicloud service.

List externalLocationDetail metadata from Oracle Cloud Infrastructure to Cloud  Service Provider for regions, Availability Zones, and Cluster Placement Group ID.

## Example Usage

```hcl
data "oci_multicloud_external_locations_metadata" "test_external_locations_metadata" {
	#Required
	subscription_id           = var.subscription_id
	subscription_service_name = var.subscription_service_name

	#Optional
	cluster_placement_group_id 	= var.cluster_placement_group_id
	compartment_id            	= var.compartment_id
	entity_type               	= var.external_locations_metadata_entity_type
	external_location			= var.external_location
	linked_compartment_id     	= var.linked_compartment_id
	logical_zone 				= var.logical_zone
}
```

## Argument Reference

The following arguments are supported:

* `cluster_placement_group_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cluster Placement Group.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud base compartment in which to list resources.  A Multicloud base compartment is an Oracle Cloud Infrastructure compartment that maps to a subscription in a Cloud Service Provider (such as Azure, AWS, or Google Cloud).  
* `entity_type` - (Optional) The resource type query (i.e. dbsystem, instance etc.)
* `external_location` - (Optional) The Cloud Service Provider region.
* `linked_compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment linked to the resource.
* `logical_zone` - (Optional) Oracle Cloud Infrastructure Logical AD to filter the response.
* `subscription_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Multicloud subscription in which to list resources.
* `subscription_service_name` - (Required) The subscription service name of the Cloud Service Provider.


## Attributes Reference

The following attributes are exported:

* `external_locations_metadatum_collection` - The list of ExternalLocationsMetadata.

### ExternalLocationsMetadata Reference

The following attributes are exported:

* `items` - List of ExternalLocationsMetadatumSummary
	* `cluster_placement_group_id` - Cluster Placement Group OCID
	* `cpg_id` - Cluster Placement Group OCID (deprecated representation)
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `external_location` - External location for CSP Region, CSP-Physical-AZ, CSP-Logical-AZ
		* `csp_logical_az` - A mapping of CSP physical availability zone to CSP logical availability zone.
		* `csp_logical_az_display_name` - User friendly display name for cspLogicalAZ
		* `csp_physical_az` - A mapping of Oracle Cloud Infrastructure site group name to CSP physical availability zone name
		* `csp_physical_az_display_name` - User friendly display name for cspPhysicalAZ
		* `csp_region` - CSP region corresponding to the given Oracle Cloud Infrastructure region
		* `csp_region_display_name` - CSP region display Name corresponding to the given Oracle Cloud Infrastructure region
		* `csp_zone_key_reference_id` - This is CSP zone key reference
			* `key_name` - KeyName for Azure=AzureSubscriptionId Aws=AwsAccountId GCP=GcpProjectName
			* `key_value` - Value of keyName GcpProjectName: A human-readable name for your project. The project name isn't used by any Google APIs. You can edit the project name at any time during or after project creation. Project names do not need to be unique. AzureSubscriptionId: A unique alphanumeric string that identifies your Azure subscription. AwsAccountId: a unique 12-digit number that identifies an Amazon Web Services (AWS) account 
		* `service_name` - The serviceName that externalLocation map object belongs to.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `oci_logical_ad` - Oracle Cloud Infrastructure logical ad name
	* `oci_physical_ad` - Oracle Cloud Infrastructure physical ad name
	* `oci_region` - Oracle Cloud Infrastructure region identifier https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
	* `partner_cloud_account_name` - User friendly name of account name for customer's subscription
	* `partner_cloud_account_url` - Direct URL to partner cloud for customer's account
	* `partner_cloud_name` - Partner Cloud Name based on service name
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
