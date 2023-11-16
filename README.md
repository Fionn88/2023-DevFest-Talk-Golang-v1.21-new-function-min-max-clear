# 2023-DevFest-Talk-Golang-v1.21-new-function-min-max-clear

This is research materials

## Let's look up the [source code](https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/math/dim.go;l=40) to understand how to use the `math.Max` and `math.Min` function.

### Max function
```
函數接收兩個float64類型的參數：x和y。
if haveArchMax 檢查一個名為 haveArchMax 的布爾變量。這個變量的目的是確定是否有一個特定於體系結構（architecture-specific）的max函數實現可用。如果 haveArchMax 為 true，則調用 archMax 函數。
如果 haveArchMax 為 false，則調用通用的 max 函數。

這是一個通用的函數，用於計算兩個浮點數中的最大值。
首先檢查一些特殊情況：
1. 如果 x 或 y 是正無窮大（IsInf(x, 1) 或 IsInf(y, 1)），函數返回正無窮大 (Inf(1))。
2. 如果 x 或 y 是非數（NaN，即 不是一個數字 ），函數返回 NaN (NaN())。
3. 如果 x 和 y 都是 0，且 x 的符號位是負的（Signbit(x) 為 true），則返回 y；否則返回 x。這種情況處理了 -0.0 和 0.0 的比較。

最後，一個簡單的比較決定了 x 和 y 中較大的數。如果 x > y，則返回 x；否則，返回 y。
總結來說，Max 函數首先檢查是否有針對特定體系結構優化的實現，如果沒有，則使用一個通用的函數來處理浮點數的特殊情況並比較兩個數值，返回較大的數值。
```

### Min function

```
接收兩個 float64 類型的參數 x 和 y。這個函數首先檢查一個全局變量 haveArchMin，如果這個變量為 true，則調用 archMin 函數（這通常是針對特定硬件架構優化的版本）。如果 haveArchMin 為 false（表示沒有特定架構的優化），則調用 min 函數。

它首先處理一些特殊情況：

1. 如果 x 或 y 中的任何一個是負無窮大（-Inf），則返回 -Inf。
2. 如果 x 或 y 是 NaN（不是數字），則返回 NaN。
3. 如果 x 和 y 都是 0（包括正零和負零），則根據 IEEE 浮點數規則，返回負零（如果 x 是負零）或正零。

比較 x 和 y 的值: 如果 x < y，則返回 x；否則，返回 y。
這個函數的設計考慮到了浮點數運算中的一些特殊情況，比如無窮大和非數字值，以及正零和負零的區別。這種設計保證了函數在處理特殊浮點數值時的正確性和一致性。
```

## Built-in min and max functions in Go 1.21 [source code](https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/builtin/builtin.go;l=215)

### Max function

1. 功能：返回一组固定数量的cmp.Ordered类型参数中的最大值。cmp.Ordered是一个接口，表示可以比较的类型（如整数、浮点数等）。
2. 参数：x T, y ...T
    x T是第一个参数，其类型为泛型T。
    y ...T表示变长参数，可以接收零个或多个类型为T的其他参数。
3. 返回值：类型为T的值，表示参数中的最大值。
4. 特殊情况：如果T是浮点类型，并且任何一个参数是NaN（不是一个数），那么max函数将返回NaN。

### Min function

1. 功能：与max函数相似，但它返回的是最小值。
2. 参数：与max函数相同，x T作为第一个参数，y ...T作为可变数量的其他参数。
3. 返回值：类型为T的值，表示参数中的最小值。
4. 特殊情况：同max，如果T是浮点类型，并且参数中包含NaN，那么min将返回NaN。

在这两个函数中，泛型T cmp.Ordered意味着这些函数可以接受任何实现了cmp.Ordered接口的类型。这种泛型的使用提高了代码的通用性和灵活性。


### 總結

例如，你可以使用max函数来比较两个整数、浮点数或甚至自定义类型（只要这些类型实现了cmp.Ordered接口）。这使得max和min函数可以在多种不同类型的数据上工作，而不需要为每种数据类型单独编写函数。

## Let's look up the [source code](https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/builtin/builtin.go;l=251) to understand how to use the `clear` function

### Clear function

```
這個函數用於清空映射（maps）和切片（slices）。讓我們分析這個函數的實作細節：

1. 函數定義: func clear[T ~[]Type | ~map[Type]Type1](t T) 是一個泛型函數，使用了 Go 1.18 引入的泛型語法。T 是一個類型參數，它被限制為只能是切片或映射的類型。
2. 類型限制:
    ~[]Type 表示 T 可以是任何與某種切片類型相似（相同底層類型）的類型。
    ~map[Type]Type1 表示 T 可以是任何與某種映射類型相似的類型。
3. 對於映射的操作:
    當 T 是映射類型時，clear 函數會刪除映射中的所有條目，使其變為空映射。
4. 對於切片的操作:
    當 T 是切片類型時，clear 函數會將切片中的所有元素設置為它們各自元素類型的零值，直到達到切片的長度。
5. 泛型約束:
    如果 T 是一個類型參數（在泛型函數或接口中），則 T 的類型集合必須只包含映射或切片類型，並且 clear 將根據類型參數的實際類型執行相應的清空操作。

總的來說，clear 函數是一個泛型函數，可以用來清空映射和切片的內容，無論它們具體是什麼類型。这种设计增加了代码的通用性和重用性。
```

- 解決的問題：

1. Slice：以前需要使用 s = s[:0] 或 s = nil 等方法來清空切片，這些方法不夠直觀且一致，可能導致內存泄漏或性能下降。使用 clear(s) 可以更簡潔高效地完成清空操作，同時保留底層內存。
2. Map：以前需要使用 for k := range m { delete(m, k) } 或 m = nil 等方法來清空映射，這些方法不夠優雅和安全，可能導致內存浪費或競爭條件。使用 clear(m) 使清空操作更容易和安全。

## Why is it designed this way?

### Python的設計理念:

1. 使用方便與靈活性： Python的設計目的是要易於閱讀和寫作。它的max函數能接受列表、元組和其他可迭代結構，這符合Python以使用者友好和靈活性為主的一般哲學。Python強調開發者的生產力和代碼的可讀性。
2. 動態類型： Python是動態類型的，這允許像max這樣的函數處理不同類型的輸入，從數字到列表，而不需要單獨的實現或重載。

### Go的設計理念:

1. 簡潔性和效率： Go的設計重點是簡潔、清晰和效率。函數只接受兩個確切的值，這簡單明了，符合Go旨在擁有簡單、不含糊的語法的目標。Go偏好明確的行為，而不是背後隱藏的「魔法」。
2. 靜態類型和性能： Go是靜態類型的，其設計通常傾向於性能和可預測性，而不是靈活性。在Go中，函數通常被設計來做一件事並且做得好。Go沒有內建的可變參數max函數（一個可以接受任意數量參數的函數），這反映了這種理念。
3. 錯誤處理和清晰度： Go強調明確的錯誤處理和清晰的代碼流程。讓max函數接受列表（在Go中是切片）作為參數可能引入更多的複雜性（例如處理空列表）和模糊性（如混合類型列表），這通常是Go所避免的。

總結來說，Python的設計目標在於開發者的便利性和靈活性，而Go的設計則聚焦於簡潔性、清晰度和性能。這些設計選擇影響了內建函數和方法的實現方式及它們接受的參數類型。

