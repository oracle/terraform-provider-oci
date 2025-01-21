// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreByoasnDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["byoasn_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(CoreByoasnResource(), fieldMap, readSingularCoreByoasn)
}

func readSingularCoreByoasn(d *schema.ResourceData, m interface{}) error {
	sync := &CoreByoasnDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreByoasnDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetByoasnResponse
}

func (s *CoreByoasnDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreByoasnDataSourceCrud) Get() error {
	request := oci_core.GetByoasnRequest{}

	if byoasnId, ok := s.D.GetOkExists("byoasn_id"); ok {
		tmp := byoasnId.(string)
		request.ByoasnId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetByoasn(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreByoasnDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Asn != nil {
		s.D.Set("asn", strconv.FormatInt(*s.Res.Asn, 10))
	}

	byoipRanges := []interface{}{}
	for _, item := range s.Res.ByoipRanges {
		byoipRanges = append(byoipRanges, ByoasnByoipRangeToMap(item))
	}
	s.D.Set("byoip_ranges", byoipRanges)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeValidated != nil {
		s.D.Set("time_validated", s.Res.TimeValidated.String())
	}

	if s.Res.ValidationToken != nil {
		s.D.Set("validation_token", *s.Res.ValidationToken)
	}

	return nil
}
