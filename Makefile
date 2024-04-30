# simple makefile
NAME=main
COVERAGE_MIN=50.0

## check: Check quality code in sonarqube.
check:
	@sonar-scanner -Dsonar.projectKey="$SONAR_PROJECT_KEY" -Dsonar.sources=. -Dsonar.host.url="$SONAR_HOST" -Dsonar.login="$SONAR_LOGIN"

## test: Run test and enforce go coverage
test:
	$(eval OUT = $(shell go test ./... -coverprofile cpx.out && cat cpx.out | grep -v "store/" > cp.out && rm cpx.out ))
	$(eval COVERAGE_CURRENT = $(shell go tool cover -func=cp.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' ))
	$(eval COVERAGE_PASSED = $(shell echo "$(COVERAGE_CURRENT) >= $(COVERAGE_MIN)" | bc -l ))

	@if [ $(COVERAGE_PASSED) == 0 ] ; then \
		echo "coverage is $(COVERAGE_CURRENT) below required threshold $(COVERAGE_MIN)"; \
		exit 2; \
    fi

	@echo "tests completed without failures!";
	@echo "coverage passed threshold: $(COVERAGE_CURRENT)%";

## coverage: Show go coverage
coverage: test
	@echo "coverage details:";
	@go tool cover -func=cp.out

## coverage-web: Show go coverage in web
coverage-web: test
	@go tool cover -html=cp.out

## bench: Run benchmark test
bench:
	go test -bench=.

## watch: development with air
watch:
	air -c .air.toml

## build: Build binary applications
build:
	@go generate ./...
	@echo building binary to ./dist/${NAME}
	@go build -o ./dist/${NAME} .

## deploy: Deploy binary to server using ansible-playbook
deploy:
	@echo "Command to deploy script distribute atrifacts to cloud, on-prem or kubernetes clusters "

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run with parameter options: "
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
