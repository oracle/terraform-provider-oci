// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceIngestTimeRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsNamespaceIngestTimeRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"condition_kind": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"field_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"field_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ingest_time_rule_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"condition_kind": {
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
									"field_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"field_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
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

func readLogAnalyticsNamespaceIngestTimeRules(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceIngestTimeRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceIngestTimeRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListIngestTimeRulesResponse
}

func (s *LogAnalyticsNamespaceIngestTimeRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceIngestTimeRulesDataSourceCrud) Get() error {
	request := oci_log_analytics.ListIngestTimeRulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if conditionKind, ok := s.D.GetOkExists("condition_kind"); ok {
		request.ConditionKind = oci_log_analytics.ListIngestTimeRulesConditionKindEnum(conditionKind.(string))
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fieldName, ok := s.D.GetOkExists("field_name"); ok {
		tmp := fieldName.(string)
		request.FieldName = &tmp
	}

	if fieldValue, ok := s.D.GetOkExists("field_value"); ok {
		tmp := fieldValue.(string)
		request.FieldValue = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_log_analytics.ListIngestTimeRulesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListIngestTimeRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIngestTimeRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsNamespaceIngestTimeRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespaceIngestTimeRulesDataSource-", LogAnalyticsNamespaceIngestTimeRulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	namespaceIngestTimeRule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, IngestTimeRuleSummaryToMap(item))
	}
	namespaceIngestTimeRule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsNamespaceIngestTimeRulesDataSource().Schema["ingest_time_rule_summary_collection"].Elem.(*schema.Resource).Schema)
		namespaceIngestTimeRule["items"] = items
	}

	resources = append(resources, namespaceIngestTimeRule)
	if err := s.D.Set("ingest_time_rule_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
