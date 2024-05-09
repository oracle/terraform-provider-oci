// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubEventsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOsManagementHubEvents,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_summary": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_summary_contains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_managed_by_autonomous_linux": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_created_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"event_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(OsManagementHubEventResource()),
						},
					},
				},
			},
		},
	}
}

func readOsManagementHubEvents(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubEventsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OsmhEventClient()

	return tfresource.ReadResource(sync)
}

type OsManagementHubEventsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_os_management_hub.EventClient
	Res    *oci_os_management_hub.ListEventsResponse
}

func (s *OsManagementHubEventsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OsManagementHubEventsDataSourceCrud) Get() error {
	request := oci_os_management_hub.ListEventsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if eventFingerprint, ok := s.D.GetOkExists("event_fingerprint"); ok {
		tmp := eventFingerprint.(string)
		request.EventFingerprint = &tmp
	}

	if eventSummary, ok := s.D.GetOkExists("event_summary"); ok {
		tmp := eventSummary.(string)
		request.EventSummary = &tmp
	}

	if eventSummaryContains, ok := s.D.GetOkExists("event_summary_contains"); ok {
		tmp := eventSummaryContains.(string)
		request.EventSummaryContains = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if isManagedByAutonomousLinux, ok := s.D.GetOkExists("is_managed_by_autonomous_linux"); ok {
		tmp := isManagedByAutonomousLinux.(bool)
		request.IsManagedByAutonomousLinux = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_os_management_hub.EventLifecycleStateEnum(state.(string))
	}

	if timeCreatedGreaterThanOrEqualTo, ok := s.D.GetOkExists("time_created_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timeCreatedLessThan, ok := s.D.GetOkExists("time_created_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreatedLessThan.(string))
		if err != nil {
			return err
		}
		request.TimeCreatedLessThan = &oci_common.SDKTime{Time: tmp}
	}

	// TODO: figure out how to support inherited field
	if type_, ok := s.D.GetOkExists("type"); ok {
		interfaces := type_.([]interface{})
		tmp := make([]oci_os_management_hub.EventTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.EventTypeEnum(interfaces[i].(string))
			}
		}
		request.Type = tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "os_management_hub")

	response, err := s.Client.ListEvents(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEvents(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OsManagementHubEventsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OsManagementHubEventsDataSource-", OsManagementHubEventsDataSource(), s.D))
	resources := []map[string]interface{}{}
	event := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EventSummaryToMap(item))
	}
	event["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OsManagementHubEventsDataSource().Schema["event_collection"].Elem.(*schema.Resource).Schema)
		event["items"] = items
	}

	resources = append(resources, event)
	if err := s.D.Set("event_collection", resources); err != nil {
		return err
	}

	return nil
}
