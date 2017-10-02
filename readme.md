Problem Sheet Number one


This readme is a record of how i went aboput solving each problem with go language.

Problem 1 
Simple helo world program that prints out a message to the screen with the Japanese characters for "Hello World".

To print out to the screen one of the commands that can be used is fmt.Println.
The fmt package needs to be imported fore this to work.

Problem 2

This problems goal was to print out the current time to the user.
The time package needs to be imported for this to work.
The time.now() method is simple put inside the a Println method to print out the current time to the user.

Problem 3

The purpose of this program is to print out Fizz for every multiple of three, buzz for multiples of five and fizzbuzz for multiples of 3 and five.

A for loop was used to achieve this with variables set to modulus 5 and modulus 3 for each iteration of the loop. If they are modulo of teh numbers they get stored in the variable for the duration of one itteration.
Then a conditional is set to check if the number is a multiple of the two numbers and finally to check if the single numbers are a multiple of the current  number.

A message is printed to the screen instead of the number when the conditions are met.

Problem 4
This program takes in a number and calculates the factoral of that number.
e.g 4! = 4 x 3 x 2 x 1.
then each digit of the answer is added together and displayed to the user.
This problem was spread over two functions to achieve this.
One function to calculate the factoral of a number and then an add function to get the sum of the digits within the number.

When 100 was input to the screen a simple int varaible could hold the number produced and output a 0.
The bigInt variable needs to used to handel large numbers.
The bigInt variables have there own methods with for loops, adding multiplying etc...

The factoral function takes ion a value and loops through the number multiplying the number by the next one in the itteration. the result is then passed on to the add function.

the add function simply gets the modulus of the number and stores it in a variable called add.
The number is then devided by ten and then this number becomes the new number at the beginning of the loop.
this repeats until th number is 0.
then the two variables are added together to give the digit sum of the number.

Problem 5

This Program Generated a random number and the user had to guess it.
A loop was used to achieve this. The loop had no conditions set it is simular to a while loop in java.
With each guess the loop would check if the the users number was too high, too low or correct. Conditionals were used to check each guess. If the user was either high or low of the number the i variables would increase by one. This was used to keep track of the total amount of attempts the user had. If the user entered the same number the counter i would decrease by one thus not being registered as an attempt. Finally if the correct number was guessed a condition was palced at the end of the loop to break out of it and end the application.

Problem 6

This application loops through a populated list of elements.
It then returns the largest and smalest element in the list.
The list package had to imported for this program.
A new list object was created to store the numbres.
PushFront method is used to populate the list.
A loop was then used to check each element in the list.
the front() method was used to get the first number in the list it was then checked agains the next number by using the next() method.
Each of these numbers werechecked against each other and the larger of the two got stored in max variable for the duration of the loop.
the same was achieved for the smallest number.

Finally the two numbers a re returned from the function and printed out to the user.


Problem 7

The pupose of this program is to check if a string is a palendrome.
rune varibles were used to manipulate the cahracters.
a loop was used to itterate over the word from the last rune (char) to the first using i >= and i--.
the reversed word and original word were then cast into strings and stored in new variables.
finally they were checked against each other for equallity and a message was printed out depending the result.

Problem 8

this is a function the sorts merges and retuns two or more lists.
Arrays were used to demonstrate this.
There are two unordered listssof different sizes.
The sortInts(VarNAme) method was used to sort each list.
A new variable was created  to hold the merged list using append(list1,list1...)
method.
the nw list was then sorted as obved.
the  result was then printed out to the user.


Problem 9
This program is uses Newtons method to find the square root of a number.
It takes in two values and reurns one number.
A form of recursion is used to solve this problem.
Firstly the initial numbers are input by the user to the function.
The values for x (the number to find the square root) and z begining at 1.
The variable z_next is used to store the result of the calculation.
A condition is used to check if the number is equal to the square root of the original number.
If it is the same then the number is returned.

If it is not the same the result of the calculation is then passed back into the function and z_next replaces z in the equation.
This continues until the number have parity.

Problem 10

This application was to take a string and reverse it out to the user.
i used the reverse part of the Palindrome to solve this problem.
Rune variables were used to maniplutae the string entered.
A for loop started at the end of the string and appended the last rune to the start of the new string.
The word or was then cast back into a string and dispalyed out to the user on the screen.



