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

	slog.Info("Create", "event", "dbCreate", "details", "productModel created successfully")

	return product, nil
}

func (pua *PrismaProductAdapter) GetById(ctx context.Context, id int) (*productv1.ProductDTO, error) {
	lookupKey := db.Product.ID.Equals(id)
	productModel, err := pua.ValidateProduct(ctx, lookupKey)
	if err != nil {
		slog.Error("GetById", "error", err)
		return nil, err
	}

	productDTO, err := pua.MapproductProfileModelToProductDTO(productModel)
	if err != nil {
		slog.Error("GetById", "error", err)
		return nil, err
	}

	slog.Info("GetById", "event", "dbGet", "details", "productModel retrieved successfully")

	return productDTO, nil
}

// GetAll returns all product accounts, we don't need this for the product Service
func (pua *PrismaProductAdapter) GetAll(ctx context.Context) ([]*productv1.ProductDTO, error) {
	products, err := pua.db.Product.FindMany().Exec(ctx)
	if err != nil {
		slog.Error("GetAll", "error", err)
		return nil, fmt.Errorf("failed to get all products: %w", err)
	}

	productDTOs := make([]*productv1.ProductDTO, 0)
	for _, product := range products {
		productDTO, err := pua.MapproductProfileModelToProductDTO(&product)
		if err != nil {
			slog.Error("GetAll", "error", err)
			return nil, err
		}

		productDTOs = append(productDTOs, productDTO)
	}

	slog.Info("GetAll", "event", "dbGetAll", "details", "products retrieved successfully")

	return productDTOs, nil
}

func (pua *PrismaProductAdapter) Update(ctx context.Context, product *productv1.ProductDTO) (*productv1.ProductDTO, error) {
	lookupKey := db.Product.ID.Equals(int(product.GetId()))
	_, err := pua.ValidateProduct(ctx, lookupKey)
	if err != nil {
		slog.Error("Update", "error", err)
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
		slog.Error("Update", "error", err)
		return nil, fmt.Errorf("failed to update account: %w", err)
	}

	slog.Info("Update", "event", "dbUpdate", "details", "productModel updated successfully")

	return product, nil
}

func (pua *PrismaProductAdapter) Delete(ctx context.Context, id int) error {
	lookupKey := db.Product.ID.Equals(id)
	_, err := pua.ValidateProduct(ctx, lookupKey)
	if err != nil {
		slog.Error("Delete", "error", err)
		return err
	}

	_, err = pua.db.Product.FindUnique(
		lookupKey,
	).Delete().Exec(ctx)

	if err != nil {
		slog.Error("Delete", "error", err)
		return fmt.Errorf("failed to delete account: %w", err)
	}

	slog.Info("Delete", "event", "dbDelete", "details", "productModel deleted successfully")

	return nil
}

// ValidateProduct checks if a product exists by their ID. Returns the product if found.
func (pua *PrismaProductAdapter) ValidateProduct(ctx context.Context, key db.ProductEqualsUniqueWhereParam) (*db.ProductModel, error) {
	product, err := pua.db.Product.FindUnique(
		key,
	).Exec(ctx)

	if err != nil || product == nil {
		slog.Error("ValidateProduct", "error", err)
		return nil, fmt.Errorf("failed to validate product: %w", err)
	}

	slog.Debug("ValidateProduct", "event", "dbValidate", "details", "productModel validated successfully")

	return product, nil
}

func (pua *PrismaProductAdapter) MapproductProfileModelToProductDTO(productModel *db.ProductModel) (*productv1.ProductDTO, error) {
	if productModel == nil {
		slog.Error("MapproductProfileModelToProductDTO", "error", errsx.NotFoundErr(errors.New("product not found")))
		return nil, errsx.NotFoundErr(errors.New("product not found"))
	}

	slog.Debug("MapproductProfileModelToProductDTO", "event", "dbMap", "details", "productModel mapped successfully")

	return &productv1.ProductDTO{
		Id:          int32(productModel.ID),
		Name:        productModel.Name,
		Description: productModel.Description,
		Price:       productModel.Price,
	}, nil
}

func (pua *PrismaProductAdapter) MapProductDTOToproductProfilesModel(product *productv1.ProductDTO) (*db.ProductModel, error) {

	slog.Debug("MapProductDTOToproductProfilesModel", "event", "dbMap", "details", "productModel mapped successfully")

	return &db.ProductModel{
		InnerProduct: db.InnerProduct{
			ID:          int(product.GetId()),
			Name:        product.GetName(),
			Description: product.GetDescription(),
			Price:       product.GetPrice(),
		},
	}, nil
}
