// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DevopsRepositoryObjectDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepositoryObject,
		Schema: map[string]*schema.Schema{
			"file_path": {
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
			"is_binary": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_bytes": {
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

func readSingularDevopsRepositoryObject(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryObjectDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryObjectDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetObjectResponse
}

func (s *DevopsRepositoryObjectDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryObjectDataSourceCrud) Get() error {
	request := oci_devops.GetObjectRequest{}

	if filePath, ok := s.D.GetOkExists("file_path"); ok {
		tmp := filePath.(string)
		request.FilePath = &tmp
	}

	if refName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := refName.(string)
		request.RefName = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetObject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepositoryObjectDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryObjectDataSource-", DevopsRepositoryObjectDataSource(), s.D))

	if s.Res.IsBinary != nil {
		s.D.Set("is_binary", *s.Res.IsBinary)
	}

	if s.Res.Sha != nil {
		s.D.Set("sha", *s.Res.Sha)
	}

	if s.Res.SizeInBytes != nil {
		s.D.Set("size_in_bytes", strconv.FormatInt(*s.Res.SizeInBytes, 10))
	}

	s.D.Set("type", s.Res.Type)

	return nil
}
