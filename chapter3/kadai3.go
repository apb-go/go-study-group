package chapter3

import "fmt"

type Dog struct{}

func (d Dog) Bark() string {
	return "わんわん"
}

type Cat struct{}

func (c Cat) Crow() string {
	return "にゃーにゃ"
}

// 課題3
// この関数の引数はxの型は不定です。
// 型がDogの場合はBow()を実行した結果
// Catの場合はCrowを実行した結果
// その他の場合はerrorを返却してください。
func Kadai3(x interface{}) (string, error) {
	dog, ok := x.(Dog)
	if ok {
		return dog.Bark(), nil
	}
	cat, ok := x.(Cat)
	if ok {
		return cat.Crow(), nil
	}
	return "", fmt.Errorf("Cast error.")
}
