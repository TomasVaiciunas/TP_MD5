# Anagram Challenge

## Algorithm
1. Remove all words form Dictionary that has a letter not in a anagram.
2. Remove all words from Dictionary that has more *same* letters a anagram.
3. Remove words from dictionary that are duplicated.
4. Add two words (with space) and see if letter match is *not more* then that letter count in anagram.
5. If #4 true, then add third word, see if length matches anagram (+ spaces), check if letter count is ok, hash it and compare hashes.

## How did it go?
Well, that was a rabbit hole for sure. I spent few full days on it but over all most of the time spent was with GoRoutines. Algorithm it self was quite easy and natural.
I never needed Goroutines before so that was a lesson learned by definitely not mastered. 

Though I have a code that finds Easy and Diffcult hashes under 30s on my i7 machine, I struggled to find elegant and effective solution that would pass possible left word combinations from three word phrase function to four word phrase function and of course do this using Goroutines. Hard phrase is - "wu lisp not statutory".

## Improvements

* Improve test coverage
* Move functions to another file
* Write a code that uses Channels to pass data between functions running that are searching for 3 and 4 word phrases using Goroutines.

## How to run it?
Make sure you have 'wordlist.dms' in same folder as main.go
Compile it using "*go build main.go*" and execute the binary. Or just run it using "*go run main.go*".

Output should look something like this but will warry depending on your machine:
```
Anagram has 18 letters
Found 1659 possible words in 101.451133ms. Moving on...
Found EASY hash: 'printout stout yawls' in 5.372153713s
Found DIFFICULT hash: ty outlaws printouts in 29.325829269s
This took 30.397019838s
```