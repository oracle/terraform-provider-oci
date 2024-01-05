// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/common"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousDatabaseSaasAdminUserResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseAutonomousDatabaseSaasAdminUser,
		Read:     readDatabaseAutonomousDatabaseSaasAdminUser,
		Delete:   deleteDatabaseAutonomousDatabaseSaasAdminUser,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"access_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"secret_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"secret_version_number": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			// Computed
			"time_saas_admin_user_enabled": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
		},
	}
}

type SaasAdminUser struct {
	AutonomousDatabaseId     *string
	Password                 *string
	SecretId                 *string
	SecretVersionNumber      *int
	AccessType               oci_database.ConfigureSaasAdminUserDetailsAccessTypeEnum
	Duration                 *int
	TimeSaasAdminUserEnabled *common.SDKTime
}

func createDatabaseAutonomousDatabaseSaasAdminUser(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseSaasAdminUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousDatabaseSaasAdminUser(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseSaasAdminUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func deleteDatabaseAutonomousDatabaseSaasAdminUser(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseSaasAdminUserResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseAutonomousDatabaseSaasAdminUserResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *SaasAdminUser
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	DisableNotFoundRetries bool
}

func (s *DatabaseAutonomousDatabaseSaasAdminUserResourceCrud) ID() string {
	return "saas-admin-user-admin-user-" + *s.Res.AutonomousDatabaseId
}

func (s *DatabaseAutonomousDatabaseSaasAdminUserResourceCrud) Create() error {
	request := oci_database.ConfigureSaasAdminUserRequest{}

	enabledTrue := true
	request.IsEnabled = &enabledTrue

	if accessType, ok := s.D.GetOkExists("access_type"); ok {
		tmp := accessType.(string)
		request.AccessType, _ = oci_database.GetMappingConfigureSaasAdminUserDetailsAccessTypeEnum(tmp)
	}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	if duration, ok := s.D.GetOkExists("duration"); ok {
		tmp := duration.(int)
		request.Duration = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok {
		tmp := secretId.(string)
		request.SecretId = &tmp
	}

	if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
		tmp := secretVersionNumber.(int)
		request.SecretVersionNumber = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ConfigureSaasAdminUser(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomousDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	saasAdminUser := SaasAdminUser{
		AutonomousDatabaseId: response.AutonomousDatabase.Id,
		Password:             request.Password,
		SecretId:             request.SecretId,
		SecretVersionNumber:  request.SecretVersionNumber,
		AccessType:           request.AccessType,
		Duration:             request.Duration,
	}
	s.Res = &saasAdminUser

	return s.Get()
}

func (s *DatabaseAutonomousDatabaseSaasAdminUserResourceCrud) Get() error {
	request := oci_database.SaasAdminUserStatusRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.SaasAdminUserStatus(context.Background(), request)
	if err != nil {
		return err
	}

	if response.IsEnabled == nil || *response.IsEnabled == false {
		return nil
	}

	saasAdminUser := SaasAdminUser{}
	saasAdminUser.TimeSaasAdminUserEnabled = response.TimeSaasAdminUserEnabled

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		saasAdminUser.AutonomousDatabaseId = &tmp
	}

	if accessType, ok := s.D.GetOkExists("access_type"); ok {
		tmp := accessType.(string)
		saasAdminUser.AccessType, _ = oci_database.GetMappingConfigureSaasAdminUserDetailsAccessTypeEnum(tmp)
	}

	if duration, ok := s.D.GetOkExists("duration"); ok {
		tmp := duration.(int)
		saasAdminUser.Duration = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		saasAdminUser.Password = &tmp
	}

	if secretId, ok := s.D.GetOkExists("secret_id"); ok {
		tmp := secretId.(string)
		saasAdminUser.SecretId = &tmp
	}

	if secretVersionNumber, ok := s.D.GetOkExists("secret_version_number"); ok {
		tmp := secretVersionNumber.(int)
		saasAdminUser.SecretVersionNumber = &tmp
	}

	s.Res = &saasAdminUser

	return nil
}

func (s *DatabaseAutonomousDatabaseSaasAdminUserResourceCrud) Delete() error {
	request := oci_database.ConfigureSaasAdminUserRequest{}

	tmp := s.D.Get("autonomous_database_id").(string)
	disabled := false
	request.AutonomousDatabaseId = &tmp
	request.IsEnabled = &disabled

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ConfigureSaasAdminUser(context.Background(), request)
	workId := response.OpcWorkRequestId
	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "autonomousDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}

	return err
}

func (s *DatabaseAutonomousDatabaseSaasAdminUserResourceCrud) SetData() error {
	if s.Res.TimeSaasAdminUserEnabled != nil {
		s.D.Set("time_saas_admin_user_enabled", s.Res.TimeSaasAdminUserEnabled.String())
	}

	return nil
}
