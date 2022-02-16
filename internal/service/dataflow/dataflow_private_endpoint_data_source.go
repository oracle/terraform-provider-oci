// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_dataflow "github.com/oracle/oci-go-sdk/v58/dataflow"
)

func DataflowPrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataflowPrivateEndpointResource(), fieldMap, readSingularDataflowPrivateEndpoint)
}

func readSingularDataflowPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowPrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

type DataflowPrivateEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dataflow.DataFlowClient
	Res    *oci_dataflow.GetPrivateEndpointResponse
}

func (s *DataflowPrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataflowPrivateEndpointDataSourceCrud) Get() error {
	request := oci_dataflow.GetPrivateEndpointRequest{}

	if privateEndpointId, ok := s.D.GetOkExists("private_endpoint_id"); ok {
		tmp := privateEndpointId.(string)
		request.PrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dataflow")

	response, err := s.Client.GetPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataflowPrivateEndpointDataSourceCrud) SetData() error {
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

	s.D.Set("dns_zones", s.Res.DnsZones)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaxHostCount != nil {
		s.D.Set("max_host_count", *s.Res.MaxHostCount)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.OwnerPrincipalId != nil {
		s.D.Set("owner_principal_id", *s.Res.OwnerPrincipalId)
	}

	if s.Res.OwnerUserName != nil {
		s.D.Set("owner_user_name", *s.Res.OwnerUserName)
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
