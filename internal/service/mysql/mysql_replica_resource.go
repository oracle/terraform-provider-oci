// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MysqlReplicaResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMysqlReplica,
		Read:     readMysqlReplica,
		Update:   updateMysqlReplica,
		Delete:   deleteMysqlReplica,
		Schema: map[string]*schema.Schema{
			// Required
			"db_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_delete_protected": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fault_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mysql_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"port_x": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMysqlReplica(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ReplicasClient()

	return tfresource.CreateResource(d, sync)
}

func readMysqlReplica(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ReplicasClient()

	return tfresource.ReadResource(sync)
}

func updateMysqlReplica(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ReplicasClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMysqlReplica(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlReplicaResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ReplicasClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MysqlReplicaResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_mysql.ReplicasClient
	Res                    *oci_mysql.Replica
	DisableNotFoundRetries bool
}

func (s *MysqlReplicaResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MysqlReplicaResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateCreating),
	}
}

func (s *MysqlReplicaResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateActive),
		string(oci_mysql.ReplicaLifecycleStateNeedsAttention),
	}
}

func (s *MysqlReplicaResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateDeleting),
	}
}

func (s *MysqlReplicaResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateDeleted),
	}
}

func (s *MysqlReplicaResourceCrud) UpdatedPending() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateUpdating),
	}
}

func (s *MysqlReplicaResourceCrud) UpdatedTarget() []string {
	return []string{
		string(oci_mysql.ReplicaLifecycleStateActive),
		string(oci_mysql.ReplicaLifecycleStateNeedsAttention),
	}
}

func (s *MysqlReplicaResourceCrud) Create() error {
	request := oci_mysql.CreateReplicaRequest{}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isDeleteProtected, ok := s.D.GetOkExists("is_delete_protected"); ok {
		tmp := isDeleteProtected.(bool)
		request.IsDeleteProtected = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.CreateReplica(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Replica
	return nil
}

func (s *MysqlReplicaResourceCrud) Get() error {
	request := oci_mysql.GetReplicaRequest{}

	tmp := s.D.Id()
	request.ReplicaId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	response, err := s.Client.GetReplica(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Replica
	return nil
}

func (s *MysqlReplicaResourceCrud) Update() error {
	request := oci_mysql.UpdateReplicaRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isDeleteProtected, ok := s.D.GetOkExists("is_delete_protected"); ok {
		tmp := isDeleteProtected.(bool)
		request.IsDeleteProtected = &tmp
	}

	tmp := s.D.Id()
	request.ReplicaId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.UpdateReplica(context.Background(), request)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *MysqlReplicaResourceCrud) Delete() error {
	request := oci_mysql.DeleteReplicaRequest{}

	tmp := s.D.Id()
	request.ReplicaId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "mysql")

	_, err := s.Client.DeleteReplica(context.Background(), request)
	return err
}

func (s *MysqlReplicaResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
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

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.IsDeleteProtected != nil {
		s.D.Set("is_delete_protected", *s.Res.IsDeleteProtected)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MysqlVersion != nil {
		s.D.Set("mysql_version", *s.Res.MysqlVersion)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.PortX != nil {
		s.D.Set("port_x", *s.Res.PortX)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
