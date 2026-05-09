package core

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_work_requests "github.com/oracle/oci-go-sdk/v65/workrequests"
)

func testByoipv6Detail(rangeID string, cidr string) oci_core.Byoipv6CidrDetails {
	return oci_core.Byoipv6CidrDetails{
		Byoipv6RangeId: &rangeID,
		Ipv6CidrBlock:  &cidr,
	}
}

type mockCoreWorkRequestClient struct {
	status               oci_work_requests.WorkRequestStatusEnum
	errorMessages        []string
	getErr               error
	listErr              error
	getRetryPolicySeen   bool
	listRetryPolicySeen  bool
	listErrorsWasInvoked bool
}

func (m *mockCoreWorkRequestClient) GetWorkRequest(_ context.Context, request oci_work_requests.GetWorkRequestRequest) (oci_work_requests.GetWorkRequestResponse, error) {
	m.getRetryPolicySeen = request.RequestMetadata.RetryPolicy != nil
	return oci_work_requests.GetWorkRequestResponse{
		WorkRequest: oci_work_requests.WorkRequest{
			Status: m.status,
		},
	}, m.getErr
}

func (m *mockCoreWorkRequestClient) ListWorkRequestErrors(_ context.Context, request oci_work_requests.ListWorkRequestErrorsRequest) (oci_work_requests.ListWorkRequestErrorsResponse, error) {
	m.listErrorsWasInvoked = true
	m.listRetryPolicySeen = request.RequestMetadata.RetryPolicy != nil
	items := make([]oci_work_requests.WorkRequestError, 0, len(m.errorMessages))
	for _, message := range m.errorMessages {
		msg := message
		items = append(items, oci_work_requests.WorkRequestError{Message: &msg})
	}
	return oci_work_requests.ListWorkRequestErrorsResponse{
		Items: items,
	}, m.listErr
}

func TestValidateCoreWorkRequestStatus(t *testing.T) {
	workRequestID := "ocid1.coreservicesworkrequest.oc1..example"

	// This covers Core work requests where the generic waiter can find the
	// expected resource/action even though the top-level work request status is
	// already FAILED or CANCELED.
	tests := []struct {
		name                 string
		client               *mockCoreWorkRequestClient
		wantErrContains      string
		wantListErrorsCalled bool
	}{
		{
			name: "succeeded work request returns nil",
			client: &mockCoreWorkRequestClient{
				status: oci_work_requests.WorkRequestStatusSucceeded,
			},
		},
		{
			name: "failed work request returns listed error",
			client: &mockCoreWorkRequestClient{
				status:        oci_work_requests.WorkRequestStatusFailed,
				errorMessages: []string{"cannot remove IPv6 CIDR while a subnet uses it"},
			},
			wantErrContains:      "cannot remove IPv6 CIDR while a subnet uses it",
			wantListErrorsCalled: true,
		},
		{
			name: "canceled work request returns error without listed details",
			client: &mockCoreWorkRequestClient{
				status: oci_work_requests.WorkRequestStatusCanceled,
			},
			wantErrContains:      "status: CANCELED",
			wantListErrorsCalled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCoreWorkRequestStatus(context.Background(), tt.client, &workRequestID, "vcn", false)

			if tt.wantErrContains == "" && err != nil {
				t.Fatalf("validateCoreWorkRequestStatus() returned error: %v", err)
			}
			if tt.wantErrContains != "" {
				if err == nil {
					t.Fatalf("validateCoreWorkRequestStatus() returned nil error, want %q", tt.wantErrContains)
				}
				if !strings.Contains(err.Error(), tt.wantErrContains) {
					t.Fatalf("validateCoreWorkRequestStatus() error = %q, want substring %q", err.Error(), tt.wantErrContains)
				}
			}
			if !tt.client.getRetryPolicySeen {
				t.Fatal("expected GetWorkRequest to use a retry policy")
			}
			if tt.client.listErrorsWasInvoked != tt.wantListErrorsCalled {
				t.Fatalf("ListWorkRequestErrors called = %v, want %v", tt.client.listErrorsWasInvoked, tt.wantListErrorsCalled)
			}
			if tt.wantListErrorsCalled && !tt.client.listRetryPolicySeen {
				t.Fatal("expected ListWorkRequestErrors to use a retry policy")
			}
		})
	}
}

func Test_computeIPv6BlocksFromBYOIPv6Details(t *testing.T) {
	var byoIpV6CidrDetailsCorrupted interface{} = []map[string]interface{}{
		{
			"ipv6cidr_block":  "2607:f590:2::/48",
			"byoipv6range_id": "randomId",
		},
	}
	var byoIpV6CidrDetailsValid interface{} = []interface{}{
		map[string]interface{}{"ipv6cidr_block": "2607:f590:2::/48", "byoipv6range_id": "randomId"},
		map[string]interface{}{"ipv6cidr_block": "2607:f590:3::/48", "byoipv6range_id": "randomId"},
	}

	var byoIpV6CidrDetailsValidAndInvalid interface{} = []interface{}{
		map[string]interface{}{"foo": "bar"},
		map[string]interface{}{"ipv6cidr_block": 30},
		map[string]interface{}{"ipv6cidr_block": "2607:f590:2::/48", "byoipv6range_id": "randomId"},
		map[string]interface{}{"ipv6cidr_block": "2607:f590:3::/48", "byoipv6range_id": "randomId"},
	}

	type args struct {
		byoIpV6CidrDetails interface{}
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test invalid input - nil",
			args: args{
				byoIpV6CidrDetails: nil,
			},
			want: []string{},
		},
		{
			name: "Test invalid format ([]map[string]interface)",
			args: args{
				byoIpV6CidrDetails: byoIpV6CidrDetailsCorrupted,
			},
			want: []string{},
		},
		{
			name: "Test invalid format (number)",
			args: args{
				byoIpV6CidrDetails: 12344,
			},
			want: []string{},
		},
		{
			name: "Test invalid format (slice without maps)",
			args: args{
				byoIpV6CidrDetails: []interface{}{1, "string", true},
			},
			want: []string{},
		},
		{
			name: "Test invalid input (no valid block)",
			args: args{
				byoIpV6CidrDetails: []interface{}{
					map[string]interface{}{"block": 42},
					map[string]interface{}{"block": nil},
				},
			},
			want: []string{},
		},
		{
			name: "Test empty input",
			args: args{
				byoIpV6CidrDetails: []interface{}{},
			},
			want: []string{},
		},
		{
			name: "Test mix of valid",
			args: args{
				byoIpV6CidrDetails: byoIpV6CidrDetailsValid,
			},
			want: []string{"2607:f590:2::/48", "2607:f590:3::/48"},
		},
		{
			name: "Test mix of valid and invalid",
			args: args{
				byoIpV6CidrDetails: byoIpV6CidrDetailsValidAndInvalid,
			},
			want: []string{"2607:f590:2::/48", "2607:f590:3::/48"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeIPv6BlocksFromBYOIPv6Details(tt.args.byoIpV6CidrDetails); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeIPv6BlocksFromBYOIPv6Details() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalizeByoipv6CidrDetailsForDiff(t *testing.T) {
	tests := []struct {
		name     string
		oldValue interface{}
		newValue interface{}
		want     []interface{}
		changed  bool
	}{
		{
			name: "suppresses matching block with different range id",
			oldValue: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2000::/64",
					"byoipv6range_id": "(known_after_apply)",
				},
			},
			newValue: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2000::/64",
					"byoipv6range_id": "ocid1.byoipv6range.oc1..example",
				},
			},
			want: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2000::/64",
					"byoipv6range_id": "(known_after_apply)",
				},
			},
			changed: true,
		},
		{
			name: "matches canonically equivalent blocks",
			oldValue: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2001:db8:1:2::/80",
					"byoipv6range_id": "(known_after_apply)",
				},
			},
			newValue: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2001:0db8:0001:0002:0000:0000:0000:0000/80",
					"byoipv6range_id": "ocid1.byoipv6range.oc1..example",
				},
			},
			want: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2001:db8:1:2::/80",
					"byoipv6range_id": "(known_after_apply)",
				},
			},
			changed: true,
		},
		{
			name: "keeps new details when block is not already present",
			oldValue: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:0200::/64",
					"byoipv6range_id": "(known_after_apply)",
				},
			},
			newValue: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2000::/64",
					"byoipv6range_id": "ocid1.byoipv6range.oc1..example",
				},
			},
			changed: false,
		},
		{
			name: "preserves config order while reusing matching state entries",
			oldValue: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:0200::/64",
					"byoipv6range_id": "range-a",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2200::/64",
					"byoipv6range_id": "range-c",
				},
			},
			newValue: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:0200::/64",
					"byoipv6range_id": "range-a",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2000::/64",
					"byoipv6range_id": "range-b",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2200::/64",
					"byoipv6range_id": "range-c",
				},
			},
			want: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:0200::/64",
					"byoipv6range_id": "range-a",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2000::/64",
					"byoipv6range_id": "range-b",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2200::/64",
					"byoipv6range_id": "range-c",
				},
			},
			changed: false,
		},
		{
			name: "preserves replacement config order instead of grouping retained entries first",
			oldValue: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2200::/64",
					"byoipv6range_id": "range-a",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2201::/64",
					"byoipv6range_id": "range-b",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2202::/64",
					"byoipv6range_id": "range-c",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2203::/64",
					"byoipv6range_id": "range-d",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2204::/64",
					"byoipv6range_id": "range-e",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2205::/64",
					"byoipv6range_id": "range-f",
				},
			},
			newValue: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2206::/64",
					"byoipv6range_id": "range-g",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2201::/64",
					"byoipv6range_id": "range-b",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2207::/64",
					"byoipv6range_id": "range-h",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2208::/64",
					"byoipv6range_id": "range-i",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2204::/64",
					"byoipv6range_id": "range-e",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2205::/64",
					"byoipv6range_id": "range-f",
				},
			},
			want: []interface{}{
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2206::/64",
					"byoipv6range_id": "range-g",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2201::/64",
					"byoipv6range_id": "range-b",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2207::/64",
					"byoipv6range_id": "range-h",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2208::/64",
					"byoipv6range_id": "range-i",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2204::/64",
					"byoipv6range_id": "range-e",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2205::/64",
					"byoipv6range_id": "range-f",
				},
			},
			changed: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, changed := normalizeByoipv6CidrDetailsForDiff(tt.oldValue, tt.newValue)
			if changed != tt.changed {
				t.Fatalf("normalizeByoipv6CidrDetailsForDiff() changed = %v, want %v", changed, tt.changed)
			}
			if !changed {
				if got != nil {
					t.Fatalf("normalizeByoipv6CidrDetailsForDiff() = %#v, want nil when unchanged", got)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("normalizeByoipv6CidrDetailsForDiff() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_convertToCanonical(t *testing.T) {
	type args struct {
		block string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test - Some octets missing values",
			args: args{
				block: "2001:db8:abcd:1:2::/80",
			},
			want: "2001:0db8:abcd:0001:0002:0000:0000:0000/80",
		},
		{
			name: "Test - Some octets missing values",
			args: args{
				block: ":db8:1:2:3/48",
			},
			want: "0000:0db8:0001:0002:0003:0000:0000:0000/48",
		},
		{
			name: "Test - Some octets missing values",
			args: args{
				block: "2001:db8:1:2:3:4:5:/26",
			},
			want: "2001:0db8:0001:0002:0003:0004:0005:0000/26",
		},
		{
			name: "Test - All octets with values",
			args: args{
				block: "2001:db8:1:2:3:4:5:6/64",
			},
			want: "2001:0db8:0001:0002:0003:0004:0005:0006/64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToCanonical(tt.args.block); got != tt.want {
				t.Errorf("convertToCanonical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEmptyString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test empty string",
			args: args{
				input: "",
			},
			want: true,
		},
		{
			name: "Test whitespace only string",
			args: args{
				input: " \t \n ",
			},
			want: true,
		},
		{
			name: "Test non-empty string",
			args: args{
				input: "hello",
			},
			want: false,
		},
		{
			name: "Test non-empty string with surrounding whitespace",
			args: args{
				input: "  hello  ",
			},
			want: false,
		},
		{
			name: "Test newline wrapped text",
			args: args{
				input: "\nvalue\n",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEmptyString(tt.args.input); got != tt.want {
				t.Errorf("isEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubnetIpv6PatchChangeSetShouldUsePatch(t *testing.T) {
	tests := []struct {
		name      string
		changeSet subnetIpv6PatchChangeSet
		want      bool
		wantErr   bool
	}{
		{
			name: "single ipv6 tail add keeps legacy api",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:1::/64"},
			},
			want: false,
		},
		{
			name: "single ipv6 tail removal keeps legacy api",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:1::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8::/64"},
			},
			want: false,
		},
		{
			name: "single ipv6 beginning add uses patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8:1::/64", "2001:db8:2::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:1::/64", "2001:db8:2::/64"},
			},
			want: true,
		},
		{
			name: "single ipv6 middle add uses patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:2::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:1::/64", "2001:db8:2::/64"},
			},
			want: true,
		},
		{
			name: "single ipv6 beginning removal uses patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:1::/64", "2001:db8:2::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8:1::/64", "2001:db8:2::/64"},
			},
			want: true,
		},
		{
			name: "single ipv6 middle removal uses patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:1::/64", "2001:db8:2::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:2::/64"},
			},
			want: true,
		},
		{
			name: "ipv6cidr_block plus ipv6cidr_blocks uses patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlockChanged:  true,
				oldIpv6CidrBlock:      "2001:db8::/64",
				newIpv6CidrBlock:      "2001:db8:1::/64",
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8:1::/64", "2001:db8:2::/64"},
			},
			want: true,
		},
		{
			name: "ipv6cidr_block change plus ipv6cidr_blocks uses patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlockChanged:  true,
				oldIpv6CidrBlock:      "2001:db8::/64",
				newIpv6CidrBlock:      "2001:db8:1::/64",
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{},
				newIpv6CidrBlocks:     []string{"2001:db8:1::/64"},
			},
			want: true,
		},
		{
			name: "ipv6cidr_block removal plus ipv6cidr_blocks uses patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlockChanged:  true,
				oldIpv6CidrBlock:      "2001:db8::/64",
				newIpv6CidrBlock:      "",
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8:1::/64", "2001:db8:2::/64"},
			},
			want: true,
		},
		{
			name: "ipv6cidr_block addition addition ipv6cidr_blocks uses patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlockChanged:  true,
				oldIpv6CidrBlock:      "",
				newIpv6CidrBlock:      "2001:db8::/64",
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8:1::/64", "2001:db8:2::/64"},
			},
			want: true,
		},
		{
			name: "multiple ipv6 adds use patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:1::/64", "2001:db8:2::/64"},
			},
			want: true,
		},
		{
			name: "ipv6cidr_block add keeps legacy api",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlockChanged: true,
				oldIpv6CidrBlock:     "",
				newIpv6CidrBlock:     "2001:db8::/64",
			},
			want: false,
		},
		{
			name: "ipv6cidr_block replacement uses patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlockChanged: true,
				oldIpv6CidrBlock:     "2001:db8::/64",
				newIpv6CidrBlock:     "2001:db8:1::/64",
			},
			want: true,
		},
		{
			name: "single ipv6 replacement uses patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8:1::/64"},
			},
			want: true,
		},
		{
			name: "multiple ipv6 replacements use patch",
			changeSet: subnetIpv6PatchChangeSet{
				ipv6CidrBlocksChanged: true,
				oldIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:1::/64"},
				newIpv6CidrBlocks:     []string{"2001:db8:2::/64", "2001:db8:3::/64"},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.changeSet.shouldUsePatch()
			if (err != nil) != tt.wantErr {
				t.Fatalf("shouldUsePatch() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Fatalf("shouldUsePatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubnetIpv6PatchChangeSetChangedFieldCount(t *testing.T) {
	changeSet := subnetIpv6PatchChangeSet{
		ipv6CidrBlockChanged:  true,
		ipv6CidrBlocksChanged: true,
	}

	if got := changeSet.changedFieldCount(); got != 2 {
		t.Fatalf("changedFieldCount() = %d, want 2", got)
	}
}

func TestSubnetIpv6PatchChangeSetBuildPatchInstructions(t *testing.T) {
	changeSet := subnetIpv6PatchChangeSet{
		ipv6CidrBlockChanged:  true,
		newIpv6CidrBlock:      "2001:db8::/64",
		ipv6CidrBlocksChanged: true,
		newIpv6CidrBlocks:     []string{"2001:db8::/64", "2001:db8:1::/64"},
	}

	got, err := changeSet.buildPatchInstructions()
	if err != nil {
		t.Fatalf("buildPatchInstructions() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("buildPatchInstructions() len = %d, want 1", len(got))
	}

	instruction, ok := got[0].(oci_core.PatchSubnetReplaceInstruction)
	if !ok {
		t.Fatalf("instruction[0] type = %T, want PatchSubnetReplaceInstruction", got[0])
	}
	if instruction.Selection == nil || *instruction.Selection != "ipv6CidrBlocks" {
		t.Fatalf("instruction[0] selection = %v, want ipv6CidrBlocks", instruction.Selection)
	}
	if instruction.Value == nil {
		t.Fatal("instruction[0] value is nil")
	}
	gotValue, ok := (*instruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("instruction[0] value type = %T, want map[string]interface{}", *instruction.Value)
	}
	wantValue := map[string]interface{}{"cidrs": []string{"2001:db8::/64", "2001:db8:1::/64"}}
	if !reflect.DeepEqual(gotValue, wantValue) {
		t.Fatalf("instruction[0] value = %#v, want %#v", gotValue, wantValue)
	}
}

func TestSubnetIpv6PatchChangeSetBuildPatchInstructionsAddsChangedScalarToBlocks(t *testing.T) {
	changeSet := subnetIpv6PatchChangeSet{
		ipv6CidrBlockChanged:  true,
		newIpv6CidrBlock:      "2001:db8::/64",
		ipv6CidrBlocksChanged: true,
		newIpv6CidrBlocks:     []string{"2001:db8:1::/64", "2001:db8:2::/64"},
	}

	got, err := changeSet.buildPatchInstructions()
	if err != nil {
		t.Fatalf("buildPatchInstructions() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("buildPatchInstructions() len = %d, want 1", len(got))
	}

	instruction, ok := got[0].(oci_core.PatchSubnetReplaceInstruction)
	if !ok {
		t.Fatalf("instruction[0] type = %T, want PatchSubnetReplaceInstruction", got[0])
	}
	if instruction.Selection == nil || *instruction.Selection != "ipv6CidrBlocks" {
		t.Fatalf("instruction[0] selection = %v, want ipv6CidrBlocks", instruction.Selection)
	}
	if instruction.Value == nil {
		t.Fatal("instruction[0] value is nil")
	}
	gotValue, ok := (*instruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("instruction[0] value type = %T, want map[string]interface{}", *instruction.Value)
	}
	wantValue := map[string]interface{}{"cidrs": []string{"2001:db8:1::/64", "2001:db8:2::/64", "2001:db8::/64"}}
	if !reflect.DeepEqual(gotValue, wantValue) {
		t.Fatalf("instruction[0] value = %#v, want %#v", gotValue, wantValue)
	}
}

func TestSubnetIpv6PatchChangeSetBuildPatchInstructionsDoesNotInventScalarForBlocksPatch(t *testing.T) {
	changeSet := subnetIpv6PatchChangeSet{
		ipv6CidrBlocksChanged: true,
		newIpv6CidrBlocks:     []string{"2001:db8:1::/64", "2001:db8:2::/64"},
	}

	got, err := changeSet.buildPatchInstructions()
	if err != nil {
		t.Fatalf("buildPatchInstructions() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("buildPatchInstructions() len = %d, want 1", len(got))
	}

	instruction, ok := got[0].(oci_core.PatchSubnetReplaceInstruction)
	if !ok {
		t.Fatalf("instruction[0] type = %T, want PatchSubnetReplaceInstruction", got[0])
	}
	if instruction.Selection == nil || *instruction.Selection != "ipv6CidrBlocks" {
		t.Fatalf("instruction[0] selection = %v, want ipv6CidrBlocks", instruction.Selection)
	}
	if instruction.Value == nil {
		t.Fatal("instruction[0] value is nil")
	}
	gotValue, ok := (*instruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("instruction[0] value type = %T, want map[string]interface{}", *instruction.Value)
	}
	wantValue := map[string]interface{}{"cidrs": []string{"2001:db8:1::/64", "2001:db8:2::/64"}}
	if !reflect.DeepEqual(gotValue, wantValue) {
		t.Fatalf("instruction[0] value = %#v, want %#v", gotValue, wantValue)
	}
}

func TestSubnetIpv6PatchChangeSetBuildPatchInstructionsUsesCurrentScalarForBlocksPatch(t *testing.T) {
	changeSet := subnetIpv6PatchChangeSet{
		currentIpv6CidrBlock:  "2001:db8::/64",
		ipv6CidrBlockInConfig: true,
		ipv6CidrBlocksChanged: true,
		newIpv6CidrBlocks:     []string{"2001:db8:1::/64", "2001:db8:2::/64"},
	}

	got, err := changeSet.buildPatchInstructions()
	if err != nil {
		t.Fatalf("buildPatchInstructions() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("buildPatchInstructions() len = %d, want 1", len(got))
	}

	firstInstruction, ok := got[0].(oci_core.PatchSubnetReplaceInstruction)
	if !ok {
		t.Fatalf("instruction[0] type = %T, want PatchSubnetReplaceInstruction", got[0])
	}
	if firstInstruction.Selection == nil || *firstInstruction.Selection != "ipv6CidrBlocks" {
		t.Fatalf("instruction[0] selection = %v, want ipv6CidrBlocks", firstInstruction.Selection)
	}
	if firstInstruction.Value == nil {
		t.Fatal("instruction[0] value is nil")
	}
	gotValue, ok := (*firstInstruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("instruction[0] value type = %T, want map[string]interface{}", *firstInstruction.Value)
	}
	wantValue := map[string]interface{}{"cidrs": []string{"2001:db8:1::/64", "2001:db8:2::/64", "2001:db8::/64"}}
	if !reflect.DeepEqual(gotValue, wantValue) {
		t.Fatalf("instruction[0] value = %#v, want %#v", gotValue, wantValue)
	}
}

func TestSubnetIpv6PatchChangeSetBuildPatchInstructionsSkipsComputedScalarForBlocksPatch(t *testing.T) {
	changeSet := subnetIpv6PatchChangeSet{
		currentIpv6CidrBlock:  "2001:db8::/64",
		ipv6CidrBlocksChanged: true,
		newIpv6CidrBlocks:     []string{"2001:db8:1::/64", "2001:db8:2::/64"},
	}

	got, err := changeSet.buildPatchInstructions()
	if err != nil {
		t.Fatalf("buildPatchInstructions() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("buildPatchInstructions() len = %d, want 1", len(got))
	}

	firstInstruction, ok := got[0].(oci_core.PatchSubnetReplaceInstruction)
	if !ok {
		t.Fatalf("instruction[0] type = %T, want PatchSubnetReplaceInstruction", got[0])
	}
	if firstInstruction.Selection == nil || *firstInstruction.Selection != "ipv6CidrBlocks" {
		t.Fatalf("instruction[0] selection = %v, want ipv6CidrBlocks", firstInstruction.Selection)
	}
	if firstInstruction.Value == nil {
		t.Fatal("instruction[0] value is nil")
	}
	gotValue, ok := (*firstInstruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("instruction[0] value type = %T, want map[string]interface{}", *firstInstruction.Value)
	}
	wantValue := map[string]interface{}{"cidrs": []string{"2001:db8:1::/64", "2001:db8:2::/64"}}
	if !reflect.DeepEqual(gotValue, wantValue) {
		t.Fatalf("instruction[0] value = %#v, want %#v", gotValue, wantValue)
	}
}

func TestBuildSubnetReplaceInstruction(t *testing.T) {
	value := map[string]interface{}{"cidr": "2001:db8::/64"}

	instruction := buildSubnetReplaceInstruction("ipv6CidrBlock", value)
	replaceInstruction, ok := instruction.(oci_core.PatchSubnetReplaceInstruction)
	if !ok {
		t.Fatalf("buildSubnetReplaceInstruction() type = %T, want PatchSubnetReplaceInstruction", instruction)
	}
	if replaceInstruction.Selection == nil || *replaceInstruction.Selection != "ipv6CidrBlock" {
		t.Fatalf("buildSubnetReplaceInstruction() selection = %v, want ipv6CidrBlock", replaceInstruction.Selection)
	}
	if replaceInstruction.Value == nil {
		t.Fatal("buildSubnetReplaceInstruction() value is nil")
	}
	gotValue, ok := (*replaceInstruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("buildSubnetReplaceInstruction() value type = %T, want map[string]interface{}", *replaceInstruction.Value)
	}
	if !reflect.DeepEqual(gotValue, value) {
		t.Fatalf("buildSubnetReplaceInstruction() value = %#v, want %#v", gotValue, value)
	}
}

func TestInterfaceSliceToStringSlice(t *testing.T) {
	tests := []struct {
		name string
		raw  interface{}
		want []string
	}{
		{
			name: "converts string slice values in order",
			raw:  []interface{}{"a", "b", "c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "preserves empty slot for nil entries",
			raw:  []interface{}{"a", nil, "c"},
			want: []string{"a", "", "c"},
		},
		{
			name: "returns empty for non slice input",
			raw:  "not-a-slice",
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interfaceSliceToStringSlice(tt.raw); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("interfaceSliceToStringSlice() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestBuildSubnetIpv6PatchChangeSetPrefersConfiguredIpv6Field(t *testing.T) {
	resourceData := schema.TestResourceDataRaw(t, CoreSubnetResource().Schema, map[string]interface{}{
		"cidr_block":     "10.0.0.0/24",
		"compartment_id": "ocid1.compartment.oc1..exampleuniqueID",
		"vcn_id":         "ocid1.vcn.oc1..exampleuniqueID",
		"ipv6cidr_block": "2607:db8::/64",
		"display_name":   "example",
	})

	if err := resourceData.Set("ipv6cidr_blocks", []string{"2607:db8::/64"}); err != nil {
		t.Fatalf("Set(ipv6cidr_blocks) error = %v", err)
	}

	if err := resourceData.Set("cidr_block", "10.0.0.0/16"); err != nil {
		t.Fatalf("Set(cidr_block) error = %v", err)
	}
	if err := resourceData.Set("ipv6cidr_block", "2607:db8:1::/64"); err != nil {
		t.Fatalf("Set(ipv6cidr_block) error = %v", err)
	}

	changeSet := buildSubnetIpv6PatchChangeSet(resourceData)

	if !changeSet.ipv6CidrBlockChanged {
		t.Fatalf("expected ipv6CidrBlockChanged to be true")
	}
	if changeSet.ipv6CidrBlocksChanged {
		t.Fatalf("expected ipv6CidrBlocksChanged to be false when only ipv6cidr_block is configured")
	}

	instructions, err := changeSet.buildPatchInstructions()
	if err != nil {
		t.Fatalf("buildPatchInstructions() error = %v", err)
	}

	if len(instructions) != 1 {
		t.Fatalf("buildPatchInstructions() len = %d, want 1", len(instructions))
	}

	firstInstruction, ok := instructions[0].(oci_core.PatchSubnetReplaceInstruction)
	if !ok {
		t.Fatalf("instruction[0] type = %T, want PatchSubnetReplaceInstruction", instructions[0])
	}
	if firstInstruction.Selection == nil || *firstInstruction.Selection != "ipv6CidrBlock" {
		t.Fatalf("instruction[0] selection = %v, want ipv6CidrBlock", firstInstruction.Selection)
	}
}

func TestVcnIpv6PatchChangeSetShouldUsePatch(t *testing.T) {
	tests := []struct {
		name      string
		changeSet vcnIpv6PatchChangeSet
		want      bool
	}{
		{
			name: "single private cidr add keeps legacy api",
			changeSet: vcnIpv6PatchChangeSet{
				ipv6PrivateCidrChanged: true,
				oldIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64"},
				newIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64"},
			},
			want: false,
		},
		{
			name: "single private cidr tail removal keeps legacy api",
			changeSet: vcnIpv6PatchChangeSet{
				ipv6PrivateCidrChanged: true,
				oldIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64"},
				newIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64"},
			},
			want: false,
		},
		{
			name: "single private cidr middle add uses patch",
			changeSet: vcnIpv6PatchChangeSet{
				ipv6PrivateCidrChanged: true,
				oldIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:3::/64"},
				newIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64", "fd00:1000:0:3::/64"},
			},
			want: true,
		},
		{
			name: "single private cidr middle removal uses patch",
			changeSet: vcnIpv6PatchChangeSet{
				ipv6PrivateCidrChanged: true,
				oldIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64", "fd00:1000:0:3::/64"},
				newIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:3::/64"},
			},
			want: true,
		},
		{
			name: "single private cidr tail replacement uses patch",
			changeSet: vcnIpv6PatchChangeSet{
				ipv6PrivateCidrChanged: true,
				oldIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64"},
				newIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:3::/64"},
			},
			want: true,
		},
		{
			name: "both fields changed use patch",
			changeSet: vcnIpv6PatchChangeSet{
				byoipv6CidrDetailsChanged: true,
				oldByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
				},
				newByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
					testByoipv6Detail("range-b", "2607:f590:0000:2000::/64"),
				},
				ipv6PrivateCidrChanged: true,
				oldIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64"},
				newIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64"},
			},
			want: true,
		},
		{
			name: "multiple private cidr adds use patch",
			changeSet: vcnIpv6PatchChangeSet{
				ipv6PrivateCidrChanged: true,
				oldIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64"},
				newIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64", "fd00:1000:0:3::/64"},
			},
			want: true,
		},
		{
			name: "single private cidr replacement uses patch",
			changeSet: vcnIpv6PatchChangeSet{
				ipv6PrivateCidrChanged: true,
				oldIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64"},
				newIpv6PrivateCidrs:    []string{"fd00:1000:0:2::/64"},
			},
			want: true,
		},
		{
			name: "single byo add keeps legacy api",
			changeSet: vcnIpv6PatchChangeSet{
				byoipv6CidrDetailsChanged: true,
				oldByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
				},
				newByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
					testByoipv6Detail("range-b", "2607:f590:0000:2000::/64"),
				},
			},
			want: false,
		},
		{
			name: "single byo tail removal keeps legacy api",
			changeSet: vcnIpv6PatchChangeSet{
				byoipv6CidrDetailsChanged: true,
				oldByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
					testByoipv6Detail("range-b", "2607:f590:0000:2000::/64"),
				},
				newByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
				},
			},
			want: false,
		},
		{
			name: "single byo middle add uses patch",
			changeSet: vcnIpv6PatchChangeSet{
				byoipv6CidrDetailsChanged: true,
				oldByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
					testByoipv6Detail("range-c", "2607:f590:0000:2200::/64"),
				},
				newByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
					testByoipv6Detail("range-b", "2607:f590:0000:2000::/64"),
					testByoipv6Detail("range-c", "2607:f590:0000:2200::/64"),
				},
			},
			want: true,
		},
		{
			name: "single byo middle removal uses patch",
			changeSet: vcnIpv6PatchChangeSet{
				byoipv6CidrDetailsChanged: true,
				oldByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
					testByoipv6Detail("range-b", "2607:f590:0000:2000::/64"),
					testByoipv6Detail("range-c", "2607:f590:0000:2200::/64"),
				},
				newByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
					testByoipv6Detail("range-c", "2607:f590:0000:2200::/64"),
				},
			},
			want: true,
		},
		{
			name: "multiple byo adds use patch",
			changeSet: vcnIpv6PatchChangeSet{
				byoipv6CidrDetailsChanged: true,
				oldByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
				},
				newByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
					testByoipv6Detail("range-b", "2607:f590:0000:2000::/64"),
					testByoipv6Detail("range-c", "2607:f590:0000:2200::/64"),
				},
			},
			want: true,
		},
		{
			name: "single byo replacement uses patch",
			changeSet: vcnIpv6PatchChangeSet{
				byoipv6CidrDetailsChanged: true,
				oldByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
				},
				newByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
					testByoipv6Detail("range-b", "2607:f590:0000:2000::/64"),
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.changeSet.shouldUsePatch()
			if err != nil {
				t.Fatalf("shouldUsePatch() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("shouldUsePatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVcnIpv6PatchChangeSetBuildPatchInstructions(t *testing.T) {
	changeSet := vcnIpv6PatchChangeSet{
		byoipv6CidrDetailsChanged: true,
		newByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
			testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
			testByoipv6Detail("range-b", "2607:f590:0000:2000::/64"),
		},
		ipv6PrivateCidrChanged: true,
		newIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64"},
	}

	got, err := changeSet.buildPatchInstructions()
	if err != nil {
		t.Fatalf("buildPatchInstructions() error = %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("buildPatchInstructions() len = %d, want 2", len(got))
	}

	firstInstruction, ok := got[0].(oci_core.PatchVcnReplaceInstruction)
	if !ok {
		t.Fatalf("instruction[0] type = %T, want PatchVcnReplaceInstruction", got[0])
	}
	if firstInstruction.Selection == nil || *firstInstruction.Selection != "byoipv6CidrDetails" {
		t.Fatalf("instruction[0] selection = %v, want byoipv6CidrDetails", firstInstruction.Selection)
	}
	firstValue, ok := (*firstInstruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("instruction[0] value type = %T, want map[string]interface{}", *firstInstruction.Value)
	}
	wantFirstValue := map[string]interface{}{
		"cidrs": []map[string]interface{}{
			{"byoipv6RangeId": "range-a", "ipv6CidrBlock": "2607:f590:0000:0200::/64"},
			{"byoipv6RangeId": "range-b", "ipv6CidrBlock": "2607:f590:0000:2000::/64"},
		},
	}
	if !reflect.DeepEqual(firstValue, wantFirstValue) {
		t.Fatalf("instruction[0] value = %#v, want %#v", firstValue, wantFirstValue)
	}

	secondInstruction, ok := got[1].(oci_core.PatchVcnReplaceInstruction)
	if !ok {
		t.Fatalf("instruction[1] type = %T, want PatchVcnReplaceInstruction", got[1])
	}
	if secondInstruction.Selection == nil || *secondInstruction.Selection != "ipv6PrivateCidrBlocks" {
		t.Fatalf("instruction[1] selection = %v, want ipv6PrivateCidrBlocks", secondInstruction.Selection)
	}
	secondValue, ok := (*secondInstruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("instruction[1] value type = %T, want map[string]interface{}", *secondInstruction.Value)
	}
	wantSecondValue := map[string]interface{}{
		"cidrs": []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64"},
	}
	if !reflect.DeepEqual(secondValue, wantSecondValue) {
		t.Fatalf("instruction[1] value = %#v, want %#v", secondValue, wantSecondValue)
	}
}

func TestVcnIpv6PatchChangeSetBuildPatchInstructionsIncludesUnchangedConfiguredFields(t *testing.T) {
	changeSet := vcnIpv6PatchChangeSet{
		byoipv6CidrDetailsPresent: true,
		currentByoipv6CidrDetails: []oci_core.Byoipv6CidrDetails{
			testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
			testByoipv6Detail("range-b", "2607:f590:0000:2000::/64"),
		},
		ipv6PrivateCidrChanged: true,
		oldIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64"},
		newIpv6PrivateCidrs:    []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64"},
		ipv6PrivateCidrPresent: true,
		currentIpv6PrivateCidrs: []string{
			"fd00:1000:0:1::/64",
			"fd00:1000:0:2::/64",
		},
	}

	got, err := changeSet.buildPatchInstructions()
	if err != nil {
		t.Fatalf("buildPatchInstructions() error = %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("buildPatchInstructions() len = %d, want 2", len(got))
	}

	firstInstruction, ok := got[0].(oci_core.PatchVcnReplaceInstruction)
	if !ok {
		t.Fatalf("instruction[0] type = %T, want PatchVcnReplaceInstruction", got[0])
	}
	if firstInstruction.Selection == nil || *firstInstruction.Selection != "byoipv6CidrDetails" {
		t.Fatalf("instruction[0] selection = %v, want byoipv6CidrDetails", firstInstruction.Selection)
	}
	firstValue, ok := (*firstInstruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("instruction[0] value type = %T, want map[string]interface{}", *firstInstruction.Value)
	}
	wantFirstValue := map[string]interface{}{
		"cidrs": []map[string]interface{}{
			{"byoipv6RangeId": "range-a", "ipv6CidrBlock": "2607:f590:0000:0200::/64"},
			{"byoipv6RangeId": "range-b", "ipv6CidrBlock": "2607:f590:0000:2000::/64"},
		},
	}
	if !reflect.DeepEqual(firstValue, wantFirstValue) {
		t.Fatalf("instruction[0] value = %#v, want %#v", firstValue, wantFirstValue)
	}

	secondInstruction, ok := got[1].(oci_core.PatchVcnReplaceInstruction)
	if !ok {
		t.Fatalf("instruction[1] type = %T, want PatchVcnReplaceInstruction", got[1])
	}
	if secondInstruction.Selection == nil || *secondInstruction.Selection != "ipv6PrivateCidrBlocks" {
		t.Fatalf("instruction[1] selection = %v, want ipv6PrivateCidrBlocks", secondInstruction.Selection)
	}
	secondValue, ok := (*secondInstruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("instruction[1] value type = %T, want map[string]interface{}", *secondInstruction.Value)
	}
	wantSecondValue := map[string]interface{}{
		"cidrs": []string{"fd00:1000:0:1::/64", "fd00:1000:0:2::/64"},
	}
	if !reflect.DeepEqual(secondValue, wantSecondValue) {
		t.Fatalf("instruction[1] value = %#v, want %#v", secondValue, wantSecondValue)
	}
}

func TestBuildVcnReplaceInstruction(t *testing.T) {
	value := map[string]interface{}{"cidrs": []string{"fd00:1000:0:1::/64"}}

	instruction := buildVcnReplaceInstruction("ipv6PrivateCidrBlocks", value)
	replaceInstruction, ok := instruction.(oci_core.PatchVcnReplaceInstruction)
	if !ok {
		t.Fatalf("buildVcnReplaceInstruction() type = %T, want PatchVcnReplaceInstruction", instruction)
	}
	if replaceInstruction.Selection == nil || *replaceInstruction.Selection != "ipv6PrivateCidrBlocks" {
		t.Fatalf("buildVcnReplaceInstruction() selection = %v, want ipv6PrivateCidrBlocks", replaceInstruction.Selection)
	}
	gotValue, ok := (*replaceInstruction.Value).(map[string]interface{})
	if !ok {
		t.Fatalf("buildVcnReplaceInstruction() value type = %T, want map[string]interface{}", *replaceInstruction.Value)
	}
	if !reflect.DeepEqual(gotValue, value) {
		t.Fatalf("buildVcnReplaceInstruction() value = %#v, want %#v", gotValue, value)
	}
}

func TestInterfaceToByoipv6CidrDetailsSlice(t *testing.T) {
	tests := []struct {
		name    string
		raw     interface{}
		want    []oci_core.Byoipv6CidrDetails
		wantErr bool
	}{
		{
			name: "converts valid details",
			raw: []interface{}{
				map[string]interface{}{
					"byoipv6range_id": "range-a",
					"ipv6cidr_block":  "2607:f590:0000:0200::/64",
				},
				map[string]interface{}{
					"byoipv6range_id": "range-b",
					"ipv6cidr_block":  "2607:f590:0000:2000::/64",
				},
			},
			want: []oci_core.Byoipv6CidrDetails{
				testByoipv6Detail("range-a", "2607:f590:0000:0200::/64"),
				testByoipv6Detail("range-b", "2607:f590:0000:2000::/64"),
			},
		},
		{
			name:    "rejects malformed entries",
			raw:     []interface{}{"not-a-map"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := interfaceToByoipv6CidrDetailsSlice(tt.raw)
			if (err != nil) != tt.wantErr {
				t.Fatalf("interfaceToByoipv6CidrDetailsSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("interfaceToByoipv6CidrDetailsSlice() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
