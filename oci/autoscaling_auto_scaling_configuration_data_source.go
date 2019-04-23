// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_autoscaling "github.com/oracle/oci-go-sdk/autoscaling"
)

func AutoscalingAutoScalingConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["auto_scaling_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(AutoscalingAutoScalingConfigurationResource(), fieldMap, readSingularAutoscalingAutoScalingConfiguration)
}

func readSingularAutoscalingAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &AutoscalingAutoScalingConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).autoScalingClient

	return ReadResource(sync)
}

type AutoscalingAutoScalingConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_autoscaling.AutoScalingClient
	Res    *oci_autoscaling.GetAutoScalingConfigurationResponse
}

func (s *AutoscalingAutoScalingConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutoscalingAutoScalingConfigurationDataSourceCrud) Get() error {
	request := oci_autoscaling.GetAutoScalingConfigurationRequest{}

	if autoScalingConfigurationId, ok := s.D.GetOkExists("auto_scaling_configuration_id"); ok {
		tmp := autoScalingConfigurationId.(string)
		request.AutoScalingConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "autoscaling")

	response, err := s.Client.GetAutoScalingConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AutoscalingAutoScalingConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CoolDownInSeconds != nil {
		s.D.Set("cool_down_in_seconds", *s.Res.CoolDownInSeconds)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	policies := []interface{}{}
	for _, item := range s.Res.Policies {
		policies = append(policies, AutoScalingPolicyToMap(item, true))
	}
	s.D.Set("policies", policies)

	if s.Res.Resource != nil {
		resourceArray := []interface{}{}
		if resourceMap := ResourceToMap(&s.Res.Resource); resourceMap != nil {
			resourceArray = append(resourceArray, resourceMap)
		}
		s.D.Set("auto_scaling_resources", resourceArray)
	} else {
		s.D.Set("auto_scaling_resources", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
