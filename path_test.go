package urlbuilder

import (
	"reflect"
	"testing"
)

func TestURLBuilder_StaticPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name  string
		input URLBuilder
		args  args
		want  URLBuilder
	}{
		{
			name:  "set static path",
			input: URLBuilder{},
			args: args{
				path: "/api/v1",
			},
			want: URLBuilder{
				Path: "/api/v1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.WithStaticPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithStaticPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLBuilder_ParameterizedPath(t *testing.T) {
	type args struct {
		path string
		args []string
	}
	tests := []struct {
		name    string
		input   URLBuilder
		args    args
		want    URLBuilder
		wantErr bool
	}{
		{
			name:  "set parameterized path",
			input: URLBuilder{},
			args: args{
				path: "/api/v1/page/?/comment/?",
				args: []string{"1", "abc"},
			},
			want: URLBuilder{
				Path: "/api/v1/page/1/comment/abc",
			},
			wantErr: false,
		},
		{
			name:  "url encodes path params",
			input: URLBuilder{},
			args: args{
				path: "/api/v1/page/?/comment/?",
				args: []string{"1", "a b/c?foo=bar"},
			},
			want: URLBuilder{
				Path: "/api/v1/page/1/comment/a%20b%2Fc%3Ffoo=bar",
			},
			wantErr: false,
		},
		{
			name:  "too few param args",
			input: URLBuilder{},
			args: args{
				path: "/api/v1/page/?/comment/?",
				args: []string{"1"},
			},
			want:    URLBuilder{},
			wantErr: true,
		},
		{
			name:  "too many param args",
			input: URLBuilder{},
			args: args{
				path: "/api/v1/page/?/comment/?",
				args: []string{"1", "abc", "def"},
			},
			want:    URLBuilder{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.WithPathWithParameters(tt.args.path, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithPathWithParameters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPathWithParameters() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLBuilder_NamedParameterPath(t *testing.T) {
	type args struct {
		path   string
		params map[string]string
	}
	tests := []struct {
		name       string
		input      URLBuilder
		args       args
		want       URLBuilder
		wantErr    bool
		wantErrStr string
	}{
		{
			name:  "set named parameter path",
			input: URLBuilder{},
			args: args{
				path: "/api/v1/page/:page/comment/:commentID",
				params: map[string]string{
					"page":      "1",
					"commentID": "abc",
				},
			},
			want: URLBuilder{
				Path: "/api/v1/page/1/comment/abc",
			},
			wantErr: false,
		},
		{
			name:  "set named parameter path with multiple versions of same param",
			input: URLBuilder{},
			args: args{
				path: "/api/v1/page/:page/page/:page",
				params: map[string]string{
					"page": "1",
				},
			},
			want: URLBuilder{
				Path: "/api/v1/page/1/page/1",
			},
			wantErr: false,
		},
		{
			name:  "set named parameter path with a param key which is a root of another",
			input: URLBuilder{},
			args: args{
				path: "/api/v1/page/:page/pageID/:pageID",
				params: map[string]string{
					"page":   "1",
					"pageID": "abc",
				},
			},
			want: URLBuilder{
				Path: "/api/v1/page/1/pageID/abc",
			},
			wantErr: false,
		},
		{
			name:  "path params are encoded",
			input: URLBuilder{},
			args: args{
				path: "/api/v1/page/:page",
				params: map[string]string{
					"page": "a b/c?foo=bar",
				},
			},
			want: URLBuilder{
				Path: "/api/v1/page/a%20b%2Fc%3Ffoo=bar",
			},
			wantErr: false,
		},
		{
			name:  "missing path params causes an error",
			input: URLBuilder{},
			args: args{
				path:   "/api/v1/page/:page/:commentID",
				params: map[string]string{},
			},
			wantErr:    true,
			wantErrStr: "path parameters in /api/v1/page/:page/:commentID not supplied :page, :commentID",
		},
		{
			name:  "extra path params causes an error",
			input: URLBuilder{},
			args: args{
				path: "/api/v1",
				params: map[string]string{
					"page": "1",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.WithPathWithNamedParameters(tt.args.path, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithPathWithNamedParameters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErrStr != "" && err.Error() != tt.wantErrStr {
				t.Errorf("WithPathWithNamedParameters() error = %v, wantErrStr %v", err, tt.wantErrStr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPathWithNamedParameters() got = %v, want %v", got, tt.want)
			}
		})
	}
}
