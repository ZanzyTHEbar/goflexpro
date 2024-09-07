package ports

import (
	productv1 "github.com/ZanzyTHEbar/goflexpro/pkgs/gen/product/v1"
)

type ProductRepoPort interface {
	DBPort[productv1.ProductDTO]
}
