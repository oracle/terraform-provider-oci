// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OdaOdaPrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oda_private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OdaOdaPrivateEndpointResource(), fieldMap, readSingularOdaOdaPrivateEndpoint)
}

func readSingularOdaOdaPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()

	return tfresource.ReadResource(sync)
}

type OdaOdaPrivateEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oda.ManagementClient
	Res    *oci_oda.GetOdaPrivateEndpointResponse
}

func (s *OdaOdaPrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OdaOdaPrivateEndpointDataSourceCrud) Get() error {
	request := oci_oda.GetOdaPrivateEndpointRequest{}

	if odaPrivateEndpointId, ok := s.D.GetOkExists("oda_private_endpoint_id"); ok {
		tmp := odaPrivateEndpointId.(string)
		request.OdaPrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oda")

	response, err := s.Client.GetOdaPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OdaOdaPrivateEndpointDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("nsg_ids", s.Res.NsgIds)
	s.D.Set("nsg_ids", s.Res.NsgIds)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
