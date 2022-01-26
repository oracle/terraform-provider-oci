// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package visual_builder

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_visual_builder "github.com/oracle/oci-go-sdk/v56/visualbuilder"
)

func VisualBuilderVbInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["vb_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(VisualBuilderVbInstanceResource(), fieldMap, readSingularVisualBuilderVbInstance)
}

func readSingularVisualBuilderVbInstance(d *schema.ResourceData, m interface{}) error {
	sync := &VisualBuilderVbInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VbInstanceClient()

	return tfresource.ReadResource(sync)
}

type VisualBuilderVbInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_visual_builder.VbInstanceClient
	Res    *oci_visual_builder.GetVbInstanceResponse
}

func (s *VisualBuilderVbInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VisualBuilderVbInstanceDataSourceCrud) Get() error {
	request := oci_visual_builder.GetVbInstanceRequest{}

	if vbInstanceId, ok := s.D.GetOkExists("vb_instance_id"); ok {
		tmp := vbInstanceId.(string)
		request.VbInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "visual_builder")

	response, err := s.Client.GetVbInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *VisualBuilderVbInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	alternateCustomEndpoints := []interface{}{}
	for _, item := range s.Res.AlternateCustomEndpoints {
		alternateCustomEndpoints = append(alternateCustomEndpoints, VbCustomEndpointDetailsToMap(&item))
	}
	s.D.Set("alternate_custom_endpoints", alternateCustomEndpoints)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("consumption_model", s.Res.ConsumptionModel)

	if s.Res.CustomEndpoint != nil {
		s.D.Set("custom_endpoint", []interface{}{VbCustomEndpointDetailsToMap(s.Res.CustomEndpoint)})
	} else {
		s.D.Set("custom_endpoint", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceUrl != nil {
		s.D.Set("instance_url", *s.Res.InstanceUrl)
	}

	if s.Res.IsVisualBuilderEnabled != nil {
		s.D.Set("is_visual_builder_enabled", *s.Res.IsVisualBuilderEnabled)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
