# alfred-repos

open local git repos in vscode

## Setup and Usage

1. Go to the releases page and download/import the workflow into alfred.
2. Edit the workflow and make sure the following variables are set:

* `REPOS_DIRECTORY`: Set to the root directory where you keep all of your local git repos.

3. Search for repos with `repo <search>`.
4. Select the repo to open with the following modifier keys:

* `Cmd + Enter` opens in vscode
* `Opt + Enter` opens in sublimetext
* `Ctrl + Enter` opens in iTerm
* `right arrow` opens universal actions menu

## FAQ

**Q: I get the error `“alfred-repos” cannot be opened because the developer cannot be verified.`.  How can i fix it?**

**A:** Thats an error gatekeeper returns when a binary isn't signed by an apple certificate.  To fix that follow these steps:

  1. Press cancel on the promp you received.
  2. Go to preferences and select `Security & Privacy`.  On the general tab, make sure "Allow apps download from: App Store and identified developers" is selected.  Beneath that you should see the text `alfred-repos was blocked from use because it is not from an identified developer`.  Press the `Allow Anyway` button and it try using the workflow again.
