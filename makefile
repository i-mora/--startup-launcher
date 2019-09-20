# make slack get=users export=true token=xoxp-739841696129-746226875648-739844863985-c1b515f143e29c38175b1bdb0c64a672

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
