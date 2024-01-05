// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v65/cloudguard"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudGuardDataSourceEventDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCloudGuardDataSourceEvent,
		Schema: map[string]*schema.Schema{
			"data_source_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region": {
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
						"comments": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_source_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"event_info": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"data_source_feed_provider": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"log_result": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"observed_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"operator": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"trigger_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_cloud_guard_data_source_event", "oci_cloud_guard_data_source_events"),
	}
}

func readSingularCloudGuardDataSourceEvent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataSourceEventDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

type CloudGuardDataSourceEventDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_guard.CloudGuardClient
	Res    *oci_cloud_guard.ListDataSourceEventsResponse
}

func (s *CloudGuardDataSourceEventDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudGuardDataSourceEventDataSourceCrud) Get() error {
	request := oci_cloud_guard.ListDataSourceEventsRequest{}

	if dataSourceId, ok := s.D.GetOkExists("data_source_id"); ok {
		tmp := dataSourceId.(string)
		request.DataSourceId = &tmp
	}

	if region, ok := s.D.GetOkExists("region"); ok {
		tmp := region.(string)
		request.Region = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_guard")

	response, err := s.Client.ListDataSourceEvents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CloudGuardDataSourceEventDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudGuardDataSourceEventDataSource-", CloudGuardDataSourceEventDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DataSourceEventSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func DataSourceEventInfoToMap(obj *oci_cloud_guard.DataSourceEventInfo) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_cloud_guard.LoggingEventInfo:
		result["data_source_feed_provider"] = "LOGGINGQUERY"

		if v.LogResult != nil {
			result["log_result"] = string(*v.LogResult)
		}

		if v.ObservedValue != nil {
			result["observed_value"] = string(*v.ObservedValue)
		}

		if v.Operator != nil {
			result["operator"] = string(*v.Operator)
		}

		if v.TriggerValue != nil {
			result["trigger_value"] = string(*v.TriggerValue)
		}
	default:
		log.Printf("[WARN] Received 'data_source_feed_provider' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DataSourceEventSummaryToMap(obj oci_cloud_guard.DataSourceEventSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Comments != nil {
		result["comments"] = string(*obj.Comments)
	}

	if obj.DataSourceId != nil {
		result["data_source_id"] = string(*obj.DataSourceId)
	}

	if obj.EventDate != nil {
		result["event_date"] = obj.EventDate.String()
	}

	if obj.EventInfo != nil {
		eventInfoArray := []interface{}{}
		if eventInfoMap := DataSourceEventInfoToMap(&obj.EventInfo); eventInfoMap != nil {
			eventInfoArray = append(eventInfoArray, eventInfoMap)
		}
		result["event_info"] = eventInfoArray
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["status"] = string(obj.Status)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
