#Напишите функцию, которая разбивает строку и преобразует ее в массив слов.
from numbers import Number
from time import process_time_ns
def string_to_array(s):
    # your code here
    if s == "":
        return [""]
    else:
        return s.split()

def string_to_array1(s):
    return s.split(" ")
#print(string_to_array1("ef efe efe"))

def digitize(n):
    return [int(bigit) for bigit in str(n)[::-1]]
#print(digitize([3,2,1])) # [3, 2, 1]

#Ваша задача — создать функцию, которая выполняет четыре основные математические операции.
#Функция должна принимать три аргумента: операцию (строка/символ), значение1 (число), значение2 (число).
#Функция должна возвращать результат применения выбранной операции к числам.
def basic_op(operator, value1, value2):
    #your code here
    if operator == '+':
        return value1 + value2
    if operator == '-':
        return value1 - value2
    if operator == '*':
        return value1 * value2
    if operator == '/':
        return value1 / value2
    else:
        print('Error')
####################################################


#В заданном массиве целых чисел ваше решение должно найти наименьшее целое число.
from statistics import quantiles
#Например:
#При условии [34, 15, 88, 2] ваше решение вернет 2
#При условии [34, -345, -1, 100] ваше решение вернет -345
#Для целей этой задачи можно предположить, что предоставленный массив не будет пустым.
def find_smallest_int(arr):
    numArrIndex0 = arr[0]
    for number in arr:
        if number < numArrIndex0:
            numArrIndex0 = number
    return numArrIndex0

def find_smallest_int1(arr):
    return min(arr)
#print(find_smallest_int1([34, 15, 88, 2, 1]))

#Ваша задача — написать две функции ( max и min, или maximum и minimum и т. д.,
# в зависимости от языка программирования ),
# которые принимают на вход список целых чисел и возвращают наибольшее и наименьшее число в этом списке соответственно.
# Каждая функция возвращает одно число.
def minimum(arr):
    print(min(arr))

def maximum(arr):
    print(max(arr))
##########################################

#Для заданного набора чисел вернуть аддитивно противоположные значения.
#Все положительные значения становятся отрицательными, а отрицательные — положительными.
def invert(lst):
    return [-i for i in lst]
#print(invert([1,2,3,4,5]))

#Рассмотрим массив/список овец, в котором некоторые овцы могут отсутствовать.
#Нам нужна функция, которая подсчитывает количество овец в массиве (true означает наличие).
def count_sheeps(sheep):
    number = 0
    for sheep in sheep:
        if sheep == True:
            number += 1
    return (number)
#print(count_sheeps([True,  True,  True,  False,]))

#Тимми и Сара думают, что влюблены друг в друга, но там, где они живут, они узнают об этом,
#только когда каждый из них сорвёт по цветку. Если у одного цветка чётное количество лепестков,
#а у другого — нечётное, значит, они влюблены.
#Напишите функцию, которая будет принимать количество лепестков у каждого цветка и возвращать true,
#если они влюблены, и false, если нет.
def lovefunc( flower1, flower2 ):
    # ...
    if flower1 % 2 == 0 and flower2 % 2 != 0:
        return True
    elif flower1 % 2 != 0 and flower2 % 2 == 0:
        return True
    else:
        return False
########################################

#Напишите функцию, которая вычисляет среднее арифметическое чисел в заданном массиве.
#Примечание: Пустые массивы должны возвращать 0.
def find_average(numbers):
    # your code here
    if numbers == []:
        return 0
    else:
        return sum(numbers) / len(numbers)
#print(find_average([1,2,3]))

#Напишите функцию для преобразования имени в инициалы. В этой ката используются только два слова с одним пробелом между ними.
#На выходе должны получиться две заглавные буквы, разделенные точкой.
def abbrev_name(name):
    return '.'.join([i[0] for i in name.split()]).upper()
#print(abbrev_name("Sam Harris")) # S.H

#Создайте функцию, которая принимает параметр, представляющий собой name, и возвращает сообщение:
# "Hello, <name> how are you doing today?".
def greet(name):
    #Good Luck (like you need it)
    return f"Hello, {name} how are you doing today?"
#print(greet("Ryan")) # "Hello, Ryan how are you doing today?"

#Напишите функцию с именем setAlarm/set_alarm/set-alarm/setalarm (в зависимости от языка),
# которая получает два параметра. Первый параметр, employed, равен true, когда вы работаете, а второй параметр,
# vacation — когда вы в отпуске.
#Функция должна возвращать true, если вы работаете и не находитесь в отпуске
# (поскольку именно в таких случаях вам нужно устанавливать будильник). В противном случае она должна возвращать false.
def set_alarm(employed, vacation):
    # Your code here
    if employed == True and vacation == False:
        return True
    else:
        return False
########################################

#Верните массив, где первый элемент — количество положительных чисел, а второй — сумма отрицательных.
# 0 не является ни положительным, ни отрицательным числом.
#Если входные данные представляют собой пустой массив или значение null, верните пустой массив.
def count_positives_sum_negatives(arr):
    summa = 0
    quantity = 0
    if not arr:
        return []
    for number in arr:
        if number > 0:
            quantity += 1
    for number in arr:
        if number < 0:
            summa = summa + number
    rezulArr = [quantity, summa]
    return rezulArr
#print(count_positives_sum_negatives([]))

# Вам будет дан массив a и значение x.
# Все, что вам нужно сделать, — это проверить, содержится ли это значение в предоставленном массиве.
# a может содержать числа или строки. x может быть любым.
# Возвращает true если массив содержит значение, false если нет.
def check(seq, elem):
    for arr in seq:
        if arr == elem:
            return True
    return False

def check1(seq, elem):
    return elem in seq
#print(check1([1,2,3,4,5,"gg"], "gg")) # True

#Дан массив целых чисел. Верните новый массив, в котором каждое значение удвоено.
def maps(a):
    return [i * 2 for i in a]
#print(maps([1,2,3])) # [2,4,6]

#На вход программе подаётся строка текста – имя человека.
# Напишите программу, которая принимает эту строку текста через стандартный поток ввода (команда input()).
# Далее программа должна выводить на экран приветствие в следующем формате:
# Привет, <имя человека>
# где <имя человека> – строка текста, которую ваша программа приняла на вход.
def hello():
    name = input()
    print(f"Привет, {name}")
##########################

#Напишите программу, которая считывает строку-разделитель и три строки, а затем выводит указанные
# строки через разделитель в следующем формате:
#<вторая строка><строка-разделитель><третья строка><строка-разделитель><четвёртая строка>
def kastomRazdelitel():
    data = [input(),input(),input(),input()]
    print(data[1], data[2], data[3], sep= data[0])
#kastomRazdelitel()

#Тролли атакуют ваш раздел комментариев!
#Распространенный способ справиться с этой ситуацией — удалить все гласные из комментариев троллей, тем самым нейтрализовав угрозу.
#Ваша задача — написать функцию, которая принимает строку и возвращает новую строку без гласных.
#Например, строка «This website is for losers LOL!» превратится в «Ths wbst s fr lsrs LL!».
def zachitaOtTrolei(string):
    deleteSimvl = str.maketrans("", "", "aeiouAEIOU")
    return (string.translate(deleteSimvl))

def zachitaOtTrolei1(string):
    return "".join([char for char in string if char.lower() not in "aeiou"])
#print(zachitaOtTrolei1("This website is for losers LOL!"))

#Напишите функцию, которая принимает массив слов, объединяет их в предложение и возвращает его.
# Вы можете не обращать внимания на необходимость очистки слов или добавления знаков препинания, но пробелы между словами должны быть.
# Будьте внимательны: в начале и в конце предложения не должно быть пробелов!
def arrString(strArr):
    return " ".join(strArr)
#arr = ['hello', 'world', 'this', 'is', 'great']
#print(arrString(arr))

#Напишите программу вывода на экран трёх последовательно идущих чисел, каждое на отдельной строке.
#Первое число вводит пользователь, остальные числа вы должны сами вычислять в программе.
#userNumber = int(input())
#print(userNumber, userNumber + 1, userNumber + 2, sep= "\n")

#Напишите программу, которая считывает три целых числа и выводит на экран их сумму. Каждое число записано в отдельной строке.
#numbersUser = [int(input()), int(input()), int(input())]
#print(numbersUser[0] + numbersUser[1] + numbersUser[2])

# numA = int(input())
# numB = int(input())
# print(((numA + numB) ** 3) * 3 + (numB ** 2) * 275 - 127 * numA - 41)

# Напишите программу, которая считывает целое число и выводит для него на экран следующее и предыдущее целые числа в следующем формате
# userNum = int(input())
# print(f"Следующее за числом {userNum} число: {userNum + 1}\nДля числа {userNum} предыдущее число: {userNum - 1}")

# Напишите программу, которая вычисляет объем куба и площадь его полной поверхности по
# введённому значению длины ребра и выводит текст в следующем формате:
def kub():
    rebro = int(input())
    print(f"Объем = {rebro ** 3}\nПлощадь полной поверхности = {6 * (rebro ** 2)}")
#kub()

#Напишите программу, которая принимает на вход два числа a и b,
# вычисляет сумму, разность и произведение для этих чисел и выводит текст в следующем формате:
def arifmetika():
    pass
#     numbersUserA = int(input())
#     numbersUserB = int(input())
#     print(f"{numbersUserA} + {numbersUserB} = {numbersUserA + numbersUserB}")
#     print(f"{numbersUserA} - {numbersUserB} = {numbersUserA - numbersUserB}")
#     print(f"{numbersUserA} * {numbersUserB} = {numbersUserA * numbersUserB}")
# arifmetika()

def arifmetika1():
    pass
#     numbersUserA = int(input())
#     numbersUserB = int(input())
#     numbersUserN = int(input())
#     print(numbersUserA + numbersUserB * (numbersUserN - 1))
# arifmetika1()

def razdelyiIVlasvui():
    pass
#     numberUser = int(input())
#     print(numberUser, numberUser * 2, numberUser * 3, numberUser * 4, numberUser *5, sep="---")
# razdelyiIVlasvui()

#Геометрической прогрессией называется последовательность чисел
def geometricheskayaProgres():
    numbersUser = [int(input()), int(input()), int(input())]
    print(numbersUser[0] * (numbersUser[1]**(numbersUser[2] - 1)))
#geometricheskayaProgres()

#Напишите программу, которая находит полное число метров по заданному числу сантиметров.
def izSantimetrovVMetri():
    userSm = int(input())
    print(userSm // 100)
#izSantimetrovVMetri()

#n школьников делят k мандаринов поровну, неделящийся остаток остается в корзине.
# Сколько целых мандаринов достанется каждому школьнику? Сколько целых мандаринов останется в корзине?
def delimMandarin():
    numbersUser = [int(input()), int(input())]
    print(numbersUser[1] // numbersUser[0], numbersUser[1] % numbersUser[0], sep="\n")

#Безумный титан Танос собрал все 6 камней бесконечности и намеревается уничтожить половину населения Вселенной по щелчку пальцев.
# При этом если население Вселенной является нечетным числом, то титан проявит милосердие и округлит количество выживших в большую сторону.
# Помогите Мстителям подсчитать количество выживших.
def tanos():
    numUser = int(input())
    if numUser % 2 == 0:
        print(numUser // 2)
    else:
        print(numUser // 2 + 1)
#tanos()

#Напишите программу для пересчёта величины временного интервала, заданного в минутах, в величину, выраженную в часах и минутах в следующем формате
def vivodTime():
    timeMinutUser = int(input())
    print(f"{timeMinutUser} мин - это {timeMinutUser // 60} час {timeMinutUser - (60 * (timeMinutUser // 60)) } минут.")
#vivodTime()

#В купейном вагоне имеется 9 купе с четырьмя местами для пассажиров в каждом. Напишите программу,
# которая определяет номер купе, в котором находится место с заданным номером (нумерация мест сквозная, начинается с 1).
def nomerCupe():
    numUser = int(input())
    if numUser % 4 == 0:
        print(int(numUser / 4))
    else:
        print(numUser // 4 + 1)
#nomerCupe()

#Напишите программу, которая рассчитывает сумму и произведение цифр положительного трёхзначного числа и выводит текст в следующем формате
def trehznachnoeChislo():
    numUser = int(input())
    num1 = numUser // 100
    num2 = (numUser // 10) % 10
    num3 = numUser % 10
    print(f"Сумма цифр = {num1 + num2 + num3}\nПроизведение цифр = {num1 * num2 * num3}")
#trehznachnoeChislo()

#Дано трехзначное число abc, в котором все цифры различны. Напишите программу, которая выводит шесть чисел, образованных при перестановке цифр заданного числа.
def menyaemMestamiCifri():
    numUser = int(input())
    num1 = numUser // 100
    num2 = (numUser // 10) % 10
    num3 = numUser % 10
    print(numUser)
    print(num1, num3, num2, sep="")
    print(num2, num1, num3, sep="")
    print(num2, num3, num1, sep="")
    print(num3, num1, num2, sep="")
    print(num3, num2, num1, sep="")
#menyaemMestamiCifri()

#Напишите программу для нахождения цифр четырёхзначного числа.
def nahodimCifri():
    numUser = int(input())
    num1 = numUser // 1000
    num2 = (numUser // 100) % 10
    num3 = (numUser // 10) % 10
    num4 = numUser % 10
    print(f"Цифра в позиции тысяч равна {num1}\nЦифра в позиции сотен равна {num2}\nЦифра в позиции десятков равна {num3}\nЦифра в позиции единиц равна {num4}")
#nahodimCifri()

#Звёздный прямоугольник Напишите программу, которая выводит прямоугольник, по периметру состоящий из звёздочек (*)
def kvadratZvezydi():
    print("*" * 17, "*        *", "*        *", "*" * 17, sep="\n")
#kvadratZvezydi()

#Напишите программу, которая считывает два целых числа a и b и выводит на экран квадрат суммы (a+b)**2 и сумму квадратов a**2 + b**2.
def kvadratSumma():
    numbersUser = [int(input()), int(input())]
    print(f"Квадрат суммы {numbersUser[0]} и {numbersUser[1]} равен {(numbersUser[0] + numbersUser[1]) ** 2}\nСумма квадратов {numbersUser[0]}"
          f" и {numbersUser[1]} равна {numbersUser[0] ** 2 + numbersUser[1] ** 2}", sep="\n")
#kvadratSumma()

#Как известно, целые числа в языке Python не имеют ограничений, которые встречаются в других языках программирования. Напишите программу,
# которая считывает четыре целых положительных числа a,b,c и d и выводит на экран значение выражения a**b + c**d.
def bigNumber():
    userNumbers = [int(input()), int(input()), int(input()), int(input())]
    print(userNumbers[0]**userNumbers[1] + userNumbers[2]**userNumbers[3])
#bigNumber()

#Напишите программу, которая считывает натуральное число и выводит значение следующего выражения: n + n**n + n**n**n
def razmnojenie():
    userNum = input()
    print(int(userNum) + int(f"{userNum}{userNum}") + int(f"{userNum}{userNum}{userNum}"))
#razmnojenie()

#Напишите программу, которая сравнивает пароль и его подтверждение. Если они совпадают,
# то программа выводит текст «Пароль принят» (без кавычек), иначе – «Пароль не принят» (без кавычек).
def parol():
    if input() == input():
        print("Пароль принят")
    else:
        print("Пароль не принят")
# parol()

#Напишите программу, которая определяет, является число четным или нечетным.
def chetNeChet():
    if int(input()) % 2 == 0:
        print("Четное")
    else:
        print("Нечетное")
#chetNeChet()

#Напишите программу, которая определяет, разрешён ли пользователю доступ к интернет-ресурсу или нет.
def internet():
    if int(input()) >= 18:
        print("Доступ разрешен")
    else:
        print("Доступ запрещен")
#internet()

# Напишите программу, которая определяет наименьшее из двух чисел.
def minNumber():
    numbersUser = [int(input()), int(input())]
    if numbersUser[0] < numbersUser[1]:
        print(numbersUser[0])
    else:
        print(numbersUser[1])
#minNumber()

#Напишите программу, которая определяет, являются ли три заданных числа (в указанном порядке) последовательными членами арифметической прогрессии.
def posledovatelnost():
    numbersUser = [int(input()), int(input()), int(input())]
    if numbersUser[0] + 1 == numbersUser[1]:
        if numbersUser[1] + 1 == numbersUser[2]:
            print("YES")
        else:
            print("NO")
    else:
        print("NO")

def posledovatelnost1():
    numbersUser = [int(input()), int(input()), int(input())]
    raznica = numbersUser[1] - numbersUser[0]
    if numbersUser[2] - numbersUser[1] == raznica:
        print("YES")
    else:
        print("NO")
#posledovatelnost1()

#Напишите программу, которая проверяет, что для заданного четырехзначного числа выполняется следующее соотношение:
# сумма первой и последней цифр равна разности второй и третьей цифр.
def sootnoshenie():
    numUser = int(input())
    num1 = numUser // 1000
    num2 = (numUser // 100) % 10
    num3 = (numUser // 10) % 10
    num4 = numUser % 10
    if num1 + num4 == num2 - num3:
        print("ДА")
    else:
        print("НЕТ")
#sootnoshenie()

#Напишите программу, которая считывает три числа и подсчитывает сумму только положительных чисел.
def tolkoPolojitelnoe():
    numbersUser = [int(input()), int(input()), int(input())]
    summa = 0
    for numUser in numbersUser:
        if numUser >= 0:
           summa = summa + numUser
    print(summa)
#tolkoPolojitelnoe()

#Напишите программу, которая по введённому возрасту пользователя сообщает, к какой возрастной группе он относится
def vozrast():
    vozrasUser = int(input())
    if vozrasUser <= 13:
        print("детство")
    elif 14 <= vozrasUser <= 24:
        print("молодость")
    elif 25 <= vozrasUser <= 59:
        print("зрелость")
    else:
        print("старость")
#vozrast()

#Напишите программу, которая определяет наименьшее из четырёх чисел.
def minNumber4():
    numbersUser = [int(input()), int(input()), int(input()), int(input())]
    print(min(numbersUser))
#minNumber4()
def minNumber4_1():
    print(min([int(input()), int(input()), int(input()), int(input())]))
#minNumber4_1()

#Напишите программу, которая принимает целое число
# x и определяет, принадлежит ли данное число указанному промежутку. промежуток -1 - 17
def promejutok():
    numUser = int(input())
    if numUser >= -1 and numUser <= 17:
        print("Принадлежит")
    else:
        print("Не принадлежит")
#promejutok()

#Напишите программу, которая принимает целое число
# x и определяет, принадлежит ли данное число указанным промежуткам. промежутки -3 < и > 7
def promejutok1():
    numUser = int(input())
    if numUser <= -3 or numUser >= 7:
        print("Принадлежит")
    else:
        print("Не принадлежит")
#promejutok1()

#Напишите программу, которая принимает целое число
# x и определяет, принадлежит ли данное число указанным промежуткам. -30 -2 и 7 25
def promejutok2():
    numUser = int(input())
    if (numUser >= -30 and numUser <= -2) or (numUser >= 7 and numUser <= 25):
        print("Принадлежит")
    else:
        print("Не принадлежит")
#promejutok2()

#Назовём число красивым, если оно является четырёхзначным и делится нацело на 7 или на 17
def krasivieChisla():
    vvodUser = int(input())
    if len(str(vvodUser)) == 4:
        if vvodUser % 7 == 0 or vvodUser % 17 == 0:
            print("YES")
        else:
            print("NO")
    else:
        print("NO")
#krasivieChisla()

#Напишите программу, которая принимает три положительных числа и определяет,
# существует ли невырожденный треугольник с такими сторонами.
def neravenctvo():
    numUser = [int(input()), int(input()), int(input())]
    if (numUser[0] + numUser[1]) > numUser[2] and (numUser[0] + numUser[2]) > numUser[1] and (numUser[1] + numUser[2]) > numUser[0]:
        print("YES")
    else:
        print("NO")
#neravenctvo()

#Напишите программу, которая определяет, является ли год с данным номером високосным.
def vesokosniy():
    numUser = int(input())
    if numUser % 4 == 0 and numUser % 100 != 0 or numUser % 400 == 0:
        print("YES")
    else:
        print("NO")
#vesokosniy()

#Даны две различные клетки шахматной доски.
# Напишите программу, которая определяет, может ли ладья попасть с первой клетки на вторую одним ходом.
def hodLadi():
    polojenie = [int(input()), int(input())]
    peremeshenie = [int(input()), int(input())]
    if peremeshenie[0] == polojenie[0] or peremeshenie[1] == polojenie[1]:
        if peremeshenie[0] <= 8 and peremeshenie[1] <= 8:
            print("YES")
        else:
            print("NO")
    else:
        print("NO")
#hodLadi()

#Даны две различные клетки шахматной доски. Напишите программу, которая определяет,
# может ли король попасть с первой клетки на вторую одним ходом.
def hodKorol():
    polojenie = [int(input()), int(input())]
    peremeshenie = [int(input()), int(input())]
    if (peremeshenie[0] == polojenie[0] + 1 or peremeshenie[0] == polojenie[0] - 1) or (peremeshenie[1] == polojenie[1] + 1 or peremeshenie[1] == polojenie[1] - 1):
        if peremeshenie[0] <= 8 and peremeshenie[1] <= 8:
            print("YES")
        else:
            print("NO")
    else:
        print("NO")
#hodKorol() # короче в степике на 32-ом тесте какая-то ошибка, степик не показывает какие данные вводились и вообще что там за тест был, по этому хз

#Зум бросил вызов Флэшу и предложил ему честный поединок в виде гонки вокруг магнетара.
#В случае проигрыша эта нейтронная звезда зарядится и уничтожит мир, поэтому Флэш решил не рисковать без причины и узнать у своего друга Циско Рамона, есть ли смысл принимать вызов.
#Циско получил данные, что скорость Зума равна n, а скорость Флэша равна k
#Напишите программу, которая должна вывести ответ Циско на вопрос Флэша.
def poedinok():
    speed = [int(input()), int(input())]
    if speed[0] <= speed[1]:
        if speed[0] == speed[1]:
            print("Don't know")
        else:
            print("YES")
    else:
        print("NO")
#poedinok()

#Напишите программу, которая классифицирует треугольник на основе длин его сторон.
# Программа должна принимать три числа, каждое из которых представляет собой длину одной из его сторон.
# В результате программа должна определить, является ли треугольник равносторонним, равнобедренным или разносторонним.
def triugolnik():
    size = [int(input()), int(input()), int(input())]
    if size[0] == size[1] == size[2]:
        print("Равносторонний")
    elif size[0] == size[1] or size[0] == size[2] or size[1] == size[2]:
        print("Равнобедренный")
    else:
        print("Разносторонний")
#triugolnik()

#Даны три различных целых числа. Напишите программу, которая находит серединное значение из этих чисел.
def seredina():
    numUser = [int(input()), int(input()), int(input())]
    print(sorted(numUser)[1])
#seredina()

#Дан порядковый номер месяца (1,2,…,12). Напишите программу, которая выводит на экран количество дней в этом месяце. Принять, что год является невисокосным.
def kolvoDnei():
    numUser = int(input())
    days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    print(days[numUser - 1])
#kolvoDnei()

#Известен вес боксёра-любителя (целое число). Известно, что вес таков, что боксер может быть отнесён к одной из трёх весовых
#Напишите программу, определяющую, в какой категории будет выступать данный боксёр.
def vesBoksera():
    numUser = int(input())
    if numUser < 60:
        print("Легкий вес")
    elif 60 <= numUser < 64:
        print("Первый полусредний вес")
    else:
        print("Полусредний вес")
#vesBoksera()

#Напишите программу, которая считывает с клавиатуры два целых числа и строку.
# Если эта строка является обозначением одной из четырёх математических операций (+, -, *, /), то выведите результат применения этой операции к введённым ранее числам,
# в противном случае выведите «Неверная операция» (без кавычек).
# Если пользователь захочет поделить на ноль, выведите текст «На ноль делить нельзя!» (без кавычек).
def kalkulator():
    numUser = [int(input()), int(input())]
    operation = input()
    if not operation in ["+", "-", "*", "/"]:
        print("Неверная операция")
        raise SystemExit
    if numUser[0] == 0 or numUser[1] == 0:
        if operation == "/":
            print("На ноль делить нельзя!")
            raise SystemExit
    if operation == "+":
        print(numUser[0] + numUser[1])
    elif operation == "-":
        print(numUser[0] - numUser[1])
    elif operation == "*":
        print(numUser[0] * numUser[1])
    else:
        print(numUser[0] / numUser[1])
#kalkulator()

#Напишите программу, которая считывает названия двух основных цветов для смешивания.
# Если пользователь вводит что-нибудь помимо названий «красный», «синий» или «желтый», то программа должна вывести сообщение об ошибке.
# В противном случае программа должна вывести название вторичного цвета, который получится в результате.
def palitra():
    vvodUser = [input(), input()]
    if vvodUser[0] == "красный" and vvodUser[1] == "синий":
        print("фиолетовый")
    elif vvodUser[0] == "красный" and vvodUser[1] == "желтый":
        print("оранжевый")
    elif vvodUser[0] == "красный" and vvodUser[1] == "красный":
        print("красный")
    elif vvodUser[0] == "синий" and vvodUser[1] == "красный":
        print("фиолетовый")
    elif vvodUser[0] == "синий" and vvodUser[1] == "желтый":
        print("зеленый")
    elif vvodUser[0] == "синий" and vvodUser[1] == "синий":
        print("синий")
    elif vvodUser[0] == "желтый" and vvodUser[1] == "красный":
        print("оранжевый")
    elif vvodUser[0] == "желтый" and vvodUser[1] == "синий":
        print("зеленый")
    elif vvodUser[0] == "желтый" and vvodUser[1] == "желтый":
        print("желтый")
    else:
        print("ошибка цвета")
#palitra()

#На колесе рулетки карманы пронумерованы от 0 до 36.
#Напишите программу, которая считывает номер кармана и показывает, является ли этот карман зеленым, красным или черным.
#Программа должна вывести сообщение об ошибке, если пользователь вводит число, которое лежит вне диапазона от 0 до 36.
def ruletka():
    numUser = int(input())
    if numUser == 0:
        print("зеленый")
    elif 1 <= numUser <= 10:
        if numUser % 2 == 0:
            print("черный")
        else:
            print("красный")
    elif 11 <= numUser <= 18:
        if numUser % 2 == 0:
            print("красный")
        else:
            print("черный")
    elif 19 <= numUser <= 28:
        if numUser % 2 == 0:
            print("черный")
        else:
            print("красный")
    elif 29 <= numUser <= 36:
        if numUser % 2 == 0:
            print("красный")
        else:
            print("черный")
    else:
        print("ошибка ввода")
#ruletka()

#На числовой прямой даны два отрезка:[a1;b1] и [a2;b2]. Напишите программу, которая находит их пересечение.
def peressechenie():
    numUser = [[int(input()), int(input())], [int(input()), int(input())]]
    if numUser[0][1] > numUser[1][0]:
        print(numUser[1][0], numUser[0][1])
    elif numUser[0][0] < numUser[1][1]:
        print(numUser[0][0], numUser[1][1])
    elif numUser[0][1] == numUser[1][0]:
        print(numUser[0][1])
    elif numUser[0][0] == numUser[1][1]:
        print(numUser[0][0])
    elif numUser[0][1] < numUser[1][0]:
        print("пустое множество")
    elif numUser[0][0] > numUser[1][1]:
        print("пустое множество")
peressechenie()
