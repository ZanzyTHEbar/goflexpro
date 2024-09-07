import { type ParentComponent, createEffect, createSignal, onMount } from 'solid-js'
import { create } from '@bufbuild/protobuf'
import {
    CreateProductRequest,
    ProductDTOSchema,
    ProductDTO,
} from '../static/types/gen/product/v1/product_pb.js'
import { connectClient } from '../api/productService'

const CreateProduct: ParentComponent = () => {
    const [product, setProduct] = createSignal<ProductDTO>()
    const [name, setName] = createSignal('')
    const [description, setDescription] = createSignal('')
    const [price, setPrice] = createSignal(0)

    const onChange = (e: Event) => {
        const target = e.target as HTMLInputElement
        const { name, value } = target
        switch (name) {
            case 'name':
                setName(value)
                break
            case 'description':
                setDescription(value)
                break
            case 'price':
                setPrice(Number(value))
                break
        }
    }

    const onSubmit = async (e: Event) => {
        e.preventDefault()

        const req = {
            $typeName: 'product.v1.CreateProductRequest',
            product: [product()],
        } as CreateProductRequest

        await connectClient.createProduct(req)
    }

    createEffect(() => {
        const newProduct = create(ProductDTOSchema, {
            name: name(),
            description: description(),
            price: price(),
        })
        setProduct(newProduct)
    })

    return (
        <form onSubmit={onSubmit}>
            <div>
                <label>Name:</label>
                <input value={name()} onInput={onChange} />
            </div>
            <br />
            <div>
                <label>Description:</label>
                <input value={description()} onInput={onChange} />
            </div>
            <br />
            <div>
                <label>Price:</label>
                <input value={price()} type="number" onInput={onChange} />
            </div>
            <br />
            <button type="submit">CreateProduct</button>
        </form>
    )
}

export default CreateProduct
