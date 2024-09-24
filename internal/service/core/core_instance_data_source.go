// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreInstanceResource(), fieldMap, readSingularCoreInstance)
}

func readSingularCoreInstance(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()
	sync.BlockStorageClient = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

type CoreInstanceDataSourceCrud struct {
	CoreInstanceResourceCrud
}

func (s *CoreInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstanceDataSourceCrud) Get() error {
	request := oci_core.GetInstanceRequest{}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Instance
	return nil
}

func (s *CoreInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AgentConfig != nil {
		s.D.Set("agent_config", []interface{}{InstanceAgentConfigToMap(s.Res.AgentConfig)})
	} else {
		s.D.Set("agent_config", nil)
	}

	if s.Res.AvailabilityConfig != nil {
		s.D.Set("availability_config", []interface{}{InstanceAvailabilityConfigToMap(s.Res.AvailabilityConfig)})
	} else {
		s.D.Set("availability_config", nil)
	}

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CapacityReservationId != nil {
		s.D.Set("capacity_reservation_id", *s.Res.CapacityReservationId)
	}

	if s.Res.ClusterPlacementGroupId != nil {
		s.D.Set("cluster_placement_group_id", *s.Res.ClusterPlacementGroupId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DedicatedVmHostId != nil {
		s.D.Set("dedicated_vm_host_id", *s.Res.DedicatedVmHostId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	// Extended metadata (a json blob) may not return with the same node order in which it
	// was originally created, the solution is to not set it here after subsequent GETS to
	// prevent inadvertent diffs or destroy/creates
	// if s.Res.ExtendedMetadata != nil {
	// // extended_metadata is an arbitrarily structured json object, `objectToMap` would not work
	// 	s.D.Set("extended_metadata", []interface{}{objectToMap(s.Res.ExtendedMetadata)})
	// }

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageId != nil {
		s.D.Set("image", *s.Res.ImageId)
	}

	if s.Res.InstanceConfigurationId != nil {
		s.D.Set("instance_configuration_id", *s.Res.InstanceConfigurationId)
	}

	if s.Res.InstanceOptions != nil {
		s.D.Set("instance_options", []interface{}{InstanceOptionsToMap(s.Res.InstanceOptions)})
	} else {
		s.D.Set("instance_options", nil)
	}

	if s.Res.IpxeScript != nil {
		s.D.Set("ipxe_script", *s.Res.IpxeScript)
	}

	if s.Res.IsCrossNumaNode != nil {
		s.D.Set("is_cross_numa_node", *s.Res.IsCrossNumaNode)
	}

	s.D.Set("launch_mode", s.Res.LaunchMode)

	if s.Res.LaunchOptions != nil {
		s.D.Set("launch_options", []interface{}{LaunchOptionsToMap(s.Res.LaunchOptions)})
	} else {
		s.D.Set("launch_options", nil)
	}

	if s.Res.Metadata != nil {
		err := s.D.Set("metadata", s.Res.Metadata)
		if err != nil {
			log.Printf("error setting metadata %q", err)
		}
	}

	if s.Res.PlatformConfig != nil {
		platformConfigArray := []interface{}{}
		if platformConfigMap := PlatformConfigToMap(&s.Res.PlatformConfig); platformConfigMap != nil {
			platformConfigArray = append(platformConfigArray, platformConfigMap)
		}
		s.D.Set("platform_config", platformConfigArray)
	} else {
		s.D.Set("platform_config", nil)
	}

	if s.Res.PreemptibleInstanceConfig != nil {
		s.D.Set("preemptible_instance_config", []interface{}{PreemptibleInstanceConfigDetailsToMap(s.Res.PreemptibleInstanceConfig)})
	} else {
		s.D.Set("preemptible_instance_config", nil)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	if s.Res.SecurityAttributes != nil {
		s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))
	}

	s.D.Set("security_attributes_state", s.Res.SecurityAttributesState)

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	if s.Res.ShapeConfig != nil {
		s.D.Set("shape_config", []interface{}{InstanceShapeConfigToMap(s.Res.ShapeConfig)})
	} else {
		s.D.Set("shape_config", nil)
	}

	bootVolume, bootVolumeErr := s.getBootVolume()
	if bootVolumeErr != nil {
		log.Printf("[WARN] Could not get the boot volume: %q", bootVolumeErr)
	}

	if s.Res.SourceDetails != nil {
		var sourceDetailsFromConfig map[string]interface{}
		if details, ok := s.D.GetOkExists("source_details"); ok {
			if tmpList := details.([]interface{}); len(tmpList) > 0 {
				sourceDetailsFromConfig = tmpList[0].(map[string]interface{})
			}
		}
		sourceDetailsArray := []interface{}{}
		if sourceDetailsMap := InstanceSourceDetailsToMap(&s.Res.SourceDetails, bootVolume, sourceDetailsFromConfig); sourceDetailsMap != nil {
			sourceDetailsArray = append(sourceDetailsArray, sourceDetailsMap)
		}
		err := s.D.Set("source_details", sourceDetailsArray)
		if err != nil {
			return err
		}
	} else {
		s.D.Set("source_details", nil)
	}

	if bootVolume != nil && bootVolume.Id != nil {
		s.D.Set("boot_volume_id", *bootVolume.Id)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeMaintenanceRebootDue != nil {
		s.D.Set("time_maintenance_reboot_due", s.Res.TimeMaintenanceRebootDue.String())
	}

	if s.Res.LifecycleState == oci_core.InstanceLifecycleStateRunning {
		vnic, vnicError := s.getPrimaryVnic()
		if vnicError != nil || vnic == nil {
			log.Printf("[WARN] Primary VNIC could not be found during instance refresh: %q", vnicError)
		} else {
			s.D.Set("hostname_label", vnic.HostnameLabel)
			s.D.Set("public_ip", vnic.PublicIp)
			s.D.Set("private_ip", vnic.PrivateIp)
			s.D.Set("subnet_id", vnic.SubnetId)
		}
	}

	return nil
}
