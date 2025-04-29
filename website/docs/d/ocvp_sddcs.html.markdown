---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_sddcs"
sidebar_current: "docs-oci-datasource-ocvp-sddcs"
description: |-
  Provides the list of Sddcs in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# Data Source: oci_ocvp_sddcs
This data source provides the list of Sddcs in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Lists the SDDCs in the specified compartment. The list can be
filtered by display name or availability domain.


## Example Usage

```hcl
data "oci_ocvp_sddcs" "test_sddcs" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compute_availability_domain = var.sddc_compute_availability_domain
	display_name = var.sddc_display_name
	state = var.sddc_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_availability_domain` - (Optional) The name of the availability domain that the Compute instances are running in.  Example: `Uocm:PHX-AD-1` 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `state` - (Optional) The lifecycle state of the resource.


## Attributes Reference

The following attributes are exported:

* `sddc_collection` - The list of sddc_collection.

### Sddc Reference

The following attributes are exported:

* `clusters_count` - The number of Clusters in the SDDC.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the SDDC. 
* `compute_availability_domain` - (**Deprecated**) The availability domain the ESXi hosts are running in. For Multi-AD SDDC, it is `multi-AD`.  Example: `Uocm:PHX-AD-1`, `multi-AD`.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A descriptive name for the SDDC. It must be unique, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information. 
* `esxi_hosts_count` - (**Deprecated**) The number of ESXi hosts in the SDDC.
* `actual_esxi_hosts_count` - (**Deprecated**) The number of actual ESXi hosts in the SDDC on the cloud. This attribute will be different when esxi Host is added to an existing SDDC.
* `esxi_software_version` - In general, this is a specific version of bundled ESXi software supported by Oracle Cloud VMware Solution (see [ListSupportedVmwareSoftwareVersions](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedVmwareSoftwareVersionSummary/ListSupportedVmwareSoftwareVersions)).

	This attribute is not guaranteed to reflect the version of software currently installed on the ESXi hosts in the SDDC. The purpose of this attribute is to show the version of software that the Oracle Cloud VMware Solution will install on any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/EsxiHost/CreateEsxiHost)  unless a different version is configured on the Cluster or ESXi host level.

	Therefore, if you upgrade the existing ESXi hosts in the SDDC to use a newer version of bundled ESXi software supported by the Oracle Cloud VMware Solution, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/UpdateSddc) to update the SDDC's `vmwareSoftwareVersion` with that new version. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hcx_fqdn` - The FQDN for HCX Manager.  Example: `hcx-my-sddc.sddc.us-phoenix-1.oraclecloud.com` 
* `hcx_mode` - HCX configuration of the SDDC.
* `hcx_on_prem_licenses` - The activation licenses to use on the on-premises HCX Enterprise appliance you site pair with HCX Manager in your VMware Solution. 
	* `activation_key` - HCX on-premise license key value.
	* `status` - status of HCX on-premise license.
	* `system_name` - Name of the system that consumed the HCX on-premise license
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC.
* `initial_configuration` - Details of SDDC initial configuration
	* `initial_cluster_configurations` - The configurations for Clusters initially created in the SDDC. 
		* `capacity_reservation_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Capacity Reservation. 
		* `compute_availability_domain` - The availability domain to create the Cluster's ESXi hosts in. For multi-AD Cluster deployment, set to `multi-AD`. 
		* `datastore_cluster_ids` - A list of datastore clusters. 
		* `datastores` - A list of datastore info for the Cluster. This value is required only when `initialHostShapeName` is a standard shape. 
			* `block_volume_ids` - A list of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of Block Storage Volumes.
			* `datastore_type` - Type of the datastore.
		* `display_name` - A descriptive name for the Cluster. Cluster name requirements are 1-22 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the region. Avoid entering confidential information. 
		* `esxi_hosts_count` - The number of ESXi hosts to create in the Cluster. You can add more hosts later (see [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost)). Creating a Cluster with a ESXi host count of 1 will be considered a single ESXi host Cluster.

			**Note:** If you later delete EXSi hosts from a production Cluster to total less than 3, you are still billed for the 3 minimum recommended ESXi hosts. Also, you cannot add more VMware workloads to the Cluster until it again has at least 3 ESXi hosts. 
		* `initial_commitment` - The billing option selected during Cluster creation. [ListSupportedCommitments](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedCommitmentSummary/ListSupportedCommitments). 
		* `initial_host_ocpu_count` - The initial OCPU count of the Cluster's ESXi hosts. 
		* `initial_host_shape_name` - The initial compute shape of the Cluster's ESXi hosts. [ListSupportedHostShapes](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedHostShapes/ListSupportedHostShapes). 
		* `instance_display_name_prefix` - A prefix used in the name of each ESXi host and Compute instance in the Cluster. If this isn't set, the Cluster's `displayName` is used as the prefix.

			For example, if the value is `myCluster`, the ESXi hosts are named `myCluster-1`, `myCluster-2`, and so on. 
		* `is_shielded_instance_enabled` - Indicates whether shielded instance is enabled for this Cluster. 
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
		* `vsphere_type` - vSphere Cluster types. 
		* `workload_network_cidr` - The CIDR block for the IP addresses that VMware VMs in the Cluster use to run application workloads.
* `initial_host_ocpu_count` - (**Deprecated**) The initial OCPU count of the SDDC's ESXi hosts.
* `initial_host_shape_name` - (**Deprecated**) The initial compute shape of the SDDC's ESXi hosts. [ListSupportedHostShapes](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedHostShapes/ListSupportedHostShapes).
* `is_hcx_enabled` - (**Deprecated**) Indicates whether HCX is enabled for this SDDC. **Deprecated**. Please use `hcx_mode` instead.
* `is_hcx_enterprise_enabled` - (**Deprecated**) Indicates whether HCX Enterprise is enabled for this SDDC.  **Deprecated**. Please use `hcx_mode` instead.
* `is_hcx_pending_downgrade` - Indicates whether SDDC is pending downgrade from HCX Enterprise to HCX Advanced.
* `is_shielded_instance_enabled` - (**Deprecated**) Indicates whether shielded instance is enabled at the SDDC level.
* `is_single_host_sddc` - Indicates whether this SDDC is designated for only single ESXi host.
* `nsx_manager_fqdn` - The FQDN for NSX Manager.  Example: `nsx-my-sddc.sddc.us-phoenix-1.oraclecloud.com`
* `nsx_manager_username` - The SDDC includes an administrator username and initial password for NSX Manager. You can change this initial username to a different value in NSX Manager.
* `state` - The current state of the SDDC.
* `time_created` - The date and time the SDDC was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the SDDC was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcenter_fqdn` - The FQDN for vCenter.  Example: `vcenter-my-sddc.sddc.us-phoenix-1.oraclecloud.com` 
* `vmware_software_version` - In general, this is a specific version of bundled VMware software supported by Oracle Cloud VMware Solution (see [ListSupportedVmwareSoftwareVersions](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedVmwareSoftwareVersionSummary/ListSupportedVmwareSoftwareVersions)).

	This attribute is not guaranteed to reflect the version of software currently installed on the ESXi hosts in the SDDC. The purpose of this attribute is to show the version of software that the Oracle Cloud VMware Solution will install on any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/EsxiHost/CreateEsxiHost).

	Therefore, if you upgrade the existing ESXi hosts in the SDDC to use a newer version of bundled VMware software supported by the Oracle Cloud VMware Solution, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/UpdateSddc) to update the SDDC's `vmwareSoftwareVersion` with that new version.

