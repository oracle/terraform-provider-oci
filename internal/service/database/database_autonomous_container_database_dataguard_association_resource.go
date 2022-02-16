// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v58/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v58/workrequests"
)

func DatabaseAutonomousContainerDatabaseDataguardAssociationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseAutonomousContainerDatabaseDataguardAssociation,
		Read:     readDatabaseAutonomousContainerDatabaseDataguardAssociation,
		Update:   updateDatabaseAutonomousContainerDatabaseDataguardAssociation,
		Delete:   deleteDatabaseAutonomousContainerDatabaseDataguardAssociation,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_container_database_dataguard_association_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"is_automatic_failover_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"apply_lag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apply_rate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_autonomous_container_database_dataguard_association_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_autonomous_container_database_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_lifecycle_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"protection_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
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
			"time_last_role_changed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_synced": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"transport_lag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseAutonomousContainerDatabaseDataguardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousContainerDatabaseDataguardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseAutonomousContainerDatabaseDataguardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseAutonomousContainerDatabaseDataguardAssociation(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.AutonomousContainerDatabaseDataguardAssociation
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) ID() string {
	return GetAutonomousContainerDatabaseDataguardAssociationCompositeId(s.D.Get("autonomous_container_database_dataguard_association_id").(string), s.D.Get("autonomous_container_database_id").(string))
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateProvisioning),
	}
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateAvailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminating),
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateUnavailable),
	}
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminated),
	}
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) Create() error {
	request := oci_database.UpdateAutonomousContainerDatabaseDataguardAssociationRequest{}

	if autonomousContainerDatabaseDataguardAssociationId, ok := s.D.GetOkExists("autonomous_container_database_dataguard_association_id"); ok {
		tmp := autonomousContainerDatabaseDataguardAssociationId.(string)
		request.AutonomousContainerDatabaseDataguardAssociationId = &tmp
	}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if isAutomaticFailoverEnabled, ok := s.D.GetOkExists("is_automatic_failover_enabled"); ok {
		tmp := isAutomaticFailoverEnabled.(bool)
		request.IsAutomaticFailoverEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousContainerDatabaseDataguardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.AutonomousContainerDatabaseDataguardAssociation

	if workId != nil {
		identifier, err := tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomousContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
		}
	}

	return s.Get()
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) Get() error {
	request := oci_database.GetAutonomousContainerDatabaseDataguardAssociationRequest{}

	if autonomousContainerDatabaseDataguardAssociationId, ok := s.D.GetOkExists("autonomous_container_database_dataguard_association_id"); ok {
		tmp := autonomousContainerDatabaseDataguardAssociationId.(string)
		request.AutonomousContainerDatabaseDataguardAssociationId = &tmp
	}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	autonomousContainerDatabaseDataguardAssociationId, autonomousContainerDatabaseId, err := parseAutonomousContainerDatabaseDataguardAssociationCompositeId(s.D.Id())
	if err == nil {
		request.AutonomousContainerDatabaseDataguardAssociationId = &autonomousContainerDatabaseDataguardAssociationId
		request.AutonomousContainerDatabaseId = &autonomousContainerDatabaseId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetAutonomousContainerDatabaseDataguardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AutonomousContainerDatabaseDataguardAssociation
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) Update() error {
	request := oci_database.UpdateAutonomousContainerDatabaseDataguardAssociationRequest{}

	if autonomousContainerDatabaseDataguardAssociationId, ok := s.D.GetOkExists("autonomous_container_database_dataguard_association_id"); ok {
		tmp := autonomousContainerDatabaseDataguardAssociationId.(string)
		request.AutonomousContainerDatabaseDataguardAssociationId = &tmp
	}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	if isAutomaticFailoverEnabled, ok := s.D.GetOkExists("is_automatic_failover_enabled"); ok {
		tmp := isAutomaticFailoverEnabled.(bool)
		request.IsAutomaticFailoverEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateAutonomousContainerDatabaseDataguardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomousContainerDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}

	return s.Get()
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationResourceCrud) SetData() error {

	autonomousContainerDatabaseDataguardAssociationId, autonomousContainerDatabaseId, err := parseAutonomousContainerDatabaseDataguardAssociationCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(autonomousContainerDatabaseDataguardAssociationId)
		s.D.Set("autonomous_container_database_id", &autonomousContainerDatabaseId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.ApplyLag != nil {
		s.D.Set("apply_lag", *s.Res.ApplyLag)
	}

	if s.Res.ApplyRate != nil {
		s.D.Set("apply_rate", *s.Res.ApplyRate)
	}

	if s.Res.AutonomousContainerDatabaseId != nil {
		s.D.Set("autonomous_container_database_id", *s.Res.AutonomousContainerDatabaseId)
	}

	if s.Res.IsAutomaticFailoverEnabled != nil {
		s.D.Set("is_automatic_failover_enabled", *s.Res.IsAutomaticFailoverEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PeerAutonomousContainerDatabaseDataguardAssociationId != nil {
		s.D.Set("peer_autonomous_container_database_dataguard_association_id", *s.Res.PeerAutonomousContainerDatabaseDataguardAssociationId)
	}

	if s.Res.PeerAutonomousContainerDatabaseId != nil {
		s.D.Set("peer_autonomous_container_database_id", *s.Res.PeerAutonomousContainerDatabaseId)
	}

	s.D.Set("peer_lifecycle_state", s.Res.PeerLifecycleState)

	s.D.Set("peer_role", s.Res.PeerRole)

	s.D.Set("protection_mode", s.Res.ProtectionMode)

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastRoleChanged != nil {
		s.D.Set("time_last_role_changed", s.Res.TimeLastRoleChanged.String())
	}

	if s.Res.TimeLastSynced != nil {
		s.D.Set("time_last_synced", s.Res.TimeLastSynced.String())
	}

	if s.Res.TransportLag != nil {
		s.D.Set("transport_lag", *s.Res.TransportLag)
	}

	return nil
}

func GetAutonomousContainerDatabaseDataguardAssociationCompositeId(autonomousContainerDatabaseDataguardAssociationId string, autonomousContainerDatabaseId string) string {
	autonomousContainerDatabaseDataguardAssociationId = url.PathEscape(autonomousContainerDatabaseDataguardAssociationId)
	autonomousContainerDatabaseId = url.PathEscape(autonomousContainerDatabaseId)
	compositeId := "autonomousContainerDatabases/" + autonomousContainerDatabaseId + "/autonomousContainerDatabaseDataguardAssociations/" + autonomousContainerDatabaseDataguardAssociationId
	return compositeId
}

func parseAutonomousContainerDatabaseDataguardAssociationCompositeId(compositeId string) (autonomousContainerDatabaseDataguardAssociationId string, autonomousContainerDatabaseId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("autonomousContainerDatabases/.*/autonomousContainerDatabaseDataguardAssociations/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	autonomousContainerDatabaseId, _ = url.PathUnescape(parts[1])
	autonomousContainerDatabaseDataguardAssociationId, _ = url.PathUnescape(parts[3])

	return
}
