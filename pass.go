package main
 
import (
    "flag"
    "fmt"
    "github.com/jchavannes/go-pgp/pgp"
    "io/ioutil"
    "os"
)


func getParams() (
  bool, 
  bool,
  string,
  string, 
) {
  
	var name string        
	var password string   
	var help bool
  var list bool

	flag.StringVar(&name, "n", "", "Specify name to password")
	flag.StringVar(&password, "p", "", "Specify password")
	flag.BoolVar(&help, "h", false, "Show help")
  flag.BoolVar(&list, "l", false, "List saved passwords")

	flag.Usage = func() {
		fmt.Printf("Usage: \n")
    fmt.Println( "pass -n facebook -p 123456")
    fmt.Println( "pass -n facebook #get facebook password")
    fmt.Println( "pass -l #list all")
		flag.PrintDefaults()
	}   

	flag.Parse()

  showUsage := (help || name == "") && !list

	if showUsage {
		flag.Usage()
		return true, false, "", ""
	}
  return false, list, name, password
}

func getDefaultStoragePath()(string){
    return os.Getenv("PASS")
}

func saveToFile(filename, payload string){
    absoluteFilepath := fmt.Sprintf("%s/%s", getDefaultStoragePath(), filename)
    fmt.Println(absoluteFilepath)
    err := ioutil.WriteFile(absoluteFilepath, []byte(payload), 0755)
    if err != nil {
       panic("[!] error on saving to file")
    }
}

func getFileLocally(filename string)([]byte, error){
  absoluteFilepath := fmt.Sprintf("%s/%s", getDefaultStoragePath(), filename)
  return ioutil.ReadFile(absoluteFilepath)
}

func listPasswords() {
  files, err := ioutil.ReadDir(getDefaultStoragePath())
	if err != nil {
    fmt.Println("[!] could not generate new keys")
    return 
  }
  

  var blacklist = map[string]bool{
    "private.pgp": true,
    "public.pgp":    true,
  }

	for _, file := range files {
    filename := file.Name()
    ignorePassword := blacklist[filename]
    if ignorePassword {
      continue
    }
    passwordName := filename[:len(filename)-4]
		fmt.Println(passwordName)
	}
}

func getKeyPair(name string)(string, string){
    localPublicKey, err := getFileLocally("public.pgp")//ioutil.ReadFile("public.pgp")
    localPrivateKey, err := getFileLocally("private.pgp")

    if err == nil {
        return string(localPublicKey), string(localPrivateKey)
    }

    pgpKeyPair, err := pgp.GenerateKeyPair(name, name, name)

    if err == nil {
        saveToFile("public.pgp", string(pgpKeyPair.PublicKey))
        saveToFile("private.pgp", string(pgpKeyPair.PrivateKey))
        return string(pgpKeyPair.PublicKey), string(pgpKeyPair.PrivateKey)
    }
    panic("[!] could not generate new keys")
}

func savePassword(name, password string) (bool) {
    filename := fmt.Sprintf("%s.pgp", name)
    
    publicKey, privateKey := getKeyPair(name)

    pubEntity, err := pgp.GetEntity([]byte(publicKey), []byte{})

    keysErrors := publicKey == "" || privateKey == ""

	if err != nil || keysErrors {
		return false
    }

	encrypted, err := pgp.Encrypt(pubEntity, []byte(password))
	if err != nil {
        return false
	}

    saveToFile(filename, string(encrypted))
    
    return true
}

func getPassword(name string) (bool, string) {
    publicKey, privateKey := getKeyPair(name)

    privEntity, err := pgp.GetEntity([]byte(publicKey), []byte(privateKey))

    keysErrors := publicKey == "" || privateKey == ""



    filename := fmt.Sprintf("%s.pgp", name)

    encrypted, err := getFileLocally(filename)

    if err != nil || keysErrors {
		  return false, ""
    }
    decrypted, err := pgp.Decrypt(privEntity, encrypted)

    
    if err != nil {
		return false, ""
    }

    return true, string(decrypted)
}

func outputValue(password string)  {
    fmt.Println(password)
}

func main() {
  
  cancelExec, 
  list,
  name, 
  password := getParams()
  
  if cancelExec {
    return;
  }

  if list {
    listPasswords()
    return;
  }

  createNewPassword := password != ""
  
  if createNewPassword {
    success := savePassword(name, password);
    switch success {
      case true:
        fmt.Println("[+] Success on saving password")
      break
      case false:
        fmt.Println("[!] Error on saving password")
      break
    }
    return;
  }

  hasPassword,
  rawPassword := getPassword(name)

  if !hasPassword {
    fmt.Println("[info] no password match with name %s", name)
    return
  } 
  
  outputValue(rawPassword)
    
   
}