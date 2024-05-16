- (main)[テスト駆動開発でGO言語を学びましょう](https://andmorefine.gitbook.io/learn-go-with-tests)
- (Sub)[Welcome to a tour of Go](https://go.dev/tour/list)

## Goリンター

[golangci-lintに入門してみる](https://miyahara.hikaru.dev/posts/20201226/)
[vscode lint setting](https://github.com/golang/vscode-go/blob/master/docs/settings.md)


```bash
$ brew install golangci/tap/golangci-lint
$ golangci-lint --version


```

## 開発ツール
vscode

名前: Go
ID: golang.go
説明: Rich Go language support for Visual Studio Code
バージョン: 0.41.4
パブリッシャー: Go Team at Google
VS Marketplace リンク: https://marketplace.visualstudio.com/items?itemName=golang.Go

.vscode/settings.json

```json
{
    "go.lintTool": "golangci-lint",
    "go.lintFlags": ["--fast"],
    "go.lintOnSave": "package",
    "go.useLanguageServer": true,
    "go.formatTool": "gofmt",
    "go.testFlags": ["-v"]
}

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

### 例関数の目的と基本的な使い方

- **例関数の定義**: Goでの例関数は、`_test.go`ファイルに記述される特別なテスト関数です。これらの関数は、`Example`という接頭辞を持ち、通常のテスト関数と同様に、パッケージのテストスイートの一部としてコンパイルされます。

- **例関数の書き方**: 例関数内では、パッケージが提供する関数を呼び出し、その動作を示します。出力はコメントとして`// Output: `の後に記述し、テスト時にこの出力が期待値と一致するか検証されます。出力が正しいと、テストは成功と見なされます。

### 具体的な例

```go
func ExampleAdd() {
    sum := Add(1, 5) // Add 関数を呼び出し
    fmt.Println(sum) // 結果を出力
    // Output: 6       // 期待される出力値
}
```

テストの実行方法
コマンド $ go test -v を使用してテストを実行すると、テストスイート内の全てのテストと例関数が実行されます。例関数も通常のテストと同様に、結果が出力コメントと一致するかどうかで評価されます。

注意点
例関数から// Output: コメントを削除すると、その関数は単にコンパイルされるだけで実際には実行されません。出力を検証するコメントがないため、テストとしての機能は果たされなくなります。

## 変数宣言

```go
package main

import "fmt"

func main() {
    // 1つの変数を宣言し、初期化します
    var a = "initial"
    // "initial"が表示されます
    fmt.Println(a)

    // 複数の変数を一度に宣言し、初期化します
    var b, c int = 1, 2
    // 1と2が表示されます
    fmt.Println(b, c)

    // Goは初期化された変数の型を推測します
    var d = true
    // trueが表示されます
    fmt.Println(d)

    // 初期化されない変数はゼロ値を持ちます。intのゼロ値は0です
    var e int
    // 0が表示されます
    fmt.Println(e)

    // 簡略化された変数宣言と初期化。関数内でのみ使用できます
    f := "apple"
    // "apple"が表示されます
    fmt.Println(f)
}

```

## 繰り返しはfor 飲み

Goで繰り返し作業を行うには、 forが必要です。 Goには while、do、 untilキーワードはなく、forのみ使用できます。


```go
package main

import "fmt"

func main() {
    // 基本的な型。単一の条件付きループ
    i := 1
    for i <= 3 {
        // 1, 2, 3が表示されます
        fmt.Println(i)
        i = i + 1
    }

    // クラシックな初期化/条件/後処理のforループ
    for j := 0; j < 3; j++ {
        // 0, 1, 2が表示されます
        fmt.Println(j)
    }

    // N回繰り返しの別の方法。これは誤りです。正しくは以下のようにスライスや配列を使います。
    nums := []int{0, 1, 2}
    for i := range nums {
        // "range 0", "range 1", "range 2"が表示されます
        fmt.Println("range", i)
    }

    // 条件なしのforループ。breakやreturnがないと無限ループします
    for {
        // "loop"が表示されます
        fmt.Println("loop")
        break
    }

    // 次のループの反復に進むcontinue
    for n := 0; n < 6; n++ {
        if n%2 == 0 {
            continue
        }
        // 1, 3, 5が表示されます
        fmt.Println(n)
    }
}

/*
rangeは配列の反復処理を行うことができる。呼び出されるたびにインデックスと値の2つの値を返します。ここではインデックスの値を無視して _ 空（スペース）の識別子を使用しています。
[5]int を期待する関数に [4]int を渡そうとしてもコンパイルできません。
*/
func Sum(numbers [5]int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}
```

## Slices

```go
package main

import "fmt"

func main() {
	// 配列の定義：固定長
	var myArray [3]int = [3]int{1, 2, 3}
	fmt.Println("配列:", myArray)

	// 配列の要素を変更（可能）
	myArray[0] = 10
	fmt.Println("変更後の配列:", myArray)

	// 配列の長さは固定
	fmt.Println("配列の長さ:", len(myArray))

	// スライスの定義：動的長
	var mySlice []int = []int{1, 2, 3}
	fmt.Println("スライス:", mySlice)

	// スライスに要素を追加（動的に長さを増やせる）
	mySlice = append(mySlice, 4)
	fmt.Println("要素を追加したスライス:", mySlice)

	// スライスの長さと容量は動的に変更可能
	fmt.Println("スライスの長さ:", len(mySlice))
	fmt.Println("スライスの容量:", cap(mySlice))
}


/*
配列: [1 2 3]
変更後の配列: [10 2 3]
配列の長さ: 3
スライス: [1 2 3]
要素を追加したスライス: [1 2 3 4]
スライスの長さ: 4
スライスの容量: 6
*/
```


# Goのベンチマーク機能

Go言語におけるベンチマークの実装は、言語の強力な特徴の一つであり、テストの記述に非常に似ています。以下は、ベンチマーク関数の基本的な構造を示す例です。

```go
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```

このコード例では、testing.Bを利用しています。これにより、ベンチマーク関数内でb.Nというカウンタにアクセスでき、このカウンタはテストが実行される回数を制御します。ベンチマークが実行されると、この関数はb.N回実行され、その実行にかかる時間が計測されます。

実際のコード実行回数（b.Nの値）は、ベンチマークフレームワークによって自動的に調整され、信頼性のある結果を得るために最適な回数が選ばれます。

ベンチマークを実行するには、次のコマンドを使用します：
```bash
go test -bench=.


# windows
go test -bench="."
```

```text
goos: darwin
goarch: amd64
pkg: github.com/quii/learn-go-with-tests/for/v4
10000000           136 ns/op
PASS
```

この出力では、「136 ns/op」という結果が得られており、これは関数が平均で136ナノ秒かかることを意味します。この結果は、関数が10000000回実行された後のものです。



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

```text
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


## Go言語のメソッド宣言

### 基本構文

Go言語におけるメソッドは、特定の型に関連付けられた関数です。メソッドを宣言するための基本的な構文は以下の通りです。

```go
MethodDecl = "func" Receiver MethodName Signature [ FunctionBody ] .
Receiver   = Parameters .
```
※「.（ドット）」は、その文法ルールの終端を表しています。
Receiverは、メソッドが操作を行う対象の型を指定します。

### レシーバーの定義
レシーバーはメソッドの実行対象となる型で、メソッド名の前にパラメータセクションとして指定されます。以下はレシーバーを使用したメソッドの例です。


```go
package main

import (
	"fmt"
)

type Point struct {
	x float64
	y float64
}

// 値渡しのメソッド（レシーバーが値）
func (p Point) Move(dx, dy float64) Point {
	p.x += dx
	p.y += dy
	return p // 新しい位置のPointを返す
}

// 参照渡しのメソッド（レシーバーがポインタ）
func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}

func main() {
	p := Point{2.0, 3.0}

	// Moveメソッドを使ってポイントを移動（値渡し）
	movedPoint := p.Move(3.0, 4.0)
	fmt.Printf("Original Point after Move: {%.1f %.1f}\n", p.x, p.y)              // 元のポイントは変わらない
	fmt.Printf("New Point after Move: {%.1f %.1f}\n", movedPoint.x, movedPoint.y) // 新しい位置のポイント

	// Scaleメソッドを使ってポイントをスケール（参照渡し）
	p.Scale(2.0)
	fmt.Printf("Point after Scale: {%.1f %.1f}\n", p.x, p.y) // 元のポイントがスケールされる
}

/*
Original Point after Move: {2.0 3.0}
New Point after Move: {5.0 7.0}
Point after Scale: {4.0 6.0}
*/

```

上記の例では、Point型にLengthとScaleという二つのメソッドが定義されており、それぞれが*Point型のレシーバーを持っています。

### ジェネリック型を持つレシーバー
Go言語のジェネリックを用いることで、より汎用的なメソッドを実装することが可能です。以下はジェネリック構造体Pairとそのメソッドの例です。

```go
package main

import (
	"fmt"
)

// Pairはジェネリックな構造体で、2つの異なる型Type1とType2の値を格納します。
type Pair[Type1, Type2 any] struct {
	a Type1
	b Type2
}

// Swapメソッドは、Pairの要素の位置を交換し、新しい型のPairを返します。
func (p Pair[Type1, Type2]) Swap() Pair[Type2, Type1] {
	return Pair[Type2, Type1]{a: p.b, b: p.a}
}

// Firstメソッドは、Pairの最初の要素（Type1）を返します。
func (p Pair[Type1, _]) First() Type1 {
	return p.a
}

func main() {
	// intとstringのペアを作成
	pair := Pair[int, string]{a: 1, b: "apple"}

	// Swapメソッドを使って、要素の順序を交換
	swappedPair := pair.Swap()
	fmt.Printf("Original Pair: {%v, %v}\n", pair.a, pair.b)
	fmt.Printf("Swapped Pair: {%v, %v}\n", swappedPair.a, swappedPair.b)

	// Firstメソッドを使って、元のペアの最初の要素を取得
	firstElement := pair.First()
	fmt.Printf("First element of the original pair: %v\n", firstElement)
}

/*
Original Pair: {1, apple}
Swapped Pair: {apple, 1}
First element of the original pair: 1
*/

```

Pair型は二つの異なる型AとBを持ち、Swapメソッドはこれらの型を逆転させた新しいPairを返します。また、FirstメソッドはPairの最初の要素を返します。


## [Interface types](https://go.dev/ref/spec#Interface_types)

インターフェイスタイプの初期化されていない変数の値はnilです。Goでは、インターフェースの解決は暗黙的です。


**インターフェース要素**
インターフェースタイプは、interfaceキーワードを使用して定義され、その本体は中括弧 {} で囲まれます。インターフェースの本体内には、メソッド要素と型要素のいずれか、または両方が含まれることがあります。

**メソッド要素**
メソッド要素は、インターフェースが要求するメソッドのシグネチャを定義します。このシグネチャには、メソッド名とパラメータ、戻り値の型が含まれます。例えば、Read([]byte) (int, error) は Read メソッドが []byte 型の引数を取り、int と error を返すことを要求するメソッド要素です。

**タイプ要素**
タイプ要素は、型項のユニオン（合併）を定義します。型項は、具体的な型（例：int、string）またはその型の基底型を指定します。基底型を指定するには、~ 記号を型の前に置きます（例：~int）。これは、int型を基底型とするすべての型を表します。タイプ要素は、インターフェースがどのような型を受け入れるかを広げるために使用されます。

```go
interface {
    Read([]byte) (int, error)   // メソッド要素
    ~int | string               // タイプ要素（intの基底型またはstring型のどちらか）
}

```


あるインターフェース（この例ではFileインターフェース）が特定のメソッド（Read, Write, Close）を要求する場合、これらのメソッドを全て実装している任意の型はそのインターフェースを実装していると見なされます。つまり、S1とS2という二つの型が同じメソッドセットを持っている場合、FileインターフェースはS1とS2の両方によって実装されることになります。
```go
package main

import (
	"fmt"
)

// FileインターフェースはRead、Write、Closeメソッドを要求します。
type File interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	Close() error
}

// S1型の定義
type S1 struct{}

func (s S1) Read(p []byte) (int, error) {
	// Readの実装
	return len(p), nil
}

func (s S1) Write(p []byte) (int, error) {
	// Writeの実装
	return len(p), nil
}

func (s S1) Close() error {
	// Closeの実装
	return nil
}

// S2型の定義
type S2 struct{}

func (s S2) Read(p []byte) (int, error) {
	// Readの実装
	return len(p), nil
}

func (s S2) Write(p []byte) (int, error) {
	// Writeの実装
	return len(p), nil
}

func (s S2) Close() error {
	// Closeの実装
	return nil
}

// 汎用的な関数でFileインターフェースを使用
func useFile(f File) {
	data := make([]byte, 100)
	f.Read(data)
	f.Write(data)
	f.Close()
	fmt.Println("Used a File interface")
}

func main() {
	var f1 File = S1{}
	var f2 File = S2{}

	useFile(f1)
	useFile(f2)
}
/*
Used a File interface
Used a File interface
*/
```

**インターフェースの埋め込み**
インターフェース内に他のインターフェースを埋め込むことができます。
ただし、埋め込みインターフェース間で同名のメソッドが存在する場合は、そのシグネチャが完全に一致する必要があります。もしシグネチャが異なる場合、それはコンパイルエラーを引き起こします。

```go
package main

import (
    "fmt"
    "io"
    "strings"
)

// ReaderインターフェースはReadメソッドを要求します。
type Reader interface {
    Read(p []byte) (n int, err error)
}

// WriterインターフェースはWriteメソッドを要求します。
type Writer interface {
    Write(p []byte) (n int, err error)
}

// ReadWriterはReaderとWriterのメソッドを含みます。
type ReadWriter interface {
    Reader
    Writer
}

// myBufferはReadWriterインターフェースを実装します。
type myBuffer struct {
    buf string
}

// Readメソッドの実装
func (b *myBuffer) Read(p []byte) (int, error) {
    n := copy(p, b.buf)
    b.buf = b.buf[n:]
    return n, nil
}

// Writeメソッドの実装
func (b *myBuffer) Write(p []byte) (int, error) {
    b.buf += string(p)
    return len(p), nil
}

func main() {
    var rw ReadWriter = &myBuffer{}

    // 文字列を書き込む
    rw.Write([]byte("Hello, world!"))
    // バッファから読み取る
    buf := make([]byte, 6)
    rw.Read(buf)
    fmt.Println(string(buf))  // "Hello," を出力

    // 残りを読み取る
    buf = make([]byte, 20)
    n, _ := rw.Read(buf)
    fmt.Println(string(buf[:n]))  // " world!" を出力
}
/*
Hello,
 world!
*/
```