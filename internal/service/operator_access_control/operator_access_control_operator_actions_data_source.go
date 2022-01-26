// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v56/operatoraccesscontrol"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OperatorAccessControlOperatorActionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOperatorAccessControlOperatorActions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"operator_action_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"component": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"customer_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"compartment_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"properties": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
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

func readOperatorAccessControlOperatorActions(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorActionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorActionsClient()

	return tfresource.ReadResource(sync)
}

type OperatorAccessControlOperatorActionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_operator_access_control.OperatorActionsClient
	Res    *oci_operator_access_control.ListOperatorActionsResponse
}

func (s *OperatorAccessControlOperatorActionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OperatorAccessControlOperatorActionsDataSourceCrud) Get() error {
	request := oci_operator_access_control.ListOperatorActionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := string(oci_operator_access_control.ResourceTypesEnum(resourceType.(string)))
		request.ResourceType = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_operator_access_control.ListOperatorActionsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "operator_access_control")

	response, err := s.Client.ListOperatorActions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	return nil
}

func (s *OperatorAccessControlOperatorActionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OperatorAccessControlOperatorActionsDataSource-", OperatorAccessControlOperatorActionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	operatorAction := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OperatorActionSummaryToMap(item))
	}
	operatorAction["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OperatorAccessControlOperatorActionsDataSource().Schema["operator_action_collection"].Elem.(*schema.Resource).Schema)
		operatorAction["items"] = items
	}

	resources = append(resources, operatorAction)
	if err := s.D.Set("operator_action_collection", resources); err != nil {
		return err
	}

	return nil
}

func OperatorActionsPropertiesToMap(obj oci_operator_access_control.OperatorActionProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func OperatorActionsSummaryToMap(obj oci_operator_access_control.OperatorActionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Component != nil {
		result["component"] = string(*obj.Component)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["resource_type"] = string(obj.ResourceType)

	result["state"] = string(obj.LifecycleState)

	return result
}
