// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CrossConnectGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCrossConnectGroups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
			"cross_connect_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CrossConnectGroupResource()),
			},
		},
	}
}

func readCrossConnectGroups(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return ReadResource(sync)
}

type CrossConnectGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListCrossConnectGroupsResponse
}

func (s *CrossConnectGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CrossConnectGroupsDataSourceCrud) Get() error {
	request := oci_core.ListCrossConnectGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.CrossConnectGroupLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListCrossConnectGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCrossConnectGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CrossConnectGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		crossConnectGroup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DisplayName != nil {
			crossConnectGroup["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			crossConnectGroup["id"] = *r.Id
		}

		crossConnectGroup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			crossConnectGroup["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, crossConnectGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CrossConnectGroupsDataSource().Schema["cross_connect_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cross_connect_groups", resources); err != nil {
		return err
	}

	return nil
}
