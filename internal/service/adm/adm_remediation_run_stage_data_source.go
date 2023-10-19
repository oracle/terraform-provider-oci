// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AdmRemediationRunStageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAdmRemediationRunStage,
		Schema: map[string]*schema.Schema{
			"remediation_run_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stage_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"audit_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"next_stage_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pipeline_properties": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"pipeline_identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pipeline_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"previous_stage_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pull_request_properties": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"pull_request_identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"pull_request_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"recommended_updates_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"summary": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularAdmRemediationRunStage(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRunStageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.ReadResource(sync)
}

type AdmRemediationRunStageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_adm.ApplicationDependencyManagementClient
	Res    *oci_adm.GetStageResponse
}

func (s *AdmRemediationRunStageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AdmRemediationRunStageDataSourceCrud) Get() error {
	request := oci_adm.GetStageRequest{}

	if remediationRunId, ok := s.D.GetOkExists("remediation_run_id"); ok {
		tmp := remediationRunId.(string)
		request.RemediationRunId = &tmp
	}

	if stageType, ok := s.D.GetOkExists("stage_type"); ok {
		enum, _ := oci_adm.GetMappingGetStageStageTypeEnum(stageType.(string))
		request.StageType = enum
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "adm")

	response, err := s.Client.GetStage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AdmRemediationRunStageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AdmRemediationRunStageDataSource-", AdmRemediationRunStageDataSource(), s.D))
	switch v := (s.Res.RemediationRunStage).(type) {
	case oci_adm.ApplyStage:
		s.D.Set("type", "APPLY")

		if v.PipelineProperties != nil {
			s.D.Set("pipeline_properties", []interface{}{PipelinePropertiesToMap(v.PipelineProperties)})
		} else {
			s.D.Set("pipeline_properties", nil)
		}

		if v.PullRequestProperties != nil {
			s.D.Set("pull_request_properties", []interface{}{PullRequestPropertiesToMap(v.PullRequestProperties)})
		} else {
			s.D.Set("pull_request_properties", nil)
		}

		s.D.Set("next_stage_type", v.NextStageType)

		s.D.Set("previous_stage_type", v.PreviousStageType)

		if v.RemediationRunId != nil {
			s.D.Set("remediation_run_id", *v.RemediationRunId)
		}

		s.D.Set("status", v.Status)

		if v.Summary != nil {
			s.D.Set("summary", *v.Summary)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeFinished != nil {
			s.D.Set("time_finished", v.TimeFinished.String())
		}

		if v.TimeStarted != nil {
			s.D.Set("time_started", v.TimeStarted.String())
		}
	case oci_adm.DetectStage:
		s.D.Set("type", "DETECT")

		if v.AuditId != nil {
			s.D.Set("audit_id", *v.AuditId)
		}

		s.D.Set("next_stage_type", v.NextStageType)

		s.D.Set("previous_stage_type", v.PreviousStageType)

		if v.RemediationRunId != nil {
			s.D.Set("remediation_run_id", *v.RemediationRunId)
		}

		s.D.Set("status", v.Status)

		if v.Summary != nil {
			s.D.Set("summary", *v.Summary)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeFinished != nil {
			s.D.Set("time_finished", v.TimeFinished.String())
		}

		if v.TimeStarted != nil {
			s.D.Set("time_started", v.TimeStarted.String())
		}
	case oci_adm.RecommendStage:
		s.D.Set("type", "RECOMMEND")

		if v.PullRequestProperties != nil {
			s.D.Set("pull_request_properties", []interface{}{PullRequestPropertiesToMap(v.PullRequestProperties)})
		} else {
			s.D.Set("pull_request_properties", nil)
		}

		if v.RecommendedUpdatesCount != nil {
			s.D.Set("recommended_updates_count", *v.RecommendedUpdatesCount)
		}

		s.D.Set("next_stage_type", v.NextStageType)

		s.D.Set("previous_stage_type", v.PreviousStageType)

		if v.RemediationRunId != nil {
			s.D.Set("remediation_run_id", *v.RemediationRunId)
		}

		s.D.Set("status", v.Status)

		if v.Summary != nil {
			s.D.Set("summary", *v.Summary)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeFinished != nil {
			s.D.Set("time_finished", v.TimeFinished.String())
		}

		if v.TimeStarted != nil {
			s.D.Set("time_started", v.TimeStarted.String())
		}
	case oci_adm.VerifyStage:
		s.D.Set("type", "VERIFY")

		if v.PipelineProperties != nil {
			s.D.Set("pipeline_properties", []interface{}{PipelinePropertiesToMap(v.PipelineProperties)})
		} else {
			s.D.Set("pipeline_properties", nil)
		}

		if v.PullRequestProperties != nil {
			s.D.Set("pull_request_properties", []interface{}{PullRequestPropertiesToMap(v.PullRequestProperties)})
		} else {
			s.D.Set("pull_request_properties", nil)
		}

		s.D.Set("next_stage_type", v.NextStageType)

		s.D.Set("previous_stage_type", v.PreviousStageType)

		if v.RemediationRunId != nil {
			s.D.Set("remediation_run_id", *v.RemediationRunId)
		}

		s.D.Set("status", v.Status)

		if v.Summary != nil {
			s.D.Set("summary", *v.Summary)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeFinished != nil {
			s.D.Set("time_finished", v.TimeFinished.String())
		}

		if v.TimeStarted != nil {
			s.D.Set("time_started", v.TimeStarted.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.RemediationRunStage)
		return nil
	}

	return nil
}
