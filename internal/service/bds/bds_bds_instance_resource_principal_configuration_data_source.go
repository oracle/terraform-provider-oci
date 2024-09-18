// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceResourcePrincipalConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bds_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["resource_principal_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BdsBdsInstanceResourcePrincipalConfigurationResource(), fieldMap, readSingularBdsBdsInstanceResourcePrincipalConfiguration)
}

func readSingularBdsBdsInstanceResourcePrincipalConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceResourcePrincipalConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceResourcePrincipalConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetResourcePrincipalConfigurationResponse
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationDataSourceCrud) Get() error {
	request := oci_bds.GetResourcePrincipalConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if resourcePrincipalConfigurationId, ok := s.D.GetOkExists("resource_principal_configuration_id"); ok {
		tmp := resourcePrincipalConfigurationId.(string)
		request.ResourcePrincipalConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetResourcePrincipalConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceResourcePrincipalConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.SessionTokenLifeSpanDurationInHours != nil {
		s.D.Set("session_token_life_span_duration_in_hours", *s.Res.SessionTokenLifeSpanDurationInHours)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeTokenExpiry != nil {
		s.D.Set("time_token_expiry", s.Res.TimeTokenExpiry.String())
	}

	if s.Res.TimeTokenRefreshed != nil {
		s.D.Set("time_token_refreshed", s.Res.TimeTokenRefreshed.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
