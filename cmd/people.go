// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"math/rand"
	"strings"
	"time"

	"io/ioutil"

	"sync"

	"github.com/mchudgins/testDataGenerator/ssn"
	"github.com/spf13/cobra"
)

type user struct {
	LastName  string
	FirstName string
	SSN       string
}

// peopleCmd represents the people command
var peopleCmd = &cobra.Command{
	Use:   "people <quantity>",
	Short: "generate random people",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "%s\n\nError:  %s\n", cmd.UsageString(), err)
			os.Exit(2)
		}
		people(i)
	},
}

var (
	femaleNames []string
	lastNames   []string
	maleNames   []string
)

func init() {
	RootCmd.AddCommand(peopleCmd)

	var err error

	lastNames, err = loadNames("lastNames.orig")
	if err != nil {
		panic("Unable to load last names from file lastNames.orig")
	}
	femaleNames, err = loadNames("femaleNames")
	if err != nil {
		panic("Unable to load last names from file femaleNames")
	}
	maleNames, err = loadNames("maleNames")
	if err != nil {
		panic("Unable to load last names from file maleNames")
	}
}

func loadNames(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return lastNames, err
	}

	d := string(data)
	return strings.Split(d, "\n"), nil
}

func selectLastName(r *rand.Rand) string {
	return lastNames[r.Int31n(int32(len(lastNames)))]
}

func selectFirstName(r *rand.Rand) string {

	// assume ratio of males to females are uniform across all age cohorts
	maleOrFemale := r.Int31n(100)
	if maleOrFemale < 49 {
		return femaleNames[r.Int31n(int32(len(femaleNames)))]
	} else {
		return maleNames[r.Int31n(int32(len(maleNames)))]
	}

}

func genRandomPerson(r *rand.Rand) user {
	var u user

	u.LastName = selectLastName(r)
	u.FirstName = selectFirstName(r)
	u.SSN = ssn.GenerateSSN(r)

	return u
}

func genPeople(count int, ch chan user, wg *sync.WaitGroup) {
	defer wg.Done()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < count; i++ {
		user := genRandomPerson(r)
		ch <- user
	}
}

func people(count int) {
	const generators = 2
	var u user
	var wg sync.WaitGroup

	users := make(chan user)

	wg.Add(generators)
	go func() {
		wg.Wait()
		close(users)
	}()

	for i := 0; i < generators; i++ {
		go genPeople(count/generators, users, &wg)
	}

	for u = range users {
		fmt.Printf("%s, %s, %s\n", u.LastName, u.FirstName, u.SSN)
	}
}
