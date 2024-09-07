import { For, type ParentComponent, createSignal, onMount } from 'solid-js'
import {
    GetProductResponse,
    GetProductRequest,
    ProductDTO,
} from '../static/types/gen/product/v1/product_pb.js'
import { connectClient } from '../api/productService'

interface ProductListReq extends GetProductRequest {}

const ProductList: ParentComponent = () => {
    const [products, setProducts] = createSignal<ProductDTO[]>([])

    onMount(async () => {
        const req = {
            $typeName: 'product.v1.GetProductRequest',
            id: [],
        } as GetProductRequest

        const res = (await connectClient.getProduct(req)) as GetProductResponse
        setProducts(res.product)
    })

    return (
        <div>
            <h1>Product List</h1>
            <ul>
                <For each={products()}>
                    {(product) => (
                        <li>
                            <h2>{product.name}</h2>
                            <p>{product.description}</p>
                        </li>
                    )}
                </For>
            </ul>
        </div>
    )
}

export default ProductList
