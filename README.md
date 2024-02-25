# Golang API Manaage User Photo's
This project is a RESTful API developed using the Go programming language and the Gin framework. It's designed to handle user authentication and manage a collection of photos. Users can register, login, update profile, upload photos, and update or delete their photos. The API uses JWT (JSON Web Tokens) for secure authentication and GORM for object-relational mapping, interfacing with a MySQL database.

## How To Use
### Prerequisites
* Go
* MySQL
* Any REST client (Postmant or Insomnia)
### Instalation
1. Clone the repository 
    - `git clone https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git` 
    - `cd repository task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git`
2. Set up the Database
    - Create MySQL database named 'your_database'
3. Install Dependecies
    - `go mod tidy`
4. Change Setup
   - On the database/setup.go file change the mysql.Open(parameter) to mysql.Open("your_db_host@tcp(localhost:your_db_port)/your_db_name?parseTime=true")
### Running
1. Start the Server
   - `go run main.go`
2. Accesing the API
Use a REST client to access the API. Here are all of the endpoints:
    - POST "/users/register"
    - POST "/users/login"
    - PUT "/users/:userId"
    - DELETE "users/:userId"
    - POST "/photos"
    - GET "/photos"
    - PUT "/photos/:photoId"
    - DELETE "/photos/:photoId"

Some endpoints need authorization, simply add `Authorization: Bearer <Your JWT Token>` in the request header.

## Screenshots
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/af896546-1156-4595-b215-d62a56d703a5)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/80aa0cbc-66f1-46d1-a7da-23cb42c0d2e4)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/e9c9a155-09ba-41f5-ba48-a224a99c12ef)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/58313af7-731d-4d3d-ac70-d6cd164f37ff)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/f6085d49-7c5c-4483-a9c8-745ffb29798b)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/db264614-3d77-4b14-81a6-4ce4b407300a)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/611a58b0-7106-40d1-86e2-0a5a74faa1ac)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/739b1f0e-1b23-4246-861e-39b1cb750c8f)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/a7ccd38f-6e90-4095-8c88-b2b284ce7ce5)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/2c76d065-ed43-4693-bf0b-6320a33ec8c0)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/7911446a-2d5e-485e-bf4e-9187a98c165c)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/d48a37ff-be0f-4292-9cef-05a7060cee90)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/2861eeb3-1755-4e64-a437-c520c30023fc)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/c45d46d0-07de-4943-b1d7-1e018c00fa94)
![image](https://github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya/assets/92701179/15989c25-1546-404d-b363-b13bcbef28ee)













