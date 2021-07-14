// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm "github.com/oracle/oci-go-sdk/v44/apmcontrolplane"
)

func init() {
	RegisterDatasource("oci_apm_apm_domain", ApmApmDomainDataSource())
}

func ApmApmDomainDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["apm_domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(ApmApmDomainResource(), fieldMap, readSingularApmApmDomain)
}

func readSingularApmApmDomain(d *schema.ResourceData, m interface{}) error {
	sync := &ApmApmDomainDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).apmDomainClient()

	return ReadResource(sync)
}

type ApmApmDomainDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm.ApmDomainClient
	Res    *oci_apm.GetApmDomainResponse
}

func (s *ApmApmDomainDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmApmDomainDataSourceCrud) Get() error {
	request := oci_apm.GetApmDomainRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "apm")

	response, err := s.Client.GetApmDomain(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmApmDomainDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataUploadEndpoint != nil {
		s.D.Set("data_upload_endpoint", *s.Res.DataUploadEndpoint)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsFreeTier != nil {
		s.D.Set("is_free_tier", *s.Res.IsFreeTier)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
