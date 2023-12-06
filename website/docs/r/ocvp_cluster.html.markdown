---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_cluster"
sidebar_current: "docs-oci-resource-ocvp-cluster"
description: |-
  Provides the Cluster resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# oci_ocvp_cluster
This resource provides the Cluster resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Create a vSphere Cluster in software-defined data center (SDDC).

Use the [WorkRequest](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/WorkRequest/) operations to track the
creation of the Cluster.

**Important:** You must configure the Cluster's networking resources with the security rules detailed in [Security Rules for Oracle Cloud VMware Solution SDDCs](https://docs.cloud.oracle.com/iaas/Content/VMware/Reference/ocvssecurityrules.htm). Otherwise, provisioning the SDDC will fail. The rules are based on the requirements set by VMware.


## Example Usage

```hcl
resource "oci_ocvp_cluster" "test_cluster" {
	#Required
	compute_availability_domain = var.cluster_compute_availability_domain
	esxi_hosts_count = var.cluster_esxi_hosts_count
	network_configuration {
		#Required
		nsx_edge_vtep_vlan_id = oci_core_vlan.test_vlan.id
		nsx_vtep_vlan_id = oci_core_vlan.test_vlan.id
		provisioning_subnet_id = oci_core_subnet.test_subnet.id
		vmotion_vlan_id = oci_core_vlan.test_vlan.id
		vsan_vlan_id = oci_core_vlan.test_vlan.id

		#Optional
		hcx_vlan_id = oci_core_vlan.test_vlan.id
		nsx_edge_uplink1vlan_id = oci_ocvp_nsx_edge_uplink1vlan.test_nsx_edge_uplink1vlan.id
		nsx_edge_uplink2vlan_id = oci_ocvp_nsx_edge_uplink2vlan.test_nsx_edge_uplink2vlan.id
		provisioning_vlan_id = oci_core_vlan.test_vlan.id
		replication_vlan_id = oci_core_vlan.test_vlan.id
		vsphere_vlan_id = oci_core_vlan.test_vlan.id
	}
	sddc_id = oci_ocvp_sddc.test_sddc.id

	#Optional
	capacity_reservation_id = oci_ocvp_capacity_reservation.test_capacity_reservation.id
	datastores {
		#Required
		block_volume_ids = var.cluster_datastores_block_volume_ids
		datastore_type = var.cluster_datastores_datastore_type
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.cluster_display_name
	esxi_software_version = var.cluster_esxi_software_version
	freeform_tags = {"Department"= "Finance"}
	initial_commitment = var.cluster_initial_commitment
	initial_host_ocpu_count = var.cluster_initial_host_ocpu_count
	initial_host_shape_name = oci_core_shape.test_shape.name
	instance_display_name_prefix = var.cluster_instance_display_name_prefix
	is_shielded_instance_enabled = var.cluster_is_shielded_instance_enabled
	vmware_software_version = var.cluster_vmware_software_version
	workload_network_cidr = var.cluster_workload_network_cidr
}
```

## Argument Reference

The following arguments are supported:

* `capacity_reservation_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Capacity Reservation. 
* `compute_availability_domain` - (Required) The availability domain to create the Cluster's ESXi hosts in. For multi-AD Cluster deployment, set to `multi-AD`. 
* `datastores` - (Optional) A list of datastore info for the Cluster. This value is required only when `initialHostShapeName` is a standard shape. 
	* `block_volume_ids` - (Required) A list of [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)s of Block Storage Volumes.
	* `datastore_type` - (Required) Type of the datastore.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A descriptive name for the Cluster. Cluster name requirements are 1-16 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the region. Avoid entering confidential information. 
* `esxi_hosts_count` - (Required) The number of ESXi hosts to create in the Cluster. You can add more hosts later (see [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost)). 

	**Note:** If you later delete EXSi hosts from a production Cluster to make SDDC total host count less than 3, you are still billed for the 3 minimum recommended  ESXi hosts. Also, you cannot add more VMware workloads to the Cluster until the  SDDC again has at least 3 ESXi hosts. 
* `esxi_software_version` - (Optional) (Updatable) The ESXi software bundle to install on the ESXi hosts in the Cluster.  Only versions under the same vmwareSoftwareVersion and have been validate by Oracle Cloud VMware Solution will be accepted. To get a list of the available versions, use [ListSupportedVmwareSoftwareVersions](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedVmwareSoftwareVersionSummary/ListSupportedVmwareSoftwareVersions). 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `initial_commitment` - (Optional) The billing option selected during Cluster creation. [ListSupportedCommitments](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedCommitmentSummary/ListSupportedCommitments). 
* `initial_host_ocpu_count` - (Optional) The initial OCPU count of the Cluster's ESXi hosts. 
* `initial_host_shape_name` - (Optional) The initial compute shape of the Cluster's ESXi hosts. [ListSupportedHostShapes](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedHostShapes/ListSupportedHostShapes). 
* `instance_display_name_prefix` - (Optional) A prefix used in the name of each ESXi host and Compute instance in the Cluster. If this isn't set, the Cluster's `displayName` is used as the prefix.

	For example, if the value is `myCluster`, the ESXi hosts are named `myCluster-1`, `myCluster-2`, and so on. 
* `is_shielded_instance_enabled` - (Optional) Indicates whether shielded instance is enabled for this Cluster. 
* `network_configuration` - (Required) (Updatable) The network configurations used by Cluster, including [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management subnet and VLANs. 
	* `hcx_vlan_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the HCX component of the VMware environment. This VLAN is a mandatory attribute  for Management Cluster when HCX is enabled.

		This attribute is not guaranteed to reflect the HCX VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the HCX VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the HCX component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/Sddc/UpdateSddc) to update the SDDC's `hcxVlanId` with that new VLAN's OCID. 
	* `nsx_edge_uplink1vlan_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the NSX Edge Uplink 1 component of the VMware environment. This VLAN is a mandatory attribute for Management Cluster.

		This attribute is not guaranteed to reflect the NSX Edge Uplink 1 VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the NSX Edge Uplink 1 VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the NSX Edge Uplink 1 component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `nsxEdgeUplink1VlanId` with that new VLAN's OCID. 
	* `nsx_edge_uplink2vlan_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC  for the NSX Edge Uplink 2 component of the VMware environment. This VLAN is a mandatory attribute for Management Cluster.

		This attribute is not guaranteed to reflect the NSX Edge Uplink 2 VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the NSX Edge Uplink 2 VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the NSX Edge Uplink 2 component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `nsxEdgeUplink2VlanId` with that new VLAN's OCID. 
	* `nsx_edge_vtep_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the NSX Edge VTEP component of the VMware environment.

		This attribute is not guaranteed to reflect the NSX Edge VTEP VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the NSX Edge VTEP VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the NSX Edge VTEP component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `nsxEdgeVTepVlanId` with that new VLAN's OCID. 
	* `nsx_vtep_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the NSX VTEP component of the VMware environment.

		This attribute is not guaranteed to reflect the NSX VTEP VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the NSX VTEP VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the NSX VTEP component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `nsxVTepVlanId` with that new VLAN's OCID. 
	* `provisioning_subnet_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management subnet used to provision the Cluster. 
	* `provisioning_vlan_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the Provisioning component of the VMware environment. 
	* `replication_vlan_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the vSphere Replication component of the VMware environment. 
	* `vmotion_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the vMotion component of the VMware environment.

		This attribute is not guaranteed to reflect the vMotion VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the vMotion VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the vMotion component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateCluster) to update the Cluster's `vmotionVlanId` with that new VLAN's OCID. 
	* `vsan_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the Cluster for the vSAN component of the VMware environment.

		This attribute is not guaranteed to reflect the vSAN VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the vSAN VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the vSAN component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Cluster/UpdateCluster) to update the Cluster's `vsanVlanId` with that new VLAN's OCID. 
	* `vsphere_vlan_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the vSphere component of the VMware environment. This VLAN is a mandatory attribute for Management Cluster.

		This attribute is not guaranteed to reflect the vSphere VLAN currently used by the ESXi hosts in the Cluster. The purpose of this attribute is to show the vSphere VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this Cluster in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

		Therefore, if you change the existing ESXi hosts in the Cluster to use a different VLAN for the vSphere component of the VMware environment, you should use [UpdateCluster](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the Cluster's `vsphereVlanId` with that new VLAN's OCID. 
* `sddc_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC that the Cluster belongs to. 
* `vmware_software_version` - (Optional) (Updatable) The VMware software bundle to install on the ESXi hosts in the Cluster. To get a list of the available versions, use [ListSupportedVmwareSoftwareVersions](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20230701/SupportedVmwareSoftwareVersionSummary/ListSupportedVmwareSoftwareVersions). 
* `workload_network_cidr` - (Optional) The CIDR block for the IP addresses that VMware VMs in the Cluster use to run application workloads. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster
	* `update` - (Defaults to 20 minutes), when updating the Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster


## Import

Clusters can be imported using the `id`, e.g.

```
$ terraform import oci_ocvp_cluster.test_cluster "id"
```

