# A golang TUI implementation of the popular word quiz Wordle!

## Future Updates

- [ ] Fix Tiny Bugs When Checking The Guess
- [ ] Make Portable GUI using Fyne (even for android and IOS)

# Example
```shell
!!!Welcome To Wordle!!!
You Have 6 Tries To Guess The Word

🟩 : Means Correct Letter In The Correct Spot

🟨 : Means Correct Letter In The Wrong Spot

⬛ : Means Letter Not Found In The Word

Enter A 5 Letter Word: sleep

🟨⬛🟨🟩⬛      (Try  1 /6):  sleep

Enter A 5 Letter Word: deads

⬛🟨⬛⬛🟨      (Try  2 /6):  deads

Enter A 5 Letter Word: argue

⬛🟨⬛🟨🟨      (Try  3 /6):  argue

Enter A 5 Letter Word: raged

🟨⬛⬛🟩⬛      (Try  4 /6):  raged

Enter A 5 Letter Word: urged

🟨🟨⬛🟩⬛      (Try  5 /6):  urged

Enter A 5 Letter Word: bibble

⬛⬛⬛⬛⬛      (Try  6 /6):  bibbl

Your Word Was  Muser : One who muses.

Want To Continue Playing (Y To Continue/ N To Stop): n

```