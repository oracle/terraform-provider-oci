// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v36/datasafe"
)

func init() {
	RegisterDatasource("oci_data_safe_on_prem_connectors", DataSafeOnPremConnectorsDataSource())
}

func DataSafeOnPremConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeOnPremConnectors,
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
				Elem:     GetDataSourceItemSchema(DataSafeOnPremConnectorResource()),
			},
		},
	}
}

func readDataSafeOnPremConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeOnPremConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataSafeClient()

	return ReadResource(sync)
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

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "data_safe")

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

	s.D.SetId(GenerateDataSourceHashID("DataSafeOnPremConnectorsDataSource-", DataSafeOnPremConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		onPremConnector := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedVersion != nil {
			onPremConnector["created_version"] = *r.CreatedVersion
		}

		if r.DefinedTags != nil {
			onPremConnector["defined_tags"] = definedTagsToMap(r.DefinedTags)
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

		if r.TimeCreated != nil {
			onPremConnector["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, onPremConnector)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DataSafeOnPremConnectorsDataSource().Schema["on_prem_connectors"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("on_prem_connectors", resources); err != nil {
		return err
	}

	return nil
}
