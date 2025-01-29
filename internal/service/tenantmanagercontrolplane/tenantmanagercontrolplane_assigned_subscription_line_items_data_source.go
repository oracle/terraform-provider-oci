// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneAssignedSubscriptionLineItemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTenantmanagercontrolplaneAssignedSubscriptionLineItems,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"assigned_subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"assigned_subscription_line_item_collection": {
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
									"billing_model": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"product_code": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"quantity": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_ended": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_started": {
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

func readTenantmanagercontrolplaneAssignedSubscriptionLineItems(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneAssignedSubscriptionLineItemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationsSubscriptionClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneAssignedSubscriptionLineItemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.SubscriptionClient
	Res    *oci_tenantmanagercontrolplane.ListAssignedSubscriptionLineItemsResponse
}

func (s *TenantmanagercontrolplaneAssignedSubscriptionLineItemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneAssignedSubscriptionLineItemsDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.ListAssignedSubscriptionLineItemsRequest{}

	if assignedSubscriptionId, ok := s.D.GetOkExists("assigned_subscription_id"); ok {
		tmp := assignedSubscriptionId.(string)
		request.AssignedSubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.ListAssignedSubscriptionLineItems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAssignedSubscriptionLineItems(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *TenantmanagercontrolplaneAssignedSubscriptionLineItemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneAssignedSubscriptionLineItemsDataSource-", TenantmanagercontrolplaneAssignedSubscriptionLineItemsDataSource(), s.D))
	resources := []map[string]interface{}{}
	assignedSubscriptionLineItem := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AssignedSubscriptionLineItemSummaryToMap(item))
	}
	assignedSubscriptionLineItem["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, TenantmanagercontrolplaneAssignedSubscriptionLineItemsDataSource().Schema["assigned_subscription_line_item_collection"].Elem.(*schema.Resource).Schema)
		assignedSubscriptionLineItem["items"] = items
	}

	resources = append(resources, assignedSubscriptionLineItem)
	if err := s.D.Set("assigned_subscription_line_item_collection", resources); err != nil {
		return err
	}

	return nil
}

func AssignedSubscriptionLineItemSummaryToMap(obj oci_tenantmanagercontrolplane.AssignedSubscriptionLineItemSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["billing_model"] = string(obj.BillingModel)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ProductCode != nil {
		result["product_code"] = string(*obj.ProductCode)
	}

	if obj.Quantity != nil {
		result["quantity"] = float32(*obj.Quantity)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}
