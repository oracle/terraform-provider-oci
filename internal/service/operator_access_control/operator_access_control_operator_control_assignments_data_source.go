// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package operator_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v58/operatoraccesscontrol"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func OperatorAccessControlOperatorControlAssignmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOperatorAccessControlOperatorControlAssignments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"operator_control_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_name": {
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
			"operator_control_assignment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OperatorAccessControlOperatorControlAssignmentResource()),
						},
					},
				},
			},
		},
	}
}

func readOperatorAccessControlOperatorControlAssignments(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlAssignmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlAssignmentClient()

	return tfresource.ReadResource(sync)
}

type OperatorAccessControlOperatorControlAssignmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_operator_access_control.OperatorControlAssignmentClient
	Res    *oci_operator_access_control.ListOperatorControlAssignmentsResponse
}

func (s *OperatorAccessControlOperatorControlAssignmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OperatorAccessControlOperatorControlAssignmentsDataSourceCrud) Get() error {
	request := oci_operator_access_control.ListOperatorControlAssignmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if operatorControlName, ok := s.D.GetOkExists("operator_control_name"); ok {
		tmp := operatorControlName.(string)
		request.OperatorControlName = &tmp
	}

	if resourceName, ok := s.D.GetOkExists("resource_name"); ok {
		tmp := resourceName.(string)
		request.ResourceName = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := string(oci_operator_access_control.ResourceTypesEnum(resourceType.(string)))
		request.ResourceType = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_operator_access_control.ListOperatorControlAssignmentsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "operator_access_control")

	response, err := s.Client.ListOperatorControlAssignments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOperatorControlAssignments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OperatorAccessControlOperatorControlAssignmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OperatorAccessControlOperatorControlAssignmentsDataSource-", OperatorAccessControlOperatorControlAssignmentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	operatorControlAssignment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OperatorControlAssignmentSummaryToMap(item))
	}
	operatorControlAssignment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OperatorAccessControlOperatorControlAssignmentsDataSource().Schema["operator_control_assignment_collection"].Elem.(*schema.Resource).Schema)
		operatorControlAssignment["items"] = items
	}

	resources = append(resources, operatorControlAssignment)
	if err := s.D.Set("operator_control_assignment_collection", resources); err != nil {
		return err
	}

	return nil
}
