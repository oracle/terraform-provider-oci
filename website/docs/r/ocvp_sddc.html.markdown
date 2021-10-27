---
subcategory: "Oracle Cloud VMware Solution"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ocvp_sddc"
sidebar_current: "docs-oci-resource-ocvp-sddc"
description: |-
  Provides the Sddc resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service
---

# oci_ocvp_sddc
This resource provides the Sddc resource in Oracle Cloud Infrastructure Oracle Cloud VMware Solution service.

Creates an Oracle Cloud VMware Solution software-defined data center (SDDC).

Use the [WorkRequest](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/WorkRequest/) operations to track the
creation of the SDDC.

**Important:** You must configure the SDDC's networking resources with the security rules detailed in [Security Rules for Oracle Cloud VMware Solution SDDCs](https://docs.cloud.oracle.com/iaas/Content/VMware/Reference/ocvssecurityrules.htm). Otherwise, provisioning the SDDC will fail. The rules are based on the requirements set by VMware.


## Example Usage

```hcl
resource "oci_ocvp_sddc" "test_sddc" {
	#Required
	compartment_id = var.compartment_id
	compute_availability_domain = var.sddc_compute_availability_domain
	esxi_hosts_count = var.sddc_esxi_hosts_count
	nsx_edge_uplink1vlan_id = oci_core_vlan.test_nsx_edge_uplink1vlan.id
	nsx_edge_uplink2vlan_id = oci_core_vlan.test_nsx_edge_uplink2vlan.id
	nsx_edge_vtep_vlan_id = oci_core_vlan.test_nsx_edge_vtep_vlan.id
	nsx_vtep_vlan_id = oci_core_vlan.test_nsx_vtep_vlan.id
	provisioning_subnet_id = oci_core_subnet.test_subnet.id
	ssh_authorized_keys = var.sddc_ssh_authorized_keys
	vmotion_vlan_id = oci_core_vlan.test_vmotion_vlan.id
	vmware_software_version = var.sddc_vmware_software_version
	vsan_vlan_id = oci_core_vlan.test_vsan_vlan.id
	vsphere_vlan_id = oci_core_vlan.test_vsphere_vlan.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.sddc_display_name
	freeform_tags = {"Department"= "Finance"}
    hcx_action = var.hcx_action
	hcx_vlan_id = oci_core_vlan.test_vlan.id
	initial_sku = var.sddc_initial_sku
	instance_display_name_prefix = var.sddc_instance_display_name_prefix
	is_hcx_enabled = var.sddc_is_hcx_enabled
	is_hcx_enterprise_enabled = var.sddc_is_hcx_enterprise_enabled
	provisioning_vlan_id = oci_core_vlan.test_vlan.id
    refresh_hcx_license_status = true
	replication_vlan_id = oci_core_vlan.test_vlan.id
    reserving_hcx_on_premise_license_keys = var.reserving_hcx_on_premise_license_keys
	workload_network_cidr = var.sddc_workload_network_cidr
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the SDDC. 
* `compute_availability_domain` - (Required) The availability domain to create the SDDC's ESXi hosts in. For multi-AD SDDC deployment, set to `multi-AD`. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A descriptive name for the SDDC. SDDC name requirements are 1-16 character length limit, Must start with a letter, Must be English letters, numbers, - only, No repeating hyphens, Must be unique within the region. Avoid entering confidential information. 
* `esxi_hosts_count` - (Required) The number of ESXi hosts to create in the SDDC. You can add more hosts later (see [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost)).

	**Note:** If you later delete EXSi hosts from the SDDC to total less than 3, you are still billed for the 3 minimum recommended EXSi hosts. Also, you cannot add more VMware workloads to the SDDC until it again has at least 3 ESXi hosts. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hcx_action` - (Optional) (Updatable) The action to be performed upon HCX licenses. "UPGRADE" will upgrade the SDDC from HCX Advanced to HCX Enterprise. "DOWNGRADE" will downgrade the SDDC from HCX Enterprise to HCX Advanced after current HCX Enterprise billing cycle end date. "CANCEL_DOWNGRADE" will cancel the pending downgrade of HCX licenses. The action will only be performed when its value is changed. This field can also be used to enable HCX Enterprise during SDDC creation. If "UPGRADE" is set during SDDC creation, the SDDC will be created with HCX Enterprise enable. Supported actions during update: UPGRADE, DOWNGRADE, CANCEL_DOWNGRADE. Supported actions during creation: UPGRADE.
* `hcx_vlan_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the HCX component of the VMware environment. This value is required only when `isHcxEnabled` is true. 
* `initial_sku` - (Optional) The billing option selected during SDDC creation. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus). 
* `instance_display_name_prefix` - (Optional) A prefix used in the name of each ESXi host and Compute instance in the SDDC. If this isn't set, the SDDC's `displayName` is used as the prefix.

	For example, if the value is `mySDDC`, the ESXi hosts are named `mySDDC-1`, `mySDDC-2`, and so on. 
* `is_hcx_enabled` - (Optional) Indicates whether to enable HCX for this SDDC. 
* `nsx_edge_uplink1vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX Edge Uplink 1 component of the VMware environment. 
* `nsx_edge_uplink2vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX Edge Uplink 2 component of the VMware environment.

	**Note:** This VLAN is reserved for future use to deploy public-facing applications on the VMware SDDC. 
* `nsx_edge_vtep_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX Edge VTEP component of the VMware environment. 
* `nsx_vtep_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the NSX VTEP component of the VMware environment. 
* `provisioning_subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management subnet to use for provisioning the SDDC. 
* `provisioning_vlan_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the Provisioning component of the VMware environment. 
* `refresh_hcx_license_status` - (Optional) (Updatable) HCX on-premise licenses status will be refreshed whenever the value of this field is changed.
* `replication_vlan_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the vSphere Replication component of the VMware environment. 
* `reserving_hcx_on_premise_license_keys` - (Optional) (Updatable) The HCX on-premise licenses to be reserved when downgrade from HCX Enterprise to HCX Advanced. It should not be provided during resource creation. It is required and can only be set when the hcx_action is "DOWNGRADE". Its value can only be changed when hcx_action is updated.
* `ssh_authorized_keys` - (Required) (Updatable) One or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on each ESXi host. Use a newline character to separate multiple keys. The SSH keys must be in the format required for the `authorized_keys` file 
* `vmotion_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the vMotion component of the VMware environment. 
* `vmware_software_version` - (Required) (Updatable) The VMware software bundle to install on the ESXi hosts in the SDDC. To get a list of the available versions, use [ListSupportedVmwareSoftwareVersions](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedVmwareSoftwareVersionSummary/ListSupportedVmwareSoftwareVersions). 
* `vsan_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the vSAN component of the VMware environment. 
* `vsphere_vlan_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN to use for the vSphere component of the VMware environment. 
* `workload_network_cidr` - (Optional) The CIDR block for the IP addresses that VMware VMs in the SDDC use to run application workloads. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the SDDC. 
* `compute_availability_domain` - The availability domain the ESXi hosts are running in. For Multi-AD SDDC, it is `multi-AD`.  Example: `Uocm:PHX-AD-1`, `multi-AD` 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A descriptive name for the SDDC. It must be unique, start with a letter, and contain only letters, digits, whitespaces, dashes and underscores. Avoid entering confidential information. 
* `esxi_hosts_count` - The number of ESXi hosts in the SDDC.
* `actual_esxi_hosts_count` - The number of actual ESXi hosts in the SDDC on the cloud. This attribute will be different when esxi Host is added to an existing SDDC.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hcx_action` - The action to be performed upon HCX licenses.
* `hcx_fqdn` - The FQDN for HCX Manager.  Example: `hcx-my-sddc.sddc.us-phoenix-1.oraclecloud.com` 
* `hcx_initial_password` - The SDDC includes an administrator username and initial password for HCX Manager. Make sure to change this initial HCX Manager password to a different value. 
* `hcx_on_prem_key` - The activation keys to use on the on-premises HCX Enterprise appliances you site pair with HCX Manager in your VMware Solution. The number of keys provided depends on the HCX license type. HCX Advanced provides 3 activation keys.  HCX Enterprise provides 10 activation keys. 
* `hcx_on_prem_licenses` - The activation licenses to use on the on-premises HCX Enterprise appliance you site pair with HCX Manager in your VMware Solution. 
	* `activation_key` - HCX on-premise license key value.
	* `status` - status of HCX on-premise license.
	* `system_name` - Name of the system that consumed the HCX on-premise license
* `hcx_private_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is the virtual IP (VIP) for HCX Manager. For information about `PrivateIp` objects, see the Core Services API. 
* `hcx_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the HCX component of the VMware environment.

	This attribute is not guaranteed to reflect the HCX VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the HCX VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the HCX component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the SDDC's `hcxVlanId` with that new VLAN's OCID. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the SDDC. 
* `initial_sku` - The billing option selected during SDDC creation. [ListSupportedSkus](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedSkuSummary/ListSupportedSkus). 
* `instance_display_name_prefix` - A prefix used in the name of each ESXi host and Compute instance in the SDDC. If this isn't set, the SDDC's `displayName` is used as the prefix.

	For example, if the value is `MySDDC`, the ESXi hosts are named `MySDDC-1`, `MySDDC-2`, and so on. 
* `is_hcx_enabled` - Indicates whether HCX is enabled for this SDDC.
* `is_hcx_enterprise_enabled` - Indicates whether HCX Enterprise is enabled for this SDDC.
* `is_hcx_pending_downgrade` - Indicates whether SDDC is pending downgrade from HCX Enterprise to HCX Advanced.
* `nsx_edge_uplink1vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the NSX Edge Uplink 1 component of the VMware environment.

	This attribute is not guaranteed to reflect the NSX Edge Uplink 1 VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the NSX Edge Uplink 1 VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the NSX Edge Uplink 1 component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the SDDC's `nsxEdgeUplink1VlanId` with that new VLAN's OCID. 
* `nsx_edge_uplink2vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the NSX Edge Uplink 2 component of the VMware environment.

	This attribute is not guaranteed to reflect the NSX Edge Uplink 2 VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the NSX Edge Uplink 2 VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the NSX Edge Uplink 2 component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the SDDC's `nsxEdgeUplink2VlanId` with that new VLAN's OCID. 
* `nsx_edge_uplink_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is the virtual IP (VIP) for the NSX Edge Uplink. Use this OCID as the route target for route table rules when setting up connectivity between the SDDC and other networks. For information about `PrivateIp` objects, see the Core Services API. 
* `nsx_edge_vtep_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the NSX Edge VTEP component of the VMware environment.

	This attribute is not guaranteed to reflect the NSX Edge VTEP VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the NSX Edge VTEP VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the NSX Edge VTEP component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the SDDC's `nsxEdgeVTepVlanId` with that new VLAN's OCID. 
* `nsx_manager_fqdn` - The FQDN for NSX Manager.  Example: `nsx-my-sddc.sddc.us-phoenix-1.oraclecloud.com` 
* `nsx_manager_initial_password` - The SDDC includes an administrator username and initial password for NSX Manager. Make sure to change this initial NSX Manager password to a different value. 
* `nsx_manager_private_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is the virtual IP (VIP) for NSX Manager. For information about `PrivateIp` objects, see the Core Services API. 
* `nsx_manager_username` - The SDDC includes an administrator username and initial password for NSX Manager. You can change this initial username to a different value in NSX Manager. 
* `nsx_overlay_segment_name` - The VMware NSX overlay workload segment to host your application. Connect to workload portgroup in vCenter to access this overlay segment. 
* `nsx_vtep_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the NSX VTEP component of the VMware environment.

	This attribute is not guaranteed to reflect the NSX VTEP VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the NSX VTEP VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the NSX VTEP component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the SDDC's `nsxVTepVlanId` with that new VLAN's OCID. 
* `provisioning_subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management subnet used to provision the SDDC. 
* `provisioning_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the Provisioning component of the VMware environment. 
* `replication_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the vSphere Replication component of the VMware environment. 
* `reserving_hcx_on_premise_license_keys` - The HCX on-premise licenses to be reserved when downgrade from HCX Enterprise to HCX Advanced.
* `ssh_authorized_keys` - One or more public SSH keys to be included in the `~/.ssh/authorized_keys` file for the default user on each ESXi host. Use a newline character to separate multiple keys. The SSH keys must be in the format required for the `authorized_keys` file.

	This attribute is not guaranteed to reflect the public SSH keys currently installed on the ESXi hosts in the SDDC. The purpose of this attribute is to show the public SSH keys that Oracle Cloud VMware Solution will install on any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you upgrade the existing ESXi hosts in the SDDC to use different SSH keys, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the SDDC's `sshAuthorizedKeys` with the new public keys. 
* `state` - The current state of the SDDC.
* `time_created` - The date and time the SDDC was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_hcx_billing_cycle_end` - The date and time current HCX Enterprise billing cycle ends, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_hcx_license_status_updated` - The date and time the SDDC's HCX on-premise license status was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the SDDC was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `vcenter_fqdn` - The FQDN for vCenter.  Example: `vcenter-my-sddc.sddc.us-phoenix-1.oraclecloud.com` 
* `vcenter_initial_password` - The SDDC includes an administrator username and initial password for vCenter. Make sure to change this initial vCenter password to a different value. 
* `vcenter_private_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `PrivateIp` object that is the virtual IP (VIP) for vCenter. For information about `PrivateIp` objects, see the Core Services API. 
* `vcenter_username` - The SDDC includes an administrator username and initial password for vCenter. You can change this initial username to a different value in vCenter. 
* `vmotion_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the vMotion component of the VMware environment.

	This attribute is not guaranteed to reflect the vMotion VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the vMotion VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the vMotion component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the SDDC's `vmotionVlanId` with that new VLAN's OCID. 
* `vmware_software_version` - In general, this is a specific version of bundled VMware software supported by Oracle Cloud VMware Solution (see [ListSupportedVmwareSoftwareVersions](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/SupportedVmwareSoftwareVersionSummary/ListSupportedVmwareSoftwareVersions)).

	This attribute is not guaranteed to reflect the version of software currently installed on the ESXi hosts in the SDDC. The purpose of this attribute is to show the version of software that the Oracle Cloud VMware Solution will install on any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you upgrade the existing ESXi hosts in the SDDC to use a newer version of bundled VMware software supported by the Oracle Cloud VMware Solution, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the SDDC's `vmwareSoftwareVersion` with that new version. 
* `vsan_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the vSAN component of the VMware environment.

	This attribute is not guaranteed to reflect the vSAN VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the vSAN VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the vSAN component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the SDDC's `vsanVlanId` with that new VLAN's OCID. 
* `vsphere_vlan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN used by the SDDC for the vSphere component of the VMware environment.

	This attribute is not guaranteed to reflect the vSphere VLAN currently used by the ESXi hosts in the SDDC. The purpose of this attribute is to show the vSphere VLAN that the Oracle Cloud VMware Solution will use for any new ESXi hosts that you *add to this SDDC in the future* with [CreateEsxiHost](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/EsxiHost/CreateEsxiHost).

	Therefore, if you change the existing ESXi hosts in the SDDC to use a different VLAN for the vSphere component of the VMware environment, you should use [UpdateSddc](https://docs.cloud.oracle.com/iaas/api/#/en/vmware/20200501/Sddc/UpdateSddc) to update the SDDC's `vsphereVlanId` with that new VLAN's OCID. 
* `workload_network_cidr` - The CIDR block for the IP addresses that VMware VMs in the SDDC use to run application workloads. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 6 hours), when creating the Sddc
	* `update` - (Defaults to 20 minutes), when updating the Sddc
	* `delete` - (Defaults to 20 minutes), when destroying the Sddc


## Import

Sddcs can be imported using the `id`, e.g.

```
$ terraform import oci_ocvp_sddc.test_sddc "id"
```

