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

func DevopsRepositoryProtectedBranchesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsRepositoryProtectedBranches,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protected_branch_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"branch_name": {
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
									"protection_levels": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"system_tags": {
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

func readDevopsRepositoryProtectedBranches(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryProtectedBranchesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryProtectedBranchesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListProtectedBranchesResponse
}

func (s *DevopsRepositoryProtectedBranchesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryProtectedBranchesDataSourceCrud) Get() error {
	request := oci_devops.ListProtectedBranchesRequest{}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.ListProtectedBranches(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListProtectedBranches(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsRepositoryProtectedBranchesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryProtectedBranchesDataSource-", DevopsRepositoryProtectedBranchesDataSource(), s.D))
	resources := []map[string]interface{}{}
	repositoryProtectedBranch := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ProtectedBranchSummaryToMap(item))
	}
	repositoryProtectedBranch["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DevopsRepositoryProtectedBranchesDataSource().Schema["protected_branch_collection"].Elem.(*schema.Resource).Schema)
		repositoryProtectedBranch["items"] = items
	}

	resources = append(resources, repositoryProtectedBranch)
	if err := s.D.Set("protected_branch_collection", resources); err != nil {
		return err
	}

	return nil
}

func ProtectedBranchSummaryToMap(obj oci_devops.ProtectedBranchSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BranchName != nil {
		result["branch_name"] = string(*obj.BranchName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	result["protection_levels"] = obj.ProtectionLevels

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
