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
  9. [条件分岐(switch文)](#anchor12)
  10. [その他の色々な書き方(defer)](#anchor13)
  11. [init関数](#anchor14)
  12. [参照型](#anchor15)
  13. [チャネル](#anchor16)
  14. [チャネルとゴルーチン](#anchor17)
  15. [ポインタ](#anchor18)
  16. [構造体](#anchor19)
  17. [インターフェース](#anchor20)
  18. [スコープ](#anchor21)
4. [テスト](#anchor22)
5. [標準パッケージ](#anchor23)
6. [DB操作](#anchor24)

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
### (応用)ラベル付きのfor
- まずは既存のfor文
```GO
func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break
			}
			fmt.Println(i, j)
		}
	}
}

>>>
0 0
0 1
0 2
1 0
1 1
1 2
2 0
2 1
2 2
3 0
3 1
3 2

```
- これはjが3の時のみ抜ける処理のため前半は1から3までループ,後半は0から2までループすることになる。
- ではjの値が3となった瞬間に全てのループを抜けるにはどうしたらいいだろうか？
  - pythonでいうところのexit()である
- それが以下になる
```Go
func main() {
Loop:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == 3 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("ループ終了")
```
- Loopと付けることで、breakの後にLoop直下の場所のfor文終了までスキップすることができる。
## 9.条件分岐(switch文) <a id="anchor12"></a>
- 基本的には他のものと一緒
```Go
n := 1
switch n{
case 1,2:
  fmt.Println("1か2です")
case 3:
  fmt.Println("3です")
defer:
  fmt.Println("1,2,3以外です")
}



ちなみに以上を簡単に書くと
switch n:=1;n{
case 1,2:
  fmt.Println("1か2です")
case 3:
  fmt.Println("3です")
defer:
  fmt.Println("1,2,3以外です")
}


また
switch{
  case n = 1:
    fmt.Println("n=1です")
}
と書くこともできる
```
### switch応用
- 型による条件分岐をすることも可能である。
  ```go
  switch x.(tipe){
    case int:
      fmt.Println
  }
  ```
  みたいな感じ
## 10.その他の色々な書き方)(defer) <a id="anchor13"></a>
- 簡単にいうとdeferをつけると処理が終わったとに(つまり最後に実行される)
- 2つ以上あるときは,スタックのようになり,最初にdeferが付いてるものが一番最後
```go
defer fmt.Println("最後")
defer fmt.Println("最後2")
fmt.Println("やあ")

>>>
やあ
最後2
最後
```
- 以上のような出力となる
## 11. init関数<a id="anchor14"></a>
- main関数が最初に呼ばれるわけだが、それよりも前に呼ぶことができる。
```Go
func init() {
	fmt.Println("やあ")
}

func main() {
	fmt.Println("テスト")

}
>>>
やあ
テスト
```
## 12. 参照型<a id="anchor15"></a>
1. スライスの復習
  - 初期化
    ```go
      sl := []int{1, 2, 3}
      fmt.Println(sl)
    ```
  - make関数
    ```go
      //make(作りたいもの,要素数,容量)
      s1 := make([]int,5,10) 

    ```
  - 追加
    ```go
      sl2 := make([]int, 5)
      sl2 = append(sl2, 3)
      fmt.Println(sl2)
    ```
  - コピー(同じ参照先を使わない配列)
    ```go
    // copy(コピー先,コピー元)
    s1 := []int{1, 3, 4}
    s2 := make([]int, 10, 10)
    fmt.Println(s2)
    copy(s2, s1)
    fmt.Println(s2)

    >>>
    [0 0 0 0 0 0 0 0 0 0]
    [1 3 4 0 0 0 0 0 0 0]

    ```
2. 辞書型
    ```go
    s1 := map[string]int{"A": 1, "B": 3}
    s1["C"] = 1
    fmt.Println(s1)
    fmt.Println(s1["A"])
    ```
  - エラーハンドリング
    ```go
    s,err = s1["D"]
    // こうすることで存在しない場合はerrにfalseが入る
    ```
  - 削除
    ```go
    delete(s1,"A")
    // こうすることで要素を削除できる
    ```
  - for文でループ
    ```go
    s1 := map[string]int{"A": 1, "B": 3}
    s1["C"] = 1
    for key, value := range s1 {
      fmt.Println(key, value)
    }
    ```
  ## 13. チャネル<a id="anchor16"></a>
  - チャネルとは、バッファ(一時的に記憶する場所)を用いて送受信をするときの中間役を担う
    - 島と島を行き来する船みたいなもん(乗客数が限られている)
  - 基本的にはチャネル作成時にどれだけの領域を確保するかを指定することができる。
    ```go
    ch := make(chan int, 2)
    // これは容量が２である
    ```
  - そして以下のようにすることでチャネル(船)に値を突っ込める
    ```
      ch <- 1
      ch <- 2 
    ```
  - また取り出すときは以下のようになる
    - また取り出す順番はキューである(最初に乗った乗客から降りるイメージ)
    ```go
    i := <- ch 
    // とすることで取り出せる。
    ```
  - 具体例
    ```go
    s1 := make(chan int, 6)
    s1 <- 1
    s1 <- 2
    num1 := <-s1
    num2 := <-s1
    fmt.Println(num1, num2)
    s1 <- 3
    fmt.Println(<-s1)

    >>>
    1 2
    3
    ```
  ### チャネルを閉じる
  - 値の入り口を閉じる、値は取り出せる（それ以上船に客を入れないようにするイメージ）
  ```go
    s1 := make(chan int, 6)
    close(s1)

    // また閉じているかの確認は第二引数から確認可能(チャネルのバッファが空でかつ、閉じていたらfalse)
    i,ok := <-s1
  ```
  ### チャネルとfor文
  - 気を付ける場所はバッファがいっぱいになったら必ずチャネルを閉めること
  ```go
  ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 3
	ch1 <- 5
	close(ch1) //for文で回す時は必ずここがいる
	for i := range ch1 {
		fmt.Println(i)
	}
  ```
  ## 14. チャネルとゴルーチン<a id="anchor17"></a>
  1. ゴルーチンとは
  - 簡単に言え並列処理を行うための仕組み
    ```go
    func hello() {
      for {
        fmt.Println("やあ")
        time.Sleep(100 * time.Millisecond)
      }
    }
    func hello2() {
      for {
        fmt.Println("やあ2")
        time.Sleep(100 * time.Millisecond)
      }
    }
    func main() {
      hello()
      hello2()
    }
    ```
  - 例えば以上の例だとhello()が無限ループのためhello2()は呼ばれないわけだが、これを呼ばれるようにするにはgoをhello()の前につけることで並行に動く関数だと明示する
    - ポイントは関数の前にgoとつけなくてはいけない(関数出ないといけない)
    ```go
    go hello()
    hello2()
    
    ```
  2. チャネルとゴルーチン
  - 並列に行うとき、2つの走っている関数間で値の受け渡しを行い続けるときにチャネルを用いる
  ## 15. ポインタ<a id="anchor18"></a>
  - ポインタとは変数や関数,配列のメモリの場所を指す
  - C言語と同じだがポインタ型が存在する
    - 基本的にはC言語とおんなじ
    ```go
    i := 13
    p := &i
    fmt.Println(p, *p)

    >> 
    0x1400001a0d8 13
    ```
    - &をつけるとメモリのアドレス値
    - ＊をつけるとポインタの参照先の値,＊がないとそのままアドレスを指す
    ## 16. 構造体<a id="anchor19"></a>
    - 簡単に言えばオブジェクト指向のclassの役割
      - 構造体を使うあたりC言語へのリスペクトを感じる
      - いかにpythonとGoでの比較を行う
      ```python
      # pyhon
      class Test:
        def __init__(self,age,name):
          self.age = age
          self.name = name
      test = Test()
      ```
      ```go
      //Go
      type User struct {
	    Age  int
	    Name string
      }
      user := User{22,"Yoshimi"}
      ```
    ### 初期化のあれこれ
    - 構造体の初期化には主に3つがある
    ```go
    var user1 User
	  user2 := User{}
    // 以上の2つは同じ意味

    user3 := new(User)
    user4 := &User{}
    // これはUserのポインタ値を返す
    ```
    ### じゃあどこで構造体のポインタを使うのか？
    - これはメソッドの実装で用いる
    - 簡単にいうと構造体のポインタを引数に取る関数を実装することでその構造体を更新することができるため、そこでメンバ変数だったりメソッドを実装することができる。
    ### メソッドの実装
    - 構造体に関数を実装する
    ```go
    func (u *User) SeyName() {
      fmt.Println(u.Name)
    }
    ```
    - 基本的には関数名の前にメソッドを実装したい構造体とそのポインタ値を実装する
    ### メソッドの埋め込み
    - 構造体の中に構造体を埋め込むことができる
    ```Go
    type T struct {
      User
    }
    user1 := T{User: User{23, "YOSHIMI"}}
    fmt.Println(user1.Name)
    ```
    - このようにいきなりNameを使うことも可能
    ### コンストラクタ
    - GOにはコンストラクタ(初期化メソッド)が存在しないため自分で作る必要がある
    - 初期化するときに自動で作られるもののため、返り値を構造体のポインタ値にする
    ```go
      func NewUser(age int, name string) *User {
    return &User{Age: age, Name: name}
  }
    ```
    - ちなみにIntelliJだと自動で生成してくれる(javaとおんなじ)
  ## 17. インターフェース<a id="anchor20"></a>
  - 簡単に言えば構造体ごとに微妙に異なるが共通している部分があるメソッドをインターフェースとして保持しておく(具体的な処理は書かずメソッド名だけを保持する)
  ```go
  type Person struct {
	Age  int
	Name string
}
type Animal struct {
	Age  int
	Kind string
}

func (p *Person) selfIntroduction() {
	fmt.Println("私は"+p.Name+"です")
}
func (a *Animal) selfIntroduction()  {
	fmt.Println("私は"+a.Kind+"です")
}
  ```
  - このように同じメソッド名だけど微妙に違う場合インターフェースを用いる
  - インターフェースは何かは具体的には書いていないがこう言うメソッドを後に呼ぶと言うことを宣言するときに使う
### 大まかな手順
1. まずはメソッドをまとめたインターフェースを作成
2. 次に作成したインターフェース型に作成した構造体を作成する
    - このとき構造体の中に存在するメソッド名をインターフェースで実装している必要がある
    - ちなみにメソッドにポインタを指定する場合は&をつける
  ```go
  type Person struct {
	Age  int
	Name string
}
type Animal struct {
	Age  int
	Kind string
}

func (p *Person) selfIntroduction() {
	fmt.Println("私は" + p.Name + "です")
}
func (a *Animal) selfIntroduction() {
	fmt.Println("私は" + a.Kind + "です")
}

type Introduction interface {
	selfIntroduction()
}

func main() {
	var intro Introduction = &Person{22, "YOSHIMI"}
	intro.selfIntroduction()

}
  ```
  ## 18. スコープ<a id="anchor21"></a>
  - Goは複数のパッケージを組み合わせて使うが他のパッケージ内から参照するのか,パッケージ内のみのスコープにするのかなど厳密に管理する必要がある
  ### パッケージ間のスコープ
  - 今回はmainと同じディレクトリのpkgの下にtest.goというファイルを作成し、そこにパッケージを作成した。
  - 大文字であれば外部から使えるし、小文字であれば使えない
  ```go
    package test

    const (
      Max = 100
      min = 1
    )
  ```
  - このとき,Maxは呼び出せるがminは呼び出せない
    - minはプライベートのため使いたかったらgetterのような関数を実装する
    ```go
    func Return_min() int {
      return min
    }
    ```
# 4. テスト<a id="anchor22"></a>
- GO言語にはtest機能というものが備わっており同じディレクトリにtestファイルを作る
  - 例えばmain.goのテストがしたければmain_test.goを作る
- 実行したい場合はターミナルでgo test とする
  ```go
  package main

  import "testing"

  //ここにデバッグ制御の文章を書く

  var Debug bool = false

  func TestOldTest(t *testing.T) {
    i := 20
    if Debug {
      t.Skip("スキップ")
    }
    v := OldTest(i)
    if !v {
      t.Errorf("%v >= %v", i, 20)
    }

  }
  ```
# 5. 標準パッケージ<a id="anchor22"></a>
## OS
- 終了
 ```go
 os.exit(1) >>>抜ける(pythonのexitと同じ)
 ```
 - ファイル処理
 ```go

 f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
```
## rand
- 疑似乱数の生成

# 6. DB操作<a id="anc
- まずはSQLightから、goのgo.modがある場所で以下をターミナルに打ち込む
  ```ターミナル
   go get github.com/mattn/go-sqlite3
  ```