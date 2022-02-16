// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package limits

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_limits "github.com/oracle/oci-go-sdk/v58/limits"

	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func LimitsQuotaDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["quota_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LimitsQuotaResource(), fieldMap, readSingularLimitsQuota)
}

func readSingularLimitsQuota(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsQuotaDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QuotasClient()

	return tfresource.ReadResource(sync)
}

type LimitsQuotaDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_limits.QuotasClient
	Res    *oci_limits.GetQuotaResponse
}

func (s *LimitsQuotaDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LimitsQuotaDataSourceCrud) Get() error {
	request := oci_limits.GetQuotaRequest{}

	if quotaId, ok := s.D.GetOkExists("quota_id"); ok {
		tmp := quotaId.(string)
		request.QuotaId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "limits")

	response, err := s.Client.GetQuota(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LimitsQuotaDataSourceCrud) SetData() error {
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("statements", s.Res.Statements)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
