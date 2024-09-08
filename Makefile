frontend: 
	@cd frontend && pnpm i && pnpm dev

backend:
	@cd backend && make run

.PHONY: frontend backend