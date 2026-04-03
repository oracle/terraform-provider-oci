package core

import (
	"reflect"
	"testing"
)

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
			name: "preserves state order for surviving entries and appends missing middle config entry",
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
					"ipv6cidr_block":  "2607:f590:0000:2200::/64",
					"byoipv6range_id": "range-c",
				},
				map[string]interface{}{
					"ipv6cidr_block":  "2607:f590:0000:2000::/64",
					"byoipv6range_id": "range-b",
				},
			},
			changed: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, changed := normalizeByoipv6CidrDetailsForDiff(tt.oldValue, tt.newValue)
			if changed != tt.changed {
				t.Fatalf("normalizeByoipv6CidrDetailsForDiff() changed = %v, want %v", changed, tt.changed)
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
