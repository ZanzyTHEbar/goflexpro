import type { Component } from 'solid-js'

import { Route, Router } from '@solidjs/router'
import ProductList from './components/ProductList'
import CreateProduct from './components/CreateProduct'

const App: Component = () => {
    return (
        <Router>
            <Route path="/" component={ProductList} />
            <Route path="/create" component={CreateProduct} />
        </Router>
    )
}

export default App
