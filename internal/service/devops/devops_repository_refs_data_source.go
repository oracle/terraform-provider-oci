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

func DevopsRepositoryRefsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsRepositoryRefs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"commit_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository_ref_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DevopsRepositoryRefResource()),
						},
					},
				},
			},
		},
	}
}

func readDevopsRepositoryRefs(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryRefsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryRefsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListRefsResponse
}

func (s *DevopsRepositoryRefsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryRefsDataSourceCrud) Get() error {
	request := oci_devops.ListRefsRequest{}

	if commitId, ok := s.D.GetOkExists("commit_id"); ok {
		tmp := commitId.(string)
		request.CommitId = &tmp
	}

	if refName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := refName.(string)
		request.RefName = &tmp
	}

	if refType, ok := s.D.GetOkExists("ref_type"); ok {
		request.RefType = oci_devops.ListRefsRefTypeEnum(refType.(string))
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.ListRefs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRefs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsRepositoryRefsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryRefsDataSource-", DevopsRepositoryRefsDataSource(), s.D))
	resources := []map[string]interface{}{}
	repositoryRef := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RepositoryRefSummaryToMap(item))
	}
	repositoryRef["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DevopsRepositoryRefsDataSource().Schema["repository_ref_collection"].Elem.(*schema.Resource).Schema)
		repositoryRef["items"] = items
	}

	resources = append(resources, repositoryRef)
	if err := s.D.Set("repository_ref_collection", resources); err != nil {
		return err
	}

	return nil
}
