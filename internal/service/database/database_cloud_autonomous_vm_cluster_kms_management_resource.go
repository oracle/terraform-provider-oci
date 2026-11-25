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

func DatabaseCloudAutonomousVmClusterKMSManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDatabaseCloudAutonomousVmClusterKMSManagementWithContext,
		ReadContext:   readDatabaseCloudAutonomousVmClusterKMSManagementWithContext,
		DeleteContext: deleteDatabaseCloudAutonomousVmClusterKMSManagementWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"cloud_autonomous_vm_cluster_id": {
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

func createDatabaseCloudAutonomousVmClusterKMSManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseCloudAutonomousVmClusterKMSManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseCloudAutonomousVmClusterKMSManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func deleteDatabaseCloudAutonomousVmClusterKMSManagementWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

type DatabaseCloudAutonomousVmClusterKMSManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.ChangeCavmKeyManagementSystemResponse
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_work_requests.WorkRequestClient
}

func (s *DatabaseCloudAutonomousVmClusterKMSManagementResourceCrud) ID() string {
	return tfresource.GenerateDataSourceHashID("DatabaseCloudAutonomousVmClusterKMSManagementResource-", DatabaseCloudAutonomousVmClusterKMSManagementResource(), s.D)
}

func (s *DatabaseCloudAutonomousVmClusterKMSManagementResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_database.ChangeCavmKeyManagementSystemRequest{}
	err := s.populateTopLevelPolymorphicChangeCavmKeyManagementSystemRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.ChangeCavmKeyManagementSystem(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response

	if workId != nil {
		var identifier *string
		var err error
		identifier, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "cloudautonomousvmcluster", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if identifier != nil {
			s.D.SetId(*identifier)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatabaseCloudAutonomousVmClusterKMSManagementResourceCrud) SetData() error {
	return nil
}

func (s *DatabaseCloudAutonomousVmClusterKMSManagementResourceCrud) populateTopLevelPolymorphicChangeCavmKeyManagementSystemRequest(request *oci_database.ChangeCavmKeyManagementSystemRequest) error {
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
		if cloudAutonomousVmClusterId, ok := s.D.GetOkExists("cloud_autonomous_vm_cluster_id"); ok {
			tmp := cloudAutonomousVmClusterId.(string)
			request.CloudAutonomousVmClusterId = &tmp
		}
		if isExclusive, ok := s.D.GetOkExists("is_exclusive"); ok {
			tmp := isExclusive.(bool)
			details.IsExclusive = &tmp
		}
		request.ChangeCavmKeyManagementSystemDetails = details
	case strings.ToLower("ORACLE_KEY_VAULT"):
		details := oci_database.OracleKeyVaultKeyManagementSystemDetails{}
		if cloudAutonomousVmClusterId, ok := s.D.GetOkExists("cloud_autonomous_vm_cluster_id"); ok {
			tmp := cloudAutonomousVmClusterId.(string)
			request.CloudAutonomousVmClusterId = &tmp
		}
		if isExclusive, ok := s.D.GetOkExists("is_exclusive"); ok {
			tmp := isExclusive.(bool)
			details.IsExclusive = &tmp
		}
		request.ChangeCavmKeyManagementSystemDetails = details
	case strings.ToLower("ORACLE_MANAGED_KEYS"):
		details := oci_database.OracleManagedKeyManagementSystemDetails{}
		if cloudAutonomousVmClusterId, ok := s.D.GetOkExists("cloud_autonomous_vm_cluster_id"); ok {
			tmp := cloudAutonomousVmClusterId.(string)
			request.CloudAutonomousVmClusterId = &tmp
		}
		if isExclusive, ok := s.D.GetOkExists("is_exclusive"); ok {
			tmp := isExclusive.(bool)
			details.IsExclusive = &tmp
		}
		request.ChangeCavmKeyManagementSystemDetails = details
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
		if cloudAutonomousVmClusterId, ok := s.D.GetOkExists("cloud_autonomous_vm_cluster_id"); ok {
			tmp := cloudAutonomousVmClusterId.(string)
			request.CloudAutonomousVmClusterId = &tmp
		}
		if isExclusive, ok := s.D.GetOkExists("is_exclusive"); ok {
			tmp := isExclusive.(bool)
			details.IsExclusive = &tmp
		}
		request.ChangeCavmKeyManagementSystemDetails = details
	default:
		return fmt.Errorf("unknown key_management_system '%v' was specified", keyManagementSystem)
	}
	return nil
}

func autonomousVmClusterKmsManagementSystemDiffSuppress(key string, old string, new string, d *schema.ResourceData) bool {
	return tfresource.EqualIgnoreCaseSuppressDiff(key, old, new, d) ||
		old == new
}
