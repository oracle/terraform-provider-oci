// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package delegate_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DelegateAccessControlDelegationControlsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDelegateAccessControlDelegationControls,
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
			"delegation_control_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DelegateAccessControlDelegationControlResource()),
						},
					},
				},
			},
		},
	}
}

func readDelegateAccessControlDelegationControls(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationControlsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

type DelegateAccessControlDelegationControlsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_delegate_access_control.DelegateAccessControlClient
	Res    *oci_delegate_access_control.ListDelegationControlsResponse
}

func (s *DelegateAccessControlDelegationControlsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DelegateAccessControlDelegationControlsDataSourceCrud) Get() error {
	request := oci_delegate_access_control.ListDelegationControlsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		request.ResourceType = oci_delegate_access_control.ListDelegationControlsResourceTypeEnum(resourceType.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_delegate_access_control.DelegationControlLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "delegate_access_control")

	response, err := s.Client.ListDelegationControls(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDelegationControls(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DelegateAccessControlDelegationControlsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DelegateAccessControlDelegationControlsDataSource-", DelegateAccessControlDelegationControlsDataSource(), s.D))
	resources := []map[string]interface{}{}
	delegationControl := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DelegationControlSummaryToMap(item))
	}
	delegationControl["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DelegateAccessControlDelegationControlsDataSource().Schema["delegation_control_summary_collection"].Elem.(*schema.Resource).Schema)
		delegationControl["items"] = items
	}

	resources = append(resources, delegationControl)
	if err := s.D.Set("delegation_control_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
