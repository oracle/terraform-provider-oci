// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentDataSourceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["data_source_key"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["management_agent_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ManagementAgentManagementAgentDataSourceResource(), fieldMap, readSingularManagementAgentManagementAgentDataSource)
}

func readSingularManagementAgentManagementAgentDataSource(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentDataSourceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentDataSourceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.GetDataSourceResponse
}

func (s *ManagementAgentManagementAgentDataSourceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentDataSourceDataSourceCrud) Get() error {
	request := oci_management_agent.GetDataSourceRequest{}

	if dataSourceKey, ok := s.D.GetOkExists("data_source_key"); ok {
		tmp := dataSourceKey.(string)
		request.DataSourceKey = &tmp
	}

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.GetDataSource(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagementAgentManagementAgentDataSourceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentManagementAgentDataSourceDataSource-", ManagementAgentManagementAgentDataSourceDataSource(), s.D))
	switch v := (s.Res.DataSource).(type) {
	case oci_management_agent.KubernetesClusterDataSource:
		s.D.Set("type", "KUBERNETES_CLUSTER")

		if v.IsDaemonSet != nil {
			s.D.Set("is_daemon_set", *v.IsDaemonSet)
		}

		if v.Namespace != nil {
			s.D.Set("namespace", *v.Namespace)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.Key != nil {
			s.D.Set("data_source_key", *v.Key)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		s.D.Set("state", v.State)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_management_agent.PrometheusEmitterDataSource:
		s.D.Set("type", "PROMETHEUS_EMITTER")

		if v.AllowMetrics != nil {
			s.D.Set("allow_metrics", *v.AllowMetrics)
		}

		if v.ConnectionTimeout != nil {
			s.D.Set("connection_timeout", *v.ConnectionTimeout)
		}

		metricDimensions := []interface{}{}
		for _, item := range v.MetricDimensions {
			metricDimensions = append(metricDimensions, MetricDimensionToMap(item))
		}
		s.D.Set("metric_dimensions", metricDimensions)

		if v.Namespace != nil {
			s.D.Set("namespace", *v.Namespace)
		}

		if v.ProxyUrl != nil {
			s.D.Set("proxy_url", *v.ProxyUrl)
		}

		if v.ReadDataLimit != nil {
			s.D.Set("read_data_limit", *v.ReadDataLimit)
		}

		if v.ReadTimeout != nil {
			s.D.Set("read_timeout", *v.ReadTimeout)
		}

		if v.ResourceGroup != nil {
			s.D.Set("resource_group", *v.ResourceGroup)
		}

		if v.ScheduleMins != nil {
			s.D.Set("schedule_mins", *v.ScheduleMins)
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.Key != nil {
			s.D.Set("data_source_key", *v.Key)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		s.D.Set("state", v.State)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res.DataSource)
		return nil
	}

	return nil
}
