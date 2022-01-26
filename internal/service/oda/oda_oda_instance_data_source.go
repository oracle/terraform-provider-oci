// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_oda "github.com/oracle/oci-go-sdk/v56/oda"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OdaOdaInstanceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oda_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OdaOdaInstanceResource(), fieldMap, readSingularOdaOdaInstance)
}

func readSingularOdaOdaInstance(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaInstanceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OdaClient()

	return tfresource.ReadResource(sync)
}

type OdaOdaInstanceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oda.OdaClient
	Res    *oci_oda.GetOdaInstanceResponse
}

func (s *OdaOdaInstanceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OdaOdaInstanceDataSourceCrud) Get() error {
	request := oci_oda.GetOdaInstanceRequest{}

	if odaInstanceId, ok := s.D.GetOkExists("oda_instance_id"); ok {
		tmp := odaInstanceId.(string)
		request.OdaInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oda")

	response, err := s.Client.GetOdaInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OdaOdaInstanceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectorUrl != nil {
		s.D.Set("connector_url", *s.Res.ConnectorUrl)
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

	s.D.Set("lifecycle_sub_state", s.Res.LifecycleSubState)

	s.D.Set("shape_name", s.Res.ShapeName)

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

	if s.Res.WebAppUrl != nil {
		s.D.Set("web_app_url", *s.Res.WebAppUrl)
	}

	return nil
}
