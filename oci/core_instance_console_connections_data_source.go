// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v46/core"
)

func init() {
	RegisterDatasource("oci_core_instance_console_connections", CoreInstanceConsoleConnectionsDataSource())
}

func CoreInstanceConsoleConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreInstanceConsoleConnections,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_console_connections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(CoreInstanceConsoleConnectionResource()),
			},
		},
	}
}

func readCoreInstanceConsoleConnections(d *schema.ResourceData, m interface{}) error {
	sync := &CoreInstanceConsoleConnectionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).computeClient()

	return ReadResource(sync)
}

type CoreInstanceConsoleConnectionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.ComputeClient
	Res    *oci_core.ListInstanceConsoleConnectionsResponse
}

func (s *CoreInstanceConsoleConnectionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreInstanceConsoleConnectionsDataSourceCrud) Get() error {
	request := oci_core.ListInstanceConsoleConnectionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if instanceId, ok := s.D.GetOkExists("instance_id"); ok {
		tmp := instanceId.(string)
		request.InstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListInstanceConsoleConnections(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListInstanceConsoleConnections(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreInstanceConsoleConnectionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("CoreInstanceConsoleConnectionsDataSource-", CoreInstanceConsoleConnectionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		instanceConsoleConnection := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ConnectionString != nil {
			instanceConsoleConnection["connection_string"] = *r.ConnectionString
		}

		if r.DefinedTags != nil {
			instanceConsoleConnection["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Fingerprint != nil {
			instanceConsoleConnection["fingerprint"] = *r.Fingerprint
		}

		instanceConsoleConnection["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			instanceConsoleConnection["id"] = *r.Id
		}

		if r.InstanceId != nil {
			instanceConsoleConnection["instance_id"] = *r.InstanceId
		}

		instanceConsoleConnection["state"] = r.LifecycleState

		if r.VncConnectionString != nil {
			instanceConsoleConnection["vnc_connection_string"] = *r.VncConnectionString
		}

		resources = append(resources, instanceConsoleConnection)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CoreInstanceConsoleConnectionsDataSource().Schema["instance_console_connections"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("instance_console_connections", resources); err != nil {
		return err
	}

	return nil
}
