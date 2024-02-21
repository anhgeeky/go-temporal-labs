# Bank Transfer

## To-do

- [x] Build cấu trúc cho Temporal (Api + Worker)
- [x] Xây dựng luồng **Thông báo** (Notification Flow)
- [ ] Xây dựng luồng **Chuyển tiền** (Transfer Flow)
  - [ ] Bổ sung require before steps (`When step 2.4, 2.5 done -> Completed`)

## Workflow

### Before: Chuyển tiền

1. Lấy thông tin session từ Session JWT (`Get session info`)
2. Lấy ds tài khoản theo session (`Get accounts`)

### Start: Chuyển tiền

1. [**Transfer Flow**] Tạo lệnh YC chuyển tiền (`Create bank transfer`) (`Start`)
  - Run [**Notification Flow**] send OTP
2. [**Transfer Flow**] Xác thực OTP (`Trigger Signal`)
  - 2.1. Kiểm tra số dư (`Check balance account`) (`Synchronize`)
  - 2.2. Kiểm tra tra tài khoản đích (`Check target account`) (`Synchronize`)
  - 2.3. Tạo giao dịch chuyển tiền (`Create new transaction`) (`When step 2.1, 2.2 done -> Continue`)
  - 2.4. Tạo giao dịch ghi nợ (`Synchronize`)
  - 2.5. Tạo giao dịch ghi có (`Synchronize`)
  - 2.6. Transfer done  (`When step 2.4, 2.5 done -> Completed`) (`Trigger [Notification Flow]`)
  - 2.7. Call subflow [**Notification Flow**] Gửi thông báo đã chuyển tiền
    - 2.7.1 Lấy thông tin `token` của các thiết bị theo tài khoản
    - 2.7.2 Push message SMS thông báo đã `Chuyển tiền Thành công` (`Parallel`)
    - 2.7.3 Push message notification vào `firebase` (`Parallel`)
    - 2.7.4 Push message internal app, reload lại màn hình hiện tại `Đang xử lý` -> `Thành công` (`Parallel`)

### End: Chuyển tiền

1. Nhận message internal app
2. Lấy thông tin kết quả chuyển tiền
3. Reload lại show kết quả `Chuyển tiền Thành công`

## APIs

- [x] Lấy DS giao dịch chuyển khoản: GET `/transfers`
- [ ] Kiểm tra số dư: GET `/accounts/:ID/balance`

## Quickstart

```bash
go run worker/main.go
go run api/main.go

# or 
sh start.sh
```

## Temporal

![Screenshot](/banktransfer/docs/assets/bank-transfer-workflows.jpg)
![Screenshot](/banktransfer/docs/assets/bank-transfer-temporal-admin-log.png)
![Screenshot](/banktransfer/docs/assets/bank-transfer-sub-workflow-temporal-admin-log.png)

## Stack for Sample

- `fiber`
- `temporal`
- `viper`

## FAQ

- <https://github.com/temporalio/temporal-ecommerce>