package resolver

import (
	"context"
	"go-template/gqlmodels"
	"reflect"
	"testing"
)

func Test_queryResolver_GetPosts(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *gqlmodels.PostPayload
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.GetPosts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResolver.GetPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResolver.GetPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
