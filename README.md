# Go言語チュートリアル
- 文法の解説と簡単なAPIを設計するところまでをまとめます。
- 独学のため間違いがあったらごめんなさい
# 目次
1. [まず最初にやること(環境構築)](#anchor1)
2. [パッケージってなんやねん](#anchor2)
3. [簡単な文法解説](#anchor3)
- 1. [変数宣言](#anchor4)
  2. [基本型](#anchor5)
  3. [配列とスライス](#anchor6)
  4. [インターフェース型](#anchor7)
  5. [定数](#anchor8)
  6. [関数](#anchor9)
  7. [条件分岐(if文)](#anchor10)
  8. [条件分岐(for文)](#anchor11)
# 1. まず最初にやること(環境構築) <a id="anchor1"></a>
1. docker-composeからGoのイメージの取得
    - 簡単にいうとbuildでイメージの作成を行い,up -dでイメージをもとにコンテナを立ち上げる。
    ```
    docker compose build 
    docker cmpose up -d　
    ```
2. modファイルの初期化
    ```
    go mod init github.com/yoshimi-I/Go_RESTAPI
    ```
- といった感じでリポジトリのurlのhttps:以降をinitの後に続けたものをターミナルに打ち込む。
- そうすることでgo.modというバージョンを管理するファイルが作られる。
  - これにパッケージをインポートいていく感じですね、Reactでいうpackage.jsonみたいなやつ
# 2. パッケージってなんやねん <a id="anchor2"></a>
```Go
package main

import "fmt"

func main() {
	fmt.Println("こんにちは")
}
```
- 上にこんにちはと出力する最小限のプログラムを書いた
- このようにGoは関数型言語であり主にmainパッケージの中のmain関数のみが読まれるため(ルール)、それ以外の関数だったり、パッケージだったりはインポートして補助的な役割で使っていくことになる
- C言語に似ていますね
## 使い方
- 今回出力するにあたり、import fmt と記載されているわけだが、これは実は
```Go
package fmt
~~~~~~~~
~~~~~~~~
func Println(a ...any) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```
というGoがデフォルトで持っている,fmtパッケージの中のPrintln関数を使っていたというわけである。
- 使う場合はimportに使いたいpackageを記載して
```
package名.関数名
```
- といったように使用する。

# 3. 簡単な文法解説　<a id="anchor3"></a>
## 1. 変数宣言　<a id="anchor4"></a>
  ```Go
  関数の外では
  var i int = 100

  関数の中では
  i := 100
  ```
  - 簡単にいうと関数の中では型を省略することができる。(基本はこっち使う)
  - 少し変わってるのは型の指定を変数宣言の後に行うこと(Javaとは違う)
  - 基本的に宣言した型はどこかしらで使わないとエラーになる
    - そのため引数を２つとるような場合は１つを破棄することで対処する
    ```
    i,_ = strconv.Atoi(s)
    ```
    みたいな感じ
## 2. 基本型　<a id="anchor5"></a>
- 基本的に型に優先順位はないので、同じ型同士でないと計算ができない(C言語と違う)
- 以下に基本的な型を表示する
```
int
float
bool
string
byte
```
- ここら辺はまぁ他の静的型付け言語と一緒なので解説は省略します。
## 3. 配列とスライス　<a id="anchor6"></a>
- Go言語の配列はC言語の配列と同じく後から大きさの変更が効かない
- そこで動的な配列として用意されているのがスライスである。
- 基本的に他の言語の配列はGO言語のスライスだと思うと良い
```Go
配列
var arr1 [3]int = [3]int{1,23,456}
arr1 := [3]int{1,23,456}

配列(要素数の省略)
arr3 := [...]int{1,2,34,5}

スライス
slice0 := []int {1, 2, 3}
slice1 := make([]int, 3, 5)
```
- ちなみに変数に配列,スライスを代入した場合はそれぞれ違う挙動を示す。
- 例えば
```Go
slice := []int{1,2,3,4,5}
slice2 := slice

>> &slice[0]: 0xc000058100
>> &slice2[0]: 0xc000058100
```
- とやるとsliceとslice2のポインタ値は別のところを指すため、これは値わたしとなっている。
- 逆にスライスは参照渡し。
### 配列からスライスの作成
```Go
array := [5]int{1, 2, 3, 4, 5}
slice := array[:]

>> reflect.TypeOf(array): [5]int
>> reflect.TypeOf(slice): []int
>> array: [1 2 3 4 5]
>> slice: [1 2 3 4 5]
>> &array[0]: 0xc0000480c0
>> &slice[0]: 0xc0000480c0
```
## 4. インターフェース型　<a id="anchor7"></a>
- 初期値はnil(他の言語でいうところのnull型)
- 全ての型と互換性を持つ

## 5.定数 <a id="anchor8"></a>
- 基本的には関数の外に書く
  ```Go
  const Pi = 3.14
  const pi = 3.1

  以下のようにもかける
  const(
    A = 1
    B = 2
    c = 3
    d = 4
  )
  ```
- このように大文字だとJavaでいうところのpublic,小文字にするとパッケージの外では参照できないprivateとなる。
- またconstは型を指定しても指定しなくても作成することができる。
  ```Go
  const a = 1
  const b int = 12
  ```
  といった具合である。
## 6.関数 <a id="anchor9"></a>
- Go言語の関数は他の言語と違ってかなり色々できる。
### 1. 基本型(見慣れたやつ)
```Go
func test(a int,b int)int{
  return a + b
}
```
- これが一番基本的な型
- 変数 型の順番で引数をとり,その後に返り値の型をとる
### 2. 引数が2つあるやつ
```Go
func test(a,b int)(int,int){
  c := a + b
  d := a-b
  return c,d
}
```
- このように返り値の型に2つ代入すると返り値を2つとることができる。
- C言語とかだとポインタを使ってたけど、そんなことしなくてもいいのはでかいなぁ
### 3. 引数にもう返り値の変数を指定するやつ
```Go
func test(a,b int)(result int){
  result = a+b
  return
}
```
- ポイントとしては返り値の場所に型だけではなく、変数名を入れることでreturnした時にその変数が必ず帰るよう指定している。
- またresultを最初に宣言しているので:=とする必要がない
### 4. 無名関数
```Go
f := func(x,y int)int{
  return x + y
}
i := f(1,2)

もっと簡単に書こうとしたら
i := func(x,y int)int{
  return x + y
}(2,3)
```
### 5. 返り値に関数を持つ関数
- 無名関数を応用することで、引数に関数を持つ関数を作成することも可能。
- ポイントは型指定のところがfunc()になる点
  - func型なんだなぁくらいに思っておけば...
```Go
func test()func(){
  return func(){
    fmt.Println("test")
  }
}
```
- returnの関数の引数も,返り値の値の型も入れる必要がある
```Go
func test(a, b int) func(int) int {
	return func(c int) int {
		return a + b + c
	}
}
func main() {
	i := test(1, 2)
	fmt.Println(i(3))
}
```
- ポイントとしては引数にとった無名関数の引数,返り値の型を2つしっかりと明記しており、returnのところに変数を書いていると言うことである。
- 上の式はiに引数を取るとその引数とtestの引数を加えたもの(今回だと3)を代入しており、そのiに引数3を取ることで3＋３を出力している。
### 6. 引数に関数を持つ関数
- 例えば2回関数を呼びたい関数を出力するとする(メガ進化のガルーラみたいに)
- そうすると引数を関数として、その関数を2回呼べばいいことになる
```Go
func double_echo(f func()) {
	f()
	f()
	fmt.Println("2回呼んだよぉ")
}
func echo() {
	fmt.Println("2回呼びたいなぁ")
}
func main() {
	double_echo(echo)
}
```
### 7. クロージャー関数
- 例えば値をnumから1つずつ追加して出力したい場合を考えると
  ```Go
  func main() {
    num := 0
    a := func() int {
      num += 1
      return num
    }
    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(a())

  }
  ```
  このようになることが容易に想像できると思う。
- ではこれをnum自体も関数の引数にとらせ、かつmainの外から呼びたいとした場合はどうしたらいいだろうか？
  - この時for文などは使わずに上のようにprintlnを3回呼ぶことにする
```Go
func test(n int) func() int {
	num := n
	return func() int {
		num += 1
		return num
	}
}
func main() {
	i := test(12)
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())

}
```
- 以上が回答となる
- 解説をすると、nという値は常に更新し続けたいため、関数の中に直接書いてしまうと常に初期化が行われてしまう。
- そのため、今回は返り値を関数とし、関数をiに入れた時点でnumを作成して、それを常に更新し続けるという関数にしたわけである
  - つまりiが生成された時点でnumも生成されているため更新し続けることが可能ということである。
- これがクロージャーである。
- ちなみにこのような1ずつ増えてくものをジェネレーターという
## 7.条件分岐(if文) <a id="anchor10"></a>
- 基本的な違いは条件を()で囲む必要がないところである。
- また条件文の前に簡易文を用いることもできる。
  - 簡易文とは変数の定義と代入などのことを指す
  ```Go
  if b := 100;b == 100{
    fmt.Println("yaa")
  }
  ```
- また簡易文の変数はif文の中のみで有効である
  ```Go
  if b := 100;b > 50{
    fmt.Prinln("50より大きい")
  }
  fmt.Println(b)
  これはエラー
  ```
## 8.条件分岐(for文) <a id="anchor11"></a>
- 基本的な書き方は他の言語と同じだが、while文がないため、これはfor文を使って書く
  ```python
  pythonで書くと

    num = 0
    while True:
      if num == 10:
        break
      print(num)
      num ++
  ```
  ```Go
  Goで書くと

    func main(){
      num := 0
      for{
        if num == 10{
          break
        }
        fmt.Println(num)
        num ++
      }
    }
  ```
- このようにGoではwhileをfor文を用いて書くのである
- Cやjavaっぽいfor文の書き方は
  ```Go
  for i:=0;i<10;i++{
    fmt.Println(i)
  }
  ```
  とすればできる