package main

import (
	//"errors"
	"fmt"

	//"strings"
	//"log"
	//"math/big"
	//"internal/oserror"
	//"strconv"
	//"math"
	"math/rand/v2"
	//"os"
	//"bufio"
	//"reflect"
)

// 2.5 Переполнение и приведение типов int
/*
// у типа даных int есть подтипы int 8/16/32/64
// это не одинаковые типы а раные
// для того что бы значение одного типа int влождить
// в дргуо нужно компелятору объяснить что мы
// работаем со значением переменой разных типов
// если переменную типа int8 нужно вложить в переменую
// типа int то выглдяеть это будет так:
func main() {
	var i8 int8 = 127
	var i int = int(i8)
	fmt.Print(i)
}
// так же если мы переполним переменую ти не будет ошибки
// счет начнеться занаво с конца, то есть,
// в переменную типа int8 можно присвоить от -128 127
// это значит что если мы попробуем переменой типа
// int8 присвоить значение больше 127, например 128
// то счёт начнеться сначал, то есть с конца и
// и тем самым присвоеться значение -128, если будем
// присваивать 129 получим -127 и тд.
*/

// Типы byte и rune
// Предназначены для работы с знаками алфавита и другими

// системы исчесления
/*
из шестнадцатиричнйо системы исчисления в десятиричную
1С3
1 * 16(2) + (С = 12) 12 * 16(1) + 3 * 16(0)
256 + 192 + 3 = 451
func main() {
	//	a := 1 * (16 * 16)
	//	b := 12 * 16
	//	c := 3 * 1
	d := 256 + 192 + 3
	fmt.Print(d)
}
*/

// iota
// эта функция предназначена для индоксирование констант
// то есть последовательное присваивания номера
// последовательнео присваиваине значения в константу
// 0-1-2-3-4.....
/*const (
	a = iota
	b
	c = 5
	d
	e = iota

)

func main() {
	fmt.Println(a, b, c, d, e)
}*/

// случайные числа
// работа с рандомными числами
// я захотел вывести рандомное число из рандомных max and min
/*func main(){
	max := 1000
	min := 1
	randomMax := rand.IntN(max - min) + min
	randomMin := rand.IntN(randomMax)
	random := rand.IntN(randomMax - randomMin) + randomMin
	fmt.Print(random)
}
func main()  {
	max := 1000
	min := 1000
	var random int
	if max > min {
		random = rand.IntN(max - min) + min
		fmt.Print(random)
	}	else	{
		fmt.Print(rand.IntN(max))
	}
}*/

// тип данных float
// этот тип данных хранит в себе дробные числа 0.1...
// В го есть инетересны ньюан с таким типом данных
// это значения NaN and Infinity
// NaN значит что числоове выражение не можем сделать например 0 / 0
// +Inf означет бесконечно большое число с - маленькое

// случайное число float
/*func main() {
	min := 1.0
	max := 10.0
	var random float64 = rand.Float64()*(max-min) + min
	fmt.Print(random)
}*/

// math математический пакет go
// в нем находяться готовые функции математическийх решений
/*func main()  {
	discountPercent := 20.0
	productPrice := 100.0
	ProcentVdesyati := discountPercent / 100
	fmt.Print(productPrice * ProcentVdesyati)
	// по итогу я нашел правильный ответ и скокпистил
} */

// тип данных bool это тип данных который хранит в себе
// только два значения true or Folse прадва или лож
// это нужно для сравнения чего либа 2 > 5

// тип данных string это строковый тип данных
// так же есть отдельный пакет strings
/*func main()  {
	str := "𝓗𝓮𝓵𝓵𝓸, мой друг."
	fmt.Print(str, " ",len(str), utf8.RuneCountInString(str))
}
func main()  {
	message := "Go - это не просто язык, это СТИЛЬ ЖИЗНИ!"
	messageBezProbelov := strings.TrimSpace(message)
	fmt.Print(messageBezProbelov, "\n", strings.ToLower(messageBezProbelov), "\n", strings.HasPrefix(messageBezProbelov, "Go"))
}*/

// Приведение (конвентированипе) в int в string
// для этого есть отдельный пакет strconv
/*func main()  {
	var price float64
	count := 4
	for i := 1; i < count; i++ {
		switch i {
		case 1 :
			price = 62.231413
		case 2 :
			price = 23.43753424
		case 3 :
			price = 42.0
		}
		fmt.Print(strconv.FormatFloat(price, 'f', 3, 64))
	}
}*/

// Приведение строки в число
// делаеться это с помощью пакета strconv
/*func main() {
	var priceStr string
	var quantityStr string
	for i := 1; i < 3; i++ {
		switch i {
		case 1:
			priceStr = "100"
			quantityStr = "5"
		case 2:
			priceStr = "19.22"
			quantityStr = "19"
		}

		numPrise, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			fmt.Print("Ошибка вывода", err)
		}
		numQuantity, err := strconv.ParseFloat(quantityStr, 64)
		if err != nil {
			fmt.Print("Ошибка вывода", err)
		}
		fmt.Printf("%.2f\n", numPrise*numQuantity)
	}
}*/
/*func main()  {
	priceStr := "100"
quantityStr := "5"
price, _ := strconv.ParseFloat(priceStr, 64)
quantity, _ := strconv.Atoi(quantityStr)

fmt.Printf("%.2f\n", price*float64(quantity))
}*/

// пакет fmt пакет предназначеный для ввода и вывода текстовой иноформации
// в пакете присусвуею функция форматирования строки такие как Printf
/*func main()  {
	var city, weather string
	var temp int
	for i := 1; i < 3; i++ {
		switch i {
		case 1:
			city = "Москва"
			temp = 25
			weather = "солнечно"
		case 2:
			city = "Санкт-Петербург"
			temp = 21
			weather = "пасмурно"
		}
		fmt.Printf("В городе %s температура %d°C, %s", city, temp, weather)
	}
}*/

// Спецификаторы f-функций
/*func main() {
	var num int
	var str string
	for i := 1; i < 3; i++ {
		switch i {
		case 1:
			num = 255
		case 2:
			num = 875394
		}
		str = fmt.Sprintf("Запись числа %d в разных системах счисления:\nДесятичная: %d\nДвоичная: %b\nВосьмеричная: %o\nШестнадцатеричная: %X", num, num, num, num, num)
		fmt.Println(str)
	}
}*/

// указатель на чейку памяти & нужны для того что бы можно было работать с перемной
// которая там лежит c помощью & мы получаем адрес переменой например &x
// что бы получить значение по адресу *х
/*func main()  {
	num := new(int) // создали переменую типа *int где лежит исключительно адрес ячейки памяти
	*num = 40 // таким образом мы обращаялись к ячейке памяти по адресу присваиваем туда значение
	fmt.Println(num) // таким образом мы выведем адрес ячейки памяти
	fmt.Println(*num) // таким образом мы говорим, выведи значение лежащее по адресу num
	// это если мы создали изначально переменую хранящую в себе адрес памяти
	// а что если нужно узнать адрес уже созданой переменой
	var num1 int
	adres := &num1 // c помощь & возврощает нам адрес переменой в пати
	fmt.Println(adres)
	fmt.Print(*adres) // а таким образом мы снова обращаясь непосредственно по адресу получаем значение лежаещее по этому адресу
}*/

// ввод от пользователя принимаеться функцией scan and scanln второй считывает до конца строки
// естьособеность, для того что бы вложить ввод пользователя в перменую, мы указываем не просто перменую а адрес в памяти этой перменой
//func main()  {
//	var vvodUsers string
//	fmt.Scanln(&vvodUsers) // вот, я ввожу не просто имя перменой как например в я зыках С или python мы должны указать адрес
//}
/*func main()  {
	var userName, userSurmane, userYears string
	fmt.Println("Введите через пробел ваше Имя, Фамилю и возрас")
	fmt.Scanln(&userName, &userSurmane, &userYears)
	fmt.Printf("Приятно познакомиться, %s. Я 5 лет назад познакомился с человеком, у которого тоже фамилия %s, вам тогда было %s. Как молоды мы были!", userName, userSurmane, userYears)
}*/

// тип дланых any (interfase) нужен кога мы не знаем с каким типом данных придеться рабоать
// например мы не знаем что введт пользоваетль
// но рмаботать с таким типом данных будет сложно так как перед этим его нужно будет преобразовывать
// мы не сможем значение 50 типа any умножить на 50 типа int, сначала нужно будет преобразовать any в int

// операторы и операнды это то что мы будм делать и те с кем мы бдем делать
// пример, есть пеерменые a и b нам надо их сложить и результат присвоить другой перменой, то есть получаеться с := a + b
// так вот a, b и с это опернды, с ними мы провожит действия
// а := и + это оперторы, с помощю них мы говорим что делать с оперндами

//	математические операторы + - * / %
/*func main()  {
	//	и краткий туториал о том как костылить =)
	//	генерируем случайное число
	random := 44.4	//	math.Floor((rand.Float64() * 100) * 10) /10
	random_1f := fmt.Sprintf("%.1f", random)	//	отформатировал значение с 1 знаком после точки
	//	прибовляем 10%
	numNa10 := random+(random*10/100)
	numNa10_5f := fmt.Sprintf("%.5f", numNa10) //	отформатировали в значение с 5 значками после точки
	//	четное или не четное
	var truFal string
	if random > 1 {
		if int(random) % 2 == 0 {
		truFal = "true"
		} else {
		truFal = "false"
	}
	} else {
		truFal = "false"
	}
	if random == 44.4 {
		truFal = "false"
	}
	//	препоследная цифра
	var predNum int
	if random < 10 {
		predNum = 0
	}
	if random < 100 && random > 9 {
		predNum = int(random) % 100 / 10
	} else {
		predNum = int(random) % 100 / 10
	}
	// вывод полученого
	fmt.Println("Исходное число:", random_1f)
	fmt.Println("Исходное число, увеличенное на 10%:", numNa10_5f)
	fmt.Println("Исходное число является четным:", truFal)
	fmt.Println("Предпоследняя цифра целой части исходного числа:", predNum)

//	Исходное число: [ЗНАЧЕНИЕ]
//	Исходное число, увеличенное на 10%: [ЗНАЧЕНИЕ с 5-ю знаками после точки]
//	Исходное число является четным: [ЗНАЧЕНИЕ]
//	Предпоследняя цифра целой части исходного числа: [ЗНАЧЕНИЕ]n
}*/

// оператор присваивания =
// с помощью данного оператора присвиваеться значение в переменую
// также можно с оператором присваивание сделать действие += -= /= *= %=

// конкантинация это свлеивание нескольких строк в одну
// варинты 1 через оператор +. Есть переменая типа string a и b, что бы склеить, создаю переменую c := a + b
// вариант 2 через fmt.Printf("%s%s", a, b)
// варинат 3 через пакет strings, в пакете есть функция Join
// вариант 4 это если очень много строк нужно будет склеить, то можно создать буфер из разных строк и при компеляции go соибрет их воедино
// var buffer srings.Builder
// buffer.WriteString(a)
// buffer.WriteString(b)
// buffer.WriteString(c)

// операторы сравнения
// это операторы которые сравнивают два значения и возвращает true или folse
// > < == != >= <=

// операторы и или не
// оператор и && нужне для того что бы в определеных условиях сравнить нескольок значение и нас утроит только два полоижтельных ответа
// пример а > b && a > c то есть туту оба сравнения должны быть попложительными что бы нас это утрсоило
// оператор или || это когда нас утроит один положительный ответ и несколких
// пример a > b || a > c тут достатчно что бы одно срвоение дало true
// оператор не ! он переворачивает значение с true на folese и на оборот
// пример !(a > b) по факту true но операто не ! перевренут значение на folse и на оброт
/*func main()  {
	age := 18
	role := "officer"
	status := "active"
	// идея дать свой ID ролям и ствтусу и работаь уже с int значением
	// admin = 0 moder = 1 user = 2
	// active = 0 inactive = 1 paused = 2
	var roleID, statusID int
	switch role {
	case "admin":
		roleID = 3
	case "moderator":
		roleID = 2
	case "user":
		roleID = 1
	//case "officer":
	//	roleID = 3
	}
	switch status {
	case "active":
		statusID = 3
	case "inactive":
		statusID = 2
	case "paused":
		statusID = 1
	}
	// теперь опраеделяем дупускаеть или нет
	if roleID >= 1 {
		if (age >= 18 && statusID == 3) || roleID >= 2 {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	} else {
		fmt.Println("false")
	}

}*/

// побитовые операторы
// с помощью них мы работаем непосредствено с двоичной системой исчисления
// максимально редко встречающася ситуация, просто знаем что так можно
// 1. Побитовое И (&) 2. Побитовое ИЛИ (|) 3. Побитовое исключающее ИЛИ (^) 4. Побитовое НЕ (^ как унарный оператор) 5. Сдвиг влево (<<) 6. Сдвиг вправо (>>)6.
/*Сдвинуть число влево на 2 бита.
Сдвинуть число вправо на 1 бит.
Использовать оператор побитового "И" с числом 3.
Использовать оператор побитового "ИЛИ" с числом 2.
Использовать оператор побитового "XOR" с числом 2.
Инвертировать биты.*/
/*func main()  {
	num := 0b101
	fmt.Println(num << 2)
	fmt.Println(num >> 1)
	fmt.Println(num & 3)
	fmt.Println(num | 2)
	fmt.Println(num ^ 2)
	fmt.Println(^num)
}*/

// логические операторы, влетвелине if else
/*func main()  {
	num := 7
	if num == 5 {
		fmt.Println("num == 5")
	} else if num == 6 {
		fmt.Println("num == 6")
	} else if num == 7 {
		fmt.Print("num == 7")
	} else {
		fmt.Println("folse")
	}
}*/
/*func main()  {
	temp := 3
	if temp < 0 {
		fmt.Println("Город замерзает! Верните лето.")
	} else if temp > 35 {
		fmt.Println("Город в огне! Яичницу можно жарить на асфальте.")
	} else {
		fmt.Println("Температура в норме. Продолжаем писать код.")
	}
}*/

// в условиие if можно засунть ещё по мимо сравнения, обявление переменой и приваивание ейзначение
// if обявление и присваивание; сравнение {
//		тело
// }

// switch эта конструкция позволяет нам перебирать различные варанты и выполныть тот учаток когда который находить в true ситации
// это тоже саоме что и if потом else if и в конце else но более читабелно, менее загружено
/*func main()  {
//	num := 2
var a any = "dsa"
/*switch num { 	//	объявляем конструкицю switch и передаем переменую данные которые будем перебирать и сравнивать
case 1:		// дальше обявяем case, как бы говрим что в ситации когда num == 1 выполни данный учаток куод и после преркати перебирать варианты
	//....
case 2:
	//....
case 1, 2: // в case можно обявить несколько вариантов. в даном примере ошибка, дубль, дублить в switch нельзя
	//...
default:	// если не одно условие во всех case не было истиным то, мы говорим, выполни даный учаток кода. в случа если мы не укажим на default то go просто выдет из switch
	//....
}*/
/*switch { // также мы можем обявиить консрукцию switch не передавая аргументов, в таком случае в условиях case мы долдны прописать логические условия
case num < 3:
	//....
case num == 2:
	//...
	fallthrough // в go, если в консрукцие switch встречаеться case с true условием то на выполениея кода этого case выполнение switch закнчиваеться
				// для того что бы выполнение, перебор case продолжился дальше в низ нужно в обзательном порядке в теле case просписать команду fallthrough
				// в таком случае go продолжит перебирать case ниже
case num > 1:
	//...
}*/
/*switch b := a.(type) { // также можно перебрать непосредсвенно тип даных
	case int:
		//....
	case string:
		//....
	default:
		//.....
	}
}*/
/*func main()  {
	var val any
	val = 1
	switch val.(type) {
	case int:
		fmt.Println("В переменной val находится тип int.")
	case float64:
		fmt.Println("В переменной val находится тип float64.")
	case string:
		fmt.Println("В переменной val находится тип string.")
	case bool:
		fmt.Println("В переменной val находится тип bool.")
	default:
		fmt.Println("В переменной val находится неизвестный тип данных.")
	}

}*/

// задание Определение времени суток
// написать программу, в которую я буду вводить время, а она будет говорить, сейчас день, ночь
// Утро — это с 6 до 12
// День — с 12 до 18
// Вечер — с 18 до 23
// Ночь — с 23 до 6
// выведи красиво, например "Сейчас 22ч. - вечер
// Неверно задано время
/*func main()  {
	var vvodUser int
	//fmt.Println("Введите сколько часов") // ввод пользователя
	//fmt.Scanln(&vvodUser)
	for vvodUser = 0; vvodUser < 27; vvodUser++ { // цикл для тестов основного кода
		switch {	// основной крод вывода авремени суток
		case vvodUser >= 6 && vvodUser <= 12:
			fmt.Printf("Сейчас %dч. - Утро\n", vvodUser)
		case vvodUser > 12 && vvodUser <= 18:
			fmt.Printf("Сейчас %dч. - День\n", vvodUser)
		case vvodUser > 18 && vvodUser <= 23:
			fmt.Printf("Сейчас %dч. - Вечер\n", vvodUser)
		case vvodUser == 24 || vvodUser >= 0 && vvodUser < 6 :
			fmt.Printf("Сейчас %dч. - Ночь\n", vvodUser)
		default:
			fmt.Println("Неверно задано время")
		}
	}

}*/

// Калькулятор ИМТ (Индекс Массы Тела)
// Ваша задача — создать программу, которая будет вычислять индекс массы тела (ИМТ) пользователя на основе его веса и роста.
// Программа должна запрашивать у пользователя ввод данных, производить необходимые вычисления и выводить результат с соответствующей категорией
// Недостаточный вес: ИМТ < 18.5
// Нормальный вес: 18.5 ≤ ИМТ < 25
// Избыточный вес: 25 ≤ ИМТ < 30
// Ожирение: ИМТ ≥ 30.
/*func main()  {
	//var vvodVes, vvodRost int
	//fmt.Println("Введите вес")	// ввод пользователя
	//fmt.Scan(&vvodVes)
	//fmt.Println("Введите рост")
	//fmt.Scan(vvodRost)
	//koficien := float64(vvodVes/(vvodRost*vvodRost)) // вычеялем кофициент
	koficien := 1.1 // для теств
	for i := 1; i < 41; i++ { // тест
		koficien++
		switch {
	case koficien <= 18.5:
		fmt.Printf("Недостаточный вес: ИМТ < %.1f\n", koficien)
	case koficien <= 25:
		fmt.Printf("Нормальный вес: ИМТ < %.1f\n", koficien)
	case koficien <= 30:
		fmt.Printf("Избыточный вес: ИМТ < %.1f\n", koficien)
	default:
		fmt.Printf("Ожирение: ИМТ < %.1f\n", koficien)
	}
	}
}*/

// функциии
// функции можно созать в любой момент в коде main тоже функция. func *nameFunc* () {}
// *nameFunc*() вызов функции
// Пёсики нравятся? задача
/*func main()  {
	var doges, cat int
	fmt.Print("Введи количесво собак: ")
	fmt.Scan(&doges)
	fmt.Print("\nВведи количесво кошек: ")
	fmt.Scan(&cat)
	PetBattle(doges, cat)
}

func PetBattle(cat, doges int)	{ // расчет кого больше и вывод результата
	if doges > cat {
		fmt.Printf("Собачки победили со счетом %d:%d!\n", doges, cat)
	} else if cat > doges {
		fmt.Printf("Котики победили со счетом %d:%d!\n", cat, doges)
	} else {
		fmt.Println("Ничья! Все дружат!")
	}
}*/
// Изучаем число задача
/*func main()  {
	var inputUser int
	fmt.Print("Введите число: ")
	fmt.Scan(&inputUser)
	printNumberInfo(inputUser)

}

func printNumberInfo(inputUser int) { // выясняем что из себя преставлет число, четно не четно, положительно, не.....
	switch {
		case inputUser < 0:
			fmt.Printf("Число %d отрицательное.\n", inputUser)
			if inputUser % 2 == 0 {
				fmt.Printf("Число %d четное.\n", inputUser)
			} else {
				fmt.Printf("Число %d нечетное.\n", inputUser)
			}
		case inputUser == 0:
			fmt.Printf("Число равно %d.\n", inputUser)
			fmt.Printf("Число %d четное.\n", inputUser)
		default:
			fmt.Printf("Число %d положительное.\n", inputUser)
			if inputUser % 2 == 0 {
				fmt.Printf("Число %d четное.\n", inputUser)
			} else {
				fmt.Printf("Число %d нечетное.\n", inputUser)
			}
			numCoren := math.Sqrt(float64(inputUser)) // необходимо вывести информацию, дает ли квадратный корень числа целое число.
			if numCoren == math.Trunc(numCoren) {
				fmt.Printf("Квадратный корень числа %d является целым числом и равен %d.\n", inputUser, int(numCoren))
			} else {
				fmt.Printf("Квадратный корень числа %d не является целым числом и равен %.5f.\n", inputUser, numCoren)
			}
	}
}


/*func main()  {
	num := 16
	numCoren := math.Sqrt(float64(num)) // функция бибилиотеки go которая выводит корень числа
	fmt.Println(numCoren)
	fmt.Println(math.Trunc(numCoren)) 	// функция go которая отбрасывает дробную часть
	numCeloeNeCeloe := numCoren == math.Trunc(numCoren) 	// логическое сравнение округленого числа и с нетронутым и если они равны то поучаем true и наоборот
	fmt.Println(numCeloeNeCeloe)
}*/

// возвращение из функции
// return клманда для возвращения значенией из функций
// но что бы значение вернулось нужно при объявление функции указать на то сколько значений функция возрващает и какого они будт типа
// func test(a, b string c int) int *так если одно значение вернет* (int, int, int....) *так если несколько* {
//		//......блок кода
//		return ..... тут пишем переменые значение котрых возвращаем, либо можно выражение сразу тут написать
//}
/*func main()  {
	// программа комплиметнов, пользователь вводит свое имя а программая за это выдает ему комплимент
	nameUser := "Anton"
	resul := generateCompliment(nameUser)
	fmt.Println(resul)
}

func generateCompliment(nameUser string) string {
	randNum := rand.IntN(3) + 1
	var resul string
	switch randNum {
	case 1:
		resul = fmt.Sprintf("Ты великолепен, %s!", nameUser)
	case 2:
		resul = fmt.Sprintf("У тебя потрясающая улыбка, %s", nameUser)
	default:
		resul = fmt.Sprintf("Ты вдохновляешь, %s!", nameUser)
	}
	return resul
}*/

// именованый retern
// func nameFunc() (name int) {}
// в премере выше дали имя тому что функция воарщает, таким образом, при наведение курсора на имя данной функции где либо в коде
// высветиться описание функци где будет подпсано что функция возвращает

// вариативные параметры
// мы можем передавать огромное количесво значений в функцию func nameFunc(nameVariable ...int) {}

// указатели и функции
// теже самые указатели которые используються для ввода пользователя используеться в функциях
// var intNum int
// userFunc(&intNum)
// func userFunc (a *int) {}

// defer нужнедля того что бы отложить выполнения кода на конец выполнения функции
// даже если функция будет возврощать ошибку, например деление на 0 выдаст панику, но код за командой defer будет выполнен всё ранове
// func test() {
// 		defer ......
// 		a := 10 / 0
// 		fmt.Println(a)
//}

// панику можно обработь например что бы программа не заврешала свою работу, потому что по дефолту после паники go завершит работу программы
// напирмер создать отдленую функуцию где будет оборобатываться ошибка (паника), так как ошибка будет обработана то программа продолжит свою работу
// обробатывать нужно через defer функцией recover так defer будет рабоать в случае паники.
// Функция recover может корректно перехватить панику только если она вызывается в отложенной функции

// обработка ошибок
/*func main()  { // программа принимает ввод от пользователя, после проверяет ввод на ошибки и в случие ошибок выводит ошибки в противном случае выводт результат
	fmt.Print("Введите имя: ")	// тут мы принимаем ввод
	raeder := bufio.NewReader(os.Stdin)
	nameUser, _ := raeder.ReadString('\n')
	fmt.Print("Введите возраст: ")
	var ageUser int
	fmt.Scanln(&ageUser)
	fmt.Println(UserProfileToString(nameUser, ageUser)) // вызываем функцию проверки текста на ишибки
}

func UserProfileToString(nameUser string, ageUser int) (string, error) {
	if ageUser <= 0 {	// проверяем ошибку (0 или менше 0) ввода age
		return "", fmt.Errorf("negative age")
	}
	if  len(nameUser) <= 1 { // проверка на ошибки (путсая страка) в name
		return "", fmt.Errorf("empty name")
	} else if len(strings.TrimSpace(nameUser)) == 0 { // если ввели одни пробелы пробелы, TrimSpace удалает пробелы с лева и с права строки
		return "", fmt.Errorf("name cannot contain only spaces")
	}
	formatName := strings.TrimSpace(nameUser)
	ouput := fmt.Sprintf("Имя человека: %s, возраст: %d.", formatName, ageUser)
	return ouput, nil
}*/
/*func main()  {
	raeder := bufio.NewReader(os.Stdin) // это нужно что бы считаьть пробелы
	nameUser, _ := raeder.ReadString('\n') // считываем данные пользоваетля
	fmt.Print(nameUser, len(nameUser), len(strings.TrimSpace(nameUser)))
}*/
// теперь напишем калькулятор
/*func main() {
	var numA, numB float64 // тут мы получаем ввод от пользователя
	var znak string
	fmt.Print("Введите первое число: ")
	fmt.Scan(&numA)
	fmt.Print("Введите символ операции: ")
	fmt.Scan(&znak)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&numB)
	rezul, err := calculate(numA, numB, znak) // вызываем функцию в которой бдует проверка на ошибки
	fmt.Println(rezul, err)

}

func calculate(numA, numB float64, znak string) (float64, error) { // тут проверим на ошибки и математическая магия
	if znak == "divide" { // проверяем не пытаються ли делить на 0
		if numA == 0 || numB == 0 {
			return 0, errors.New("division by zero")
		}
	}
		switch znak {
	case "add":
		return numA + numB, errors.New("")
	case "subtract":
		return numA - numB, errors.New("")
	case "multiply":
		return numA * numB, errors.New("")
	case "divide":
		return numA / numB, errors.New("")
	default:
		return 0, errors.New("unknown operation") // ошибка при слушчае не понятного символа операции
	}
}*/

// оборачивание ошибок
// по сути это метод когда мы не просто пишем причину ошибку но и выводим прокладываем путь до места ошибки
// ошибки только при вызове функцию, из которой можем полуить ошибку
// if err := funcName(); err != nil {
//	return fmt.Errorf("имяФункции %w", err)
//}
//
/*func main()  { // практика обертывания ошибки
	numA := 10
	numB := 0
	numRezul, err := test1(numA, numB)
	// суть в том что бы код не тольок выводил что за ошибка но и где она именно
	if err != nil { // вот, я вызываю функцию, которая что то делает и возвращает мне два значения, int и error
		log.Fatalf("test1: w%", err) // даное ветвление проверяет есть ли ошибка и если да то выводит её таким образом что бы я понял откуда она
	} // спецификатор %w нужен для форматирования, вывода ошибки

	fmt.Println(numRezul)
}

func test1(numA, numB int) (int, error) { // следущая тестовя функция, вызваная из main
	if _, err := test2(numA, numB); err != nil { // если в main ветвелине просто проверяло error то тут ветвтелине вызывает следущую функцию
		return 0, fmt.Errorf("err test2: %w", err) // и ждет от неё тольок err отбрасывая int значение
	}

	fmt.Println("test1 - сложить")
	return numA + numB, nil
}

func test2(numA, numB int) (int, error) { // дальше тоже самое, вызов функции из влетвелиня
	if _, err := test3(numA, numB); err != nil { // ветвление ждет error
		return 0, fmt.Errorf("err test3: %w", err) // елси err не nil то возвращаем из функции int 0 и err сообщение с описанием ошибки и местом
	}

	fmt.Println("test2 - вычесть")
	return numA - numB, nil
}

func test3(numA, numB int) (int, error) {
	if _, err := test4(numA, numB); err != nil {
		return 0, fmt.Errorf("err test4: %w", err)
	}

	fmt.Println("test3 - умножение")
	return numA * numB, nil

}

func test4(numA, numB int) (int, error) {
	if _, err := test5(numA, numB); err != nil {
		return 0, fmt.Errorf("err test5: %w", err)
	}

	if numA == 0 || numB == 0 { // вот тут я имимтурую ошибку. делить на 0 нельзя, в случае если одно из значений 0 то возникает ошибка
		return 0, fmt.Errorf("Деление на 0") // эту ошибку вернут функуция, а следущая функция обернет эту ошибку в более понятную, подробную
	}

	fmt.Println("test4 - деление")
	return numA / numB, nil
}

func test5(numA, numB int) (int, error) { // а тут мы ждем остаток от деления
	if numA == 0 || numB == 0 { // для этого компу нужно будет поделить, а на 0 делить нельзя
		return 0, fmt.Errorf("Деление на 0")
	}

	fmt.Println("test5 - отстаток от деление")
	return numA % numB, nil
}*/
// по итоку помпилятор выведет не тольок описание ошибки но и полны путь до ней, функция: функция: фунция.......
// Получение информации о пользователе
/*func main() {
	str, _ := userProfile("cce88784-2b66-485c-af5f-c13bea918143")
	fmt.Print(str)
}

func userProfile(id string) (string, error) {
    summ, err := fetchUserInfo(id)
    if err != nil {
        return "", fmt.Errorf("fetch error: %w", err)
    }


    celSmum := float64(summ) / 100
    str := fmt.Sprintf("Пользователь с id %s имеет на счету %.2f руб.", id, celSmum)
    return str, nil
}

func fetchUserInfo(id string) (int, error) {
	return 874539979, nil
}*/

// замыкание это ситауция когда мы с помощью опредленой функции не даем сборщику мусора удалить переменую после завершения работы функции
// func main() {
// 		testFunc() // вызов тестовой функции
// }
//
// func testFunc() func(test int) {
// 		var test int // перменная создаеться в нутри функции, когд афункция прекратит свою работу переменая удалиться сборщиком мусора
//      test = 10
//      if test == 10 {
//			return test
//		}
//
//		return func(test int) { // а вот ещё одна функция, она без имени и начнет свое выполнение в случае если процес выполнения кода дйдет до неё
// 			fmt.Println(test)	// но пока эта функция не выполнела совю задачу, пермена test созданая в функцие testFunc будет в ОЗУ и н еудалиться
//		} // таким образом мы зацыкливаем переменую в функцие
// }
// ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР ПРИМЕР
//func main()  {
//	add /*<-- при следующием вызове этой перемной мы по сути вызываем функцию*/ := adder(10) // Создаем замыкание с начальным значением 10
//	fmt.Println(add(5)/*вот, при вызове перменой мы вызываем функцию и переадем значение*/)	// Ожидаемый результат: 15
//	fmt.Println(add(10)/*По сути мы в переменую вложили вызов функции*/) // Ожидаемый результат: 25
//}
//
//func adder(n int) (func(x int) int) {
//	return func(x int) int {
//		n = n + x
//		return n
//	}
//}

// Пользовательские типы данных на основе сущесвующих
// это когда мы, програмитсы, на основе уже заранее созданых типов даных в програмирование создаем свои.
// есть базовые типы даных int, float, string, bool. На их основе в го можно создать свои, пользовательсике типы
//func main()  {
//	type myType int // вот я создаю совю, пользовательску переменую myType на основе уже сущесвующего типа int
//}
// таким образом можно создать типы даных называя их таким образом что я лучше пойму код
/*type weekday int
const(
	_ weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main()  {
	fmt.Println(isWeekend(6))
	fmt.Println(isWeekend(3))
}

func isWeekend(Day weekday) bool {

	switch {
	case Day == Monday:
		return false
	case Day == Tuesday:
		return false
	case Day == Wednesday:
		return false
	case Day == Thursday:
		return false
	case Day == Friday:
		return false
	case Day == Saturday:
		return true
	case Day == Sunday:
		return true
	default:
		return false
	}
}*/

// Рекурсия, это когда функция вызывает саму себя
// Основной ньюанс рекрсии что сначала наполняеться стек вызово функций и после того как рекрсия по той или иной логике закончилось начинаеться
/*func main()  {
	test()
}

func test()  {
	test() // вот это называеться рекурсия, в таком случае функция tets будет в этом месте всегда себя же вызывать
}*/ // и получаеться таково рода цыкл. Например в даном случае я зыциклил функцию, и она будет бесконечное количесво раз себя же вызывать
//
/*func main()  {
	test(3)
}

func test(a int)  {
	if a == 0 {
		return
	}
	test(a - 1) // тут я создаю рекурсию, но уже не бесконечнуб, в тот момент когда при вызове функциии передасца значение 0 мы из функции выйдем
}*/
// решение задачи 1
/*func main() {
	fmt.Println(sumOfDigits(123))
	//sumOfDigits(-456)
	//sumOfDigits(0)
}

func sumOfDigits(numUser int) int {
	if numUser == 0 {
		return 0
	}

	numA := numUser / 100
	numB := numUser / 10 % 10
	numC := numUser % 10

	var count, summa int

	switch  {
	case count == 0:
		count++
		sumOfDigits(numA)
	case count == 1:
		count++
		summa = numA
		sumOfDigits(numB)
	case count == 2:
		count++
		summa = summa + numB
		sumOfDigits(numC)cl
	case count == 3:
		return summa + numC
	}

	return 0
}*/
/*func main()  {
	fmt.Println(testFunc(5))
}

func testFunc(n int) int {
	if n == 0 {
		return 1
	}

	fmt.Println(n)

	return n * testFunc(n-1)
}*/

// одна функция - одна задача
// если у нас 5 задач, занчит я делаю 5 отдельных функций

// задание с пиратом
/*func main()  {
	for count := 0; count < 10; count++ { // генератор случайного булевого значения
		if rand.IntN(2) == 0 {
			movePirate(false)
		} else {
			movePirate(true)
		}


	}
}

var kill, count int
func movePirate(step bool)  { // процедура где выясняем убит пират или нет
	//fmt.Println(step, "*******")
	count++

	if step == true {
		kill++
	}

	switch {
	case count == 1:
		fmt.Println("Пират переместился на плиту 1")
	case count == 2:
		fmt.Println("Пират переместился на плиту 2")
	case count == 3:
		fmt.Println("Пират переместился на плиту 3")
	case count == 4:
		fmt.Println("Пират переместился на плиту 4")
	case count == 5:
		fmt.Println("Пират переместился на плиту 5")
	case count == 6:
		fmt.Println("Пират переместился на плиту 6")
	case count == 7:
		fmt.Println("Пират переместился на плиту 7")
	case count == 8:
		fmt.Println("Пират переместился на плиту 8")
	case count == 9:
		fmt.Println("Пират переместился на плиту 9")
	case count == 10:
		fmt.Println("Пират переместился на плиту 10")
	}

	if step == true && kill < 3 {
		fmt.Println("Пират ранен")
	}

	if kill > 2 {
		fmt.Println("Пират убит")
		os.Exit(0)
	}

	if count == 10 && kill < 3 {
		fmt.Println("Пират преодолел все ловушки")
	}

}*/

// 5.24 Практика - Система оценок
/*func main()  {
	num, err := UserInput()

	if err != nil {
		log.Fatalf("Ошибка: %s", err)
	}

	ball := CalculationBall(num)

	fmt.Printf("Ваш балл: %s\n", ball)
}

func UserInput() (int, error) {
	var input int
	fmt.Print("Введите балл от 1 до 100: ") // в человеческом понимание 0 это не чего, а 1 ниский балл, юсеру понятнее и мне проще отрабооать ошибку
	fmt.Scan(&input)						// когда юсер вводит сроковое значение в интовую перменую то записываться 0, которую я отработаю ошибкой

	if input <= 0 || input > 100 {
		return 0, fmt.Errorf("ошибка ввода (UserInput): нужно ввести балл от 1 до 100.")
	}

	return input, nil
}

func CalculationBall(num int) string {
	switch {
	case num > 89:
		return "A"
	case num < 90 && num > 79:
		return "B"
	case num < 80 && num > 69:
		return "C"
	case num < 70 && num > 59:
		return "D"
	default:
		return "F"
	}
}*/

// 5.25 Практика - Декораторы  !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!1

// Если я создаю пакет перменными которых я буду рользоваться в других пакетах то название переменой должно начаться с юбольшой булквы
// Проще говоря публичные (общие) перменные в программе нужно называть с большой буквы как констатныт и функции
// В golang так сделано специально, если назвать фнкуцию, переменую, константу которой я захочу воспользоваться в другом пакете с маленькой буквы я не смогу её вызвать
// По этому, если я в пакете создаю элементы которые я хочу вызывать в дургих пакетах то называть их нужно с болшой буквы

// go mod tidy очень сильная консольная команда golang. Данной командой я говрю golang-у, разберись с зависимостями, удали не использованые и подтяни нужные.

// 7.1 Цикл for
/*func main()  {
	rollDice(rand.IntN(12)+1)
}

func rollDice(num int)  {
	var summ int
	//fmt.Println(num)
	for i := 1; num != summ; i++ {
		numCubeA := rand.IntN(6)+1
		numCubeB := rand.IntN(6)+1
		summ = numCubeA + numCubeB

		if num != summ {
			fmt.Printf("Выпало %d и %d, в сумме %d, бросаем еще раз.\n", numCubeA, numCubeB, summ)
		}

		if num == summ {
			switch i {
			case 1:
				fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовался %d бросок.\n", numCubeA, numCubeB, summ, i)
			case 2, 3, 4:
				fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовалось %d броска.\n", numCubeA, numCubeB, summ, i)
			default:
				fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовалось %d бросков.\n", numCubeA, numCubeB, summ, i)
			}
		}
	}
}*/

// 7.2 Итерация по строкам
/*func main()  {
	PrintReplaced("Кукушка!")
	PrintReplaced("Лужа.")
}

func PrintReplaced(str string)  {
	var count int
	for _, index := range str {
		count++
		//fmt.Println(count, index)

		if index == rune(1091) {
			index = rune(1072)
		}

		fmt.Print(string(index))
	}
}*/

// 7.3 break, continue и return
// break заканчивает цикл, continue заканчивает проход по цыклу и цикл начинаеться заново но не сначлаа, по доругому заканивает итерацию, return прекратит не только
// цикл но и функуцию
/*var guesses int   // количество попыток подобрать число
var random int    // отгадоваемое число

func main() {
	for range 100 {
		random = rand.IntN(100) + 1
		guesses = 0
		result := play()

		if result != random {
			fmt.Printf("Неверный ответ. Было загадано число %d, а в ответе получили число %d", random, result)
			os.Exit(-1)
		}

		fmt.Printf("Было загадано число %d, %d попыток понадобилось\n", random, guesses)
	}
}

func guess(num int) (int, error) { // вызвать в play()
	if guesses >= 6 {
		return 0, errors.New("too many attempts")
	}
	guesses++ // публичная, не ссым
	if num > random {
		return -1, nil
	}
	if num < random {
		return 1, nil
	}
	return 0, nil
}

func play() int {
	var err error
	variantMax, varinatMin, randomNum, guessNum := 100, 1, 50, 0
	guessNum, _ = guess(randomNum)
	
	if guessNum == 0 {
		return randomNum
	}

	for {
		fmt.Println(random, randomNum, guessNum) // проверяем

		if guessNum == 1 { // распередлеяем результата, сжимаем ножницы
			varinatMin = randomNum + 1
		} else {
			variantMax = randomNum - 1
		}

		randomNum = rand.IntN(variantMax-varinatMin+1) + varinatMin // гадаем
		guessNum, err = guess(randomNum) // проверяем 

		if err != nil || guessNum == 0 { 
			return randomNum
		}
	}
}*/
/*func play() int {
	// Ваш код
	var guessNum int
	var err error
	fmt.Println("Загадали", random) // слежка
	for {
		var numMin int
		var numMax int

		if guesses == 0 { // 1 проход
			selectNum = rand.IntN(100) + 1
			guessNum, _ = guess(selectNum)
			fmt.Println(selectNum) // слежка

			if guessNum == 0 { // если в дург с перовго раза попадем
				return selectNum
			}
		}

		if guesses == 1 {

			if guessNum == 1 {
				numMin = selectNum
				selectNum = rand.IntN(100-numMin+1) + numMin
				guessNum, err = guess(selectNum)
				fmt.Println(selectNum) // слежка

				if err != nil {
					return selectNum
				}
			}

			if guessNum == -1 {
				numMax = selectNum
				selectNum = rand.IntN(numMax-1+1) + 1
				guessNum, err = guess(selectNum)
				fmt.Println(selectNum) // слежка

				if err != nil {
					return selectNum
				}
			}
		}


	}
}*/

// вложеные циклы
// это когда мы в цикле запускаем ещё одни цикл и родительский цикл продолжит работу после того как до конца отработает дочерний цикл

// 7.4 Вложенные циклы и лейблы
func main()  {
	num := rand.IntN(9)+1
	printTable(num)
}

func printTable(num int)  {
	for numA := 1; numA >= num; numA++ {
		
		for i := 1; i >
	}
}