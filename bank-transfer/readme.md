# Bank Transfer

## Workflow

### Chuyển tiền

1. Lấy thông tin session từ Session JWT (`Get session info`)
2. Lấy ds tài khoản theo session (`Get accounts`)
3. Tạo lệnh YC chuyển tiền (`Create bank transfer`)
4. Kiểm tra số dư (`Check balance account`)
5. Kiểm tra tra tài khoản đích (`Check target account`)
6. Tạo giao dịch chuyển tiền (`Create new transaction`)
7. Tạo giao dịch ghi nợ
8. Tạo giao dịch ghi có
9. Gửi thông báo đã chuyển tiền

## APIs

- 1. Lấy DS giao dịch chuyển khoản: GET `/bank-transfer/transfers`
- 2. Kiểm tra số dư: GET `/accounts/:ID/balance`

## Quickstart

```bash
go run worker/main.go
go run api/main.go
```

## Stack for Sample

- `fiber`
- `temporal`

## FAQ

- <https://github.com/temporalio/temporal-ecommerce>