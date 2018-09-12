package main

import "fmt"

type Stereotype int

type AudioOutput int

const (
	OutMute   AudioOutput = iota // 0
	OutMono                      // 1
	OutStereo                    // 2
	_
	_
	OutSurround // 5
)

const (
	TypicalNoob           Stereotype = iota + 1001 // 0
	TypicalHipster                                 // 1
	TypicalUnixWizard                              // 2
	TypicalStartupFounder                          // 3
)

// iota 可以做更多事情，而不仅仅是 increment。更精确地说，iota 总是用于 increment，但是它可以用于表达式，在常量中的存储结果值。

// 这里我们创建一个常量用于位掩码。

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

// 这个工作是因为当你在一个 const 组中仅仅有一个标示符在一行的时候，它将使用增长的 iota 取得前面的表达式并且再运用它，。在 Go 语言的 spec 中， 这就是所谓的隐性重复最后一个非空的表达式列表。

// 如果你对鸡蛋，巧克力和海鲜过敏，把这些 bits 翻转到 “on” 的位置（从左到右映射 bits）。然后你将得到一个 bit 值 00010011，它对应十进制的 19。

// fmt.Println(IgEggs | IgChocolate | IgShellfish)

// 这是在 Effective Go 中一个非常好定义数量级的示例：

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

// 当你在把两个常量定义在一行的时候会发生什么？
// Banana 的值是什么？ 2 还是 3？ Durian 的值又是？

const (
	Apple, Banana = iota + 1000, iota
	Cherimoya, Durian
	Elderberry, Fig
)

// iota 在下一行增长，而不是立即取得它的引用。

// Apple: 1
// Banana: 2
// Cherimoya: 2
// Durian: 3
// Elderberry: 3
// Fig: 4
// 这搞砸了，因为现在你的常量有相同的值。

// 因此，对的

func CountAllTheThings(i int) string {
	return fmt.Sprintf("there are %d things", i)
}

// Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants. Its value is the index of the respective ConstSpec in that constant declaration, starting at zero. It can be used to construct a set of related constants:

const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

const (
	a = 1 << iota // a == 1  (iota == 0)
	b = 1 << iota // b == 2  (iota == 1)
	c = 3         // c == 3  (iota == 2, unused)
	d = 1 << iota // d == 8  (iota == 3)
)

const (
	u         = iota * 42 // u == 0     (untyped integer constant)
	v float64 = iota * 42 // v == 42.0  (float64 constant)
	w         = iota * 42 // w == 84    (untyped integer constant)
)

const x = iota // x == 0
const y = iota // y == 0
// By definition, multiple uses of iota in the same ConstSpec all have the same value:

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0  (iota == 0)
	bit1, mask1                          // bit1 == 2, mask1 == 1  (iota == 1)
	_, _                                 //                        (iota == 2, unused)
	bit3, mask3                          // bit3 == 8, mask3 == 7  (iota == 3)
)

// This last example exploits the implicit repetition of the last non-empty expression list.

func main() {
	//	n := TypicalHipster
	//一个戏剧性的转折 尽管如此。你可以传递一个数值常量，然后它能工作。
	// 这是因为常量在 Go 中是弱类型直到它使用在一个严格的上下文环境中。
	fmt.Println(CountAllTheThings(0))
	//	fmt.Println(CountAllTheThings(n))
	//cannot use n (type Stereotype) as type int in argument to CountAllTheThings
	//

	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)

	fmt.Println(TypicalNoob,
		TypicalHipster,
		TypicalUnixWizard, // 2
		TypicalStartupFounder)
}
