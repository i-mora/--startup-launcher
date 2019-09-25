ARGS:=
ifneq (, ${export})
	ARGS := ${ARGS} --export=$(export)
endif
ifneq (, ${out-dir})
	ARGS := ${ARGS} --out-dir=$(out-dir)
endif

ifneq (, ${get})
	ARGS := ${ARGS} --get=$(get)
endif
ifneq (, ${post})
	ARGS := ${ARGS} --post=$(post)
endif

ifneq (, ${user})
	ARGS := ${ARGS} --user=$(user)
endif
ifneq (, ${token})
	ARGS := ${ARGS} --token=$(token)
endif
ifneq (, ${url})
	ARGS := ${ARGS} --url=$(url)
endif

.DEFAULT:
	@go run main.go $@ ${ARGS}

%.txt:
	@echo "$*"
	@[ -f $* ] || echo "I'm a file called $*.txt" > $*.txt
	cat $*.txt
