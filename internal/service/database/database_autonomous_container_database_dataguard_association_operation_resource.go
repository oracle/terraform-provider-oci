package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		CreateContext: createDatabaseAutonomousContainerDatabaseDataguardAssociationOperationWithContext,
		ReadContext:   readDatabaseAutonomousContainerDatabaseDataguardAssociationOperationWithContext,
		DeleteContext: deleteDatabaseAutonomousContainerDatabaseDataguardAssociationOperationWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"autonomous_container_database_dataguard_association_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"operation": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

type DatabaseAutonomousContainerDatabaseDataguardAssociationOperation struct {
	// The OCID of the Data Safe private endpoint.
	Id *string `mandatory:"true" json:"id"`
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud) ID() string {
	return fmt.Sprint(utils.GetStringHashcode(s.D.Get("autonomous_container_database_id").(string)))
}

func createDatabaseAutonomousContainerDatabaseDataguardAssociationOperationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDatabaseAutonomousContainerDatabaseDataguardAssociationOperationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func deleteDatabaseAutonomousContainerDatabaseDataguardAssociationOperationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud) GetWithContext(ctx context.Context) error {
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud) SetData() error {
	return nil
}

type DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	WorkRequestClient      *oci_work_requests.WorkRequestClient
	DisableNotFoundRetries bool
	Res                    *DatabaseAutonomousContainerDatabaseDataguardAssociationOperation
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud) CreateWithContext(ctx context.Context) error {
	return s.dataguardOperation(ctx)
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud) dataguardOperation(ctx context.Context) error {
	dataguardAssociationId := ""
	tmpId := ""
	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmpId = autonomousContainerDatabaseId.(string)
	}
	if autonomousContainerDatabaseDataguardAssociationId, ok := s.D.GetOkExists("autonomous_container_database_dataguard_association_id"); ok {
		dataguardAssociationId = autonomousContainerDatabaseDataguardAssociationId.(string)
	}
	if operation, ok := s.D.GetOkExists("operation"); ok {
		if strings.ToLower(operation.(string)) == "switchover" {
			switchoverRequest := oci_database.SwitchoverAutonomousContainerDatabaseDataguardAssociationRequest{}
			switchoverRequest.AutonomousContainerDatabaseDataguardAssociationId = &dataguardAssociationId
			switchoverRequest.AutonomousContainerDatabaseId = &tmpId
			switchoverRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
			//switchoverRequest.RequestMetadata.RetryPolicy.
			response, err := s.Client.SwitchoverAutonomousContainerDatabaseDataguardAssociation(ctx, switchoverRequest)
			if err != nil {
				return err
			}
			workId := response.OpcWorkRequestId
			if workId != nil {
				_, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
				if err != nil {
					return err
				}
			}
		}
		if strings.ToLower(operation.(string)) == "failover" {
			failoverRequest := oci_database.FailoverAutonomousContainerDatabaseDataguardAssociationRequest{}
			failoverRequest.AutonomousContainerDatabaseDataguardAssociationId = &dataguardAssociationId
			failoverRequest.AutonomousContainerDatabaseId = &tmpId
			failoverRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
			//failoverRequest.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 2
			response, err := s.Client.FailoverAutonomousContainerDatabaseDataguardAssociation(ctx, failoverRequest)
			if err != nil {
				return err
			}
			workId := response.OpcWorkRequestId
			if workId != nil {
				_, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
				if err != nil {
					return err
				}
			}
		}
		if strings.ToLower(operation.(string)) == "reinstate" {
			reinstateRequest := oci_database.ReinstateAutonomousContainerDatabaseDataguardAssociationRequest{}
			reinstateRequest.AutonomousContainerDatabaseDataguardAssociationId = &dataguardAssociationId
			reinstateRequest.AutonomousContainerDatabaseId = &tmpId
			reinstateRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")
			//reinstateRequest.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 2
			response, err := s.Client.ReinstateAutonomousContainerDatabaseDataguardAssociation(ctx, reinstateRequest)
			if err != nil {
				return err
			}
			workId := response.OpcWorkRequestId
			if workId != nil {
				_, err = tfresource.WaitForWorkRequestWithErrorHandlingAndContext(ctx, s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
