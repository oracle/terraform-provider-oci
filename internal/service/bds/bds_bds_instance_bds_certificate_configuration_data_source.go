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

func BdsBdsInstanceBdsCertificateConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bds_certificate_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["bds_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BdsBdsInstanceBdsCertificateConfigurationResource(), fieldMap, readSingularBdsBdsInstanceBdsCertificateConfiguration)
}

func readSingularBdsBdsInstanceBdsCertificateConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceBdsCertificateConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceBdsCertificateConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetBdsCertificateConfigurationResponse
}

func (s *BdsBdsInstanceBdsCertificateConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceBdsCertificateConfigurationDataSourceCrud) Get() error {
	request := oci_bds.GetBdsCertificateConfigurationRequest{}

	if bdsCertificateConfigurationId, ok := s.D.GetOkExists("bds_certificate_configuration_id"); ok {
		tmp := bdsCertificateConfigurationId.(string)
		request.BdsCertificateConfigurationId = &tmp
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetBdsCertificateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceBdsCertificateConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CertificateAuthorityId != nil {
		s.D.Set("certificate_authority_id", *s.Res.CertificateAuthorityId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsDefaultConfiguration != nil {
		s.D.Set("is_default_configuration", *s.Res.IsDefaultConfiguration)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastRefreshedOrGenerated != nil {
		s.D.Set("time_last_refreshed_or_generated", s.Res.TimeLastRefreshedOrGenerated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
