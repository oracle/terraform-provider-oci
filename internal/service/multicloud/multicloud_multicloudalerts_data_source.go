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

func MulticloudMulticloudalertsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readMulticloudMulticloudalertsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"alert_function_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
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
			"severity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_service_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicloud_alert_collection": {
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
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"multicloud_alerts": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"additional_parameters": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"alert_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"alert_status": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"alert_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"compartment_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"defined_tags": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"description": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"display_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"freeform_tags": {
													Type:     schema.TypeMap,
													Computed: true,
													Elem:     schema.TypeString,
												},
												"function_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"resource_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"resource_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"severity": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"source": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"source_region": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"lifecycle_state": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"subscription_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"subscription_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"system_tags": {
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
											},
										},
									},
									"multicloudalert_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"lifecycle_state": {
										Type:     schema.TypeString,
										Computed: true,
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

func readMulticloudMulticloudalertsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &MulticloudMulticloudalertsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MulticloudAlertsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type MulticloudMulticloudalertsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.MulticloudAlertsClient
	Res    *oci_multicloud.ListMulticloudAlertsResponse
}

func (s *MulticloudMulticloudalertsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudMulticloudalertsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_multicloud.ListMulticloudAlertsRequest{}

	if alertFunctionName, ok := s.D.GetOkExists("alert_function_name"); ok {
		tmp := alertFunctionName.(string)
		request.AlertFunctionName = &tmp
	}

	if alertStatus, ok := s.D.GetOkExists("alert_status"); ok {
		request.AlertStatus = oci_multicloud.ListMulticloudAlertsAlertStatusEnum(alertStatus.(string))
	}

	if alertType, ok := s.D.GetOkExists("alert_type"); ok {
		tmp := alertType.(string)
		request.AlertType = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		tmp := resourceType.(string)
		request.ResourceType = &tmp
	}

	if severity, ok := s.D.GetOkExists("severity"); ok {
		request.Severity = oci_multicloud.ListMulticloudAlertsSeverityEnum(severity.(string))
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if subscriptionServiceName, ok := s.D.GetOkExists("subscription_service_name"); ok {
		request.SubscriptionServiceName = oci_multicloud.ListMulticloudAlertsSubscriptionServiceNameEnum(subscriptionServiceName.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.ListMulticloudAlerts(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMulticloudAlerts(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MulticloudMulticloudalertsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudMulticloudalertsDataSource-", MulticloudMulticloudalertsDataSource(), s.D))
	resources := []map[string]interface{}{}
	multicloudalert := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MulticloudAlertSummaryToMap(item))
	}
	multicloudalert["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MulticloudMulticloudalertsDataSource().Schema["multicloud_alert_collection"].Elem.(*schema.Resource).Schema)
		multicloudalert["items"] = items
	}

	resources = append(resources, multicloudalert)
	if err := s.D.Set("multicloud_alert_collection", resources); err != nil {
		return err
	}

	return nil
}

func MulticloudAlertToMap(obj oci_multicloud.MulticloudAlert) map[string]interface{} {
	result := map[string]interface{}{}

	result["additional_parameters"] = obj.AdditionalParameters

	if obj.AlertId != nil {
		result["alert_id"] = string(*obj.AlertId)
	}

	result["alert_status"] = string(obj.AlertStatus)

	if obj.AlertType != nil {
		result["alert_type"] = string(*obj.AlertType)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.FunctionType != nil {
		result["function_type"] = string(*obj.FunctionType)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ResourceId != nil {
		result["resource_id"] = string(*obj.ResourceId)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["severity"] = string(obj.Severity)

	if obj.Source != nil {
		result["source"] = string(*obj.Source)
	}

	if obj.SourceRegion != nil {
		result["source_region"] = string(*obj.SourceRegion)
	}

	result["lifecycle_state"] = string(obj.LifecycleState)

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	result["subscription_type"] = string(obj.SubscriptionType)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func MulticloudAlertSummaryToMap(obj oci_multicloud.MulticloudAlertSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	multicloudAlerts := []interface{}{}
	for _, item := range obj.MulticloudAlerts {
		multicloudAlerts = append(multicloudAlerts, MulticloudAlertToMap(item))
	}
	result["multicloud_alerts"] = multicloudAlerts

	if obj.Count != nil {
		result["multicloudalert_count"] = int(*obj.Count)
	}

	result["lifecycle_state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
