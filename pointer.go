package main

import "fmt"

// Person は人物を表す構造体。
type Person struct {
	Name string
	Age  int
}

// ポインタ型とポインタ型変数
func PointerTypeAndPointerTypeVariable()  {
	fmt.Printf("ポインタ型とポインタ型変数 PointerTypeAndPointerTypeVariable()\n")

	/*
		ポインタ変数とは
		メモリ上のアドレスを値として入れられる変数のこと

		& を使うことで、ポインタ型を生成することができる。
		Person型の変数pを &p とすると、Personへのポインタである *Person型 の値を生み出すことができる。
		&p は、pのアドレスという。

		Goでは、構造体内のメソッド内で、構造体のフィールドの情報を変更するときには、この参照渡しをよく利用する。こことかが参考になる。
	*/

	var p *Person

	//	ポインタ型の変数を宣言する
	//	pがポインタ型変数
	//	*Personポインタ型
	p = &Person{
		Name: "一郎",
		Age:  10,
	}

	// 変数pがポインタ変数となり、実際にpにはアドレスが格納されている。
	fmt.Printf("p :%+v　", p)
	fmt.Printf("変数pに格納されているアドレス :%p\n", p)
}

// デリファレンス
func DeReference()  {
	fmt.Printf("デリファレンス DeReference()\n")

	/*
		pはポインタではなく、Person型の値である。
		p2 := p
		は、Person型の値コピーしてp2に格納しているので、p2で書き換えを行っても、それがpに反映されることはない。これを値渡しという。

		p3 := &p
		は、*Person型(Personへのポインタである *Person型)をp3に格納しているので
		p3はpのアドレス(Personへのポインタである *Person型)を持っていることになる。
		従って、p3で書き換えを行うと、その変更はpに反映される。これを参照渡しという。
	 */

	// 値として、pに代入
	p := Person{
		Name: "一郎",
		Age:  10,
	}

	fmt.Printf("最初のp :%+v\n", p)

	// 値渡し
	p2 := p
	p2.Name = "二郎"
	p2.Age = 20
	// pではなく
	fmt.Printf("値渡し　p2で二郎に書き換えを行なったはずのp :%+v\n", p)

	// 参照渡し
	// &pで*Person(Personのポインタ型)を生成する
	// p3はpのアドレスが格納されている状態になる
	p3 := &p
	p3.Name = "二郎"
	p3.Age = 20

	fmt.Printf("参照渡し　p3で二郎に書き換えを行なったp :%+v\n", p)
}

// ポインタ型が格納された変数
func VariableStoredPointerType()  {
	fmt.Printf("ポインタ型が格納された変数 VariableStoredPointerType()\n")

	/*
		& を使うことで、ポインタ型を生成することができた。
		では、& を使って生成されたポインタ型を格納した変数はどう扱うか。
	 */

	name := "一郎"
	fmt.Printf("name :%v\n", name)

	namePoint := &name
	// namePointは、&nameが格納されているだけなので、stringへのポインタである*string型の値が格納されている。
	fmt.Printf("namePoint :%v\n", namePoint)
	// namePointが指している変数は、"*namePoint"という感じで、"*"をつけて表す。
	fmt.Printf("namePoint :%v\n", *namePoint)

	/*
		コードに示したように namePoint には &name が格納されている。
		& は、ポインタ型を生成するので&name は、stringへのポインタである *string型 の値(アドレス)が格納されている。
		よって、 namePoint を表示すると *string型 の値である name のアドレスが格納されていることがわかる。

		では、namePoint の元となっている name の変数に格納されている値(ここでは「太郎」)は、どのように取得すれば良いか。
		そのような場合は、 *namePoint のように変数名の前に * をつければ良い。
		なお、ここが紛らわしいところなのだが、 *namePoint 自体も変数なので、これに代入することもできる。
	 */

	*namePoint = "二郎"
	// *namePointに値を代入することもできる。
	fmt.Printf("*namePointに二郎を代入後の*namePoint :%v\n", *namePoint)
	// 再代入したところで、namePointに格納されている*string型の値(アドレス)自体は、変わらない
	fmt.Printf("*namePointに二郎を代入後のnamePoint :%v\n", namePoint)
	// stringへのポインタである*string型の値(nameに格納されている値)を書き換えたので、nameの値も変更される。
	fmt.Printf("*namePointに二郎を代入後のname :%v\n", name)

	/*
		ここで注意すべきことは、
		*namePoint に値を代入すると、nameの値も書き変わるということだ。
		これはなぜか?
		*namePoint には、 &name (stringへのポインタである*string型の値が格納されているからであり、
		それを *namePoint = "二郎" で書き換えているので、当然 name の値も書き変わるということである。
	 */
}

func main() {
	// PointerTypeAndPointerTypeVariable()
	// DeReference()
	// VariableStoredPointerType()
}
