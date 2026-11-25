// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousVmClusterKMSManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseAutonomousVmClusterKMSManagementWithContext,
		ReadContext:   readDatabaseAutonomousVmClusterKMSManagementWithContext,
		DeleteContext: deleteDatabaseAutonomousVmClusterKMSManagementWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_management_system": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: autonomousVmClusterKmsManagementSystemDiffSuppress,
				ValidateFunc: validation.StringInSlice([]string{
					"OCI_VAULT",
					"ORACLE_KEY_VAULT",
					"ORACLE_MANAGED_KEYS",
					"THALES",
				}, true),
			},

			// Optional
			"is_exclusive": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"key_store_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"thales_hsm_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createDatabaseAutonomousVmClusterKMSManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseAutonomousVmClusterKMSManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseAutonomousVmClusterKMSManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteDatabaseAutonomousVmClusterKMSManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type DatabaseAutonomousVmClusterKMSManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ChangeAvmKeyManagementSystemResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseAutonomousVmClusterKMSManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseAutonomousVmClusterKMSManagementResource-", DatabaseAutonomousVmClusterKMSManagementResource(), s.D)
}

func (s *DatabaseAutonomousVmClusterKMSManagementResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database.ChangeAvmKeyManagementSystemRequest{}
	err := s.populateTopLevelPolymorphicChangeAvmKeyManagementSystemRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeAvmKeyManagementSystem(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "autonomousvmcluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseAutonomousVmClusterKMSManagementResourceCrud) SetData() error {
	return nil
}

func (s *DatabaseAutonomousVmClusterKMSManagementResourceCrud) populateTopLevelPolymorphicChangeAvmKeyManagementSystemRequest(request *oci_database.ChangeAvmKeyManagementSystemRequest) error {
	//discriminator
	keyManagementSystemRaw, ok := s.D.GetOkExists("key_management_system")
	var keyManagementSystem string
	if ok {
		keyManagementSystem = keyManagementSystemRaw.(string)
	} else {
		keyManagementSystem = "" // default value
	}
	switch strings.ToLower(keyManagementSystem) {
	case strings.ToLower("OCI_VAULT"):
		details := oci_database.OciVaultKeyManagementSystemDetails{}
		if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
			tmp := autonomousVmClusterId.(string)
			request.AutonomousVmClusterId = &tmp
		}
		if isExclusive, ok := s.D.GetOkExists("is_exclusive"); ok {
			tmp := isExclusive.(bool)
			details.IsExclusive = &tmp
		}
		request.ChangeAvmKeyManagementSystemDetails = details
	case strings.ToLower("ORACLE_KEY_VAULT"):
		details := oci_database.OracleKeyVaultKeyManagementSystemDetails{}
		if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
			tmp := autonomousVmClusterId.(string)
			request.AutonomousVmClusterId = &tmp
		}
		if isExclusive, ok := s.D.GetOkExists("is_exclusive"); ok {
			tmp := isExclusive.(bool)
			details.IsExclusive = &tmp
		}
		request.ChangeAvmKeyManagementSystemDetails = details
	case strings.ToLower("ORACLE_MANAGED_KEYS"):
		details := oci_database.OracleManagedKeyManagementSystemDetails{}
		if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
			tmp := autonomousVmClusterId.(string)
			request.AutonomousVmClusterId = &tmp
		}
		if isExclusive, ok := s.D.GetOkExists("is_exclusive"); ok {
			tmp := isExclusive.(bool)
			details.IsExclusive = &tmp
		}
		request.ChangeAvmKeyManagementSystemDetails = details
	case strings.ToLower("THALES"):
		details := oci_database.ThalesKeyManagementSystemDetails{}
		if keyStoreId, ok := s.D.GetOkExists("key_store_id"); ok {
			tmp := keyStoreId.(string)
			details.KeyStoreId = &tmp
		}
		if thalesHsmPort, ok := s.D.GetOkExists("thales_hsm_port"); ok {
			tmp := thalesHsmPort.(int)
			details.ThalesHsmPort = &tmp
		}
		if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
			tmp := autonomousVmClusterId.(string)
			request.AutonomousVmClusterId = &tmp
		}
		if isExclusive, ok := s.D.GetOkExists("is_exclusive"); ok {
			tmp := isExclusive.(bool)
			details.IsExclusive = &tmp
		}
		request.ChangeAvmKeyManagementSystemDetails = details
	default:
		return fmt.Errorf("unknown key_management_system '%v' was specified", keyManagementSystem)
	}
	return nil
}
