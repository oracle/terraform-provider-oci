// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeAuditEventAnalyticDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeAuditEventAnalytic,
		Schema: map[string]*schema.Schema{
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_by": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"query_time_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scim_query": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"summary_field": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_ended": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dimensions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"audit_event_time": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"audit_type": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"client_hostname": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"client_id": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"client_program": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"db_user_name": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"event_name": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"object_type": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"target_class": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"target_id": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"target_name": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric_name": {
							Type:     schema.TypeString,
							Computed: true,
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_data_safe_audit_event", "oci_data_safe_audit_events"),
	}
}

func readSingularDataSafeAuditEventAnalytic(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditEventAnalyticDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeAuditEventAnalyticDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListAuditEventAnalyticsResponse
}

func (s *DataSafeAuditEventAnalyticDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeAuditEventAnalyticDataSourceCrud) Get() error {
	request := oci_data_safe.ListAuditEventAnalyticsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListAuditEventAnalyticsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		interfaces := groupBy.([]interface{})
		tmp := make([]oci_data_safe.ListAuditEventAnalyticsGroupByEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.ListAuditEventAnalyticsGroupByEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_by") {
			request.GroupBy = tmp
		}
	}

	if queryTimeZone, ok := s.D.GetOkExists("query_time_zone"); ok {
		tmp := queryTimeZone.(string)
		request.QueryTimeZone = &tmp
	}

	if scimQuery, ok := s.D.GetOkExists("scim_query"); ok {
		tmp := scimQuery.(string)
		request.ScimQuery = &tmp
	}

	if summaryField, ok := s.D.GetOkExists("summary_field"); ok {
		interfaces := summaryField.([]interface{})
		tmp := make([]oci_data_safe.ListAuditEventAnalyticsSummaryFieldEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_data_safe.ListAuditEventAnalyticsSummaryFieldEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("summary_field") {
			request.SummaryField = tmp
		}
	}

	if timeEnded, ok := s.D.GetOkExists("time_ended"); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnded.(string))
		if err != nil {
			return err
		}
		request.TimeEnded = &oci_common.SDKTime{Time: tmp}
	}

	if timeStarted, ok := s.D.GetOkExists("time_started"); ok {
		tmp, err := time.Parse(time.RFC3339, timeStarted.(string))
		if err != nil {
			return err
		}
		request.TimeStarted = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListAuditEventAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeAuditEventAnalyticDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeAuditEventAnalyticDataSource-", DataSafeAuditEventAnalyticDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AuditEventAggregationItemsToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func AuditEventAggregationDimensionsToMap(obj *oci_data_safe.AuditEventAggregationDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	result["audit_event_time"] = obj.AuditEventTime

	result["audit_type"] = obj.AuditType

	result["client_hostname"] = obj.ClientHostname

	result["client_id"] = obj.ClientId

	result["client_program"] = obj.ClientProgram

	result["db_user_name"] = obj.DbUserName

	result["event_name"] = obj.EventName

	result["object_type"] = obj.ObjectType

	result["target_class"] = obj.TargetClass

	result["target_id"] = obj.TargetId

	result["target_name"] = obj.TargetName

	return result
}

func AuditEventAggregationItemsToMap(obj oci_data_safe.AuditEventAggregationItems) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = strconv.FormatInt(*obj.Count, 10)
	}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{AuditEventAggregationDimensionsToMap(obj.Dimensions)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.MetricName != nil {
		result["metric_name"] = string(*obj.MetricName)
	}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}
