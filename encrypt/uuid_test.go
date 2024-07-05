package encrypt_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/supernarsi/gotool/encrypt"
)

func TestStringToUUID(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "t1", input: "TEST_12345678", want: "43fa6f2f-cb55-5b26-5edf-ca673e75e4b8"},
		{name: "t2", input: "0987654321", want: "17756315-ebd4-7b71-1035-9fc7b168179b"},
		{name: "t3", input: " ", want: "36a9e7f1-c95b-82ff-b997-43e0c5c4ce95"},
		{name: "t4", input: "", want: "e3b0c442-98fc-1c14-9afb-f4c8996fb924"},
		{name: "t5", input: "000", want: "2ac9a674-6aca-543a-f8df-f39894cfe817"},
		{name: "t6", input: "aaaaaaaaaaaaaaaaaaaaaa", want: "ec7c494d-f6d2-a7ea-3666-8d656e6b8979"},
		{name: "t7", input: "hello world and golang", want: "15e1afa6-f7e0-c3f3-8ea8-953f68fb6747"},
		{name: "t8", input: "KJLSDUIHFNUE18723JUIWji", want: "115576c2-59ab-5dc9-2c8d-f8bc1552f92b"},
		{name: "t9", input: "☀️🌛中文试一下", want: "305bd630-0c52-53db-8ab7-8f7d4b554a9c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encrypt.StringToUuid(tt.input); got != tt.want {
				t.Errorf("StringToUuid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToUUIDV5(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "t1", input: "TEST_LMC1202407051515584UX3Y4E3", want: "d76e5fbe-9ab0-5c12-a5c0-97abc13b1b02"},
		{name: "t2", input: "LMC1202407051515584UX3Y000", want: "4a4bcceb-d06b-510e-9a64-6e203efe5e9b"},
		{name: "t3", input: "LMC1202407051515584UX3Y000123123123", want: "50da975d-fba3-51fe-9a1a-851d85949c8d"},
		{name: "t4", input: "LMC1202407051515584UX3Y00012312312333333", want: "39b7ab5d-a37d-59d5-8cc0-b2aa02644f24"},
		{name: "t5", input: "", want: "6e27877b-6f4b-5895-9a0f-9a16080460b6"},
		{name: "t6", input: " ", want: "5cf0d00c-7c2f-5439-afc0-345445a1ee22"},
		{name: "t7", input: "  ", want: "f4ff705e-cff3-59ff-bb24-c5ecb29e833e"},
	}

	namespace := uuid.Must(uuid.Parse("123e4567-e89b-12d3-a456-426614174000"))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encrypt.StringToUuidByNamespace(namespace, tt.input); got != tt.want {
				t.Errorf("StringToUuidByNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
