// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseApplicationVipsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseApplicationVips,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cloud_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_vips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseApplicationVipResource()),
			},
		},
	}
}

func readDatabaseApplicationVips(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseApplicationVipsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseApplicationVipsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListApplicationVipsResponse
}

func (s *DatabaseApplicationVipsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseApplicationVipsDataSourceCrud) Get() error {
	request := oci_database.ListApplicationVipsRequest{}

	if cloudVmClusterId, ok := s.D.GetOkExists("cloud_vm_cluster_id"); ok {
		tmp := cloudVmClusterId.(string)
		request.CloudVmClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ApplicationVipSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListApplicationVips(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListApplicationVips(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseApplicationVipsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseApplicationVipsDataSource-", DatabaseApplicationVipsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		applicationVip := map[string]interface{}{
			"cloud_vm_cluster_id": *r.CloudVmClusterId,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			applicationVip["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		applicationVip["freeform_tags"] = r.FreeformTags

		if r.HostnameLabel != nil {
			applicationVip["hostname_label"] = *r.HostnameLabel
		}

		if r.Id != nil {
			applicationVip["id"] = *r.Id
		}

		if r.IpAddress != nil {
			applicationVip["ip_address"] = *r.IpAddress
		}

		if r.LifecycleDetails != nil {
			applicationVip["lifecycle_details"] = *r.LifecycleDetails
		}

		applicationVip["state"] = r.LifecycleState

		if r.SubnetId != nil {
			applicationVip["subnet_id"] = *r.SubnetId
		}

		if r.TimeAssigned != nil {
			applicationVip["time_assigned"] = r.TimeAssigned.String()
		}

		resources = append(resources, applicationVip)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseApplicationVipsDataSource().Schema["application_vips"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("application_vips", resources); err != nil {
		return err
	}

	return nil
}
