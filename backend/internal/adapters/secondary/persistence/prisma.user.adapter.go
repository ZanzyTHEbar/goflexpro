package persistence

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/ZanzyTHEbar/goflexpro/internal/dto/db"
	"github.com/ZanzyTHEbar/goflexpro/pkgs/errsx"
	productv1 "github.com/ZanzyTHEbar/goflexpro/pkgs/gen/product/v1"
)

type PrismaProductAdapter struct {
	db *db.PrismaClient
}

func NewPrismaProductAdapter(db *db.PrismaClient) *PrismaProductAdapter {
	return &PrismaProductAdapter{db: db}
}

// Creates a new product
func (pua *PrismaProductAdapter) Create(ctx context.Context, product *productv1.ProductDTO) (*productv1.ProductDTO, error) {
	// Check if the product exists
	lookupKey := db.Product.ID.Equals(int(product.GetId()))
	_, err := pua.ValidateProduct(ctx, lookupKey)
	if err == nil {
		return nil, fmt.Errorf("product already exists")
	}

	// Create a new product
	productModel, err := pua.db.Product.CreateOne(
		db.Product.Name.Set(product.GetName()),
		db.Product.Description.Set(product.GetDescription()),
		db.Product.Price.Set(product.GetPrice()),
	).Exec(ctx)

	if err != nil || productModel == nil {
		return nil, fmt.Errorf("failed to create productModel: %w", err)
	}

	slog.Info("event", "addUser", "productModel created successfully")

	return product, nil
}

func (pua *PrismaProductAdapter) GetById(ctx context.Context, id int) (*productv1.ProductDTO, error) {
	lookupKey := db.Product.ID.Equals(id)
	productModel, err := pua.ValidateProduct(ctx, lookupKey)
	if err != nil {
		return nil, err
	}

	productDTO, err := pua.MapUserProfileModelToProductDTO(productModel)
	if err != nil {
		return nil, err
	}

	return productDTO, nil
}

// GetAll returns all user accounts, we don't need this for the User Service
func (pua *PrismaProductAdapter) GetAll(ctx context.Context) ([]*productv1.ProductDTO, error) {
	return nil, nil
}

func (pua *PrismaProductAdapter) Update(ctx context.Context, product *productv1.ProductDTO) (*productv1.ProductDTO, error) {
	lookupKey := db.Product.ID.Equals(int(product.GetId()))
	_, err := pua.ValidateProduct(ctx, lookupKey)
	if err != nil {
		return nil, err
	}

	productModel, err := pua.db.Product.UpsertOne(
		lookupKey,
	).Update(
		db.Product.Name.Set(product.GetName()),
		db.Product.Description.Set(product.GetDescription()),
		db.Product.Price.Set(product.GetPrice()),
	).Exec(ctx)

	if err != nil || productModel == nil {
		return nil, fmt.Errorf("failed to update account: %w", err)
	}

	return product, nil
}

func (pua *PrismaProductAdapter) Delete(ctx context.Context, id int) error {
	lookupKey := db.Product.ID.Equals(id)
	_, err := pua.ValidateProduct(ctx, lookupKey)
	if err != nil {
		return err
	}

	_, err = pua.db.Product.FindUnique(
		lookupKey,
	).Delete().Exec(ctx)

	if err != nil {
		return fmt.Errorf("failed to delete account: %w", err)
	}

	return nil
}

// ValidateProduct checks if a user exists by their ID. Returns the user if found.
func (pua *PrismaProductAdapter) ValidateProduct(ctx context.Context, key db.ProductEqualsUniqueWhereParam) (*db.ProductModel, error) {
	user, err := pua.db.Product.FindUnique(
		key,
	).Exec(ctx)

	if err != nil || user == nil {
		return nil, fmt.Errorf("failed to validate user: %w", err)
	}

	return user, nil
}

func (pua *PrismaProductAdapter) MapUserProfileModelToProductDTO(productModel *db.ProductModel) (*productv1.ProductDTO, error) {
	if productModel == nil {
		return nil, errsx.NotFoundErr(errors.New("product not found"))
	}

	return &productv1.ProductDTO{
		Id:          int32(productModel.ID),
		Name:        productModel.Name,
		Description: productModel.Description,
		Price:       productModel.Price,
	}, nil
}

func (pua *PrismaProductAdapter) MapProductDTOToUserProfilesModel(product *productv1.ProductDTO) (*db.ProductModel, error) {
	return &db.ProductModel{
		InnerProduct: db.InnerProduct{
			ID:          int(product.GetId()),
			Name:        product.GetName(),
			Description: product.GetDescription(),
			Price:       product.GetPrice(),
		},
	}, nil
}
