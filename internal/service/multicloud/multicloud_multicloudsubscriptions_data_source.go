// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package multicloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_multicloud "github.com/oracle/oci-go-sdk/v65/multicloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MulticloudMulticloudsubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readMulticloudMulticloudsubscriptionsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicloud_subscription_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Computed
									"classic_subscription_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subscription_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"partner_cloud_account_identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_linked_date": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"payment_plan": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"active_commitment": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_end_date": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"lifecycle_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"csp_additional_properties": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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

func readMulticloudMulticloudsubscriptionsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MulticloudMulticloudsubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MulticloudsubscriptionsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type MulticloudMulticloudsubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.MulticloudsubscriptionsClient
	Res    *oci_multicloud.ListMulticloudSubscriptionsResponse
}

func (s *MulticloudMulticloudsubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudMulticloudsubscriptionsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_multicloud.ListMulticloudSubscriptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.ListMulticloudSubscriptions(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMulticloudSubscriptions(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MulticloudMulticloudsubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudMulticloudsubscriptionsDataSource-", MulticloudMulticloudsubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	multicloudsubscription := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MulticloudSubscriptionSummaryToMap(item))
	}
	multicloudsubscription["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MulticloudMulticloudsubscriptionsDataSource().Schema["multicloud_subscription_collection"].Elem.(*schema.Resource).Schema)
		multicloudsubscription["items"] = items
	}

	resources = append(resources, multicloudsubscription)
	if err := s.D.Set("multicloud_subscription_collection", resources); err != nil {
		return err
	}

	return nil
}

func MulticloudSubscriptionSummaryToMap(obj oci_multicloud.MulticloudSubscriptionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClassicSubscriptionId != nil {
		result["classic_subscription_id"] = string(*obj.ClassicSubscriptionId)
	}

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	if obj.PartnerCloudAccountIdentifier != nil {
		result["partner_cloud_account_identifier"] = string(*obj.PartnerCloudAccountIdentifier)
	}

	result["service_name"] = string(obj.ServiceName)

	if obj.TimeLinkedDate != nil {
		result["time_linked_date"] = obj.TimeLinkedDate.String()
	}

	if obj.PaymentPlan != nil {
		result["payment_plan"] = string(*obj.PaymentPlan)
	}

	if obj.ActiveCommitment != nil {
		result["active_commitment"] = string(*obj.ActiveCommitment)
	}

	if obj.TimeEndDate != nil {
		result["time_end_date"] = obj.TimeEndDate.String()
	}

	result["lifecycle_state"] = string(obj.LifecycleState)

	result["csp_additional_properties"] = obj.CspAdditionalProperties

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
