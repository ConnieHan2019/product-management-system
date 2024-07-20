package service

import (
	"product-management-system/pkg/dtos"
	"product-management-system/pkg/model"
	"reflect"
	"testing"

	"github.com/go-logr/logr"
	"gorm.io/gorm"
)

func TestProductService_CreateProduct(t *testing.T) {
	type fields struct {
		DB     *gorm.DB
		Logger logr.Logger
	}
	type args struct {
		productEntity *dtos.Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ProductService{
				DB:     tt.fields.DB,
				Logger: tt.fields.Logger,
			}
			got, err := ps.CreateProduct(tt.args.productEntity)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
