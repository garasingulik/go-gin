NAME=app
VERSION=0.0.1

install:
	@go mod vendor

build:
	@echo "Cleaning output folder ..."
	@rm -rf output
	@echo "Installing dependencies to vendor folder ..."
	@go mod vendor
	@echo "Compiling $(NAME) ..."
	@go build -o output/$(NAME)

build_image:
	@docker build --build-arg APP_NAME=$(NAME) -t $(NAME):$(VERSION) .

clean:
	@rm -rf output

test:
	@go mod vendor && go test ./...

run:
	@./output/$(NAME)
