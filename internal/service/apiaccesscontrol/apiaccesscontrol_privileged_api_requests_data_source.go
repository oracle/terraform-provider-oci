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

func ApiaccesscontrolPrivilegedApiRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApiaccesscontrolPrivilegedApiRequests,
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
			"resource_id": {
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
			"privileged_api_request_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ApiaccesscontrolPrivilegedApiRequestResource()),
						},
					},
				},
			},
		},
	}
}

func readApiaccesscontrolPrivilegedApiRequests(d *schema.ResourceData, m interface{}) error {
	sync := &ApiaccesscontrolPrivilegedApiRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PrivilegedApiRequestsClient()

	return tfresource.ReadResource(sync)
}

type ApiaccesscontrolPrivilegedApiRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apiaccesscontrol.PrivilegedApiRequestsClient
	Res    *oci_apiaccesscontrol.ListPrivilegedApiRequestsResponse
}

func (s *ApiaccesscontrolPrivilegedApiRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApiaccesscontrolPrivilegedApiRequestsDataSourceCrud) Get() error {
	request := oci_apiaccesscontrol.ListPrivilegedApiRequestsRequest{}

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

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if state, ok := s.D.GetOkExists("lifecycle_state"); ok {
		request.LifecycleState = oci_apiaccesscontrol.PrivilegedApiRequestLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apiaccesscontrol")

	response, err := s.Client.ListPrivilegedApiRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPrivilegedApiRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApiaccesscontrolPrivilegedApiRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ApiaccesscontrolPrivilegedApiRequestsDataSource-", ApiaccesscontrolPrivilegedApiRequestsDataSource(), s.D))
	resources := []map[string]interface{}{}
	privilegedApiRequest := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PrivilegedApiRequestSummaryToMap(item))
	}
	privilegedApiRequest["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ApiaccesscontrolPrivilegedApiRequestsDataSource().Schema["privileged_api_request_collection"].Elem.(*schema.Resource).Schema)
		privilegedApiRequest["items"] = items
	}

	resources = append(resources, privilegedApiRequest)
	if err := s.D.Set("privileged_api_request_collection", resources); err != nil {
		return err
	}

	return nil
}
