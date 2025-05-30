// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apiaccesscontrol

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apiaccesscontrol "github.com/oracle/oci-go-sdk/v65/apiaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApiaccesscontrolPrivilegedApiControlsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApiaccesscontrolPrivilegedApiControls,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
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
			"privileged_api_control_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ApiaccesscontrolPrivilegedApiControlResource()),
						},
					},
				},
			},
		},
	}
}

func readApiaccesscontrolPrivilegedApiControls(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolPrivilegedApiControlsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivilegedApiControlClient()

	return tfresource.ReadResource(sync)
}

type ApiaccesscontrolPrivilegedApiControlsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apiaccesscontrol.PrivilegedApiControlClient
	Res    *oci_apiaccesscontrol.ListPrivilegedApiControlsResponse
}

func (s *ApiaccesscontrolPrivilegedApiControlsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApiaccesscontrolPrivilegedApiControlsDataSourceCrud) Get() error {
	request := oci_apiaccesscontrol.ListPrivilegedApiControlsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if state, ok := s.D.GetOkExists("lifecycle_state"); ok {
		request.LifecycleState = oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apiaccesscontrol")

	response, err := s.Client.ListPrivilegedApiControls(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPrivilegedApiControls(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApiaccesscontrolPrivilegedApiControlsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApiaccesscontrolPrivilegedApiControlsDataSource-", ApiaccesscontrolPrivilegedApiControlsDataSource(), s.D))
	resources := []map[string]interface{}{}
	privilegedApiControl := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PrivilegedApiControlSummaryToMap(item))
	}
	privilegedApiControl["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApiaccesscontrolPrivilegedApiControlsDataSource().Schema["privileged_api_control_collection"].Elem.(*schema.Resource).Schema)
		privilegedApiControl["items"] = items
	}

	resources = append(resources, privilegedApiControl)
	if err := s.D.Set("privileged_api_control_collection", resources); err != nil {
		return err
	}

	return nil
}
