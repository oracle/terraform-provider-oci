// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package logging

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_logging "github.com/oracle/oci-go-sdk/v65/logging"
)

func LoggingLogGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoggingLogGroups,
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
			"is_compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"log_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(LoggingLogGroupResource()),
			},
		},
	}
}

func readLoggingLogGroups(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingLogGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.ReadResource(sync)
}

type LoggingLogGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_logging.LoggingManagementClient
	Res    *oci_logging.ListLogGroupsResponse
}

func (s *LoggingLogGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoggingLogGroupsDataSourceCrud) Get() error {
	request := oci_logging.ListLogGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isCompartmentIdInSubtree, ok := s.D.GetOkExists("is_compartment_id_in_subtree"); ok {
		tmp := isCompartmentIdInSubtree.(bool)
		request.IsCompartmentIdInSubtree = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "logging")

	response, err := s.Client.ListLogGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLogGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LoggingLogGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoggingLogGroupsDataSource-", LoggingLogGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		logGroup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			logGroup["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			logGroup["description"] = *r.Description
		}

		if r.DisplayName != nil {
			logGroup["display_name"] = *r.DisplayName
		}

		logGroup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			logGroup["id"] = *r.Id
		}

		logGroup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			logGroup["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastModified != nil {
			logGroup["time_last_modified"] = r.TimeLastModified.String()
		}

		resources = append(resources, logGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LoggingLogGroupsDataSource().Schema["log_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("log_groups", resources); err != nil {
		return err
	}

	return nil
}
