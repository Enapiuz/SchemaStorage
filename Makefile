all:
	@echo 123

fmt:
	govendor fmt +local

test:
	govendor test +local