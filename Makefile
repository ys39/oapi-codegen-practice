# 変数定義
GO       := go
OAPI_GEN := oapi-codegen
OAPI_CFG := ./doc/config.yaml
OAPI_SRC := ./doc/openapi.yaml
PKG      := ./...

# 生成されたコードの出力ディレクトリ
GENERATED_DIR := ./api
GENERATED_FILE := $(GENERATED_DIR)/interface.go

# デフォルトターゲット
all: generate build

# oapi-codegen を使ってGoコードを生成
generate:
	@echo "Generating code from OpenAPI spec..."
	$(OAPI_GEN) -config $(OAPI_CFG) $(OAPI_SRC)

# 依存関係を解決 (go mod tidy)
deps:
	@echo "Tidying up dependencies..."
	$(GO) mod tidy

# プロジェクトをビルド
build:
	@echo "Building the project..."
	$(GO) build -o bin/server main.go

# プロジェクトをテスト
test:
	@echo "Running tests..."
	$(GO) test $(PKG) -v

# コードをフォーマット
fmt:
	@echo "Formatting the code..."
	$(GO) fmt $(PKG)

# クリーンアップ
clean:
	@echo "Cleaning up..."
	rm -rf $(GENERATED_FILE) bin/

# 起動
run:
	@echo "Running the server..."
	$(GO) run main.go

.PHONY: all generate deps build test lint fmt clean run
