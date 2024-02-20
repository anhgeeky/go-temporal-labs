# Bank Transfer

## To-do

- [ ] Build cấu trúc cho Temporal (Api + Worker)
- [ ] Xây dựng luồng **Chuyển tiền**

## Workflow

### Chuyển tiền

1. Lấy thông tin session từ Session JWT (`Get session info`)
2. Lấy ds tài khoản theo session (`Get accounts`)
3. Tạo lệnh YC chuyển tiền (`Create bank transfer`) (`Start`)
4. Kiểm tra số dư (`Check balance account`) (`Parallel`)
5. Kiểm tra tra tài khoản đích (`Check target account`) (`Parallel`)
6. Tạo giao dịch chuyển tiền (`Create new transaction`) (`When step 4,5 done -> Continue`)
7. Tạo giao dịch ghi nợ (`Parallel`)
8. Tạo giao dịch ghi có (`Parallel`)
9. Gửi thông báo đã chuyển tiền (`Completed`)

## APIs

- 1. Lấy DS giao dịch chuyển khoản: GET `/transfers`
- 2. Kiểm tra số dư: GET `/accounts/:ID/balance`

## Quickstart

```bash
go run worker/main.go
go run api/main.go
```

## Temporal

![Screenshot](/banktransfer/docs/assets/bank-transfer-temporal-admin-log.png)

## Stack for Sample

- `fiber`
- `temporal`

## FAQ

- <https://github.com/temporalio/temporal-ecommerce>