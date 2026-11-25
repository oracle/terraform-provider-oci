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

func DistributedDatabaseDistributedDatabasesAddGdsControlNodeResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("12h"),
			Update: tfresource.GetTimeoutDuration("12h"),
			Delete: tfresource.GetTimeoutDuration("12h"),
		},
		CreateContext: createDistributedDatabaseDistributedDatabasesAddGdsControlNodeWithContext,
		ReadContext:   readDistributedDatabaseDistributedDatabasesAddGdsControlNodeWithContext,
		DeleteContext: deleteDistributedDatabaseDistributedDatabasesAddGdsControlNodeWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"distributed_database_id": {
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

func createDistributedDatabaseDistributedDatabasesAddGdsControlNodeWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DistributedDatabaseDistributedDatabasesAddGdsControlNodeResourceCrud{
		Client:            m.(*client.OracleClients).DistributedDbServiceClient(),
		WorkRequestClient: m.(*client.OracleClients).DistributedDatabaseDistributedDbWorkRequestServiceClient(),
	}
	sync.D = d
	if err := sync.CreateWithContext(ctx); err != nil {
		return tfresource.HandleDiagError(m, err)
	}
	d.SetId(distributedDatabaseActionID("DistributedDatabaseGdsControlNodeAction-", d, "distributed_database_id", "public_ssh_key", "subnet_id"))
	return nil
}

func readDistributedDatabaseDistributedDatabasesAddGdsControlNodeWithContext(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {
	return nil
}

func deleteDistributedDatabaseDistributedDatabasesAddGdsControlNodeWithContext(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

type DistributedDatabaseDistributedDatabasesAddGdsControlNodeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_distributed_database.DistributedDbServiceClient
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_distributed_database.DistributedDbWorkRequestServiceClient
}

func (s *DistributedDatabaseDistributedDatabasesAddGdsControlNodeResourceCrud) createAction(ctx context.Context) error {
	request := oci_distributed_database.AddDistributedDatabaseGdsControlNodeRequest{}
	distributedDatabaseID := s.D.Get("distributed_database_id").(string)
	if err := waitForDistributedDatabaseActionReady(ctx, s.D, s.Client, distributedDatabaseID); err != nil {
		return err
	}

	request.DistributedDatabaseId = &distributedDatabaseID
	publicSshKey := s.D.Get("public_ssh_key").(string)
	request.PublicSshKey = &publicSshKey
	if subnetID, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetID.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "distributed_database")
	response, err := s.Client.AddDistributedDatabaseGdsControlNode(ctx, request)
	if err != nil {
		return err
	}

	workID := response.OpcWorkRequestId
	if workID == nil || strings.TrimSpace(*workID) == "" {
		return fmt.Errorf("missing opc-work-request-id for AddDistributedDatabaseGdsControlNode")
	}
	if _, err := distributedDatabaseWaitForWorkRequest(ctx, workID, "distributeddatabase", oci_distributed_database.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries, s.WorkRequestClient); err != nil {
		return err
	}
	return waitForDistributedDatabaseActionReady(ctx, s.D, s.Client, distributedDatabaseID)
}

func (s *DistributedDatabaseDistributedDatabasesAddGdsControlNodeResourceCrud) CreateWithContext(ctx context.Context) error {
	return s.createAction(ctx)
}
