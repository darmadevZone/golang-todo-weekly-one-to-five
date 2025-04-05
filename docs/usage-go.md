# Usage-Golang

- ポインターについて
    [ポインターについて https://qiita.com/Sekky0905/items/447efa04a95e3fec217f](https://qiita.com/Sekky0905/items/447efa04a95e3fec217f)

    ```
    // ポインタ型の変数を宣言する
    // pがポインタ変数
    // *Personポインタ型
    var p *Person

    p = &Person{
        Name: "太郎",
        Age:  20,
    }
    fmt.Printf("変数pに格納されているアドレス :%p", p) //0x1040a0d0
    ```

## Http Error Status
