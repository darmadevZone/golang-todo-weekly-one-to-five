package docs

import "fmt"

type Person struct{
	Name string
	Age uint
}

func PointerFunc(){

	p := Person{
		Name: "Takechi",
		Age: 20,
	}
	fmt.Printf("最初のp :%+v\n",p)
	fmt.Printf("最初のp Address :%p\n", &p)
	
	p2 := p
	p2.Name = "二郎"
	p2.Age = 21

	fmt.Printf("p2アドレス :%p\n", &p2)
	fmt.Printf("p2で二郎に書き換えを行なったはずのp :%+v\n", p)

	// &pで*Person(Personのポインタ型)を生成する
	// p3はpのアドレスが格納されている状態になる
	p3 := &p
	p3.Name = "二郎"
	p3.Age = 21

	fmt.Printf("p3アドレス :%p\n", &p3)
	fmt.Printf("p3で二郎に書き換えを行なったp :%+v\n", p)
}
