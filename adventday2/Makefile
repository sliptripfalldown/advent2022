.DEFAULT_GOAL = help
ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

help:
> @grep -hE '^[A-Za-z0-9_ \-]*?:.*##.*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

go: ## Runs Go version of Advent of Code
> @cd ./go && go run main.go
.PHONY: go

python: ## Runs Python version of Advent of Code
> @cd ./python && python3 main.py
.PHONY: python

#rust: ## Runs Rust version of Advent of Code
#> @cd ./rust && rustc main.rs && ./main.exe
#.PHONY: clean rust

#clean: ## Cleans all project folders
#> @cd ./rust && rm main.exe main.pdb
#.PHONY: clean