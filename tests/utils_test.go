package tests

import (
	"reflect"
	"testing"

	helpers "tebakaja_lb_proxy/proxy/helpers"
)

func IsStringReflect(x interface{}) bool {
	return reflect.TypeOf(x).Kind() == reflect.String
}

func IsStringTypeAssertion(x interface{}) bool {
	_, ok := x.(string)
	return ok
}

func TestGetEndpointByRestService(t *testing.T) {
	tests := []struct {
		service string
		want    bool
	}{
		{"crypto", true},
		{"national", true},
		{"stock", true},
	}

	for _, tt := range tests {
		t.Run(tt.service, func(t *testing.T) {
			endpoint := helpers.GetEndpointService(tt.service)

			if got := IsStringReflect(endpoint); got != tt.want {
				t.Errorf("IsStringReflect(%v) = %v, want %v", endpoint, got, tt.want)
			}

			if got := IsStringTypeAssertion(endpoint); got != tt.want {
				t.Errorf("IsStringTypeAssertion(%v) = %v, want %v", endpoint, got, tt.want)
			}
		})
	}
}