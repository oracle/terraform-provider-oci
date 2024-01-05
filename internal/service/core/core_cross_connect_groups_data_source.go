// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreCrossConnectGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreCrossConnectGroups,
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
			"cross_connect_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreCrossConnectGroupResource()),
			},
		},
	}
}

func readCoreCrossConnectGroups(d *schema.ResourceData, m interface{}) error {
	sync := &CoreCrossConnectGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()

	return tfresource.ReadResource(sync)
}

type CoreCrossConnectGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListCrossConnectGroupsResponse
}

func (s *CoreCrossConnectGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreCrossConnectGroupsDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

func (s *CoreCrossConnectGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreCrossConnectGroupsDataSource-", CoreCrossConnectGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		crossConnectGroup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CustomerReferenceName != nil {
			crossConnectGroup["customer_reference_name"] = *r.CustomerReferenceName
		}

		if r.DefinedTags != nil {
			crossConnectGroup["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			crossConnectGroup["display_name"] = *r.DisplayName
		}

		crossConnectGroup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			crossConnectGroup["id"] = *r.Id
		}

		if r.MacsecProperties != nil {
			crossConnectGroup["macsec_properties"] = []interface{}{MacsecPropertiesToMap(r.MacsecProperties)}
		} else {
			crossConnectGroup["macsec_properties"] = nil
		}

		if r.OciLogicalDeviceName != nil {
			crossConnectGroup["oci_logical_device_name"] = *r.OciLogicalDeviceName
		}

		if r.OciPhysicalDeviceName != nil {
			crossConnectGroup["oci_physical_device_name"] = *r.OciPhysicalDeviceName
		}

		crossConnectGroup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			crossConnectGroup["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, crossConnectGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreCrossConnectGroupsDataSource().Schema["cross_connect_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cross_connect_groups", resources); err != nil {
		return err
	}

	return nil
}
