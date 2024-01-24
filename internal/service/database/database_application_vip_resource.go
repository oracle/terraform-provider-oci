// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func DatabaseApplicationVipResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseApplicationVip,
		Read:     readDatabaseApplicationVip,
		Delete:   deleteDatabaseApplicationVip,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"hostname_label": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"db_node_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_assigned": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseApplicationVip(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseApplicationVipResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseApplicationVip(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseApplicationVipResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func deleteDatabaseApplicationVip(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseApplicationVipResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.DeleteResource(d, sync)
}

type DatabaseApplicationVipResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ApplicationVip
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseApplicationVipResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseApplicationVipResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.ApplicationVipLifecycleStateProvisioning),
	}
}

func (s *DatabaseApplicationVipResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.ApplicationVipLifecycleStateAvailable),
	}
}

func (s *DatabaseApplicationVipResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.ApplicationVipLifecycleStateTerminating),
	}
}

func (s *DatabaseApplicationVipResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.ApplicationVipLifecycleStateTerminated),
	}
}

func (s *DatabaseApplicationVipResourceCrud) Create() error {
	request := oci_database.CreateApplicationVipRequest{}

	if cloudVmClusterId, ok := s.D.GetOkExists("cloud_vm_cluster_id"); ok {
		tmp := cloudVmClusterId.(string)
		request.CloudVmClusterId = &tmp
	}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	if hostnameLabel, ok := s.D.GetOkExists("hostname_label"); ok {
		tmp := hostnameLabel.(string)
		request.HostnameLabel = &tmp
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateApplicationVip(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	if workId != nil {
		var identifier *string
		var err error
		identifier = response.Id
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "applicationvip", oci_work_requests.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	s.Res = &response.ApplicationVip
	return nil
}

func (s *DatabaseApplicationVipResourceCrud) Get() error {
	request := oci_database.GetApplicationVipRequest{}

	tmp := s.D.Id()
	request.ApplicationVipId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetApplicationVip(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApplicationVip
	return nil
}

func (s *DatabaseApplicationVipResourceCrud) Delete() error {
	request := oci_database.DeleteApplicationVipRequest{}

	tmp := s.D.Id()
	request.ApplicationVipId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DeleteApplicationVip(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "applicationVip", oci_work_requests.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DatabaseApplicationVipResourceCrud) SetData() error {
	if s.Res.CloudVmClusterId != nil {
		s.D.Set("cloud_vm_cluster_id", *s.Res.CloudVmClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostnameLabel != nil {
		s.D.Set("hostname_label", *s.Res.HostnameLabel)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeAssigned != nil {
		s.D.Set("time_assigned", s.Res.TimeAssigned.String())
	}

	return nil
}
