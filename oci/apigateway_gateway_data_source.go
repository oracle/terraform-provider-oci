// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/v25/apigateway"
)

func init() {
	RegisterDatasource("oci_apigateway_gateway", ApigatewayGatewayDataSource())
}

func ApigatewayGatewayDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["gateway_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(ApigatewayGatewayResource(), fieldMap, readSingularApigatewayGateway)
}

func readSingularApigatewayGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayGatewayDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).gatewayClient()

	return ReadResource(sync)
}

type ApigatewayGatewayDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apigateway.GatewayClient
	Res    *oci_apigateway.GetGatewayResponse
}

func (s *ApigatewayGatewayDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApigatewayGatewayDataSourceCrud) Get() error {
	request := oci_apigateway.GetGatewayRequest{}

	if gatewayId, ok := s.D.GetOkExists("gateway_id"); ok {
		tmp := gatewayId.(string)
		request.GatewayId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "apigateway")

	response, err := s.Client.GetGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApigatewayGatewayDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("endpoint_type", s.Res.EndpointType)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

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
