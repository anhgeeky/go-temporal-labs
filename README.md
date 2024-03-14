# Temporal Labs

## To-do

- [x] Build cấu trúc cho Temporal (Api + Worker)
- [x] Xây dựng luồng **Thông báo** (Notification Flow)
- [ ] Xây dựng luồng **Chuyển tiền** (Transfer Flow)
  - [ ] Bổ sung require before steps (`When step 2.4, 2.5 done -> Completed`)
- [x] Tái cấu trúc project for Temporal
  - [x] Chia nhỏ submodule cho `workflow`, `api`
  - [x] Config `namespace` for temporal
  - [x] Bổ sung thêm các features chung cho `workflow` core
- [ ] Saga for Temporal
  - [x] Saga sample with `REST Api`
  - [ ] Saga sample with `Kafka Event Driven`
- [ ] Add or Remove 1 activity
  - Follow: <https://community.temporal.io/t/update-activity-and-or-workflow-inputs/4972/5>
  - Temporal chỉ chạy từng activity, có `STOP` cluster, khi chạy lại vẫn còn `Running` thì sẽ chạy lại
  - Nếu có add or remove 1 activity thì sẽ load lại các activity đã update (add, remove, update) -> Chạy tiếp tục

## Issues

- [ ] 1. chỉnh workflow áp dụng kafka (Database có thể produce vào Request & Response, MService chỉ 1 chiều nhận từ kafka)
- [ ] 2. áp dụng commit hub
- [ ] 3. interface luồng để dev ko bị sót khi code rollback

## Quickstart

```bash
go run ./banktransfer/cmd/worker/main.go
# or
go run ./notification/cmd/worker/main.go
# or 
sh start-worker.sh
```

## Saga (Temporal + Kafka + Microservices)

![Screenshot](/docs/assets/saga-workflows-sample.png)

## Workers

- `banktransfer`
- `notification`
- `onboarding`

<!-- ## Bank Transfer Workflow (Implement Saga)

### Before: Chuyển tiền

1. Lấy thông tin session từ Session JWT (`Get session info`)
2. Lấy ds tài khoản theo session (`Get accounts`)

### Start: Chuyển tiền

1. [**Transfer Flow**] Tạo lệnh YC chuyển tiền (`Create bank transfer`) (`Start`)
  - Run [**Notification Flow**] send OTP
2. [**Transfer Flow**] Xác thực OTP (`Trigger Signal`)
  - 2.1. Kiểm tra số dư (`CheckBalance`) (`Synchronize`)
  - 2.2. Kiểm tra tra tài khoản đích (`CheckTargetAccount`) (`Synchronize`)
  - 2.3. Tạo giao dịch chuyển tiền (`CreateTransaction`) (`When step 2.1, 2.2 done -> Continue`)
  - 2.4. Tạo giao dịch ghi nợ (`WriteCreditAccount`) (`Synchronize`)
  - 2.5. Tạo giao dịch ghi có (`WriteDebitAccount`) (`Synchronize`)
  - 2.6. Transfer done  (`When step 2.4, 2.5 done -> Completed`) (`Trigger [Notification Flow]`)
  - 2.7. Call subflow [**Notification Flow**] Gửi thông báo đã chuyển tiền
    - 2.7.1 Lấy thông tin `token` của các thiết bị theo tài khoản
    - 2.7.2 Push message SMS thông báo đã `Chuyển tiền Thành công` (`Parallel`)
    - 2.7.3 Push message notification vào `firebase` (`Parallel`)
    - 2.7.4 Push message internal app, reload lại màn hình hiện tại `Đang xử lý` -> `Thành công` (`Parallel`)

### End: Chuyển tiền

1. Nhận message internal app
2. Lấy thông tin kết quả chuyển tiền
3. Reload lại show kết quả `Chuyển tiền Thành công` -->

<!-- ## APIs

- [x] Lấy DS giao dịch chuyển khoản: GET `/transfers/:workflowID`
- [x] Kiểm tra số dư: GET `/accounts/:workflowID/balance`
- [ ] Kiểm tra tra tài khoản đích (`CheckTargetAccount`)
- [x] Tạo giao dịch chuyển tiền (`CreateTransaction`): POST `/transfers/:workflowID/transactions`
- [x] Tạo giao dịch ghi nợ (`WriteCreditAccount`): POST `/transfers/:workflowID/credit-accounts`
- [x] Tạo giao dịch ghi có (`WriteDebitAccount`): POST `/transfers/:workflowID/debit-accounts`
- [x] Add new activity for test: POST `/transfers/:workflowID/new-activity`
- [x] [Rollback] Tạo giao dịch chuyển tiền (`CreateTransactionCompensation`): POST `/transfers/:workflowID/transactions/rollback`
- [x] [Rollback] Tạo giao dịch ghi nợ (`WriteCreditAccountCompensation`): POST `/transfers/:workflowID/credit-accounts/rollback`
- [x] [Rollback] Tạo giao dịch ghi có (`WriteDebitAccountCompensation`): POST `/transfers/:workflowID/debit-accounts/rollback`
- [x] [Rollback] Add new activity for test: POST `/transfers/:workflowID/new-activity/rollback` -->

## Transfer Flow Activities

![Screenshot](/docs/assets/OCB-Fund-Transfer-Demo.png)

## Temporal Versioning

- Issue downgrade `2.0` -> `1.0` -> Error all workflows
![Screenshot](/docs/assets/temporal-worker-versioning-error-downgrade-version.jpg)

<!-- ## Saga

![Screenshot](/docs/assets/bank-transfer-saga-pattern-log.png)
-->

## Worker versioning

![Screenshot](/docs/assets/temporal-versioning-1.0.png)
![Screenshot](/docs/assets/temporal-versioning-2.0.png)
![Screenshot](/docs/assets/temporal-versioning-3.0.png)
![Screenshot](/docs/assets/temporal-versioning-4.0.png)

## Stack

- `fiber`: <https://github.com/gofiber/fiber>
- `temporal`: <https://github.com/temporalio/temporal>
- `viper`: <https://github.com/spf13/viper>

## FAQ

- Temporal for Docker: <https://github.com/temporalio/docker-compose>
- <https://github.com/temporalio/samples-go>
