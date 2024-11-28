FRONTEND_DIR := frontend
PNPM := pnpm
SHADCN := npx shadcn@latest

front/dev: front/install
	$(PNPM) --prefix $(FRONTEND_DIR) run dev

front/install:
	$(PNPM) --prefix $(FRONTEND_DIR) install

add/ui/%:
	$(SHADCN) add $* --cwd $(FRONTEND_DIR)