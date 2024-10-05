# Simple Go Struct Method and Mutation

### Overview
This Go program demonstrates the use of struct and struct methods. It features the use of pointers with structs and struct methods to display and overwrite/clear (mutate) the data in the original struct without unnecessarily occupying extra memory spaces. The features to display/clear data were written in a way that they are methods of the struct created in this program. 

### Program manual
When run, the program will present the user with 4 choices which include Input data, Display data, Clear data, and Exit. The user can choose one of these choices by inputing a number preceeding each choice. The program validates the input at this step and will keep asking the user for input until it gets a valid input. The following choices work according to the following descriptions:

1. Input data: The user will be asked to input a student's first name, last name, (string) and score (floating number). There will not be a data validation process in this step. The score will be 0 by default if the input is not a floating number. The program keeps the input information at this step in a struct.

2. Display data: The program reads the information stored in the created struct and display it accordingly. If the user skips #1 and perform this task once the program is run, the default information for each data type will be displayed ("" for string and 0 for floating number).

3. Clear data: The program clear data in the created struct. The score which is of the floating number type will be reset to 0.

4. The program exits with a Goodbye message.

Except for the choice #4, the program will ask if the user wants to continue to perform another tasks (any among the 4 choices) after finishing each task. The user will confirm this question by inputting "Y" or "N" (case-insensitive). The program will continue to work until the user input "N" or "n".

### Code structrue
The program is divided into 2 main parts: the main program codes and the functions which include the function to display the menu, function for getting information from the user and store it in a struct, function to display the information, and function to clear the information in the created struct. 

Each function will be called according to the chosen choice. Once the user completes inputting information, the functions that displays and clears the information will get the original struct pointer as their argument and read/modify the original struct (created by the function that gets user input) directly to prevent unnecessary momory usage.

### Program flow
The program flow is shown as the following numbered list:

1. The program presents the user with 4 choices:

    1.1 Input Data

    1.2 Display Data

    1.3 Clear Data

    1.4 Exit

2. The user can choose any of each choice by inputting the preceeding number. There is no order for choosing any choice. But if the user chooses to display data first, the default data for strig ("") and floating number (0) types will be displayed. If the user chooses to input the data, the user inputs will be stored in a struct with the field names "first name", "last name", and "score". If the user chooses to display or clear the data, the original struct will be read/modify directly according to the use of pointer with struct data type.

3. After each choice, the program will ask if the user wants to perform another operation. The user can also exit the program at this step by inputting "N" or "n"

4. If the user chooses to exit, the program will exit with a goodbye message.
