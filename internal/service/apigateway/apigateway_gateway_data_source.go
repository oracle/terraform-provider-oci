// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apigateway "github.com/oracle/oci-go-sdk/v56/apigateway"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ApigatewayGatewayDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["gateway_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApigatewayGatewayResource(), fieldMap, readSingularApigatewayGateway)
}

func readSingularApigatewayGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayGatewayDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GatewayClient()

	return tfresource.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apigateway")

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

	caBundles := []interface{}{}
	for _, item := range s.Res.CaBundles {
		caBundles = append(caBundles, CaBundleToMap(item))
	}
	s.D.Set("ca_bundles", caBundles)

	if s.Res.CertificateId != nil {
		s.D.Set("certificate_id", *s.Res.CertificateId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("endpoint_type", s.Res.EndpointType)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	ipAddresses := []interface{}{}
	for _, item := range s.Res.IpAddresses {
		ipAddresses = append(ipAddresses, GatewayIpAddressToMap(item))
	}
	s.D.Set("ip_addresses", ipAddresses)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("network_security_group_ids", s.Res.NetworkSecurityGroupIds)

	if s.Res.ResponseCacheDetails != nil {
		responseCacheDetailsArray := []interface{}{}
		if responseCacheDetailsMap := ResponseCacheDetailsToMap(&s.Res.ResponseCacheDetails); responseCacheDetailsMap != nil {
			responseCacheDetailsArray = append(responseCacheDetailsArray, responseCacheDetailsMap)
		}
		s.D.Set("response_cache_details", responseCacheDetailsArray)
	} else {
		s.D.Set("response_cache_details", nil)
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
