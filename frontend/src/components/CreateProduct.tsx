import { type ParentComponent, createEffect, createSignal } from 'solid-js'
import { CreateProductRequest, ProductDTO } from '@static/types/gen/product/v1/product_pb'
import { connectClient } from '@src/api/productService'

const CreateProduct: ParentComponent = () => {
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

        const newProduct: ProductDTO = new ProductDTO({
            name: name(),
            description: description(),
            price: price(),
        })

        const req = new CreateProductRequest({
            product: [newProduct],
        })

        console.log('req', req)

        const res = await connectClient.createProduct(req)
        console.log('res', res)
    }

    return (
        <form onSubmit={onSubmit}>
            <div>
                <label>Name:</label>
                <input name="name" value={name()} onInput={onChange} />
            </div>
            <br />
            <div>
                <label>Description:</label>
                <input name="description" value={description()} onInput={onChange} />
            </div>
            <br />
            <div>
                <label>Price:</label>
                <input name="price" value={price()} type="number" onInput={onChange} />
            </div>
            <br />
            <button type="submit">CreateProduct</button>
        </form>
    )
}

export default CreateProduct
