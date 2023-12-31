# Project created to practice iterating, working with structs, interfaces, json & files and goroutines

### Iterating
        Create a function called "removeDuplicates" that takes a slice of integers as input and returns a new slice with duplicate values removed. The order of the elements in the resulting slice should be maintained.

        Example Output: Input: [1 2 2 3 4 4 4 5] Output: [1 2 3 4 5]

        Create a function called "countOccurrences" that takes a slice of strings as input and returns a map where the keys are the unique strings from the slice, and the values are the number of occurrences of each string in the slice.

        Input: ["apple" "banana" "apple" "orange" "banana" "banana"] Output: map[apple:2 banana:3 orange:1]

        Create a function called "mergeMaps" that takes two maps with string keys and integer values as input and returns a new map that contains the combined key-value pairs from both input maps. If a key exists in both input maps, the value in the resulting map should be the sum of the corresponding values from the input maps.

        Input Map 1: map[apple:2 banana:3] Input Map 2: map[banana:4 orange:1] Output: map[apple:2 banana:7 orange:1]

        Write a main function that demonstrates the usage of the above functions. Create sample input data and call each function with appropriate arguments. Print the results to the console.

        Each function should have also unit tests.

### Working with structs
        Define a struct type called "Rectangle" to represent a rectangle. The struct should have the following fields:

        Width (float64): representing the width of the rectangle.
        Height (float64): representing the height of the rectangle.
        Create a method called "Area" for the Rectangle struct that calculates and returns the area of the rectangle.

        Create a method called "Perimeter" for the Rectangle struct that calculates and returns the perimeter of the rectangle.

        Define a struct type called "Circle" to represent a circle. The struct should have the following fields: Radius (float64): representing the radius of the circle.

        Create a method called "Area" for the Circle struct that calculates and returns the area of the circle.

        Create a method called "Circumference" for the Circle struct that calculates and returns the circumference of the circle.

        Write a main function that demonstrates the usage of the Rectangle and Circle structs. Create instances of both structs, initialize their fields with appropriate values, and call the "Area" and "Perimeter" methods for the Rectangle struct and the "Area" and "Circumference" methods for the Circle struct. Print the results to the console.

-Note:

        Use the built-in "math" package for mathematical calculations. The formula for calculating the area of a rectangle is: area = width * height. The formula for calculating the perimeter of a rectangle is: perimeter = 2 * (width + height). The formula for calculating the area of a circle is: area = π * radius^2. The formula for calculating the circumference of a circle is: circumference = 2 * π * radius. You can assume the value of π to be 3.14159.

        Example Output: Rectangle Width: 5.0 Height: 3.0 Area: 15.0 Perimeter: 16.0

        Circle Radius: 2.5 Area: 19.6349375 Circumference: 15.70796

### Interfaces
        Define an interface called "Shape" with the following methods:

        Area() float64: returns the area of the shape.
        Perimeter() float64: returns the perimeter of the shape.
        Create two struct types, "Rectangle" and "Circle," that implement the "Shape" interface.

        Implement the "Area" and "Perimeter" methods for the Rectangle struct based on its width and height.

        Implement the "Area" and "Perimeter" methods for the Circle struct based on its radius.

        Write a function called "PrintShapeDetails" that takes a Shape interface as input and prints the area and perimeter of the shape.

        Write a main function that demonstrates the usage of the "Shape" interface and the "PrintShapeDetails" function. Create instances of both Rectangle and Circle, and pass them to the "PrintShapeDetails" function. Ensure that the function correctly prints the area and perimeter of each shape.

        Example Output: Rectangle Width: 5.0 Height: 3.0 Area: 15.0 Perimeter: 16.0

        Circle Radius: 2.5 Area: 19.6349375 Circumference: 15.70796

### Working with json & files
        Define a struct called "Person" with the following fields:

        Name (string): representing the person's name.
        Age (int): representing the person's age.
        Email (string): representing the person's email address.
        Create a function called "encodeJSON" that takes a Person struct as input and returns the JSON representation of the struct as a byte slice.

        Create a function called "decodeJSON" that takes a byte slice (containing JSON data) as input and returns a Person struct.

        Write a main function that demonstrates the usage of the "encodeJSON" and "decodeJSON" functions. Create an instance of the Person struct, encode it to JSON, and then decode it back to a struct. Print the decoded struct's fields to the console.

        Example: (Output) {Name: "John Doe", Age: 25, Email: "johndoe@example.com"}

        Decoded Person struct: Name: John Doe Age: 25 Email: johndoe@example.com


 ### Simple Concurrency Goroutines
  
        Create a function called "printMessage" that takes a string as input and prints it to the console.

        In the main function, create a goroutine that invokes the "printMessage" function with the argument "Hello, World!".

        After starting the goroutine, print the message "Goroutine started!" to the console from the main function.

        Use a delay or wait mechanism to ensure that the goroutine has enough time to execute before the main function terminates.

        Verify that the output displays "Goroutine started!" before the message "Hello, World!".

        Example Output: Goroutine started! Hello, World!

-  Practical Concurency
  
        Define a function called "processFile" that takes a filename (string) as input and simulates some processing on the file. For this example, you can simply print the filename and a message indicating that processing is being done.

        In the main function, create a slice of filenames (at least 5 filenames) that represent files to be processed concurrently.

        Use a loop to iterate through the slice of filenames.

        For each filename, launch a goroutine that calls the "processFile" function with the respective filename as an argument.

        Use a delay or wait mechanism (such as time.Sleep()) to ensure that all goroutines have enough time to complete their processing before the main function terminates.

        Example Output: Assuming the filenames in the slice are ["file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"], the output could be:

        Processing file: file1.txt Processing file: file2.txt Processing file: file3.txt Processing file: file4.txt Processing file: file5.txt

- Note:

        You can modify the "processFile" function to include more meaningful processing, such as reading the contents of the file, performing calculations, or any other desired operations. The actual execution order and interleaving of the messages may vary due to the concurrent nature of goroutines.