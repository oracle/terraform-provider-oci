// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v65/optimizer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OptimizerResourceActionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOptimizerResourceActions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"child_tenancy_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"include_organization": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"include_resource_metadata": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommendation_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommendation_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_action_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OptimizerResourceActionResource()),
						},
					},
				},
			},
		},
	}
}

func readOptimizerResourceActions(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerResourceActionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

type OptimizerResourceActionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_optimizer.OptimizerClient
	Res    *oci_optimizer.ListResourceActionsResponse
}

func (s *OptimizerResourceActionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OptimizerResourceActionsDataSourceCrud) Get() error {
	request := oci_optimizer.ListResourceActionsRequest{}

	if childTenancyIds, ok := s.D.GetOkExists("child_tenancy_ids"); ok {
		interfaces := childTenancyIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("child_tenancy_ids") {
			request.ChildTenancyIds = tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if includeOrganization, ok := s.D.GetOkExists("include_organization"); ok {
		tmp := includeOrganization.(bool)
		request.IncludeOrganization = &tmp
	}

	if includeResourceMetadata, ok := s.D.GetOkExists("include_resource_metadata"); ok {
		tmp := includeResourceMetadata.(bool)
		request.IncludeResourceMetadata = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if recommendationId, ok := s.D.GetOkExists("recommendation_id"); ok {
		tmp := recommendationId.(string)
		request.RecommendationId = &tmp
	}

	if recommendationName, ok := s.D.GetOkExists("recommendation_name"); ok {
		tmp := recommendationName.(string)
		request.RecommendationName = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_optimizer.ListResourceActionsLifecycleStateEnum(state.(string))
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_optimizer.ListResourceActionsStatusEnum(status.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "optimizer")

	response, err := s.Client.ListResourceActions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListResourceActions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OptimizerResourceActionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OptimizerResourceActionsDataSource-", OptimizerResourceActionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	resourceAction := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResourceActionSummaryToMap(item))
	}
	resourceAction["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OptimizerResourceActionsDataSource().Schema["resource_action_collection"].Elem.(*schema.Resource).Schema)
		resourceAction["items"] = items
	}

	resources = append(resources, resourceAction)
	if err := s.D.Set("resource_action_collection", resources); err != nil {
		return err
	}

	return nil
}
