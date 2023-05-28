BASE_DIR := $(PWD)
FRONTEND_DIR := $(BASE_DIR)/frontend
API_DIR := $(BASE_DIR)/api

define LOGO
\033[31m ____              _                         _        \033[0m
\033[31m| __ )  ___   ___ | | ___ __ ___   __ _ _ __| | _____ \033[0m
\033[31m|  _ \ / _ \ / _ \| |/ / '_ ` _ \ / _` | '__| |/ / __|\033[0m
\033[31m| |_) | (_) | (_) |   <| | | | | | (_| | |  |   <\__ \\ \033[0m
\033[31m|____/ \___/ \___/|_|\_\_| |_| |_|\__,_|_|  |_|\_\___/\033[0m
endef
export LOGO

.PHONY: logo
logo:
	@clear && \
	echo "$$LOGO" && \
	echo ""

.PHONY: info
info:
	@echo "Author: Juyong Shin" && \
	echo ""

check-node:
ifneq ($(shell node --version 2>&1 | grep "v18" > /dev/null; printf $$?),0)
	@echo "No node or version not match" && exit 1
endif

check-python3:
ifneq ($(shell python3 --version 2>&1 | grep "Python" > /dev/null; printf $$?),0)
	@echo "No python or version not match" && exit 1
endif

check-poetry:
ifneq ($(shell poetry --version 2>&1 | grep "Poetry" > /dev/null; printf $$?),0)
	@echo "No poetry or version not match" && exit 1
endif

dependencies: check-node check-python3 check-poetry
	@cd frontend && \
	node --version
	@cd api && \
	python3 --version && \
	poetry about && \
	poetry env info && \
	echo ""

.PHONY: envs
envs:
	@echo BASE_DIR: $(BASE_DIR) && \
	echo FRONTEND_DIR: $(FRONTEND_DIR) && \
	echo API_DIR: $(API_DIR) && \
	echo ""

.PHONY: header
header: logo info dependencies envs

run-fe: header
	@cd frontend && \
    npm run dev

run-api: header
	@cd api && \
	poetry run uvicorn main:app --reload

api-test: header
	@cd api && \
	poetry check && \
	poetry install --with test && \
	poetry run pytest -v -rA

api-doc: header
	@cd api && \
	poetry run sphinx-apidoc -F -o $(API_DIR)/docs $(API_DIR) && \
	poetry run make -C $(API_DIR)/docs html
