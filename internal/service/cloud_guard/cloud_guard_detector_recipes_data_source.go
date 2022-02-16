// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v58/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CloudGuardDetectorRecipesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardDetectorRecipes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_metadata_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"detector_recipe_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(CloudGuardDetectorRecipeResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudGuardDetectorRecipes(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDetectorRecipesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardDetectorRecipesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListDetectorRecipesResponse
}

func (s *CloudGuardDetectorRecipesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardDetectorRecipesDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListDetectorRecipesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_cloud_guard.ListDetectorRecipesAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if resourceMetadataOnly, ok := s.D.GetOkExists("resource_metadata_only"); ok {
		tmp := resourceMetadataOnly.(bool)
		request.ResourceMetadataOnly = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_guard.ListDetectorRecipesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListDetectorRecipes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDetectorRecipes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardDetectorRecipesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardDetectorRecipesDataSource-", CloudGuardDetectorRecipesDataSource(), s.D))
	resources := []map[string]interface{}{}
	detectorRecipe := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DetectorRecipeSummaryToMap(item))
	}
	detectorRecipe["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardDetectorRecipesDataSource().Schema["detector_recipe_collection"].Elem.(*schema.Resource).Schema)
		detectorRecipe["items"] = items
	}

	resources = append(resources, detectorRecipe)
	if err := s.D.Set("detector_recipe_collection", resources); err != nil {
		return err
	}

	return nil
}
