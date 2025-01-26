## Description

```API request handler``` handles concurrent API requests for multiple company data.

## Installation & Usage

1. Clone the repository: `git clone https://github.com/imofficialvivek/api-request-handler.git` 
2. Change the directory and checkout to develop branch
```bash
cd api-request-handler
git checkout develop
```
3. Execute below command to run the program
```bash
go run main.go
```
4. Run below command in another terminal
```bash
curl "http://localhost:8080/api/company/financials?companyId=123"
curl "http://localhost:8080/api/sales/data?companyId=123"
```

## Run test (Optional)
1. Execute below command to run the specific test 
```bash
go test -v test\handler_test.go
go test -v test\services_test.go
```