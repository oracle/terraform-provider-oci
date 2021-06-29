// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v43/opsi"
)

func init() {
	RegisterDatasource("oci_opsi_host_insight", OpsiHostInsightDataSource())
}

func OpsiHostInsightDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["host_insight_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(OpsiHostInsightResource(), fieldMap, readSingularOpsiHostInsight)
}

func readSingularOpsiHostInsight(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiHostInsightDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).operationsInsightsClient()

	return ReadResource(sync)
}

type OpsiHostInsightDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.HostInsight
}

func (s *OpsiHostInsightDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiHostInsightDataSourceCrud) Get() error {
	request := oci_opsi.GetHostInsightRequest{}

	if hostInsightId, ok := s.D.GetOkExists("host_insight_id"); ok {
		tmp := hostInsightId.(string)
		request.HostInsightId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "opsi")

	response, err := s.Client.GetHostInsight(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.HostInsight
	return nil
}

func (s *OpsiHostInsightDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	switch v := (*s.Res).(type) {
	case oci_opsi.MacsManagedExternalHostInsight:
		s.D.Set("entity_source", "MACS_MANAGED_EXTERNAL_HOST")

		s.D.SetId(*v.GetId())

		if v.GetCompartmentId() != nil {
			s.D.Set("compartment_id", *v.GetCompartmentId())
		}

		if v.GetDefinedTags() != nil {
			s.D.Set("defined_tags", definedTagsToMap(v.GetDefinedTags()))
		}

		s.D.Set("freeform_tags", v.GetFreeformTags())

		if v.GetHostDisplayName() != nil {
			s.D.Set("host_display_name", *v.GetHostDisplayName())
		}

		if v.GetHostName() != nil {
			s.D.Set("host_name", *v.GetHostName())
		}

		if v.GetHostType() != nil {
			s.D.Set("host_type", *v.GetHostType())
		}

		if v.GetLifecycleDetails() != nil {
			s.D.Set("lifecycle_details", *v.GetLifecycleDetails())
		}

		if v.GetProcessorCount() != nil {
			s.D.Set("processor_count", *v.GetProcessorCount())
		}

		s.D.Set("state", v.GetLifecycleState())

		s.D.Set("status", v.GetStatus())

		if v.GetSystemTags() != nil {
			s.D.Set("system_tags", systemTagsToMap(v.GetSystemTags()))
		}

		if v.GetTimeCreated() != nil {
			s.D.Set("time_created", v.GetTimeCreated().String())
		}

		if v.GetTimeUpdated() != nil {
			s.D.Set("time_updated", v.GetTimeUpdated().String())
		}

	default:
		log.Printf("[WARN] Received 'entity_source' of unknown type %v", *s.Res)
		return nil
	}

	return nil
}
