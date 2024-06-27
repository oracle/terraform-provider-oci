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

func DelegateAccessControlDelegationSubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDelegateAccessControlDelegationSubscriptions,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"delegation_subscription_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DelegateAccessControlDelegationSubscriptionResource()),
						},
					},
				},
			},
		},
	}
}

func readDelegateAccessControlDelegationSubscriptions(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationSubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

type DelegateAccessControlDelegationSubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_delegate_access_control.DelegateAccessControlClient
	Res    *oci_delegate_access_control.ListDelegationSubscriptionsResponse
}

func (s *DelegateAccessControlDelegationSubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DelegateAccessControlDelegationSubscriptionsDataSourceCrud) Get() error {
	request := oci_delegate_access_control.ListDelegationSubscriptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_delegate_access_control.DelegationSubscriptionLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "delegate_access_control")

	response, err := s.Client.ListDelegationSubscriptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDelegationSubscriptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DelegateAccessControlDelegationSubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DelegateAccessControlDelegationSubscriptionsDataSource-", DelegateAccessControlDelegationSubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	delegationSubscription := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DelegationSubscriptionSummaryToMap(item))
	}
	delegationSubscription["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DelegateAccessControlDelegationSubscriptionsDataSource().Schema["delegation_subscription_summary_collection"].Elem.(*schema.Resource).Schema)
		delegationSubscription["items"] = items
	}

	resources = append(resources, delegationSubscription)
	if err := s.D.Set("delegation_subscription_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
