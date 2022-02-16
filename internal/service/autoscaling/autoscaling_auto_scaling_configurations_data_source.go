// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package autoscaling

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_auto_scaling "github.com/oracle/oci-go-sdk/v58/autoscaling"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func AutoScalingAutoScalingConfigurationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAutoScalingAutoScalingConfigurations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
				Elem:     tfresource.GetDataSourceItemSchema(AutoScalingAutoScalingConfigurationResource()),
			},
		},
	}
}

func readAutoScalingAutoScalingConfigurations(d *schema.ResourceData, m interface{}) error {
	sync := &AutoScalingAutoScalingConfigurationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AutoScalingClient()

	return tfresource.ReadResource(sync)
}

type AutoScalingAutoScalingConfigurationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_auto_scaling.AutoScalingClient
	Res    *oci_auto_scaling.ListAutoScalingConfigurationsResponse
}

func (s *AutoScalingAutoScalingConfigurationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AutoScalingAutoScalingConfigurationsDataSourceCrud) Get() error {
	request := oci_auto_scaling.ListAutoScalingConfigurationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "auto_scaling")

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

func (s *AutoScalingAutoScalingConfigurationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AutoScalingAutoScalingConfigurationsDataSource-", AutoScalingAutoScalingConfigurationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autoScalingConfiguration := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CoolDownInSeconds != nil {
			autoScalingConfiguration["cool_down_in_seconds"] = *r.CoolDownInSeconds
		}

		if r.DefinedTags != nil {
			autoScalingConfiguration["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			autoScalingConfiguration["display_name"] = *r.DisplayName
		}

		autoScalingConfiguration["freeform_tags"] = r.FreeformTags

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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, AutoScalingAutoScalingConfigurationsDataSource().Schema["auto_scaling_configurations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("auto_scaling_configurations", resources); err != nil {
		return err
	}

	return nil
}
