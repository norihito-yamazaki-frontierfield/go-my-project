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


最も一般的な形で、インターフェース要素は任意の型項T、もしくは基底型Tを指定する形式の~T、または項の和t1|t2|…|tnとして表されます
```go
package main

import (
	"fmt"
)

type MyInt int

type EnhancedInt int

// EnhancedInt のための String メソッド
func (e EnhancedInt) String() string {
	return fmt.Sprintf("EnhancedInt: %d", e)
}

type Float32 float32
type Float64 float64

// displayStringable は IntStringer 型制約を満たす任意の型 T の値を表示します
func displayStringable[T IntStringer](value T) {
	fmt.Println("IntStringer:", value.String())
}

// displayFloat は Float 型制約を満たす任意の型 T の値を表示します
func displayFloat[T Float](value T) {
	fmt.Printf("Float: %v\n", value)
}

// IntStringer は、基底型が int で String メソッドを実装する型のインターフェースです
type IntStringer interface {
	~int
	String() string
}

// Float は float32 または float64 の型制約を持つインターフェースです
type Float interface {
	~float32 | ~float64
}

func main() {
	var ei EnhancedInt = 42
	displayStringable(ei) // EnhancedInt は IntStringer を満たす

	var f32 Float32 = 3.14
	displayFloat(f32) // Float32 は Float を満たす
}

/*
IntStringer: EnhancedInt: 42
Float: 3.14

*/
```


# Goのアクセス修飾子


Goでは、シンボル（変数var、タイプtype、関数func）が小文字の記号で始まっている場合は、それは定義されているパッケージの外側のプライベートなものです。

## 例

```go
package mypackage

// PublicFunction は他のパッケージからもアクセス可能な公開関数です。
func PublicFunction() {
    // 何か処理を行う
}

// privateFunction はmypackage内でのみアクセス可能な非公開関数です。
func privateFunction() {
    // 何か処理を行う
}

```


# Go言語の `&` シンボルについて

Go言語において、`&` シンボルはアドレス演算子として使用され、特定の変数のメモリアドレスを取得するのに使います。この機能は、ポインタを通じて変数の参照を渡す際に重要です。

## アドレス演算子 `&`

アドレス演算子 `&` は、変数の前に置かれることで、その変数が格納されているメモリのアドレスを返します。これにより、変数の実際の場所を指し示すポインタが得られます。

### 使用例

```go
package main

import "fmt"

func main() {
	var a int = 58
	fmt.Println("Value of a:", a)    // aの値
	fmt.Println("Address of a:", &a) // aのメモリアドレス
}

/*
Value of a: 58
Address of a: 0xc00010e010
*/

```

この例では、整数型の変数 a が定義されており、&a によってそのメモリアドレスが表示されます。

### ポインタとの関連
& シンボルで取得したアドレスは、ポインタ変数に格納することができます。ポインタは、そのアドレスに格納されている値にアクセスしたり、変更したりするために使用されます。


```go
package main

import "fmt"

func main() {
	var a int = 100
	var p *int = &a
	fmt.Println("Value of a:", *p) // ポインタpを通じてaの値にアクセス
}
// Value of a: 100

```

このコードでは、`a` のアドレスをポインタ `p` に格納し、`*p` で `a` の値にアクセスしています。

このように、`&` シンボルはポインタと密接に関連しており、Goのポインタを理解するために不可欠な部分です。


```go
type Wallet struct {
    balance int
}

// ポインタレシーバーを使用してbalanceを返す
func (w *Wallet) Balance() int {
    return w.balance  // 自動的に (*w).balance として扱われる
}

```


## Method values


```go
package main

import "fmt"

// Person 構造体定義
type Person struct {
	Name string
	Age  int
}

// Greet メソッドは、Personのインスタンスに対して挨拶を行います。
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	// Personのインスタンスを作成
	alice := Person{Name: "Alice", Age: 30}

	// Greetメソッドのメソッド値を取得
	greetFunc := alice.Greet

	// 変数aliceのプロパティを変更
	alice.Name = "Alicia"
	alice.Age = 31

	// メソッド値を呼び出す
	greetFunc() // "Hello, my name is Alice and I am 30 years old."が出力される

	// 直接メソッドを呼び出すと、更新された情報が出力される
	alice.Greet() // "Hello, my name is Alicia and I am 31 years old."が出力される
}

/*

Hello, my name is Alice and I am 30 years old.
Hello, my name is Alicia and I am 31 years old.

*/
```


pt := &t という操作により、pt は t のアドレス、つまりポインタを保持しているため、pt と t は同じメモリ領域を指しています。したがって、pt を通じて行われる変更は t に影響を及ぼし、その逆も同様です。
```go

package main

import "fmt"

type T struct {
	a int
}

// 値レシーバを使ったメソッド
func (tv T) Mv() {
	fmt.Printf("Mv: 値レシーバの値は %d\n", tv.a)
}

// ポインタレシーバを使ったメソッド
func (tp *T) Mp() {
	fmt.Printf("Mp: ポインタレシーバの値は %d\n", tp.a)
}

func main() {
	t := T{a: 10}
	pt := &t

	t.a = 20

	// ポインタから値レシーバのメソッドを呼び出し
	pt.Mv() // 自動的にポインタが指す値 (*pt) にアクセスして実行

	// 値からポインタレシーバのメソッドを呼び出し
	t.Mp() // 自動的にtのアドレス (&t) が取られて実行
}

/*
Mv: 値レシーバの値は 20
Mp: ポインタレシーバの値は 20

*/
```

## errcheck

```bash
go install github.com/kisielk/errcheck@latest

ls /Users/n.yamazaki/go/bin/errcheck
chmod +x /Users/n.yamazaki/go/bin/errcheck

export PATH=$PATH:/Users/n.yamazaki/go/bin
#コードを含むディレクトリ内で 
errcheck .
```

- エラーの見逃し防止:
errcheck は、エラーチェックが行われていない関数呼び出しを検出します。これにより、潜在的なバグや予期しない動作を防ぐことができます。

- シンプルなインターフェース:
コマンドラインツールとして動作し、指定したディレクトリ内の全ての Go ソースコードを解析します。

## nil
ポインタはnilにすることができます

関数が何かへのポインターを返すとき、それがnilであるかどうかを確認する必要があります。そうでない場合、ランタイム例外が発生する可能性があります。コンパイラーはここでは役立ちません。

欠落している可能性のある値を説明する場合に役立ちます


## [map](https://go.dev/blog/maps)

マップの興味深い特性は、マップをポインタとして渡さなくても変更できることです。

```go

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

```


```go
// この初期化はだめ　マップがnil値になる
var m map[string]string

// OK
var dictionary = map[string]string{}
// OR
var dictionary = make(map[string]string)
```


## [センチネルエラーの問題点](https://dave.cheney.net/2016/04/07/constant-errors)

1. 変更可能な公開変数: io.EOFは公開変数であり、その値を変更することが可能です。このため、異なるパッケージやモジュールからこの値が変更されると、予期せぬ挙動やデバッグが困難な問題が発生する可能性があります。
2. シングルトンではあるが定数ではない: io.EOFはシングルトンのように扱われるが、定数としての性質（不変性や一意性）を持っていません。たとえ同じ文字列でエラーを新たに作成したとしても、io.EOFと同一ではありません。


### 改善案

```go
type Error string

func (e Error) Error() string { return string(e) }

const err = Error("EOF")
fmt.Println(err == Error("EOF")) // true
```

ただし、センチネルエラーは基本的に使用すべきではありません。


## 依存性注入

https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection

このページを簡単に要約

1. 標準出力への書き込みを抽象化する
Goのfmt.Printf関数は標準出力に直接書き込むため、テストが難しいです。代わりに、io.Writerインターフェースを利用して出力先を抽象化します。


```go
func Greet(writer io.Writer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}

```
この関数はどんなio.Writerも受け入れるため、テスト時にはバッファなどの代替出力先に置き換えることができます。

2. テストの実装

bytes.Bufferを使ってGreet関数をテストする方法です。このバッファはio.Writerを実装しており、テスト中に関数の出力をキャプチャします。


```go
func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Chris")

    got := buffer.String()
    want := "Hello, Chris"
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

```


3. リアルな使用例

Greet関数はHTTPサーバーのハンドラー内で再利用することができます。http.ResponseWriterはio.Writerを実装しているため、HTTP応答として直接使用することができます。

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet sends a personalised greeting to writer.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler says Hello, world over HTTP.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))

	if err != nil {
		fmt.Println(err)
	}
}
```


## mock

interfaceを通じててUTする時なのでMockingを行う

```go
// main.go
package main

import (
    "fmt"
    "io"
    "time"
)

// Sleeper interface defines a method for sleeping.
type Sleeper interface {
    Sleep(duration time.Duration)
}

// DefaultSleeper is a Sleeper that uses the real time.Sleep function.
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep(duration time.Duration) {
    time.Sleep(duration)
}

// Countdown counts down from a specified number to zero and then prints a final word.
func Countdown(out io.Writer, sleeper Sleeper) {
    for i := 3; i > 0; i-- {
        sleeper.Sleep(1 * time.Second)
        fmt.Fprintln(out, i)
    }
    sleeper.Sleep(1 * time.Second)
    fmt.Fprint(out, "Go!")
}

func main() {
    sleeper := &DefaultSleeper{}
    Countdown(os.Stdout, sleeper)
}

// main_test.go
package main

import (
    "bytes"
    "testing"
    "time"
)

// MockSleeper tracks the number of Sleep calls and the durations it was called with.
type MockSleeper struct {
    Calls     int
    Durations []time.Duration
}

func (m *MockSleeper) Sleep(duration time.Duration) {
    m.Calls++
    m.Durations = append(m.Durations, duration)
}

func TestCountdown(t *testing.T) {
    buffer := &bytes.Buffer{}
    mockSleeper := &MockSleeper{}

    Countdown(buffer, mockSleeper)

    got := buffer.String()
    want := "3\n2\n1\nGo!"

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }

    if mockSleeper.Calls != 4 {
        t.Errorf("not enough calls to sleeper, want 4 got %d", mockSleeper.Calls)
    }
}



```



「[でも、模試やテストのせいで生活が苦しくなってきました!](https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/mocking#demoyatesutonoseidegashikunattekimashita)」いいこと書いてありました。

- リファクタリングの定義では、コードは変更されますが、動作は同じです。理論的にリファクタリングを行うことに決めた場合は、テストを変更せずにコミットを実行できるはずです。だからテストを書くときは自問してください
  - 必要な動作や実装の詳細をテストしていますか？
  - このコードをリファクタリングする場合、テストに多くの変更を加える必要がありますか？
- Goではプライベート関数をテストできますが、プライベート関数は実装に関係しているため、避けたいと思います。
- テストが3つ以上のモックで動作している場合、それは危険信号であるように感じます（デザインを再検討する時間）
- スパイは注意して使用してください。スパイを使用すると、作成中のアルゴリズムの内部を確認できます。これは非常に便利ですが、テストコードと実装の間の結合がより緊密になることを意味します。 これらをスパイする場合は、これらの詳細に注意してください

##  [ボックルおじさんの「モックするとき」の記事](https://blog.cleancoder.com/uncle-bob/2014/05/10/WhenToMock.html)

### テストダブルの種類

1. **ダミー (Dummy)**:
   - ダミーは、操作が実際には不要である場合にテストで使用されます。ダミーは通常、メソッドが呼ばれないことが前提で、呼ばれるとエラー（例：`NullPointerException`）を引き起こすように設計されています。

2. **スタブ (Stub)**:
   - スタブは、テスト中に必要な特定の応答を返すように設定されたオブジェクトです。例えば、ユーザー認証が成功することを前提としたテストでは、常に`true`を返すスタブを使用します。

3. **スパイ (Spy)**:
   - スパイは、それが呼ばれたかどうかを記録するスタブの拡張です。これにより、システムが期待通りに特定のメソッドを呼び出しているかを検証できます。

4. **モック (Mock)**:
   - モックは、特定のメソッドがどのように、いつ、どれだけの頻度で呼ばれるかを検証するために使用されます。モックは行動の検証に重点を置き、しばしば検証ロジックを含みます。

5. **フェイク (Fake)**:
   - フェイクは、実際のビジネスロジックを模倣するより複雑なテストダブルです。例えば、特定のユーザー名でのみ認証を許可するなど、異なる入力に基づいて異なる動作をすることができます。
