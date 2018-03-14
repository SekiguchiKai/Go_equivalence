# Goにおける等値と等価の考察
Goにおいての等値と等価を考察してみた。
ツッコミがあったら、教えてください。

## 等値と等価
まず、そもそもの等値と等価を整理してみる。

### 等値
> 指しているものが「完全に同一の存在」であること
> (つまり同じアドレスを指していること)

引用元 : 中山 清喬(2014/9/22)『スッキリわかるJava入門 実践編 第2版 スッキリわかるシリーズ』 インプレスジャパン
Javaでいうところの `hoge == foo`

### 等価
> 指している2つのものが「同じ内容」であること
> (同じアドレスを指していなくてもよい)

引用元 : 中山 清喬(2014/9/22)『スッキリわかるJava入門 実践編 第2版 スッキリわかるシリーズ』 インプレスジャパン
Javaでいうところの `hoge.equals(foo)`


## reflect.DeepEqual
以下の引用文は[reflect - The Go Programming Language](https://golang.org/pkg/reflect/#DeepEqual)から引用させていただいている。

> Struct values are deeply equal if their corresponding fields, both exported and unexported, are deeply equal.

(意訳)構造体の値は、対応するフィールド（エクスポートと非エクスポートの両方(大文字と小文字)）がdeeply equalならば、deeply equalになる。

> Pointer values are deeply equal if they are equal using Go's == operator or if they point to deeply equal values.

(意訳)ポインタの値は、Goの==演算子を使用して等しい場合、またはそれらがdeeply equalな値を指す場合deeply equalになる。

> Slice values are deeply equal when all of the following are true: they are both nil or both non-nil, 
they have the same length, and either they point to the same initial entry of the same underlying array (that is,
 &x[0] == &y[0]) or their corresponding elements (up to length) are deeply equal.
 Note that a non-nil empty slice and a nil slice (for example, []byte{} and []byte(nil)) are not deeply equal.

(意訳)Sliceの値は、次の条件がすべて成立している場合には、等しくなる。
1. Sliceの値は、nilまたはnon-nilの両方である
2. 同じ長さである
3. 元になる配列が（つまり＆x [0 ] ==＆y [0]）同一か、対応する要素が全部deeply equalである
! 空でない(non-nil)Sliceと空(nil)のSliceと無限スライス（たとえば、[]byte{}と[] byte（nil））はdeeply equalにならないことに注意。


> Other values - numbers, bools, strings, and channels - are deeply equal if they are equal using Go's == operator.

(意訳)他の値（数値、bool、文字列、およびチャンネル）は、Go ==演算子を使用して等しい場合、等しくなる。

簡単にいうと、普通の型には == で比較した感じで、コンポジット型は再帰的に比較した感じ。

### struct1==struct1と&struct1==&struct1とreflect.DeepEqual
### 等値を確認したい時
&structでもって、比較する。アドレスを比較し、指しているものが「完全に同一か」どうかを確認する。`&struct1 == &struct2` を用いる。
`reflect.DeepEqual(&struct1, &struct2)` は一見等値に見えるが、そうではないので注意!

 
### 等価を確認したい時
structでもって、比較する。アドレスは別で良くて、単純に指しているものが「同じ内容か」どうかを確認する。`struct1 == struct2` か `reflect.DeepEqual(struct1, struct2)` だが、後者を使用した方がよさそう。

### 

## 単純な構造体の比較
単純な構造体(ここでは、構造体のフィールドにSliceやポインタ型を含まない物を単純な構造体と呼ぶことにする)同士の等値と等価

### 実装

```go
	fmt.Println(bar)

	p1 := Person{
		Name: "太郎",
		Age:  20,
	}

	p2 := Person{
		Name: "太郎",
		Age:  20,
	}

	fmt.Println("単純な構造体の比較")
	fmt.Printf("p1 == p2 : 等価(メモリが別でも値が一緒だったらOK) : %t\n", p1 == p2)
	fmt.Printf("&p1 == &p2 : 等値 : %t\n", &p1 == &p2)
	fmt.Printf("reflect.DeepEqual(p1, p2) : 等価 : %t\n", reflect.DeepEqual(p1, p2))
	fmt.Printf("reflect.DeepEqual(&p1, &p2)  : %t\n", reflect.DeepEqual(&p1, &p2))

	fmt.Println(bar)
	fmt.Println("単純な構造体の比較2")
	p3 := p1
	fmt.Printf("p1 == p3 : 等価(メモリが別でも値が一緒だったらOK) : %t\n", p1 == p3)
	fmt.Printf("&p1 == &p3 : 等値 : %t\n", &p1 == &p3)
	fmt.Printf("reflect.DeepEqual(p1, p3) : 等価 : %t\n", reflect.DeepEqual(p1, p3))
	fmt.Printf("reflect.DeepEqual(&p1, &p3) : %t\n", reflect.DeepEqual(&p1, &p3))

```

### 結果

```
===========================================================
単純な構造体の比較
p1 == p2 : 等価(メモリが別でも値が一緒だったらOK) : true
&p1 == &p2 : 等値 : false
reflect.DeepEqual(p1, p2) : 等価 : true
reflect.DeepEqual(&p1, &p2)  : true
===========================================================
単純な構造体の比較2
p1 == p3 : 等価(メモリが別でも値が一緒だったらOK) : true
&p1 == &p3 : 等値 : false
reflect.DeepEqual(p1, p3) : 等価 : true
reflect.DeepEqual(&p1, &p3) : true
===========================================================
```


## 構造体のポインタ型の比較
構造体のポインタ型同士の等値と等価

### 実装

```go
	fmt.Println(bar)
	fmt.Println("構造体のポインタ型の比較(同一ポインタ)")

	p1 := Person{
		Name: "太郎",
		Age:  20,
	}

	p2 := Person{
		Name: "太郎",
		Age:  20,
	}

	p3 := &p1
	fmt.Printf("&p1 == p3 : 等値 : %t\n", &p1 == p3)
	fmt.Printf("reflect.DeepEqual(&p1, p3) : 等価 : %t\n", reflect.DeepEqual(&p1, p3))

	fmt.Println(bar)
	fmt.Println("構造体のポインタ型の比較(別ポインタ)")
	p4 := &p2
	fmt.Printf("&p1 == p5 : 等値 : %t\n", &p1 == p4)
	fmt.Printf("reflect.DeepEqual(&p1, p4) : 等価 : %t\n", reflect.DeepEqual(&p1, p4))
```

### 結果

```
===========================================================
構造体のポインタ型の比較(同一ポインタ)
&p1 == p3 : 等値 : true
reflect.DeepEqual(&p1, p3) : 等価 : true
===========================================================
構造体のポインタ型の比較(別ポインタ)
&p1 == p5 : 等値 : false
reflect.DeepEqual(&p1, p4) : 等価 : true
===========================================================
```


## フィールドにSliceが入っている構造体の比較
フィールドにSliceが入っている構造体同士の比較
フィールドにSliceが入っている構造体同士の比較で、比較する場合、構造体自体は異なる物なのに、構造体の中に存在するSliceが同一アドレスの場合、`&struct1 == &struct2` でも、trueになるので注意。

### 実装

```go
	fmt.Println(bar)

	sa := []int{0, 1, 2, 3, 4, 5}
	sb := []int{0, 1, 2, 3, 4, 5}

	s1 := StructIncludingSlice{
		Numbers: sa,
	}

	s2 := StructIncludingSlice{
		Numbers: sb,
	}

	fmt.Println("フィールドにSliceが入っている構造体の比較")
	// fmt.Printf("等価(メモリが別でも値が一緒だったらOK) : %t\n", s1 == s2) // struct containing []int cannot be compared

	fmt.Println("s1 == s2 : 等値(エラー) : struct containing []int cannot be compared")
	fmt.Printf("&s1 == &s2 : 等値 : %t\n", &s1 == &s2)
	fmt.Printf("reflect.DeepEqual(s1, s2) : 等価 : %t\n", reflect.DeepEqual(s1, s2))
	fmt.Printf("reflect.DeepEqual(&s1, &s2) : %t\n", reflect.DeepEqual(&s1, &s2))

	fmt.Println(bar)
	fmt.Println("フィールドにSliceが入っている構造体の比較(Sliceが同じ物)")
	s3 := StructIncludingSlice{
		Numbers: sa,
	}

	fmt.Println("s1 == s3 : 等値(エラー) : struct containing []int cannot be compared")
	fmt.Printf("&s1 == &s3 : 等値 : %t\n", &s1 == &s3)
	fmt.Printf("reflect.DeepEqual(s1, s3) : 等価 : %t\n", reflect.DeepEqual(s1, s3))
	fmt.Printf("reflect.DeepEqual(&s1, &s3)  : %t\n", reflect.DeepEqual(&s1, &s3))

	fmt.Println(bar)
	fmt.Println("フィールドにSliceが入っている構造体の比較(同じアドレスの場合)")
	s4 := &s1

	fmt.Println("s1 == s4 : 等値(エラー) : struct containing []int cannot be compared")
	fmt.Printf("&s1 == s4 : 等値 : %t\n", &s1 == s4)
	fmt.Printf("reflect.DeepEqual(s1, s4) : 等価 : %t\n", reflect.DeepEqual(s1, s4))
	fmt.Printf("reflect.DeepEqual(&s1, s4)  : %t\n", reflect.DeepEqual(&s1, s4))
```

### 結果


```
===========================================================
フィールドにSliceが入っている構造体の比較
s1 == s2 : 等値(エラー) : struct containing []int cannot be compared
&s1 == &s2 : 等値 : false
reflect.DeepEqual(s1, s2) : 等価 : true
reflect.DeepEqual(&s1, &s2) : true
===========================================================
フィールドにSliceが入っている構造体の比較(Sliceが同じ物)
s1 == s3 : 等値(エラー) : struct containing []int cannot be compared
&s1 == &s3 : 等値 : false
reflect.DeepEqual(s1, s3) : 等価 : true
reflect.DeepEqual(&s1, &s3)  : true
===========================================================
フィールドにSliceが入っている構造体の比較(同じアドレスの場合)
s1 == s4 : 等値(エラー) : struct containing []int cannot be compared
&s1 == s4 : 等値 : true
reflect.DeepEqual(s1, s4) : 等価 : false
reflect.DeepEqual(&s1, s4)  : true
===========================================================
```

### 実装
```go
fmt.Println(bar)
	fmt.Println("StructIncludingPointerの実験")
	num1 := 1
	num2 := 1

	s1 := StructIncludingPointer{
		Pointer: &num1,
	}

	s2 := StructIncludingPointer{
		Pointer: &num2,
	}

	// プロパティにポインタ（スライス含む）が入っていると動作
	fmt.Println("フィールドに異なるアドレスを指す同じ数字が入っている構造体の比較")
	fmt.Printf("s1 == s2 : この場合、構造体のフィールドの等値での比較になる : %t\n", s1 == s2)
	fmt.Printf("&s1 == &s2 : %t\n", &s1 == &s2)
	fmt.Printf("reflect.DeepEqual(s1, s2) : 等価 : %t\n", reflect.DeepEqual(s1, s2))
	fmt.Printf("reflect.DeepEqual(&s1, &s2) : %t\n", reflect.DeepEqual(&s1, &s2))
```

### 結果

```
===========================================================
StructIncludingPointerの実験
フィールドに異なるアドレスを指す同じ数字が入っている構造体の比較
s1 == s2 : この場合、構造体のフィールドの等値での比較になる : false
&s1 == &s2 : false
reflect.DeepEqual(s1, s2) : 等価 : true
reflect.DeepEqual(&s1, &s2) : true
===========================================================
```

## フィールドに異なるアドレスを指す同じ数字が入っている構造体の比較
フィールドに異なるアドレスを指す同じ数字が入っている構造体同士の比較

### 実装

```go
	fmt.Println(bar)

	numX := 2
	s3 := StructIncludingPointer{
		Pointer: &numX,
	}

	s4 := StructIncludingPointer{
		Pointer: &numX,
	}

	fmt.Println("フィールドに同一アドレスを指す同じ数字が入っている構造体の比較")
	fmt.Printf("s3 == s4 : 構造体のフィールドの等値での比較になる : %t\n", s3 == s4)
	fmt.Printf("&s3 == &s4 : %t\n", &s3 == &s4)
	fmt.Printf("reflect.DeepEqual(s3, s4) : 等価 : %t\n", reflect.DeepEqual(s3, s4))
	fmt.Printf("reflect.DeepEqual(&s3, &s4)  : %t\n", reflect.DeepEqual(&s3, &s4))

	fmt.Println(bar)
```

### 結果

```
===========================================================
フィールドに同一アドレスを指す同じ数字が入っている構造体の比較
s3 == s4 : 構造体のフィールドの等値での比較になる : true
&s3 == &s4 : false
reflect.DeepEqual(s3, s4) : 等価 : true
reflect.DeepEqual(&s3, &s4)  : true
===========================================================
```


## 参考文献
中山 清喬(2014/9/22)『スッキリわかるJava入門 実践編 第2版 スッキリわかるシリーズ』 インプレスジャパン

## 参考にさせていただいた記事
[reflect - The Go Programming Language](https://golang.org/pkg/reflect/)


