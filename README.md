# 2023-DevFest-Talk-Golang-v1.21-new-function-min-max-clear

This is research materials

## Let's look up the [source code](https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/math/dim.go;l=40) to understand how to use the `math.Max` and `math.Min` function.

### Max function

1. 參數：x 和 y。
2. if haveArchMax 檢查一個名為 [haveArchMax](https://go.dev/src/math/dim_noasm.go) 的 boolean。這個變量的目的是確定是否有一個特定於體系結構（architecture-specific）的 max function 實現可用。如果 haveArchMax 為 true，則調用 archMax function。
如果 haveArchMax 為 false，則調用通用的 max function。
3. 功能：用於計算兩個浮點數中的最大值。
特殊情況：
    1. 如果 x 或 y 是正無窮大（IsInf (x, 1) 或 IsInf (y, 1)）， function 返回正無窮大 (Inf (1))。
    2. 如果 x 或 y 是非數（NaN，即 不是一個數字 ），function 返回 NaN (NaN ())。
    3. 如果 x 和 y 都是 0，且 x 的符號位是負的（Signbit (x) 為 true），則返回 y；否則返回 x。這種情況處理了 -0.0 和 0.0 的比較。
4. 最後，一個簡單的比較決定了 x 和 y 中較大的數。如果 x > y，則返回 x；否則，返回 y。
總結來說，Max function 首先檢查是否有針對特定體系結構優化的實現，如果沒有，則使用一個通用的 function 來處理浮點數的特殊情況並比較兩個數值，返回較大的數值。


### Min function


1. 參數：x 和 y。
2. 這個 function 首先檢查一個全局變量 haveArchMin，如果這個變量為 true，則調用 archMin  function（這通常是針對特定硬件架構優化的版本）。如果 haveArchMin 為 false（表示沒有特定架構的優化），則調用 min function。
3. 功能：用於計算兩個浮點數中的最小值。
特殊情況：
    1. 如果 x 或 y 中的任何一個是負無窮大（-Inf），則返回 -Inf。
    2. 如果 x 或 y 是 NaN（不是數字），則返回 NaN。
    3. 如果 x 和 y 都是 0（包括正零和負零），則根據 IEEE 浮點數規則，返回負零（如果 x 是負零）或正零。
4.  最後，一個簡單的比較決定了 x 和 y 的值：如果 x < y，則返回 x；否則，返回 y。這個 function 的設計考慮到了浮點數運算中的一些特殊情況，比如無窮大和非數字值，以及正零和負零的區別。這種設計保證了 function 在處理特殊浮點數值時的正確性和一致性。

## Let's look up the [source code](https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/builtin/builtin.go;l=215) to understand Built-in `min` and `max` functions in Go 1.21 

### Max function

1. 功能：返回一组固定数量的 cmp.Ordered 類型參數中的最大值。cmp.Ordered 是一個 interface，表示可以比較的類型（如 int、float 等）。
2. 參數：x T, y ...T
    x T 是第一個參數，其類型為泛型 T。
    y ...T 表示變長參數，可以接收零個或多個類型為 T 的其他參數。
3. 返回值：類型為 T 的值，表示參數中的最大值。
4. 特殊情况：如果 T 是 float，并且任何一個參數是 NaN（不是一個數），那麽 max function 將返回 NaN。

### Min function

1. 功能：與 max function 相似，但它返回的是最小值。
2. 參數：與 max function 相同，x T 作為第一個參數，y ...T 作為可變數量的其他參數。
3. 返回值：類型為 T 的值，表示參數中的最小值。
4. 特殊情况：同 max，如果 T 是 float 類型，并且參數中包含 NaN，那麽 min 將返回 NaN。

### 總結

在這兩個 function 中，泛型 T [cmp.Ordered](https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/cmp/cmp.go;l=18) 意味著這些 function 可以接受任何實現了 cmp.Ordered interface 的 類型。這種泛型的使用提高了程式碼的通用性和靈活性。
例如，你可以使用 max function 来比較兩個 int、float 或甚至自定義類型（只要這些類型實現了 cmp.Ordered 接口）。這使得 max 和 min function 可以在多種不同類型的數據上工作，而不需要為每種數據類型單獨寫 function。

## Let's look up the [source code](https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/builtin/builtin.go;l=251) to understand how to use the `clear` function

### Clear function


1. 功能：這個 function 用於清空映射（maps）和切片（slices）
2. function 定義: func clear [T ~[] Type | ~map [Type] Type1](t T) 是一個泛型 function，使用了 Go 1.18 引入的泛型語法。T 是一個類型參數，它被限制為只能是 slice 或 map 的類型。
3. 類型限制:
    ~[] Type 表示 T 可以是任何與某種 slice 類型相似（相同底層類型）的類型。
    ~map [Type] Type1 表示 T 可以是任何與某種 map 類型相似的類型。
4. 對於 map 的操作:
    當 T 是 map 類型時，clear function 會刪除 map 中的所有條目，使其變為空 map。
5. 對於 slice 的操作:
    當 T 是 slice 類型時，clear function 會將 slice 中的所有元素設置為它們各自元素類型的零值，直到達到 slice 的長度。
6. 泛型約束:
    如果 T 是一個類型參數（在泛型 function 或接口中），則 T 的類型集合必須只包含 map 或 slice 類型，並且 clear 將根據類型參數的實際類型執行相應的清空操作。

總的來說，clear function 是一個泛型 function，可以用來清空 map 和 slice 的內容，無論它們具體是什麼類型。這種設計增加了程式碼的通用性和重用性。

- 解決的問題：

1. Slice：以前需要使用 s = s [:0] 或 s = nil 等方法來清空 slice，這些方法不夠直觀且一致，可能導致內存泄漏或性能下降。使用 clear (s) 可以更簡潔高效地完成清空操作，同時保留底層內存。
2. Map：以前需要使用 for k := range m {delete (m, k) } 或 m = nil 等方法來清空 map，這些方法不夠優雅和安全，可能導致內存浪費或競爭條件。使用 clear (m) 使清空操作更容易和安全。

```
	s := []int{1, 2, 3, 4, 5}
	// Approach 1
	s = s[:0]
	// Approach 2
	s = nil

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// Approach 1
	for k := range m {
		delete(m, k)
	}
	// Approach 2
	m = nil
```

## Why is it designed this way?

### Python 的設計理念:

1. 使用方便與靈活性： Python 的設計目的是要易於閱讀和寫作。它的 max function 能接受列表、元組和其他可迭代結構，這符合 Python 以使用者友好和靈活性為主的一般哲學。Python 強調開發者的生產力和代碼的可讀性。
2. 動態類型：Python 是動態類型的，這允許像 max 這樣的 function 處理不同類型的輸入，從數字到列表，而不需要單獨的實現或重載。

### Go 的設計理念:

1. 簡潔性和效率： Go 的設計重點是簡潔、清晰和效率。 function 只接受兩個確切的值，這簡單明了，符合 Go 旨在擁有簡單、不含糊的語法的目標。Go 偏好明確的行為，而不是背後隱藏的「魔法」。
2. 靜態類型和性能： Go 是靜態類型的，其設計通常傾向於性能和可預測性，而不是靈活性。在 Go 中，function 通常被設計來做一件事並且做得好。Go 沒有內建的可變參數 max function（一個可以接受任意數量參數的 function），這反映了這種理念。
3. 錯誤處理和清晰度： Go 強調明確的錯誤處理和清晰的代碼流程。讓 max function 接受列表（在 Go 中是 slice ）作為參數可能引入更多的複雜性（例如處理空列表）和模糊性（如混合類型列表），這通常是 Go 所避免的。

總結來說，Python 的設計目標在於開發者的便利性和靈活性，而 Go 的設計則聚焦於簡潔性、清晰度和性能。這些設計選擇影響了內建 function 和方法的實現方式及它們接受的參數類型。
