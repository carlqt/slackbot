package slackbot

import (
	"fmt"
	"math/rand"
	"time"
)

var advice = [...]string{
	"To save money when cooking don't use olive oil, use the oil you would use for your car engine.",
	"During an interview ask how strict their sexual harassment policy is.",
	"To save time don't look before crossing the street.",
	"If you kick a police dog it gives you money.",
	"Girlfriend accuses you of checking out other women? Tell her you're not even attracted to thin women",
	"Feel unproductive at work? Cocaine.",
	"To test if a battery still works lick it.",
	"PEMDAS in math class means 'Please Excuse My Dope Ass Swag'.",
	"Real men don't wear pink. They wear crocs.",
	"Every good decision starts with a line of coke.",
	"It's not child abuse if you kill the child.",
	"Fight childhood obesity by beating up fat kids.",
	"Every zoo is a petting zoo.",
}

func BadAdvice() string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(advice))
	return fmt.Sprint(advice[n])
}
