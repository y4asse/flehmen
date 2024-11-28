FRONTEND_DIR := frontend
PNPM := pnpm
SHADCN := npx shadcn@latest

front/dev: front/install
	$(PNPM) --prefix $(FRONTEND_DIR) run dev

front/install:
	$(PNPM) --prefix $(FRONTEND_DIR) install

ui/add/%:
	$(SHADCN) add $* --cwd $(FRONTEND_DIR)

pnpm/add/%:
	$(PNPM) --prefix $(FRONTEND_DIR) add $*