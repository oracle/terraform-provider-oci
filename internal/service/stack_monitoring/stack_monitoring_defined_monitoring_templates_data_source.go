// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringDefinedMonitoringTemplatesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readStackMonitoringDefinedMonitoringTemplates,
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
			"resource_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"defined_monitoring_template_collection": {
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
									"composite_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_alarm_conditions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"condition_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"conditions": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"body": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"query": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"severity": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"should_append_note": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"should_append_url": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"trigger_delay": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"metric_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_type": {
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
					},
				},
			},
		},
	}
}

func readStackMonitoringDefinedMonitoringTemplates(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringDefinedMonitoringTemplatesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringDefinedMonitoringTemplatesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.ListDefinedMonitoringTemplatesResponse
}

func (s *StackMonitoringDefinedMonitoringTemplatesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringDefinedMonitoringTemplatesDataSourceCrud) Get() error {
	request := oci_stack_monitoring.ListDefinedMonitoringTemplatesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if resourceTypes, ok := s.D.GetOkExists("resource_types"); ok {
		interfaces := resourceTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("resource_types") {
			request.ResourceTypes = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.ListDefinedMonitoringTemplates(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDefinedMonitoringTemplates(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *StackMonitoringDefinedMonitoringTemplatesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("StackMonitoringDefinedMonitoringTemplatesDataSource-", StackMonitoringDefinedMonitoringTemplatesDataSource(), s.D))
	resources := []map[string]interface{}{}
	definedMonitoringTemplate := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DefinedMonitoringTemplateSummaryToMap(item))
	}
	definedMonitoringTemplate["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, StackMonitoringDefinedMonitoringTemplatesDataSource().Schema["defined_monitoring_template_collection"].Elem.(*schema.Resource).Schema)
		definedMonitoringTemplate["items"] = items
	}

	resources = append(resources, definedMonitoringTemplate)
	if err := s.D.Set("defined_monitoring_template_collection", resources); err != nil {
		return err
	}

	return nil
}

func DefinedAlarmConditionToMap(obj oci_stack_monitoring.DefinedAlarmCondition) map[string]interface{} {
	result := map[string]interface{}{}

	result["condition_type"] = string(obj.ConditionType)

	conditions := []interface{}{}
	for _, item := range obj.Conditions {
		conditions = append(conditions, ConditionToMap(item))
	}
	result["conditions"] = conditions

	if obj.MetricName != nil {
		result["metric_name"] = string(*obj.MetricName)
	}

	return result
}

func DefinedMonitoringTemplateSummaryToMap(obj oci_stack_monitoring.DefinedMonitoringTemplateSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompositeType != nil {
		result["composite_type"] = string(*obj.CompositeType)
	}

	definedAlarmConditions := []interface{}{}
	for _, item := range obj.DefinedAlarmConditions {
		definedAlarmConditions = append(definedAlarmConditions, DefinedAlarmConditionToMap(item))
	}
	result["defined_alarm_conditions"] = definedAlarmConditions

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

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
