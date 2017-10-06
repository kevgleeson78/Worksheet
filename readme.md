Problem Sheet Number one


This readme is a record of how I went about solving each problem with go language.

Problem 1 
Hello world

Simple hello world program that prints out a message to the screen with the Japanese characters for "Hello World".

To print out to the screen one of the commands that can be used is fmt.Println.
The fmt package needs to be imported fore this to work.

Problem 2
Time

This problems goal was to print out the current time to the user.
The time package needs to be imported for this to work.
The time.now() method is simple put inside the a Println method to print out the current time to the user.

Problem 3
Fizzbuzz

The purpose of this program is to print out Fizz for every multiple of three, buzz for multiples of five and fizzbuzz for multiples of 3 and five.

A for loop was used to achieve this with variables set to modulus 5 and modulus 3 for each iteration of the loop. If they are the modulo of the number’s they get stored in the variable for the duration of one iteration.
Then a conditional is set to check if the number is a multiple of the two numbers and finally to check if the single numbers are a multiple of the current number.

A message is printed to the screen instead of the number when the conditions are met.

Problem 4
Sum of Factorial

This program takes in a number and calculates the factorial of that number.
e.g. 4! = 4 x 3 x 2 x 1.
Then each digit of the answer is added together and displayed to the user.
This problem was spread over two functions to achieve this.
One function to calculate the factorial of a number and then an add function to get the sum of the digits within the number.

When 100 is input to the function a simple int variable could hold the number and produced 0.
The bigInt variable needs to used to handle large numbers.
The bigInt variables have there own methods with for loops, adding multiplying etc...

The factorial function takes in a value and loops through the number multiplying the number by the next value of i in the iteration. the result is then passed on to the add function.

The add function simply gets the modulus of the number and stores it in a variable called add.
The number is then divided by ten and then this number becomes the new number at the beginning of the loop.
this repeats until the number is 0.
Then the two variables are added together to give the digit sum of the number.

Problem 5
Guessing Game

This Program Generated a random number and the user must attempt to guess it.
A loop is used to achieve this. The loop has no conditions set it is like a while loop in java.
With each guess the loop will check if the  users number is too high, too low or correct. Conditionals are used to check each guess. If the user input is either too high or is too low of the value of the i variable will increase by one. This is used to keep track of the total amount of attempts the user has had. If the user has entered the same number in succession the counter i will decrease by one thus not being registered as an attempt. Finally, if the correct number was guessed a condition was placed at the end of the loop to break out and end the application.

Problem 6
Min Max element of list.

This application loops through a populated list of elements.
It then returns the largest and smallest element in the list.
The list package had to be imported for this program.
A new list object was created to store the numbers.
PushFront method is used to populate the list.
A loop is then used to check each element in the list.
The front() method is used to get the first number in the list and then checked against the next number by using the next() method.
Each of these numbers are checked against each other and the larger of the two get stored in the max variable for the duration of the loop.
the same was achieved for the smallest number.

Finally, the two numbers are returned from the function and printed out to the user.


Problem 7
Palindrome

The purpose of this program is to check if a string is a palindrome.
rune variables are used to manipulate the characters.
 A loop is used to iterate over the word from the last rune (char) to the first using i >= and i--.
the reversed word and original word were then cast into strings and stored in new variables.
finally, they were checked against each other for equality and a message is printed out depending the result.

Problem 8
Merge/Sort List

This is a function the sorts merges and returns two or more lists.
Arrays are used to demonstrate this.
There are two unordered lists of different sizes.
The sortInts(VarNAme) method is used to sort each list.
A new variable was created to hold the merged list using append(list1,list1...)
method.
The new list is then sorted as above.
The result is then printed out to the user.


Problem 9
Newtons Method

This program is uses Newtons method to find the square root of a number.
It takes in two values and returns one number.
A form of recursion is used to solve this problem.
Firstly, the initial numbers are input by the user to the function.
The values for x (the number to find the square root) and z beginning at 1.
The variable z_next is used to store the result of the calculation.
A condition is used to check if the number z_next is equal to the square root of the number entered.

If it is not the same the result of the calculation is then passed back into the function and z_next replaces z in the equation.
If z_next is equal to the square root of the number, then the number is returned and the function returns the final number and the loop terminates.
Each result is printed out to the terminal.


Problem 10
Reverse String

This application takes a string and reverse it out to the user.
I reused the reverse function of the Palindrome program to solve this problem.
Rune variables were used to manipulate the string entered.
A for loop started at the end of the string and appended the last rune to the start of the new string.
The word is then cast back into a string and displayed out to the user on the screen.



