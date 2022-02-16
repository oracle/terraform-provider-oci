// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v58/loganalytics"
)

func LogAnalyticsLogAnalyticsPreferenceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsLogAnalyticsPreference,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
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
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularLogAnalyticsLogAnalyticsPreference(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsPreferenceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsLogAnalyticsPreferenceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetPreferencesResponse
}

func (s *LogAnalyticsLogAnalyticsPreferenceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsPreferenceDataSourceCrud) Get() error {
	request := oci_log_analytics.GetPreferencesRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetPreferences(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsLogAnalyticsPreferenceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LogAnalyticsLogAnalyticsPreferenceDataSource-", LogAnalyticsLogAnalyticsPreferenceDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, LogAnalyticsPreferenceToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func LogAnalyticsPreferenceToMap(obj oci_log_analytics.LogAnalyticsPreference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}
