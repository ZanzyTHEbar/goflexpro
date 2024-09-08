import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { ProductService } from "@static/types/gen/product/v1/product_connect";

// The transport defines what type of endpoint we're hitting.
// In our example we'll be communicating with a Connect endpoint.
// If your endpoint only supports gRPC-web, make sure to use
// `createGrpcWebTransport` instead.
const transport = createConnectTransport({
    baseUrl: "http://localhost:9000",
});

// Set up the Connect client
// Here we make the client itself, combining the service
// definition with the transport.
export const connectClient = createPromiseClient(ProductService, transport);
