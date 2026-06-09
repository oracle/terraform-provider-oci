// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreCaptureFilterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["capture_filter_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(CoreCaptureFilterResource(), fieldMap, readSingularCoreCaptureFilterWithContext)
}

func readSingularCoreCaptureFilterWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &CoreCaptureFilterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type CoreCaptureFilterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetCaptureFilterResponse
}

func (s *CoreCaptureFilterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCaptureFilterDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_core.GetCaptureFilterRequest{}

	if captureFilterId, ok := s.D.GetOkExists("capture_filter_id"); ok {
		tmp := captureFilterId.(string)
		request.CaptureFilterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetCaptureFilter(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreCaptureFilterDataSourceCrud) SetData() error {
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

	s.D.Set("filter_type", s.Res.FilterType)

	flowLogCaptureFilterRules := []interface{}{}
	for _, item := range s.Res.FlowLogCaptureFilterRules {
		flowLogCaptureFilterRules = append(flowLogCaptureFilterRules, FlowLogCaptureFilterRuleDetailsToMap(item))
	}
	s.D.Set("flow_log_capture_filter_rules", flowLogCaptureFilterRules)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	vtapCaptureFilterRules := []interface{}{}
	for _, item := range s.Res.VtapCaptureFilterRules {
		vtapCaptureFilterRules = append(vtapCaptureFilterRules, VtapCaptureFilterRuleDetailsToMap(item))
	}
	s.D.Set("vtap_capture_filter_rules", vtapCaptureFilterRules)

	return nil
}
