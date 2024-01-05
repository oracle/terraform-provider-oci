// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsBuildRunsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsBuildRuns,
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
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"build_run_summary_collection": {
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
									"build_pipeline_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"build_run_arguments": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"items": {
													Type:     schema.TypeList,
													Required: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"name": {
																Type:     schema.TypeString,
																Required: true,
																ForceNew: true,
															},
															"value": {
																Type:     schema.TypeString,
																Required: true,
																ForceNew: true,
															},

															// Optional

															// Computed
														},
													},
												},

												// Optional

												// Computed
											},
										},
									},
									"commit_info": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"commit_hash": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"repository_branch": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"repository_url": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},

												// Optional

												// Computed
											},
										},
									},
									"defined_tags": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},

									// Computed
									"build_run_progress_summary": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"time_finished": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_started": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"build_run_source": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"repository_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"source_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"trigger_info": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"actions": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"build_pipeline_id": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"filter": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"events": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																					"include": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								// Required

																								// Optional

																								// Computed
																								"base_ref": {
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																								"head_ref": {
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},
																					"trigger_source": {
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																		"type": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"display_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"trigger_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
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
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readDevopsBuildRuns(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildRunsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsBuildRunsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListBuildRunsResponse
}

func (s *DevopsBuildRunsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsBuildRunsDataSourceCrud) Get() error {
	request := oci_devops.ListBuildRunsRequest{}

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

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_devops.BuildRunLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.ListBuildRuns(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBuildRuns(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsBuildRunsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsBuildRunsDataSource-", DevopsBuildRunsDataSource(), s.D))
	resources := []map[string]interface{}{}
	buildRun := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BuildRunSummaryToMap(item))
	}
	buildRun["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DevopsBuildRunsDataSource().Schema["build_run_summary_collection"].Elem.(*schema.Resource).Schema)
		buildRun["items"] = items
	}

	resources = append(resources, buildRun)
	if err := s.D.Set("build_run_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
