### oapi-codegen を用いて OpenAPI から Gin の API サーバのコードを自動生成する

1. プロジェクトを clone

    ```bash
    git clone git@github.com:ys39/oapi-codegen-practice.git go-oapi-codegen
    cd go-oapi-codegen
    ```

2. OpenAPI から Gin の API サーバのコードを自動生成する

    ```bash
    make generate
    ```

3. 起動

    ```
    make run
    ```

-   その他のコマンドは`Makefile`を参照
