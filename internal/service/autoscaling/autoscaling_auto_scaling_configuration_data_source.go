// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package autoscaling

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_auto_scaling "github.com/oracle/oci-go-sdk/v56/autoscaling"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func AutoScalingAutoScalingConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["auto_scaling_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(AutoScalingAutoScalingConfigurationResource(), fieldMap, readSingularAutoScalingAutoScalingConfiguration)
}

func readSingularAutoScalingAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AutoScalingAutoScalingConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AutoScalingClient()

	return tfresource.ReadResource(sync)
}

type AutoScalingAutoScalingConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_auto_scaling.AutoScalingClient
	Res    *oci_auto_scaling.GetAutoScalingConfigurationResponse
}

func (s *AutoScalingAutoScalingConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutoScalingAutoScalingConfigurationDataSourceCrud) Get() error {
	request := oci_auto_scaling.GetAutoScalingConfigurationRequest{}

	if autoScalingConfigurationId, ok := s.D.GetOkExists("auto_scaling_configuration_id"); ok {
		tmp := autoScalingConfigurationId.(string)
		request.AutoScalingConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "auto_scaling")

	response, err := s.Client.GetAutoScalingConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AutoScalingAutoScalingConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Resource != nil {
		autoScalingResourcesArray := []interface{}{}
		if autoScalingResourcesMap := ResourceToMap(&s.Res.Resource); autoScalingResourcesMap != nil {
			autoScalingResourcesArray = append(autoScalingResourcesArray, autoScalingResourcesMap)
		}
		s.D.Set("auto_scaling_resources", autoScalingResourcesArray)
	} else {
		s.D.Set("auto_scaling_resources", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CoolDownInSeconds != nil {
		s.D.Set("cool_down_in_seconds", *s.Res.CoolDownInSeconds)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.MaxResourceCount != nil {
		s.D.Set("max_resource_count", *s.Res.MaxResourceCount)
	}

	if s.Res.MinResourceCount != nil {
		s.D.Set("min_resource_count", *s.Res.MinResourceCount)
	}

	policies := []interface{}{}
	for _, item := range s.Res.Policies {
		policies = append(policies, AutoScalingPolicyToMap(item, true))
	}
	s.D.Set("policies", policies)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
