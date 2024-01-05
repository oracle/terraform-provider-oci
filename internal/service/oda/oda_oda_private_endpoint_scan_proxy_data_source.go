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

func OdaOdaPrivateEndpointScanProxyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["oda_private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["oda_private_endpoint_scan_proxy_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OdaOdaPrivateEndpointScanProxyResource(), fieldMap, readSingularOdaOdaPrivateEndpointScanProxy)
}

func readSingularOdaOdaPrivateEndpointScanProxy(d *schema.ResourceData, m interface{}) error {
	sync := &OdaOdaPrivateEndpointScanProxyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementClient()

	return tfresource.ReadResource(sync)
}

type OdaOdaPrivateEndpointScanProxyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oda.ManagementClient
	Res    *oci_oda.GetOdaPrivateEndpointScanProxyResponse
}

func (s *OdaOdaPrivateEndpointScanProxyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OdaOdaPrivateEndpointScanProxyDataSourceCrud) Get() error {
	request := oci_oda.GetOdaPrivateEndpointScanProxyRequest{}

	if odaPrivateEndpointId, ok := s.D.GetOkExists("oda_private_endpoint_id"); ok {
		tmp := odaPrivateEndpointId.(string)
		request.OdaPrivateEndpointId = &tmp
	}

	if odaPrivateEndpointScanProxyId, ok := s.D.GetOkExists("oda_private_endpoint_scan_proxy_id"); ok {
		tmp := odaPrivateEndpointScanProxyId.(string)
		request.OdaPrivateEndpointScanProxyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oda")

	response, err := s.Client.GetOdaPrivateEndpointScanProxy(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OdaOdaPrivateEndpointScanProxyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("protocol", s.Res.Protocol)

	scanListenerInfos := []interface{}{}
	for _, item := range s.Res.ScanListenerInfos {
		scanListenerInfos = append(scanListenerInfos, ScanListenerInfoToMap(item))
	}
	s.D.Set("scan_listener_infos", scanListenerInfos)

	s.D.Set("scan_listener_type", s.Res.ScanListenerType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
