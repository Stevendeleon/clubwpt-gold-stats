# ClubWPT Gold Stats CLI

A command-line tool written in Go to help players on ClubWPT Gold calculate exactly how many consecutive hands they need to fold to reach a target VPIP percentage.

## Features
* Calculate the exact number of folds needed to drop your VPIP step-by-step.
* See how your PFR (Pre-Flop Raise) dynamically adjusts alongside your VPIP.
* Set custom target VPIP goals (defaults to 20%).

## Requirements

* Go

## How to use

* Clone Repo
* build or run the `cmd/cli/main.go` and pass the flags for your current total hands, vpip, and pfr
    * you can run `-h` to see args + shorthands 
