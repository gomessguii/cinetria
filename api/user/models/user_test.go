package models

import "testing"

func TestCpfIsValid(t *testing.T) {
	testCases := []struct {
		cpf  string
		want bool
	}{
		{cpf: "", want: false},
		{cpf: "477.679.580", want: false},
		{cpf: "477.679.580-95", want: false},
		{cpf: "477.679.580-94", want: true},
		{cpf: "47767958094", want: true},
	}

	for _, testCase := range testCases {
		user := NewBuilder().WithCpf(testCase.cpf).Build()
		got := user.IsCpfValid()
		if got != testCase.want {
			t.Errorf("want %t got %t", testCase.want, got)
		}
	}
}
