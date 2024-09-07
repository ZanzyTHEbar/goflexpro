# GoFlexPro - Flexible, Fast, and Schema-based Stack for Go

GoFlexPro is a powerful (_and opinionated_) stack designed for developers building fast, flexible, and schema-based applications in Go. With `Prisma-Client-Go` and `BufBuild` at its core, it provides an efficient way to manage both the API and database layers using type-safe, schema-driven tools.

With a lot of hype in the Go community around `htmx` and SSR, GoFlexPro is a breath of fresh air for developers looking to build robust, scalable, and maintainable backend systems in Go without having to compromise on the DX of a modern stack.

Some projects simply _do not require SSR or htmx_, and for those projects, GoFlexPro is a perfect choice, with the flexability to custom tooling on-top later (such as SSR for API endpoints).

My goal in creating the GoFlexPro stack is to provide a solid foundation for building Go applications that are both performant and developer-friendly. This is a stack that I use for my own projects, and I hope it will be useful for others as well.

## Key Features

- Full shared-type type-safety for API, database, and frontend layers - one source of truth.
- gRPC and HTTP support with BufBuild-Connect
- Cross-platform ORM with Prisma-Client-Go
- Minimized boilerplate for faster development
- Schema-driven development for consistency and maintainability

## Why GoFlexPro?

This stack eliminates the need for glue code and boilerplate, letting you focus on service implementation. It's designed for modern backend development, offering both flexibility and speed while keeping everything strictly typed and shared across your application layers.

## Design Decisions

GoFlexPro is designed to be a robust, scalable, and maintainable stack for Go applications. To achieve this, I have made some key design decisions to aid in scalability, performance, and developer productivity:

**Go**: Go is a perfect choice for building scalable, performant backend systems. It's known for its simplicity, speed, and concurrency support. By using Go, you’re already setting a solid foundation for performance and scalability.

[**Prisma-Client-Go**](https://goprisma.org/docs): Prisma is a powerful ORM with a strong emphasis on developer experience, type safety, and performance. Prisma’s ability to generate schema-based, type-safe clients is a huge plus, especially in reducing errors and improving productivity. It also brings multi-database support, which gives your stack flexibility in working across different database systems. The recent development of `Prisma-Client-Go` for the Go ecosystem is a welcome addition, making it easier to integrate Prisma into Go projects.

**BufBuild-[`Connect RPC`](https://connectrpc.com/) & [`BufBuild-Buf`](https://buf.build/docs/installation)**: Schema-driven generation of type-compliant clients and services. Leveraging gRPC and HTTP simultaneously makes the stack highly flexible for communication protocols. Using BufBuild-Buf to automate the protoc workflow is a nice touch for DX, minimizing setup hassles while maintaining a strict contract-first approach. This protocol flexibility allows for a future-proof system that can evolve easily.

**Schema-Driven, Type-Safe Development**: The type-safe and schema-driven nature of both the API (via BufBuild) and the ORM (via Prisma) is an excellent strategy. It minimizes runtime errors, ensures consistency, and adds an extra layer of safety that boosts developer confidence. This approach also contributes to long-term maintainability.

**Reduced Boilerplate**: One of the most attractive aspects of this stack is how it minimizes boilerplate code. The use of Prisma and BufBuild tools automates much of the tedious work, letting developers focus on actual business logic and service implementation. This leads to faster development cycles and fewer chances of human error.

### Technical Perspective

1. **Type-Safe, Schema-Driven Development**: By using `Prisma-Client-Go` and `BufBuild`, the stack enforces a type-safe and schema-driven approach. This minimizes runtime errors, ensures consistency, and boosts developer confidence.
    - Creates a single source of truth for the API and database schemas, allowing code generation to facilitate communication between the layers.
2. **Domain Driven Design**: The stack is designed to follow domain-driven design principles, with clear separation of concerns and a focus on business logic.
    - Encourages a modular, maintainable codebase that is easier to understand and extend.
    - Implements the idea of Data Transfer Objects (DTOs) to encapsulate data and reduce coupling between layers.
3. **Hexangonal Architecture (a.k.a: Ports & Adapters Pattern)**: The stack follows a hexagonal architecture, separating the core business logic from the infrastructure and delivery mechanisms.
    - Promotes testability, maintainability, and flexibility in the application design.
    - Allows for easy swapping of components and adaptors without affecting the core logic.
4. **gRPC and HTTP Support**: The stack leverages the awesome protocol developed by the `bufbuild` team; `Connect RPC`. `Connect` supports both gRPC and HTTP communication protocols, providing flexibility in how services interact with each other.
    - Enables efficient, high-performance communication between services using gRPC.
    - Allows for broader compatibility and ease of use with HTTP endpoints.
    - No need for swagger or openAPI documentation, as the schema is the source of truth.
    - Fully cross-language compatible, enabling seamless integration with other services.
5. **Automated Code Generation**: The stack uses `BufBuild` to automate the generation of protocol buffers, a `Connect` client, and `Connect` services, reducing boilerplate code and ensuring consistency across services.
    - Minimizes manual work and potential errors in maintaining protocol definitions.
    - Provides a contract-first approach to API development, ensuring that services adhere to a common schema.
6. **Prisma ORM**: The stack uses `Prisma-Client-Go` as the ORM for database interactions, providing a type-safe, schema-driven way to interact with the database.
    - Reduces the complexity of database interactions and ensures type safety.
    - Supports multiple databases, allowing for flexibility in choosing the right database for the application without requiring a total rewrite.

### Business Perspective

**Professional**: The “pro” in GoFlexPro is not just in the name—it’s embedded in the stack's design. GoFlexPro is a serious, production-grade setup that prioritizes long-term maintainability and performance. There’s a level of intentionality here, especially with the type-safe and schema-driven approach, that speaks to a professional mindset.

**Balance of Flexibility and Efficiency**: This stack achieves a very solid balance between flexibility and raw performance. `Connect`'s dual support for gRPC and HTTP and Prisma’s multi-database compatibility give you options without sacrificing speed.

**Modern and Future-Proof**: We are using modern tools, built on well-tested industry accepted practices and protocols, that are both powerful and scalable. BufBuild is increasingly becoming a standard for Protobuf management, and Prisma has a strong track record of community support and updates. This ensures GoFlexPro will continue to evolve with the industry, rather than becoming obsolete.

### Potential Areas of Consideration

Learning Curve: For developers not familiar with the specific tools like Prisma or BufBuild, there may be a learning curve. However, this is a minor issue if you're building with experienced Go developers or willing to invest a bit in onboarding.

Prisma in Go: Prisma is traditionally more popular in JavaScript/TypeScript ecosystems, so the Go implementation (Prisma-Client-Go) is relatively newer and may have some less-explored edges. It would be worth keeping an eye on community adoption and support over time. Consider supporting the developer of Prisma-Client-Go on [GitHub](https://github.com/steebchen/prisma-client-go).

## Quick Start

## Example Project

## Contributing

## License
