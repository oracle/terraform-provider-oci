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

func DevopsRepositoryFileDiffDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepositoryFileDiff,
		Schema: map[string]*schema.Schema{
			"base_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file_path": {
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
			"target_version": {
				Type:     schema.TypeString,
				Required: true,
			},
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
		},
	}
}

func readSingularDevopsRepositoryFileDiff(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryFileDiffDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryFileDiffDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetRepoFileDiffResponse
}

func (s *DevopsRepositoryFileDiffDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryFileDiffDataSourceCrud) Get() error {
	request := oci_devops.GetRepoFileDiffRequest{}

	if baseVersion, ok := s.D.GetOkExists("base_version"); ok {
		tmp := baseVersion.(string)
		request.BaseVersion = &tmp
	}

	if filePath, ok := s.D.GetOkExists("file_path"); ok {
		tmp := filePath.(string)
		request.FilePath = &tmp
	}

	if isComparisonFromMergeBase, ok := s.D.GetOkExists("is_comparison_from_merge_base"); ok {
		tmp := isComparisonFromMergeBase.(bool)
		request.IsComparisonFromMergeBase = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if targetVersion, ok := s.D.GetOkExists("target_version"); ok {
		tmp := targetVersion.(string)
		request.TargetVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetRepoFileDiff(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepositoryFileDiffDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryFileDiffDataSource-", DevopsRepositoryFileDiffDataSource(), s.D))

	if s.Res.AreConflictsInFile != nil {
		s.D.Set("are_conflicts_in_file", *s.Res.AreConflictsInFile)
	}

	changes := []interface{}{}
	for _, item := range s.Res.Changes {
		changes = append(changes, DiffChunkToMap(item))
	}
	s.D.Set("changes", changes)

	if s.Res.IsBinary != nil {
		s.D.Set("is_binary", *s.Res.IsBinary)
	}

	if s.Res.IsLarge != nil {
		s.D.Set("is_large", *s.Res.IsLarge)
	}

	if s.Res.NewId != nil {
		s.D.Set("new_id", *s.Res.NewId)
	}

	if s.Res.NewPath != nil {
		s.D.Set("new_path", *s.Res.NewPath)
	}

	if s.Res.OldId != nil {
		s.D.Set("old_id", *s.Res.OldId)
	}

	if s.Res.OldPath != nil {
		s.D.Set("old_path", *s.Res.OldPath)
	}

	return nil
}
