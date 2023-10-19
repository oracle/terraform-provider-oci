// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package adm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AdmRemediationRunStagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readAdmRemediationRunStages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"remediation_run_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remediation_run_stage_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

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
									"remediation_run_id": {
										Type:     schema.TypeString,
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
							},
						},
					},
				},
			},
		},
	}
}

func readAdmRemediationRunStages(d *schema.ResourceData, m interface{}) error {
	sync := &AdmRemediationRunStagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApplicationDependencyManagementClient()

	return tfresource.ReadResource(sync)
}

type AdmRemediationRunStagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_adm.ApplicationDependencyManagementClient
	Res    *oci_adm.ListStagesResponse
}

func (s *AdmRemediationRunStagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AdmRemediationRunStagesDataSourceCrud) Get() error {
	request := oci_adm.ListStagesRequest{}

	if remediationRunId, ok := s.D.GetOkExists("remediation_run_id"); ok {
		tmp := remediationRunId.(string)
		request.RemediationRunId = &tmp
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_adm.RemediationRunStageStatusEnum(status.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_adm.ListStagesTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "adm")

	response, err := s.Client.ListStages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListStages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *AdmRemediationRunStagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("AdmRemediationRunStagesDataSource-", AdmRemediationRunStagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	remediationRunStage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RemediationRunStageSummaryToMap(item))
	}
	remediationRunStage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, AdmRemediationRunStagesDataSource().Schema["remediation_run_stage_collection"].Elem.(*schema.Resource).Schema)
		remediationRunStage["items"] = items
	}

	resources = append(resources, remediationRunStage)
	if err := s.D.Set("remediation_run_stage_collection", resources); err != nil {
		return err
	}

	return nil
}

func PipelinePropertiesToMap(obj *oci_adm.PipelineProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PipelineIdentifier != nil {
		result["pipeline_identifier"] = string(*obj.PipelineIdentifier)
	}

	if obj.PipelineUrl != nil {
		result["pipeline_url"] = string(*obj.PipelineUrl)
	}

	return result
}

func PullRequestPropertiesToMap(obj *oci_adm.PullRequestProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PullRequestIdentifier != nil {
		result["pull_request_identifier"] = string(*obj.PullRequestIdentifier)
	}

	if obj.PullRequestUrl != nil {
		result["pull_request_url"] = string(*obj.PullRequestUrl)
	}

	return result
}

func RemediationRunStageSummaryToMap(obj oci_adm.RemediationRunStageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RemediationRunId != nil {
		result["remediation_run_id"] = string(*obj.RemediationRunId)
	}

	result["status"] = string(obj.Status)

	if obj.Summary != nil {
		result["summary"] = string(*obj.Summary)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	result["type"] = string(obj.Type)

	return result
}
