// @generated by protoc-gen-es v2.0.0
// @generated from file product/v1/product.proto (package product.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file product/v1/product.proto.
 */
export declare const file_product_v1_product: GenFile;

/**
 * @generated from message product.v1.ProductDTO
 */
export declare type ProductDTO = Message<"product.v1.ProductDTO"> & {
  /**
   * @generated from field: int32 id = 1;
   */
  id: number;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: string description = 3;
   */
  description: string;

  /**
   * @generated from field: double price = 4;
   */
  price: number;

  /**
   * @generated from field: string created_at = 5;
   */
  createdAt: string;
};

/**
 * Describes the message product.v1.ProductDTO.
 * Use `create(ProductDTOSchema)` to create a new message.
 */
export declare const ProductDTOSchema: GenMessage<ProductDTO>;

/**
 * @generated from message product.v1.CreateProductRequest
 */
export declare type CreateProductRequest = Message<"product.v1.CreateProductRequest"> & {
  /**
   * @generated from field: repeated product.v1.ProductDTO product = 1;
   */
  product: ProductDTO[];
};

/**
 * Describes the message product.v1.CreateProductRequest.
 * Use `create(CreateProductRequestSchema)` to create a new message.
 */
export declare const CreateProductRequestSchema: GenMessage<CreateProductRequest>;

/**
 * @generated from message product.v1.CreateProductResponse
 */
export declare type CreateProductResponse = Message<"product.v1.CreateProductResponse"> & {
  /**
   * @generated from field: repeated bool success = 1;
   */
  success: boolean[];
};

/**
 * Describes the message product.v1.CreateProductResponse.
 * Use `create(CreateProductResponseSchema)` to create a new message.
 */
export declare const CreateProductResponseSchema: GenMessage<CreateProductResponse>;

/**
 * @generated from message product.v1.GetProductRequest
 */
export declare type GetProductRequest = Message<"product.v1.GetProductRequest"> & {
  /**
   * @generated from field: repeated int32 id = 1;
   */
  id: number[];
};

/**
 * Describes the message product.v1.GetProductRequest.
 * Use `create(GetProductRequestSchema)` to create a new message.
 */
export declare const GetProductRequestSchema: GenMessage<GetProductRequest>;

/**
 * @generated from message product.v1.GetProductResponse
 */
export declare type GetProductResponse = Message<"product.v1.GetProductResponse"> & {
  /**
   * @generated from field: repeated product.v1.ProductDTO product = 1;
   */
  product: ProductDTO[];
};

/**
 * Describes the message product.v1.GetProductResponse.
 * Use `create(GetProductResponseSchema)` to create a new message.
 */
export declare const GetProductResponseSchema: GenMessage<GetProductResponse>;

/**
 * @generated from message product.v1.UpdateProductRequest
 */
export declare type UpdateProductRequest = Message<"product.v1.UpdateProductRequest"> & {
  /**
   * @generated from field: repeated product.v1.ProductDTO product = 1;
   */
  product: ProductDTO[];
};

/**
 * Describes the message product.v1.UpdateProductRequest.
 * Use `create(UpdateProductRequestSchema)` to create a new message.
 */
export declare const UpdateProductRequestSchema: GenMessage<UpdateProductRequest>;

/**
 * @generated from message product.v1.UpdateProductResponse
 */
export declare type UpdateProductResponse = Message<"product.v1.UpdateProductResponse"> & {
  /**
   * @generated from field: repeated bool success = 1;
   */
  success: boolean[];
};

/**
 * Describes the message product.v1.UpdateProductResponse.
 * Use `create(UpdateProductResponseSchema)` to create a new message.
 */
export declare const UpdateProductResponseSchema: GenMessage<UpdateProductResponse>;

/**
 * @generated from message product.v1.DeleteProductRequest
 */
export declare type DeleteProductRequest = Message<"product.v1.DeleteProductRequest"> & {
  /**
   * @generated from field: repeated int32 id = 1;
   */
  id: number[];
};

/**
 * Describes the message product.v1.DeleteProductRequest.
 * Use `create(DeleteProductRequestSchema)` to create a new message.
 */
export declare const DeleteProductRequestSchema: GenMessage<DeleteProductRequest>;

/**
 * @generated from message product.v1.DeleteProductResponse
 */
export declare type DeleteProductResponse = Message<"product.v1.DeleteProductResponse"> & {
  /**
   * @generated from field: repeated bool success = 1;
   */
  success: boolean[];
};

/**
 * Describes the message product.v1.DeleteProductResponse.
 * Use `create(DeleteProductResponseSchema)` to create a new message.
 */
export declare const DeleteProductResponseSchema: GenMessage<DeleteProductResponse>;

/**
 * @generated from service product.v1.ProductService
 */
export declare const ProductService: GenService<{
  /**
   * @generated from rpc product.v1.ProductService.CreateProduct
   */
  createProduct: {
    methodKind: "unary";
    input: typeof CreateProductRequestSchema;
    output: typeof CreateProductResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.GetProduct
   */
  getProduct: {
    methodKind: "unary";
    input: typeof GetProductRequestSchema;
    output: typeof GetProductResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.UpdateProduct
   */
  updateProduct: {
    methodKind: "unary";
    input: typeof UpdateProductRequestSchema;
    output: typeof UpdateProductResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.DeleteProduct
   */
  deleteProduct: {
    methodKind: "unary";
    input: typeof DeleteProductRequestSchema;
    output: typeof DeleteProductResponseSchema;
  },
}>;
