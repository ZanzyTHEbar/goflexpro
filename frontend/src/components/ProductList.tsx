import { type ParentComponent, createSignal, onMount } from 'solid-js'
import { GetProductResponse, ProductDTO } from '../static/types/gen/product/v1/product_pb'
import { connectClient } from '../api/productService'

const ProductList: ParentComponent = () => {
    const [products, setProducts] = createSignal<ProductDTO[]>([])

    onMount(async () => {
        const res = await connectClient.getProduct({ id: [1] })
        setProducts(res.product)
    })

    return (
        <div>
            <h1>Product List</h1>
        </div>
    )
}
