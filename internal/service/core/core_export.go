package core

import (
	"context"
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportCoreInstancePoolInstanceHints.GetIdFn = getCoreInstancePoolInstanceId
	exportCoreNetworkSecurityGroupSecurityRuleHints.GetIdFn = getCoreNetworkSecurityGroupSecurityRuleId
	exportCoreDrgRouteTableRouteRuleHints.GetIdFn = getCoreDrgRouteTableRouteRuleId
	exportCoreBootVolumeHints.ProcessDiscoveredResourcesFn = filterSourcedBootVolumes
	exportCoreCrossConnectGroupHints.DiscoverableLifecycleStates = append(exportCoreCrossConnectGroupHints.DiscoverableLifecycleStates, string(oci_core.CrossConnectGroupLifecycleStateInactive))
	exportCoreDhcpOptionsHints.ProcessDiscoveredResourcesFn = processDefaultDhcpOptions
	exportCoreImageHints.ProcessDiscoveredResourcesFn = filterCustomImages

	exportCoreInstanceHints.DiscoverableLifecycleStates = append(exportCoreInstanceHints.DiscoverableLifecycleStates, string(oci_core.InstanceLifecycleStateStopped))
	exportCoreInstanceHints.ProcessDiscoveredResourcesFn = processInstances
	exportCorePublicIpHints.ProcessDiscoveredResourcesFn = processCorePublicIp
	exportCorePrivateIpHints.ProcessDiscoveredResourcesFn = processPrivateIps
	exportCoreInstanceHints.RequireResourceRefresh = true
	exportCoreNetworkSecurityGroupSecurityRuleHints.DatasourceClass = "oci_core_network_security_group_security_rules"
	exportCoreNetworkSecurityGroupSecurityRuleHints.DatasourceItemsAttr = "security_rules"
	exportCoreNetworkSecurityGroupSecurityRuleHints.ProcessDiscoveredResourcesFn = processNetworkSecurityGroupRules
	exportCoreRouteTableHints.ProcessDiscoveredResourcesFn = processDefaultRouteTables
	exportCoreSecurityListHints.ProcessDiscoveredResourcesFn = processDefaultSecurityLists
	exportCoreVcnHints.ProcessDiscoveredResourcesFn = processCoreVcns
	exportCoreVnicAttachmentHints.RequireResourceRefresh = true
	exportCoreVnicAttachmentHints.ProcessDiscoveredResourcesFn = filterSecondaryVnicAttachments
	exportCoreVolumeGroupHints.ProcessDiscoveredResourcesFn = processVolumeGroups

	exportCoreDrgRouteTableRouteRuleHints.DatasourceClass = "oci_core_drg_route_table_route_rules"
	exportCoreDrgRouteTableRouteRuleHints.DatasourceItemsAttr = "drg_route_rules"
	exportCoreDrgRouteTableRouteRuleHints.ProcessDiscoveredResourcesFn = processDrgRouteTableRouteRules
	exportCoreDrgRouteDistributionHints.ProcessDiscoveredResourcesFn = processDrgRouteDistributions
	tf_export.RegisterCompartmentGraphs("core", coreResourceGraph)
	tf_export.RegisterRelatedResourcesGraph("oci_core_instance", relatedcoreinstance)
	tf_export.RegisterRelatedResourcesGraph("oci_core_volume_attachment", relatedcorevolumeattachment)
	tf_export.BuildAvailabilityResourceGraph("oci_identity_availability_domain", customAssociationCoreIdentityAvailabilityDomain)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func processDrgRouteTableRouteRules(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}
		resource.SourceAttributes["drg_route_table_id"] = resource.Parent.Id
	}
	return resources, nil
}

func processDrgRouteDistributions(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {

	// filtering out export drg route distributions, if the drg route distributions is of type "export", we don't consider it.
	results := []*tf_export.OCIResource{}
	for _, resource := range resources {
		if resource == nil {
			continue
		}

		if resource.SourceAttributes["distribution_type"] != nil && resource.SourceAttributes["distribution_type"].(string) == "EXPORT" {
			continue // skip over export drg route distributions
		}
		results = append(results, resource)
	}
	return results, nil
}

func processVolumeGroups(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	// Replace the volume Group's source details volume list with the actual volume list
	// The source details only captures the list of volumes that were known when the Group was created.
	// Additional volumes may have been added since and should be part of the source_details that we generate.
	// TODO: This is a shortcoming that should be addressed by the service and/or the Terraform
	for _, group := range resources {
		volumeIdsRaw, exists := group.SourceAttributes["volume_ids"]
		if !exists {
			continue
		}

		if volumeIds, ok := volumeIdsRaw.([]interface{}); ok && len(volumeIds) > 0 {
			sourceDetailsRaw, detailsExist := group.SourceAttributes["source_details"]
			if !detailsExist {
				continue
			}

			sourceDetails := sourceDetailsRaw.([]interface{})[0].(map[string]interface{})
			sourceDetails["volume_ids"] = volumeIds
		}
	}

	return resources, nil
}

func filterSecondaryVnicAttachments(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}

	for _, attachment := range resources {
		// Filter out any primary vnics, as it's not necessary to Create separate TF resources for those.
		datasourceSchema := tf_export.DatasourcesMap["oci_core_vnic"]
		if vnicReadFn := datasourceSchema.Read; vnicReadFn != nil {
			d := datasourceSchema.TestResourceData()
			d.Set("vnic_id", attachment.SourceAttributes["vnic_id"].(string))
			if err := vnicReadFn(d, ctx.Clients); err != nil {
				return results, err
			}

			if isPrimaryVnic, ok := d.GetOkExists("is_primary"); ok && isPrimaryVnic.(bool) {
				continue
			}
		}
		results = append(results, attachment)
	}

	return results, nil
}

func processCoreVcns(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	// remove deprecated cidr_block field from discovered vcns,
	// either cidr_block or cidr_blocks should be specified in config
	// service returns the cidr_block value in cidr_blocks field
	for _, resource := range resources {
		// Adding default DHCP Options reference for VCN to referenceMap to be utilized during processDefaultDhcpOptions()
		if defaultDhcpOptionsId, exists := resource.SourceAttributes["default_dhcp_options_id"]; exists {
			if defaultDhcpOptionsIdStr, ok := defaultDhcpOptionsId.(string); ok {
				tf_export.RefMapLock.Lock()
				tf_export.ReferenceMap[defaultDhcpOptionsIdStr] = tf_export.TfHclVersionvar.GetDoubleExpHclString(resource.GetTerraformReference(), "default_dhcp_options_id")
				tf_export.RefMapLock.Unlock()
			}
		}

		// Adding default Route Table reference for VCN to referenceMap to be utilized during processDefaultRouteTables()
		if defaultRouteTableId, exists := resource.SourceAttributes["default_route_table_id"]; exists {
			if defaultRouteTableIdStr, ok := defaultRouteTableId.(string); ok {
				tf_export.RefMapLock.Lock()
				tf_export.ReferenceMap[defaultRouteTableIdStr] = tf_export.TfHclVersionvar.GetDoubleExpHclString(resource.GetTerraformReference(), "default_route_table_id")
				tf_export.RefMapLock.Unlock()
			}
		}

		// Adding default SecurityList reference for VCN to referenceMap to be utilized during processDefaultSecurityLists()
		if defaultSecurityListId, exists := resource.SourceAttributes["default_security_list_id"]; exists {
			if defaultSecurityListIdStr, ok := defaultSecurityListId.(string); ok {
				tf_export.RefMapLock.Lock()
				tf_export.ReferenceMap[defaultSecurityListIdStr] = tf_export.TfHclVersionvar.GetDoubleExpHclString(resource.GetTerraformReference(), "default_security_list_id")
				tf_export.RefMapLock.Unlock()
			}
		}

		if _, ok := resource.SourceAttributes["cidr_block"].(string); ok {
			delete(resource.SourceAttributes, "cidr_block")
		}
	}
	return resources, nil
}

func processDefaultSecurityLists(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	// Default security lists need to be handled as default resources
	for _, resource := range resources {
		if resource.SourceAttributes["vcn_id"] != nil {
			vcnId := resource.SourceAttributes["vcn_id"].(string)
			request := oci_core.GetVcnRequest{}
			request.VcnId = &vcnId
			response, err := ctx.Clients.VirtualNetworkClient().GetVcn(context.Background(), request)

			if err != nil {
				return resources, err
			}

			if response.Vcn.DefaultSecurityListId != nil && resource.Id == *response.Vcn.DefaultSecurityListId {
				resource.SourceAttributes["manage_default_resource_id"] = resource.Id
				resource.TerraformResource.TerraformClass = "oci_core_default_security_list"

				if referenceVal, exists := tf_export.ReferenceMap[resource.Id]; exists {
					resource.TerraformResource.TerraformReferenceIdString = referenceVal
				}
			}
		}
	}
	return resources, nil
}

func processDefaultRouteTables(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	// Default route tables need to be handled as default resources
	for _, resource := range resources {
		if resource.SourceAttributes["vcn_id"] != nil {
			vcnId := resource.SourceAttributes["vcn_id"].(string)
			request := oci_core.GetVcnRequest{}
			request.VcnId = &vcnId
			response, err := ctx.Clients.VirtualNetworkClient().GetVcn(context.Background(), request)

			if err != nil {
				return resources, err
			}

			if response.Vcn.DefaultRouteTableId != nil && resource.Id == *response.Vcn.DefaultRouteTableId {
				resource.SourceAttributes["manage_default_resource_id"] = resource.Id
				resource.TerraformResource.TerraformClass = "oci_core_default_route_table"

				if referenceVal, exists := tf_export.ReferenceMap[resource.Id]; exists {
					resource.TerraformResource.TerraformReferenceIdString = referenceVal
				}
			}
		}
	}
	return resources, nil
}

func processNetworkSecurityGroupRules(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	for _, resource := range resources {
		if resource.Parent == nil {
			continue
		}

		resource.SourceAttributes["network_security_group_id"] = resource.Parent.Id
	}
	return resources, nil
}

func processPrivateIps(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	privateIps := []*tf_export.OCIResource{}

	for _, privateIp := range resources {

		if privateIp.HasFreeformTag(tf_export.ResourceCreatedByInstancePool) {
			continue
		}

		// OKE will add tagging support, for now we rely on Automatic default tags for tenancies created after December 17, 2019
		if privateIp.HasDefinedTag(tf_export.OracleTagsCreatedBy, tf_export.OkeTagValue) {
			continue
		}

		privateIps = append(privateIps, privateIp)
	}

	return privateIps, nil
}

func processCorePublicIp(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	publicIps := []*tf_export.OCIResource{}

	for _, publicIp := range resources {

		if lifeTime, exists := publicIp.SourceAttributes["lifetime"].(string); exists {
			// this is public IP created by NAT gateway
			if lifeTime == "EPHEMERAL" {
				continue
			}
		}
		publicIps = append(publicIps, publicIp)
	}

	return publicIps, nil
}

func processInstances(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}

	for _, instance := range resources {
		// Omit any resources that were launched by an instance pool. Those shouldn't be managed by Terraform as they are created
		// and managed through the instance pool resource instead.
		if instance.HasFreeformTag(tf_export.ResourceCreatedByInstancePool) {
			continue
		}

		// OKE will add tagging support, for now we rely on Automatic default tags for tenancies created after December 17, 2019
		if instance.HasDefinedTag(tf_export.OracleTagsCreatedBy, tf_export.OkeTagValue) {
			continue
		}

		// Ensure the boot volume created by this instance can be referenced elsewhere by adding it to the reference map
		if bootVolumeId, exists := instance.SourceAttributes["boot_volume_id"]; exists {
			if bootVolumeIdStr, ok := bootVolumeId.(string); ok {
				tf_export.RefMapLock.Lock()
				tf_export.ReferenceMap[bootVolumeIdStr] = tf_export.TfHclVersionvar.GetDoubleExpHclString(instance.GetTerraformReference(), "boot_volume_id")
				tf_export.RefMapLock.Unlock()
			}
		}

		if rawSourceDetailsList, sourceDetailsExist := instance.SourceAttributes["source_details"]; sourceDetailsExist {
			if sourceDetailList, ok := rawSourceDetailsList.([]interface{}); ok && len(sourceDetailList) > 0 {
				if sourceDetails, ok := sourceDetailList[0].(map[string]interface{}); ok {
					if imageId, ok := instance.SourceAttributes["image"].(string); ok {
						sourceDetails["source_id"] = imageId

						// The image OCID may be different if it's in a different tenancy or region, add a variable for users to specify
						// TODO: handle nested attribute better instead of hardcode
						imageVarName := utils.GetVarNameFromAttributeOfResources("source_details.source_id", instance.TerraformClass, instance.TerraformName)
						tf_export.Vars[imageVarName] = fmt.Sprintf("\"%s\"", imageId)
						tf_export.RefMapLock.Lock()
						tf_export.ReferenceMap[imageId] = tf_export.TfHclVersionvar.GetVarHclString(imageVarName)
						tf_export.RefMapLock.Unlock()
					}

					// Workaround for service limitation. Service returns 47GB size for boot volume but LaunchInstance can only
					// accept sizes 50GB and above. If such a situation arises, fall back to service default values for boot volume size.
					if bootVolumeSizeInGbs, exists := sourceDetails["boot_volume_size_in_gbs"]; exists {
						bootVolumeSize, err := strconv.ParseInt(bootVolumeSizeInGbs.(string), 10, 64)
						if err != nil {
							return resources, err
						}

						if bootVolumeSize < 50 {
							delete(sourceDetails, "boot_volume_size_in_gbs")
						}
					}
				}
			}
		}

		results = append(results, instance)
	}

	return results, nil
}

func filterCustomImages(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}

	// Filter out official images that are predefined by Oracle. We cannot manage such images in Terraform.
	// Official images have a null or empty compartment ID.
	for _, image := range resources {
		compartmentId, exists := image.SourceAttributes["compartment_id"]
		if !exists {
			continue
		}

		if compartmentIdString, ok := compartmentId.(string); !ok || len(compartmentIdString) == 0 {
			continue
		}

		results = append(results, image)
	}

	return results, nil
}

func processDefaultDhcpOptions(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	// Default dhcp options need to be handled as default resources
	for _, resource := range resources {
		if resource.SourceAttributes["vcn_id"] != nil {
			vcnId := resource.SourceAttributes["vcn_id"].(string)
			request := oci_core.GetVcnRequest{}
			request.VcnId = &vcnId
			response, err := ctx.Clients.VirtualNetworkClient().GetVcn(context.Background(), request)

			if err != nil {
				return resources, err
			}

			if response.Vcn.DefaultDhcpOptionsId != nil && resource.Id == *response.Vcn.DefaultDhcpOptionsId {
				resource.SourceAttributes["manage_default_resource_id"] = resource.Id
				resource.TerraformResource.TerraformClass = "oci_core_default_dhcp_options"

				if referenceVal, exists := tf_export.ReferenceMap[resource.Id]; exists {
					resource.TerraformResource.TerraformReferenceIdString = referenceVal
				}
			}
		}
	}
	return resources, nil
}
func filterSourcedBootVolumes(ctx *tf_export.ResourceDiscoveryContext, resources []*tf_export.OCIResource) ([]*tf_export.OCIResource, error) {
	results := []*tf_export.OCIResource{}

	// Filter out boot volumes that don't have source details. We cannot Create boot volumes unless they have source details.
	for _, bootVolume := range resources {
		sourceDetails, exists := bootVolume.SourceAttributes["source_details"]
		if !exists {
			continue
		}

		if sourceDetailsList, ok := sourceDetails.([]interface{}); !ok || len(sourceDetailsList) == 0 {
			continue
		}

		results = append(results, bootVolume)
	}

	return results, nil
}

func getCoreInstancePoolInstanceId(resource *tf_export.OCIResource) (string, error) {

	instancePoolId := resource.Parent.Id
	instanceId := resource.SourceAttributes["instance_id"].(string)
	return GetInstancePoolInstanceCompositeId(instancePoolId, instanceId), nil
}

func getCoreNetworkSecurityGroupSecurityRuleId(resource *tf_export.OCIResource) (string, error) {

	networkSecurityGroupId := resource.Parent.Id
	securityRuleId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find id for Core NetworkSecurityGroupSecurityRule")
	}
	return GetNetworkSecurityGroupSecurityRuleCompositeId(networkSecurityGroupId, securityRuleId), nil
}

func getCoreDrgRouteTableRouteRuleId(resource *tf_export.OCIResource) (string, error) {

	drgRouteTableId := resource.Parent.Id
	drgRouteRuleId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find drgRouteTableId for Core DrgRouteTableRouteRule")
	}
	return GetDrgRouteTableRouteRuleCompositeId(drgRouteTableId, drgRouteRuleId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportCoreBootVolumeBackupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_boot_volume_backup",
	DatasourceClass:      "oci_core_boot_volume_backups",
	DatasourceItemsAttr:  "boot_volume_backups",
	ResourceAbbreviation: "boot_volume_backup",
	DiscoverableLifecycleStates: []string{
		string(oci_core.BootVolumeBackupLifecycleStateAvailable),
	},
}

var exportCoreBootVolumeHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_boot_volume",
	DatasourceClass:      "oci_core_boot_volumes",
	DatasourceItemsAttr:  "boot_volumes",
	ResourceAbbreviation: "boot_volume",
	DiscoverableLifecycleStates: []string{
		string(oci_core.BootVolumeLifecycleStateAvailable),
	},
}

var exportCoreConsoleHistoryHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_console_history",
	DatasourceClass:      "oci_core_console_histories",
	DatasourceItemsAttr:  "console_histories",
	ResourceAbbreviation: "console_history",
	DiscoverableLifecycleStates: []string{
		string(oci_core.ConsoleHistoryLifecycleStateSucceeded),
	},
}

var exportCoreClusterNetworkHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_core_cluster_network",
	DatasourceClass:        "oci_core_cluster_networks",
	DatasourceItemsAttr:    "cluster_networks",
	ResourceAbbreviation:   "cluster_network",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_core.ClusterNetworkLifecycleStateRunning),
	},
}

var exportCoreComputeImageCapabilitySchemaHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_core_compute_image_capability_schema",
	DatasourceClass:        "oci_core_compute_image_capability_schemas",
	DatasourceItemsAttr:    "compute_image_capability_schemas",
	ResourceAbbreviation:   "compute_image_capability_schema",
	RequireResourceRefresh: true,
}

var exportCoreComputeCapacityReservationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_core_compute_capacity_reservation",
	DatasourceClass:        "oci_core_compute_capacity_reservations",
	DatasourceItemsAttr:    "compute_capacity_reservations",
	ResourceAbbreviation:   "compute_capacity_reservation",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_core.ComputeCapacityReservationLifecycleStateActive),
	},
}

var exportCoreCpeHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_cpe",
	DatasourceClass:      "oci_core_cpes",
	DatasourceItemsAttr:  "cpes",
	ResourceAbbreviation: "cpe",
}

var exportCoreCrossConnectGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_cross_connect_group",
	DatasourceClass:      "oci_core_cross_connect_groups",
	DatasourceItemsAttr:  "cross_connect_groups",
	ResourceAbbreviation: "cross_connect_group",
	DiscoverableLifecycleStates: []string{
		string(oci_core.CrossConnectGroupLifecycleStateProvisioned),
	},
}

var exportCoreCrossConnectHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_cross_connect",
	DatasourceClass:      "oci_core_cross_connects",
	DatasourceItemsAttr:  "cross_connects",
	ResourceAbbreviation: "cross_connect",
	DiscoverableLifecycleStates: []string{
		string(oci_core.CrossConnectLifecycleStatePendingCustomer),
		string(oci_core.CrossConnectLifecycleStateProvisioned),
	},
}

var exportCoreDhcpOptionsHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_dhcp_options",
	DatasourceClass:      "oci_core_dhcp_options",
	DatasourceItemsAttr:  "options",
	ResourceAbbreviation: "dhcp_options",
	DiscoverableLifecycleStates: []string{
		string(oci_core.DhcpOptionsLifecycleStateAvailable),
	},
}

var exportCoreDrgAttachmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_drg_attachment",
	DatasourceClass:      "oci_core_drg_attachments",
	DatasourceItemsAttr:  "drg_attachments",
	ResourceAbbreviation: "drg_attachment",
	DiscoverableLifecycleStates: []string{
		string(oci_core.DrgAttachmentLifecycleStateAttached),
	},
}

var exportCoreDrgHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_drg",
	DatasourceClass:      "oci_core_drgs",
	DatasourceItemsAttr:  "drgs",
	ResourceAbbreviation: "drg",
	DiscoverableLifecycleStates: []string{
		string(oci_core.DrgLifecycleStateAvailable),
	},
}

var exportCoreDedicatedVmHostHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_core_dedicated_vm_host",
	DatasourceClass:        "oci_core_dedicated_vm_hosts",
	DatasourceItemsAttr:    "dedicated_vm_hosts",
	ResourceAbbreviation:   "dedicated_vm_host",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_core.DedicatedVmHostLifecycleStateActive),
	},
}

var exportCoreImageHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_image",
	DatasourceClass:      "oci_core_images",
	DatasourceItemsAttr:  "images",
	ResourceAbbreviation: "image",
	DiscoverableLifecycleStates: []string{
		string(oci_core.ImageLifecycleStateAvailable),
	},
}

var exportCoreInstanceConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_core_instance_configuration",
	DatasourceClass:        "oci_core_instance_configurations",
	DatasourceItemsAttr:    "instance_configurations",
	ResourceAbbreviation:   "instance_configuration",
	RequireResourceRefresh: true,
}

var exportCoreInstanceConsoleConnectionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_instance_console_connection",
	DatasourceClass:      "oci_core_instance_console_connections",
	DatasourceItemsAttr:  "instance_console_connections",
	ResourceAbbreviation: "instance_console_connection",
	DiscoverableLifecycleStates: []string{
		string(oci_core.InstanceConsoleConnectionLifecycleStateActive),
	},
}

var exportCoreInstancePoolInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_instance_pool_instance",
	DatasourceClass:      "oci_core_instance_pool_instances",
	DatasourceItemsAttr:  "instances",
	ResourceAbbreviation: "instance_pool_instance",
	DiscoverableLifecycleStates: []string{
		string(oci_core.InstancePoolInstanceLifecycleStateActive),
	},
}

var exportCoreInstancePoolHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_core_instance_pool",
	DatasourceClass:        "oci_core_instance_pools",
	DatasourceItemsAttr:    "instance_pools",
	ResourceAbbreviation:   "instance_pool",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_core.InstancePoolLifecycleStateRunning),
	},
}

var exportCoreInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_instance",
	DatasourceClass:      "oci_core_instances",
	DatasourceItemsAttr:  "instances",
	ResourceAbbreviation: "instance",
	DiscoverableLifecycleStates: []string{
		string(oci_core.InstanceLifecycleStateRunning),
	},
}

var exportCoreInternetGatewayHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_internet_gateway",
	DatasourceClass:      "oci_core_internet_gateways",
	DatasourceItemsAttr:  "gateways",
	ResourceAbbreviation: "internet_gateway",
	DiscoverableLifecycleStates: []string{
		string(oci_core.InternetGatewayLifecycleStateAvailable),
	},
}

var exportCoreIpSecConnectionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_ipsec",
	DatasourceClass:      "oci_core_ipsec_connections",
	DatasourceItemsAttr:  "connections",
	ResourceAbbreviation: "ip_sec_connection",
	DiscoverableLifecycleStates: []string{
		string(oci_core.IpSecConnectionLifecycleStateAvailable),
	},
}

var exportCoreLocalPeeringGatewayHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_local_peering_gateway",
	DatasourceClass:      "oci_core_local_peering_gateways",
	DatasourceItemsAttr:  "local_peering_gateways",
	ResourceAbbreviation: "local_peering_gateway",
	DiscoverableLifecycleStates: []string{
		string(oci_core.LocalPeeringGatewayLifecycleStateAvailable),
	},
}

var exportCoreNatGatewayHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_nat_gateway",
	DatasourceClass:      "oci_core_nat_gateways",
	DatasourceItemsAttr:  "nat_gateways",
	ResourceAbbreviation: "nat_gateway",
	DiscoverableLifecycleStates: []string{
		string(oci_core.NatGatewayLifecycleStateAvailable),
	},
}

var exportCoreNetworkSecurityGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_network_security_group",
	DatasourceClass:      "oci_core_network_security_groups",
	DatasourceItemsAttr:  "network_security_groups",
	ResourceAbbreviation: "network_security_group",
	DiscoverableLifecycleStates: []string{
		string(oci_core.NetworkSecurityGroupLifecycleStateAvailable),
	},
}

var exportCoreNetworkSecurityGroupSecurityRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_network_security_group_security_rule",
	ResourceAbbreviation: "network_security_group_security_rule",
}

var exportCorePrivateIpHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_private_ip",
	DatasourceClass:      "oci_core_private_ips",
	DatasourceItemsAttr:  "private_ips",
	ResourceAbbreviation: "private_ip",
}

var exportCorePublicIpHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_public_ip",
	DatasourceClass:      "oci_core_public_ips",
	DatasourceItemsAttr:  "public_ips",
	ResourceAbbreviation: "public_ip",
	DiscoverableLifecycleStates: []string{
		string(oci_core.PublicIpLifecycleStateAvailable),
		string(oci_core.PublicIpLifecycleStateAssigned),
	},
}

var exportCoreRemotePeeringConnectionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_remote_peering_connection",
	DatasourceClass:      "oci_core_remote_peering_connections",
	DatasourceItemsAttr:  "remote_peering_connections",
	ResourceAbbreviation: "remote_peering_connection",
	DiscoverableLifecycleStates: []string{
		string(oci_core.RemotePeeringConnectionLifecycleStateAvailable),
	},
}

var exportCoreRouteTableHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_route_table",
	DatasourceClass:      "oci_core_route_tables",
	DatasourceItemsAttr:  "route_tables",
	ResourceAbbreviation: "route_table",
	DiscoverableLifecycleStates: []string{
		string(oci_core.RouteTableLifecycleStateAvailable),
	},
}

var exportCoreSecurityListHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_security_list",
	DatasourceClass:      "oci_core_security_lists",
	DatasourceItemsAttr:  "security_lists",
	ResourceAbbreviation: "security_list",
	DiscoverableLifecycleStates: []string{
		string(oci_core.SecurityListLifecycleStateAvailable),
	},
}

var exportCoreServiceGatewayHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_service_gateway",
	DatasourceClass:      "oci_core_service_gateways",
	DatasourceItemsAttr:  "service_gateways",
	ResourceAbbreviation: "service_gateway",
	DiscoverableLifecycleStates: []string{
		string(oci_core.ServiceGatewayLifecycleStateAvailable),
	},
}

var exportCoreSubnetHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_subnet",
	DatasourceClass:      "oci_core_subnets",
	DatasourceItemsAttr:  "subnets",
	ResourceAbbreviation: "subnet",
	DiscoverableLifecycleStates: []string{
		string(oci_core.SubnetLifecycleStateAvailable),
	},
}

var exportCoreVcnHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_vcn",
	DatasourceClass:      "oci_core_vcns",
	DatasourceItemsAttr:  "virtual_networks",
	ResourceAbbreviation: "vcn",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VcnLifecycleStateAvailable),
	},
}

var exportCoreVlanHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_vlan",
	DatasourceClass:      "oci_core_vlans",
	DatasourceItemsAttr:  "vlans",
	ResourceAbbreviation: "vlan",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VlanLifecycleStateAvailable),
	},
}

var exportCoreVirtualCircuitHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_virtual_circuit",
	DatasourceClass:      "oci_core_virtual_circuits",
	DatasourceItemsAttr:  "virtual_circuits",
	ResourceAbbreviation: "virtual_circuit",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VirtualCircuitLifecycleStatePendingProvider),
		string(oci_core.VirtualCircuitLifecycleStateProvisioned),
	},
}

var exportCoreVnicAttachmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_vnic_attachment",
	DatasourceClass:      "oci_core_vnic_attachments",
	DatasourceItemsAttr:  "vnic_attachments",
	ResourceAbbreviation: "vnic_attachment",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VnicAttachmentLifecycleStateAttached),
	},
}

var exportCoreVolumeAttachmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_volume_attachment",
	DatasourceClass:      "oci_core_volume_attachments",
	DatasourceItemsAttr:  "volume_attachments",
	ResourceAbbreviation: "volume_attachment",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VolumeAttachmentLifecycleStateAttached),
	},
}

var exportCoreVolumeBackupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_volume_backup",
	DatasourceClass:      "oci_core_volume_backups",
	DatasourceItemsAttr:  "volume_backups",
	ResourceAbbreviation: "volume_backup",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VolumeBackupLifecycleStateAvailable),
	},
}

var exportCoreVolumeBackupPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_volume_backup_policy",
	DatasourceClass:      "oci_core_volume_backup_policies",
	DatasourceItemsAttr:  "volume_backup_policies",
	ResourceAbbreviation: "volume_backup_policy",
}

var exportCoreVolumeBackupPolicyAssignmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_volume_backup_policy_assignment",
	DatasourceClass:      "oci_core_volume_backup_policy_assignments",
	DatasourceItemsAttr:  "volume_backup_policy_assignments",
	ResourceAbbreviation: "volume_backup_policy_assignment",
}

var exportCoreVolumeGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_volume_group",
	DatasourceClass:      "oci_core_volume_groups",
	DatasourceItemsAttr:  "volume_groups",
	ResourceAbbreviation: "volume_group",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VolumeGroupLifecycleStateAvailable),
	},
}

var exportCoreVolumeGroupBackupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_volume_group_backup",
	DatasourceClass:      "oci_core_volume_group_backups",
	DatasourceItemsAttr:  "volume_group_backups",
	ResourceAbbreviation: "volume_group_backup",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VolumeGroupBackupLifecycleStateCommitted),
		string(oci_core.VolumeGroupBackupLifecycleStateAvailable),
	},
}

var exportCoreVolumeHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_volume",
	DatasourceClass:      "oci_core_volumes",
	DatasourceItemsAttr:  "volumes",
	ResourceAbbreviation: "volume",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VolumeLifecycleStateAvailable),
	},
}

var exportCorePublicIpPoolHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_core_public_ip_pool",
	DatasourceClass:        "oci_core_public_ip_pools",
	DatasourceItemsAttr:    "public_ip_pool_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "public_ip_pool",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_core.PublicIpPoolLifecycleStateActive),
	},
}

var exportCoreIpv6Hints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_ipv6",
	DatasourceClass:      "oci_core_ipv6s",
	DatasourceItemsAttr:  "ipv6s",
	ResourceAbbreviation: "ipv6",
	DiscoverableLifecycleStates: []string{
		string(oci_core.Ipv6LifecycleStateAvailable),
	},
}

var exportCoreDrgRouteTableHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_drg_route_table",
	DatasourceClass:      "oci_core_drg_route_tables",
	DatasourceItemsAttr:  "drg_route_tables",
	ResourceAbbreviation: "drg_route_table",
	DiscoverableLifecycleStates: []string{
		string(oci_core.DrgRouteTableLifecycleStateAvailable),
	},
}

var exportCoreDrgRouteDistributionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_drg_route_distribution",
	DatasourceClass:      "oci_core_drg_route_distributions",
	DatasourceItemsAttr:  "drg_route_distributions",
	ResourceAbbreviation: "drg_route_distribution",
	DiscoverableLifecycleStates: []string{
		string(oci_core.DrgRouteDistributionLifecycleStateAvailable),
	},
}

var exportCoreDrgRouteTableRouteRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_drg_route_table_route_rule",
	ResourceAbbreviation: "drg_route_table_route_rule",
}

var exportCoreCaptureFilterHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_capture_filter",
	DatasourceClass:      "oci_core_capture_filters",
	DatasourceItemsAttr:  "capture_filters",
	ResourceAbbreviation: "capture_filter",
	DiscoverableLifecycleStates: []string{
		string(oci_core.CaptureFilterLifecycleStateAvailable),
	},
}

var exportCoreVtapHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_vtap",
	DatasourceClass:      "oci_core_vtaps",
	DatasourceItemsAttr:  "vtaps",
	ResourceAbbreviation: "vtap",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VtapLifecycleStateAvailable),
	},
}

var exportCoreComputeClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_core_compute_cluster",
	DatasourceClass:        "oci_core_compute_clusters",
	DatasourceItemsAttr:    "compute_cluster_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "compute_cluster",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_core.ComputeClusterLifecycleStateActive),
	},
}

var exportCoreComputeCapacityReportHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_compute_capacity_report",
	ResourceAbbreviation: "compute_capacity_report",
}

var exportCoreComputeCapacityTopologyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_core_compute_capacity_topology",
	DatasourceClass:        "oci_core_compute_capacity_topologies",
	DatasourceItemsAttr:    "compute_capacity_topology_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "compute_capacity_topology",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_core.ComputeCapacityTopologyLifecycleStateActive),
	},
}

var coreResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCoreBootVolumeBackupHints},
		{TerraformResourceHints: exportCoreBootVolumeHints},
		{TerraformResourceHints: exportCoreConsoleHistoryHints},
		{TerraformResourceHints: exportCoreClusterNetworkHints},
		{TerraformResourceHints: exportCoreComputeImageCapabilitySchemaHints},
		{TerraformResourceHints: exportCoreComputeCapacityReservationHints},
		{TerraformResourceHints: exportCoreCpeHints},
		{TerraformResourceHints: exportCoreCrossConnectGroupHints},
		{TerraformResourceHints: exportCoreCrossConnectHints},
		{TerraformResourceHints: exportCoreDhcpOptionsHints},
		{TerraformResourceHints: exportCoreNatGatewayHints},
		{TerraformResourceHints: exportCoreDrgAttachmentHints},
		{TerraformResourceHints: exportCoreDrgHints},
		{TerraformResourceHints: exportCoreDedicatedVmHostHints},
		{TerraformResourceHints: exportCoreImageHints},
		{TerraformResourceHints: exportCoreInstanceConfigurationHints},
		{TerraformResourceHints: exportCoreInstanceConsoleConnectionHints},
		{TerraformResourceHints: exportCoreInstancePoolHints},
		{TerraformResourceHints: exportCoreInstanceHints},
		{TerraformResourceHints: exportCoreInternetGatewayHints},
		{TerraformResourceHints: exportCoreIpSecConnectionHints},
		{TerraformResourceHints: exportCoreLocalPeeringGatewayHints},
		{TerraformResourceHints: exportCoreNetworkSecurityGroupHints},
		{
			TerraformResourceHints: exportCorePublicIpHints,
			DatasourceQueryParams: map[string]string{
				"scope": "'REGION'",
			},
		},
		{TerraformResourceHints: exportCoreRemotePeeringConnectionHints},
		{TerraformResourceHints: exportCoreRouteTableHints},
		{TerraformResourceHints: exportCoreSecurityListHints},
		{TerraformResourceHints: exportCoreServiceGatewayHints},
		{TerraformResourceHints: exportCoreSubnetHints},
		{TerraformResourceHints: exportCoreVcnHints},
		{TerraformResourceHints: exportCoreVlanHints},
		{TerraformResourceHints: exportCoreVirtualCircuitHints},
		{TerraformResourceHints: exportCoreVolumeAttachmentHints},
		{TerraformResourceHints: exportCoreVolumeBackupHints},
		{TerraformResourceHints: exportCoreVolumeBackupPolicyHints},
		{TerraformResourceHints: exportCoreVolumeGroupHints},
		{TerraformResourceHints: exportCoreVolumeGroupBackupHints},
		{TerraformResourceHints: exportCoreVolumeHints},
		{TerraformResourceHints: exportCorePublicIpPoolHints},
		{TerraformResourceHints: exportCoreCaptureFilterHints},
		{TerraformResourceHints: exportCoreVtapHints},
		{TerraformResourceHints: exportCoreComputeClusterHints},
		{TerraformResourceHints: exportCoreComputeCapacityTopologyHints},
	},
	"oci_core_boot_volume": {
		{
			TerraformResourceHints: exportCoreVolumeBackupPolicyAssignmentHints,
			DatasourceQueryParams: map[string]string{
				"asset_id": "id",
			},
		},
	},
	"oci_core_drg": {
		{
			TerraformResourceHints: exportCoreDrgRouteDistributionHints,
			DatasourceQueryParams: map[string]string{
				"drg_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreDrgRouteTableHints,
			DatasourceQueryParams: map[string]string{
				"drg_id": "id",
			},
		},
	},
	"oci_core_instance_pool": {
		{
			TerraformResourceHints: exportCoreInstancePoolInstanceHints,
			DatasourceQueryParams: map[string]string{
				"instance_pool_id": "id",
			},
		},
	},
	"oci_core_instance": {
		{
			TerraformResourceHints: exportCoreVnicAttachmentHints,
			DatasourceQueryParams: map[string]string{
				"instance_id": "id",
			},
		},
		{
			TerraformResourceHints: exportCoreVolumeBackupPolicyAssignmentHints,
			DatasourceQueryParams: map[string]string{
				"asset_id": "boot_volume_id",
			},
		},
	},
	"oci_core_network_security_group": {
		{
			TerraformResourceHints: exportCoreNetworkSecurityGroupSecurityRuleHints,
			DatasourceQueryParams: map[string]string{
				"network_security_group_id": "id",
			},
		},
	},
	"oci_core_subnet": {
		{
			TerraformResourceHints: exportCorePrivateIpHints,
			DatasourceQueryParams: map[string]string{
				"subnet_id": "id",
			},
		},
	},
	"oci_core_volume": {
		{
			TerraformResourceHints: exportCoreVolumeBackupPolicyAssignmentHints,
			DatasourceQueryParams: map[string]string{
				"asset_id": "id",
			},
		},
	},
	"oci_core_drg_route_table": {
		{
			TerraformResourceHints: exportCoreDrgRouteTableRouteRuleHints,
			DatasourceQueryParams: map[string]string{
				"drg_route_table_id": "id",
			},
		},
	},
}

var relatedcoreinstance = []tf_export.TerraformResourceAssociation{
	{
		TerraformResourceHints: exportCoreVolumeAttachmentHints,
		DatasourceQueryParams: map[string]string{
			"availability_domain": "availability_domain",
			"instance_id":         "id",
		},
	}, {
		TerraformResourceHints: exportCoreVnicAttachmentHints,
		DatasourceQueryParams: map[string]string{
			"instance_id": "id",
		},
	},
}

var relatedcorevolumeattachment = []tf_export.TerraformResourceAssociation{
	{
		TerraformResourceHints: exportCoreVolumeClosureHints,
		DatasourceQueryParams: map[string]string{
			"volume_id": "volume_id",
		},
	},
}

// Separate hints for closure as we want to discover volumes related to an instance only
// and list data source for these give volumes for an AD or a volume Group only

var exportCoreVolumeClosureHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_core_volume",
	DatasourceClass:      "oci_core_volume",
	ResourceAbbreviation: "volume",
	DiscoverableLifecycleStates: []string{
		string(oci_core.VolumeLifecycleStateAvailable),
	},
}

var customAssociationCoreIdentityAvailabilityDomain = []tf_export.TerraformResourceAssociation{
	{
		TerraformResourceHints: exportCoreBootVolumeHints,
		DatasourceQueryParams: map[string]string{
			"availability_domain": "name",
		},
	},
}
