// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotDigitalTwinAdapterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["digital_twin_adapter_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(IotDigitalTwinAdapterResource(), fieldMap, readSingularIotDigitalTwinAdapterWithContext)
}

func readSingularIotDigitalTwinAdapterWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinAdapterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type IotDigitalTwinAdapterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_iot.IotClient
	Res    *oci_iot.GetDigitalTwinAdapterResponse
}

func (s *IotDigitalTwinAdapterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IotDigitalTwinAdapterDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetDigitalTwinAdapterRequest{}

	if digitalTwinAdapterId, ok := s.D.GetOkExists("digital_twin_adapter_id"); ok {
		tmp := digitalTwinAdapterId.(string)
		request.DigitalTwinAdapterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "iot")

	response, err := s.Client.GetDigitalTwinAdapter(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IotDigitalTwinAdapterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DigitalTwinModelId != nil {
		s.D.Set("digital_twin_model_id", *s.Res.DigitalTwinModelId)
	}

	if s.Res.DigitalTwinModelSpecUri != nil {
		s.D.Set("digital_twin_model_spec_uri", *s.Res.DigitalTwinModelSpecUri)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InboundEnvelope != nil {
		s.D.Set("inbound_envelope", []interface{}{DigitalTwinAdapterInboundEnvelopeToMap(s.Res.InboundEnvelope)})
	} else {
		s.D.Set("inbound_envelope", nil)
	}

	inboundRoutes := []interface{}{}
	for _, item := range s.Res.InboundRoutes {
		inboundRoutes = append(inboundRoutes, DigitalTwinAdapterInboundRouteToMap(item))
	}
	s.D.Set("inbound_routes", inboundRoutes)

	if s.Res.IotDomainId != nil {
		s.D.Set("iot_domain_id", *s.Res.IotDomainId)
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

	return nil
}
