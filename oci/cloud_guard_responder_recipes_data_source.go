// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v46/cloudguard"
)

func init() {
	RegisterDatasource("oci_cloud_guard_responder_recipes", CloudGuardResponderRecipesDataSource())
}

func CloudGuardResponderRecipesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudGuardResponderRecipes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
			"responder_recipe_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(CloudGuardResponderRecipeResource()),
						},
					},
				},
			},
		},
	}
}

func readCloudGuardResponderRecipes(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardResponderRecipesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).cloudGuardClient()

	return ReadResource(sync)
}

type CloudGuardResponderRecipesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListResponderRecipesResponse
}

func (s *CloudGuardResponderRecipesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardResponderRecipesDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListResponderRecipesRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_cloud_guard.ListResponderRecipesAccessLevelEnum(accessLevel.(string))
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
		request.LifecycleState = oci_cloud_guard.ListResponderRecipesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListResponderRecipes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResponderRecipes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudGuardResponderRecipesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CloudGuardResponderRecipesDataSource-", CloudGuardResponderRecipesDataSource(), s.D))
	resources := []map[string]interface{}{}
	responderRecipe := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResponderRecipeSummaryToMap(item))
	}
	responderRecipe["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, CloudGuardResponderRecipesDataSource().Schema["responder_recipe_collection"].Elem.(*schema.Resource).Schema)
		responderRecipe["items"] = items
	}

	resources = append(resources, responderRecipe)
	if err := s.D.Set("responder_recipe_collection", resources); err != nil {
		return err
	}

	return nil
}
