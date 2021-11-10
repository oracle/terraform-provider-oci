// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v51/devops"
)

func init() {
	RegisterDatasource("oci_devops_repository_archive_content", DevopsRepositoryArchiveContentDataSource())
}

func DevopsRepositoryArchiveContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepositoryArchiveContent,
		Schema: map[string]*schema.Schema{
			"format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularDevopsRepositoryArchiveContent(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryArchiveContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).devopsClient()

	return ReadResource(sync)
}

type DevopsRepositoryArchiveContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetRepositoryArchiveContentResponse
}

func (s *DevopsRepositoryArchiveContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryArchiveContentDataSourceCrud) Get() error {
	request := oci_devops.GetRepositoryArchiveContentRequest{}

	if format, ok := s.D.GetOkExists("format"); ok {
		tmp := format.(string)
		request.Format = &tmp
	}

	if refName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := refName.(string)
		request.RefName = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "devops")

	response, err := s.Client.GetRepositoryArchiveContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepositoryArchiveContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DevopsRepositoryArchiveContentDataSource-", DevopsRepositoryArchiveContentDataSource(), s.D))

	return nil
}
