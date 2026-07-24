#Напишите функцию, которая разбивает строку и преобразует ее в массив слов.
from numbers import Number
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
  # TODO May the force be with you
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

print(23 // 7, 20 // 5, 2 // 5, 123 // 10, -123 // 10)