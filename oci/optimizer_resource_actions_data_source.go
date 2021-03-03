// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_optimizer "github.com/oracle/oci-go-sdk/v36/optimizer"
)

func init() {
	RegisterDatasource("oci_optimizer_resource_actions", OptimizerResourceActionsDataSource())
}

func OptimizerResourceActionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOptimizerResourceActions,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommendation_id": {
				Type:     schema.TypeString,
				Required: true,
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
							Elem:     GetDataSourceItemSchema(OptimizerResourceActionResource()),
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
	sync.Client = m.(*OracleClients).optimizerClient()

	return ReadResource(sync)
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

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if recommendationId, ok := s.D.GetOkExists("recommendation_id"); ok {
		tmp := recommendationId.(string)
		request.RecommendationId = &tmp
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "optimizer")

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

	s.D.SetId(GenerateDataSourceHashID("OptimizerResourceActionsDataSource-", OptimizerResourceActionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	resourceAction := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ResourceActionSummaryToMap(item))
	}
	resourceAction["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, OptimizerResourceActionsDataSource().Schema["resource_action_collection"].Elem.(*schema.Resource).Schema)
		resourceAction["items"] = items
	}

	resources = append(resources, resourceAction)
	if err := s.D.Set("resource_action_collection", resources); err != nil {
		return err
	}

	return nil
}
