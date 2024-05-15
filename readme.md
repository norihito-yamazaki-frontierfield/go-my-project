- (main)[テスト駆動開発でGO言語を学びましょう](https://andmorefine.gitbook.io/learn-go-with-tests)
- (Sub)[Welcome to a tour of Go](https://go.dev/tour/list)

## Goリンター

[golangci-lintに入門してみる](https://miyahara.hikaru.dev/posts/20201226/)
[vscode lint setting](https://github.com/golang/vscode-go/blob/master/docs/settings.md)


```bash
$ brew install golangci/tap/golangci-lint
$ golangci-lint --version


```

## 実行


```bash
go run hello.go
```

### ut実行


```bash
go test 
```


```go
// hello_test.go

package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, world"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

- `xxx_test.go`のような名前のファイルにある必要があります。
- テスト関数は`Test`という単語で始まる必要があります。
- テスト関数は1つの引数のみをとります。`t *testing.T`
- `*testing.T`型を使うには、他のファイルの`fmt`と同じように`import "testing"` が必要です。

## Go Modules: `go mod init <modulepath>` の推奨理由

#### 1. 一意性の確保

`modulepath` に `github.com/username/project` の形式を使用することで、他のプロジェクトと名前空間が重複しないようにできます。これにより、同名のパッケージが他のプロジェクトに存在しても問題が発生しません。

- 詳細：[Go Modules Reference](https://golang.org/ref/mod)&#8203;``【oaicite:2】``&#8203;

#### 2. 依存関係の明確化

モジュールパスにリポジトリのURLを含めることで、依存関係がどこから取得されるのかが明確になります。これにより、Goのツールチェーンは正しいリポジトリからソースコードを自動的に取得できます。

- 詳細：[Organizing a Go Module](https://golang.org/doc/modules/organizing)&#8203;``【oaicite:1】``&#8203;

#### 3. 自動インポートのサポート

`go get` コマンドを使用する際に、モジュールパスにリポジトリのURLが含まれていると、そのリポジトリから必要なコードを自動的にインポートすることができます。

- 詳細：[Managing Dependencies](https://golang.org/doc/modules/managing-dependencies)&#8203;``【oaicite:0】``&#8203;

### 例: モジュールの初期化

以下は、Go モジュールを `github.com/username/project` として初期化する例です：

```sh
go mod init github.com/username/project
```

このコマンドにより、go.mod ファイルが次のように作成されます：

```go
module github.com/username/project

go 1.22.2
```

さらに、プロジェクト内のパッケージは次のようにインポートできます：

```go
import "github.com/username/project/pkg/subpkg"
```

### 参考リンク

[Go Modules Reference](https://go.dev/ref/mod)
[Go Wiki: Go Modules](https://go.dev/wiki/)


# [Go Modules: go.modファイルの基本](https://go.dev/doc/modules/gomod-ref#exclude)

Go モジュールは、`go.mod` ファイルによって定義されます。このファイルには、モジュールのプロパティ（他のモジュールや Go のバージョンに対する依存関係など）が記述されています。

## プロパティ

- **モジュールパス**: 現在のモジュールのモジュールパス。これはモジュールのコードがダウンロードされる場所であり、モジュールのバージョン番号と組み合わせて一意の識別子となります。また、モジュール内の全てのパッケージのパッケージパスのプレフィックスでもあります。
- **Go のバージョン**: 現在のモジュールに必要な最低限の Go のバージョン。
- **依存関係**: 現在のモジュールに必要な他のモジュールの最低バージョンのリスト。
- **置き換えと除外**: 必要なモジュールを別のモジュールのバージョンまたはローカルディレクトリに置き換える、または特定のバージョンを除外する指示（オプション）。

## go.mod ファイルの生成

`go mod init` コマンドを実行すると、Go は `go.mod` ファイルを生成します。以下の例では、モジュールのモジュールパスを `example/mymodule` に設定して `go.mod` ファイルを作成します：

```sh
$ go mod init example/mymodule
```

## 依存関係の管理

依存関係を管理するために Go ツールを使用する際、Go ツールは `go.mod` ファイルに変更を加えます。詳細については、[依存関係の管理](https://example.com)を参照してください。

### 参照

- [依存関係の管理](https://example.com)
- `go.mod` ファイルに関連する詳細および制約については、[Go モジュールリファレンス](https://example.com)を参照してください。

## 例: go.mod ファイル

以下は、`go.mod` ファイルの例です。それぞれの指示についてコメントを付けて解説します。

```go

// リポジトリパスを使用することはベストプラクティスです。
module example.com/mymodule // モジュールパス: このモジュールの一意の識別子であり、パッケージパスのプレフィックスにもなります。(バージョン0または1のモジュールの宣言:)
// module example.com/mymodule/v2 // バージョン2のモジュールのモジュールパス

go 1.14 // 必須のGoのバージョン: このモジュールが動作するために必要な最低限のGoのバージョン。

require (
    example.com/othermodule v1.2.3 // 依存モジュール: このモジュールが依存する他のモジュールとそのバージョン。
    example.com/thismodule v1.2.3  // 依存モジュール: 別の依存するモジュールとそのバージョン。
    example.com/thatmodule v1.2.3  // 依存モジュール: さらに別の依存するモジュールとそのバージョン。
)

replace example.com/thatmodule => ../thatmodule // モジュールの置き換え: 特定のモジュールをローカルディレクトリや別のバージョンに置き換え。

exclude example.com/thismodule v1.3.0 // バージョンの除外: 特定のバージョンのモジュールを除外。
```