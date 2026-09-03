# Copy question templates into each chapter myAnswer directory.
# Existing files (your answers) are never overwritten.
#
#   make next-question          copy the next missing question into myAnswer/
#   make questions              copy missing templates into myAnswer/
#   make clean-unchanged-questions        delete myAnswer files that still match the template
#   make clean-all-questions    delete every myAnswer copy (templates stay)

CHAPTERS := $(wildcard [0-9][0-9]-*)
ANSWERS  := myAnswer

.PHONY: help next-question questions clean-unchanged-questions clean-all-questions watch
.DEFAULT_GOAL := help

help:
	@echo "make next-question          copy the next missing question into $(ANSWERS)/ and print the path"
	@echo "make questions              copy questions/* into $(ANSWERS)/ if the file does not exist"
	@echo "make clean-unchanged-questions        remove unchanged $(ANSWERS) files; keep ones you edited"
	@echo "make clean-all-questions    remove all copies from $(ANSWERS)/ (does not touch questions/)"
	@echo "make watch                  go run a myAnswer file whenever you save it"

# First question whose myAnswer copy is missing. Companion files that share
# the same id (e.g. 11.2-add.go and 11.2-add_test.go) are copied together.
next-question:
	@found=0; \
	for ch in $(CHAPTERS); do \
		if [ ! -d "$$ch/questions" ]; then continue; fi; \
		mkdir -p "$$ch/$(ANSWERS)"; \
		for src in "$$ch"/questions/*; do \
			[ -f "$$src" ] || continue; \
			base=$$(basename "$$src"); \
			case "$$base" in README.md|README|*.expected.txt) continue ;; esac; \
			dest="$$ch/$(ANSWERS)/$$base"; \
			if [ -e "$$dest" ]; then continue; fi; \
			id=$${base%%-*}; \
			for s in "$$ch"/questions/"$$id"-*; do \
				[ -f "$$s" ] || continue; \
				b=$$(basename "$$s"); \
				case "$$b" in *.expected.txt) continue ;; esac; \
				d="$$ch/$(ANSWERS)/$$b"; \
				if [ -e "$$d" ]; then continue; fi; \
				cp "$$s" "$$d"; \
				echo "copied $$d"; \
				found=1; \
			done; \
			break 2; \
		done; \
	done; \
	if [ "$$found" -eq 0 ]; then \
		echo "no next question (all already copied)"; \
	fi

questions:
	@copied=0; skipped=0; \
	for ch in $(CHAPTERS); do \
		if [ ! -d "$$ch/questions" ]; then continue; fi; \
		mkdir -p "$$ch/$(ANSWERS)"; \
		for src in "$$ch"/questions/*; do \
			[ -f "$$src" ] || continue; \
			base=$$(basename "$$src"); \
			case "$$base" in README.md|README|*.expected.txt) continue ;; esac; \
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
			case "$$base" in README.md|README|*.expected.txt) continue ;; esac; \
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
			case "$$base" in README.md|README|*.expected.txt) continue ;; esac; \
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
