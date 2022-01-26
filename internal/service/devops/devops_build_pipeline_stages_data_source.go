// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"
)

func DevopsBuildPipelineStagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsBuildPipelineStages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"build_pipeline_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"build_pipeline_stage_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"project_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"connection_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"build_pipeline_stage_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"deploy_pipeline_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_pass_all_parameters_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"build_pipeline_stage_predecessor_collection": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"items": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"build_pipeline_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"build_source_collection": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"items": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"connection_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"branch": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"connection_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"repository_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"repository_url": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"build_spec_file": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"deliver_artifact_collection": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"items": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"artifact_id": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"artifact_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"image": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"primary_build_source": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"stage_execution_timeout_in_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"wait_criteria": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"wait_duration": {
													Type:     schema.TypeString,
													Required: true,
												},
												"wait_type": {
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
				},
			},
		},
	}
}

func readDevopsBuildPipelineStages(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildPipelineStagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsBuildPipelineStagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListBuildPipelineStagesResponse
}

func (s *DevopsBuildPipelineStagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsBuildPipelineStagesDataSourceCrud) Get() error {
	request := oci_devops.ListBuildPipelineStagesRequest{}

	if buildPipelineId, ok := s.D.GetOkExists("build_pipeline_id"); ok {
		tmp := buildPipelineId.(string)
		request.BuildPipelineId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_devops.BuildPipelineStageLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.ListBuildPipelineStages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBuildPipelineStages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsBuildPipelineStagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsBuildPipelineStagesDataSource-", DevopsBuildPipelineStagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	buildPipelineStage := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BuildPipelineStageSummaryToMap(item))
	}
	buildPipelineStage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DevopsBuildPipelineStagesDataSource().Schema["build_pipeline_stage_collection"].Elem.(*schema.Resource).Schema)
		buildPipelineStage["items"] = items
	}

	resources = append(resources, buildPipelineStage)
	if err := s.D.Set("build_pipeline_stage_collection", resources); err != nil {
		return err
	}

	return nil
}
