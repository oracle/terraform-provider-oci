// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiOperationsInsightsPrivateEndpointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["operations_insights_private_endpoint_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpsiOperationsInsightsPrivateEndpointResource(), fieldMap, readSingularOpsiOperationsInsightsPrivateEndpoint)
}

func readSingularOpsiOperationsInsightsPrivateEndpoint(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsPrivateEndpointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

type OpsiOperationsInsightsPrivateEndpointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetOperationsInsightsPrivateEndpointResponse
}

func (s *OpsiOperationsInsightsPrivateEndpointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiOperationsInsightsPrivateEndpointDataSourceCrud) Get() error {
	request := oci_opsi.GetOperationsInsightsPrivateEndpointRequest{}

	if operationsInsightsPrivateEndpointId, ok := s.D.GetOkExists("operations_insights_private_endpoint_id"); ok {
		tmp := operationsInsightsPrivateEndpointId.(string)
		request.OperationsInsightsPrivateEndpointId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetOperationsInsightsPrivateEndpoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiOperationsInsightsPrivateEndpointDataSourceCrud) SetData() error {
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

	if s.Res.IsUsedForRacDbs != nil {
		s.D.Set("is_used_for_rac_dbs", *s.Res.IsUsedForRacDbs)
	} else {
		s.D.Set("is_used_for_rac_dbs", "false")
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.PrivateEndpointStatusDetails != nil {
		s.D.Set("private_endpoint_status_details", *s.Res.PrivateEndpointStatusDetails)
	}

	if s.Res.PrivateIp != nil {
		s.D.Set("private_ip", *s.Res.PrivateIp)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
