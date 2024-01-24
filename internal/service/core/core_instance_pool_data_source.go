// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreInstancePoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["instance_pool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreInstancePoolResource(), fieldMap, readSingularCoreInstancePool)
}

func readSingularCoreInstancePool(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstancePoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ComputeManagementClient()

	return tfresource.ReadResource(sync)
}

type CoreInstancePoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeManagementClient
	Res    *oci_core.GetInstancePoolResponse
}

func (s *CoreInstancePoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstancePoolDataSourceCrud) Get() error {
	request := oci_core.GetInstancePoolRequest{}

	if instancePoolId, ok := s.D.GetOkExists("instance_pool_id"); ok {
		tmp := instancePoolId.(string)
		request.InstancePoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetInstancePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreInstancePoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceConfigurationId != nil {
		s.D.Set("instance_configuration_id", *s.Res.InstanceConfigurationId)
	}

	if s.Res.InstanceDisplayNameFormatter != nil {
		s.D.Set("instance_display_name_formatter", *s.Res.InstanceDisplayNameFormatter)
	}

	if s.Res.InstanceHostnameFormatter != nil {
		s.D.Set("instance_hostname_formatter", *s.Res.InstanceHostnameFormatter)
	}

	loadBalancers := []interface{}{}
	for _, item := range s.Res.LoadBalancers {
		if item.LifecycleState != oci_core.InstancePoolLoadBalancerAttachmentLifecycleStateDetached {
			loadBalancers = append(loadBalancers, InstancePoolLoadBalancerAttachmentToMap(item))
		}
	}
	s.D.Set("load_balancers", loadBalancers)

	placementConfigurations := []interface{}{}
	for _, item := range s.Res.PlacementConfigurations {
		placementConfigurations = append(placementConfigurations, InstancePoolPlacementConfigurationToMap(item))
	}
	s.D.Set("placement_configurations", placementConfigurations)

	if s.Res.Size != nil {
		s.D.Set("size", *s.Res.Size)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
