package services

import (
	"context"

	"connectrpc.com/connect"
	"github.com/ZanzyTHEbar/goflexpro/internal/ports"
	productv1 "github.com/ZanzyTHEbar/goflexpro/pkgs/gen/product/v1"
)

type ProductService struct {
	productRepo ports.ProductRepoPort
}

func NewProductService(productRepo ports.ProductRepoPort) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

// CreateProduct creates a new product
func (ps *ProductService) CreateProduct(ctx context.Context, req *connect.Request[productv1.CreateProductRequest]) (*connect.Response[productv1.CreateProductResponse], error) {
	// The request comes in as an array of products.

	// req.Msg.GetProduct() returns a slice of pointers to ProductDTO
	newProducts := req.Msg.GetProduct()
	success := make([]bool, 0)

	for _, product := range newProducts {
		_, err := ps.productRepo.Create(ctx, product)
		if err != nil {
			success = append(success, false)
			continue
		}
		success = append(success, true)
	}

	// construct the response
	res := connect.NewResponse(&productv1.CreateProductResponse{
		Success: success,
	})

	return res, nil
}

// GetProduct gets a product by id
func (ps *ProductService) GetProduct(ctx context.Context, req *connect.Request[productv1.GetProductRequest]) (*connect.Response[productv1.GetProductResponse], error) {
	// The request comes in as an array of product ids

	// req.Msg.GetId() returns a slice of strings
	productIDs := req.Msg.GetId()
	products := make([]*productv1.ProductDTO, 0)

	for _, id := range productIDs {
		product, err := ps.productRepo.GetById(ctx, int(id))
		if err != nil {
			// Handle the error
			continue
		}
		products = append(products, product)
	}

	// construct the response
	res := connect.NewResponse(&productv1.GetProductResponse{
		Product: products,
	})

	return res, nil
}

// UpdateProduct updates a product
func (ps *ProductService) UpdateProduct(ctx context.Context, req *connect.Request[productv1.UpdateProductRequest]) (*connect.Response[productv1.UpdateProductResponse], error) {
	// The request comes in as an array of products.

	// req.Msg.GetProduct() returns a slice of pointers to ProductDTO
	updatedProducts := req.Msg.GetProduct()
	success := make([]bool, 0)

	for _, product := range updatedProducts {
		_, err := ps.productRepo.Update(ctx, product)
		if err != nil {
			success = append(success, false)
			continue
		}
		success = append(success, true)
	}

	// construct the response
	res := connect.NewResponse(&productv1.UpdateProductResponse{
		Success: success,
	})

	return res, nil

}

// DeleteProduct deletes a product
func (ps *ProductService) DeleteProduct(ctx context.Context, req *connect.Request[productv1.DeleteProductRequest]) (*connect.Response[productv1.DeleteProductResponse], error) {
	// The request comes in as an array of product ids

	// req.Msg.GetId() returns a slice of strings
	productIDs := req.Msg.GetId()
	success := make([]bool, 0)

	for _, id := range productIDs {
		err := ps.productRepo.Delete(ctx, int(id))
		if err != nil {
			success = append(success, false)
			continue
		}
		success = append(success, true)
	}

	// construct the response
	res := connect.NewResponse(&productv1.DeleteProductResponse{
		Success: success,
	})

	return res, nil

}
