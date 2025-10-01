// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiGenerativeAiPrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["generative_ai_private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GenerativeAiGenerativeAiPrivateEndpointResource(), fieldMap, readSingularGenerativeAiGenerativeAiPrivateEndpoint)
}

func readSingularGenerativeAiGenerativeAiPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiGenerativeAiPrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.ReadResource(sync)
}

type GenerativeAiGenerativeAiPrivateEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.GetGenerativeAiPrivateEndpointResponse
}

func (s *GenerativeAiGenerativeAiPrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiGenerativeAiPrivateEndpointDataSourceCrud) Get() error {
	request := oci_generative_ai.GetGenerativeAiPrivateEndpointRequest{}

	if generativeAiPrivateEndpointId, ok := s.D.GetOkExists("generative_ai_private_endpoint_id"); ok {
		tmp := generativeAiPrivateEndpointId.(string)
		request.GenerativeAiPrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.GetGenerativeAiPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiGenerativeAiPrivateEndpointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Fqdn != nil {
		s.D.Set("fqdn", *s.Res.Fqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.PrivateEndpointIp != nil {
		s.D.Set("private_endpoint_ip", *s.Res.PrivateEndpointIp)
	}

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
