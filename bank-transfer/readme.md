# Bank Transfer

## Quickstart

```bash
env STRIPE_PRIVATE_KEY=stripe-key-here env MAILGUN_DOMAIN=mailgun-domain-here env MAILGUN_PRIVATE_KEY=mailgun-private-key-here go run worker/main.go
env STRIPE_PRIVATE_KEY=stripe-key-here env MAILGUN_DOMAIN=mailgun-domain-here env MAILGUN_PRIVATE_KEY=mailgun-private-key-here env PORT=3001 go run api/main.go
```

## FAQ

- <https://github.com/temporalio/temporal-ecommerce>