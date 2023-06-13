# TODO

## Milestone A - MVP

- [x] feat: Edit an expense
- [x] chore: Create Dockerfile & docker compose
- [x] chore: Implement CD with Github Actions

## Milestone B - Creation & Sharing

- [x] feat: Save visited Trikounts & display on home page (using cookies)
- [x] feat: Force user to copy link when creating the Trikount
- [x] feat: Add multiple members when creating a new Trikount
- [x] feat: Delete a Trikount (settings)

## Milestone C - Total & Refunds

- [x] feat: Display total spend by project
- [x] feat: Mark some expenses as refunds
- [x] feat: Create a 1-click refund button

## Milestone E - Production Ready

- [x] fix: Fix missing favicon.ico
- [x] feat: Improve UI (mobile & desktop)
- [x] feat: Delete expense
- [x] feat: See expense date in details page
- [x] feat: Rename a Trikount
- [x] feat: Show cost per person
- [x] chore: Fix typos in wording
- [x] feat: Create a dedicated "share" button
- [x] feat: Show the most recent expense first
- [x] feat: By default, all participants should be selected in a new expense
- [x] bug: It's impossible to add decimals on iPhone (cannot input a comma)
- [x] bug: It's possible to open expenses from other projects by editing the url
- [x] bug: It's possible to add an expense without specifying a "paid_by"
- [x] feat: Ask which is current user + set a cookie
- [x] bug: Refunds are not pre-filled
- [x] feat: Copy link in tutorial before closing
- [x] feat: Move "Ajouter un participant" to Settings
- [x] feat: Add disconnect button from settings
- [x] bug: Values are sometimes displayed with many decimals (ex: 2.6666666666)
- [x] fix: Use userId instead of username for identification
- [x] bug: Refunds are taken into account in "total spent per member"

## Milestone F - Cleanup & UI
- [ ] chore: Write a README.md
- [ ] chore: Refactor the project
- [ ] feat: Make pages visually different (helps to navigate)
- [x] bug: Editing an expense shows too many digits
- [ ] feat: Show how much each expense impacts your balance
- [ ] feat: Add notifications for each added expense/refund
- [ ] feat: Highlight my information in balance.html
- [ ] feat: Rework UI with a Neobrutalism style
- [ ] feat: Validate user input (project, expense, member)

## Backlog

- [ ] chore: Add unit tests
- [ ] feat: Add "parts" to expenses
- [ ] feat: Admin script for listing projects and members
- [ ] feat: Integrate Lydia for refunds
- [ ] feat: Create an account with Webauthn (webauthn.io)