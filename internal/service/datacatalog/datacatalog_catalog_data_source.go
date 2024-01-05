// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacatalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v65/datacatalog"
	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatacatalogCatalogDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["catalog_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatacatalogCatalogResource(), fieldMap, readSingularDatacatalogCatalog)
}

func readSingularDatacatalogCatalog(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogCatalogDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataCatalogClient()

	return tfresource.ReadResource(sync)
}

type DatacatalogCatalogDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacatalog.DataCatalogClient
	Res    *oci_datacatalog.GetCatalogResponse
}

func (s *DatacatalogCatalogDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatacatalogCatalogDataSourceCrud) Get() error {
	request := oci_datacatalog.GetCatalogRequest{}

	if catalogId, ok := s.D.GetOkExists("catalog_id"); ok {
		tmp := catalogId.(string)
		request.CatalogId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacatalog")

	response, err := s.Client.GetCatalog(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatacatalogCatalogDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("attached_catalog_private_endpoints", s.Res.AttachedCatalogPrivateEndpoints)

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMapCatalog(item))
	}
	s.D.Set("locks", locks)

	if s.Res.NumberOfObjects != nil {
		s.D.Set("number_of_objects", *s.Res.NumberOfObjects)
	}

	if s.Res.ServiceApiUrl != nil {
		s.D.Set("service_api_url", *s.Res.ServiceApiUrl)
	}

	if s.Res.ServiceConsoleUrl != nil {
		s.D.Set("service_console_url", *s.Res.ServiceConsoleUrl)
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
