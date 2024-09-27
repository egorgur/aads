# lint with Ruff
.PHONY: lint
lint:
	poetry run ruff check --fix

.PHONY: ipython
ipython: 
	poetry run ipython