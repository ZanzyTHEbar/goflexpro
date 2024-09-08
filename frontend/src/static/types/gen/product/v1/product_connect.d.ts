// @generated by protoc-gen-connect-es v1.4.0
// @generated from file product/v1/product.proto (package product.v1, syntax proto3)
/* eslint-disable */


import { CreateProductRequest, CreateProductResponse, DeleteProductRequest, DeleteProductResponse, GetProductRequest, GetProductResponse, UpdateProductRequest, UpdateProductResponse } from "./product_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service product.v1.ProductService
 */
export declare const ProductService: {
  readonly typeName: "product.v1.ProductService",
  readonly methods: {
    /**
     * @generated from rpc product.v1.ProductService.CreateProduct
     */
    readonly createProduct: {
      readonly name: "CreateProduct",
      readonly I: typeof CreateProductRequest,
      readonly O: typeof CreateProductResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc product.v1.ProductService.GetProduct
     */
    readonly getProduct: {
      readonly name: "GetProduct",
      readonly I: typeof GetProductRequest,
      readonly O: typeof GetProductResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc product.v1.ProductService.UpdateProduct
     */
    readonly updateProduct: {
      readonly name: "UpdateProduct",
      readonly I: typeof UpdateProductRequest,
      readonly O: typeof UpdateProductResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc product.v1.ProductService.DeleteProduct
     */
    readonly deleteProduct: {
      readonly name: "DeleteProduct",
      readonly I: typeof DeleteProductRequest,
      readonly O: typeof DeleteProductResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

