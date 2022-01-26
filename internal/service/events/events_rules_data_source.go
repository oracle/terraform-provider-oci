// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package events

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_events "github.com/oracle/oci-go-sdk/v56/events"
)

func EventsRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readEventsRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(EventsRuleResource()),
			},
		},
	}
}

func readEventsRules(d *schema.ResourceData, m interface{}) error {
	sync := &EventsRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EventsClient()

	return tfresource.ReadResource(sync)
}

type EventsRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_events.EventsClient
	Res    *oci_events.ListRulesResponse
}

func (s *EventsRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EventsRulesDataSourceCrud) Get() error {
	request := oci_events.ListRulesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_events.RuleLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "events")

	response, err := s.Client.ListRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *EventsRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("EventsRulesDataSource-", EventsRulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		rule := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.Condition != nil {
			rule["condition"] = *r.Condition
		}

		if r.DefinedTags != nil {
			rule["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			rule["description"] = *r.Description
		}

		if r.DisplayName != nil {
			rule["display_name"] = *r.DisplayName
		}

		rule["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			rule["id"] = *r.Id
		}

		if r.IsEnabled != nil {
			rule["is_enabled"] = *r.IsEnabled
		}

		rule["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			rule["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, rule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, EventsRulesDataSource().Schema["rules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("rules", resources); err != nil {
		return err
	}

	return nil
}
