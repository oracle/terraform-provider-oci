// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ons

import (
	"context"
	"fmt"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_ons "github.com/oracle/oci-go-sdk/v58/ons"
)

func OnsSubscriptionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOnsSubscriptions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"topic_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscriptions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     modifySubscriptionSchema(tfresource.GetDataSourceItemSchema(OnsSubscriptionResource())),
			},
		},
	}
}

func readOnsSubscriptions(d *schema.ResourceData, m interface{}) error {
	sync := &OnsSubscriptionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NotificationDataPlaneClient()

	return tfresource.ReadResource(sync)
}

type OnsSubscriptionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ons.NotificationDataPlaneClient
	Res    *oci_ons.ListSubscriptionsResponse
}

func (s *OnsSubscriptionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OnsSubscriptionsDataSourceCrud) Get() error {
	request := oci_ons.ListSubscriptionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if topicId, ok := s.D.GetOkExists("topic_id"); ok {
		tmp := topicId.(string)
		request.TopicId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ons")

	response, err := s.Client.ListSubscriptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSubscriptions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OnsSubscriptionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OnsSubscriptionsDataSource-", OnsSubscriptionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		subscription := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedTime != nil {
			subscription["created_time"] = strconv.FormatInt(*r.CreatedTime, 10)
		}

		if r.DefinedTags != nil {
			subscription["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DeliveryPolicy != nil {
			subscription["delivery_policy"] = []interface{}{DeliveryPolicyToMap(r.DeliveryPolicy)}
		} else {
			subscription["delivery_policy"] = nil
		}

		if r.Endpoint != nil {
			subscription["endpoint"] = *r.Endpoint
		}

		if r.Etag != nil {
			subscription["etag"] = *r.Etag
		}

		subscription["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			subscription["id"] = *r.Id
		}

		if r.Protocol != nil {
			subscription["protocol"] = *r.Protocol
		}

		subscription["state"] = r.LifecycleState

		if r.TopicId != nil {
			subscription["topic_id"] = *r.TopicId
		}

		resources = append(resources, subscription)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OnsSubscriptionsDataSource().Schema["subscriptions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("subscriptions", resources); err != nil {
		return err
	}

	return nil
}

func parseDeliveryPolicy(policy interface{}) string {
	backoffRetryPolicy := policy.(map[string]interface{})["backoff_retry_policy"].([]interface{})
	maxRetryDuration := backoffRetryPolicy[0].(map[string]interface{})["max_retry_duration"]
	policyType := backoffRetryPolicy[0].(map[string]interface{})["policy_type"]
	return fmt.Sprintf("{\"backoffRetryPolicy\":{\"maxRetryDuration\":%v,\"policyType\":\"%v\"}}", maxRetryDuration, policyType)
}

func modifySubscriptionSchema(resourceSchema *schema.Resource) *schema.Resource {
	resourceSchema.Schema["delivery_policy"] = &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Required

				// Optional

				// Computed
				"backoff_retry_policy": {
					Type:     schema.TypeList,
					Computed: true,
					MaxItems: 1,
					MinItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							// Required

							// Optional

							// Computed
							"max_retry_duration": {
								Type:     schema.TypeInt,
								Computed: true,
							},
							"policy_type": {
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
	resourceSchema.Schema["created_time"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	return resourceSchema
}
