package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/stevendeleon/clubwpt-gold-stats/pkg/stats"
)

func main() {
	var hands, vpip, pfr, target float64

	// Hands (-hands, -n)
	flag.Float64Var(&hands, "hands", 0, "Total hands played")
	flag.Float64Var(&hands, "n", 0, "Total hands played (shorthand)")

	// VPIP (-vpip, -v)
	flag.Float64Var(&vpip, "vpip", 0.0, "Current VPIP percentage")
	flag.Float64Var(&vpip, "v", 0.0, "Current VPIP percentage (shorthand)")

	// PFR (-pfr, -p)
	flag.Float64Var(&pfr, "pfr", 0.0, "Current PFR percentage")
	flag.Float64Var(&pfr, "p", 0.0, "Current PFR percentage (shorthand)")

	// Target (-target, -t)
	flag.Float64Var(&target, "target", 20.0, "Target VPIP percentage to reach")
	flag.Float64Var(&target, "t", 20.0, "Target VPIP percentage to reach (shorthand)")

	flag.Usage = func() {
		fmt.Println("ClubWPT Gold Player percentage stats- Calculate how many folds you need to reach a target VPIP.")
		fmt.Println("\nUsage:")
		fmt.Println("./bin/main -n=1500 -v=28.5 -p=12.2 -t=18.0")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if hands <= 0 || vpip <= 0 {
		fmt.Println("Error: Please provide valid stats.")
		flag.Usage() 
		os.Exit(1)
	}

	results := stats.CalculateDrops(hands, vpip, pfr, target)

	if len(results) == 0 {
		fmt.Println("You are already at or below your target VPIP. Keep it up!")
		os.Exit(0)
	}

	// Print Header
	fmt.Println("===============================================================")
	fmt.Printf("CURRENT STATS: %.0f Hands | %.2f%% VPIP | %.2f%% PFR\n", hands, vpip, pfr)
	fmt.Printf("TARGET VPIP:   %.2f%%\n", target)
	fmt.Println("===============================================================")
	fmt.Printf("%-13s | %-15s | %-13s | %-13s\n", "Target VPIP %", "Folds Needed", "New Total", "Resulting PFR")
	fmt.Println("---------------------------------------------------------------")

	// Iterate over the returned results slice and print
	for _, r := range results {
		fmt.Printf("%-13.0f | %-15.0f | %-13.0f | %-13.2f%%\n", 
			r.TargetVPIP, 
			r.FoldsNeeded, 
			r.NewTotalHands, 
			r.ResultingPFR)
	}
	fmt.Println("===============================================================")
}
