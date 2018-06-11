// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_file_storage "github.com/oracle/oci-go-sdk/filestorage"
)

func MountTargetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createMountTarget,
		Read:     readMountTarget,
		Update:   updateMountTarget,
		Delete:   deleteMountTarget,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname_label": {
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
			"export_set_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMountTarget(d *schema.ResourceData, m interface{}) error {
	sync := &MountTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.CreateResource(d, sync)
}

func readMountTarget(d *schema.ResourceData, m interface{}) error {
	sync := &MountTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.ReadResource(sync)
}

func updateMountTarget(d *schema.ResourceData, m interface{}) error {
	sync := &MountTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient

	return crud.UpdateResource(d, sync)
}

func deleteMountTarget(d *schema.ResourceData, m interface{}) error {
	sync := &MountTargetResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type MountTargetResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.MountTarget
	DisableNotFoundRetries bool
}

func (s *MountTargetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MountTargetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_file_storage.MountTargetLifecycleStateCreating),
	}
}

func (s *MountTargetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_file_storage.MountTargetLifecycleStateActive),
		string(oci_file_storage.MountTargetLifecycleStateFailed),
	}
}

func (s *MountTargetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_file_storage.MountTargetLifecycleStateDeleting),
	}
}

func (s *MountTargetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_file_storage.MountTargetLifecycleStateDeleted),
		string(oci_file_storage.MountTargetLifecycleStateFailed),
	}
}

func (s *MountTargetResourceCrud) Create() error {
	request := oci_file_storage.CreateMountTargetRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateMountTarget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MountTarget
	return nil
}

func (s *MountTargetResourceCrud) Get() error {
	request := oci_file_storage.GetMountTargetRequest{}

	tmp := s.D.Id()
	request.MountTargetId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetMountTarget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MountTarget
	return nil
}

func (s *MountTargetResourceCrud) Update() error {
	request := oci_file_storage.UpdateMountTargetRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.MountTargetId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateMountTarget(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MountTarget
	return nil
}

func (s *MountTargetResourceCrud) Delete() error {
	request := oci_file_storage.DeleteMountTargetRequest{}

	tmp := s.D.Id()
	request.MountTargetId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteMountTarget(context.Background(), request)
	return err
}

func (s *MountTargetResourceCrud) SetData() {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExportSetId != nil {
		s.D.Set("export_set_id", *s.Res.ExportSetId)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("private_ip_ids", s.Res.PrivateIpIds)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
