// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_autoscaling "github.com/oracle/oci-go-sdk/autoscaling"
)

func AutoscalingAutoScalingConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAutoscalingAutoScalingConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_scaling_configurations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(AutoscalingAutoScalingConfigurationResource()),
			},
		},
	}
}

func readAutoscalingAutoScalingConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &AutoscalingAutoScalingConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).autoScalingClient

	return ReadResource(sync)
}

type AutoscalingAutoScalingConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_autoscaling.AutoScalingClient
	Res    *oci_autoscaling.ListAutoScalingConfigurationsResponse
}

func (s *AutoscalingAutoScalingConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutoscalingAutoScalingConfigurationsDataSourceCrud) Get() error {
	request := oci_autoscaling.ListAutoScalingConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "autoscaling")

	response, err := s.Client.ListAutoScalingConfigurations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutoScalingConfigurations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AutoscalingAutoScalingConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autoScalingConfiguration := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CoolDownInSeconds != nil {
			autoScalingConfiguration["cool_down_in_seconds"] = *r.CoolDownInSeconds
		}

		if r.DisplayName != nil {
			autoScalingConfiguration["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			autoScalingConfiguration["id"] = *r.Id
		}

		if r.IsEnabled != nil {
			autoScalingConfiguration["is_enabled"] = *r.IsEnabled
		}

		if r.Resource != nil {
			resourceArray := []interface{}{}
			if resourceMap := ResourceToMap(&r.Resource); resourceMap != nil {
				resourceArray = append(resourceArray, resourceMap)
			}
			autoScalingConfiguration["auto_scaling_resources"] = resourceArray
		} else {
			autoScalingConfiguration["auto_scaling_resources"] = nil
		}

		if r.TimeCreated != nil {
			autoScalingConfiguration["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, autoScalingConfiguration)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, AutoscalingAutoScalingConfigurationsDataSource().Schema["auto_scaling_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("auto_scaling_configurations", resources); err != nil {
		return err
	}

	return nil
}
