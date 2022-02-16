package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	oci_work_requests "github.com/oracle/oci-go-sdk/v58/workrequests"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		Create: createDatabaseAutonomousContainerDatabaseDataguardAssociationOperation,
		Read:   readDatabaseAutonomousContainerDatabaseDataguardAssociationOperation,
		Delete: deleteDatabaseAutonomousContainerDatabaseDataguardAssociationOperation,
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
	return fmt.Sprint(hashcode.String(s.D.Get("autonomous_container_database_id").(string)))
}

func createDatabaseAutonomousContainerDatabaseDataguardAssociationOperation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.CreateResource(d, sync)
}

func readDatabaseAutonomousContainerDatabaseDataguardAssociationOperation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WorkRequestClient

	return tfresource.ReadResource(sync)
}

func deleteDatabaseAutonomousContainerDatabaseDataguardAssociationOperation(d *schema.ResourceData, m interface{}) error {
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud) Get() error {
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

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud) Create() error {
	return s.dataguardOperation()
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResourceCrud) dataguardOperation() error {
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
			switchoverRequest.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 2
			response, err := s.Client.SwitchoverAutonomousContainerDatabaseDataguardAssociation(context.Background(), switchoverRequest)
			if err != nil {
				return err
			}
			workId := response.OpcWorkRequestId
			if workId != nil {
				_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
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
			failoverRequest.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 2
			response, err := s.Client.FailoverAutonomousContainerDatabaseDataguardAssociation(context.Background(), failoverRequest)
			if err != nil {
				return err
			}
			workId := response.OpcWorkRequestId
			if workId != nil {
				_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
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
			reinstateRequest.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 2
			response, err := s.Client.ReinstateAutonomousContainerDatabaseDataguardAssociation(context.Background(), reinstateRequest)
			if err != nil {
				return err
			}
			workId := response.OpcWorkRequestId
			if workId != nil {
				_, err = tfresource.WaitForWorkRequestWithErrorHandling(s.WorkRequestClient, workId, "database", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
