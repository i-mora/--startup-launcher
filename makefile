ARGS:=
ifneq (, ${export})
	ARGS := ${ARGS} --export=$(export)
endif
ifneq (, ${get})
	ARGS := ${ARGS} --get=$(get)
endif
ifneq (, ${token})
	ARGS := ${ARGS} --token=$(token)
endif
ifneq (, ${out-dir})
	ARGS := ${ARGS} --out-dir=$(out-dir)
endif

.DEFAULT:
	@go run main.go $@ ${ARGS}
