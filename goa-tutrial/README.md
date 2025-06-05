# Hello Goa - Goaフレームワーク チュートリアル

このプロジェクトは、Goaフレームワークを使用したAPIサーバーの作成方法を学習するためのチュートリアルです。

## Goaとは

Goaは、Go言語でマイクロサービスとAPIを構築するためのフレームワークです。設計駆動開発（Design-Driven Development）のアプローチを採用し、まずAPIの設計を定義してからコードを生成します。

## 環境要件

- Go 1.23.3 以上
- Goa v3

## プロジェクト構造

```
hello-goa/
├── design/          # API設計ファイル
│   └── design.go    # API設計定義
├── gen/             # Goaによって生成されるコード
│   ├── hello/       # サービス関連の生成コード
│   └── http/        # HTTP関連の生成コード
├── cmd/             # メインアプリケーション
│   ├── hello/       # HTTPサーバー
│   └── hello-cli/   # CLIクライアント
├── hello.go         # サービス実装
├── go.mod
├── go.sum
└── README.md
```

## セットアップ

### 1. Goaのインストール

```bash
go install goa.design/goa/v3/cmd/goa@v3
```

### 2. 依存関係のインストール

```bash
go mod tidy
```

## チュートリアル手順

### Step 1: API設計の作成

`design/design.go` ファイルでAPIの設計を定義しています。

### Step 2: コード生成

```bash
goa gen hello/design
```

このコマンドで、`gen/` ディレクトリに以下のファイルが生成されます：
- サービスインターフェース
- HTTPハンドラー
- クライアントコード
- OpenAPI仕様書

### Step 3: サービス実装

`hello.go` でサービスロジックを実装しています。

### Step 4: アプリケーションの実行

```bash
# コード生成（必要に応じて）
goa gen hello/design

# HTTPサーバーの起動
go run ./cmd/hello

# またはCLIクライアントの実行
go run ./cmd/hello-cli
```

## APIテスト

### cURLでのテスト

```bash
# 基本的なリクエスト
curl http://localhost:8080/hello/Alice

# 日本語での挨拶
curl http://localhost:8080/hello/太郎
```

### 期待される応答

```json
"こんにちは、Aliceさん!"
"こんにちは、太郎さん!"
```

### 文字化け対策

PowerShellで日本語が文字化けする場合は、以下のいずれかを試してください：

1. **文字エンコーディングを変更**
   ```powershell
   chcp 65001
   curl http://localhost:8080/hello/Alice
   ```

2. **ブラウザでアクセス**
   ```
   http://localhost:8080/hello/Alice
   ```

3. **Invoke-RestMethodを使用**
   ```powershell
   Invoke-RestMethod -Uri http://localhost:8080/hello/Alice
   ```

## 生成されるファイル

Goaは以下のファイルを自動生成します：

- `gen/hello/service.go` - サービスインターフェース
- `gen/hello/endpoints.go` - エンドポイント定義
- `gen/http/hello/server/server.go` - HTTPサーバー実装
- `gen/http/hello/client/client.go` - HTTPクライアント実装
- `gen/http/openapi.json` - OpenAPI仕様書（JSON形式）
- `gen/http/openapi.yaml` - OpenAPI仕様書（YAML形式）

## 利用可能なコマンド

### HTTPサーバー
```bash
# デフォルト設定で起動
go run ./cmd/hello

# ポートを指定して起動
go run ./cmd/hello --http-port 8090

# ヘルプを表示
go run ./cmd/hello --help
```

### CLIクライアント
```bash
# CLIクライアントでAPIを呼び出し
go run ./cmd/hello-cli hello say-hello --name="テスト"

# ヘルプを表示
go run ./cmd/hello-cli --help
```

## OpenAPI仕様書

生成されたOpenAPI仕様書は以下の場所にあります：
- JSON形式: `gen/http/openapi.json`
- YAML形式: `gen/http/openapi.yaml`
- OpenAPI 3.0形式: `gen/http/openapi3.json`, `gen/http/openapi3.yaml`

これらの仕様書は、SwaggerやPostmanなどのツールでAPIドキュメントとして利用できます。

## 参考リンク

- [Goa公式ドキュメント](https://goa.design/)
- [Goa GitHubリポジトリ](https://github.com/goadesign/goa)
- [Goa Examples](https://github.com/goadesign/examples)

## トラブルシューティング

### よくある問題

1. **ポートが使用中**
   ```bash
   # 別のポートを指定
   go run ./cmd/hello --http-port 8090
   ```

2. **依存関係の問題**
   ```bash
   go mod tidy
   go clean -modcache
   ```

3. **コード生成エラー**
   ```bash
   # design.goファイルの構文を確認
   goa gen hello/design
   ```

4. **文字化け**
   - PowerShellの場合: `chcp 65001` で文字エンコーディングを変更
   - ブラウザまたは他のHTTPクライアントツールを使用

## 次のステップ

このチュートリアルを完了したら、以下を試してみてください：

1. より複雑なAPIエンドポイントの追加
2. リクエスト/レスポンスの構造体定義
3. バリデーションの追加
4. ミドルウェアの実装
5. エラーハンドリングの改善
6. テストの作成
7. データベース連携