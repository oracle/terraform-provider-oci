package core

import (
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func buildCoreVcnDnsResolverAssociationDataSourceCrudForSetDataTest(t *testing.T, state oci_core.VcnDnsResolverAssociationLifecycleStateEnum) *CoreVcnDnsResolverAssociationDataSourceCrud {
	t.Helper()

	d := schema.TestResourceDataRaw(t, CoreVcnDnsResolverAssociationDataSource().Schema, map[string]interface{}{
		"vcn_id": "ocid1.vcn.oc1..exampleuniqueID",
	})

	dnsResolverId := "ocid1.dnsresolver.oc1..exampleuniqueID"
	return &CoreVcnDnsResolverAssociationDataSourceCrud{
		D: d,
		Res: &oci_core.GetVcnDnsResolverAssociationResponse{
			VcnDnsResolverAssociation: oci_core.VcnDnsResolverAssociation{
				DnsResolverId:  &dnsResolverId,
				LifecycleState: state,
			},
		},
	}
}

// Covers the transient provisioning path where SetData waits before storing the resolver ID.
func TestCoreVcnDnsResolverAssociationDataSourceCrudSetDataWaitsForAvailable(t *testing.T) {
	originalWaitForCondition := waitForVcnDnsResolverAssociationDataSourceCondition
	defer func() {
		waitForVcnDnsResolverAssociationDataSourceCondition = originalWaitForCondition
	}()

	crud := buildCoreVcnDnsResolverAssociationDataSourceCrudForSetDataTest(t, oci_core.VcnDnsResolverAssociationLifecycleStateProvisioning)
	waitCalled := false
	waitForVcnDnsResolverAssociationDataSourceCondition = func(s tfresource.ResourceFetcher, resourceChangedFunc func() bool, timeout time.Duration) error {
		waitCalled = true
		if s != crud {
			t.Fatalf("expected wait to receive crud instance")
		}
		if resourceChangedFunc() {
			t.Fatalf("expected resource condition to be false before the association becomes available")
		}

		crud.Res.LifecycleState = oci_core.VcnDnsResolverAssociationLifecycleStateAvailable
		if !resourceChangedFunc() {
			t.Fatalf("expected resource condition to be true once the association is available")
		}
		return nil
	}

	if err := crud.SetData(); err != nil {
		t.Fatalf("SetData() returned error: %v", err)
	}

	if !waitCalled {
		t.Fatalf("expected SetData() to wait when association is not available")
	}
	if got := crud.D.Get("state").(string); got != string(oci_core.VcnDnsResolverAssociationLifecycleStateAvailable) {
		t.Fatalf("expected state to be %q, got %q", oci_core.VcnDnsResolverAssociationLifecycleStateAvailable, got)
	}
	if got := crud.D.Get("dns_resolver_id").(string); got != "ocid1.dnsresolver.oc1..exampleuniqueID" {
		t.Fatalf("expected dns_resolver_id to be set, got %q", got)
	}
	if crud.D.Id() == "" {
		t.Fatalf("expected data source ID to be set")
	}
}

// Covers terminal states so SetData stops waiting but does not save incomplete association data.
func TestCoreVcnDnsResolverAssociationDataSourceCrudSetDataReturnsUnexpectedStateForTerminalState(t *testing.T) {
	originalWaitForCondition := waitForVcnDnsResolverAssociationDataSourceCondition
	defer func() {
		waitForVcnDnsResolverAssociationDataSourceCondition = originalWaitForCondition
	}()

	crud := buildCoreVcnDnsResolverAssociationDataSourceCrudForSetDataTest(t, oci_core.VcnDnsResolverAssociationLifecycleStateProvisioning)
	waitCalled := false
	waitForVcnDnsResolverAssociationDataSourceCondition = func(s tfresource.ResourceFetcher, resourceChangedFunc func() bool, timeout time.Duration) error {
		waitCalled = true
		crud.Res.LifecycleState = oci_core.VcnDnsResolverAssociationLifecycleStateTerminated
		if !resourceChangedFunc() {
			t.Fatalf("expected resource condition to stop waiting for terminal state")
		}
		return nil
	}

	err := crud.SetData()
	if err == nil {
		t.Fatalf("expected SetData() to return an unexpected state error")
	}
	if !waitCalled {
		t.Fatalf("expected SetData() to wait before returning terminal state error")
	}
	if !strings.Contains(err.Error(), "unexpected state: TERMINATED") {
		t.Fatalf("expected unexpected state error for TERMINATED, got %q", err)
	}
	if crud.D.Id() != "" {
		t.Fatalf("expected data source ID to remain unset")
	}
}
