// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsJmsPluginsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readJmsJmsPlugins,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"agent_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"availability_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_last_seen_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_registered_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jms_plugin_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(JmsJmsPluginResource()),
						},
					},
				},
			},
		},
	}
}

func readJmsJmsPlugins(d *schema.ResourceData, m interface{}) error {
	sync := &JmsJmsPluginsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type JmsJmsPluginsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_jms.JavaManagementServiceClient
	Res    *oci_jms.ListJmsPluginsResponse
}

func (s *JmsJmsPluginsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *JmsJmsPluginsDataSourceCrud) Get() error {
	request := oci_jms.ListJmsPluginsRequest{}

	if agentId, ok := s.D.GetOkExists("agent_id"); ok {
		tmp := agentId.(string)
		request.AgentId = &tmp
	}

	if availabilityStatus, ok := s.D.GetOkExists("availability_status"); ok {
		request.AvailabilityStatus = oci_jms.ListJmsPluginsAvailabilityStatusEnum(availabilityStatus.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if hostnameContains, ok := s.D.GetOkExists("hostname_contains"); ok {
		tmp := hostnameContains.(string)
		request.HostnameContains = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_jms.ListJmsPluginsLifecycleStateEnum(state.(string))
	}

	if timeLastSeenLessThanOrEqualTo, ok := s.D.GetOkExists("time_last_seen_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeLastSeenLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeLastSeenLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeRegisteredLessThanOrEqualTo, ok := s.D.GetOkExists("time_registered_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeRegisteredLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeRegisteredLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "jms")

	response, err := s.Client.ListJmsPlugins(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListJmsPlugins(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *JmsJmsPluginsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("JmsJmsPluginsDataSource-", JmsJmsPluginsDataSource(), s.D))
	resources := []map[string]interface{}{}
	jmsPlugin := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, JmsPluginSummaryToMap(item))
	}
	jmsPlugin["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, JmsJmsPluginsDataSource().Schema["jms_plugin_collection"].Elem.(*schema.Resource).Schema)
		jmsPlugin["items"] = items
	}

	resources = append(resources, jmsPlugin)
	if err := s.D.Set("jms_plugin_collection", resources); err != nil {
		return err
	}

	return nil
}
