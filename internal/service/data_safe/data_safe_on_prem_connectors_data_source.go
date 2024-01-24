// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeOnPremConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeOnPremConnectors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"access_level": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"on_prem_connector_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"on_prem_connector_lifecycle_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"on_prem_connectors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DataSafeOnPremConnectorResource()),
			},
		},
	}
}

func readDataSafeOnPremConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeOnPremConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeOnPremConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListOnPremConnectorsResponse
}

func (s *DataSafeOnPremConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeOnPremConnectorsDataSourceCrud) Get() error {
	request := oci_data_safe.ListOnPremConnectorsRequest{}

	if accessLevel, ok := s.D.GetOkExists("access_level"); ok {
		request.AccessLevel = oci_data_safe.ListOnPremConnectorsAccessLevelEnum(accessLevel.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if onPremConnectorId, ok := s.D.GetOkExists("id"); ok {
		tmp := onPremConnectorId.(string)
		request.OnPremConnectorId = &tmp
	}

	if onPremConnectorLifecycleState, ok := s.D.GetOkExists("on_prem_connector_lifecycle_state"); ok {
		request.OnPremConnectorLifecycleState = oci_data_safe.ListOnPremConnectorsOnPremConnectorLifecycleStateEnum(onPremConnectorLifecycleState.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListOnPremConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOnPremConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeOnPremConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeOnPremConnectorsDataSource-", DataSafeOnPremConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		onPremConnector := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedVersion != nil {
			onPremConnector["created_version"] = *r.CreatedVersion
		}

		if r.DefinedTags != nil {
			onPremConnector["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			onPremConnector["description"] = *r.Description
		}

		if r.DisplayName != nil {
			onPremConnector["display_name"] = *r.DisplayName
		}

		onPremConnector["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			onPremConnector["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			onPremConnector["lifecycle_details"] = *r.LifecycleDetails
		}

		onPremConnector["state"] = r.LifecycleState

		if r.SystemTags != nil {
			onPremConnector["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			onPremConnector["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, onPremConnector)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DataSafeOnPremConnectorsDataSource().Schema["on_prem_connectors"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("on_prem_connectors", resources); err != nil {
		return err
	}

	return nil
}
