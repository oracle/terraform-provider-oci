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

func CoreCaptureFiltersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreCaptureFilters,
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
			"filter_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"capture_filters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreCaptureFilterResource()),
			},
		},
	}
}

func readCoreCaptureFilters(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCaptureFiltersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreCaptureFiltersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListCaptureFiltersResponse
}

func (s *CoreCaptureFiltersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCaptureFiltersDataSourceCrud) Get() error {
	request := oci_core.ListCaptureFiltersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if filterType, ok := s.D.GetOkExists("filter_type"); ok {
		request.FilterType = oci_core.CaptureFilterFilterTypeEnum(filterType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.CaptureFilterLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListCaptureFilters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCaptureFilters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreCaptureFiltersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreCaptureFiltersDataSource-", CoreCaptureFiltersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		captureFilter := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			captureFilter["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			captureFilter["display_name"] = *r.DisplayName
		}

		captureFilter["filter_type"] = r.FilterType

		flowLogCaptureFilterRules := []interface{}{}
		for _, item := range r.FlowLogCaptureFilterRules {
			flowLogCaptureFilterRules = append(flowLogCaptureFilterRules, FlowLogCaptureFilterRuleDetailsToMap(item))
		}
		captureFilter["flow_log_capture_filter_rules"] = flowLogCaptureFilterRules

		captureFilter["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			captureFilter["id"] = *r.Id
		}

		captureFilter["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			captureFilter["time_created"] = r.TimeCreated.String()
		}

		vtapCaptureFilterRules := []interface{}{}
		for _, item := range r.VtapCaptureFilterRules {
			vtapCaptureFilterRules = append(vtapCaptureFilterRules, VtapCaptureFilterRuleDetailsToMap(item))
		}
		captureFilter["vtap_capture_filter_rules"] = vtapCaptureFilterRules

		resources = append(resources, captureFilter)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreCaptureFiltersDataSource().Schema["capture_filters"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("capture_filters", resources); err != nil {
		return err
	}

	return nil
}
