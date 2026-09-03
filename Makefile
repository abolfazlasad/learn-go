# Copy question templates into each chapter myAnswer directory.
# Existing files (your answers) are never overwritten.
#
#   make questions         copy missing templates into myAnswer/
#   make clean-questions   delete working copies in myAnswer/ (templates stay)

CHAPTERS := $(wildcard [0-9][0-9]-*)
ANSWERS  := myAnswer

.PHONY: help questions clean-questions
.DEFAULT_GOAL := help

help:
	@echo "make questions         copy questions/* into $(ANSWERS)/ if the file does not exist"
	@echo "make clean-questions   remove copies from $(ANSWERS)/ (does not touch questions/)"

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

clean-questions:
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
