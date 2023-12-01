---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_cluster"
sidebar_current: "docs-oci-datasource-ocvp-cluster"
description: |-
  Provides details about a specific Cluster in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_cluster
This data source provides details about a specific Cluster resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Gets the specified Cluster's information.

## Example Usage

```hcl
data "oci_ocvp_cluster" "test_cluster" {
	#Required
	cluster_id = oci_ocvp_cluster.test_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cluster. 


## Attributes Reference

The following attributes are exported:

* `capacity_reservation_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Capacity Reservation. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Cluster. 
* `compute_availability_domain` - The availability domain the ESXi hosts are running in. For Multi-AD Cluster, it is `multi-AD`.  Example: `Uocm:PHX-AD-1`, `multi-AD` 
* `datastores` - Datastores used for the Cluster. 
	* `block_volume_ids` - A list of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of Block Storage Volumes.
	* `capacity` - Size of the Block Storage Volume in GB.
	* `datastore_type` - Type of the datastore.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A descriptive name for the Cluster. It must be unique, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information. 
* `esxi_hosts_count` - The number of ESXi hosts in the Cluster.
* `esxi_software_version` - In general, this is a specific version of bundled ESXi software supported by Oracle Cloud VMware Solution (see [ListSupportedVmwareSoftwareVersions](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedVmwareSoftwareVersionSummary/ListSupportedVmwareSoftwareVersions)).

	This attribute is not guaranteed to reflect the version of software currently installed on the ESXi hosts in the SDDC. The purpose of this attribute is to show the version of software that the Oracle Cloud VMware Solution will install on any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/EsxiHost/CreateEsxiHost)  unless a different version is configured on the ESXi host level.

	Therefore, if you upgrade the existing ESXi hosts in the Cluster to use a newer version of bundled ESXi software supported by the Oracle Cloud VMware Solution, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `esxiSoftwareVersion` with that new version. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cluster. 
* `initial_commitment` - The billing option selected during Cluster creation. [ListSupportedCommitments](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedCommitmentSummary/ListSupportedCommitments). 
* `initial_host_ocpu_count` - The initial OCPU count of the Cluster's ESXi hosts. 
* `initial_host_shape_name` - The initial compute shape of the Cluster's ESXi hosts. [ListSupportedHostShapes](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedHostShapes/ListSupportedHostShapes). 
* `instance_display_name_prefix` - A prefix used in the name of each ESXi host and Compute instance in the Cluster. If this isn't set, the Cluster's `displayName` is used as the prefix.

	For example, if the value is `MyCluster`, the ESXi hosts are named `MyCluster-1`, `MyCluster-2`, and so on. 
* `is_shielded_instance_enabled` - Indicates whether shielded instance is enabled at the Cluster level. 
* `network_configuration` - The network configurations used by Cluster, including [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management subnet and VLANs. 
	* `hcx_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the HCX component of the VMware environment. This VLAN is a mandatory attribute  for Management Cluster when HCX is enabled.

		This attribute is not guaranteed to reflect the HCX VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the HCX VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the HCX component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/UpdateSddc) to update the SDDC's `hcxVlanId` with that new VLAN's OCID. 
	* `nsx_edge_uplink1vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the NSX Edge Uplink 1 component of the VMware environment. This VLAN is a mandatory attribute for Management Cluster.

		This attribute is not guaranteed to reflect the NSX Edge Uplink 1 VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the NSX Edge Uplink 1 VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the NSX Edge Uplink 1 component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `nsxEdgeUplink1VlanId` with that new VLAN's OCID. 
	* `nsx_edge_uplink2vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC  for the NSX Edge Uplink 2 component of the VMware environment. This VLAN is a mandatory attribute for Management Cluster.

		This attribute is not guaranteed to reflect the NSX Edge Uplink 2 VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the NSX Edge Uplink 2 VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the NSX Edge Uplink 2 component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `nsxEdgeUplink2VlanId` with that new VLAN's OCID. 
	* `nsx_edge_vtep_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the NSX Edge VTEP component of the VMware environment.

		This attribute is not guaranteed to reflect the NSX Edge VTEP VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the NSX Edge VTEP VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the NSX Edge VTEP component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `nsxEdgeVTepVlanId` with that new VLAN's OCID. 
	* `nsx_vtep_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the NSX VTEP component of the VMware environment.

		This attribute is not guaranteed to reflect the NSX VTEP VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the NSX VTEP VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the NSX VTEP component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `nsxVTepVlanId` with that new VLAN's OCID. 
	* `provisioning_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management subnet used to provision the Cluster. 
	* `provisioning_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the Provisioning component of the VMware environment. 
	* `replication_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the vSphere Replication component of the VMware environment. 
	* `vmotion_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the vMotion component of the VMware environment.

		This attribute is not guaranteed to reflect the vMotion VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the vMotion VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the vMotion component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateCluster) to update the Cluster's `vmotionVlanId` with that new VLAN's OCID. 
	* `vsan_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the vSAN component of the VMware environment.

		This attribute is not guaranteed to reflect the vSAN VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the vSAN VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the vSAN component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `vsanVlanId` with that new VLAN's OCID. 
	* `vsphere_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the vSphere component of the VMware environment. This VLAN is a mandatory attribute for Management Cluster.

		This attribute is not guaranteed to reflect the vSphere VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the vSphere VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the vSphere component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the Cluster's `vsphereVlanId` with that new VLAN's OCID. 
* `sddc_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC that the Cluster belongs to. 
* `state` - The current state of the Cluster.
* `time_created` - The date and time the Cluster was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the Cluster was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `upgrade_licenses` - The vSphere licenses to use when upgrading the Cluster. 
	* `license_key` - vSphere license key value.
	* `license_type` - vSphere license type.
* `vmware_software_version` - In general, this is a specific version of bundled VMware software supported by Oracle Cloud VMware Solution (see [ListSupportedVmwareSoftwareVersions](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedVmwareSoftwareVersionSummary/ListSupportedVmwareSoftwareVersions)).

	This attribute is not guaranteed to reflect the version of software currently installed on the ESXi hosts in the Cluster. The purpose of this attribute is to show the version of software that the Oracle Cloud VMware Solution will install on any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you upgrade the existing ESXi hosts in the Cluster to use a newer version of bundled VMware software supported by the Oracle Cloud VMware Solution, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `vmwareSoftwareVersion` with that new version. 
* `vsphere_type` - vSphere Cluster types. 
* `vsphere_upgrade_objects` - The links to binary objects needed to upgrade vSphere. 
	* `download_link` - Binary object download link.
	* `link_description` - Binary object description.
* `workload_network_cidr` - The CIDR block for the IP addresses that VMware VMs in the SDDC use to run application workloads. 

