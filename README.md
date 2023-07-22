# 📅 Atur.in - Course Scheduler
> Backend side of Course Scheduler using Dynamic Programming with React Typescript and Golang

## General Information
Atur.in is a simple course scheduler to plan what course to take based on user prediction. The program will proceed an input from forms to setup what the configuration on filtering process then receive the best plan on what course to take using popular effective algorithm, dynamic programming. User can also add more Fakultas, Jurusan, and Mata Kuliah by filling the forms or upload a .json file for batch. The program also provide a good chart illustration on mata kuliah composition based on the course recommendation. Furthermore, the project information is also provided for future improvements.

## Project Structure
```bash
.
├─── algorithm
│   └─── algorithm.go
├─── db
│   └─── migration
├─── middleware
│   └─── handlers.go
├─── models
│   └─── models.go
├─── repository
│   └─── repository.go
├─── router.go
│   └─── router.go
├─── .env
├─── .env.example
├─── .gitignore
├─── docker-compose.yml
├─── Dockerfile
├─── go.mod
├─── go.sum
├─── main.go
└─── README.md
```

## Prerequisites
- gorilla/mux (v 1.8.0) to handle routing
- net/http to handle request and responseWriter
- encoding/json to parse the json POST body
- driver/posgres as postgresql driver
- Docker

## Algorithms
This section will explain what is Dynamic Programming and how it used to determine the best course recommendation on spesific semester based on minimum and maximum credits (sks) allowed on consecutive semester.

### What is Dynamic Programming
Dynamic programming is a technique used in computer science and mathematics to solve problems by breaking them down into overlapping subproblems and efficiently reusing the solutions to these subproblems. It is typically used when a problem exhibits the property of overlapping subproblems and optimal substructure.

The key idea behind dynamic programming is to store the solutions to subproblems in a table or memoization array so that they can be reused instead of recomputed. This approach can significantly improve the efficiency of algorithms, especially for problems that have an exponential number of overlapping subproblems.

Dynamic programming can be applied to a wide range of problems, such as optimization problems, counting problems, and decision problems. It is commonly used in various fields including algorithms, artificial intelligence, operations research, and bioinformatics.

### How do the algorithm work
The general steps for solving a problem using dynamic programming are as follows:
1. Identify the problem and determine if it exhibits the properties of overlapping subproblems and optimal substructure.
2. Define the recursive relationship or recurrence relation that relates the solution to larger subproblems to the solutions of smaller subproblems.
3. Decide on the order in which the subproblems should be solved. This can be done using either a top-down approach (memoization) or a bottom-up approach (tabulation).
4. Solve the subproblems in a systematic way, storing the solutions in a table or memoization array.
5. Build the solution to the original problem using the solutions to the subproblems.
By using dynamic programming, it is often possible to significantly improve the time complexity of algorithms and solve problems that would be otherwise infeasible to solve efficiently.

### Applying algorithm to course scheduling
By adapting the dynamic programming approach, here's how to bring recommendation based on user prediction over spesific course, credits limit, and semester taken.
1. **Input:** Gather the necessary input, which includes the list of courses with their corresponding prediction scores and credits needed for the specific semester. Also, take note of the credit limit for the semester (in this case the minimum and maximum range).
2. **Define the subproblems:** In this case, the subproblems can be defined as finding the best score achievable with a given number of credits within the semester. We can consider the number of credits remaining as the state for our dynamic programming solution.
3. **Define the recurrence relation:** The recurrence relation specifies how to calculate the solution to a subproblem based on the solutions to smaller subproblems. In this case, we can define it as follows:
    1. Let ```DP[i]``` be the maximum score achievable with $i$ credits within the semester.
    2. For each course $j$, if the credits needed for course $j$ is less than or equal to $i$, we can consider taking that course. The score obtained would be the sum of the prediction score for course $j$ and the maximum score achievable with ```i - credits_needed[j]``` credits.
    3. Therefore, the recurrence relation can be defined as: ```DP[i] = max(DP[i], DP[i - credits_needed[j]] + prediction_score[j])``` for each course $j$.
4. **Initialize the table:** Create a table or memoization array to store the solutions to the subproblems. Initialize it with a base case, such as ```DP[0] = 0```, indicating that the maximum score achievable with 0 credits within the semester is 0.
5. **Build the solution using dynamic programming:** Iterate over the number of credits from 1 to the credit limit for the semester. For each credit value, iterate over the list of courses. Calculate the maximum score achievable with the current number of credits using the recurrence relation defined in step 3. Update the table accordingly.
6. **Retrieve the final solution:** After completing the dynamic programming calculations, the maximum achievable score can be obtained from the last entry in the table, the ```DP[credit_limit]```.
7. **Track the selected courses:** To determine which courses were selected to achieve the maximum score, we can maintain an additional array or data structure alongside the table. Whenever we update a value in the table, store the corresponding course that contributed to that maximum score.
8. **Output:** Output the maximum achievable score and the selected courses for the semester.

### How the "total score" is calculated
Course prediction valid value is A, AB, B, BC, C, D, and E. Here is the table value of every index mapped to score
``` bash
A   : 4.0
AB  : 3.5
B   : 3.0
BC  : 2.5
C   : 2.0
D   : 1.0
E   : 0.0
```
While every course have their own credits (sks), the total is calculated as sum of product between mapped value of course prediction and the sks divided by total sks. Here is the example.<br/>
Consider having a list of course (Mata Kuliah) like this.
| Course name                   | Prediction    | Credits     |
| ----------------------------- | ------------- |-------------|
| Strategi Algoritma            | A             | 3           |
| Aljabar Linear dan Geomteri   | AB            | 3           |
| Algoritma dan Struktur Data   | B             | 4           |

Then the total value is 
$$\frac{(4.0 \times 3) + (3.5 \times 3) + (3.0 \times 4)}{3 + 3 + 4} = 3.45$$

### Algorithm time and space complexity
- The algorithm uses a two-dimensional array dp of size $(n + 1) \times (C + 1)$ to store the maximum achievable score for each subproblem. The initialization of this array takes $O(n \times C)$ time.
- It uses a nested loop to iterate over each course $(i)$ and each credit limit $(j)$ from $1$ to $C$. For each subproblem, it calculates the maximum achievable score based on the previous subproblems. The calculation for each subproblem takes constant time $(O(1))$.
Therefore, the total time taken to fill the entire dp array is $O(n \times C)$.
The space complexity of the algorithm is also $O(n \times C)$ since it uses the dp array of size $(n + 1) \times (C + 1)$ and the selectedCourses array of the same size.

### Fakultas and Jurusan implementation
Due to the bonus task, Fakultas have one-to-many relation with Jurusan. This make Jurusan needs to save Fakultas ID to make a better data binding based on foreign key references. The implementation is used as follows :
1. User with Jurusan $X$ could take all course from Fakultas $Y$, where $Y$ is the Fakultas which holds Jurusan $X$.
2. User don't need to input Fakultas when adding new Mata Kuliah because the Jurusan save the Fakultas ID and make the data binding so easy to do.

## How to Compile and Run the Program
Before running the service, make sure you have docker installed. Click [this link](https://docs.docker.com/get-docker/) for installation.<br />
Clone this repository from terminal with this command
``` bash
$ git clone https://github.com/mikeleo03/Course-Scheduler_Backend.git
```
### Run the application on development server
Compile the program by running the following *command*
``` bash
$ docker-compose up -d
```
If you do it correctly, the pogram should be running on localhost:8080.

### Run the application after doing updates
To run the program after doing updates, you can add a build tag by using this *command*
``` bash
$ docker-compose up -d --build
```

## Available Scripts
In the project directory, you can run:

### `go run main.go`

Ths runs the app in the development mode.

The page will reload if you make edits.<br />
You will also see any lint errors in the console. You can also use the environment by appyling the basic .env configuration on .env.example file.<br />
Also don't forget to uncomment the section in repository go that have comment "Uncomment for local development"

## References
- https://go.dev/doc/
- https://www.postgresql.org/docs/
- https://www.yugabyte.com/
- https://www.docker.com/

## Contributors
<a href = "https://github.com/mikeleo03/markdown-editor/graphs/contributors">
  <img src = "https://contrib.rocks/image?repo=mikeleo03/markdown-editor"/>
</a>