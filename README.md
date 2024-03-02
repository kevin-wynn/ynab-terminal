# ynab-terminal

**This is a work in progress**

YNAB budget right in your terminal! Get quick views of your budgets, balances, and accounts in your terminal view.

## TODO/Roadmap:

- Fix when rounding for currency, if the final int is a `0` it doesn't show
- Support more currency than just USD
- Set up tview and build a dashboard

  - Left hand panel for budgets/accounts
  - Center panel table for showing budget
  - Center panel swappable for account/transactions

- Set up a config to store access tokens for users instead of using .env files
- Allow for user to define "quick view items" to select categories and accounts to return by calling `ynab-terminal --quick`
- Prompt user to set up alias in dotfile for just calling `ynab` command

## Dep Docs:

tview docs: https://pkg.go.dev/github.com/rivo/tview#pkg-overview
ynab docs: https://api.ynab.com/
