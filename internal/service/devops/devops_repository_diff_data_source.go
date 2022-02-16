// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v58/devops"
)

func DevopsRepositoryDiffDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepositoryDiff,
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

func readSingularDevopsRepositoryDiff(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryDiffDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryDiffDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetFileDiffResponse
}

func (s *DevopsRepositoryDiffDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryDiffDataSourceCrud) Get() error {
	request := oci_devops.GetFileDiffRequest{}

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

	response, err := s.Client.GetFileDiff(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepositoryDiffDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryDiffDataSource-", DevopsRepositoryDiffDataSource(), s.D))

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

func DiffChunkToMap(obj oci_devops.DiffChunk) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseLine != nil {
		result["base_line"] = int(*obj.BaseLine)
	}

	if obj.BaseSpan != nil {
		result["base_span"] = int(*obj.BaseSpan)
	}

	diffSections := []interface{}{}
	for _, item := range obj.DiffSections {
		diffSections = append(diffSections, DiffSectionToMap(item))
	}
	result["diff_sections"] = diffSections

	if obj.TargetLine != nil {
		result["target_line"] = int(*obj.TargetLine)
	}

	if obj.TargetSpan != nil {
		result["target_span"] = int(*obj.TargetSpan)
	}

	return result
}

func DiffLineDetailsToMap(obj oci_devops.DiffLineDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseLine != nil {
		result["base_line"] = int(*obj.BaseLine)
	}

	result["conflict_marker"] = string(obj.ConflictMarker)

	if obj.LineContent != nil {
		result["line_content"] = string(*obj.LineContent)
	}

	if obj.TargetLine != nil {
		result["target_line"] = int(*obj.TargetLine)
	}

	return result
}

func DiffSectionToMap(obj oci_devops.DiffSection) map[string]interface{} {
	result := map[string]interface{}{}

	lines := []interface{}{}
	for _, item := range obj.Lines {
		lines = append(lines, DiffLineDetailsToMap(item))
	}
	result["lines"] = lines

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func DiffSummaryToMap(obj oci_devops.DiffSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AreConflictsInFile != nil {
		result["are_conflicts_in_file"] = bool(*obj.AreConflictsInFile)
	}

	changes := []interface{}{}
	for _, item := range obj.Changes {
		changes = append(changes, DiffChunkToMap(item))
	}
	result["changes"] = changes

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.IsBinary != nil {
		result["is_binary"] = bool(*obj.IsBinary)
	}

	if obj.IsLarge != nil {
		result["is_large"] = bool(*obj.IsLarge)
	}

	if obj.NewId != nil {
		result["new_id"] = string(*obj.NewId)
	}

	if obj.NewPath != nil {
		result["new_path"] = string(*obj.NewPath)
	}

	if obj.OldId != nil {
		result["old_id"] = string(*obj.OldId)
	}

	if obj.OldPath != nil {
		result["old_path"] = string(*obj.OldPath)
	}

	return result
}
