package chapter3

// 課題1
// 以下のstructにgetterとsetterを実装してください。
// Getterの関数名ID, Name
// Setterの関数名SetID, SetName
type Kadai1 struct {
	id   int
	name string
}

func (i *Kadai1) ID() int {
	return i.id
}

func (i *Kadai1) Name() string {
	return i.name
}

func (i *Kadai1) SetName(str string) {
	i.name = str
}

func (i *Kadai1) SetID(id int) {
	i.id = id
}
