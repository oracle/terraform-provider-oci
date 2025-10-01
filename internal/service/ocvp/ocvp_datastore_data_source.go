// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpDatastoreDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["datastore_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OcvpDatastoreResource(), fieldMap, readSingularOcvpDatastore)
}

func readSingularOcvpDatastore(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpDatastoreDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatastoreClient()

	return tfresource.ReadResource(sync)
}

type OcvpDatastoreDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.DatastoreClient
	Res    *oci_ocvp.GetDatastoreResponse
}

func (s *OcvpDatastoreDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpDatastoreDataSourceCrud) Get() error {
	request := oci_ocvp.GetDatastoreRequest{}

	if datastoreId, ok := s.D.GetOkExists("datastore_id"); ok {
		tmp := datastoreId.(string)
		request.DatastoreId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.GetDatastore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpDatastoreDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	blockVolumeDetails := []interface{}{}
	for _, item := range s.Res.BlockVolumeDetails {
		blockVolumeDetails = append(blockVolumeDetails, BlockVolumeDetailsToMap(item))
	}
	s.D.Set("block_volume_details", blockVolumeDetails)

	s.D.Set("block_volume_ids", s.Res.BlockVolumeIds)

	if s.Res.CapacityInGBs != nil {
		s.D.Set("capacity_in_gbs", *s.Res.CapacityInGBs)
	}

	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.SddcId != nil {
		s.D.Set("sddc_id", *s.Res.SddcId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
