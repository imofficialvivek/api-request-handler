## Description

```API request handler``` handles concurrent API requests for multiple company data.

## Installation & Usage

```Note:``` This project is built using ```Go 1.23.5```. Go installation is must to run the project as shown below

1. Clone the repository: `git clone https://github.com/imofficialvivek/api-request-handler.git` 
2. Change the directory and checkout to develop branch
```bash
cd api-request-handler
git checkout develop
```
3. Execute below command in windows to run the program
```bash
go run cmd\main.go
```
4. Run below command in another terminal
```bash
curl "http://localhost:8080/api/company/financials?companyId=123"
curl "http://localhost:8080/api/sales/data?companyId=123"
curl "http://localhost:8080/api/employee/stats?companyId=123"
```


```Note:``` If you want to run the program without installing Go then just execute the windows executable file called ```api-request-handler.exe``` (Windows OS is required), provided in the exe folder of repo. 

Run below command in another terminal
```bash
curl "http://localhost:8080/api/company/financials?companyId=123"
curl "http://localhost:8080/api/sales/data?companyId=123"
curl "http://localhost:8080/api/emloyee/stats?companyId=123"
```

## Run test (Optional)
1. Execute below command to run the specific test 
```bash
go test -v test\cache_test.go
go test -v test\handler_test.go
go test -v test\services_test.go
```


```Note:``` For more information, refer to documentation ```API-Request-Handler.pdf``` available in doc folder.