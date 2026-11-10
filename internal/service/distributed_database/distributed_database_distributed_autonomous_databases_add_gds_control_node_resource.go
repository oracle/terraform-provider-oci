// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		CreateContext: createDistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeWithContext,
		ReadContext:   readDistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeWithContext,
		DeleteContext: deleteDistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"distributed_autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public_ssh_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
		},
	}
}

func createDistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DistributedAutonomousDbServiceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient()

	if err := sync.CreateWithContext(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	d.SetId(distributedDatabaseActionID("DistributedAutonomousDatabaseGdsControlNodeAction-", d, "distributed_autonomous_database_id", "public_ssh_key", "subnet_id"))
	return nil
}

func readDistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeWithContext(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return nil
}

func deleteDistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeWithContext(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

type DistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_distributed_database.DistributedAutonomousDbServiceClient
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_distributed_database.DistributedDbWorkRequestServiceClient
}

func (s *DistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeResourceCrud) createAction(ctx context.Context) error {
	request := oci_distributed_database.AddDistributedAutonomousDatabaseGdsControlNodeRequest{}
	distributedAutonomousDatabaseID := s.D.Get("distributed_autonomous_database_id").(string)
	if err := waitForDistributedAutonomousDatabaseActionReady(ctx, s.D, s.Client, distributedAutonomousDatabaseID); err != nil {
		return err
	}

	request.DistributedAutonomousDatabaseId = &distributedAutonomousDatabaseID
	if publicSshKey, ok := s.D.GetOkExists("public_ssh_key"); ok {
		tmp := publicSshKey.(string)
		request.PublicSshKey = &tmp
	}
	if subnetID, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetID.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")
	response, err := s.Client.AddDistributedAutonomousDatabaseGdsControlNode(ctx, request)
	if err != nil {
		return err
	}

	workID := response.OpcWorkRequestId
	if workID == nil || strings.TrimSpace(*workID) == "" {
		return fmt.Errorf("missing opc-work-request-id for AddDistributedAutonomousDatabaseGdsControlNode")
	}
	if _, err := distributedAutonomousDatabaseWaitForWorkRequest(ctx, workID, "distributedautonomousdatabase", oci_distributed_database.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.WorkRequestClient); err != nil {
		return err
	}
	return waitForDistributedAutonomousDatabaseActionReady(ctx, s.D, s.Client, distributedAutonomousDatabaseID)
}

func (s *DistributedDatabaseDistributedAutonomousDatabasesAddGdsControlNodeResourceCrud) CreateWithContext(ctx context.Context) error {
	return s.createAction(ctx)
}
