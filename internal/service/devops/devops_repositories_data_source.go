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

func DevopsRepositoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsRepositories,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_collection": {
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
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"project_id": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"repository_type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"default_branch": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"defined_tags": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"description": {
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
									"parent_repository_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"mirror_repository_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Optional
												"connector_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"repository_url": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"trigger_schedule": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"schedule_type": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional
															"custom_schedule": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},

									// Computed
									"branch_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"commit_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"http_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_details": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"project_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"size_in_bytes": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssh_url": {
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
									"trigger_build_events": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
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

func readDevopsRepositories(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListRepositoriesResponse
}

func (s *DevopsRepositoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoriesDataSourceCrud) Get() error {
	request := oci_devops.ListRepositoriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if page, ok := s.D.GetOkExists("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	if sortBy, ok := s.D.GetOkExists("sort_by"); ok {
		request.SortBy = oci_devops.ListRepositoriesSortByEnum(sortBy.(string))
	}

	if sortOrder, ok := s.D.GetOkExists("sort_order"); ok {
		request.SortOrder = oci_devops.ListRepositoriesSortOrderEnum(sortOrder.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_devops.RepositoryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.ListRepositories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRepositories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsRepositoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoriesDataSource-", DevopsRepositoriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	repository := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DevopsRepositorySummaryToMap(item))
	}
	repository["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DevopsRepositoriesDataSource().Schema["repository_collection"].Elem.(*schema.Resource).Schema)
		repository["items"] = items
	}

	resources = append(resources, repository)
	if err := s.D.Set("repository_collection", resources); err != nil {
		return err
	}

	return nil
}
