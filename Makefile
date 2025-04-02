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
	N=
	@if [ -z "$(N)" ]; then \
		echo "$(RED)N must be not empty$(NC)"; \
		exit 1; \
	fi
	@echo "Generating easy solution"
	@leetscrape -n $(N) sol -t easy -l go -o CompetitiveProgramming/leetcode/problem
medium:
	N=
	@if [ -z "$(N)" ]; then \
		echo "$(RED)N must be not empty$(NC)"; \
		exit 1; \
	fi
	@echo "Generating medium solution"
	@leetscrape -n $(N) sol -t medium -l go -o CompetitiveProgramming/leetcode/problem
hard:
	N=
	@if [ -z "$(N)" ]; then \
		echo "$(RED)N must be not empty$(NC)"; \
		exit 1; \
	fi
	@echo "Generating hard solution"
	@leetscrape -n $(N) sol -t hard -l go -o CompetitiveProgramming/leetcode/problem
