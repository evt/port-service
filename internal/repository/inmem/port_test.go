package inmem

import (
	"testing"
	"time"

	"github.com/evt/port-service/internal/domain"
	"github.com/stretchr/testify/require"
)

func Test_portStoreToDomain(t *testing.T) {
	type args struct {
		p *Port
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Port
		wantErr bool
	}{
		{
			name: "should return error when store port is nil",
			args: args{
				p: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should return domain port when store port is not nil",
			args: args{
				p: newTestStorePort(t),
			},
			want:    newTestDomainPort(t),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := portStoreToDomain(tt.args.p)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

const testString = "test"

func newTestStorePort(t *testing.T) *Port {
	t.Helper()
	return &Port{
		ID:          testString,
		Name:        testString,
		Code:        testString,
		City:        testString,
		Country:     testString,
		Alias:       []string{testString},
		Regions:     []string{testString},
		Coordinates: []float64{1.0, 2.0},
		Province:    testString,
		Timezone:    testString,
		Unlocs:      []string{testString},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func newTestDomainPort(t *testing.T) *domain.Port {
	t.Helper()
	port, err := domain.NewPort(testString, testString, testString, testString, testString,
		[]string{testString}, []string{testString}, []float64{1.0, 2.0}, testString, testString, []string{testString})
	require.NoError(t, err)
	return port
}
