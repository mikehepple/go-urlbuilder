package urlbuilder

import (
	"reflect"
	"testing"
)

func TestURLBuilder_WithQuery(t *testing.T) {
	type args struct {
		key         string
		value       string
		extraValues []string
	}
	tests := []struct {
		name    string
		input   URLBuilder
		args    args
		want    URLBuilder
		wantErr bool
	}{
		{
			name: "set a query val",
			args: args{key: "foo", value: "bar"},
			want: URLBuilder{
				RawQuery: "foo=bar",
			},
			wantErr: false,
		},
		{
			name: "set a query val with multiple values",
			args: args{key: "foo", value: "bar", extraValues: []string{"baz"}},
			want: URLBuilder{
				RawQuery: "foo=bar&foo=baz",
			},
			wantErr: false,
		},
		{
			name: "query keys and vals are encoded",
			args: args{key: "?/", value: "#&/"},
			want: URLBuilder{
				RawQuery: "%3F%2F=%23%26%2F",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.WithQuery(tt.args.key, tt.args.value, tt.args.extraValues...)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}
