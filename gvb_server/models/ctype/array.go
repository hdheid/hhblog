package ctype

import (
	"database/sql/driver"
	"strings"
)

/*
当我们在使用 Go 语言进行数据库操作时，经常会使用 database/sql 包。这个包提供了一套通用的接口，用于与各种数据库进行交互。其中，Scanner 和 Valuer 接口分别用于将数据库的值扫描到自定义类型中，以及将自定义类型的值转换为数据库可存储的值。

在您给出的代码中，定义了一个类型别名 Array，它是一个字符串切片类型。然后，为 Array 类型定义了两个方法：Scan 和 Value。

Scan 方法实现了 database/sql 包中的 Scanner 接口的方法。这个接口要求实现一个 Scan 方法，用于将数据库查询结果中的值扫描到自定义类型中。在这个方法中，value 参数表示数据库查询结果中的值。首先，通过类型断言将其转换为字节切片类型 []byte。然后，通过判断字节切片是否为空字符串来确定如何处理值。如果值为空字符串，则将 Array 类型切片设置为空切片。否则，将值按照换行符拆分成字符串切片，并将结果赋值给 Array 类型切片。

Value 方法实现了 database/sql 包中的 Valuer 接口的方法。这个接口要求实现一个 Value 方法，用于将自定义类型的值转换为数据库可存储的值。在这个方法中，将字符串切片转换为以换行符连接的字符串，并返回这个字符串作为 driver.Value。

通过定义这两个方法，Array 类型可以在使用 database/sql 包进行数据库操作时，被正确地处理和转换。您可以将 Array 类型作为查询结果的字段类型，或将其作为参数传递给数据库查询的参数。这样，Scan 方法将负责将数据库的值扫描到 Array 类型中，而 Value 方法将负责将 Array 类型的值转换为数据库可存储的值。

这样，您可以在数据库操作中方便地使用 Array 类型，并通过这两个方法实现自定义类型与数据库之间的交互。


目的：
这样做的目的是为了在使用 Go 语言进行数据库操作时，能够方便地处理自定义类型与数据库之间的转换。

当我们使用 database/sql 包进行数据库操作时，数据库查询结果的值通常是基本类型（如字符串、整数等）或标准库提供的类型（如 time.Time）。但是，有时候我们可能希望在应用程序中使用自定义的复杂类型，而不仅仅是基本类型。

通过实现 Scanner 和 Valuer 接口的方法，我们可以告诉 database/sql 包如何将数据库的值转换为自定义类型，以及如何将自定义类型的值转换为数据库可存储的值。这样，我们可以在应用程序中使用自定义类型，而无需手动进行类型转换。

在您给出的代码中，通过为 Array 类型定义 Scan 和 Value 方法，我们可以将数据库查询结果中的值扫描到 Array 类型中，或将 Array 类型的值转换为数据库可存储的值。这样，您可以在数据库操作中直接使用 Array 类型，而无需手动编写转换逻辑。

这种方式的好处是可以提高代码的可读性和可维护性。通过使用自定义类型和实现 Scanner 和 Valuer 接口的方法，我们可以更清晰地表达代码的意图，减少了手动类型转换的复杂性，并提供了更好的类型安全性。
*/

type Array []string

func (t *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) == "" {
		*t = []string{}
		return nil
	}
	*t = strings.Split(string(v), "\n")
	return nil
}

func (t Array) Value() (driver.Value, error) {
	// 将数字转换为值
	return strings.Join(t, "\n"), nil
}
