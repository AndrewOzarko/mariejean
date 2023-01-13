package cli

import (
	"encoding/json"
	"fmt"
	"github.com/andrewozarko/mariejean/internal/laravel"
	"github.com/andrewozarko/mariejean/pkg/cmdExec"
	"github.com/andrewozarko/mariejean/pkg/cmdRunner"
	"github.com/briandowns/spinner"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type CreateCommand struct {
}

func (cc *CreateCommand) Run(args cmdRunner.PreparedArgs) {
	if len(args.Args) < 3 {
		log.Fatalln("Create command require two parameters: TEMPLATE_NAME and PROJECT_NAME")
	}
	templateName := args.Args[1]
	projectName := args.Args[2]

	if _, ok := args.Options["force"]; ok {
		err := os.RemoveAll(projectName)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err := os.Stat(projectName); !os.IsNotExist(err) {
			log.Fatalf("Remove %v directory or use --force option.", projectName)
		}
	}

	switch templateName {
	case "laravel/laravel":
		LaravelInstall(projectName)
		break
	case "ozarko/mariejean":
		LaravelInstall(projectName)

		ComposerInstall("nwidart/laravel-modules", projectName)
		ConfigureLaravelModules(projectName)

		ComposerInstall("laravel/passport", projectName)
		ComposerInstall("league/fractal", projectName)
		ComposerInstall("prettus/l5-repository", projectName)
		ComposerInstall("spatie/laravel-fractal", projectName)

		ComposerInstall("--dev roave/security-advisories", projectName)
		ComposerInstall("--dev vimeo/psalm", projectName)

		RemoveMigrationsFromDefaultPath(projectName)
		RemoveUserModel(projectName)

		RestoreLaravelRoot(projectName)
		break
	default:
		log.Fatalf("%v is undefined template name", templateName)
		break
	}
}

func RemoveUserModel(projectName string) {
	// remove default User Model
	log.Println("remove default user models: ")
	if _, err := os.Stat("./" + projectName + "/app/Models"); err == nil {
		log.Println("/app/Models")
		err := os.RemoveAll("./" + projectName + "/app/Models")

		if err != nil {
			log.Println(err)
		}
	}
}

func RemoveMigrationsFromDefaultPath(projectName string) {
	// remove migration from default path
	log.Println("remove migration from default path: ")
	files, err := os.ReadDir("./" + projectName + "/database/migrations")
	if err != nil {
		log.Println(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), "create_users_table") ||
			strings.Contains(f.Name(), "create_password_resets_table") {
			log.Println(f.Name())
			err := os.Remove("./" + projectName + "/database/migrations/" + f.Name())
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func RestoreLaravelRoot(projectName string) {
	for _, a := range laravel.AssetNames() {
		log.Println(fmt.Sprintf("copy: %v", a))
		err := laravel.RestoreAsset(projectName, a)

		if err != nil {
			log.Println(err)
		}
	}
}

func ConfigureLaravelModules(projectName string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	s := spinner.New(spinner.CharSets[12], 100*time.Millisecond)
	defer log.Println(fmt.Sprintf("nwidart/laravel-modules configurated"))
	s.Start()
	go func() {
		defer wg.Done()
		jsonFile, err := os.Open(fmt.Sprintf("%v/composer.json", projectName))
		if err != nil {
			log.Println("Can't open composer.json")
			log.Fatalln(err)
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

		//modify composer json, autoload section
		result["autoload"].(map[string]interface{})["psr-4"].(map[string]interface{})["App\\Modules\\"] = "app/Modules/"
		log.Printf("%v/composer.json - modified", projectName)

		//save json
		file, _ := json.MarshalIndent(result, "", " ")
		_ = os.WriteFile(fmt.Sprintf("%v/composer.json", projectName), file, 0644)
		log.Printf("%v/composer.json - rewritten", projectName)

	}()
	wg.Wait()
	s.Stop()
}

func LaravelInstall(projectName string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	s := spinner.New(spinner.CharSets[12], 100*time.Millisecond)
	s.FinalMSG = "Laravel installed!\n"
	s.Start()
	go func() {
		defer wg.Done()
		out, errout, err := cmdExec.Shellout(fmt.Sprintf("composer create-project %v %v", "laravel/laravel", projectName))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(out, errout)
	}()
	wg.Wait()
	s.Stop()
}

func ComposerInstall(packageName string, projectName string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	s := spinner.New(spinner.CharSets[12], 100*time.Millisecond)
	defer log.Println(fmt.Sprintf("%v installed", packageName))
	s.Start()
	go func() {
		defer wg.Done()
		out, errout, err := cmdExec.Shellout(fmt.Sprintf("cd %v && composer require %v", projectName, packageName))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(out, errout)
	}()
	wg.Wait()
	s.Stop()
}
