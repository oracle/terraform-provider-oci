// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

//fake

var (
	externalAssetKey = "terraform_provider_test_external_asset_key"
	assetSourceId    = `${oci_cloud_bridge_asset_source.test_asset_source.id}`
	assetId          = `${oci_cloud_bridge_asset.test_asset.id}`
	testDateTime     = time.Now().UTC().Truncate(time.Millisecond)

	inventoryAssetClassName    = "com.oracle.pic.ocb.discovery.model.OlvmStorageDomainAssetDetails"
	inventoryAssetClassVersion = "0"
	inventoryAssetDisplayName  = "kvm-storagedomain1"
	inventoryAssetDisplayName2 = "kvm-storagedomain1-updated"
	inventoryAssetExternalKey  = "1a10b288-f688-47d2-b38d-22dede44ba8a"
	inventoryAssetSourceKey    = "https://11.0.11.131:443/ovirt-engine/api"
	inventoryAssetDetailsJson  = `{"olvmStorageDomain":{"availableSpaceInBytes":643171352576,"storageDomainName":"kvm-storagedomain1"}}`
	inventoryAssetDetailsTf    = `{\"olvmStorageDomain\":{\"availableSpaceInBytes\":643171352576,\"storageDomainName\":\"kvm-storagedomain1\"}}`

	//VMWARE_VM
	CloudBridgeAssetRequiredOnlyResource = CloudBridgeAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Required, acctest.Create, CloudBridgeAssetRepresentation)

	CloudBridgeAssetResourceConfig = CloudBridgeAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Update, CloudBridgeAssetRepresentation)

	CloudBridgeCloudBridgeAssetSingularDataSourceRepresentation = map[string]interface{}{
		"asset_id": acctest.Representation{RepType: acctest.Required, Create: assetId},
	}

	CloudBridgeCloudBridgeAssetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"asset_id":           acctest.Representation{RepType: acctest.Optional, Create: assetId},
		"asset_type":         acctest.Representation{RepType: acctest.Optional, Create: `VMWARE_VM`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"external_asset_key": acctest.Representation{RepType: acctest.Optional, Create: externalAssetKey},
		"inventory_id":       acctest.Representation{RepType: acctest.Optional, Create: inventoryId},
		"source_key":         acctest.Representation{RepType: acctest.Optional, Create: `sourceKey`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetDataSourceFilterRepresentation}}
	CloudBridgeAssetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{assetId}},
	}

	CloudBridgeAssetRepresentation = map[string]interface{}{
		"asset_type":         acctest.Representation{RepType: acctest.Required, Create: `VMWARE_VM`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"external_asset_key": acctest.Representation{RepType: acctest.Required, Create: externalAssetKey},
		"inventory_id":       acctest.Representation{RepType: acctest.Required, Create: inventoryId},
		"source_key":         acctest.Representation{RepType: acctest.Required, Create: `sourceKey`},
		"asset_source_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{assetSourceId}, Update: []string{assetSourceId}},
		"compute":            acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetComputeRepresentation},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"vmware_vm":          acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetVmwareVmRepresentation},
		"vm":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetVmRepresentation},
		"vmware_vcenter":     acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetVmwareVCenterRepresentation},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}
	CloudBridgeAssetComputeRepresentation = map[string]interface{}{
		"connected_networks":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"cores_count":                acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"cpu_model":                  acctest.Representation{RepType: acctest.Optional, Create: `cpuModel`, Update: `cpuModel2`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"disks":                      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetComputeDisksRepresentation},
		"disks_count":                acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"dns_name":                   acctest.Representation{RepType: acctest.Optional, Create: `dnsName`, Update: `dnsName2`},
		"firmware":                   acctest.Representation{RepType: acctest.Optional, Create: `firmware`, Update: `firmware2`},
		"gpu_devices":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetComputeGpuDevicesRepresentation},
		"gpu_devices_count":          acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"guest_state":                acctest.Representation{RepType: acctest.Optional, Create: `guestState`, Update: `guestState2`},
		"hardware_version":           acctest.Representation{RepType: acctest.Optional, Create: `hardwareVersion`, Update: `hardwareVersion2`},
		"host_name":                  acctest.Representation{RepType: acctest.Optional, Create: `hostName`, Update: `hostName2`},
		"is_pmem_enabled":            acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_tpm_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"latency_sensitivity":        acctest.Representation{RepType: acctest.Optional, Create: `latencySensitivity`, Update: `latencySensitivity2`},
		"memory_in_mbs":              acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"nics":                       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetComputeNicsRepresentation},
		"nics_count":                 acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"nvdimm_controller":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetComputeNvdimmControllerRepresentation},
		"nvdimms":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetComputeNvdimmsRepresentation},
		"operating_system":           acctest.Representation{RepType: acctest.Optional, Create: `operatingSystem`, Update: `operatingSystem2`},
		"operating_system_version":   acctest.Representation{RepType: acctest.Optional, Create: `operatingSystemVersion`, Update: `operatingSystemVersion2`},
		"pmem_in_mbs":                acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"power_state":                acctest.Representation{RepType: acctest.Optional, Create: `powerState`, Update: `powerState2`},
		"primary_ip":                 acctest.Representation{RepType: acctest.Optional, Create: `primaryIp`, Update: `primaryIp2`},
		"scsi_controller":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetComputeScsiControllerRepresentation},
		"storage_provisioned_in_mbs": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"threads_per_core_count":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	CloudBridgeAssetVmRepresentation = map[string]interface{}{
		"hypervisor_host":    acctest.Representation{RepType: acctest.Required, Create: `hypervisorHost`, Update: `hypervisorHost2`},
		"hypervisor_vendor":  acctest.Representation{RepType: acctest.Required, Create: `hypervisorVendor`, Update: `hypervisorVendor2`},
		"hypervisor_version": acctest.Representation{RepType: acctest.Required, Create: `hypervisorVersion`, Update: `hypervisorVersion2`},
	}
	CloudBridgeAssetVmwareVCenterRepresentation = map[string]interface{}{
		"data_center":     acctest.Representation{RepType: acctest.Required, Create: `dataCenter`, Update: `dataCenter2`},
		"vcenter_key":     acctest.Representation{RepType: acctest.Required, Create: `vcenterKey`, Update: `vcenterKey2`},
		"vcenter_version": acctest.Representation{RepType: acctest.Required, Create: `vcenterVersion`, Update: `vcenterVersion2`},
	}
	CloudBridgeAssetVmwareVmRepresentation = map[string]interface{}{
		"cluster":                           acctest.Representation{RepType: acctest.Optional, Create: `cluster`, Update: `cluster2`},
		"customer_fields":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`customerFields`}, Update: []string{`customerFields2`}},
		"customer_tags":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetVmwareVmCustomerTagsRepresentation},
		"fault_tolerance_bandwidth":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"fault_tolerance_secondary_latency": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"fault_tolerance_state":             acctest.Representation{RepType: acctest.Optional, Create: `faultToleranceState`, Update: `faultToleranceState2`},
		"instance_uuid":                     acctest.Representation{RepType: acctest.Required, Create: `instanceUuid`, Update: `instanceUuid2`},
		"is_disks_cbt_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_disks_uuid_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"path":                              acctest.Representation{RepType: acctest.Required, Create: `path`, Update: `path2`},
		"vmware_tools_status":               acctest.Representation{RepType: acctest.Required, Create: `vmwareToolsStatus`, Update: `vmwareToolsStatus2`},
	}
	CloudBridgeAssetComputeDisksRepresentation = map[string]interface{}{
		"boot_order":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"is_cbt_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"location":        acctest.Representation{RepType: acctest.Optional, Create: `location`, Update: `location2`},
		"name":            acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"persistent_mode": acctest.Representation{RepType: acctest.Optional, Create: `persistentMode`, Update: `persistentMode2`},
		"size_in_mbs":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"uuid":            acctest.Representation{RepType: acctest.Optional, Create: `uuid`, Update: `uuid2`},
		"uuid_lun":        acctest.Representation{RepType: acctest.Optional, Create: `uuidLun`, Update: `uuidLun2`},
	}
	CloudBridgeAssetComputeGpuDevicesRepresentation = map[string]interface{}{
		"cores_count":   acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"manufacturer":  acctest.Representation{RepType: acctest.Optional, Create: `manufacturer`, Update: `manufacturer2`},
		"memory_in_mbs": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"name":          acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
	}
	CloudBridgeAssetComputeNicsRepresentation = map[string]interface{}{
		"ip_addresses":     acctest.Representation{RepType: acctest.Optional, Create: []string{`ipAddresses`}, Update: []string{`ipAddresses2`}},
		"label":            acctest.Representation{RepType: acctest.Optional, Create: `label`, Update: `label2`},
		"mac_address":      acctest.Representation{RepType: acctest.Optional, Create: `macAddress`, Update: `macAddress2`},
		"mac_address_type": acctest.Representation{RepType: acctest.Optional, Create: `macAddressType`, Update: `macAddressType2`},
		"network_name":     acctest.Representation{RepType: acctest.Optional, Create: `networkName`, Update: `networkName2`},
		"switch_name":      acctest.Representation{RepType: acctest.Optional, Create: `switchName`, Update: `switchName2`},
	}
	CloudBridgeAssetComputeNvdimmControllerRepresentation = map[string]interface{}{
		"bus_number": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"label":      acctest.Representation{RepType: acctest.Optional, Create: `label`, Update: `label2`},
	}
	CloudBridgeAssetComputeNvdimmsRepresentation = map[string]interface{}{
		"controller_key": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"label":          acctest.Representation{RepType: acctest.Optional, Create: `label`, Update: `label2`},
		"unit_number":    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	CloudBridgeAssetComputeScsiControllerRepresentation = map[string]interface{}{
		"label":       acctest.Representation{RepType: acctest.Optional, Create: `label`, Update: `label2`},
		"shared_bus":  acctest.Representation{RepType: acctest.Optional, Create: `sharedBus`, Update: `sharedBus2`},
		"unit_number": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}
	CloudBridgeAssetVmwareVmCustomerTagsRepresentation = map[string]interface{}{
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
	}

	//AWS_EC2
	CloudBridgeAwsEc2AssetRequiredOnlyResource = CloudBridgeAwsAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Required, acctest.Create, CloudBridgeAwsEc2AssetRepresentation)

	CloudBridgeAwsEc2AssetResourceConfig = CloudBridgeAwsAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Update, CloudBridgeAwsEc2AssetRepresentation)

	CloudBridgeAwsEc2AssetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"asset_id":           acctest.Representation{RepType: acctest.Optional, Create: assetId},
		"asset_type":         acctest.Representation{RepType: acctest.Optional, Create: `AWS_EC2`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"external_asset_key": acctest.Representation{RepType: acctest.Optional, Create: externalAssetKey},
		"inventory_id":       acctest.Representation{RepType: acctest.Optional, Create: inventoryId},
		"source_key":         acctest.Representation{RepType: acctest.Optional, Create: `sourceKey`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetDataSourceFilterRepresentation}}

	CloudBridgeAwsEc2AssetRepresentation = map[string]interface{}{
		"asset_type":                acctest.Representation{RepType: acctest.Required, Create: `AWS_EC2`},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"external_asset_key":        acctest.Representation{RepType: acctest.Required, Create: externalAssetKey},
		"inventory_id":              acctest.Representation{RepType: acctest.Required, Create: inventoryId},
		"source_key":                acctest.Representation{RepType: acctest.Required, Create: `sourceKey`},
		"asset_source_ids":          acctest.Representation{RepType: acctest.Optional, Create: []string{assetSourceId}, Update: []string{assetSourceId}},
		"compute":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetComputeRepresentation},
		"vm":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetVmRepresentation},
		"aws_ec2":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetAwsEc2Representation},
		"aws_ec2cost":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEc2CostRepresentation},
		"attached_ebs_volumes_cost": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAttachedEbsVolumesCostRepresentation},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}
	CloudBridgeAssetAwsEc2Representation = map[string]interface{}{
		"architecture": acctest.Representation{RepType: acctest.Required, Create: `architecture`, Update: `architecture2`},
		"are_elastic_inference_accelerators_present": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"boot_mode":                acctest.Representation{RepType: acctest.Optional, Create: `bootMode`, Update: `bootMode2`},
		"capacity_reservation_key": acctest.Representation{RepType: acctest.Optional, Create: `capacityReservationKey`, Update: `capacityReservationKey2`},
		"image_key":                acctest.Representation{RepType: acctest.Optional, Create: `imageKey`, Update: `imageKey2`},
		"instance_key":             acctest.Representation{RepType: acctest.Required, Create: `instanceKey`, Update: `instanceKey2`},
		"instance_lifecycle":       acctest.Representation{RepType: acctest.Optional, Create: `instanceLifecycle`, Update: `instanceLifecycle2`},
		"instance_type":            acctest.Representation{RepType: acctest.Required, Create: `instanceType`, Update: `instanceType2`},
		"ip_address":               acctest.Representation{RepType: acctest.Optional, Create: `ipAddress`, Update: `ipAddress2`},
		"ipv6address":              acctest.Representation{RepType: acctest.Optional, Create: `ipv6Address`, Update: `ipv6Address2`},
		"is_enclave_options":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_hibernation_options":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_source_dest_check":     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_spot_instance":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"kernel_key":               acctest.Representation{RepType: acctest.Optional, Create: `kernelKey`, Update: `kernelKey2`},
		"licenses":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`licenses`}, Update: []string{`licenses2`}},
		"maintenance_options":      acctest.Representation{RepType: acctest.Optional, Create: `maintenanceOptions`, Update: `maintenanceOptions2`},
		"monitoring":               acctest.Representation{RepType: acctest.Optional, Create: `monitoring`, Update: `monitoring2`},
		"network_interfaces":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEc2NetworkInterfacesRepresentation},
		"placement":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEc2PlacementRepresentation},
		"private_dns_name":         acctest.Representation{RepType: acctest.Optional, Create: `privateDnsName`, Update: `privateDnsName2`},
		"private_ip_address":       acctest.Representation{RepType: acctest.Optional, Create: `privateIpAddress`, Update: `privateIpAddress2`},
		"root_device_name":         acctest.Representation{RepType: acctest.Required, Create: `rootDeviceName`, Update: `rootDeviceName2`},
		"root_device_type":         acctest.Representation{RepType: acctest.Optional, Create: `rootDeviceType`, Update: `rootDeviceType2`},
		"security_groups":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEc2SecurityGroupsRepresentation},
		"sriov_net_support":        acctest.Representation{RepType: acctest.Optional, Create: `sriovNetSupport`, Update: `sriovNetSupport2`},
		"state":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetAwsEc2StateRepresentation},
		"subnet_key":               acctest.Representation{RepType: acctest.Optional, Create: `subnetKey`, Update: `subnetKey2`},
		"tags":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEc2TagsRepresentation},
		"time_launch":              acctest.Representation{RepType: acctest.Optional, Create: testDateTime.Format(time.RFC3339Nano), Update: testDateTime.Format(time.RFC3339Nano)},
		"tpm_support":              acctest.Representation{RepType: acctest.Optional, Create: `tpmSupport`, Update: `tpmSupport2`},
		"virtualization_type":      acctest.Representation{RepType: acctest.Optional, Create: `virtualizationType`, Update: `virtualizationType2`},
		"vpc_key":                  acctest.Representation{RepType: acctest.Optional, Create: `vpcKey`, Update: `vpcKey2`},
	}
	CloudBridgeAssetAwsEc2CostRepresentation = map[string]interface{}{
		"amount":        acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `2.0`},
		"currency_code": acctest.Representation{RepType: acctest.Optional, Create: `currencyCode`, Update: `currencyCode2`},
	}
	CloudBridgeAssetAttachedEbsVolumesCostRepresentation = map[string]interface{}{
		"amount":        acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `2.0`},
		"currency_code": acctest.Representation{RepType: acctest.Optional, Create: `currencyCode`, Update: `currencyCode2`},
	}
	CloudBridgeAssetAwsEc2NetworkInterfacesRepresentation = map[string]interface{}{
		"association":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEc2NetworkInterfacesAssociationRepresentation},
		"attachment":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEc2NetworkInterfacesAttachmentRepresentation},
		"description":           acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"interface_type":        acctest.Representation{RepType: acctest.Optional, Create: `interfaceType`, Update: `interfaceType2`},
		"ipv4prefixes":          acctest.Representation{RepType: acctest.Optional, Create: []string{`ipv4Prefixes`}, Update: []string{`ipv4Prefixes2`}},
		"ipv6addresses":         acctest.Representation{RepType: acctest.Optional, Create: []string{`ipv6Addresses`}, Update: []string{`ipv6Addresses2`}},
		"ipv6prefixes":          acctest.Representation{RepType: acctest.Optional, Create: []string{`ipv6Prefixes`}, Update: []string{`ipv6Prefixes2`}},
		"is_source_dest_check":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"mac_address":           acctest.Representation{RepType: acctest.Optional, Create: `macAddress`, Update: `macAddress2`},
		"network_interface_key": acctest.Representation{RepType: acctest.Optional, Create: `networkInterfaceKey`, Update: `networkInterfaceKey2`},
		"owner_key":             acctest.Representation{RepType: acctest.Optional, Create: `ownerKey`, Update: `ownerKey2`},
		"private_ip_addresses":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEc2NetworkInterfacesPrivateIpAddressesRepresentation},
		"security_groups":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEc2NetworkInterfacesSecurityGroupsRepresentation},
		"status":                acctest.Representation{RepType: acctest.Optional, Create: `status`, Update: `status2`},
		"subnet_key":            acctest.Representation{RepType: acctest.Optional, Create: `subnetKey`, Update: `subnetKey2`},
	}
	CloudBridgeAssetAwsEc2PlacementRepresentation = map[string]interface{}{
		"affinity":                acctest.Representation{RepType: acctest.Optional, Create: `affinity`, Update: `affinity2`},
		"availability_zone":       acctest.Representation{RepType: acctest.Optional, Create: `availabilityZone`, Update: `availabilityZone2`},
		"group_name":              acctest.Representation{RepType: acctest.Optional, Create: `groupName1`, Update: `groupName2`},
		"host_key":                acctest.Representation{RepType: acctest.Optional, Create: `hostKey`, Update: `hostKey2`},
		"host_resource_group_arn": acctest.Representation{RepType: acctest.Optional, Create: `hostResourceGroupArn`, Update: `hostResourceGroupArn2`},
		"partition_number":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"spread_domain":           acctest.Representation{RepType: acctest.Optional, Create: `spreadDomain`, Update: `spreadDomain2`},
		"tenancy":                 acctest.Representation{RepType: acctest.Optional, Create: `tenancy`, Update: `tenancy2`},
	}
	CloudBridgeAssetAwsEc2SecurityGroupsRepresentation = map[string]interface{}{
		"group_key":  acctest.Representation{RepType: acctest.Optional, Create: `groupKey`, Update: `groupKey2`},
		"group_name": acctest.Representation{RepType: acctest.Optional, Create: `groupName1`, Update: `groupName2`},
	}
	CloudBridgeAssetAwsEc2StateRepresentation = map[string]interface{}{
		"code": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"name": acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
	}
	CloudBridgeAssetAwsEc2TagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}
	CloudBridgeAssetAwsEc2NetworkInterfacesAssociationRepresentation = map[string]interface{}{
		"carrier_ip":        acctest.Representation{RepType: acctest.Optional, Create: `carrierIp`, Update: `carrierIp2`},
		"customer_owned_ip": acctest.Representation{RepType: acctest.Optional, Create: `customerOwnedIp`, Update: `customerOwnedIp2`},
		"ip_owner_key":      acctest.Representation{RepType: acctest.Optional, Create: `ipOwnerKey`, Update: `ipOwnerKey2`},
		"public_dns_name":   acctest.Representation{RepType: acctest.Optional, Create: `publicDnsName`, Update: `publicDnsName2`},
		"public_ip":         acctest.Representation{RepType: acctest.Optional, Create: `publicIp`, Update: `publicIp2`},
	}
	CloudBridgeAssetAwsEc2NetworkInterfacesAttachmentRepresentation = map[string]interface{}{
		"attachment_key":           acctest.Representation{RepType: acctest.Optional, Create: `attachmentKey`, Update: `attachmentKey2`},
		"device_index":             acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"is_delete_on_termination": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"network_card_index":       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"status":                   acctest.Representation{RepType: acctest.Optional, Create: `status`, Update: `status2`},
		"time_attach":              acctest.Representation{RepType: acctest.Optional, Create: testDateTime.Format(time.RFC3339Nano), Update: testDateTime.Format(time.RFC3339Nano)},
	}
	CloudBridgeAssetAwsEc2NetworkInterfacesPrivateIpAddressesRepresentation = map[string]interface{}{
		"association":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEc2NetworkInterfacesPrivateIpAddressesAssociationRepresentation},
		"is_primary":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"private_dns_name":   acctest.Representation{RepType: acctest.Optional, Create: `privateDnsName`, Update: `privateDnsName2`},
		"private_ip_address": acctest.Representation{RepType: acctest.Optional, Create: `privateIpAddress`, Update: `privateIpAddress2`},
	}
	CloudBridgeAssetAwsEc2NetworkInterfacesSecurityGroupsRepresentation = map[string]interface{}{
		"group_key":  acctest.Representation{RepType: acctest.Optional, Create: `groupKey`, Update: `groupKey2`},
		"group_name": acctest.Representation{RepType: acctest.Optional, Create: `groupName1`, Update: `groupName2`},
	}
	CloudBridgeAssetAwsEc2NetworkInterfacesPrivateIpAddressesAssociationRepresentation = map[string]interface{}{
		"carrier_ip":        acctest.Representation{RepType: acctest.Optional, Create: `carrierIp`, Update: `carrierIp2`},
		"customer_owned_ip": acctest.Representation{RepType: acctest.Optional, Create: `customerOwnedIp`, Update: `customerOwnedIp2`},
		"ip_owner_key":      acctest.Representation{RepType: acctest.Optional, Create: `ipOwnerKey`, Update: `ipOwnerKey2`},
		"public_dns_name":   acctest.Representation{RepType: acctest.Optional, Create: `publicDnsName`, Update: `publicDnsName2`},
		"public_ip":         acctest.Representation{RepType: acctest.Optional, Create: `publicIp`, Update: `publicIp2`},
	}
	//AWS_EBS
	CloudBridgeAwsEbsAssetRequiredOnlyResource = CloudBridgeAwsAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Required, acctest.Create, CloudBridgeAwsEbsAssetRepresentation)

	CloudBridgeAwsEbsAssetResourceConfig = CloudBridgeAwsAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Update, CloudBridgeAwsEbsAssetRepresentation)

	CloudBridgeAwsEbsAssetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"asset_id":           acctest.Representation{RepType: acctest.Optional, Create: assetId},
		"asset_type":         acctest.Representation{RepType: acctest.Optional, Create: `AWS_EBS`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"external_asset_key": acctest.Representation{RepType: acctest.Optional, Create: externalAssetKey},
		"inventory_id":       acctest.Representation{RepType: acctest.Optional, Create: inventoryId},
		"source_key":         acctest.Representation{RepType: acctest.Optional, Create: `sourceKey`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetDataSourceFilterRepresentation}}

	CloudBridgeAwsEbsAssetRepresentation = map[string]interface{}{
		"asset_type":         acctest.Representation{RepType: acctest.Required, Create: `AWS_EBS`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"external_asset_key": acctest.Representation{RepType: acctest.Required, Create: externalAssetKey},
		"inventory_id":       acctest.Representation{RepType: acctest.Required, Create: inventoryId},
		"source_key":         acctest.Representation{RepType: acctest.Required, Create: `sourceKey`},
		"aws_ebs":            acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetAwsEbsRepresentation},
		"asset_source_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{assetSourceId}, Update: []string{assetSourceId}},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}
	CloudBridgeAssetAwsEbsRepresentation = map[string]interface{}{
		"attachments":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEbsAttachmentsRepresentation},
		"availability_zone":       acctest.Representation{RepType: acctest.Optional, Create: `availabilityZone`, Update: `availabilityZone2`},
		"iops":                    acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"is_encrypted":            acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"is_multi_attach_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"size_in_gi_bs":           acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"status":                  acctest.Representation{RepType: acctest.Optional, Create: `status`, Update: `status2`},
		"tags":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetAwsEbsTagsRepresentation},
		"throughput":              acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"volume_key":              acctest.Representation{RepType: acctest.Required, Create: `volumeKey`, Update: `volumeKey2`},
		"volume_type":             acctest.Representation{RepType: acctest.Required, Create: `volumeType`, Update: `volumeType2`},
	}
	CloudBridgeAssetAwsEbsAttachmentsRepresentation = map[string]interface{}{
		"device":                   acctest.Representation{RepType: acctest.Optional, Create: `device`, Update: `device2`},
		"instance_key":             acctest.Representation{RepType: acctest.Optional, Create: `instanceKey`, Update: `instanceKey2`},
		"is_delete_on_termination": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"status":                   acctest.Representation{RepType: acctest.Optional, Create: `status`, Update: `status2`},
		"volume_key":               acctest.Representation{RepType: acctest.Optional, Create: `volumeKey`, Update: `volumeKey2`},
	}
	CloudBridgeAssetAwsEbsTagsRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Optional, Create: `key`, Update: `key2`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `value`, Update: `value2`},
	}

	CloudBridgeInventoryAssetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"asset_id":           acctest.Representation{RepType: acctest.Optional, Create: assetId},
		"asset_type":         acctest.Representation{RepType: acctest.Optional, Create: `INVENTORY_ASSET`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: inventoryAssetDisplayName2},
		"external_asset_key": acctest.Representation{RepType: acctest.Optional, Create: inventoryAssetExternalKey},
		"inventory_id":       acctest.Representation{RepType: acctest.Optional, Create: inventoryId},
		"source_key":         acctest.Representation{RepType: acctest.Optional, Create: inventoryAssetSourceKey},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetDataSourceFilterRepresentation}}

	CloudBridgeInventoryAssetRequiredOnlyResource = CloudBridgeOlvmAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Required, acctest.Create, CloudBridgeInventoryAssetRepresentation)

	CloudBridgeInventoryAssetResourceConfig = CloudBridgeOlvmAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Update, CloudBridgeInventoryAssetRepresentation)

	CloudBridgeInventoryAssetRepresentation = map[string]interface{}{
		"asset_type":          acctest.Representation{RepType: acctest.Required, Create: `INVENTORY_ASSET`},
		"asset_class_name":    acctest.Representation{RepType: acctest.Required, Create: inventoryAssetClassName},
		"asset_class_version": acctest.Representation{RepType: acctest.Required, Create: inventoryAssetClassVersion},
		"asset_details":       acctest.Representation{RepType: acctest.Required, Create: inventoryAssetDetailsTf, Update: inventoryAssetDetailsTf},
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"external_asset_key":  acctest.Representation{RepType: acctest.Required, Create: inventoryAssetExternalKey},
		"inventory_id":        acctest.Representation{RepType: acctest.Required, Create: inventoryId},
		"source_key":          acctest.Representation{RepType: acctest.Required, Create: inventoryAssetSourceKey},
		"asset_source_ids":    acctest.Representation{RepType: acctest.Optional, Create: []string{assetSourceId}, Update: []string{assetSourceId}},
		"display_name":        acctest.Representation{RepType: acctest.Optional, Create: inventoryAssetDisplayName, Update: inventoryAssetDisplayName2},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}

	CloudBridgeAssetResourceDependencies = CloudBridgeAssetSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Required, acctest.Create, CloudBridgeAssetSourceRepresentation)

	CloudBridgeOlvmAssetResourceDependencies = CloudBridgeAssetSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Required, acctest.Create, CloudBridgeOlvmAssetSourceRepresentation)

	CloudBridgeAwsAssetResourceDependencies = CloudBridgeAssetSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset_source", "test_asset_source", acctest.Required, acctest.Create, CloudBridgeAwsAssetSourceRepresentation)
)

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	vaultSecretId := utils.GetEnvSettingWithBlankDefault("vaultSecretId")
	vaultSecretIdVariableStr := fmt.Sprintf("variable \"vaultSecretId\" { default = \"%s\" }\n", vaultSecretId)

	inventoryId := utils.GetEnvSettingWithBlankDefault("inventoryId")
	inventoryIdVariableStr := fmt.Sprintf("variable \"inventoryId\" { default = \"%s\" }\n", inventoryId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	variableStr := compartmentIdVariableStr + inventoryIdVariableStr + vaultSecretIdVariableStr

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_bridge_asset.test_asset"
	datasourceName := "data.oci_cloud_bridge_assets.test_assets"
	singularDatasourceName := "data.oci_cloud_bridge_asset.test_asset"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+CloudBridgeAssetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Create, CloudBridgeAssetRepresentation), "cloudbridge", "asset", t)

	acctest.ResourceTest(t, testAccCheckCloudBridgeAssetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variableStr + CloudBridgeAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Required, acctest.Create, CloudBridgeAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_type", "VMWARE_VM"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttr(resourceName, "compute.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + variableStr + CloudBridgeAssetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + variableStr + CloudBridgeAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Create, CloudBridgeAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "VMWARE_VM"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.connected_networks", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cores_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cpu_model", "cpuModel"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.boot_order", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.is_cbt_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.location", "location"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.persistent_mode", "persistentMode"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.size_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid", "uuid"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid_lun", "uuidLun"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.dns_name", "dnsName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.firmware", "firmware"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.cores_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.manufacturer", "manufacturer"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.memory_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.guest_state", "guestState"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.hardware_version", "hardwareVersion"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.host_name", "hostName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_pmem_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_tpm_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.latency_sensitivity", "latencySensitivity"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.memory_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address", "macAddress"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address_type", "macAddressType"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.network_name", "networkName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.switch_name", "switchName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.bus_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.controller_key", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.unit_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system", "operatingSystem"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system_version", "operatingSystemVersion"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.pmem_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.power_state", "powerState"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.primary_ip", "primaryIp"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.shared_bus", "sharedBus"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.unit_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.storage_provisioned_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.threads_per_core_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.cluster", "cluster"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_fields.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_tags.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_tags.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.fault_tolerance_bandwidth", "10"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.fault_tolerance_secondary_latency", "10"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.fault_tolerance_state", "faultToleranceState"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.instance_uuid", "instanceUuid"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.is_disks_cbt_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.is_disks_uuid_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.path", "path"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.vmware_tools_status", "vmwareToolsStatus"),
				resource.TestCheckResourceAttr(resourceName, "vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_host", "hypervisorHost"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_vendor", "hypervisorVendor"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_version", "hypervisorVersion"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.0.data_center", "dataCenter"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.0.vcenter_key", "vcenterKey"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.0.vcenter_version", "vcenterVersion"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + variableStr + compartmentIdUVariableStr + CloudBridgeAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudBridgeAssetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "VMWARE_VM"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "compute.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.connected_networks", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cores_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cpu_model", "cpuModel"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.boot_order", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.is_cbt_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.location", "location"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.persistent_mode", "persistentMode"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.size_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid", "uuid"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid_lun", "uuidLun"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.dns_name", "dnsName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.firmware", "firmware"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.cores_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.manufacturer", "manufacturer"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.memory_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.guest_state", "guestState"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.hardware_version", "hardwareVersion"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.host_name", "hostName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_pmem_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_tpm_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.latency_sensitivity", "latencySensitivity"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.memory_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address", "macAddress"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address_type", "macAddressType"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.network_name", "networkName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.switch_name", "switchName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.bus_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.controller_key", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.unit_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system", "operatingSystem"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system_version", "operatingSystemVersion"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.pmem_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.power_state", "powerState"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.primary_ip", "primaryIp"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.shared_bus", "sharedBus"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.unit_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.storage_provisioned_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.threads_per_core_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.cluster", "cluster"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_fields.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_tags.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_tags.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.fault_tolerance_bandwidth", "10"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.fault_tolerance_secondary_latency", "10"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.fault_tolerance_state", "faultToleranceState"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.instance_uuid", "instanceUuid"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.is_disks_cbt_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.is_disks_uuid_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.path", "path"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.vmware_tools_status", "vmwareToolsStatus"),
				resource.TestCheckResourceAttr(resourceName, "vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_host", "hypervisorHost"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_vendor", "hypervisorVendor"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_version", "hypervisorVersion"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.0.data_center", "dataCenter"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.0.vcenter_key", "vcenterKey"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.0.vcenter_version", "vcenterVersion"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + variableStr + CloudBridgeAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Update, CloudBridgeAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "VMWARE_VM"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.connected_networks", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cores_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cpu_model", "cpuModel2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.boot_order", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.is_cbt_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.location", "location2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.persistent_mode", "persistentMode2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.size_in_mbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid", "uuid2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid_lun", "uuidLun2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.dns_name", "dnsName2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.firmware", "firmware2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.cores_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.manufacturer", "manufacturer2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.memory_in_mbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.guest_state", "guestState2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.hardware_version", "hardwareVersion2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.host_name", "hostName2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_pmem_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_tpm_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.latency_sensitivity", "latencySensitivity2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.memory_in_mbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.label", "label2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address", "macAddress2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address_type", "macAddressType2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.network_name", "networkName2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.switch_name", "switchName2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.bus_number", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.label", "label2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.controller_key", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.label", "label2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.unit_number", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system", "operatingSystem2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system_version", "operatingSystemVersion2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.pmem_in_mbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.power_state", "powerState2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.primary_ip", "primaryIp2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.label", "label2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.shared_bus", "sharedBus2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.unit_number", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.storage_provisioned_in_mbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.threads_per_core_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.cluster", "cluster2"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_fields.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_tags.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.customer_tags.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.fault_tolerance_bandwidth", "11"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.fault_tolerance_secondary_latency", "11"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.fault_tolerance_state", "faultToleranceState2"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.instance_uuid", "instanceUuid2"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.is_disks_cbt_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.is_disks_uuid_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.path", "path2"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vm.0.vmware_tools_status", "vmwareToolsStatus2"),
				resource.TestCheckResourceAttr(resourceName, "vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_host", "hypervisorHost2"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_vendor", "hypervisorVendor2"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_version", "hypervisorVersion2"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.0.data_center", "dataCenter2"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.0.vcenter_key", "vcenterKey2"),
				resource.TestCheckResourceAttr(resourceName, "vmware_vcenter.0.vcenter_version", "vcenterVersion2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_assets", "test_assets", acctest.Optional, acctest.Update, CloudBridgeCloudBridgeAssetDataSourceRepresentation) +
				variableStr + CloudBridgeAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Update, CloudBridgeAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "asset_id"),
				resource.TestCheckResourceAttr(datasourceName, "asset_type", "VMWARE_VM"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(datasourceName, "inventory_id"),
				resource.TestCheckResourceAttr(datasourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "asset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "asset_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Required, acctest.Create, CloudBridgeCloudBridgeAssetSingularDataSourceRepresentation) +
				variableStr + CloudBridgeAssetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "asset_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "asset_type", "VMWARE_VM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.connected_networks", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.cores_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.cpu_model", "cpuModel2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.boot_order", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.is_cbt_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.location", "location2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.persistent_mode", "persistentMode2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.size_in_mbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.uuid", "uuid2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.uuid_lun", "uuidLun2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.dns_name", "dnsName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.firmware", "firmware2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.0.cores_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.0.manufacturer", "manufacturer2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.0.memory_in_mbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.guest_state", "guestState2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.hardware_version", "hardwareVersion2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.host_name", "hostName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.is_pmem_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.is_tpm_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.latency_sensitivity", "latencySensitivity2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.memory_in_mbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.label", "label2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.mac_address", "macAddress2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.mac_address_type", "macAddressType2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.network_name", "networkName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.switch_name", "switchName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimm_controller.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimm_controller.0.bus_number", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimm_controller.0.label", "label2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimms.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimms.0.controller_key", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimms.0.label", "label2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimms.0.unit_number", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.operating_system", "operatingSystem2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.operating_system_version", "operatingSystemVersion2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.pmem_in_mbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.power_state", "powerState2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.primary_ip", "primaryIp2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.scsi_controller.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.scsi_controller.0.label", "label2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.scsi_controller.0.shared_bus", "sharedBus2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.scsi_controller.0.unit_number", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.storage_provisioned_in_mbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.threads_per_core_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "environment_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.cluster", "cluster2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.customer_fields.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.customer_tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.customer_tags.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.customer_tags.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.fault_tolerance_bandwidth", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.fault_tolerance_secondary_latency", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.fault_tolerance_state", "faultToleranceState2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.instance_uuid", "instanceUuid2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.is_disks_cbt_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.is_disks_uuid_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.path", "path2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vm.0.vmware_tools_status", "vmwareToolsStatus2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm.0.hypervisor_host", "hypervisorHost2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm.0.hypervisor_vendor", "hypervisorVendor2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm.0.hypervisor_version", "hypervisorVersion2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vcenter.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vcenter.0.data_center", "dataCenter2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vcenter.0.vcenter_key", "vcenterKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_vcenter.0.vcenter_version", "vcenterVersion2"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudBridgeAssetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func TestCloudBridgeAwsEc2AssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeAwsEc2AssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	vaultSecretId := utils.GetEnvSettingWithBlankDefault("vaultSecretId")
	vaultSecretIdVariableStr := fmt.Sprintf("variable \"vaultSecretId\" { default = \"%s\" }\n", vaultSecretId)

	awsAccountKey := utils.GetEnvSettingWithBlankDefault("awsAccountKey")
	awsAccountKeyVariableStr := fmt.Sprintf("variable \"awsAccountKey\" { default = \"%s\" }\n", awsAccountKey)

	inventoryId := utils.GetEnvSettingWithBlankDefault("inventoryId")
	inventoryIdVariableStr := fmt.Sprintf("variable \"inventoryId\" { default = \"%s\" }\n", inventoryId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	variableStr := compartmentIdVariableStr + vaultSecretIdVariableStr + awsAccountKeyVariableStr + inventoryIdVariableStr

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_bridge_asset.test_asset"
	datasourceName := "data.oci_cloud_bridge_assets.test_assets"
	singularDatasourceName := "data.oci_cloud_bridge_asset.test_asset"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+CloudBridgeAwsAssetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
			acctest.Optional, acctest.Create, CloudBridgeAwsEc2AssetRepresentation), "cloudbridge", "asset", t)

	acctest.ResourceTest(t, testAccCheckCloudBridgeAssetDestroy, []resource.TestStep{
		// AWS EC2 verify Create
		{
			Config: config + variableStr + CloudBridgeAwsAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
					acctest.Required, acctest.Create, CloudBridgeAwsEc2AssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_type", "AWS_EC2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.architecture", "architecture"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_key", "instanceKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_type", "instanceType"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.root_device_name", "rootDeviceName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.state.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttr(resourceName, "vm.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// AWS EC2 delete before next Create
		{
			Config: config + variableStr + CloudBridgeAwsAssetResourceDependencies,
		},
		// AWS EC2 verify Create with optionals
		{
			Config: config + variableStr + CloudBridgeAwsAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
					acctest.Optional, acctest.Create, CloudBridgeAwsEc2AssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "AWS_EC2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "attached_ebs_volumes_cost.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attached_ebs_volumes_cost.0.amount", "1"),
				resource.TestCheckResourceAttr(resourceName, "attached_ebs_volumes_cost.0.currency_code", "currencyCode"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.architecture", "architecture"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.are_elastic_inference_accelerators_present", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.boot_mode", "bootMode"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.capacity_reservation_key", "capacityReservationKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.image_key", "imageKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_key", "instanceKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_lifecycle", "instanceLifecycle"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_type", "instanceType"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.ip_address", "ipAddress"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.ipv6address", "ipv6Address"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_enclave_options", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_hibernation_options", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_source_dest_check", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_spot_instance", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.kernel_key", "kernelKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.licenses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.maintenance_options", "maintenanceOptions"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.monitoring", "monitoring"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.carrier_ip", "carrierIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.customer_owned_ip", "customerOwnedIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.ip_owner_key", "ipOwnerKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.public_dns_name", "publicDnsName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.public_ip", "publicIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.attachment_key", "attachmentKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.device_index", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.is_delete_on_termination", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.network_card_index", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.status", "status"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.time_attach", testDateTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.interface_type", "interfaceType"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.ipv4prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.ipv6addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.ipv6prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.is_source_dest_check", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.mac_address", "macAddress"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.network_interface_key", "networkInterfaceKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.owner_key", "ownerKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.carrier_ip", "carrierIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.customer_owned_ip", "customerOwnedIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.ip_owner_key", "ipOwnerKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.public_dns_name", "publicDnsName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.public_ip", "publicIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.is_primary", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.private_dns_name", "privateDnsName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.private_ip_address", "privateIpAddress"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.security_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.security_groups.0.group_key", "groupKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.security_groups.0.group_name", "groupName1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.status", "status"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.subnet_key", "subnetKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.affinity", "affinity"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.availability_zone", "availabilityZone"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.group_name", "groupName1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.host_key", "hostKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.host_resource_group_arn", "hostResourceGroupArn"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.partition_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.spread_domain", "spreadDomain"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.tenancy", "tenancy"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.private_dns_name", "privateDnsName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.private_ip_address", "privateIpAddress"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.root_device_name", "rootDeviceName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.root_device_type", "rootDeviceType"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.security_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.security_groups.0.group_key", "groupKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.security_groups.0.group_name", "groupName1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.sriov_net_support", "sriovNetSupport"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.state.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.state.0.code", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.state.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.subnet_key", "subnetKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.time_launch", testDateTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tpm_support", "tpmSupport"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.virtualization_type", "virtualizationType"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.vpc_key", "vpcKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2cost.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2cost.0.amount", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2cost.0.currency_code", "currencyCode"),
				resource.TestCheckResourceAttr(resourceName, "compute.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.connected_networks", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cores_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cpu_model", "cpuModel"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.boot_order", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.is_cbt_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.location", "location"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.persistent_mode", "persistentMode"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.size_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid", "uuid"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid_lun", "uuidLun"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.dns_name", "dnsName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.firmware", "firmware"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.cores_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.manufacturer", "manufacturer"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.memory_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.guest_state", "guestState"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.hardware_version", "hardwareVersion"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.host_name", "hostName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_pmem_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_tpm_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.latency_sensitivity", "latencySensitivity"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.memory_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address", "macAddress"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address_type", "macAddressType"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.network_name", "networkName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.switch_name", "switchName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.bus_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.controller_key", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.unit_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system", "operatingSystem"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system_version", "operatingSystemVersion"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.pmem_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.power_state", "powerState"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.primary_ip", "primaryIp"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.shared_bus", "sharedBus"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.unit_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.storage_provisioned_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.threads_per_core_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_host", "hypervisorHost"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_vendor", "hypervisorVendor"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_version", "hypervisorVersion"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// AWS EC2 verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + variableStr + compartmentIdUVariableStr + CloudBridgeAwsAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudBridgeAwsEc2AssetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "AWS_EC2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "attached_ebs_volumes_cost.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attached_ebs_volumes_cost.0.amount", "1"),
				resource.TestCheckResourceAttr(resourceName, "attached_ebs_volumes_cost.0.currency_code", "currencyCode"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.architecture", "architecture"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.are_elastic_inference_accelerators_present", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.boot_mode", "bootMode"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.capacity_reservation_key", "capacityReservationKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.image_key", "imageKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_key", "instanceKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_lifecycle", "instanceLifecycle"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_type", "instanceType"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.ip_address", "ipAddress"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.ipv6address", "ipv6Address"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_enclave_options", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_hibernation_options", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_source_dest_check", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_spot_instance", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.kernel_key", "kernelKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.licenses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.maintenance_options", "maintenanceOptions"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.monitoring", "monitoring"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.carrier_ip", "carrierIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.customer_owned_ip", "customerOwnedIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.ip_owner_key", "ipOwnerKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.public_dns_name", "publicDnsName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.public_ip", "publicIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.attachment_key", "attachmentKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.device_index", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.is_delete_on_termination", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.network_card_index", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.status", "status"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.time_attach", testDateTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.interface_type", "interfaceType"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.ipv4prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.ipv6addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.ipv6prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.is_source_dest_check", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.mac_address", "macAddress"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.network_interface_key", "networkInterfaceKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.owner_key", "ownerKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.carrier_ip", "carrierIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.customer_owned_ip", "customerOwnedIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.ip_owner_key", "ipOwnerKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.public_dns_name", "publicDnsName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.public_ip", "publicIp"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.is_primary", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.private_dns_name", "privateDnsName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.private_ip_address", "privateIpAddress"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.security_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.security_groups.0.group_key", "groupKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.security_groups.0.group_name", "groupName1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.status", "status"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.subnet_key", "subnetKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.affinity", "affinity"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.availability_zone", "availabilityZone"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.group_name", "groupName1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.host_key", "hostKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.host_resource_group_arn", "hostResourceGroupArn"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.partition_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.spread_domain", "spreadDomain"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.tenancy", "tenancy"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.private_dns_name", "privateDnsName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.private_ip_address", "privateIpAddress"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.root_device_name", "rootDeviceName"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.root_device_type", "rootDeviceType"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.security_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.security_groups.0.group_key", "groupKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.security_groups.0.group_name", "groupName1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.sriov_net_support", "sriovNetSupport"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.state.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.state.0.code", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.state.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.subnet_key", "subnetKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.time_launch", testDateTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tpm_support", "tpmSupport"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.virtualization_type", "virtualizationType"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.vpc_key", "vpcKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2cost.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2cost.0.amount", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2cost.0.currency_code", "currencyCode"),
				resource.TestCheckResourceAttr(resourceName, "compute.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.connected_networks", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cores_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cpu_model", "cpuModel"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.boot_order", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.is_cbt_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.location", "location"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.persistent_mode", "persistentMode"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.size_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid", "uuid"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid_lun", "uuidLun"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.dns_name", "dnsName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.firmware", "firmware"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.cores_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.manufacturer", "manufacturer"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.memory_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.guest_state", "guestState"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.hardware_version", "hardwareVersion"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.host_name", "hostName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_pmem_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_tpm_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.latency_sensitivity", "latencySensitivity"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.memory_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address", "macAddress"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address_type", "macAddressType"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.network_name", "networkName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.switch_name", "switchName"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.bus_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.controller_key", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.unit_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system", "operatingSystem"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system_version", "operatingSystemVersion"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.pmem_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.power_state", "powerState"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.primary_ip", "primaryIp"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.label", "label"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.shared_bus", "sharedBus"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.unit_number", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.storage_provisioned_in_mbs", "10"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.threads_per_core_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_host", "hypervisorHost"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_vendor", "hypervisorVendor"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_version", "hypervisorVersion"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// AWS EC2 verify updates to updatable parameters
		{
			Config: config + variableStr + CloudBridgeAwsAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
					acctest.Optional, acctest.Update, CloudBridgeAwsEc2AssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "AWS_EC2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "attached_ebs_volumes_cost.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "attached_ebs_volumes_cost.0.amount", "2"),
				resource.TestCheckResourceAttr(resourceName, "attached_ebs_volumes_cost.0.currency_code", "currencyCode2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.architecture", "architecture2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.are_elastic_inference_accelerators_present", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.boot_mode", "bootMode2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.capacity_reservation_key", "capacityReservationKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.image_key", "imageKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_key", "instanceKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_lifecycle", "instanceLifecycle2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.instance_type", "instanceType2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.ip_address", "ipAddress2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.ipv6address", "ipv6Address2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_enclave_options", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_hibernation_options", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_source_dest_check", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.is_spot_instance", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.kernel_key", "kernelKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.licenses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.maintenance_options", "maintenanceOptions2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.monitoring", "monitoring2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.carrier_ip", "carrierIp2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.customer_owned_ip", "customerOwnedIp2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.ip_owner_key", "ipOwnerKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.public_dns_name", "publicDnsName2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.association.0.public_ip", "publicIp2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.attachment_key", "attachmentKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.device_index", "11"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.is_delete_on_termination", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.network_card_index", "11"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.status", "status2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.attachment.0.time_attach", testDateTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.interface_type", "interfaceType2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.ipv4prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.ipv6addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.ipv6prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.is_source_dest_check", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.mac_address", "macAddress2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.network_interface_key", "networkInterfaceKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.owner_key", "ownerKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.carrier_ip", "carrierIp2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.customer_owned_ip", "customerOwnedIp2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.ip_owner_key", "ipOwnerKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.public_dns_name", "publicDnsName2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.public_ip", "publicIp2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.is_primary", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.private_dns_name", "privateDnsName2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.private_ip_address", "privateIpAddress2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.security_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.security_groups.0.group_key", "groupKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.security_groups.0.group_name", "groupName2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.status", "status2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.network_interfaces.0.subnet_key", "subnetKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.affinity", "affinity2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.availability_zone", "availabilityZone2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.group_name", "groupName2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.host_key", "hostKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.host_resource_group_arn", "hostResourceGroupArn2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.partition_number", "11"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.spread_domain", "spreadDomain2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.placement.0.tenancy", "tenancy2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.private_dns_name", "privateDnsName2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.private_ip_address", "privateIpAddress2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.root_device_name", "rootDeviceName2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.root_device_type", "rootDeviceType2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.security_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.security_groups.0.group_key", "groupKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.security_groups.0.group_name", "groupName2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.sriov_net_support", "sriovNetSupport2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.state.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.state.0.code", "11"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.state.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.subnet_key", "subnetKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.time_launch", testDateTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.tpm_support", "tpmSupport2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.virtualization_type", "virtualizationType2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2.0.vpc_key", "vpcKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2cost.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2cost.0.amount", "2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ec2cost.0.currency_code", "currencyCode2"),
				resource.TestCheckResourceAttr(resourceName, "compute.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.connected_networks", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cores_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.cpu_model", "cpuModel2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.boot_order", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.is_cbt_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.location", "location2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.persistent_mode", "persistentMode2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.size_in_mbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid", "uuid2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks.0.uuid_lun", "uuidLun2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.disks_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.dns_name", "dnsName2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.firmware", "firmware2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.cores_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.manufacturer", "manufacturer2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.memory_in_mbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.gpu_devices_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.guest_state", "guestState2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.hardware_version", "hardwareVersion2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.host_name", "hostName2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_pmem_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.is_tpm_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.latency_sensitivity", "latencySensitivity2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.memory_in_mbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.label", "label2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address", "macAddress2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.mac_address_type", "macAddressType2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.network_name", "networkName2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics.0.switch_name", "switchName2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nics_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.bus_number", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimm_controller.0.label", "label2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.controller_key", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.label", "label2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.nvdimms.0.unit_number", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system", "operatingSystem2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.operating_system_version", "operatingSystemVersion2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.pmem_in_mbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.power_state", "powerState2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.primary_ip", "primaryIp2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.label", "label2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.shared_bus", "sharedBus2"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.scsi_controller.0.unit_number", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.storage_provisioned_in_mbs", "11"),
				resource.TestCheckResourceAttr(resourceName, "compute.0.threads_per_core_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "vm.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_host", "hypervisorHost2"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_vendor", "hypervisorVendor2"),
				resource.TestCheckResourceAttr(resourceName, "vm.0.hypervisor_version", "hypervisorVersion2"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_assets", "test_assets",
					acctest.Optional, acctest.Update, CloudBridgeAwsEc2AssetDataSourceRepresentation) +
				variableStr + CloudBridgeAwsAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
					acctest.Optional, acctest.Update, CloudBridgeAwsEc2AssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "asset_id"),
				resource.TestCheckResourceAttr(datasourceName, "asset_type", "AWS_EC2"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(datasourceName, "inventory_id"),
				resource.TestCheckResourceAttr(datasourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "asset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "asset_collection.0.items.#", "1"),
			),
		},
		// AWS EC2 verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
					acctest.Required, acctest.Create, CloudBridgeCloudBridgeAssetSingularDataSourceRepresentation) +
				variableStr + CloudBridgeAwsEc2AssetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "asset_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "asset_type", "AWS_EC2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "attached_ebs_volumes_cost.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attached_ebs_volumes_cost.0.amount", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "attached_ebs_volumes_cost.0.currency_code", "currencyCode2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.architecture", "architecture2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.are_elastic_inference_accelerators_present", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.boot_mode", "bootMode2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.capacity_reservation_key", "capacityReservationKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.image_key", "imageKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.instance_key", "instanceKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.instance_lifecycle", "instanceLifecycle2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.instance_type", "instanceType2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.ip_address", "ipAddress2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.ipv6address", "ipv6Address2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.is_enclave_options", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.is_hibernation_options", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.is_source_dest_check", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.is_spot_instance", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.kernel_key", "kernelKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.licenses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.maintenance_options", "maintenanceOptions2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.monitoring", "monitoring2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.association.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.association.0.carrier_ip", "carrierIp2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.association.0.customer_owned_ip", "customerOwnedIp2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.association.0.ip_owner_key", "ipOwnerKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.association.0.public_dns_name", "publicDnsName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.association.0.public_ip", "publicIp2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.attachment.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.attachment.0.attachment_key", "attachmentKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.attachment.0.device_index", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.attachment.0.is_delete_on_termination", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.attachment.0.network_card_index", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.attachment.0.status", "status2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.attachment.0.time_attach", testDateTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.interface_type", "interfaceType2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.ipv4prefixes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.ipv6addresses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.ipv6prefixes.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.is_source_dest_check", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.mac_address", "macAddress2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.network_interface_key", "networkInterfaceKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.owner_key", "ownerKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.carrier_ip", "carrierIp2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.customer_owned_ip", "customerOwnedIp2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.ip_owner_key", "ipOwnerKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.public_dns_name", "publicDnsName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.association.0.public_ip", "publicIp2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.is_primary", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.private_dns_name", "privateDnsName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.private_ip_addresses.0.private_ip_address", "privateIpAddress2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.security_groups.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.security_groups.0.group_key", "groupKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.security_groups.0.group_name", "groupName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.status", "status2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.network_interfaces.0.subnet_key", "subnetKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.placement.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.placement.0.affinity", "affinity2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.placement.0.availability_zone", "availabilityZone2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.placement.0.group_name", "groupName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.placement.0.host_key", "hostKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.placement.0.host_resource_group_arn", "hostResourceGroupArn2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.placement.0.partition_number", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.placement.0.spread_domain", "spreadDomain2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.placement.0.tenancy", "tenancy2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.private_dns_name", "privateDnsName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.private_ip_address", "privateIpAddress2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.root_device_name", "rootDeviceName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.root_device_type", "rootDeviceType2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.security_groups.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.security_groups.0.group_key", "groupKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.security_groups.0.group_name", "groupName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.sriov_net_support", "sriovNetSupport2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.state.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.state.0.code", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.state.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.subnet_key", "subnetKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.tags.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.tags.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.time_launch", testDateTime.Format(time.RFC3339Nano)),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.tpm_support", "tpmSupport2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.virtualization_type", "virtualizationType2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2.0.vpc_key", "vpcKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2cost.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2cost.0.amount", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ec2cost.0.currency_code", "currencyCode2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.connected_networks", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.cores_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.cpu_model", "cpuModel2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.boot_order", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.is_cbt_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.location", "location2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.persistent_mode", "persistentMode2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.size_in_mbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.uuid", "uuid2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks.0.uuid_lun", "uuidLun2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.disks_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.dns_name", "dnsName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.firmware", "firmware2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.0.cores_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.0.description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.0.manufacturer", "manufacturer2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.0.memory_in_mbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.gpu_devices_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.guest_state", "guestState2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.hardware_version", "hardwareVersion2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.host_name", "hostName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.is_pmem_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.is_tpm_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.latency_sensitivity", "latencySensitivity2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.memory_in_mbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.ip_addresses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.label", "label2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.mac_address", "macAddress2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.mac_address_type", "macAddressType2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.network_name", "networkName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics.0.switch_name", "switchName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nics_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimm_controller.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimm_controller.0.bus_number", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimm_controller.0.label", "label2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimms.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimms.0.controller_key", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimms.0.label", "label2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.nvdimms.0.unit_number", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.operating_system", "operatingSystem2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.operating_system_version", "operatingSystemVersion2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.pmem_in_mbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.power_state", "powerState2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.primary_ip", "primaryIp2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.scsi_controller.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.scsi_controller.0.label", "label2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.scsi_controller.0.shared_bus", "sharedBus2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.scsi_controller.0.unit_number", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.storage_provisioned_in_mbs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute.0.threads_per_core_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "inventory_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm.0.hypervisor_host", "hypervisorHost2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm.0.hypervisor_vendor", "hypervisorVendor2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm.0.hypervisor_version", "hypervisorVersion2"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudBridgeAwsEc2AssetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func TestCloudBridgeAwsEbsAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeAwsEbsAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	vaultSecretId := utils.GetEnvSettingWithBlankDefault("vaultSecretId")
	vaultSecretIdVariableStr := fmt.Sprintf("variable \"vaultSecretId\" { default = \"%s\" }\n", vaultSecretId)

	awsAccountKey := utils.GetEnvSettingWithBlankDefault("awsAccountKey")
	awsAccountKeyVariableStr := fmt.Sprintf("variable \"awsAccountKey\" { default = \"%s\" }\n", awsAccountKey)

	inventoryId := utils.GetEnvSettingWithBlankDefault("inventoryId")
	inventoryIdVariableStr := fmt.Sprintf("variable \"inventoryId\" { default = \"%s\" }\n", inventoryId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	variableStr := compartmentIdVariableStr + vaultSecretIdVariableStr + awsAccountKeyVariableStr + inventoryIdVariableStr

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_bridge_asset.test_asset"
	datasourceName := "data.oci_cloud_bridge_assets.test_assets"
	singularDatasourceName := "data.oci_cloud_bridge_asset.test_asset"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+CloudBridgeAwsAssetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
			acctest.Optional, acctest.Create, CloudBridgeAwsEbsAssetRepresentation), "cloudbridge", "asset", t)

	acctest.ResourceTest(t, testAccCheckCloudBridgeAssetDestroy, []resource.TestStep{
		// AWS EBS verify Create
		{
			Config: config + variableStr + CloudBridgeAwsAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset",
					"test_asset", acctest.Required, acctest.Create, CloudBridgeAwsEbsAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_type", "AWS_EBS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.is_encrypted", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.is_multi_attach_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.size_in_gi_bs", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.volume_key", "volumeKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.volume_type", "volumeType"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// AWS EBS delete before next Create
		{
			Config: config + variableStr + CloudBridgeAwsAssetResourceDependencies,
		},
		// AWS EBS verify Create with optionals
		{
			Config: config + variableStr + CloudBridgeAwsAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
					acctest.Optional, acctest.Create, CloudBridgeAwsEbsAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "AWS_EBS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.device", "device"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.instance_key", "instanceKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.is_delete_on_termination", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.status", "status"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.volume_key", "volumeKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.availability_zone", "availabilityZone"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.iops", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.is_encrypted", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.is_multi_attach_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.size_in_gi_bs", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.status", "status"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.throughput", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.volume_key", "volumeKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.volume_type", "volumeType"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// AWS EBS verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + variableStr + compartmentIdUVariableStr + CloudBridgeAwsAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudBridgeAwsEbsAssetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "AWS_EBS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.device", "device"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.instance_key", "instanceKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.is_delete_on_termination", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.status", "status"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.volume_key", "volumeKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.availability_zone", "availabilityZone"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.iops", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.is_encrypted", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.is_multi_attach_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.size_in_gi_bs", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.status", "status"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.tags.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.tags.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.throughput", "10"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.volume_key", "volumeKey"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.volume_type", "volumeType"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// AWS EBS verify updates to updatable parameters
		{
			Config: config + variableStr + CloudBridgeAwsAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
					acctest.Optional, acctest.Update, CloudBridgeAwsEbsAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "AWS_EBS"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.device", "device2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.instance_key", "instanceKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.is_delete_on_termination", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.status", "status2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.attachments.0.volume_key", "volumeKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.availability_zone", "availabilityZone2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.iops", "11"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.is_encrypted", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.is_multi_attach_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.size_in_gi_bs", "11"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.status", "status2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.tags.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.tags.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.tags.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.throughput", "11"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.volume_key", "volumeKey2"),
				resource.TestCheckResourceAttr(resourceName, "aws_ebs.0.volume_type", "volumeType2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// AWS EBS verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_assets", "test_assets",
					acctest.Optional, acctest.Update, CloudBridgeAwsEbsAssetDataSourceRepresentation) +
				variableStr + CloudBridgeAwsAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
					acctest.Optional, acctest.Update, CloudBridgeAwsEbsAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "asset_id"),
				resource.TestCheckResourceAttr(datasourceName, "asset_type", "AWS_EBS"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(datasourceName, "inventory_id"),
				resource.TestCheckResourceAttr(datasourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "asset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "asset_collection.0.items.#", "1"),
			),
		},
		// AWS EBS verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
					acctest.Required, acctest.Create, CloudBridgeCloudBridgeAssetSingularDataSourceRepresentation) +
				variableStr + CloudBridgeAwsEbsAssetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "asset_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "asset_type", "AWS_EBS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.attachments.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.attachments.0.device", "device2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.attachments.0.instance_key", "instanceKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.attachments.0.is_delete_on_termination", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.attachments.0.status", "status2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.attachments.0.volume_key", "volumeKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.availability_zone", "availabilityZone2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.iops", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.is_encrypted", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.is_multi_attach_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.size_in_gi_bs", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.status", "status2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.tags.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.tags.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.tags.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.throughput", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.volume_key", "volumeKey2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "aws_ebs.0.volume_type", "volumeType2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "external_asset_key", externalAssetKey),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "inventory_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_key", "sourceKey"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// AWS EBS verify resource import
		{
			Config:                  config + CloudBridgeAwsEbsAssetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeInventoryAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeInventoryAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	vaultSecretId := utils.GetEnvSettingWithBlankDefault("vaultSecretId")
	vaultSecretIdVariableStr := fmt.Sprintf("variable \"vaultSecretId\" { default = \"%s\" }\n", vaultSecretId)

	inventoryId := utils.GetEnvSettingWithBlankDefault("inventoryId")
	inventoryIdVariableStr := fmt.Sprintf("variable \"inventoryId\" { default = \"%s\" }\n", inventoryId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	variableStr := compartmentIdVariableStr + inventoryIdVariableStr + vaultSecretIdVariableStr

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_bridge_asset.test_asset"
	datasourceName := "data.oci_cloud_bridge_assets.test_assets"
	singularDatasourceName := "data.oci_cloud_bridge_asset.test_asset"

	var resId, resId2 string
	acctest.SaveConfigContent(config+variableStr+CloudBridgeOlvmAssetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Create, CloudBridgeInventoryAssetRepresentation), "cloudbridge", "asset", t)

	acctest.ResourceTest(t, testAccCheckCloudBridgeAssetDestroy, []resource.TestStep{
		{
			Config: config + variableStr + CloudBridgeOlvmAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Required, acctest.Create, CloudBridgeInventoryAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_type", "INVENTORY_ASSET"),
				resource.TestCheckResourceAttr(resourceName, "asset_class_name", inventoryAssetClassName),
				resource.TestCheckResourceAttr(resourceName, "asset_class_version", inventoryAssetClassVersion),
				acctest.TestCheckJsonResourceAttr(resourceName, "asset_details", inventoryAssetDetailsJson),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", inventoryAssetExternalKey),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", inventoryAssetSourceKey),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		{
			Config: config + variableStr + CloudBridgeOlvmAssetResourceDependencies,
		},
		{
			Config: config + variableStr + CloudBridgeOlvmAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Create, CloudBridgeInventoryAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "INVENTORY_ASSET"),
				resource.TestCheckResourceAttr(resourceName, "asset_class_name", inventoryAssetClassName),
				resource.TestCheckResourceAttr(resourceName, "asset_class_version", inventoryAssetClassVersion),
				acctest.TestCheckJsonResourceAttr(resourceName, "asset_details", inventoryAssetDetailsJson),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", inventoryAssetExternalKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", inventoryAssetSourceKey),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		{
			Config: config + variableStr + compartmentIdUVariableStr + CloudBridgeOlvmAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudBridgeInventoryAssetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "INVENTORY_ASSET"),
				resource.TestCheckResourceAttr(resourceName, "asset_class_name", inventoryAssetClassName),
				resource.TestCheckResourceAttr(resourceName, "asset_class_version", inventoryAssetClassVersion),
				acctest.TestCheckJsonResourceAttr(resourceName, "asset_details", inventoryAssetDetailsJson),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", inventoryAssetExternalKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", inventoryAssetSourceKey),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		{
			Config: config + variableStr + CloudBridgeOlvmAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Update, CloudBridgeInventoryAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "asset_type", "INVENTORY_ASSET"),
				resource.TestCheckResourceAttr(resourceName, "asset_class_name", inventoryAssetClassName),
				resource.TestCheckResourceAttr(resourceName, "asset_class_version", inventoryAssetClassVersion),
				acctest.TestCheckJsonResourceAttr(resourceName, "asset_details", inventoryAssetDetailsJson),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "external_asset_key", inventoryAssetExternalKey),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_id"),
				resource.TestCheckResourceAttr(resourceName, "source_key", inventoryAssetSourceKey),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_assets", "test_assets",
					acctest.Optional, acctest.Update, CloudBridgeInventoryAssetDataSourceRepresentation) +
				variableStr + CloudBridgeOlvmAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Update, CloudBridgeInventoryAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "asset_id"),
				resource.TestCheckResourceAttr(datasourceName, "asset_type", "INVENTORY_ASSET"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "display_name"),
				resource.TestCheckResourceAttr(datasourceName, "external_asset_key", inventoryAssetExternalKey),
				resource.TestCheckResourceAttrSet(datasourceName, "inventory_id"),
				resource.TestCheckResourceAttr(datasourceName, "source_key", inventoryAssetSourceKey),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "asset_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "asset_collection.0.items.#", "1"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset",
					acctest.Required, acctest.Create, CloudBridgeCloudBridgeAssetSingularDataSourceRepresentation) +
				variableStr + CloudBridgeOlvmAssetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Update, CloudBridgeInventoryAssetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "asset_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "asset_source_ids.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "asset_type", "INVENTORY_ASSET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "asset_class_name", inventoryAssetClassName),
				resource.TestCheckResourceAttr(singularDatasourceName, "asset_class_version", inventoryAssetClassVersion),
				acctest.TestCheckJsonResourceAttr(singularDatasourceName, "asset_details", inventoryAssetDetailsJson),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "external_asset_key", inventoryAssetExternalKey),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "inventory_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_key", inventoryAssetSourceKey),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		{
			Config:                  config + CloudBridgeInventoryAssetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"asset_details"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudBridgeAssetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).InventoryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_bridge_asset" {
			noResourceFound = false
			request := oci_cloud_bridge.GetAssetRequest{}

			tmp := rs.Primary.ID
			request.AssetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")

			response, err := client.GetAsset(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_bridge.AssetLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("CloudBridgeAsset") {
		resource.AddTestSweepers("CloudBridgeAsset", &resource.Sweeper{
			Name:         "CloudBridgeAsset",
			Dependencies: acctest.DependencyGraph["asset"],
			F:            sweepCloudBridgeAssetResource,
		})
	}
}

func sweepCloudBridgeAssetResource(compartment string) error {
	inventoryClient := acctest.GetTestClients(&schema.ResourceData{}).InventoryClient()
	assetIds, err := getCloudBridgeAssetIds(compartment)
	if err != nil {
		return err
	}
	for _, assetId := range assetIds {
		if ok := acctest.SweeperDefaultResourceId[assetId]; !ok {
			deleteAssetRequest := oci_cloud_bridge.DeleteAssetRequest{}

			deleteAssetRequest.AssetId = &assetId

			deleteAssetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")
			_, err := inventoryClient.DeleteAsset(context.Background(), deleteAssetRequest)
			if err != nil {
				fmt.Printf("Error deleting Asset %s %s, It is possible that the resource is already deleted. Please verify manually \n", assetId, err)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &assetId, CloudBridgeAssetSweepWaitCondition, time.Duration(3*time.Minute),
				CloudBridgeAssetSweepResponseFetchOperation, "cloud_bridge", true)
		}
	}
	return nil
}

func getCloudBridgeAssetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AssetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	inventoryClient := acctest.GetTestClients(&schema.ResourceData{}).InventoryClient()

	listAssetsRequest := oci_cloud_bridge.ListAssetsRequest{}
	listAssetsRequest.CompartmentId = &compartmentId
	listAssetsRequest.LifecycleState = oci_cloud_bridge.AssetLifecycleStateActive
	listAssetsResponse, err := inventoryClient.ListAssets(context.Background(), listAssetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Asset list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, asset := range listAssetsResponse.Items {
		id := *asset.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AssetId", id)
	}
	return resourceIds, nil
}

func CloudBridgeAssetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if assetResponse, ok := response.Response.(oci_cloud_bridge.GetAssetResponse); ok {
		return assetResponse.GetLifecycleState() != oci_cloud_bridge.AssetLifecycleStateDeleted
	}
	return false
}

func CloudBridgeAssetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.InventoryClient().GetAsset(context.Background(), oci_cloud_bridge.GetAssetRequest{
		AssetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
