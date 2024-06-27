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

func DelegateAccessControlDelegationControlResourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDelegateAccessControlDelegationControlResources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"delegation_control_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"delegation_control_resource_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_status": {
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

func readDelegateAccessControlDelegationControlResources(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationControlResourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

type DelegateAccessControlDelegationControlResourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_delegate_access_control.DelegateAccessControlClient
	Res    *oci_delegate_access_control.ListDelegationControlResourcesResponse
}

func (s *DelegateAccessControlDelegationControlResourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DelegateAccessControlDelegationControlResourcesDataSourceCrud) Get() error {
	request := oci_delegate_access_control.ListDelegationControlResourcesRequest{}

	if delegationControlId, ok := s.D.GetOkExists("delegation_control_id"); ok {
		tmp := delegationControlId.(string)
		request.DelegationControlId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "delegate_access_control")

	response, err := s.Client.ListDelegationControlResources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDelegationControlResources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DelegateAccessControlDelegationControlResourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DelegateAccessControlDelegationControlResourcesDataSource-", DelegateAccessControlDelegationControlResourcesDataSource(), s.D))
	resources := []map[string]interface{}{}
	delegationControlResource := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DelegationControlResourceSummaryToMap(item))
	}
	delegationControlResource["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DelegateAccessControlDelegationControlResourcesDataSource().Schema["delegation_control_resource_collection"].Elem.(*schema.Resource).Schema)
		delegationControlResource["items"] = items
	}

	resources = append(resources, delegationControlResource)
	if err := s.D.Set("delegation_control_resource_collection", resources); err != nil {
		return err
	}

	return nil
}

func DelegationControlResourceSummaryToMap(obj oci_delegate_access_control.DelegationControlResourceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["resource_status"] = string(obj.ResourceStatus)

	return result
}
