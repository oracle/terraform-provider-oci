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

func OperatorAccessControlOperatorControlsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOperatorAccessControlOperatorControls,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
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
			"operator_control_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OperatorAccessControlOperatorControlResource()),
						},
					},
				},
			},
		},
	}
}

func readOperatorAccessControlOperatorControls(d *schema.ResourceData, m interface{}) error {
	sync := &OperatorAccessControlOperatorControlsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperatorControlClient()

	return tfresource.ReadResource(sync)
}

type OperatorAccessControlOperatorControlsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_operator_access_control.OperatorControlClient
	Res    *oci_operator_access_control.ListOperatorControlsResponse
}

func (s *OperatorAccessControlOperatorControlsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OperatorAccessControlOperatorControlsDataSourceCrud) Get() error {
	request := oci_operator_access_control.ListOperatorControlsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := string(oci_operator_access_control.ResourceTypesEnum(resourceType.(string)))
		request.ResourceType = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_operator_access_control.ListOperatorControlsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "operator_access_control")

	response, err := s.Client.ListOperatorControls(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	// Operator Controls is not supporting serverside pagination
	// request.Page = s.Res.OpcNextPage

	// for request.Page != nil {
	// 	listResponse, err := s.Client.ListOperatorControls(context.Background(), request)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	s.Res.Items = append(s.Res.Items, listResponse.Items...)
	// 	request.Page = listResponse.OpcNextPage
	// }

	return nil
}

func (s *OperatorAccessControlOperatorControlsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OperatorAccessControlOperatorControlsDataSource-", OperatorAccessControlOperatorControlsDataSource(), s.D))
	resources := []map[string]interface{}{}
	operatorControl := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OperatorControlSummaryToMap(item))
	}
	operatorControl["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OperatorAccessControlOperatorControlsDataSource().Schema["operator_control_collection"].Elem.(*schema.Resource).Schema)
		operatorControl["items"] = items
	}

	resources = append(resources, operatorControl)
	if err := s.D.Set("operator_control_collection", resources); err != nil {
		return err
	}

	return nil
}
