// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_monitoring "github.com/oracle/oci-go-sdk/monitoring"
)

func MonitoringAlarmDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMonitoringAlarm,
		Schema: map[string]*schema.Schema{
			"alarm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"body": {
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
			"destinations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"is_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"metric_compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metric_compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pending_duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"query": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repeat_notification_duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resolution": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"suppression": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"time_suppress_from": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: timeDiffSuppressFunction,
						},
						"time_suppress_until": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: timeDiffSuppressFunction,
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
	}
}

func readSingularMonitoringAlarm(d *schema.ResourceData, m interface{}) error {
	sync := &MonitoringAlarmDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).monitoringClient

	return ReadResource(sync)
}

type MonitoringAlarmDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_monitoring.MonitoringClient
	Res    *oci_monitoring.GetAlarmResponse
}

func (s *MonitoringAlarmDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MonitoringAlarmDataSourceCrud) Get() error {
	request := oci_monitoring.GetAlarmRequest{}

	if alarmId, ok := s.D.GetOkExists("alarm_id"); ok {
		tmp := alarmId.(string)
		request.AlarmId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "monitoring")

	response, err := s.Client.GetAlarm(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MonitoringAlarmDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Body != nil {
		s.D.Set("body", *s.Res.Body)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("destinations", s.Res.Destinations)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.MetricCompartmentId != nil {
		s.D.Set("metric_compartment_id", *s.Res.MetricCompartmentId)
	}

	if s.Res.MetricCompartmentIdInSubtree != nil {
		s.D.Set("metric_compartment_id_in_subtree", *s.Res.MetricCompartmentIdInSubtree)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.PendingDuration != nil {
		s.D.Set("pending_duration", *s.Res.PendingDuration)
	}

	if s.Res.Query != nil {
		s.D.Set("query", *s.Res.Query)
	}

	if s.Res.RepeatNotificationDuration != nil {
		s.D.Set("repeat_notification_duration", *s.Res.RepeatNotificationDuration)
	}

	if s.Res.Resolution != nil {
		s.D.Set("resolution", *s.Res.Resolution)
	}

	s.D.Set("severity", s.Res.Severity)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.Suppression != nil {
		s.D.Set("suppression", []interface{}{SuppressionToMap(s.Res.Suppression)})
	} else {
		s.D.Set("suppression", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
