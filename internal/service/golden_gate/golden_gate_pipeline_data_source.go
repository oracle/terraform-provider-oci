// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGatePipelineDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["pipeline_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GoldenGatePipelineResource(), fieldMap, readSingularGoldenGatePipeline)
}

func readSingularGoldenGatePipeline(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGatePipelineDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGatePipelineDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.GetPipelineResponse
}

func (s *GoldenGatePipelineDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGatePipelineDataSourceCrud) Get() error {
	request := oci_golden_gate.GetPipelineRequest{}

	if pipelineId, ok := s.D.GetOkExists("pipeline_id"); ok {
		tmp := pipelineId.(string)
		request.PipelineId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.GetPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGatePipelineDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.Pipeline).(type) {
	case oci_golden_gate.ZeroEtlPipeline:
		s.D.Set("recipe_type", "ZERO_ETL")

		mappingRules := []interface{}{}
		for _, item := range v.MappingRules {
			mappingRules = append(mappingRules, MappingRuleToMap(item))
		}
		s.D.Set("mapping_rules", mappingRules)

		if v.ProcessOptions != nil {
			s.D.Set("process_options", []interface{}{ProcessOptionsToMap(v.ProcessOptions)})
		} else {
			s.D.Set("process_options", nil)
		}

		if v.TimeLastRecorded != nil {
			s.D.Set("time_last_recorded", v.TimeLastRecorded.Format(time.RFC3339Nano))
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.CpuCoreCount != nil {
			s.D.Set("cpu_core_count", *v.CpuCoreCount)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsAutoScalingEnabled != nil {
			s.D.Set("is_auto_scaling_enabled", *v.IsAutoScalingEnabled)
		}

		s.D.Set("license_model", v.LicenseModel)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("lifecycle_sub_state", v.LifecycleSubState)

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		if v.PipelineDiagnosticData != nil {
			s.D.Set("pipeline_diagnostic_data", []interface{}{PipelineDiagnosticDataToMap(v.PipelineDiagnosticData)})
		} else {
			s.D.Set("pipeline_diagnostic_data", nil)
		}

		if v.SourceConnectionDetails != nil {
			s.D.Set("source_connection_details", []interface{}{SourcePipelineConnectionDetailsToMap(v.SourceConnectionDetails)})
		} else {
			s.D.Set("source_connection_details", nil)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TargetConnectionDetails != nil {
			s.D.Set("target_connection_details", []interface{}{TargetPipelineConnectionDetailsToMap(v.TargetConnectionDetails)})
		} else {
			s.D.Set("target_connection_details", nil)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'recipe_type' of unknown type %v", s.Res.Pipeline)
		return nil
	}

	return nil
}
