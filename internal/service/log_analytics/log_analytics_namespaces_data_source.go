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

func LogAnalyticsNamespacesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsNamespaces,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace_collection": {
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
									"is_onboarded": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"namespace": {
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

func readLogAnalyticsNamespaces(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespacesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespacesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListNamespacesResponse
}

func (s *LogAnalyticsNamespacesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespacesDataSourceCrud) Get() error {
	request := oci_log_analytics.ListNamespacesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListNamespaces(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespacesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsNamespacesDataSource-", LogAnalyticsNamespacesDataSource(), s.D))
	resources := []map[string]interface{}{}
	namespace := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, NamespaceSummaryToMap(item))
	}
	namespace["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsNamespacesDataSource().Schema["namespace_collection"].Elem.(*schema.Resource).Schema)
		namespace["items"] = items
	}

	resources = append(resources, namespace)
	if err := s.D.Set("namespace_collection", resources); err != nil {
		return err
	}

	return nil
}

func NamespaceSummaryToMap(obj oci_log_analytics.NamespaceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.IsOnboarded != nil {
		result["is_onboarded"] = bool(*obj.IsOnboarded)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	return result
}
