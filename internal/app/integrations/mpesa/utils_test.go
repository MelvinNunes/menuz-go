package mpesa

import (
	"testing"
)

func Test_generateThirdPartyReference(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testing invalid upper length", args: args{length: 16}, want: invalidCodeLength},
		{name: "testing invalid down length", args: args{length: 4}, want: invalidCodeLength},
		{name: "testing valid code length", args: args{length: 11}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateThirdPartyReference(tt.args.length); tt.want != "" && got != tt.want {
				t.Errorf("generateThirdPartyReference() = %v, want %v", got, tt.want)
			}
			if got := generateThirdPartyReference(tt.args.length); got != invalidCodeLength && len(got) != tt.args.length {
				t.Errorf("The length of the code must be equal to the length of the reference!")
			}
		})
	}
}

func Test_isValidMpesaContact(t *testing.T) {
	type args struct {
		contact string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Invalid length", args: args{contact: "258842807039"}, want: false},
		{name: "Invalid length 2", args: args{contact: "258862807039"}, want: false},
		{name: "Invalid length 3", args: args{contact: "842807"}, want: false},
		{name: "Invalid Vodacom Contact", args: args{contact: "862807039"}, want: false},
		{name: "Valid Vodacom Contact", args: args{contact: "842807039"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidMpesaContact(tt.args.contact); got != tt.want {
				t.Errorf("isValidMpesaContact() = %v, want %v", got, tt.want)
			}
		})
	}
}
