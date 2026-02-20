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
