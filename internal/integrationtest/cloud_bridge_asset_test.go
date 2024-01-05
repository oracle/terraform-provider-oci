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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	assetSourceId    = `${var.assetSourceId}`

	CloudBridgeAssetRequiredOnlyResource = CloudBridgeAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Required, acctest.Create, CloudBridgeAssetRepresentation)

	CloudBridgeAssetResourceConfig = CloudBridgeAssetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_asset", "test_asset", acctest.Optional, acctest.Update, CloudBridgeAssetRepresentation)

	CloudBridgeCloudBridgeAssetSingularDataSourceRepresentation = map[string]interface{}{
		"asset_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_bridge_asset.test_asset.id}`},
	}

	CloudBridgeCloudBridgeAssetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"asset_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_bridge_asset.test_asset.id}`},
		"asset_type":         acctest.Representation{RepType: acctest.Optional, Create: `VMWARE_VM`},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"external_asset_key": acctest.Representation{RepType: acctest.Optional, Create: externalAssetKey},
		"inventory_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_bridge_inventory.test_inventory.id}`},
		"source_key":         acctest.Representation{RepType: acctest.Optional, Create: `sourceKey`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAssetDataSourceFilterRepresentation}}
	CloudBridgeAssetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_bridge_asset.test_asset.id}`}},
	}

	CloudBridgeAssetRepresentation = map[string]interface{}{
		"asset_type":         acctest.Representation{RepType: acctest.Required, Create: `VMWARE_VM`},
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"external_asset_key": acctest.Representation{RepType: acctest.Required, Create: externalAssetKey},
		"inventory_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_bridge_inventory.test_inventory.id}`},
		"source_key":         acctest.Representation{RepType: acctest.Required, Create: `sourceKey`},
		"asset_source_ids":   acctest.Representation{RepType: acctest.Optional, Create: []string{assetSourceId}, Update: []string{assetSourceId}},
		"compute":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetComputeRepresentation},
		"display_name":       acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"vmware_vm":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetVmwareVmRepresentation},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}
	CloudBridgeAssetComputeRepresentation = map[string]interface{}{
		"connected_networks":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"cores_count":                acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
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
		"memory_in_mbs":              acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
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
		"hypervisor_host":    acctest.Representation{RepType: acctest.Optional, Create: `hypervisorHost`, Update: `hypervisorHost2`},
		"hypervisor_vendor":  acctest.Representation{RepType: acctest.Optional, Create: `hypervisorVendor`, Update: `hypervisorVendor2`},
		"hypervisor_version": acctest.Representation{RepType: acctest.Optional, Create: `hypervisorVersion`, Update: `hypervisorVersion2`},
	}
	CloudBridgeAssetVmwareVCenterRepresentation = map[string]interface{}{
		"data_center":     acctest.Representation{RepType: acctest.Optional, Create: `dataCenter`, Update: `dataCenter2`},
		"vcenter_key":     acctest.Representation{RepType: acctest.Optional, Create: `vcenterKey`, Update: `vcenterKey2`},
		"vcenter_version": acctest.Representation{RepType: acctest.Optional, Create: `vcenterVersion`, Update: `vcenterVersion2`},
	}
	CloudBridgeAssetVmwareVmRepresentation = map[string]interface{}{
		"cluster":                           acctest.Representation{RepType: acctest.Optional, Create: `cluster`, Update: `cluster2`},
		"customer_fields":                   acctest.Representation{RepType: acctest.Optional, Create: []string{`customerFields`}, Update: []string{`customerFields2`}},
		"customer_tags":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: CloudBridgeAssetVmwareVmCustomerTagsRepresentation},
		"fault_tolerance_bandwidth":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"fault_tolerance_secondary_latency": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"fault_tolerance_state":             acctest.Representation{RepType: acctest.Optional, Create: `faultToleranceState`, Update: `faultToleranceState2`},
		"instance_uuid":                     acctest.Representation{RepType: acctest.Optional, Create: `instanceUuid`, Update: `instanceUuid2`},
		"is_disks_cbt_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_disks_uuid_enabled":             acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"path":                              acctest.Representation{RepType: acctest.Optional, Create: `path`, Update: `path2`},
		"vmware_tools_status":               acctest.Representation{RepType: acctest.Optional, Create: `vmwareToolsStatus`, Update: `vmwareToolsStatus2`},
	}
	CloudBridgeAssetComputeDisksRepresentation = map[string]interface{}{
		"boot_order":      acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
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

	CloudBridgeAssetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Required, acctest.Create, CloudBridgeInventoryRepresentation)
)

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeAssetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeAssetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	assetSourceId := utils.GetEnvSettingWithBlankDefault("assetSourceId")
	assetSourceIdVariableStr := fmt.Sprintf("variable \"assetSourceId\" { default = \"%s\" }\n", assetSourceId)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	variableStr := compartmentIdVariableStr + assetSourceIdVariableStr

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
