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

func DevopsRepoFileLineDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepoFileLine,
		Schema: map[string]*schema.Schema{
			"file_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"revision": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start_line_number": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Computed
			"lines": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"line_content": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_number": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDevopsRepoFileLine(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepoFileLineDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepoFileLineDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetRepoFileLinesResponse
}

func (s *DevopsRepoFileLineDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepoFileLineDataSourceCrud) Get() error {
	request := oci_devops.GetRepoFileLinesRequest{}

	if filePath, ok := s.D.GetOkExists("file_path"); ok {
		tmp := filePath.(string)
		request.FilePath = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if revision, ok := s.D.GetOkExists("revision"); ok {
		tmp := revision.(string)
		request.Revision = &tmp
	}

	if startLineNumber, ok := s.D.GetOkExists("start_line_number"); ok {
		tmp := startLineNumber.(int)
		request.StartLineNumber = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetRepoFileLines(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepoFileLineDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepoFileLineDataSource-", DevopsRepoFileLineDataSource(), s.D))

	lines := []interface{}{}
	for _, item := range s.Res.Lines {
		lines = append(lines, RepoFileLineDetailsToMap(item))
	}
	s.D.Set("lines", lines)

	return nil
}

func RepoFileLineDetailsToMap(obj oci_devops.FileLineDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LineContent != nil {
		result["line_content"] = string(*obj.LineContent)
	}

	if obj.LineNumber != nil {
		result["line_number"] = int(*obj.LineNumber)
	}

	return result
}
