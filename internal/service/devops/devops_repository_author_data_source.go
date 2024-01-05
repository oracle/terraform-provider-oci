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

func DevopsRepositoryAuthorDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepositoryAuthor,
		Schema: map[string]*schema.Schema{
			"ref_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"author_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_devops_repository_author", "oci_devops_repository_authors"),
	}
}

func readSingularDevopsRepositoryAuthor(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryAuthorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryAuthorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListAuthorsResponse
}

func (s *DevopsRepositoryAuthorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryAuthorDataSourceCrud) Get() error {
	request := oci_devops.ListAuthorsRequest{}

	if refName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := refName.(string)
		request.RefName = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.ListAuthors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepositoryAuthorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryAuthorDataSource-", DevopsRepositoryAuthorDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RepositoryAuthorSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func RepositoryAuthorSummaryToMap(obj oci_devops.RepositoryAuthorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuthorName != nil {
		result["author_name"] = string(*obj.AuthorName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	return result
}
