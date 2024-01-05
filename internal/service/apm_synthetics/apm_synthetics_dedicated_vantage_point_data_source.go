// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func ApmSyntheticsDedicatedVantagePointDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["apm_domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["dedicated_vantage_point_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApmSyntheticsDedicatedVantagePointResource(), fieldMap, readSingularApmSyntheticsDedicatedVantagePoint)
}

func readSingularApmSyntheticsDedicatedVantagePoint(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsDedicatedVantagePointDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsDedicatedVantagePointDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.GetDedicatedVantagePointResponse
}

func (s *ApmSyntheticsDedicatedVantagePointDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsDedicatedVantagePointDataSourceCrud) Get() error {
	request := oci_apm_synthetics.GetDedicatedVantagePointRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if dedicatedVantagePointId, ok := s.D.GetOkExists("dedicated_vantage_point_id"); ok {
		tmp := dedicatedVantagePointId.(string)
		dedicatedVantagePointId, apmDomainId, err := parseDedicatedVantagePointCompositeId(tmp)
		if err == nil {
			request.DedicatedVantagePointId = &dedicatedVantagePointId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.GetDedicatedVantagePoint(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmSyntheticsDedicatedVantagePointDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GetDedicatedVantagePointCompositeId(*s.Res.Id, s.D.Get("apm_domain_id").(string)))

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DvpStackDetails != nil {
		dvpStackDetailsArray := []interface{}{}
		if dvpStackDetailsMap := DvpStackDetailsToMap(&s.Res.DvpStackDetails); dvpStackDetailsMap != nil {
			dvpStackDetailsArray = append(dvpStackDetailsArray, dvpStackDetailsMap)
		}
		s.D.Set("dvp_stack_details", dvpStackDetailsArray)
	} else {
		s.D.Set("dvp_stack_details", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MonitorStatusCountMap != nil {
		s.D.Set("monitor_status_count_map", []interface{}{MonitorStatusCountMapToMap(s.Res.MonitorStatusCountMap)})
	} else {
		s.D.Set("monitor_status_count_map", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	s.D.Set("status", s.Res.Status)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
