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

