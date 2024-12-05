FRONTEND_DIR := frontend
PNPM := pnpm
SHADCN := npx shadcn@latest
OPENAPI := npx openapi-typescript
OPENAPI_PATH := swagger.example.yml

front/dev: front/install
	$(PNPM) --prefix $(FRONTEND_DIR) run dev

front/install:
	$(PNPM) --prefix $(FRONTEND_DIR) install

ui/add/%:
	$(SHADCN) add $* --cwd $(FRONTEND_DIR)

pnpm/add/%:
	$(PNPM) --prefix $(FRONTEND_DIR) add $*

front/api/generate:
	$(OPENAPI) ${OPENAPI_PATH} --output $(FRONTEND_DIR)/src/api/schema.ts
