// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsOnPremiseVantagePointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["apm_domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["on_premise_vantage_point_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApmSyntheticsOnPremiseVantagePointResource(), fieldMap, readSingularApmSyntheticsOnPremiseVantagePoint)
}

func readSingularApmSyntheticsOnPremiseVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsOnPremiseVantagePointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsOnPremiseVantagePointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.GetOnPremiseVantagePointResponse
}

func (s *ApmSyntheticsOnPremiseVantagePointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsOnPremiseVantagePointDataSourceCrud) Get() error {
	request := oci_apm_synthetics.GetOnPremiseVantagePointRequest{}

	if onPremiseVantagePointCompositeId, ok := s.D.GetOkExists("on_premise_vantage_point_id"); ok {
		tmp := onPremiseVantagePointCompositeId.(string)
		onPremiseVantagePointId, apmDomainId, err := parseOnPremiseVantagePointCompositeId(tmp)
		if err == nil {
			request.OnPremiseVantagePointId = &onPremiseVantagePointId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse onPremiseVantagePointCompositeId: %s", tmp)
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.GetOnPremiseVantagePoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmSyntheticsOnPremiseVantagePointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)
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

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.WorkersSummary != nil {
		s.D.Set("workers_summary", []interface{}{WorkersSummaryToMap(s.Res.WorkersSummary)})
	} else {
		s.D.Set("workers_summary", nil)
	}

	return nil
}
