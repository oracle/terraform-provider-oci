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

func DelegateAccessControlDelegatedResourceAccessRequestHistoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDelegateAccessControlDelegatedResourceAccessRequestHistories,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"delegated_resource_access_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"delegated_resource_access_request_history_collection": {
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
									"comment": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"request_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"timestamp": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"user_id": {
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

func readDelegateAccessControlDelegatedResourceAccessRequestHistories(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegatedResourceAccessRequestHistoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

type DelegateAccessControlDelegatedResourceAccessRequestHistoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_delegate_access_control.DelegateAccessControlClient
	Res    *oci_delegate_access_control.ListDelegatedResourceAccessRequestHistoriesResponse
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestHistoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestHistoriesDataSourceCrud) Get() error {
	request := oci_delegate_access_control.ListDelegatedResourceAccessRequestHistoriesRequest{}

	if delegatedResourceAccessRequestId, ok := s.D.GetOkExists("delegated_resource_access_request_id"); ok {
		tmp := delegatedResourceAccessRequestId.(string)
		request.DelegatedResourceAccessRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "delegate_access_control")

	response, err := s.Client.ListDelegatedResourceAccessRequestHistories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDelegatedResourceAccessRequestHistories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DelegateAccessControlDelegatedResourceAccessRequestHistoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DelegateAccessControlDelegatedResourceAccessRequestHistoriesDataSource-", DelegateAccessControlDelegatedResourceAccessRequestHistoriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	delegatedResourceAccessRequestHistory := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DelegatedResourceAccessRequestHistorySummaryToMap(item))
	}
	delegatedResourceAccessRequestHistory["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DelegateAccessControlDelegatedResourceAccessRequestHistoriesDataSource().Schema["delegated_resource_access_request_history_collection"].Elem.(*schema.Resource).Schema)
		delegatedResourceAccessRequestHistory["items"] = items
	}

	resources = append(resources, delegatedResourceAccessRequestHistory)
	if err := s.D.Set("delegated_resource_access_request_history_collection", resources); err != nil {
		return err
	}

	return nil
}

func DelegatedResourceAccessRequestHistorySummaryToMap(obj oci_delegate_access_control.DelegatedResourceAccessRequestHistorySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Comment != nil {
		result["comment"] = string(*obj.Comment)
	}

	result["request_status"] = string(obj.RequestStatus)

	result["state"] = string(obj.LifecycleState)

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.String()
	}

	if obj.UserId != nil {
		result["user_id"] = string(*obj.UserId)
	}

	return result
}
