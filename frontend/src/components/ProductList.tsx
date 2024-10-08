import { For, type ParentComponent, createSignal, onMount } from 'solid-js'
import { GetProductRequest, ProductDTO } from '@static/types/gen/product/v1/product_pb'
import { connectClient } from '@src/api/productService'

const ProductList: ParentComponent = () => {
    const [products, setProducts] = createSignal<ProductDTO[]>([])

    onMount(async () => {
        const req = new GetProductRequest({
            id: [],
        })

        const res = await connectClient.getProduct(req)

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
                            <p>{product.price}</p>
                        </li>
                    )}
                </For>
            </ul>
        </div>
    )
}

export default ProductList
