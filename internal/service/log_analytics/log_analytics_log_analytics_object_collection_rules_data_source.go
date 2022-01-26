// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v56/loganalytics"
)

func LogAnalyticsLogAnalyticsObjectCollectionRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsLogAnalyticsObjectCollectionRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
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
			"log_analytics_object_collection_rule_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(LogAnalyticsLogAnalyticsObjectCollectionRuleResource()),
						},
					},
				},
			},
		},
	}
}

func readLogAnalyticsLogAnalyticsObjectCollectionRules(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsObjectCollectionRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsLogAnalyticsObjectCollectionRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListLogAnalyticsObjectCollectionRulesResponse
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRulesDataSourceCrud) Get() error {
	request := oci_log_analytics.ListLogAnalyticsObjectCollectionRulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_log_analytics.ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListLogAnalyticsObjectCollectionRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLogAnalyticsObjectCollectionRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsLogAnalyticsObjectCollectionRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsLogAnalyticsObjectCollectionRulesDataSource-", LogAnalyticsLogAnalyticsObjectCollectionRulesDataSource(), s.D))
	resources := []map[string]interface{}{}
	logAnalyticsObjectCollectionRule := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LogAnalyticsObjectCollectionRuleSummaryToMap(item))
	}
	logAnalyticsObjectCollectionRule["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsLogAnalyticsObjectCollectionRulesDataSource().Schema["log_analytics_object_collection_rule_collection"].Elem.(*schema.Resource).Schema)
		logAnalyticsObjectCollectionRule["items"] = items
	}

	resources = append(resources, logAnalyticsObjectCollectionRule)
	if err := s.D.Set("log_analytics_object_collection_rule_collection", resources); err != nil {
		return err
	}

	return nil
}
