
$(eval TASK_ID := $(shell curl -s http://localhost:8080/tasks | jq -r '.[-1].id'))

id:
	@echo $(TASK_ID)

get_tasks:
	@curl -s http://localhost:8080/tasks | jq


create:
	curl -X POST -s -H "Content-Type: application/json" http://localhost:8080/tasks\
	 -d '{"title":"${title}", "workLength": 1500, "longBreakLength": 600, "shortBreakLength": 300}'
	$(eval TASK_ID = $(shell curl -s http://localhost:8080/tasks | jq -r '.[-1].id'))



