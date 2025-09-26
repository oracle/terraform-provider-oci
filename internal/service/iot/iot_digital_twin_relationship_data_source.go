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

func IotDigitalTwinRelationshipDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["digital_twin_relationship_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(IotDigitalTwinRelationshipResource(), fieldMap, readSingularIotDigitalTwinRelationshipWithContext)
}

func readSingularIotDigitalTwinRelationshipWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinRelationshipDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type IotDigitalTwinRelationshipDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_iot.IotClient
	Res    *oci_iot.GetDigitalTwinRelationshipResponse
}

func (s *IotDigitalTwinRelationshipDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IotDigitalTwinRelationshipDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetDigitalTwinRelationshipRequest{}

	if digitalTwinRelationshipId, ok := s.D.GetOkExists("digital_twin_relationship_id"); ok {
		tmp := digitalTwinRelationshipId.(string)
		request.DigitalTwinRelationshipId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "iot")

	response, err := s.Client.GetDigitalTwinRelationship(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IotDigitalTwinRelationshipDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Content != nil {
		contentStr, err := MapToJsonString(s.Res.Content)
		if err != nil {
			return err
		}
		s.D.Set("content", contentStr)
	}

	if s.Res.ContentPath != nil {
		s.D.Set("content_path", *s.Res.ContentPath)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IotDomainId != nil {
		s.D.Set("iot_domain_id", *s.Res.IotDomainId)
	}

	if s.Res.SourceDigitalTwinInstanceId != nil {
		s.D.Set("source_digital_twin_instance_id", *s.Res.SourceDigitalTwinInstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetDigitalTwinInstanceId != nil {
		s.D.Set("target_digital_twin_instance_id", *s.Res.TargetDigitalTwinInstanceId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
