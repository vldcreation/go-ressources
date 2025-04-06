.PHONY: AddSolution

#COLOR CONSTANTS
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[0;33m
NC=\033[0m # No Color

AddSolution:
	@echo "Adding new solution to the project"
	@bash add-solution.sh

.PHONY: easy

easy:
	n=
	N=
	@if [ -z "$(n)" ] && [ -z "$(N)" ]; then \
		echo "$(RED)Either n or N must be not empty$(NC)"; \
		exit 1; \
	fi
	@echo "Generating easy solution"
	@if [ ! -z "$(n)" ]; then \
		leetscrape -n $(n) sol -t easy -l go -o CompetitiveProgramming/leetcode/problem; \
	elif [ ! -z "$(N)" ]; then \
		leetscrape -N $(N) sol -t easy -l go -o CompetitiveProgramming/leetcode/problem; \
	fi

medium:
	n=
	N=
	@if [ -z "$(n)" ] && [ -z "$(N)" ]; then \
		echo "$(RED)Either n or N must be not empty$(NC)"; \
		exit 1; \
	fi
	@echo "Generating medium solution"
	@if [ ! -z "$(n)" ]; then \
		leetscrape -n $(n) sol -t medium -l go -o CompetitiveProgramming/leetcode/problem; \
	elif [ ! -z "$(N)" ]; then \
		leetscrape -N $(N) sol -t medium -l go -o CompetitiveProgramming/leetcode/problem; \
	fi

hard:
	n=
	N=
	@if [ -z "$(n)" ] && [ -z "$(N)" ]; then \
		echo "$(RED)Either n or N must be not empty$(NC)"; \
		exit 1; \
	fi
	@echo "Generating hard solution"
	@if [ ! -z "$(n)" ]; then \
		leetscrape -n $(n) sol -t hard -l go -o CompetitiveProgramming/leetcode/problem; \
	elif [ ! -z "$(N)" ]; then \
		leetscrape -N $(N) sol -t hard -l go -o CompetitiveProgramming/leetcode/problem; \
	fi
