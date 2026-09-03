# Copy question templates into each chapter myAnswer directory.
# Existing files (your answers) are never overwritten.
#
#   make questions              copy missing templates into myAnswer/
#   make clean-unchanged-questions        delete myAnswer files that still match the template
#   make clean-all-questions    delete every myAnswer copy (templates stay)

CHAPTERS := $(wildcard [0-9][0-9]-*)
ANSWERS  := myAnswer

.PHONY: help questions clean-unchanged-questions clean-all-questions watch
.DEFAULT_GOAL := help

help:
	@echo "make questions              copy questions/* into $(ANSWERS)/ if the file does not exist"
	@echo "make clean-unchanged-questions        remove unchanged $(ANSWERS) files; keep ones you edited"
	@echo "make clean-all-questions    remove all copies from $(ANSWERS)/ (does not touch questions/)"
	@echo "make watch                  go run a myAnswer file whenever you save it"

questions:
	@copied=0; skipped=0; \
	for ch in $(CHAPTERS); do \
		if [ ! -d "$$ch/questions" ]; then continue; fi; \
		mkdir -p "$$ch/$(ANSWERS)"; \
		for src in "$$ch"/questions/*; do \
			[ -f "$$src" ] || continue; \
			base=$$(basename "$$src"); \
			case "$$base" in README.md|README) continue ;; esac; \
			dest="$$ch/$(ANSWERS)/$$base"; \
			if [ -e "$$dest" ]; then \
				echo "skip $$dest (exists)"; \
				skipped=$$((skipped + 1)); \
			else \
				cp "$$src" "$$dest"; \
				echo "copied $$dest"; \
				copied=$$((copied + 1)); \
			fi; \
		done; \
	done; \
	echo "done: $$copied copied, $$skipped skipped"

clean-unchanged-questions:
	@removed=0; kept=0; \
	for ch in $(CHAPTERS); do \
		if [ ! -d "$$ch/questions" ]; then continue; fi; \
		for src in "$$ch"/questions/*; do \
			[ -f "$$src" ] || continue; \
			base=$$(basename "$$src"); \
			case "$$base" in README.md|README) continue ;; esac; \
			dest="$$ch/$(ANSWERS)/$$base"; \
			if [ ! -e "$$dest" ]; then continue; fi; \
			if cmp -s "$$src" "$$dest"; then \
				rm -f "$$dest"; \
				echo "removed $$dest (unchanged)"; \
				removed=$$((removed + 1)); \
			else \
				echo "keep $$dest (edited)"; \
				kept=$$((kept + 1)); \
			fi; \
		done; \
	done; \
	echo "done: $$removed removed, $$kept kept"

clean-all-questions:
	@removed=0; \
	for ch in $(CHAPTERS); do \
		if [ ! -d "$$ch/questions" ]; then continue; fi; \
		for src in "$$ch"/questions/*; do \
			[ -f "$$src" ] || continue; \
			base=$$(basename "$$src"); \
			case "$$base" in README.md|README) continue ;; esac; \
			dest="$$ch/$(ANSWERS)/$$base"; \
			if [ -e "$$dest" ]; then \
				rm -f "$$dest"; \
				echo "removed $$dest"; \
				removed=$$((removed + 1)); \
			fi; \
		done; \
	done; \
	echo "done: $$removed removed"

watch:
	@bash "$(CURDIR)/watch.sh"
