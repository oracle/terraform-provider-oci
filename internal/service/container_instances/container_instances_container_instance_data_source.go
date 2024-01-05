// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package container_instances

import (
	"context"
	"strconv"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_container_instances "github.com/oracle/oci-go-sdk/v65/containerinstances"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ContainerInstancesContainerInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["container_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ContainerInstancesContainerInstanceResource(), fieldMap, readSingularContainerInstancesContainerInstance)
}

func readSingularContainerInstancesContainerInstance(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerInstancesContainerInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ContainerInstanceClient()
	sync.VirtualNetworkClient = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type ContainerInstancesContainerInstanceDataSourceCrud struct {
	D                    *schema.ResourceData
	Client               *oci_container_instances.ContainerInstanceClient
	VirtualNetworkClient *oci_core.VirtualNetworkClient
	Res                  *oci_container_instances.GetContainerInstanceResponse
}

func (s *ContainerInstancesContainerInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerInstancesContainerInstanceDataSourceCrud) Get() error {
	request := oci_container_instances.GetContainerInstanceRequest{}

	if containerInstanceId, ok := s.D.GetOkExists("container_instance_id"); ok {
		tmp := containerInstanceId.(string)
		request.ContainerInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerinstance")

	response, err := s.Client.GetContainerInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerInstancesContainerInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ContainerCount != nil {
		s.D.Set("container_count", *s.Res.ContainerCount)
	}

	s.D.Set("container_restart_policy", s.Res.ContainerRestartPolicy)

	containers := []interface{}{}
	for _, item := range s.Res.Containers {
		result := map[string]interface{}{}

		if item.ContainerId != nil {
			result["container_id"] = string(*item.ContainerId)

			request := oci_container_instances.GetContainerRequest{}
			request.ContainerId = item.ContainerId

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerinstance")

			response, err := s.Client.GetContainer(context.Background(), request)
			if err != nil {
				return err
			}

			container := response.Container
			result = ContainerToMap(container)
		}
		containers = append(containers, result)
	}
	s.D.Set("containers", containers)

	vnics := []interface{}{}
	for _, item := range s.Res.Vnics {

		result := map[string]interface{}{}

		if item.VnicId != nil {
			result["vnic_id"] = string(*item.VnicId)

			request := oci_core.GetVnicRequest{}
			request.VnicId = item.VnicId

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "containerinstance")

			response, err := s.VirtualNetworkClient.GetVnic(context.Background(), request)
			if err != nil {
				return err
			}

			vnic := response.Vnic
			result = VnicDetailsToMap(vnic, true)
		}
		vnics = append(vnics, result)
	}
	s.D.Set("vnics", vnics)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DnsConfig != nil {
		s.D.Set("dns_config", []interface{}{ContainerDnsConfigToMap(s.Res.DnsConfig)})
	} else {
		s.D.Set("dns_config", nil)
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GracefulShutdownTimeoutInSeconds != nil {
		s.D.Set("graceful_shutdown_timeout_in_seconds", strconv.FormatInt(*s.Res.GracefulShutdownTimeoutInSeconds, 10))
	}

	imagePullSecrets := []interface{}{}
	for index, item := range s.Res.ImagePullSecrets {
		imagePullSecrets = append(imagePullSecrets, ImagePullSecretToMap(item, s.D, index))
	}
	s.D.Set("image_pull_secrets", imagePullSecrets)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	if s.Res.ShapeConfig != nil {
		s.D.Set("shape_config", []interface{}{ContainerInstanceShapeConfigToMap(s.Res.ShapeConfig)})
	} else {
		s.D.Set("shape_config", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VolumeCount != nil {
		s.D.Set("volume_count", *s.Res.VolumeCount)
	}

	volumes := []interface{}{}
	for index, item := range s.Res.Volumes {
		volumes = append(volumes, ContainerVolumeToMap(item, s.D, index))
	}
	s.D.Set("volumes", volumes)

	return nil
}
