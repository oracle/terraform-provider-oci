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

func ManagementAgentManagementAgentDataSourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagementAgentManagementAgentDataSources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"management_agent_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ManagementAgentManagementAgentDataSourceSummary(),
			},
		},
	}
}

func ManagementAgentManagementAgentDataSourceSummary() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data_source_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func readManagementAgentManagementAgentDataSources(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentDataSourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentDataSourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.ListDataSourcesResponse
}

func (s *ManagementAgentManagementAgentDataSourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentDataSourcesDataSourceCrud) Get() error {
	request := oci_management_agent.ListDataSourcesRequest{}

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		request.Name = []string{name.(string)}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.ListDataSources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataSources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagementAgentManagementAgentDataSourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentManagementAgentDataSourcesDataSource-", ManagementAgentManagementAgentDataSourcesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		result := map[string]interface{}{}
		switch v := (r).(type) {
		case oci_management_agent.KubernetesClusterDataSourceSummary:
			result["type"] = "KUBERNETES_CLUSTER"

			if v.IsDaemonSet != nil {
				result["is_daemon_set"] = bool(*v.IsDaemonSet)
			}

			if v.Key != nil {
				result["data_source_key"] = string(*v.Key)
			}

			if v.Name != nil {
				result["name"] = string(*v.Name)
			}
		case oci_management_agent.PrometheusEmitterDataSourceSummary:
			result["type"] = "PROMETHEUS_EMITTER"

			if v.Key != nil {
				result["data_source_key"] = string(*v.Key)
			}

			if v.Name != nil {
				result["name"] = string(*v.Name)
			}
		default:
			log.Printf("[WARN] Received 'type' of unknown type %v", r)
			return nil
		}

		resources = append(resources, result)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ManagementAgentManagementAgentDataSourcesDataSource().Schema["data_sources"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("data_sources", resources); err != nil {
		return err
	}

	return nil
}
