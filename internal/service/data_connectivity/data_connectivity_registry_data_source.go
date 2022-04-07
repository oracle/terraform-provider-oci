// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_connectivity

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v65/dataconnectivity"
)

func DataConnectivityRegistryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["registry_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DataConnectivityRegistryResource(), fieldMap, readSingularDataConnectivityRegistry)
}

func readSingularDataConnectivityRegistry(d *schema.ResourceData, m interface{}) error {
	sync := &DataConnectivityRegistryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataConnectivityManagementClient()

	return tfresource.ReadResource(sync)
}

type DataConnectivityRegistryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_connectivity.DataConnectivityManagementClient
	Res    *oci_data_connectivity.GetRegistryResponse
}

func (s *DataConnectivityRegistryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataConnectivityRegistryDataSourceCrud) Get() error {
	request := oci_data_connectivity.GetRegistryRequest{}

	if registryId, ok := s.D.GetOkExists("registry_id"); ok {
		tmp := registryId.(string)
		request.RegistryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_connectivity")

	response, err := s.Client.GetRegistry(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataConnectivityRegistryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StateMessage != nil {
		s.D.Set("state_message", *s.Res.StateMessage)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UpdatedBy != nil {
		s.D.Set("updated_by", *s.Res.UpdatedBy)
	}

	return nil
}
