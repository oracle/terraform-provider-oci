// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bastion

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_bastion "github.com/oracle/oci-go-sdk/v58/bastion"
)

func BastionSessionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBastionSessions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bastion_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sessions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(BastionSessionResource()),
			},
		},
	}
}

func readBastionSessions(d *schema.ResourceData, m interface{}) error {
	sync := &BastionSessionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()

	return tfresource.ReadResource(sync)
}

type BastionSessionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bastion.BastionClient
	Res    *oci_bastion.ListSessionsResponse
}

func (s *BastionSessionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BastionSessionsDataSourceCrud) Get() error {
	request := oci_bastion.ListSessionsRequest{}

	if bastionId, ok := s.D.GetOkExists("bastion_id"); ok {
		tmp := bastionId.(string)
		request.BastionId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if sessionId, ok := s.D.GetOkExists("id"); ok {
		tmp := sessionId.(string)
		request.SessionId = &tmp
	}

	if sessionLifecycleState, ok := s.D.GetOkExists("session_lifecycle_state"); ok {
		request.SessionLifecycleState = oci_bastion.ListSessionsSessionLifecycleStateEnum(sessionLifecycleState.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bastion")

	response, err := s.Client.ListSessions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSessions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BastionSessionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BastionSessionsDataSource-", BastionSessionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		session := map[string]interface{}{
			"bastion_id": *r.BastionId,
		}

		if r.BastionName != nil {
			session["bastion_name"] = *r.BastionName
		}

		if r.DisplayName != nil {
			session["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			session["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			session["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.SessionTtlInSeconds != nil {
			session["session_ttl_in_seconds"] = *r.SessionTtlInSeconds
		}

		session["state"] = r.LifecycleState

		if r.TargetResourceDetails != nil {
			targetResourceDetailsArray := []interface{}{}
			if targetResourceDetailsMap := TargetResourceDetailsToMap(&r.TargetResourceDetails); targetResourceDetailsMap != nil {
				targetResourceDetailsArray = append(targetResourceDetailsArray, targetResourceDetailsMap)
			}
			session["target_resource_details"] = targetResourceDetailsArray
		} else {
			session["target_resource_details"] = nil
		}

		if r.TimeCreated != nil {
			session["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			session["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, session)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, BastionSessionsDataSource().Schema["sessions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("sessions", resources); err != nil {
		return err
	}

	return nil
}
