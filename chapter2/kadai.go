package chapter2

import "fmt"

// 引数のスライスsliceの要素数が
// 0の場合、0とエラー
// 2以下の場合、要素を掛け算
// 3以上の場合、要素を足し算
// を返却。正常終了時、errorはnilでよい
func Calc(slice []int) (int, error) {
	// TODO Q1
	// ヒント：エラーにも色々な生成方法があるが、ここではシンプルにfmtパッケージの
	// fmt.Errorf(“invalid op=%s”, op) などでエラー内容を返却するのがよい
	// https://golang.org/pkg/fmt/#Errorf

	length := len(slice)

	if length == 0 {
		return 0, fmt.Errorf("invalid len=%d", length)
	} else if length < 2 {
		return slice[0], nil
	} else if length < 3 {
		return slice[0] * slice[1], nil
	}
	r := 0
	for _, num := range slice {
		r = r + num
	}

	return r, nil
}

type Number struct {
	index int
}

// 構造体Numberを3つの要素数から成るスライスにして返却
// 3つの要素の中身は[{1} {2} {3}]とし、append関数を使用すること
func Numbers() []Number {
	// TODO Q2
	var slice []Number
	for i := 0; i < 3; i++ {
		slice = append(slice, Number{i + 1})
	}
	return slice
}

// 引数mをforで回し、「値」部分だけの和を返却
// キーに「yon」が含まれる場合は、キー「yon」に関連する値は除外すること
func CalcMap(m map[string]int) int {
	// TODO Q3
	n := 0
	for k, v := range m {
		if k == "yon" {
			continue
		}
		n += v
	}
	return n
}

type Model struct {
	Value int
}

// 与えられたスライスのModel全てのValueに5を足す破壊的な関数を作成
func Add(models []Model) {
	for i := range models {
		models[i].Value += 5
	}
}

// 引数のスライスには重複な値が格納されているのでユニークな値のスライスに加工して返却
// 順序はスライスに格納されている順番のまま返却すること
// ex) 引数:[]slice{21,21,4,5} 戻り値:[]int{21,4,5}
func Unique(slice []int) []int {
	// TODO Q5
	m := make(map[int]bool)
	var u []int
	for _, v := range slice {
		if !m[v] {
			m[v] = true
			u = append(u, v)
		}
	}
	return u
}

// 連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返却
func Fibonacci() func() int {
	// TODO Q6 オプション
	i, j := 0, 1
	return func() int {
		k := i
		i = j
		j = k + j
		return k
	}
}
