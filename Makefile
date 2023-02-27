.PHONY: worker
worker: ## Start the worker
	go run ./cmd/worker/

.PHONY: ex1
ex1: ## Start a shell with the Temporal CLI
	docker exec temporal-admin-tools tctl wf run --taskqueue workshop --execution_timeout 30 --workflow_id my-first-wf --workflow_type example01 -i 1 -i 2
.PHONY: ex2
ex2: ## Start a shell with the Temporal CLI
	docker exec temporal-admin-tools tctl wf run --taskqueue workshop --execution_timeout 30 --workflow_type example02 -i '{"A": 1, "B":2'}

.PHONY: ex4
ex4:
	docker exec temporal-admin-tools tctl wf run --taskqueue workshop --execution_timeout 30 --workflow_type example04 -i '{"Num": 3}'

.PHONY: ex5
ex5:
	docker exec temporal-admin-tools tctl wf run --taskqueue workshop --execution_timeout 300 --workflow_type example05

.PHONY: ex6
ex6:
	docker exec temporal-admin-tools tctl wf run --taskqueue workshop --execution_timeout 300 --workflow_type example06 -i '{"A": 1, "B":2'}