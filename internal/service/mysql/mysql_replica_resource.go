// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"
	"fmt"

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
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
		},
		Create: createMysqlReplica,
		Read:   readMysqlReplica,
		Update: updateMysqlReplica,
		Delete: deleteMysqlReplica,
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
			"replica_overrides": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"configuration_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mysql_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"shape_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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
			"configuration_id": {
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
			"secure_connections": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"certificate_generation_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"shape_name": {
				Type:     schema.TypeString,
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

	if replicaOverrides, ok := s.D.GetOkExists("replica_overrides"); ok {
		if tmpList := replicaOverrides.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replica_overrides", 0)
			tmp, err := s.mapToReplicaOverrides(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ReplicaOverrides = &tmp
		}
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

	if replicaOverrides, ok := s.D.GetOkExists("replica_overrides"); ok {
		if tmpList := replicaOverrides.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "replica_overrides", 0)
			tmp, err := s.mapToReplicaOverrides(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ReplicaOverrides = &tmp
		}
	}

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

	if s.Res.ConfigurationId != nil {
		s.D.Set("configuration_id", *s.Res.ConfigurationId)
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

	if s.Res.ReplicaOverrides != nil {
		s.D.Set("replica_overrides", []interface{}{ReplicaOverridesToMap(s.Res.ReplicaOverrides)})
	} else {
		s.D.Set("replica_overrides", nil)
	}

	if s.Res.SecureConnections != nil {
		s.D.Set("secure_connections", []interface{}{SecureConnectionDetailsToMap(s.Res.SecureConnections)})
	} else {
		s.D.Set("secure_connections", nil)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape_name", *s.Res.ShapeName)
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

func (s *MysqlReplicaResourceCrud) mapToReplicaOverrides(fieldKeyFormat string) (oci_mysql.ReplicaOverrides, error) {
	result := oci_mysql.ReplicaOverrides{}

	configurationIdField := fmt.Sprintf(fieldKeyFormat, "configuration_id")
	if configurationId, ok := s.D.GetOkExists(configurationIdField); ok && s.D.HasChange(configurationIdField) {
		tmp := configurationId.(string)
		result.ConfigurationId = &tmp
	}

	mysqlVersionField := fmt.Sprintf(fieldKeyFormat, "mysql_version")
	if mysqlVersion, ok := s.D.GetOkExists(mysqlVersionField); ok && s.D.HasChange(mysqlVersionField) {
		tmp := mysqlVersion.(string)
		result.MysqlVersion = &tmp
	}

	shapeNameField := fmt.Sprintf(fieldKeyFormat, "shape_name")
	if shapeName, ok := s.D.GetOkExists(shapeNameField); ok && s.D.HasChange(shapeNameField) {
		tmp := shapeName.(string)
		result.ShapeName = &tmp
	}

	return result, nil
}

func ReplicaOverridesToMap(obj *oci_mysql.ReplicaOverrides) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigurationId != nil {
		result["configuration_id"] = string(*obj.ConfigurationId)
	}

	if obj.MysqlVersion != nil {
		result["mysql_version"] = string(*obj.MysqlVersion)
	}

	if obj.ShapeName != nil {
		result["shape_name"] = string(*obj.ShapeName)
	}

	return result
}
