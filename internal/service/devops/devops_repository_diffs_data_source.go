// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DevopsRepositoryDiffsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsRepositoryDiffs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"base_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_comparison_from_merge_base": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_repository_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"diff_collection": {
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
									"are_conflicts_in_file": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"changes": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"base_line": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"base_span": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"diff_sections": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"lines": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"base_line": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"conflict_marker": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"line_content": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"target_line": {
																			Type:     schema.TypeInt,
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
												"target_line": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"target_span": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
									"is_binary": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_large": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"new_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"new_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"old_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"old_path": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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

func readDevopsRepositoryDiffs(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryDiffsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryDiffsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListCommitDiffsResponse
}

func (s *DevopsRepositoryDiffsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryDiffsDataSourceCrud) Get() error {
	request := oci_devops.ListCommitDiffsRequest{}

	if baseVersion, ok := s.D.GetOkExists("base_version"); ok {
		tmp := baseVersion.(string)
		request.BaseVersion = &tmp
	}

	if isComparisonFromMergeBase, ok := s.D.GetOkExists("is_comparison_from_merge_base"); ok {
		tmp := isComparisonFromMergeBase.(bool)
		request.IsComparisonFromMergeBase = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if targetRepositoryId, ok := s.D.GetOkExists("target_repository_id"); ok {
		tmp := targetRepositoryId.(string)
		request.TargetRepositoryId = &tmp
	}

	if targetVersion, ok := s.D.GetOkExists("target_version"); ok {
		tmp := targetVersion.(string)
		request.TargetVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.ListCommitDiffs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCommitDiffs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsRepositoryDiffsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryDiffsDataSource-", DevopsRepositoryDiffsDataSource(), s.D))
	resources := []map[string]interface{}{}
	repositoryDiff := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DiffSummaryToMap(item))
	}
	repositoryDiff["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DevopsRepositoryDiffsDataSource().Schema["diff_collection"].Elem.(*schema.Resource).Schema)
		repositoryDiff["items"] = items
	}

	resources = append(resources, repositoryDiff)
	if err := s.D.Set("diff_collection", resources); err != nil {
		return err
	}

	return nil
}
