# Copy question templates into each chapter directory.
# Existing files (your answers) are never overwritten.
#
#   make questions         copy missing templates into chapter dirs
#   make clean-questions   delete those working copies (templates stay)

CHAPTERS := $(wildcard [0-9][0-9]-*)

.PHONY: help questions clean-questions
.DEFAULT_GOAL := help

help:
	@echo "make questions         copy questions/*.go into the chapter dir if the file does not exist"
	@echo "make clean-questions   remove those copies from the chapter dir (does not touch questions/)"

questions:
	@copied=0; skipped=0; \
	for ch in $(CHAPTERS); do \
		if [ ! -d "$$ch/questions" ]; then continue; fi; \
		for src in "$$ch"/questions/*; do \
			[ -f "$$src" ] || continue; \
			base=$$(basename "$$src"); \
			case "$$base" in README.md|README) continue ;; esac; \
			dest="$$ch/$$base"; \
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
			dest="$$ch/$$base"; \
			if [ -e "$$dest" ]; then \
				rm -f "$$dest"; \
				echo "removed $$dest"; \
				removed=$$((removed + 1)); \
			fi; \
		done; \
	done; \
	echo "done: $$removed removed"
